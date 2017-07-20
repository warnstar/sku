package Tsi

import (
	"github.com/leesper/holmes"
	"net"
	"time"
)

func sendRunning(conn net.Conn) {
	go func() {
		isContinue := true
		isReceiveDataContinue := true

		for {
			if isContinue {
				sendMsg := <-ToSendChan

				switch sendMsg.Type {
				case TSI_SERVER_EXIT:
					_, err := conn.Write([]byte(TSI_STOP))
					println("关闭TSI连接")
					if err != nil {
						println(err.Error())
					} else {
						//断开与tsi服务器连接
						conn.Close()
					}

					tsiConn := <-TsiClientChan
					tsiConn.IsRunning = false
					tsiConn.Conn = nil
					TsiClientChan <- tsiConn

					isContinue = false
				case TSI_SERVER_START:
					//重置运行时状态变量
					resetChan()

					_, err := conn.Write([]byte(TSI_START))
					if err != nil {
						holmes.Errorln(err.Error())
					} else {
						holmes.Infoln("启动TSI成功")
					}
				case TSI_SERVER_STOP:
					//关闭数据接收
					ControlTsi(TSI_SERVER_RECEIVE_DATA_STOP, "")

					_, err := conn.Write([]byte(TSI_STOP))
					if err != nil {
						println(err.Error())
						//关闭连接
						ControlTsi(TSI_SERVER_EXIT, "")
					} else {
						holmes.Infoln("关闭TSI成功")
					}
				case TSI_SERVER_RECEIVE_DATA_START:
					//开启请求接收tsi数据线程
					go func() {
						isReceiveDataContinue = true
						holmes.Infoln("启动TSI数据接收线程")
						for {
							if isReceiveDataContinue {
								//每隔1s发送一次请求命令
								_, err := conn.Write([]byte(TSI_RECEIVE))
								if err != nil {
									println(err.Error())
									//关闭连接
									ControlTsi(TSI_SERVER_EXIT, "")
								}
								time.Sleep(time.Second)
							}
						}
					}()
				case TSI_SERVER_RECEIVE_DATA_STOP:
					isReceiveDataContinue = false
				}
			}
		}
	}()

}
