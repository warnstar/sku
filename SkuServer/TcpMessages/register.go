package TcpMessages

import (
	"github.com/leesper/tao"
	"sku/SkuServer/TcpMessages/ClientInfo"
	"sku/SkuServer/TcpMessages/ClientTimeSync"
)

func init() {

	/**
	========== 与 pi 交互===========
	*/
	//上报pi信息
	tao.Register(ClientInfo.Message{}.MessageNumber(), ClientInfo.DeserializeMessage, ClientInfo.ProcessMessage)

	//时间同步确认
	tao.Register(ClientTimeSync.Message{}.MessageNumber(), ClientTimeSync.DeserializeMessage, ClientTimeSync.ProcessMessage)

}
