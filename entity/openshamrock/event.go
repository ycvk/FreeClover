package openShamrockEntity

type GroupFileInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	BusID int64  `json:"busid"`
	URL   string `json:"url"`
}
type PrivateFileInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	URL    string `json:"url"`
	SubID  string `json:"sub_id"`
	Expire int64  `json:"expire"`
}
type OfflineFile struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	URL  string `json:"url"`
}
type PokeDetail struct {
	Action       string `json:"action"`
	Suffix       string `json:"suffix,omitempty"`
	ActionImgURL string `json:"action_img_url,omitempty"`
}

//type OpenShamrockEntity struct {
//	GroupFileInfo
//	PrivateFileInfo
//	OfflineFile
//	PokeDetail
//	SendMessage
//	MessageHistory
//	UserProfile
//	Common
//	LoginInfo
//	PhoneModel
//	Device
//	StrangerInfo
//	FriendList
//	UnidirectionalFriend
//	GroupDetail
//	GroupInfo
//	GroupList
//	UserInfo
//	GroupMemberInfo
//	GroupMemberList
//	HonorInfo
//	GroupHonorInfo
//	GroupSystemMsg
//	InvitedRequest
//	JoinRequest
//	FriendSystemMsg
//	GroupEssenceMsg
//	IsBlacklist
//	CommonMsg
//	PrivateMsg
//	GroupMsg
//	Sender
//	Msg
//	HistoryMsg
//	GroupHistoryMsg
//	ForwardMsg
//	FileInfo
//	EnableSendImage
//	OcrImage
//	AudioMsg
//	EnableSendAudio
//	GroupAvatar
//	GroupNotice
//	GroupProhibitedMember
//	SendPrivateFile
//	SendGroupFile
//	GroupFileSystemInfo
//	GroupRootFileList
//	File
//	Folder
//	GroupFiles
//	GroupFileUrl
//	UploadCache
//	DownloadCache
//	Battery
//	StartTime
//}
