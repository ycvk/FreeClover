package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock"
	entity "github.com/oliverkirk-sudo/FreeClover/entity/openshamrock"
)

// User 用户相关
type User struct {
	transceiver adapter.Transceiver
}

// DeleteFriend 该接口用于删除好友。
func (o User) DeleteFriend(userId int64) entity.Common {
	endpoint := "delete_friend"
	panic("Not implemented: " + endpoint)
}

// DeleteUnidirectionFriend 该接口用于删除单向好友。
func (o User) DeleteUnidirectionFriend(userId int64) entity.Common {
	endpoint := "delete_unidirectional_friend"
	panic("Not implemented: " + endpoint)
}

// RevokeMsg 该接口用于撤回消息。
func (o User) RevokeMsg(messageId int32) entity.Common {
	endpoint := "delete_msg"
	values, err := json.Marshal(map[string]interface{}{
		"message_id": messageId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}
