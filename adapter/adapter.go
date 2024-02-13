package adapter

import entity "github.com/ycvk/FreeClover/entity/openshamrock"

// Adapter 适配器结构体
type Adapter struct {
	Account                account
	Contacts               contacts
	User                   user
	Message                message
	Resource               resource
	Process                process
	Group                  group
	File                   file
	OpenShamrockSpecialApi openShamrockSpecialApi
}

type account interface { //账户相关
	GetBotInfo() entity.LoginInfo                        //该接口用于获取 QQ 用户的登录号信息。
	SetBotInfo(profile entity.UserProfile) entity.Common //该接口用于设置 QQ 用户的个人资料信息。
	GetOnlineType(model string) entity.PhoneModel        //该接口用于获取 QQ 在线机型。
	SetOnlineType(model, manu string) entity.Common      //该接口用于设置 QQ 在线机型。
	GetOnlineClientList(noCache bool) entity.Device      //该接口用于获取当前账号在线客户端列表。
}
type contacts interface { //联系人相关
	GetStrangerInfo(userId int64) entity.StrangerInfo                  //该接口用于获取陌生人的信息。
	GetFriendsList(refresh bool) entity.FriendList                     //该接口用于获取好友列表。
	GetUnidirectionFriendsList() entity.UnidirectionalFriend           //该接口用于获取单向好友列表。
	GetGroupInfo(groupId int64) entity.GroupInfo                       //该接口用于获取群信息。
	GetGroupsList() entity.GroupList                                   //该接口用于获取群列表。
	GetGroupMemberInfo(groupId, userId int64) entity.GroupMemberInfo   //该接口用于获取群成员信息。
	GetGroupMemberList(groupId int64) entity.GroupMemberList           //该接口用于获取群成员列表。
	GetGroupHonorInfo(groupId int64) entity.GroupHonorInfo             //该接口用于获取群荣誉信息。
	GetGroupSystemInfo() entity.GroupFileSystemInfo                    //该接口用于获取群系统消息。
	GetFriendSystemInfo() entity.FriendSystemMsg                       //该接口用于获取好友系统消息。
	GetGroupEssentialMessageList(groupId int64) entity.GroupEssenceMsg //该接口用于获取精华消息列表。
	CheckInBlacklist(userId int64) entity.IsBlacklist                  //获取好友/陌生人是否处于黑名单列表。
}
type user interface { //用户相关
	DeleteFriend(userId int64) entity.Common             //该接口用于删除好友。
	DeleteUnidirectionFriend(userId int64) entity.Common //该接口用于删除单向好友。
	RevokeMsg(messageId int32) entity.Common             //该接口用于撤回消息。
}

