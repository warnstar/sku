package clientInfo

import (
	"context"

	"github.com/leesper/tao"
	"encoding/json"
	"sku/pi"
)

// Message defines the echo message.
type Message struct {
	Type string
	Content struct{
		Name string
		ConnectNow int
	}
}

// Serialize serializes Message into bytes.
func (em Message) Serialize() ([]byte, error) {
	msg, err:= json.Marshal(em)
	return msg, err
}

// MessageNumber returns message type number.
func (em Message) MessageNumber() int32 {
	return 1001
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
	msg := tao.MessageFromContext(ctx).(Message)

	conn.Write(msg)

	onePi := new(pi.Pi)
	onePi.Info = new(pi.Info)
	onePi.Info.Name = msg.Content.Name
	onePi.Info.ConnectNow = msg.Content.ConnectNow

	server := <-pi.PiServer
	server.AddPi(onePi)
	pi.PiServer <- server
}

