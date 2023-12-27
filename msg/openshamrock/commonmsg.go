package openshamrockmsg

import (
	entity "github.com/oliverkirk-sudo/FreeClover/entity/openshamrock"
)

type Message struct {
}

func (m Message) At(qq string) entity.MessageItem {
	return entity.MessageItem{
		Type: "at",
		Data: entity.AtData{Qq: qq},
	}
}

func (m Message) AtAll() entity.MessageItem {
	return entity.MessageItem{
		Type: "at",
		Data: entity.AtData{Qq: "all"},
	}
}

func (m Message) Face(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "face",
		Data: entity.FaceData{Id: id},
	}
}

func (m Message) Replay(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "reply",
		Data: entity.ReplyData{Id: id},
	}
}

func (m Message) Image(file string) entity.MessageItem {
	return entity.MessageItem{
		Type: "image",
		Data: entity.ImageData{File: file},
	}
}

func (m Message) Audio(file string) entity.MessageItem {
	return entity.MessageItem{
		Type: "record",
		Data: entity.ImageData{File: file},
	}
}

func (m Message) Video(file string) entity.MessageItem {
	return entity.MessageItem{
		Type: "video",
		Data: entity.VideoData{File: file},
	}
}

func (m Message) BasketBall(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "basketball",
		Data: entity.BasketballData{Id: id},
	}
}

func (m Message) Rps(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "new_rps",
		Data: entity.BasketballData{Id: id},
	}
}

func (m Message) Dict(id int) entity.MessageItem {
	return entity.MessageItem{
		Type: "new_dice",
		Data: entity.BasketballData{Id: id},
	}
}

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

func (m Message) Touch(id int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "touch",
		Data: entity.TouchData{Id: id},
	}
}

func (m Message) Music(musicType string, id int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "music",
		Data: entity.MusicData{Id: id, Type: musicType},
	}
}

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

func (m Message) Weather(city string) entity.MessageItem {
	return entity.MessageItem{
		Type: "weather",
		Data: entity.WeatherData{
			City: city,
		},
	}
}

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

func (m Message) Gift(qq, id int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "gift",
		Data: entity.GiftData{
			Qq: qq,
			Id: id,
		},
	}
}

func (m Message) ForWard(id string) entity.MessageItem {
	return entity.MessageItem{
		Type: "forward",
		Data: entity.ForwardData{
			Id: id,
		},
	}
}

func (m Message) ForwardNode(id int64) entity.MessageItem {
	return entity.MessageItem{
		Type: "forward",
		Data: entity.ForwardNodeData{
			Id: id,
		},
	}
}

func (m Message) XML() entity.MessageItem {
	panic("Not implemented")
}

func (m Message) JSON(data string) entity.MessageItem {
	return entity.MessageItem{
		Type: "json",
		Data: entity.JsonData{
			Data: data,
		},
	}
}

func (m Message) TextToAudio() entity.MessageItem {
	panic("Not implemented")
}

func (m Message) Plain(text string) entity.MessageItem {
	return entity.MessageItem{
		Type: "text",
		Data: entity.TextData{Text: text},
	}
}

func (m Message) AppendMessageList(msg ...entity.MessageItem) []entity.MessageItem {
	var msgList []entity.MessageItem
	for _, item := range msg {
		msgList = append(msgList, item)
	}
	return msgList
}
