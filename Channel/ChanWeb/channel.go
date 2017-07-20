package ChanWeb

import (
	"sku/WebServer/WebKey"
	"sku/WebServer/WebRun"
	"sku/SkuServer/TcpMessages/TcpKey"
)

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

var ToWebChan = make(chan *Message)

func SendWeb(msgType string, msgContent interface{}) {
	msg := Message{Type: msgType, Content: msgContent}

	ToWebChan <- &msg
}

func init() {
	//消息处理 - 发送到WebSocket服务器的消息
	go func() {
		for {
			msg := <-ToWebChan
			go func(){
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
				case TcpKey.TYPE_CLIENT_TSI_TEST_MODULE_RESULT:
					WebRun.SendToClient(msg)
				case WebKey.WEB_CAN_START_TSI_TEST:
					WebRun.SendToClient(msg)
				}
			}()
		}
	}()
}
