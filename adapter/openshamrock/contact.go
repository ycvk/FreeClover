package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock"
	entity "github.com/oliverkirk-sudo/FreeClover/entity/openshamrock"
)

// Contacts 联系人相关
type Contacts struct {
	transceiver adapter.Transceiver
}

// GetStrangerInfo 该接口用于获取陌生人的信息。
func (o Contacts) GetStrangerInfo(userId int64) entity.StrangerInfo {
	endpoint := "get_stranger_info"
	values, err := json.Marshal(map[string]interface{}{
		"user_id": userId,
	})
	if err != nil {
		return entity.StrangerInfo{}
	}
	return *processJson[entity.StrangerInfo](endpoint, values, o.transceiver)
}

// GetFriendsList 该接口用于获取好友列表。
func (o Contacts) GetFriendsList(refresh bool) entity.FriendList {
	endpoint := "get_friend_list"
	values, err := json.Marshal(map[string]interface{}{
		"refresh": refresh,
	})
	if err != nil {
		return entity.FriendList{}
	}
	return *processJson[entity.FriendList](endpoint, values, o.transceiver)
}

// GetUnidirectionFriendsList 该接口用于获取单向好友列表。
func (o Contacts) GetUnidirectionFriendsList() entity.UnidirectionalFriend {
	endpoint := "get_unidirectional_friend_list"
	panic("Not implemented: " + endpoint)
}

// GetGroupInfo 该接口用于获取群信息。
func (o Contacts) GetGroupInfo(groupId int64) entity.GroupInfo {
	endpoint := "get_group_info"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.GroupInfo{}
	}
	return *processJson[entity.GroupInfo](endpoint, values, o.transceiver)
}

// GetGroupsList 该接口用于获取群列表。
func (o Contacts) GetGroupsList() entity.GroupList {
	endpoint := "get_group_list"
	return *processJson[entity.GroupList](endpoint, nil, o.transceiver)
}

// GetGroupMemberInfo 该接口用于获取群成员信息。
func (o Contacts) GetGroupMemberInfo(groupId, userId int64) entity.GroupMemberInfo {
	endpoint := "get_group_member_info"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
	})
	if err != nil {
		return entity.GroupMemberInfo{}
	}
	return *processJson[entity.GroupMemberInfo](endpoint, values, o.transceiver)
}

// GetGroupMemberList 该接口用于获取群成员列表。
func (o Contacts) GetGroupMemberList(groupId int64) entity.GroupMemberList {
	endpoint := "get_group_member_list"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.GroupMemberList{}
	}
	return *processJson[entity.GroupMemberList](endpoint, values, o.transceiver)
}

// GetGroupHonorInfo 该接口用于获取群荣誉信息。
func (o Contacts) GetGroupHonorInfo(groupId int64) entity.GroupHonorInfo {
	endpoint := "get_group_honor_info"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.GroupHonorInfo{}
	}
	return *processJson[entity.GroupHonorInfo](endpoint, values, o.transceiver)
}

// GetGroupSystemInfo 该接口用于获取群系统消息。
func (o Contacts) GetGroupSystemInfo() entity.GroupFileSystemInfo {
	endpoint := "get_group_system_msg"
	return *processJson[entity.GroupFileSystemInfo](endpoint, nil, o.transceiver)
}

// GetFriendSystemInfo 该接口用于获取好友系统消息。
func (o Contacts) GetFriendSystemInfo() entity.FriendSystemMsg {
	endpoint := "get_friend_system_msg"
	return *processJson[entity.FriendSystemMsg](endpoint, nil, o.transceiver)
}

// GetGroupEssentialMessageList 该接口用于获取精华消息列表。
func (o Contacts) GetGroupEssentialMessageList(groupId int64) entity.GroupEssenceMsg {
	endpoint := "get_essence_msg_list"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.GroupEssenceMsg{}
	}
	return *processJson[entity.GroupEssenceMsg](endpoint, values, o.transceiver)
}

// CheckInBlacklist 获取好友/陌生人是否处于黑名单列表。
func (o Contacts) CheckInBlacklist(userId int64) entity.IsBlacklist {
	endpoint := "is_blacklist_uin"
	values, err := json.Marshal(map[string]interface{}{
		"user_id": userId,
	})
	if err != nil {
		return entity.IsBlacklist{}
	}
	return *processJson[entity.IsBlacklist](endpoint, values, o.transceiver)
}
