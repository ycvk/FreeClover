package openShamrockHttp

import (
	"errors"
	"github.com/ycvk/FreeClover/log"
	"github.com/ycvk/FreeClover/utils"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpWebHookDriver struct {
	url       string
	authToken string
	MsgQueue  utils.MsgQueue
}

func (h *HttpWebHookDriver) Connect(url string, authToken string) error {
	log.Log.Debug("[WebHook] 初始化连接")
	h.MsgQueue = utils.NewOriginMsgQueue()
	h.url = url
	h.authToken = authToken
	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Log.Debug("Error reading request body:", err)
			return
		}
		h.MsgQueue.AddMessageInList(body)
	})
	go func() {
		if strings.HasPrefix(url, "http") {
			url = strings.Split(url, "//")[1]
		}
		log.Log.Debug("[WebHook] 启动服务器，监听地址:", url, "/hook")
		err := http.ListenAndServe(url, nil)
		if err != nil {
			log.Log.Fatal("[WebHook] 服务器启动失败:", err)
		}
	}()
	log.Log.Debug("[WebHook] 启动了回调http,服务地址为：", url)
	return nil
}

func (h *HttpWebHookDriver) SendFormRequest(data url.Values, endpoint string) ([]byte, error) {
	log.Log.Warning("[WebHook] 不支持发送消息")
	return nil, errors.New("[WebHook] 不支持发送消息")
}
func (h *HttpWebHookDriver) SendJsonRequest(data []byte, endpoint string) ([]byte, error) {
	log.Log.Warning("[WebHook] 不支持发送消息")
	return nil, errors.New("[WebHook] 不支持发送消息")
}

func (h *HttpWebHookDriver) SendFileRequest(data []byte, endpoint string) ([]byte, error) {
	log.Log.Warning("[WebHook] 不支持发送消息")
	return nil, errors.New("[WebHook] 不支持发送消息")
}
