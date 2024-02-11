package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/ycvk/FreeClover/driver/openshamrock"
	entity "github.com/ycvk/FreeClover/entity/openshamrock"
)

// Resource 资源相关
type Resource struct {
	transceiver adapter.Transceiver
}

// GetImage 用于获取图片，只能获取已缓存的图片。
func (o Resource) GetImage(file string) entity.FileInfo {
	endpoint := "get_image"
	values, err := json.Marshal(map[string]interface{}{
		"file": file,
	})
	if err != nil {
		return entity.FileInfo{}
	}
	return *processJson[entity.FileInfo](endpoint, values, o.transceiver)
}

// CheckSendImage 该接口用于检查是否可以发送图片。
func (o Resource) CheckSendImage() entity.EnableSendImage {
	endpoint := "can_send_image"
	panic("Not implemented: " + endpoint)
}

// OcrImage 该接口用于图片 OCR。
func (o Resource) OcrImage(image string) entity.OcrImage {
	endpoint := "ocr_image"
	panic("Not implemented: " + endpoint)
}

// GetAudio 该接口用于获取语音。
func (o Resource) GetAudio(file, outFormat string) entity.AudioMsg {
	endpoint := "get_record"
	values, err := json.Marshal(map[string]interface{}{
		"file":       file,
		"out_format": outFormat,
	})
	if err != nil {
		return entity.AudioMsg{}
	}
	return *processJson[entity.AudioMsg](endpoint, values, o.transceiver)
}

// CheckSendAudio 该接口用于检查是否可以发送语音。
func (o Resource) CheckSendAudio() entity.EnableSendAudio {
	endpoint := "can_send_record"
	panic("Not implemented: " + endpoint)
}

// GetFile 该接口用于获取文件。
func (o Resource) GetFile() {
	panic("Not implemented")
}

// GetVideo 该接口用于获取视频。
func (o Resource) GetVideo() {
	panic("Not implemented")
}

// GetThumbnail 该接口用于获取缩略图。
func (o Resource) GetThumbnail() {
	panic("Not implemented")
}
