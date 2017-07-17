package pi

import "fmt"


//全局服务器变量
var PiServer = make(chan *Server, 1)

type Server struct {
	TsiServerAddress string
	IsAllConnected bool
	IsAllTimeSync bool
	PiNum int
	Pis [] *Pi
}

func init() {
	PiServer <- new(Server)
}

func (s *Server) AddPi(pi *Pi) {
	if len(s.Pis) == 0 {
		//此时无pi链接
		s.Pis = append(s.Pis, pi)
	} else {
		//此时已有pi链接
		isNew := true
		for _, onePi := range s.Pis {
			//此pi已存在
			if onePi.Info.Name == pi.Info.Name {
				onePi = pi
				isNew = false
				break
			}
		}

		//新pi
		if isNew {
			s.Pis = append(s.Pis, pi)
		}
	}
}

func (s *Server) PrintInfo() {
	fmt.Printf("总链接的Pi数量：%d\n", len(s.Pis))

	for _, onePi := range s.Pis {
		fmt.Printf("%v \n", onePi.Info)
	}
}
