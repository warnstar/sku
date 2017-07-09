package tcpServer

import (
	"github.com/leesper/tao"
)

// Message defines the echo message.
type Message struct {
	Content string
}

// Serialize serializes Message into bytes.
func (em Message) Serialize() ([]byte, error) {
	return []byte(em.Content), nil
}

// MessageNumber returns message type number.
func (em Message) MessageNumber() int32 {
	return 1
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	println("==========DeserializeMessage=========data: " + string(data))
	if data == nil {
		return nil, tao.ErrNilData
	}
	msg := string(data)
	echo := Message{
		Content: msg,
	}
	return echo, nil
}
