package SkuRun

import (
	"github.com/leesper/holmes"
	"github.com/leesper/tao"
)

func SendToAllPi(toPiMsg tao.Message) {

	//获取tcp服务器对象
	tcpServer := <-PiServer
	PiServer <- tcpServer

	for _, thisPi := range tcpServer.Pis {
		err := (*(thisPi.ConnWriter)).Write(toPiMsg)
		if err != nil {
			holmes.Errorln(err.Error(),thisPi.Info)
		}
	}

}