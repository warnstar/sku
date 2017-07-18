package WebMessages

import (
	"golang.org/x/net/websocket"
	"github.com/leesper/holmes"
	"encoding/json"
	"sku/WebServer/WebKey"
	"sku/WebServer/WebMessages/WebUser"
)


type Message struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func Register(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			holmes.Errorf("WebSocket 接收消息失败：%s、n" ,err.Error())
			break
		}

		var message Message

		err := json.Unmarshal([]byte(reply), &message)
		if err != nil {
			holmes.Errorf("此WebSocket消息无法解析：%s\n" ,reply)
		}

		switch message.Type  {
		case WebKey.WEB_USER:
			WebUser.ProcessMessage(ws, message.Content)
		case WebKey.WEB_TSI_NOW_DATA:
		case WebKey.WEB_CLIENT_CONNECT_COMPLETE:
		case WebKey.WEB_CLIENT_TIME_SYNC_COMPLETE:
		case WebKey.WEB_CLIENT_CONNECT_AND_TIME_SYNC_CHECK:
		case WebKey.WEB_CLIENT_TREE_DATA:
		case WebKey.WEB_CLIENT_EXIT:
		case WebKey.WEB_CAN_START_TSI_TEST:
		case WebKey.WEB_TSI_TEST_MODULE_RESULT:
		}
	}
}
