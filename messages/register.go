package messages

import (
	"sku/messages/demoMessage"
	"github.com/leesper/tao"
	"sku/messages/clientInfo"
)

func init() {

	tao.Register(demoMessage.Message{}.MessageNumber(), demoMessage.DeserializeMessage, demoMessage.ProcessMessage)


	//上报pi信息
	tao.Register(clientInfo.Message{}.MessageNumber(), clientInfo.DeserializeMessage, clientInfo.ProcessMessage)
}