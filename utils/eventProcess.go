package utils

import (
	"encoding/json"
	event "github.com/oliverkirk-sudo/FreeClover/event/openshamrock"
	"github.com/oliverkirk-sudo/FreeClover/log"
)

type MsgQueue struct {
	MsgList chan map[string]interface{}
}

func (m *MsgQueue) AddMessageInList(message []byte) {
	var heartBeat event.HeartBeat
	var data map[string]interface{}
	json.Unmarshal(message, &heartBeat)
	json.Unmarshal(message, &data)
	if heartBeat.MetaEventType == "heartbeat" {
		log.Log.Debug("接收到心跳包")
		return
	} else if heartBeat.MetaEventType == "lifecycle" {
		log.Log.Debug("进入生命循环")
		return
	}
	if data["post_type"] != nil {
		log.Log.Debug("接收到事件：", data["post_type"])
	}
	if data["status"] != nil {
		log.Log.Debug("接收到请求返回：", data["echo"])
	}
	m.MsgList <- data
}

func (m *MsgQueue) GetMessageItemBlock() map[string]interface{} {
	return <-m.MsgList
}
func (m *MsgQueue) GetMessageItem() map[string]interface{} {
	select {
	case t := <-m.MsgList:
		return t
	default:
		return map[string]interface{}{}
	}
}
func (m *MsgQueue) IsEmpty() bool {
	return len(m.MsgList) == 0
}
func (m *MsgQueue) eventDetail(data map[string]interface{}) {
	postType := data["post_type"].(string)
	println(postType)
	switch postType {
	case event.ReceiveMsgPostType:
		messageType := data["message_type"].(string)
		switch messageType {
		case event.PrivateMessageType:
		case event.GroupMessageType:
		default:
		}
	case event.SentMsgPostType:
		messageType := data["message_type"].(string)
		switch messageType {
		case event.PrivateMessageType:
		case event.GroupMessageType:
		default:
		}
	case event.NoticePostType:
		messageType := data["notice_type"].(string)
		switch messageType {
		case event.FriendRecallNoticeType:
		case event.GroupRecallNoticeType:
		case event.GroupIncreaseNoticeType:
		case event.GroupDecreaseNoticeType:
		case event.GroupAdminNoticeType:
		case event.GroupUploadNoticeType:
		case event.PrivateUploadNoticeType:
		case event.GroupBanNoticeType:
		case event.GroupCardNoticeType:
		case event.FriendAddNoticeType:
		case event.OfflineFileNoticeType:
		case event.EssenceNoticeType:
		case event.ClientStatusNoticeType:
		case event.NotifyNoticeType:
			subType := data["sub_type"].(string)
			switch subType {
			case event.PokeNoticeNotifySubType:
			case event.LuckyKingNoticeNotifySubType:
			case event.HonorNoticeNotifySubType:
			case event.TitleNoticeNotifySubType:
			default:
			}
		default:
		}
	case event.RequestPostType:
		requestType := data["request_type"].(string)
		switch requestType {
		case event.FriendRequestType:

		case event.GroupRequestType:
		}
	default:
	}
}
