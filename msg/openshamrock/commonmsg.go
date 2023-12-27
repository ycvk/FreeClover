package openshamrockmsg

import (
	entity "github.com/oliverkirk-sudo/FreeClover/entity/openshamrock"
)

type Message struct {
}

// At at消息，接收一个qq号码
func (m Message) At(qq string) entity.MessageItem {
	return entity.MessageItem{
		Type: "at",
		Data: entity.AtData{Qq: qq},
	}
}

// AtAll at全体消息
func (m Message) AtAll() entity.MessageItem {
	return entity.MessageItem{
		Type: "at",
		Data: entity.AtData{Qq: "all"},
	}
}

// Face 表情消息，接收一个表情id，表情id在这https://github.com/richardchien/coolq-http-api/wiki/%E8%A1%A8%E6%83%85-CQ-%E7%A0%81-ID-%E8%A1%A8
func (m Message) Face(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "face",
		Data: entity.FaceData{Id: id},
	}
}

// Replay 回复消息，接收一个消息id
func (m Message) Replay(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "reply",
		Data: entity.ReplyData{Id: id},
	}
}

// Image 图片消息，接收一个消息路径，file://,http://,https://,base64://
func (m Message) Image(file string) entity.MessageItem {
	return entity.MessageItem{
		Type: "image",
		Data: entity.ImageData{File: file},
	}
}

// Audio 语音消息，接收一个消息路径，file://,http://,https://,base64://
func (m Message) Audio(file string) entity.MessageItem {
	return entity.MessageItem{
		Type: "record",
		Data: entity.ImageData{File: file},
	}
}

// Video 视频消息，接收一个消息路径，file://,http://,https://,base64://
func (m Message) Video(file string) entity.MessageItem {
	return entity.MessageItem{
		Type: "video",
		Data: entity.VideoData{File: file},
	}
}

// BasketBall 篮球消息，接收一个类型id，5 没中, 4 擦边没中, 3 卡框, 2 擦边中, 1 正中
func (m Message) BasketBall(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "basketball",
		Data: entity.BasketballData{Id: id},
	}
}

// Rps 猜拳消息，接收一个猜拳id，锤 3 剪 2 布 1
func (m Message) Rps(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "new_rps",
		Data: entity.BasketballData{Id: id},
	}
}

// Dict 骰子消息，接收一个骰子id，点数 ID
func (m Message) Dict(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "new_dice",
		Data: entity.BasketballData{Id: id},
	}
}

// Poke 戳一戳消息，接收戳一戳类型，戳一戳 ID，戳一戳强度(1-5 默认1)
func (m Message) Poke(pokeType, id, strength int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "poke",
		Data: entity.PokeData{
			Type:     pokeType,
			Id:       id,
			Strength: strength,
		},
	}
}

// Touch 戳一戳(双击头像)消息，接收一个qq号码
func (m Message) Touch(id int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "touch",
		Data: entity.TouchData{Id: id},
	}
}

// Music 音乐消息，接收音乐类型(qq/163)，音乐id
func (m Message) Music(musicType string, id int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "music",
		Data: entity.MusicData{Id: id, Type: musicType},
	}
}

// MusicCustom 自定义音乐消息，接收音乐类型(自定义请使用 custom),跳转链接,音乐音频链接,标题,歌手,封面图片链接
func (m Message) MusicCustom(musicType, url, audio, title, singer, image string) entity.MessageItem {
	return entity.MessageItem{
		Type: "music",
		Data: entity.MusicCustomData{
			Type:   musicType,
			Url:    url,
			Audio:  audio,
			Title:  title,
			Singer: singer,
			Image:  image,
		},
	}
}

// Weather 天气消息，接收一个城市名称
func (m Message) Weather(city string) entity.MessageItem {
	return entity.MessageItem{
		Type: "weather",
		Data: entity.WeatherData{
			City: city,
		},
	}
}

// Location 位置消息，接收纬度，经度，标题，内容
func (m Message) Location(lat float64, lon float64, title string, content string) entity.MessageItem {
	return entity.MessageItem{
		Type: "location",
		Data: entity.LocationData{
			Lat:     lat,
			Lon:     lon,
			Title:   title,
			Content: content,
		},
	}
}

// Share 链接消息，接收链接地址，标题，内容，图片链接，文件链接
func (m Message) Share(url, title, content, image, file string) entity.MessageItem {
	return entity.MessageItem{
		Type: "share",
		Data: entity.ShareData{
			Url:     url,
			Title:   title,
			Content: content,
			Image:   image,
			File:    file,
		},
	}
}

// Gift 礼物消息，接收QQ号，礼物ID
func (m Message) Gift(qq, id int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "gift",
		Data: entity.GiftData{
			Qq: qq,
			Id: id,
		},
	}
}

// ForWard 合并转发消息，接收一个合并转发resid
func (m Message) ForWard(id string) entity.MessageItem {
	return entity.MessageItem{
		Type: "forward",
		Data: entity.ForwardData{
			Id: id,
		},
	}
}

// ForwardNode 合并转发(节点)消息，接收一个消息ID
func (m Message) ForwardNode(id int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "forward",
		Data: entity.ForwardNodeData{
			Id: id,
		},
	}
}

// XML XML消息，未实现
func (m Message) XML() entity.MessageItem {
	panic("Not implemented")
}

// JSON json消息，接收一段json文本
func (m Message) JSON(data string) entity.MessageItem {
	return entity.MessageItem{
		Type: "json",
		Data: entity.JsonData{
			Data: data,
		},
	}
}

// TextToAudio 文本转语音消息，未实现
func (m Message) TextToAudio() entity.MessageItem {
	panic("Not implemented")
}

// Plain 普通文本消息，接收一段文本
func (m Message) Plain(text string) entity.MessageItem {
	return entity.MessageItem{
		Type: "text",
		Data: entity.TextData{Text: text},
	}
}

// AppendMessageList 添加消息到消息列表，接收一系列消息，返回消息列表
func (m Message) AppendMessageList(msg ...entity.MessageItem) []entity.MessageItem {
	var msgList []entity.MessageItem
	for _, item := range msg {
		msgList = append(msgList, item)
	}
	return msgList
}
