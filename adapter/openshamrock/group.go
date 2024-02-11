package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/ycvk/FreeClover/driver/openshamrock"
	entity "github.com/ycvk/FreeClover/entity/openshamrock"
)

// Group 群组相关
type Group struct {
	transceiver adapter.Transceiver
}

// SetGroupName 该接口用于设置群名。
func (o Group) SetGroupName(groupId int64, groupName string) entity.Common {
	endpoint := "set_group_name"
	values, err := json.Marshal(map[string]interface{}{
		"group_id":   groupId,
		"group_name": groupName,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SetGroupAvatar 该接口用于设置群头像。
func (o Group) SetGroupAvatar(groupId int64, file string, cache int) entity.Common {
	endpoint := "set_group_portrait"
	panic("Not implemented: " + endpoint)
}

// SerGroupAdmin 该接口用于设置群管理员。
func (o Group) SerGroupAdmin(groupId int64, userId int64, enable bool) entity.Common {
	endpoint := "set_group_admin"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
		"enable":   enable,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SetGroupNote 该接口用于设置群备注。
func (o Group) SetGroupNote(groupId int64, userId int64, card string) entity.Common {
	endpoint := "set_group_card"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
		"card":     card,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SetGroupExclusiveTitle 该接口用于设置群组专属头衔。
func (o Group) SetGroupExclusiveTitle(groupId int64, userId int64, specialTitle string) entity.Common {
	endpoint := "set_group_special_title"
	values, err := json.Marshal(map[string]interface{}{
		"group_id":      groupId,
		"user_id":       userId,
		"special_title": specialTitle,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SetGroupSingleForbidSpeak 该接口用于群单人禁言。
func (o Group) SetGroupSingleForbidSpeak(groupId int64, userId int64, duration int64) entity.Common {
	endpoint := "set_group_ban"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
		"duration": duration,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SetGroupForbidSpeak 该接口用于群全员禁言。
func (o Group) SetGroupForbidSpeak(groupId int64, enable bool) entity.Common {
	endpoint := "set_group_whole_ban"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"enable":   enable,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SetGroupEssentialMessage 该接口用于设置精华消息。
func (o Group) SetGroupEssentialMessage(messageId int32) entity.Common {
	endpoint := "set_essence_msg"
	values, err := json.Marshal(map[string]interface{}{
		"message_id": messageId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// DeleteGroupEssentialMessage 该接口用于移出精华消息。
func (o Group) DeleteGroupEssentialMessage(messageId int32) entity.Common {
	endpoint := "delete_essence_msg"
	values, err := json.Marshal(map[string]interface{}{
		"message_id": messageId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SendGroupSign 该接口用于群打卡。
func (o Group) SendGroupSign(groupId int64) entity.Common {
	endpoint := "send_group_sign"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SendGroupNotice 该接口用于发送群公告。
func (o Group) SendGroupNotice(groupId int64, content string, image string) entity.Common {
	endpoint := "_send_group_notice"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"content":  content,
		"image":    image,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// GetGroupNotice 该接口用于获取群公告。
func (o Group) GetGroupNotice(groupId int64) entity.GroupNotice {
	endpoint := "_get_group_notice"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.GroupNotice{}
	}
	return *processJson[entity.GroupNotice](endpoint, values, o.transceiver)
}

// SetGroupKick 该接口用于群组踢人。
func (o Group) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) entity.Common {
	endpoint := "set_group_kick"
	values, err := json.Marshal(map[string]interface{}{
		"group_id":           groupId,
		"user_id":            userId,
		"reject_add_request": rejectAddRequest,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// SetGroupLeave 该接口用于退出群组。
func (o Group) SetGroupLeave(groupId int64) entity.Common {
	endpoint := "set_group_leave"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// GroupTouch 该接口用于群戳一戳。
func (o Group) GroupTouch(groupId int64, userId int64) entity.Common {
	endpoint := "group_touch"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// GetForbidSpeakList 用于获取群聊被禁言人列表。
func (o Group) GetForbidSpeakList(groupId int64) entity.GroupProhibitedMember {
	endpoint := "get_prohibited_member_list"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.GroupProhibitedMember{}
	}
	return *processJson[entity.GroupProhibitedMember](endpoint, values, o.transceiver)
}
