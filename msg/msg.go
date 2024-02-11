package msg

import entity "github.com/ycvk/FreeClover/entity/openshamrock"

type MessageType interface {
	At(qq string) entity.MessageItem
	AtAll() entity.MessageItem
	Face(id int) entity.MessageItem
	Replay(id int) entity.MessageItem
	Image(file string) entity.MessageItem
	Audio(file string) entity.MessageItem
	Video(file string) entity.MessageItem
	BasketBall(id int) entity.MessageItem
	Rps(id int) entity.MessageItem
	Dict(id int) entity.MessageItem
	Poke(pokeType, id, strength int64) entity.MessageItem
	Touch(id int64) entity.MessageItem
	Music(musicType string, id int64) entity.MessageItem
	MusicCustom(musicType, url, audio, title, singer, image string) entity.MessageItem
	Weather(city string) entity.MessageItem
	Location(lat float64, lon float64, title string, content string) entity.MessageItem
	Share(url, title, content, image, file string) entity.MessageItem
	Gift(qq, id int64) entity.MessageItem
	ForWard(id string) entity.MessageItem
	ForwardNode(id int64) entity.MessageItem
	XML() entity.MessageItem
	JSON(data string) entity.MessageItem
	TextToAudio() entity.MessageItem
	Plain(text string) entity.MessageItem
	AppendMessageList(msg ...entity.MessageItem) []entity.MessageItem
}
