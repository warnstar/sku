package ClientGetFileTsiTest

import (
	"context"

	"github.com/leesper/holmes"
	"github.com/leesper/tao"
	"sku/SkuServer/SkuRun"
)

// Message defines the echo message.
type Message string

// Serialize serializes Message into bytes.
func (em Message) Serialize() ([]byte, error) {
	return []byte(string(em)), nil
}

// MessageNumber returns message type number.
func (em Message) MessageNumber() int32 {
	return 1107
}

// MessageType returns message type .
func (em Message) MessageType() string {
	return "client_get_file_tsi_test"
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	return Message(string(data)), err
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
