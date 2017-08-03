package WebMessages

import (
	"encoding/json"
	"github.com/warnstar/holmes"
	"golang.org/x/net/websocket"
	"sku/SkuServer/Tsi"
	"sku/WebServer/WebKey"
	"sku/WebServer/WebMessages/WebUser"
	"sku/Channel/ChanTcp"
	"time"
)

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

func Register(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			holmes.Errorf("WebSocket 接收消息失败：%s、n", err.Error())
			break
		}
		holmes.Infof("websocket 接收:%v\n", reply)
		var message Message

		err := json.Unmarshal([]byte(reply), &message)
		if err != nil {
			holmes.Errorf("此WebSocket消息无法解析：%s\n", reply)
		}

		switch message.Type {
		case WebKey.WEB_USER:
			WebUser.ProcessMessage(ws, message.Content)
		case WebKey.WEB_TSI_CHECK:
			//将旧tsi文件迁移到历史文件夹
			Tsi.MoveFileToHistory()

			//处理 tsi 检查
			tsiChan := <-Tsi.TsiClientChan
			tsiChan.Type = Tsi.TSI_RUN_TYPE_CHECK
			Tsi.TsiClientChan <- tsiChan

			if !tsiChan.IsRunning {
				Tsi.Connect()
			}

			if tsiChan.IsRunning {
				//等待1s
				time.Sleep(time.Second)

				//开启读取tsi数据
				Tsi.ControlTsi(Tsi.TSI_SERVER_START, "")
				Tsi.ControlTsi(Tsi.TSI_SERVER_RECEIVE_DATA_START, "")
			}
		case WebKey.WEB_TSI_TEST_PRE:
			//处理 tsi 校验
			tsiChan := <-Tsi.TsiClientChan
			tsiChan.Type = Tsi.TSI_RUN_TYPE_TEST_PRE
			Tsi.TsiClientChan <- tsiChan

			if !tsiChan.IsRunning {
				Tsi.Connect()
			}

			if tsiChan.IsRunning {
				//等待1s
				time.Sleep(time.Second)

				holmes.Infof("tsi校准启动:%v\n", tsiChan.IsRunning)

				//开启读取tsi数据
				Tsi.ControlTsi(Tsi.TSI_SERVER_START, "")
				Tsi.ControlTsi(Tsi.TSI_SERVER_RECEIVE_DATA_START, "")
			}
		case WebKey.WEB_TSI_TEST:
			//处理 tsi 校验
			tsiChan := <-Tsi.TsiClientChan
			tsiChan.Type = Tsi.TSI_RUN_TYPE_TEST
			Tsi.TsiClientChan <- tsiChan

			if !tsiChan.IsRunning {
				Tsi.Connect()
			}

			if tsiChan.IsRunning {
				//等待1s
				time.Sleep(time.Second)

				holmes.Infof("tsi测试启动:%v\n", tsiChan.IsRunning)

				//开启读取tsi数据
				Tsi.ControlTsi(Tsi.TSI_SERVER_START, "")
				Tsi.ControlTsi(Tsi.TSI_SERVER_RECEIVE_DATA_START, "")
			}
		case WebKey.WEB_CLIENT_CONNECT_AND_TIME_SYNC_CHECK:
			ChanTcp.SendTcp(WebKey.WEB_CLIENT_CONNECT_AND_TIME_SYNC_CHECK, "")
		case WebKey.WEB_CLIENT_EXIT:

			tsiChan := <-Tsi.TsiClientChan
			tsiChan.Type = Tsi.TSI_RUN_TYPE_TEST
			Tsi.TsiClientChan <- tsiChan
			if tsiChan.IsRunning {
				//tsi客户端关闭
				Tsi.ControlTsi(Tsi.TSI_SERVER_EXIT,"")
			}


			//通知tcp服务器关闭
			ChanTcp.SendTcp(WebKey.WEB_CLIENT_EXIT, "")

		}
	}
}
