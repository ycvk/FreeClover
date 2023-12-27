package client

import (
	openShamrockApi "github.com/oliverkirk-sudo/FreeClover/adapter/openshamrock"
	openShamrockHttp "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock/rhttp"
	openShamrockWs "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock/rws"
	"github.com/oliverkirk-sudo/FreeClover/log"
)

type BotClient struct {
	HttpApi            *openShamrockApi.OpenShamrockApi
	HttpWebHookApi     *openShamrockApi.OpenShamrockApi
	WsApi              *openShamrockApi.OpenShamrockApi
	ReverseWsApi       *openShamrockApi.OpenShamrockApi
	HttpAdapter        *openShamrockHttp.HttpDriver
	WsAdapter          *openShamrockWs.WsDriver
	HttpWebhookAdapter *openShamrockHttp.HttpWebHookDriver
	ReverseWsAdapter   *openShamrockWs.ReverseWsDriver
}

func GetBotClient() *BotClient {
	client := BotClient{}
	return &client
}

func (b *BotClient) WithOpenShamrockHttpWebHookClient(url, authToken string) *BotClient {
	log.Log.Debug("添加回调HTTP监听")
	adapter := openShamrockHttp.HttpWebHookDriver{}
	api := openShamrockApi.NewOpenShamrockApi(url, authToken, &adapter)
	if api == nil {
		log.Log.Warning("[WebHook] 连接失败")
	}
	b.HttpWebhookAdapter = &adapter
	b.HttpWebHookApi = api
	return b
}
func (b *BotClient) WithOpenShamrockHttpClient(url, authToken string) *BotClient {
	log.Log.Debug("添加主动HTTP请求")
	adapter := openShamrockHttp.HttpDriver{}
	api := openShamrockApi.NewOpenShamrockApi(url, authToken, &adapter)
	if api == nil {
		log.Log.Warning("[HTTP] 连接失败")
	}
	b.HttpAdapter = &adapter
	b.HttpApi = api
	return b
}
func (b *BotClient) WithOpenShamrockReverseWsClient(url, authToken string) *BotClient {
	log.Log.Debug("添加被动WebSocket监听")
	adapter := openShamrockWs.ReverseWsDriver{}
	api := openShamrockApi.NewOpenShamrockApi(url, authToken, &adapter)
	if api == nil {
		log.Log.Warning("[ReverseWebSocket] 连接失败")
	}
	b.ReverseWsAdapter = &adapter
	b.ReverseWsApi = api
	return b
}
func (b *BotClient) WithOpenShamrockWsClient(url, authToken string) *BotClient {
	log.Log.Debug("添加主动WebSocket监听")
	adapter := openShamrockWs.WsDriver{}
	api := openShamrockApi.NewOpenShamrockApi(url, authToken, &adapter)
	if api == nil {
		log.Log.Warning("[WebSocket] 连接失败")
	}
	b.WsAdapter = &adapter
	b.WsApi = api
	return b
}
