package Tsi

import (
	"net"
	"sku/WebServer/WebKey"
	"github.com/warnstar/holmes"
	"sku/Channel/ChanWeb"
	"sku/SkuServer/SkuRun"
)

type Message struct {
	Type    string
	Content string
}

type TsiConnect struct {
	Conn      *net.TCPConn
	IsRunning bool
	IsStartedReceive bool
	Type      string
	RecvNum 	int
}

// 共享变量 - tsi 链接对象
var TsiClientChan = make(chan *TsiConnect, 1)

// 通讯变量 - TSI控制信号
var ToSendChan = make(chan *Message)

// 共享变量 - tsi数据接收处理时的暂存变量
var TsiRunningStatusChan = make(chan *TsiRunningStatus, 1)

func resetChan() {
	//重置 共享变量 - tsi数据接收处理时的暂存变量
	<-TsiRunningStatusChan
	TsiRunningStatusChan <- new(TsiRunningStatus)
}

func init() {
	TsiClientChan <- new(TsiConnect)
	TsiRunningStatusChan <- new(TsiRunningStatus)
}


func ControlTsi(msgType string, msgContent string) {
	msg := Message{msgType, msgContent}

	ToSendChan <- &msg
}

func Connect() bool {
	server := <-SkuRun.PiServer
	SkuRun.PiServer <- server

	addr, err := net.ResolveTCPAddr("tcp", server.TsiServerAddress + ":3602")
	if err != nil {
		holmes.Debugf("Tsi 服务器连接失败:%s\n" ,err.Error())
		ChanWeb.SendWeb(WebKey.WEB_TSI_CHECK, WebKey.FAIL)
		return false
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		holmes.Debugf("Tsi 服务器连接失败:%s\n" ,err.Error())
		ChanWeb.SendWeb(WebKey.WEB_TSI_CHECK, WebKey.FAIL)
		return false
	}

	tsiConn := <-TsiClientChan
	tsiConn.Conn = conn
	tsiConn.IsRunning = true

	TsiClientChan <- tsiConn

	//启动客户端发送线程
	sendRunning(conn)

	return true
}
