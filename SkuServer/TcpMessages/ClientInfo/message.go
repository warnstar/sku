package ClientInfo

import (
	"context"

	"github.com/leesper/tao"
	"encoding/json"
	"github.com/leesper/holmes"
	"time"
	"sku/SkuServer/SkuPi"
	"sku/SkuServer/SkuRun"
	"sku/SkuServer/TcpMessages/ClientTimeSync"
	"sku/Channel/ChanWebTcp"
	"sku/WebServer/WebKey"
)

// Message defines the echo message.
type Message struct {
	Type string `json:"type"`
	Content struct{
		Name string
		ConnectNow int
	} `json:"content"`
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

// MessageType returns message type .
func (em Message) MessageType() string {
	return "client_info"
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

	//pi信息
	onePi := new(SkuPi.Pi)
	onePi.ConnId = tao.NetIDFromContext(ctx)
	onePi.Info = new(SkuPi.Info)
	onePi.Info.Name = msg.Content.Name
	onePi.Info.ConnectNow = msg.Content.ConnectNow

	server := <-SkuRun.PiServer

	//上报pi信息
	server.AddPi(*onePi)

	//发起时间同步
	timeSyncMsg := ClientTimeSync.Message{}
	timeSyncMsg.Type = ClientTimeSync.Message{}.MessageType()
	timeSyncMsg.Content = time.Now().Unix()
	conn.Write(timeSyncMsg)


	if server.CheckPiAllConnected() {
		holmes.Infoln("全部客户端已连接")

		ChanWebTcp.SendWebLog(WebKey.LOG_TYPE_SERVER,"全部客户端已经连接")
	}


	SkuRun.PiServer <- server
}

