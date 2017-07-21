package ClientExit

import (
	"context"

	"encoding/json"
	"github.com/leesper/holmes"
	"github.com/leesper/tao"
	"sku/SkuServer/SkuRun"
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
	return 1005
}

// MessageType returns message type .
func (em Message) MessageType() string {
	return "client_exit"
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
	SkuRun.PiServer <- server

	_, err := server.GetPiByConnId(connId)
	if err != nil {
		holmes.Errorln("client-time-sync: 当前链接对应的pi不存在")
		return
	}

}
