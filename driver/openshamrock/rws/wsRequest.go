package openShamrockWs

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/oliverkirk-sudo/FreeClover/log"
	"github.com/oliverkirk-sudo/FreeClover/utils"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
)

type WsDriver struct {
	url       string
	authToken string
	Instance  *websocket.Conn
	MsgQueue  utils.MsgQueue
}

func (w *WsDriver) Connect(url string, authToken string) error {
	log.Log.Debug("[WebSocket] 初始化连接")
	var err error
	w.url = url
	w.authToken = authToken
	w.MsgQueue = utils.MsgQueue{MsgList: make(chan map[string]interface{}, 256)}
	w.Instance, err = w.createWebSocketConnection()
	log.Log.Debug("[WebSocket] 启动了WebSocket连接")
	w.pollingEvents()
	return err
}
func (w *WsDriver) createWebSocketConnection() (*websocket.Conn, error) {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+w.authToken)
	conn, _, err := websocket.DefaultDialer.Dial(w.url, headers)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func (w *WsDriver) SendFormRequest(data url.Values, endpoint string) ([]byte, error) {
	body := struct {
		Action string      `json:"action"`
		Params interface{} `json:"params"`
		Echo   string      `json:"echo"`
	}{
		Action: endpoint,
		Params: data,
		Echo:   "",
	}
	err := w.Instance.WriteJSON(body)
	if err != nil {
		return nil, err
	}
	log.Log.Debug("[WebSocket] Form请求发送成功")
	if strings.HasSuffix(w.url, "adapter") {
		_, message, err := w.Instance.ReadMessage()
		if err != nil {
			return nil, err
		}
		return message, nil
	}
	return nil, nil
}
func (w *WsDriver) SendJsonRequest(data []byte, endpoint string) ([]byte, error) {
	log.Log.Debug("[WebSocket] 发送json消息")
	var originData map[string]interface{}
	newData := make(map[string]interface{})
	newData["action"] = endpoint
	newData["echo"] = endpoint
	if data != nil {
		err := json.Unmarshal(data, &originData)
		if err != nil {
			return nil, err
		}
		newData["params"] = originData

	} else {
		newData["params"] = struct{}{}
	}
	if err := w.Instance.WriteJSON(newData); err != nil {
		return nil, err
	}
	log.Log.Debug("[WebSocket] Json请求发送成功")
	if strings.HasSuffix(w.url, "api") {
		_, message, err := w.Instance.ReadMessage()
		if err != nil {
			return nil, err
		}
		return message, nil
	}
	return nil, nil
}
func (w *WsDriver) reConnect() *websocket.Conn {
	connection, err := w.createWebSocketConnection()
	if err != nil {
		log.Log.Warning("[WebSocket] 重连失败: ", err)
		return nil
	}
	return connection
}
func (w *WsDriver) pollingEvents() {
	if strings.HasSuffix(w.url, "api") {
		log.Log.Debug("[WebSocket] api接口无需轮询")
		return
	}
	log.Log.Debug("[WebSocket] 开启轮询")
	interrupt := make(chan os.Signal)
	go func() {
		for {
			select {
			case <-interrupt:
				log.Log.Debug("[WebSocket] 接收到中断信号，关闭轮询")
				w.Instance.Close()
				runtime.Goexit()
			case <-time.After(time.Duration(1) * time.Second):
				done := make(chan int)
				messageBody := make(chan []byte)
				go w.receiveHandler(messageBody, done)
				select {
				case <-time.After(time.Duration(30) * time.Second):
					log.Log.Debug("[WebSocket] 心跳包超时，尝试重连")
					conn := w.reConnect()
					for conn == nil {
						conn = w.reConnect()
					}
					w.Instance = conn
					log.Log.Debug("[WebSocket] 重新连接成功")
				case <-done:
					w.MsgQueue.AddMessageInList(<-messageBody)
				}
			}
		}
	}()
}

func (w *WsDriver) receiveHandler(m chan<- []byte, ch chan<- int) {
	_, message, err := w.Instance.ReadMessage()
	if err != nil {
		log.Log.Warning("[WebSocket] read: ", err)
	} else {
		ch <- 1
		m <- message
	}
}
