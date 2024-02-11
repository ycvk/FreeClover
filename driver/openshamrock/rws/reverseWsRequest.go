package openShamrockWs

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/oliverkirk-sudo/FreeClover/log"
	"github.com/oliverkirk-sudo/FreeClover/utils"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var upgrade = websocket.Upgrader{
	HandshakeTimeout: time.Second * 30,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
}

type ReverseWsDriver struct {
	url       string
	authToken string
	MsgQueue  utils.MsgQueue
}

func (r *ReverseWsDriver) Connect(url string, authToken string) error {
	log.Log.Debug("[ReverseWebSocket] 初始化连接")
	r.url = url
	r.authToken = authToken
	r.MsgQueue = utils.NewOriginMsgQueue()
	http.HandleFunc("/ws", r.handler)
	go func() {
		log.Log.Debug("[ReverseWebSocket] WebSocket服务器启动,地址为: ", url, "/ws")
		if strings.HasPrefix(url, "ws") {
			url = strings.Split(url, "//")[1]
		}
		log.Log.Fatal(http.ListenAndServe(url, nil))
	}()
	log.Log.Debug("[ReverseWebSocket] 启动了反向WebSocket")
	return nil
}

func (r *ReverseWsDriver) SendFormRequest(data url.Values, endpoint string) ([]byte, error) {
	log.Log.Warning("[ReverseWebSocket] 不支持发送消息")
	return nil, errors.New("[ReverseWebSocket] 不支持发送消息")
}
func (r *ReverseWsDriver) SendJsonRequest(data []byte, endpoint string) ([]byte, error) {
	log.Log.Warning("[ReverseWebSocket] 不支持发送消息")
	return nil, errors.New("[ReverseWebSocket] 不支持发送消息")
}
func (r *ReverseWsDriver) handler(w http.ResponseWriter, q *http.Request) {
	auth := q.Header.Get("Authorization")
	if auth != "bearer "+r.authToken {
		log.Log.Warning("[ReverseWebSocket] 身份验证失败")
		return
	}
	conn, err := upgrade.Upgrade(w, q, nil)
	if err != nil {
		return
	}
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return
		}
		r.MsgQueue.AddMessageInList(message)
	}
}
