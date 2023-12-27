package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock"
	entity "github.com/oliverkirk-sudo/FreeClover/entity/openshamrock"
)

// Account 账户相关
type Account struct {
	transceiver adapter.Transceiver
}

// GetBotInfo 获取机器人信息
func (o Account) GetBotInfo() entity.LoginInfo {
	endpoint := "get_login_info"
	return *processJson[entity.LoginInfo](endpoint, nil, o.transceiver)
}

// SetBotInfo 设置机器人信息
func (o Account) SetBotInfo(profile entity.UserProfile) entity.Common {
	endpoint := "set_qq_profile"
	values, err := json.Marshal(profile)
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// GetOnlineType 该接口用于获取 QQ 在线机型。
func (o Account) GetOnlineType(model string) entity.PhoneModel {
	endpoint := "_get_model_show"
	values, err := json.Marshal(map[string]interface{}{
		"model": model,
	})
	if err != nil {
		return entity.PhoneModel{}
	}
	return *processJson[entity.PhoneModel](endpoint, values, o.transceiver)
}

// SetOnlineType 该接口用于设置 QQ 在线机型。
func (o Account) SetOnlineType(model, manu string) entity.Common {
	endpoint := "_set_model_show"
	values, err := json.Marshal(map[string]interface{}{
		"model": model,
		"manu":  manu,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// GetOnlineClientList 该接口用于获取当前账号在线客户端列表。
func (o Account) GetOnlineClientList(noCache bool) entity.Device {
	endpoint := "get_online_clients"
	panic("Not implemented: " + endpoint)
}
