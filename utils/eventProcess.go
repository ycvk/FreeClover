package utils

import (
	"encoding/json"
	event "github.com/ycvk/FreeClover/event/openshamrock"
	"github.com/ycvk/FreeClover/log"
)

// MsgQueue 消息通道接口
type MsgQueue interface {
	AddMessageInList(message []byte)
	GetMessageItemBlock() map[string]interface{}
	GetMessageItem() map[string]interface{}
	IsEmpty() bool
}

// OriginMsgQueue  消息通道
type OriginMsgQueue struct {
	MsgList chan map[string]interface{}
}

func NewOriginMsgQueue() MsgQueue {
	return &OriginMsgQueue{
		MsgList: make(chan map[string]interface{}, 256),
	}
}

// AddMessageInList 添加消息到消息通道中
func (m *OriginMsgQueue) AddMessageInList(message []byte) {
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

// GetMessageItemBlock 以阻塞的方式获取消息通道中的内容，队列为空时阻塞
func (m *OriginMsgQueue) GetMessageItemBlock() map[string]interface{} {
	return <-m.MsgList
}

// GetMessageItem 以非阻塞的方式获取消息通道中的内容，队列为空时返回空对象
func (m *OriginMsgQueue) GetMessageItem() map[string]interface{} {
	select {
	case t := <-m.MsgList:
		return t
	default:
		return map[string]interface{}{}
	}
}

// IsEmpty 判断队列是否为空
func (m *OriginMsgQueue) IsEmpty() bool {
	return len(m.MsgList) == 0
}
func (m *OriginMsgQueue) eventDetail(data map[string]interface{}) {
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
