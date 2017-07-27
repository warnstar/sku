package ClientGetFileInfo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/warnstar/holmes"
	"github.com/warnstar/tao"
	"sku/SkuServer/SkuRun"
	"sku/WebServer/WebKey"
	"sku/Channel/ChanWeb"
	"sku/SkuServer/Tsi"
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
	return 1101
}

// MessageType returns message type .
func (em Message) MessageType() string {
	return "client_get_file_info"
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

	thisPi, err := server.GetPiByConnId(connId)
	if err != nil {
		holmes.Errorln("client-time-sync: 当前链接对应的pi不存在")
		return
	}

	msg := tao.MessageFromContext(ctx).(Message)
	fileName := msg.Content.(string)

	msg.Content, err = Tsi.GetFileInfo(fileName)

	if err != nil {
		holmes.Errorf("获取文件信息错误：%v(%v)\n",err.Error(), fileName)
		return
	}

	//将文件发送至pi
	err = conn.Write(msg)

	if err != nil {
		holmes.Errorf("发送文件消息失败：%v\n",err.Error())
	}

	//通知浏览器-pi已经获取文件信息
	ChanWeb.SendWebLog(WebKey.LOG_TYPE_CLIENT, fmt.Sprintf("%v已获取文件信息", thisPi.Info.Name))
}
