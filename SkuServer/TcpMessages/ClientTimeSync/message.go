package ClientTimeSync

import (
	"context"

	"encoding/json"
	"fmt"
	"github.com/leesper/holmes"
	"github.com/leesper/tao"
	"sku/SkuServer/SkuRun"
	"sku/WebServer/WebKey"
	"sku/Channel/ChanWeb"
)

// Message defines the echo message.
type Message struct {
	Type    string `json:"type"`
	Content int64  `json:"content"`
}

// Serialize serializes Message into bytes.
func (em Message) Serialize() ([]byte, error) {
	msg, err := json.Marshal(em)
	return msg, err
}

// MessageNumber returns message type number.
func (em Message) MessageNumber() int32 {
	return 1002
}

// MessageType returns message type .
func (em Message) MessageType() string {
	return "client_time_sync"
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	if data == nil {
		return nil, tao.ErrNilData
	}

	var body Message

	err = json.Unmarshal(data, &body)

	return body, err
}

// ProcessMessage process the logic of echo message.
func ProcessMessage(ctx context.Context, conn tao.WriteCloser) {
	connId := tao.NetIDFromContext(ctx)

	server := <-SkuRun.PiServer
	thisPi, err := server.GetPiByConnId(connId)
	if err != nil {
		holmes.Errorln("client-time-sync: 当前链接对应的pi不存在")
		return
	}

	thisPi.IsTimeSync = true
	server.UpdatePiByConnId(connId, thisPi)

	//通知浏览器-客户端已经时间同步
	ChanWeb.SendWebLog(WebKey.LOG_TYPE_CLIENT, fmt.Sprintf("%v已时间同步", thisPi.Info.Name))

	if server.CheckPiAllTimeSync() {
		holmes.Infoln("全部已经时间同步")

		ChanWeb.SendWebLog(WebKey.LOG_TYPE_SERVER, "全部客户端已经时间同步")

		//通知用户
		ChanWeb.SendWeb(WebKey.WEB_CLIENT_TIME_SYNC_COMPLETE, "")
	}

	SkuRun.PiServer <- server

	//发送pi树形数据到浏览器
	ChanWeb.SendWeb(WebKey.WEB_CLIENT_TREE_DATA, SkuRun.GetClientTree())
}
