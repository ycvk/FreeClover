package openShamrockEntity

type CQCode struct {
	CQ string
}
type MessageItem struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type MessagesList struct {
	Message []MessageItem `json:"message"`
}

type TextData struct {
	Text string `json:"text"`
}
type FaceData struct {
	Id int `json:"id"`
}
type AtData struct {
	Qq string `json:"qq"`
}
type ReplyData struct {
	Id int `json:"id"`
}
type ImageData struct {
	File string `json:"file"`
}
type AudioData struct {
	File string `json:"file"`
}
type VideoData struct {
	File string `json:"file"`
}
type BasketballData struct {
	Id int `json:"id"`
}
type NewRpsData struct {
	Id int `json:"id"`
}
type NewDict struct {
	Id int `json:"id"`
}
type PokeData struct {
	Type     int64 `json:"type"`
	Id       int64 `json:"id"`
	Strength int64 `json:"strength"`
}
type TouchData struct {
	Id int64 `json:"id"`
}
type MusicData struct {
	Type string `json:"type"`
	Id   int64  `json:"id"`
}
type MusicCustomData struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Audio  string `json:"audio"`
	Title  string `json:"title"`
	Singer string `json:"singer"`
	Image  string `json:"image"`
}
type WeatherData struct {
	City string `json:"city"`
}
type LocationData struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
}
type ShareData struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
	File    string `json:"file"`
}
type GiftData struct {
	Qq int64 `json:"qq"`
	Id int64 `json:"id"`
}
type ForwardData struct {
	Id string `json:"id"`
}
type ForwardNodeData struct {
	Id int64 `json:"id"`
}
type JsonData struct {
	Data string `json:"data"`
}
