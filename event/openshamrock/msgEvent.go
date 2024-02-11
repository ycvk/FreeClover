package openShamrockEvent

import (
	entity "github.com/ycvk/FreeClover/entity/openshamrock"
)

type messageSender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card,omitempty"`
	Level    string `json:"level,omitempty"`
	Role     string `json:"role,omitempty"`
	Title    string `json:"title,omitempty"`
}
type MessageEvent struct {
	BaseReport
	MessageType string               `json:"message_type"`
	SubType     string               `json:"sub_type"`
	MessageID   int64                `json:"message_id"`
	UserID      int64                `json:"user_id"`
	Message     []entity.MessageItem `json:"message"`
	RawMessage  string               `json:"raw_message"`
	Font        int32                `json:"font"`
	Sender      messageSender        `json:"sender"`
	GroupID     int64                `json:"group_id"`
	TargetID    int64                `json:"target_id"`
	TempSource  int32                `json:"temp_source"`
	PeerID      int64                `json:"peer_id"`
}

type QuickReply struct {
	Reply      string
	AutoEscape bool
	AutoReply  bool
}

type QuickOperation struct {
	Reply       string
	AutoEscape  bool
	AtSender    bool
	Delete      bool
	Kick        bool
	Ban         bool
	BanDuration int64
}
