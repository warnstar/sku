package SkuRun

import (
	"errors"
	"sku/SkuServer/SkuPi"
	"github.com/leesper/holmes"
	"fmt"
	"sku/base/config"
)

type Server struct {
	TsiServerAddress string
	IsAllConnected   bool
	IsAllTimeSync    bool
	IsAllWriteKb	 bool
	IsAllSendResult  bool
	PiCurNum         int
	PiMaxNum         int
	Pis              []SkuPi.Pi
}

// 共享变量 - 全局服务器变量
var PiServer = make(chan *Server, 1)

func init() {

	s := Server{PiMaxNum: 2, TsiServerAddress: config.Ini.String("tsi.host")}
	PiServer <- &s
}

func ResetServer() {
	//取出服务器全局变量
	<- PiServer

	//关闭所有pi的链接
	//for _,oldPis := range oldPiServer.Pis {
	//	oldPis.ConnWriter.Close()
	//}
	//holmes.Infoln("重置Tcp服务器")

	s := Server{PiMaxNum: 2, TsiServerAddress: config.Ini.String("tsi.host")}

	PiServer <- &s

}

func (s *Server) AddPi(pi SkuPi.Pi) {
	//此时已有pi链接
	isNew := true
	for _, onePi := range s.Pis {
		//此pi已存在
		if onePi.ConnId == pi.ConnId {
			onePi = pi
			isNew = false
			break
		}
	}

	//新pi
	if isNew {
		s.Pis = append(s.Pis, pi)
		s.PiCurNum++
		if s.PiCurNum >= s.PiMaxNum {
			s.IsAllConnected = true
		}
	}
}

func (s *Server) GetPiByConnId(connId int64) (pi SkuPi.Pi, err error) {
	for _, onePi := range s.Pis {

		//此pi已存在
		if onePi.ConnId == connId {
			return onePi, nil
			break
		}
	}

	return pi, errors.New("object not found")
}

func (s *Server) UpdatePiByConnId(connId int64, pi SkuPi.Pi) {
	for index, onePi := range s.Pis {
		if onePi.ConnId == connId {
			s.Pis[index] = pi
			break
		}
	}
}

func GetClientTree() (treeData [] ClientTreeNode) {
	s := <- PiServer

	for k,thisPi := range s.Pis {
		if thisPi.Ctx.Err() != nil {
			holmes.Debugln(thisPi,thisPi.Ctx.Err())

			// 去除已断开的链接
			s.Pis = append(s.Pis[:k], s.Pis[k+1:]...)
			s.PiCurNum--
			continue
		} else {
			oneNode := new(ClientTreeNode)
			oneNode.Fd = thisPi.ConnId
			oneNode.Label = thisPi.Info.Name

			connModules := fmt.Sprintf("模块数：%v",thisPi.Info.ConnectNow)
			isTimeSync := fmt.Sprintf("时间同步：%v",thisPi.IsTimeSync)

			oneNode.Children = append(oneNode.Children, ClientTreeNode{connModules,0,nil})
			oneNode.Children = append(oneNode.Children, ClientTreeNode{isTimeSync,0,nil})

			treeData = append(treeData, *oneNode)
		}
	}

	PiServer <- s

	return treeData
}

func (s *Server) PrintInfo() {
	holmes.Infof("总链接的Pi数量：%d\n", len(s.Pis))
	for _, onePi := range s.Pis {
		holmes.Infof("%v \n", onePi.Info)
	}
}
