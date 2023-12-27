package openShamrockEvent

type QuickAddFriend struct {
	Approve bool
	Remark  string
}

type QuickAddGroup struct {
	Approve bool
	Reason  string
}
type FriendRequest struct {
	BaseReport
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
}
type GroupRequest struct {
	BaseReport
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
	SubType string `json:"sub_type"`
}
