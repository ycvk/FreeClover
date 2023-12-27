package openShamrockEvent

const (
	ReceiveMsgPostType string = "message"
	SentMsgPostType    string = "message_sent"
	NoticePostType     string = "notice"
	RequestPostType    string = "request"
	MetaPostType       string = "meta_event"

	PrivateMessageType string = "private"
	GroupMessageType   string = "group"

	FriendMessageSubType string = "friend"
	NormalMessageSubType string = "normal"
	GroupMessageSubType  string = "group"

	GroupUploadNoticeType   string = "group_upload"
	GroupAdminNoticeType    string = "group_admin"
	GroupDecreaseNoticeType string = "group_decrease"
	GroupIncreaseNoticeType string = "group_increase"
	GroupBanNoticeType      string = "group_ban"
	GroupRecallNoticeType   string = "group_recall"
	GroupCardNoticeType     string = "group_card"
	PrivateUploadNoticeType string = "private_upload"
	FriendAddNoticeType     string = "friend_add"
	FriendRecallNoticeType  string = "friend_recall"
	OfflineFileNoticeType   string = "offline_file"
	ClientStatusNoticeType  string = "client_status"
	EssenceNoticeType       string = "essence"
	NotifyNoticeType        string = "notify"

	HonorNoticeNotifySubType     string = "honor"
	PokeNoticeNotifySubType      string = "poke"
	LuckyKingNoticeNotifySubType string = "lucky_king"
	TitleNoticeNotifySubType     string = "title"
	ConnectSubType               string = "connect"

	FriendRequestType string = "friend"
	GroupRequestType  string = "group"
)

type BaseReport struct {
	Time     int64  `json:"time"`
	SelfId   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}
type HeartBeat struct {
	Time          int    `json:"time"`
	SelfId        int64  `json:"self_id"`
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
	Status        struct {
		Self struct {
			Platform string `json:"platform"`
			UserId   int64  `json:"user_id"`
		} `json:"self"`
		Online   bool   `json:"online"`
		Good     bool   `json:"good"`
		QqStatus string `json:"qq.status"`
	} `json:"status"`
	Interval int `json:"interval"`
}
