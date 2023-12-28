# FreeClover

**基于Go的[OpenShamrock](https://github.com/whitechi73/OpenShamrock)登录框架，兼容OneBot11/12的接口**

## 连接方式

- [x] HTTP
- [x] 被动HTTP
- [x] 主动WebSocket
- [x] 被动WebSocket

## 使用方法
### 安装
`go get github.com/oliverkirk-sudo/FreeClover`

```golang
package main

import (
	"github.com/duke-git/lancet/v2/formatter"
	"github.com/oliverkirk-sudo/FreeClover/client"
	entity "github.com/oliverkirk-sudo/FreeClover/entity/openshamrock"
)

func main() {
	bot := client.GetAdapterClient()
	bot.WithOpenShamrockHttpClient("http://192.168.2.181:5801", "auth_token")

	ad := bot.HttpAdapter //http适配器

	msgList := ad.Message.AppendMessageList(
		ad.Message.Plain("测试文本"),
		ad.Message.Face(123),
	)
	resp := ad.Api.Message.SendPrivateMessage(00000,false,msgList)
	pretty, err := formatter.Pretty(resp)
	if err != nil {
		return
	}
	println(pretty)
}
```
## 许可证
**AGPL-3.0**