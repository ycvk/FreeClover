package openShamrockEntity

type SendMessage struct {
	MessageType string        `json:"message_type"`
	UserID      int64         `json:"user_id"`
	GroupID     int64         `json:"group_id"`
	DiscussID   int64         `json:"discuss_id"`
	Message     []MessageItem `json:"message"`
	AutoEscape  bool          `json:"auto_escape,omitempty"`
}
type MessageHistory struct {
	MessageType string `json:"message_type"`
	UserID      int64  `json:"user_id,omitempty"`
	GroupID     int64  `json:"group_id,omitempty"`
	Count       int32  `json:"count,omitempty"`
	MessageSeq  int32  `json:"message_seq,omitempty"`
}
type UserProfile struct {
	Nickname     string `json:"nickname"`
	Company      string `json:"company"`
	Email        string `json:"email"`
	College      string `json:"college"`
	PersonalNote string `json:"personal_note"`
	Age          int32  `json:"age,omitempty"`
	Birthday     string `json:"birthday,omitempty"`
}
type SendPrivateMessage struct {
	UserId     int64         `json:"user_id"`
	Message    []MessageItem `json:"message"`
	AutoEscape bool          `json:"auto_escape"`
}
type SendGroupMessage struct {
	GroupId    int64         `json:"group_id"`
	Message    []MessageItem `json:"message"`
	AutoEscape bool          `json:"auto_escape"`
}
type SendGroupForwardMessage struct {
	GroupId  int64                 `json:"group_id"`
	Messages []GroupForwardMessage `json:"messages"`
}
type GroupForwardMessage struct {
	Type string                  `json:"type"`
	Data GroupForwardMessageData `json:"data"`
}
type GroupForwardMessageData struct {
	Content []MessageItem `json:"content"`
}
type SendPrivateForwardMessage struct {
	UserId  int64         `json:"user_id"`
	Message []MessageItem `json:"message"`
}
