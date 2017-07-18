package WebRun

import (
	"golang.org/x/net/websocket"
	"encoding/json"
	"github.com/leesper/holmes"
)

type Server struct {
	Host string `json:"host"`
	Port int `json:"port"`
	TsiHost string `json:"tsi_host"`
	ClientModuleNum int `json:"client_module_num"`
	ClientNum int `json:"client_num"`
	ConnId int64
	ClientWs *websocket.Conn
}

func (s *Server) ConnectSkuServer() {

}

//全局服务器变量
var ServerChan = make(chan *Server, 1)

func SendToClient(msg interface{}) {
	wsServer := <-ServerChan
	ServerChan <- wsServer
	
	msgInfo, err := json.Marshal(msg)
	if err != nil {
		holmes.Errorln("发送至浏览器客户端失败")
	}
	wsServer.ClientWs.Write(msgInfo)


}


func init() {

	s := Server{}
	ServerChan <- &s
}