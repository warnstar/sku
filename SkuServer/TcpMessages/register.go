package TcpMessages

import (
	"github.com/leesper/tao"
	"sku/SkuServer/TcpMessages/ClientInfo"
	"sku/SkuServer/TcpMessages/ClientTimeSync"
	"sku/SkuServer/TcpMessages/ClientTsiTestPreStart"
	"sku/SkuServer/TcpMessages/ClientTsiTestPreStop"
	"sku/SkuServer/TcpMessages/ClientTsiTestStart"
	"sku/SkuServer/TcpMessages/ClientTsiTestStop"
	"sku/SkuServer/TcpMessages/ClientTsiTestModuleResult"
	"sku/SkuServer/TcpMessages/ClientModuleDac"
	"sku/SkuServer/TcpMessages/ClientWriteKbFinish"
)

func init() {

	/**
	========== 与 pi 交互===========
	*/
	//上报pi信息
	tao.Register(ClientInfo.Message{}.MessageNumber(), ClientInfo.DeserializeMessage, ClientInfo.ProcessMessage)

	//时间同步确认
	tao.Register(ClientTimeSync.Message{}.MessageNumber(), ClientTimeSync.DeserializeMessage, ClientTimeSync.ProcessMessage)

	//客户端TSI校验启动
	tao.Register(ClientTsiTestPreStart.Message{}.MessageNumber(), ClientTsiTestPreStart.DeserializeMessage, ClientTsiTestPreStart.ProcessMessage)

	//客户端TSI校验关闭
	tao.Register(ClientTsiTestPreStop.Message{}.MessageNumber(), ClientTsiTestPreStop.DeserializeMessage, ClientTsiTestPreStop.ProcessMessage)

	//客户端TSI测试启动
	tao.Register(ClientTsiTestStart.Message{}.MessageNumber(), ClientTsiTestStart.DeserializeMessage, ClientTsiTestStart.ProcessMessage)

	//客户端TSI测试关闭
	tao.Register(ClientTsiTestStop.Message{}.MessageNumber(), ClientTsiTestStop.DeserializeMessage, ClientTsiTestStop.ProcessMessage)

	//客户端上报DAC
	tao.Register(ClientModuleDac.Message{}.MessageNumber(), ClientModuleDac.DeserializeMessage, ClientModuleDac.ProcessMessage)

	//客户端写入KB完成
	tao.Register(ClientWriteKbFinish.Message{}.MessageNumber(), ClientWriteKbFinish.DeserializeMessage, ClientWriteKbFinish.ProcessMessage)

	//客户端上报Tsi测试分析结果
	tao.Register(ClientTsiTestModuleResult.Message{}.MessageNumber(), ClientTsiTestModuleResult.DeserializeMessage, ClientTsiTestModuleResult.ProcessMessage)

}
