package client

import (
	openShamrockApi "github.com/oliverkirk-sudo/FreeClover/adapter/openshamrock"
	openShamrockHttp "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock/rhttp"
	openShamrockWs "github.com/oliverkirk-sudo/FreeClover/driver/openshamrock/rws"
	"github.com/oliverkirk-sudo/FreeClover/log"
)

// AdapterClient 适配器客户端结构体
type AdapterClient struct {
	HttpAdapter        *openShamrockApi.OpenShamrockAdapter
	HttpWebHookAdapter *openShamrockApi.OpenShamrockAdapter
	WsAdapter          *openShamrockApi.OpenShamrockAdapter
	ReverseWsAdapter   *openShamrockApi.OpenShamrockAdapter
	HttpDriver         *openShamrockHttp.HttpDriver
	WsDriver           *openShamrockWs.WsDriver
	HttpWebHookDriver  *openShamrockHttp.HttpWebHookDriver
	ReverseWsDriver    *openShamrockWs.ReverseWsDriver
}

// GetAdapterClient 获取一个适配器客户端
func GetAdapterClient() *AdapterClient {
	client := new(AdapterClient)
	return client
}

// WithOpenShamrockHttpWebHookClient 为适配器添加webhook通信方法
func (b *AdapterClient) WithOpenShamrockHttpWebHookClient(url, authToken string) *AdapterClient {
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

// WithOpenShamrockHttpClient 为适配器添加HTTP通信方法
func (b *AdapterClient) WithOpenShamrockHttpClient(url, authToken string) *AdapterClient {
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

// WithOpenShamrockReverseWsClient 为适配器添加反向WebSocket通信方法
func (b *AdapterClient) WithOpenShamrockReverseWsClient(url, authToken string) *AdapterClient {
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

// WithOpenShamrockWsClient 为适配器添加正向WebSocket通信方法
func (b *AdapterClient) WithOpenShamrockWsClient(url, authToken string) *AdapterClient {
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
