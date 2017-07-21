package ClientTsiTestModuleResult

import (
	"context"

	"encoding/json"
	"fmt"
	"github.com/leesper/holmes"
	"github.com/leesper/tao"
	"sku/SkuServer/SkuRun"
	"sku/WebServer/WebKey"
	"sku/Channel/ChanWeb"
	"math"
)

// Message defines the echo message.
type Message struct {
	Type    string `json:"type"`
	Content [] Module `json:"content"`
}

type Module struct {
	Status string `json:"status"`
	ModuleId int `json:"module_id"`
	Info [] struct{
		Stage string `json:"stage"`
		Total int `json:"total"`
		Error int `json:"error"`
		Proportion float64 `json:"proportion"`
	} `json:"info"`
}

// Serialize serializes Message into bytes.
func (em Message) Serialize() ([]byte, error) {
	msg, err := json.Marshal(em)
	return msg, err
}

// MessageNumber returns message type number.
func (em Message) MessageNumber() int32 {
	return 1207
}

// MessageType returns message type .
func (em Message) MessageType() string {
	return "client_tsi_test_module_result"
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

	msg := tao.MessageFromContext(ctx).(Message)

	//取tcp服务器 变量
	server := <-SkuRun.PiServer

	thisPi, err := server.GetPiByConnId(connId)

	if err != nil {
		holmes.Errorln("client-time-sync: 当前链接对应的pi不存在")
		return
	}

	//设置pi 上报分析结果
	thisPi.IsSendResult = true
	server.UpdatePiByConnId(connId, thisPi)

	//设置tcp服务器 变量
	SkuRun.PiServer <- server

	//分析的结果数据加工
	for k,module := range msg.Content {
		msg.Content[k].Status = "success"

		for kk,state := range module.Info {
			if state.Error != 0 {
				msg.Content[k].Status = "error"
			}

			var errRate float64
			if state.Total != 0 {
				errRate = float64(state.Error)/float64(state.Total)
			} else {
				errRate = 0
				msg.Content[k].Status = "error"
			}
			errRate = math.Trunc(errRate*1e2 + 0.5)*1e-2
			msg.Content[k].Info[kk].Proportion = errRate * 100
		}
	}

	result := struct {
		Pid string `json:"pid"`
		Modules [] Module `json:"modules"`
	}{thisPi.Info.Name,msg.Content}

	//将结果发送至客户端
	ChanWeb.SendWeb(WebKey.WEB_TSI_TEST_MODULE_RESULT, result)

	//通知浏览器-pi已上传测试结果
	ChanWeb.SendWebLog(WebKey.LOG_TYPE_CLIENT, fmt.Sprintf("%v已分析完毕", thisPi.Info.Name))
}
