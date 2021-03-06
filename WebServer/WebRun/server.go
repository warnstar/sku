package WebRun

import (
	"encoding/json"
	"github.com/warnstar/holmes"
	"golang.org/x/net/websocket"
)

type Server struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	TsiHost         string `json:"tsi_host"`
	ClientModuleNum int    `json:"client_module_num"`
	ClientNum       int    `json:"client_num"`
	ConnId          int64
	ClientWs        *websocket.Conn
}


//全局服务器变量
var ServerChan = make(chan *Server, 1)


func init() {
	s := Server{}
	ServerChan <- &s
}

func ResetServer() {
	<-ServerChan

	s := Server{}
	ServerChan <- &s

	holmes.Infoln("重置websocket服务器")
}

func SendToClient(msg interface{}) {
	wsServer := <-ServerChan
	ServerChan <- wsServer

	if wsServer.ClientWs != nil {
		msgInfo, err := json.Marshal(msg)
		if err != nil {
			holmes.Errorln("发送至浏览器客户端失败")
		}
		wsServer.ClientWs.Write(msgInfo)

	} else {
		holmes.Infof("浏览器客户端未连接")
	}

}