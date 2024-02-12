package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/ycvk/FreeClover/driver/openshamrock"
	entity "github.com/ycvk/FreeClover/entity/openshamrock"
	"io"
	"os"
)

// OpenShamrock Shamrock特殊接口相关
type OpenShamrock struct {
	transceiver adapter.Transceiver
}

// SwitchAccount 该接口用于切换账户。
func (o OpenShamrock) SwitchAccount(userId int64) entity.Common {
	endpoint := "switch_account"
	values, err := json.Marshal(map[string]interface{}{
		"user_id": userId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// UploadFileToCache 该接口用于上传文件至缓存目录。
func (o OpenShamrock) UploadFileToCache(filePath string) entity.UploadCache {
	endpoint := "upload_file"
	// 打开文件并将文件转为byte[]
	file, err := os.Open(filePath)
	if err != nil {
		return entity.UploadCache{}
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return entity.UploadCache{}
	}
	return *processJson[entity.UploadCache](endpoint, bytes, o.transceiver, RequestTypeFile)
}

// DownloadFileToCache 该接口用于让设备下载链接的文件
func (o OpenShamrock) DownloadFileToCache(fileUrl string, name string, base64 string, threadCnt int32, headers map[string]interface{}) entity.DownloadCache {
	endpoint := "download_file"
	values, err := json.Marshal(map[string]interface{}{
		"url":        fileUrl,
		"name":       name,
		"base64":     base64,
		"thread_cnt": threadCnt,
		"headers":    headers,
	})
	if err != nil {
		return entity.DownloadCache{}
	}
	return *processJson[entity.DownloadCache](endpoint, values, o.transceiver)
}

// GetDeviceBattery 该接口用于获取设置电量
func (o OpenShamrock) GetDeviceBattery() entity.Battery {
	endpoint := "get_device_battery"
	return *processJson[entity.Battery](endpoint, nil, o.transceiver)
}

// GetStartTime 该接口用于获取设备启动时间，WS不可以
func (o OpenShamrock) GetStartTime() entity.StartTime {
	endpoint := "get_start_time"
	return *processJson[entity.StartTime](endpoint, nil, o.transceiver)
}

// GetLog 该接口用于获取文本日志
func (o OpenShamrock) GetLog() string {
	endpoint := "log"
	panic("Not implemented: " + endpoint)
	//return *processJson[string](endpoint, nil, o.transceiver)
}

func (o OpenShamrock) RunShell(cmd interface{}, dir string) {
	endpoint := ""
	panic("Not implemented: " + endpoint)
}

// StopShamrock 该接口用于关闭Shamrock
func (o OpenShamrock) StopShamrock() {
	endpoint := "shut"
	processJson[entity.Common](endpoint, nil, o.transceiver)
}

// GetWeatherCityCode 该接口用于获取天气城市代码。
func (o OpenShamrock) GetWeatherCityCode(city string) entity.WeatherCityCode {
	endpoint := "get_weather_city_code"
	values, err := json.Marshal(map[string]interface{}{
		"city": city,
	})
	if err != nil {
		return entity.WeatherCityCode{}
	}
	return *processJson[entity.WeatherCityCode](endpoint, values, o.transceiver)
}

// GetWeather 该接口用于获取天气，该接口调用来自QQ官方服务。
func (o OpenShamrock) GetWeather(code int64, city string) entity.Weather {
	endpoint := "get_weather_city_code"
	values, err := json.Marshal(map[string]interface{}{
		"city": city,
		"code": code,
	})
	if err != nil {
		return entity.Weather{}
	}
	return *processJson[entity.Weather](endpoint, values, o.transceiver)

}

// UploadGroupImage 该接口用于上传群聊图片, 注意该接口是上传群消息的图片，不是群文件，也不是群相册。
func (o OpenShamrock) UploadGroupImage(pic string, sender int64, troop int64) entity.Common {
	endpoint := "upload_group_image"
	values, err := json.Marshal(map[string]interface{}{
		"pic":    pic,
		"sender": sender,
		"troop":  troop,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}
