package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/ycvk/FreeClover/driver/openshamrock"
	entity "github.com/ycvk/FreeClover/entity/openshamrock"
)

// Process 处理相关
type Process struct {
	transceiver adapter.Transceiver
}

// AddFriendRequest 该接口用于处理加好友请求。
func (o Process) AddFriendRequest(flag string, approve bool, remark string) entity.Common {
	endpoint := "set_friend_add_request"
	values, err := json.Marshal(map[string]interface{}{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// AddGroupRequest 该接口用于处理加群请求／邀请。
func (o Process) AddGroupRequest(flag string, subType string, approve bool, reason string) entity.Common {
	endpoint := "set_group_add_request"
	values, err := json.Marshal(map[string]interface{}{
		"flag":     flag,
		"approve":  approve,
		"sub_type": subType,
		"reason":   reason,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}
