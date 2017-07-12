package messages

import (
	"sku/messages/demoMessage"
	"github.com/leesper/tao"
)

func init() {

	tao.Register(demoMessage.Message{}.MessageNumber(), demoMessage.DeserializeMessage, demoMessage.ProcessMessage)
}