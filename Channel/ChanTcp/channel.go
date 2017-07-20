package ChanTcp

import (
	"sku/SkuServer/SkuRun"
	"sku/WebServer/WebKey"
	"sku/WebServer/WebRun"
	"sku/SkuServer/TcpMessages/TcpKey"
	"sku/SkuServer/TcpMessages/ClientTsiTestPreStart"
	"sku/Channel/ChanWeb"
	"sku/SkuServer/TcpMessages/ClientTsiTestStart"
	"sku/SkuServer/TcpMessages/ClientTsiTestStop"
	"sku/SkuServer/TcpMessages/ClientTsiTestPreStop"
	"sku/SkuServer/TcpMessages"
)

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

var ToTcpChan = make(chan *Message)

func SendTcp(msgType string, msgContent interface{}) {
	msg := Message{Type: msgType, Content: msgContent}

	ToTcpChan <- &msg
}

func init() {
	//消息处理 - 发送到TCP服务器的消息
	go func() {
		for {
			msg := <-ToTcpChan
			//fmt.Printf("tcpServer channel 消息接收：%v\n", *msg)
			go func(){
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
					//检测客户端连接与时间同步

					//获取tcp服务器对象
					tcpServer := <-SkuRun.PiServer
					SkuRun.PiServer <- tcpServer

					if tcpServer.IsAllConnected {
						ChanWeb.SendWeb(WebKey.WEB_CLIENT_CONNECT_COMPLETE, "")

						if tcpServer.IsAllTimeSync {
							ChanWeb.SendWeb(WebKey.WEB_CLIENT_TIME_SYNC_COMPLETE, "")
						}
					}
				case TcpKey.TYPE_CLIENT_TSI_TEST_PRE_START:
					//通知全部pi启动tsi校准
					toPiMsg := new(ClientTsiTestPreStart.Message)
					toPiMsg.Type = toPiMsg.MessageType()

					//发送消息到所有pi
					TcpMessages.SendToAllPi(toPiMsg)

					ChanWeb.SendWebLog(WebKey.LOG_TYPE_SERVER,"通知客户端--启动TSI校准")
				case TcpKey.TYPE_CLIENT_TSI_TEST_PRE_STOP:
					//通知全部pi关闭tsi校准
					toPiMsg := new(ClientTsiTestPreStop.Message)
					toPiMsg.Type = toPiMsg.MessageType()

					//发送消息到所有pi
					TcpMessages.SendToAllPi(toPiMsg)

					ChanWeb.SendWebLog(WebKey.LOG_TYPE_SERVER,"通知客户端--关闭TSI校准")
				case TcpKey.TYPE_CLIENT_TSI_TEST_START:
					//通知全部pi启动tsi测试
					toPiMsg := new(ClientTsiTestStart.Message)
					toPiMsg.Type = toPiMsg.MessageType()

					//发送消息到所有pi
					TcpMessages.SendToAllPi(toPiMsg)

					ChanWeb.SendWebLog(WebKey.LOG_TYPE_SERVER,"通知客户端--启动TSI测试")
				case TcpKey.TYPE_CLIENT_TSI_TEST_STOP:
					//通知全部pi关闭tsi测试
					toPiMsg := new(ClientTsiTestStop.Message)
					toPiMsg.Type = toPiMsg.MessageType()

					//发送消息到所有pi
					TcpMessages.SendToAllPi(toPiMsg)

					ChanWeb.SendWebLog(WebKey.LOG_TYPE_SERVER,"通知客户端--关闭TSI测试")
				}
			}()
		}
	}()
}