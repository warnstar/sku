package SkuRun

import (
	"errors"
	"fmt"
	"sku/SkuServer/SkuPi"
)

type Server struct {
	TsiServerAddress string
	IsAllConnected   bool
	IsAllTimeSync    bool
	PiCurNum         int
	PiMaxNum         int
	Pis              []SkuPi.Pi
}

// 共享变量 - 全局服务器变量
var PiServer = make(chan *Server, 1)

func init() {
	s := Server{PiMaxNum: 2, TsiServerAddress: "172.16.15.214"}
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

func (s *Server) CheckPiAllTimeSync() bool {
	timeSyncNum := 0
	for _, onePi := range s.Pis {
		if onePi.IsTimeSync {
			timeSyncNum++
		}
	}

	if timeSyncNum >= s.PiMaxNum {
		s.IsAllTimeSync = true

		return true
	}

	return false
}

func (s *Server) CheckPiAllConnected() bool {
	return s.PiCurNum >= s.PiMaxNum
}

func (s *Server) PrintInfo() {
	fmt.Printf("总链接的Pi数量：%d\n", len(s.Pis))

	for _, onePi := range s.Pis {
		fmt.Printf("%v \n", onePi.Info)
	}
}
