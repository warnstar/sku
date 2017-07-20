package ChanWebTcp

import (
	"sku/SkuServer/SkuRun"
	"sku/WebServer/WebKey"
	"sku/WebServer/WebRun"
)

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

var ToWebChan = make(chan *Message)
var ToTcpChan = make(chan *Message)

func SendTcp(msgType string, msgContent interface{}) {
	msg := Message{Type: msgType, Content: msgContent}

	ToTcpChan <- &msg
}

func SendWeb(msgType string, msgContent interface{}) {
	msg := Message{Type: msgType, Content: msgContent}

	ToWebChan <- &msg
}

func init() {
	go func() {
		for {
			msg := <-ToWebChan

			//fmt.Printf("websocket channel 消息接收：%v\n", *msg)
			switch msg.Type {
			case WebKey.WEB_CLIENT_LOG:
				//发送至浏览器客户端
				WebRun.SendToClient(msg)
			case WebKey.WEB_CLIENT_CONNECT_COMPLETE:
				WebRun.SendToClient(msg)
			case WebKey.WEB_CLIENT_TIME_SYNC_COMPLETE:
				WebRun.SendToClient(msg)
			case WebKey.WEB_TSI_CHECK:
				WebRun.SendToClient(msg)
			case WebKey.WEB_TSI_NOW_DATA:
				WebRun.SendToClient(msg)
			}
		}
	}()

	go func() {
		for {
			msg := <-ToTcpChan
			//fmt.Printf("tcpServer channel 消息接收：%v\n", *msg)

			switch msg.Type {
			case WebKey.WEB_USER:
				//设置服务器信息 （tsi服务器地址，最大客户端数量）
				server := <-SkuRun.PiServer
				webServer := <-WebRun.ServerChan
				WebRun.ServerChan <- webServer

				server.PiMaxNum = webServer.ClientNum
				server.TsiServerAddress = webServer.TsiHost
				SkuRun.PiServer <- server
			case WebKey.WEB_CLIENT_CONNECT_AND_TIME_SYNC_CHECK:
				tcpServer := <-SkuRun.PiServer
				SkuRun.PiServer <- tcpServer

				if tcpServer.IsAllConnected {
					SendWeb(WebKey.WEB_CLIENT_CONNECT_COMPLETE, "")

					if tcpServer.IsAllTimeSync {
						SendWeb(WebKey.WEB_CLIENT_TIME_SYNC_COMPLETE, "")
					}
				}
			}
		}
	}()
}
