package openShamrockEvent

import entity "github.com/ycvk/FreeClover/entity/openshamrock"

type FriendRecallEvent struct {
	BaseReport
	UserID     int64 `json:"user_id"`
	OperatorID int64 `json:"operator_id"`
	MessageID  int64 `json:"message_id"`
}

type GroupRecallEvent struct {
	BaseReport
	GroupID    int64 `json:"group_id"`
	UserID     int64 `json:"user_id"`
	OperatorID int64 `json:"operator_id"`
	MessageID  int64 `json:"message_id"`
}
type GroupIncreaseEvent struct {
	BaseReport
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
	SubType    string `json:"sub_type"`
}

type GroupDecreaseEvent struct {
	BaseReport
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
	SubType    string `json:"sub_type"`
}
type GroupAdminEvent struct {
	BaseReport
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	SubType string `json:"sub_type"`
}
type GroupUploadEvent struct {
	BaseReport
	GroupID int64                `json:"group_id"`
	UserID  int64                `json:"user_id"`
	File    entity.GroupFileInfo `json:"file"`
}
type PrivateUploadEvent struct {
	BaseReport
	UserID      int64                  `json:"user_id"`
	Sender      int64                  `json:"sender"`
	PrivateFile entity.PrivateFileInfo `json:"private_file"`
}
type GroupBanEvent struct {
	BaseReport
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
	Duration   int64  `json:"duration"`
	SubType    string `json:"sub_type"`
}
type GroupCardChangeEvent struct {
	BaseReport
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	CardNew string `json:"card_new"`
	CardOld string `json:"card_old"`
}
type FriendAddEvent struct {
	BaseReport
	UserID int64 `json:"user_id"`
}
type OfflineFileEvent struct {
	BaseReport
	UserID int64              `json:"user_id"`
	File   entity.OfflineFile `json:"file"`
}
type PokeEvent struct {
	BaseReport
	UserID     int64             `json:"user_id"`
	SenderID   int64             `json:"sender_id,omitempty"` // 仅私聊时存在
	GroupID    int64             `json:"group_id,omitempty"`  // 仅群聊时存在
	TargetID   int64             `json:"target_id"`
	PokeDetail entity.PokeDetail `json:"poke_detail"`
}
type EssenceEvent struct {
	BaseReport
	GroupID    int64  `json:"group_id"`
	SenderID   int64  `json:"sender_id"`
	OperatorID int64  `json:"operator_id"`
	MessageID  int64  `json:"message_id"`
	SubType    string `json:"sub_type"`
}
type GroupLuckyKingEvent struct {
	BaseReport
	GroupID  int64 `json:"group_id"`
	UserID   int64 `json:"user_id"`
	TargetID int64 `json:"target_id"`
}

type GroupHonorChangeEvent struct {
	BaseReport
	GroupID   int64  `json:"group_id"`
	UserID    int64  `json:"user_id"`
	HonorType string `json:"honor_type"`
}
type GroupTitleChangeEvent struct {
	BaseReport
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
}
type ClientStatusEvent struct {
	BaseReport
	Client interface{} `json:"client"`
	Online bool        `json:"online"`
}
