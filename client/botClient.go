package client

import (
	openShamrockApi "github.com/oliverkirk-sudo/FreeClover/adapter/openshamrock"
	openShamrockHttp "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock/rhttp"
	openShamrockWs "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock/rws"
	"github.com/oliverkirk-sudo/FreeClover/log"
)

type BotClient struct {
	HttpAdapter        *openShamrockApi.OpenShamrockApi
	HttpWebHookAdapter *openShamrockApi.OpenShamrockApi
	WsAdapter          *openShamrockApi.OpenShamrockApi
	ReverseWsAdapter   *openShamrockApi.OpenShamrockApi
	HttpDriver         *openShamrockHttp.HttpDriver
	WsDriver           *openShamrockWs.WsDriver
	HttpWebHookDriver  *openShamrockHttp.HttpWebHookDriver
	ReverseWsDriver    *openShamrockWs.ReverseWsDriver
}

func GetBotClient() *BotClient {
	client := BotClient{}
	return &client
}

func (b *BotClient) WithOpenShamrockHttpWebHookClient(url, authToken string) *BotClient {
	log.Log.Debug("添加回调HTTP监听")
	driver := openShamrockHttp.HttpWebHookDriver{}
	adapter := openShamrockApi.NewOpenShamrockDriver(url, authToken, &driver)
	if adapter == nil {
		log.Log.Warning("[WebHook] 连接失败")
	}
	b.HttpWebHookAdapter = adapter
	b.HttpWebHookDriver = &driver
	return b
}
func (b *BotClient) WithOpenShamrockHttpClient(url, authToken string) *BotClient {
	log.Log.Debug("添加主动HTTP请求")
	driver := openShamrockHttp.HttpDriver{}
	adapter := openShamrockApi.NewOpenShamrockDriver(url, authToken, &driver)
	if adapter == nil {
		log.Log.Warning("[HTTP] 连接失败")
	}
	b.HttpAdapter = adapter
	b.HttpDriver = &driver
	return b
}
func (b *BotClient) WithOpenShamrockReverseWsClient(url, authToken string) *BotClient {
	log.Log.Debug("添加被动WebSocket监听")
	driver := openShamrockWs.ReverseWsDriver{}
	adapter := openShamrockApi.NewOpenShamrockDriver(url, authToken, &driver)
	if adapter == nil {
		log.Log.Warning("[ReverseWebSocket] 连接失败")
	}
	b.ReverseWsAdapter = adapter
	b.ReverseWsDriver = &driver
	return b
}
func (b *BotClient) WithOpenShamrockWsClient(url, authToken string) *BotClient {
	log.Log.Debug("添加主动WebSocket监听")
	driver := openShamrockWs.WsDriver{}
	adapter := openShamrockApi.NewOpenShamrockDriver(url, authToken, &driver)
	if adapter == nil {
		log.Log.Warning("[WebSocket] 连接失败")
	}
	b.WsAdapter = adapter
	b.WsDriver = &driver
	return b
}
