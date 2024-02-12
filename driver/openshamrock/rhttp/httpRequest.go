package openShamrockHttp

import (
	"bytes"
	http "github.com/bogdanfinn/fhttp"
	"github.com/ycvk/FreeClover/log"
	"io"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
	"time"
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
	var response *http.Response
	var err error
	if data != nil {
		log.Log.Debug("[HTTP] 发送了form数据")
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
	var response *http.Response
	var err error
	if data != nil {
		log.Log.Debug("[HTTP] 发送了json数据")
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

func (h *HttpDriver) SendFileRequest(data []byte, endpoint string) ([]byte, error) {
	var response *http.Response
	var err error
	if data != nil {
		log.Log.Debug("[HTTP] 发送了文件")
		buffer := bytes.NewBuffer(data)
		multiPartWriter := multipart.NewWriter(buffer)
		// 创建一个新的表单文件字段
		fileWriter, err := multiPartWriter.CreateFormFile("file", strconv.FormatInt(time.Now().UnixMilli(), 10))
		if err != nil {
			return nil, err
		}
		// 将文件内容拷贝到表单文件字段
		_, err = io.Copy(fileWriter, buffer)
		if err != nil {
			return nil, err
		}
		// 关闭多部分写入器
		multiPartWriter.Close()
		// 创建一个新的请求
		req, err := http.NewRequest("POST", h.url+"/"+endpoint, buffer)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
		req.Header.Set("Authorization", "Bearer "+h.authToken)
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
