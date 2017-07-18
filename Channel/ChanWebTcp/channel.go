package ChanWebTcp

import (
	"sku/WebServer/WebKey"
	"sku/WebServer/WebRun"
)

type Message struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

var ToWeb = make(chan *Message)
var ToTcp = make(chan *Message)

func SendTcp(msgType string, msgContent interface{}) {
	msg := Message{Type:msgType,Content:msgContent}

	ToTcp <- &msg
}


func SendWeb(msgType string, msgContent interface{}) {
	msg := Message{Type:msgType,Content:msgContent}

	ToWeb <- &msg
}



func init() {
	go func(){
		for {
			msg := <- ToWeb
			
			//fmt.Printf("websocket channel 消息接收：%v\n", *msg)
			switch msg.Type {
			case WebKey.WEB_CLIENT_LOG:
				//发送至浏览器客户端
				WebRun.SendToClient(msg)
			}
		}
	}()

	go func(){
		for {
			msg := <- ToTcp
			//fmt.Printf("tcpServer channel 消息接收：%v\n", *msg)

			switch msg.Type {
			case WebKey.WEB_USER:

			}
		}
	}()
}