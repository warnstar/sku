package Tsi

import (
	"github.com/warnstar/holmes"
	"net"
	"time"
)

const CONTROL_SIGNAL_CLOSE = "control_signal_close"
const CONTROL_SIGNAL_NULL  = "control_signal_null"

func sendRunning(conn net.Conn) {
	go func() {
		isContinue := true
		isReceiveDataContinue := true

		holmes.Infoln("TSI 操作信号 - 启动线程")
		for {
			if isContinue {
				sendMsg := <-ToSendChan

				go func(){
					switch sendMsg.Type {
					case CONTROL_SIGNAL_NULL:
					case CONTROL_SIGNAL_CLOSE:
					case TSI_SERVER_EXIT:
						isReceiveDataContinue = false
						isContinue = false

						ControlTsi(CONTROL_SIGNAL_NULL,"")


						_, err := conn.Write([]byte(TSI_STOP))
						if err != nil {
							holmes.Errorf("TSI 服务器 关闭失败：%v\n",err.Error())
						} else {
							holmes.Infoln("TSI 服务器 关闭成功")

							//断开与tsi服务器连接
							err := conn.Close()
							if err != nil {
								holmes.Errorf("TSI 会话 关闭失败：%v\n",err.Error())
							} else {
								holmes.Infoln("TSI 会话 关闭成功")
							}
						}

						<-TsiClientChan
						TsiClientChan <- new(TsiConnect)
					case TSI_SERVER_START:
						//重置运行时状态变量
						resetChan()

						_, err := conn.Write([]byte(TSI_START))
						if err != nil {
							holmes.Errorf("TSI 服务器启动失败：%v\n",err.Error())
						} else {
							holmes.Infoln("TSI 服务器 启动成功")
						}
					case TSI_SERVER_STOP:
						//结束 tsi数据接收请求线程
						ControlTsi(TSI_SERVER_RECEIVE_DATA_STOP, "")

						//关闭tsi服务器
						_, err := conn.Write([]byte(TSI_STOP))
						if err != nil {
							holmes.Errorf("TSI 服务器关闭失败：%v\n", err.Error())

							//关闭连接
							ControlTsi(TSI_SERVER_EXIT, "")
						} else {
							holmes.Infoln("TSI 服务器 关闭成功")
						}
					case TSI_SERVER_RECEIVE_DATA_START:
						//开启请求接收tsi数据线程
						go func() {
							isReceiveDataContinue = true
							holmes.Infoln("TSI 发送接收请求 - 启动线程")

							tsiClient := <-TsiClientChan
							tsiClient.RecvNum = 0
							tsiClient.IsStartedReceive = true
							TsiClientChan <- tsiClient

							//启动 TSI数据接收线程
							recvRunning(conn)

							for {
								if isReceiveDataContinue {
									//每隔1s发送一次请求命令
									_, err := conn.Write([]byte(TSI_RECEIVE))
									if err != nil {
										holmes.Errorln(err.Error())
										isReceiveDataContinue = false
										//关闭连接
										ControlTsi(TSI_SERVER_EXIT, "")
									}
									time.Sleep(time.Second)
								} else {
									tsiClient := <-TsiClientChan
									tsiClient.RecvNum = 0
									tsiClient.IsStartedReceive = false
									TsiClientChan <- tsiClient

									holmes.Infoln("TSI 发送接收请求 - 结束线程")
									break
								}
							}
						}()
					case TSI_SERVER_RECEIVE_DATA_STOP:
						isReceiveDataContinue = false
					}
				}()

			} else {
				break
			}
		}

		holmes.Infoln("TSI 操作信号 - 结束线程")
	}()

}