type message interface { //消息相关
	SendPrivateMessage(userId int64, autoEscape bool, message []entity.MessageItem) entity.PrivateMsg                                               //该接口用于发送私聊消息。
	SendGroupMessage(groupId int64, autoEscape bool, message []entity.MessageItem) entity.GroupMsg                                                  //该接口用于发送群聊消息。
	SendMessage(messageType string, userId int64, groupId int64, discussId int64, autoEscape bool, message []entity.MessageItem) entity.SendMessage //该接口用于发送消息。
	GetMessage(messageId int32) entity.Msg                                                                                                          //该接口用于获取消息。
	GetHistoryMessage(messageType string, userId int64, groupId int64, count int32, messageSeq int32) entity.MessageHistory                         //该接口用于获取历史消息。
	GetGroupHistoryMessage(groupId int64, count int32, messageSeq int32) entity.GroupHistoryMsg                                                     //该接口用于获取群聊历史消息。
	DeleteLocalMessageCache(messageType string, userId int64, groupId int64) entity.Common                                                          //该接口用于清除本地消息缓存。
	GetForwardMessage(id string) entity.ForwardMsg                                                                                                  //该接口用于获取合并转发内容。
	SendGroupForwardMessage(groupId int64, message []entity.GroupForwardMessage) entity.Common                                                      //该接口用于发送群聊合并转发。
	SendPrivateForwardMessage(userId int64, message []entity.MessageItem) entity.Common                                                             //该接口用于发送私聊合并转发。
}
type resource interface { //资源相关
	GetImage(file string) entity.FileInfo            //用于获取图片，只能获取已缓存的图片。
	CheckSendImage() entity.EnableSendImage          //该接口用于检查是否可以发送图片。
	OcrImage(image string) entity.OcrImage           //该接口用于图片 OCR。
	GetAudio(file, outFormat string) entity.AudioMsg //该接口用于获取语音。
	CheckSendAudio() entity.EnableSendAudio          //该接口用于检查是否可以发送语音。
	GetFile()                                        //该接口用于获取文件。
	GetVideo()                                       //该接口用于获取视频。
	GetThumbnail()                                   //该接口用于获取缩略图。
}
type process interface { //处理相关
	AddFriendRequest(flag string, approve bool, remark string) entity.Common                //该接口用于处理加好友请求。
	AddGroupRequest(flag string, subType string, approve bool, reason string) entity.Common //该接口用于处理加群请求／邀请。
}
type group interface { //群聊相关
	SetGroupName(groupId int64, groupName string) entity.Common                            //该接口用于设置群名。
	SetGroupAvatar(groupId int64, file string, cache int) entity.Common                    //该接口用于设置群头像。
	SerGroupAdmin(groupId int64, userId int64, enable bool) entity.Common                  //该接口用于设置群管理员。
	SetGroupNote(groupId int64, userId int64, card string) entity.Common                   //该接口用于设置群备注。
	SetGroupExclusiveTitle(groupId int64, userId int64, specialTitle string) entity.Common //该接口用于设置群组专属头衔。
	SetGroupSingleForbidSpeak(groupId int64, userId int64, duration int64) entity.Common   //该接口用于群单人禁言。
	SetGroupForbidSpeak(groupId int64, enable bool) entity.Common                          //该接口用于群全员禁言。
	SetGroupEssentialMessage(messageId int32) entity.Common                                //该接口用于设置精华消息。
	DeleteGroupEssentialMessage(messageId int32) entity.Common                             //该接口用于移出精华消息。
	SendGroupSign(groupId int64) entity.Common                                             //该接口用于群打卡。
	SendGroupNotice(groupId int64, content string, image string) entity.Common             //该接口用于发送群公告。
	GetGroupNotice(groupId int64) entity.GroupNotice                                       //该接口用于获取群公告。
	SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) entity.Common         //该接口用于群组踢人。
	SetGroupLeave(groupId int64) entity.Common                                             //该接口用于退出群组。
	GroupTouch(groupId int64, userId int64) entity.Common                                  //该接口用于群戳一戳。
	GetForbidSpeakList(groupId int64) entity.GroupProhibitedMember                         //用于获取群聊被禁言人列表。
}
type file interface {
	SendPrivateFile(userId int64, fileUri string, name string) entity.SendPrivateFile //该接口用于上传私聊文件。
	SendGroupFile(groupId int64, fileUri string, name string) entity.SendGroupFile    //该接口用于上传群文件。
	DeleteGroupFile(groupId int64, fileId string, busid int32) entity.Common          //该接口用于删除群文件。
	CreateGroupFileFolder(msgId int32) entity.Common                                  //该接口用于创建群文件文件夹。
	DeleteGroupFolder(groupId int64, folderId string) entity.Common                   //该接口用于删除群文件文件夹。
	GetGroupFileSystemInfo(groupId int64) entity.GroupFileSystemInfo                  //该接口用于获取群文件系统信息。
	GetGroupRootFiles(groupId int64) entity.GroupRootFileList                         //该接口用于获取群根目录文件列表。
	GetGroupFilesFolder(groupId int64, folderId string) entity.GroupFiles             //该接口用于获取群子目录文件列表。
	GetGroupFileUrl(groupId int64, fileId string, busid int32) entity.GroupFileUrl    //该接口用于获取群文件资源链接。
}

type openShamrockSpecialApi interface {
	SwitchAccount(userId int64) entity.Common
	UploadFileToCache(filePath string) entity.UploadCache
	DownloadFileToCache(fileUrl string, name string, base64 string, threadCnt int32, headers map[string]interface{}) entity.DownloadCache
	GetDeviceBattery() entity.Battery
	GetStartTime() entity.StartTime
	GetLog() string
	RunShell(cmd interface{}, dir string)
	StopShamrock()
	GetWeatherCityCode(city string) entity.WeatherCityCode
	GetWeather(code int64, city string) entity.Weather
	UploadGroupImage(pic string, sender int64, troop int64) entity.Common
}
