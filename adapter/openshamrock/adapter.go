package openShamrockApi

import (
	"encoding/json"
	"github.com/oliverkirk-sudo/FreeClover/adapter"
	driver "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock"
	"github.com/oliverkirk-sudo/FreeClover/log"
	"github.com/oliverkirk-sudo/FreeClover/msg"
	openshamrockmsg "github.com/oliverkirk-sudo/FreeClover/msg/openshamrock"
)

type OpenShamrockApi struct {
	transceiver driver.Transceiver
	Url         string
	Api         adapter.Adapter
	Message     msg.MessageType
}

func NewOpenShamrockDriver(url string, authToken string, transceiver driver.Transceiver) *OpenShamrockApi {
	err := transceiver.Connect(url, authToken)
	if err != nil {
		log.Log.Warning("[WebSocket] 连接失败: ", err)
		return nil
	}
	apis := adapter.Adapter{
		Account:                Account{transceiver},
		Contacts:               Contacts{transceiver},
		User:                   User{transceiver},
		Message:                Message{transceiver},
		Resource:               Resource{transceiver},
		Process:                Process{transceiver},
		Group:                  Group{transceiver},
		File:                   File{transceiver},
		OpenShamrockSpecialApi: OpenShamrock{transceiver},
	}
	return &OpenShamrockApi{Url: url, Api: apis, transceiver: transceiver, Message: openshamrockmsg.Message{}}
}
func processJson[T interface{}](endpoint string, data []byte, transceiver driver.Transceiver) *T {
	log.Log.Debug("[ProcessJson] 当前端点：" + endpoint)
	var (
		a    T
		resp []byte
		err  error
	)
	if data != nil {
		resp, err = transceiver.SendJsonRequest(data, endpoint)
		if err != nil {
			log.Log.Warning("[ProcessJson] 解析返回值错误")
			return &a
		}
	} else {
		resp, err = transceiver.SendJsonRequest(nil, endpoint)
		if err != nil {
			log.Log.Warning("[ProcessJson] 解析返回值错误")
			return &a
		}
	}
	if resp != nil {
		err = json.Unmarshal(resp, &a)
		if err != nil {
			log.Log.Warning("[ProcessJson] json解析时出现问题: ", err)
			return &a
		}
		return &a
	}
	return &a
}
