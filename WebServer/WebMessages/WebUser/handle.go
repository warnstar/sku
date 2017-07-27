package WebUser

import (
	"github.com/warnstar/holmes"
	"golang.org/x/net/websocket"
	"sku/WebServer/WebKey"
	"sku/WebServer/WebRun"
	"strconv"
	"sku/Channel/ChanWeb"
	"sku/Channel/ChanTcp"
)

type Message struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	TsiHost         string `json:"tsi_host"`
	ClientModuleNum int    `json:"client_module_num"`
	ClientNum       int    `json:"client_num"`
}

func (m Message) MessageType() string {
	return WebKey.WEB_USER
}

func unmarshal(content interface{}) (msg Message) {
	data := content.(map[string]interface{})
	var err error

	msg.ClientModuleNum, err = strconv.Atoi(data["client_module_num"].(string))
	if err != nil {
		holmes.Errorf("web-user 获取内容参数失败：%v\n", "client_module_num")
	}

	msg.ClientNum, err = strconv.Atoi(data["client_num"].(string))
	if err != nil {
		holmes.Errorf("web-user 获取内容参数失败：%v\n", "client_num")
	}

	msg.Port, err = strconv.Atoi(data["port"].(string))
	if err != nil {
		holmes.Errorf("web-user 获取内容参数失败：%v\n", "port")
	}

	msg.Host = data["host"].(string)
	msg.TsiHost = data["tsi_host"].(string)

	return msg
}

func ProcessMessage(ws *websocket.Conn, content interface{}) {
	msg := unmarshal(content)

	wsServer := <-WebRun.ServerChan

	//关闭旧链接
	if wsServer.ClientWs != nil {
		wsServer.ClientWs.Close()
	}

	wsServer.ClientWs = ws
	wsServer.TsiHost = msg.TsiHost
	wsServer.Host = msg.Host
	wsServer.Port = msg.Port
	wsServer.ClientNum = msg.ClientNum
	wsServer.ClientModuleNum = msg.ClientModuleNum
	WebRun.ServerChan <- wsServer

	//通知Tcp服务器
	ChanTcp.SendTcp(msg.MessageType(), msg)

	//通知浏览器
	ChanWeb.SendWebLog(WebKey.LOG_TYPE_SERVER, "已连接服务器")
}
