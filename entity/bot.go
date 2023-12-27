package entity

type BotInfo struct {
	UserId     int64
	NickName   string
	FriendList []struct {
		UserId          int64  `json:"user_id"`
		UserName        string `json:"user_name"`
		UserDisplayname string `json:"user_displayname"`
		UserRemark      string `json:"user_remark"`
		Age             int    `json:"age"`
		Gender          int    `json:"gender"`
		GroupId         int    `json:"group_id"`
		Platform        string `json:"platform"`
		TermType        int    `json:"term_type"`
	}
	GroupList []struct {
		GroupId        int    `json:"group_id"`
		GroupName      string `json:"group_name"`
		GroupRemark    string `json:"group_remark"`
		GroupUin       int    `json:"group_uin"`
		Admins         []int  `json:"admins"`
		ClassText      string `json:"class_text"`
		IsFrozen       bool   `json:"is_frozen"`
		MaxMember      int    `json:"max_member"`
		MemberNum      int    `json:"member_num"`
		MemberCount    int    `json:"member_count"`
		MaxMemberCount int    `json:"max_member_count"`
	}
}
