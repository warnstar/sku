package Tsi

import (
	"net"
	"sku/WebServer/WebKey"
	"sku/Channel/ChanWebTcp"
)

type Message struct {
	Type string
	Content string
}

type TsiConnect struct {
	Conn *net.TCPConn
	IsRunning bool
	Type string
}

// 共享变量 - tsi 链接对象
var TsiClientChan = make(chan *TsiConnect, 1)

// 通讯变量 - TSI控制信号
var ToSendChan = make(chan *Message)

// 共享变量 - tsi数据接收处理时的暂存变量
var TsiRunningStatusChan = make(chan *TsiRunningStatus, 1)

func resetChan() {
	//重置 共享变量 - tsi数据接收处理时的暂存变量
	<- TsiRunningStatusChan
	TsiRunningStatusChan <- new(TsiRunningStatus)
}

func init() {
	TsiClientChan <- new(TsiConnect)
	TsiRunningStatusChan <- new(TsiRunningStatus)
}

func ControlTsi(msgType string, msgContent string) {
	msg := Message{msgType,msgContent}
	ToSendChan <- &msg
}

func Connect() {
	addr, err := net.ResolveTCPAddr("tcp", "172.16.15.214:3602")
	if err !=nil {
		println(err.Error())
		ChanWebTcp.SendWeb(WebKey.WEB_TSI_CHECK,WebKey.FAIL)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err !=nil {
		println(err.Error())
		ChanWebTcp.SendWeb(WebKey.WEB_TSI_CHECK,WebKey.FAIL)
	}

	tsiConn := <-TsiClientChan
	tsiConn.Conn = conn
	tsiConn.IsRunning = true

	TsiClientChan <- tsiConn

	//启动客户端发送线程
	sendRunning(conn)

	//启动客户端接收线程
	recvRunning(conn)
}