package openShamrockHttp

import (
	"bytes"
	http "github.com/bogdanfinn/fhttp"
	"github.com/oliverkirk-sudo/FreeClover/log"
	"io"
	"net/url"
	"strings"
)

type HttpDriver struct {
	url       string
	authToken string
	client    http.Client
}

func (h *HttpDriver) Connect(url string, authToken string) error {
	log.Log.Debug("[HTTP] 初始化连接")
	h.url = url
	h.authToken = authToken
	h.client = http.Client{}
	log.Log.Debug("[HTTP] 启动了主动HTTP")
	return nil
}

func (h *HttpDriver) SendFormRequest(data url.Values, endpoint string) ([]byte, error) {
	log.Log.Debug("[HTTP] 发送了form数据")
	var response *http.Response
	var err error
	if data != nil {
		req, err := http.NewRequest("POST", h.url+"/"+endpoint, strings.NewReader(data.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", "Bearer "+h.authToken)
		response, err = h.client.Do(req)
		if err != nil {
			return nil, err
		}
	} else {
		req, err := http.NewRequest("POST", h.url+"/"+endpoint, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer "+h.authToken)
		response, err = h.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
func (h *HttpDriver) SendJsonRequest(data []byte, endpoint string) ([]byte, error) {
	log.Log.Debug("[HTTP] 发送了json数据")
	var response *http.Response
	var err error
	if data != nil {
		req, err := http.NewRequest("POST", h.url+"/"+endpoint, bytes.NewBuffer(data))
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+h.authToken)
		response, err = h.client.Do(req)
		if err != nil {
			return nil, err
		}
	} else {
		req, err := http.NewRequest("POST", h.url+"/"+endpoint, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer "+h.authToken)
		response, err = h.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
