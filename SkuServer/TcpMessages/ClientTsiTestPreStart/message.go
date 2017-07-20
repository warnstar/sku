package ClientTsiTestPreStart

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
	Content interface{}  `json:"content"`
}

// Serialize serializes Message into bytes.
func (em Message) Serialize() ([]byte, error) {
	msg, err := json.Marshal(em)
	return msg, err
}

// MessageNumber returns message type number.
func (em Message) MessageNumber() int32 {
	return 1201
}

// MessageType returns message type .
func (em Message) MessageType() string {
	return "client_tsi_test_pre_start"
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

	//取tcp服务器 变量
	server := <-SkuRun.PiServer

	thisPi, err := server.GetPiByConnId(connId)
	if err != nil {
		holmes.Errorln("client-time-sync: 当前链接对应的pi不存在")
		return
	}

	//设置客户端校准状态
	thisPi.IsTsiPreStart = true
	server.UpdatePiByConnId(connId, thisPi)

	//设置tcp服务器 变量
	SkuRun.PiServer <- server

	//通知浏览器-客户端已经时间同步
	ChanWeb.SendWebLog(WebKey.LOG_TYPE_CLIENT, fmt.Sprintf("%v确认已启动TSI校准", thisPi.Info.Name))
}
