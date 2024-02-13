package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/ycvk/FreeClover/driver/openshamrock"
	entity "github.com/ycvk/FreeClover/entity/openshamrock"
)

// Message 消息相关
type Message struct {
	transceiver adapter.Transceiver
}

// SendPrivateMessage 该接口用于发送私聊消息。
func (o Message) SendPrivateMessage(userId int64, autoEscape bool, message ...entity.MessageItem) entity.PrivateMsg {
	endpoint := "send_private_msg"
	values, err := json.Marshal(entity.SendPrivateMessage{
		UserId:     userId,
		Message:    message,
		AutoEscape: autoEscape,
	})
	if err != nil {
		return entity.PrivateMsg{}
	}
	return *processJson[entity.PrivateMsg](endpoint, values, o.transceiver)
}

// SendGroupMessage 该接口用于发送群聊消息。
func (o Message) SendGroupMessage(groupId int64, autoEscape bool, message ...entity.MessageItem) entity.GroupMsg {
	endpoint := "send_group_msg"
	values, err := json.Marshal(entity.SendGroupMessage{
		GroupId:    groupId,
		Message:    message,
		AutoEscape: autoEscape,
	})
	if err != nil {
		return entity.GroupMsg{}
	}
	return *processJson[entity.GroupMsg](endpoint, values, o.transceiver)
}

// SendMessage 该接口用于发送消息。
func (o Message) SendMessage(messageType string, userId int64, groupId int64, discussId int64, autoEscape bool, message []entity.MessageItem) entity.SendMessage {
	endpoint := "send_msg"
	values, err := json.Marshal(entity.SendMessage{
		MessageType: messageType,
		UserID:      userId,
		GroupID:     groupId,
		DiscussID:   discussId,
		Message:     message,
		AutoEscape:  autoEscape,
	})
	if err != nil {
		return entity.SendMessage{}
	}
	return *processJson[entity.SendMessage](endpoint, values, o.transceiver)
}

// GetMessage 该接口用于获取消息。
func (o Message) GetMessage(messageId int32) entity.Msg {
	endpoint := "get_msg"
	values, err := json.Marshal(map[string]interface{}{
		"message_id": messageId,
	})
	if err != nil {
		return entity.Msg{}
	}
	return *processJson[entity.Msg](endpoint, values, o.transceiver)
}

// GetHistoryMessage 该接口用于获取历史消息。
func (o Message) GetHistoryMessage(messageType string, userId int64, groupId int64, count int32, messageSeq int32) entity.MessageHistory {
	endpoint := "get_history_msg"
	values, err := json.Marshal(entity.MessageHistory{
		MessageType: messageType,
		UserID:      userId,
		GroupID:     groupId,
		Count:       count,
		MessageSeq:  messageSeq,
	})
	if err != nil {
		return entity.MessageHistory{}
	}
	return *processJson[entity.MessageHistory](endpoint, values, o.transceiver)
}

// GetGroupHistoryMessage 该接口用于获取群聊历史消息。
func (o Message) GetGroupHistoryMessage(groupId int64, count int32, messageSeq int32) entity.GroupHistoryMsg {
	endpoint := "get_group_msg_history"
	values, err := json.Marshal(map[string]interface{}{
		"group_id":    groupId,
		"count":       count,
		"message_seq": messageSeq,
	})
	if err != nil {
		return entity.GroupHistoryMsg{}
	}
	return *processJson[entity.GroupHistoryMsg](endpoint, values, o.transceiver)
}

// DeleteLocalMessageCache 该接口用于清除本地消息缓存。
func (o Message) DeleteLocalMessageCache(messageType string, userId int64, groupId int64) entity.Common {
	endpoint := "clear_msgs"
	values, err := json.Marshal(map[string]interface{}{
		"group_id":     groupId,
		"user_id":      userId,
		"message_type": messageType,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// GetForwardMessage 该接口用于获取合并转发内容。
func (o Message) GetForwardMessage(id string) entity.ForwardMsg {
	endpoint := "get_forward_msg"
	values, err := json.Marshal(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return entity.ForwardMsg{}
	}
	return *processJson[entity.ForwardMsg](endpoint, values, o.transceiver)
}

// SendGroupForwardMessage 该接口用于发送群聊合并转发。
func (o Message) SendGroupForwardMessage(groupId int64, messages ...entity.GroupForwardMessage) entity.Common {
	endpoint := "send_group_forward_msg"
	values, err := json.Marshal(entity.SendGroupForwardMessage{
		GroupId:  groupId,
		Messages: messages,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SendPrivateForwardMessage 该接口用于发送私聊合并转发。
func (o Message) SendPrivateForwardMessage(userId int64, message ...entity.MessageItem) entity.Common {
	endpoint := "send_private_forward_msg"
	values, err := json.Marshal(entity.SendPrivateForwardMessage{
		UserId:  userId,
		Message: message,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}
