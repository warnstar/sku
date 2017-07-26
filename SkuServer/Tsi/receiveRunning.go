package Tsi

import (
	"bufio"
	"net"
	"sku/SkuServer/TcpMessages/TcpKey"
	"sku/WebServer/WebKey"
	"strconv"
	"strings"
	"sku/Channel/ChanWeb"
	"sku/Channel/ChanTcp"
	"github.com/leesper/holmes"
)


/**
	tsi数据接收线程
 */
func recvRunning(conn net.Conn) {
	go func() {
		reader := bufio.NewReader(conn)
		holmes.Infoln("TSI 数据接收--启动线程")

		for {
			tsiChan := <- TsiClientChan
			TsiClientChan <- tsiChan

			if !tsiChan.IsStartedReceive {
				holmes.Infoln("TSI 数据接收--结束线程")
				break
			}

			msg, err := reader.ReadString('\n')
			println(msg)
			if err != nil {
				holmes.Errorf("TSI 数据接收 -- read数据错误：%s\n",err.Error())
			}
			res := strings.Split(msg, ",")

			if len(res) >= 2 {
				tsiStr := res[1]
				tsiStr = strings.Replace(tsiStr, "\n", "", -1)

				tsiNumFloat, err := strconv.ParseFloat(tsiStr, 32)
				if err != nil {
					holmes.Errorf("TSI数据接收：%s\n",err.Error())
				}
				tsiNum := int(tsiNumFloat * 1000)
				analysisTsi(tsiNum)
			}
		}
	}()
}

type TsiRunningStatus struct {
	IsRunning    bool
	AtRunningNum int

	AtStartNum int
	AtStopNum  int
}

/**
	tsi数据分析
 */
func analysisTsi(pm25 int) {
	tsiChan := <-TsiClientChan
	tsiChan.RecvNum++
	TsiClientChan <- tsiChan

	tsiRunStatus := <-TsiRunningStatusChan

	if tsiChan.Type == TSI_RUN_TYPE_CHECK {
		if pm25 > 0 {
			if tsiChan.RecvNum >= 3 {
				ChanWeb.SendWeb(WebKey.WEB_TSI_CHECK, WebKey.SUCCESS)

				// 关闭数据接收
				ControlTsi(TSI_SERVER_STOP, "")
			}
		}
	} else {
		//发送当前tsi值到浏览器
		ChanWeb.SendWeb(WebKey.WEB_TSI_NOW_DATA, pm25)

		if tsiRunStatus.IsRunning {
			if pm25 >= TSI_START_POINT {
				// pm25 >= 500
				// 废弃
			} else if pm25 >= TSI_STOP_POINT {
				// 14 < pm25 < 500
				if tsiRunStatus.AtStartNum >= TSI_FLAG_TIMES {
					// 有效数据--处理

					//将数据写入文件
					SaveFile(tsiChan.Type, pm25)
				} else {
					tsiRunStatus.AtStartNum++
					if tsiRunStatus.AtStartNum >= TSI_FLAG_TIMES {
						//发送日志 到浏览器
						ChanWeb.SendWebLog(WebKey.LOG_TYPE_TSI, "开始收集数据任务")

						// 通知 服务器，开始数据收集任务
						if tsiChan.Type == TSI_RUN_TYPE_TEST_PRE {
							ChanTcp.SendTcp(TcpKey.TYPE_CLIENT_TSI_TEST_PRE_START, "")
						} else {
							ChanTcp.SendTcp(TcpKey.TYPE_CLIENT_TSI_TEST_START, "")
						}
					}
				}
			} else if pm25 > 0 {
				// 0 < pm25 < 14
				if tsiRunStatus.AtStopNum >= TSI_FLAG_TIMES {
					// 停止采集TSI
					ControlTsi(TSI_SERVER_STOP, "")

					// 通知 服务器，收集数据任务完成
					if tsiChan.Type == TSI_RUN_TYPE_TEST_PRE {
						ChanTcp.SendTcp(TcpKey.TYPE_CLIENT_TSI_TEST_PRE_STOP, "")

						//告知浏览器 tsi校准已完毕
						ChanWeb.SendWeb(WebKey.WEB_TSI_TEST_PRE,"")
					} else {
						ChanTcp.SendTcp(TcpKey.TYPE_CLIENT_TSI_TEST_STOP, "")

						//告知浏览器 tsi测试已完毕
						ChanWeb.SendWeb(WebKey.WEB_TSI_TEST,"")
					}

					//发送日志 到浏览器
					ChanWeb.SendWebLog(WebKey.LOG_TYPE_TSI, "处理完成, 关闭收集数据任务")
				} else {
					tsiRunStatus.AtStopNum++
				}
			} else {
				// 0 <= pm25
			}
		} else {
			if pm25 >= TSI_START_POINT {
				if tsiRunStatus.AtRunningNum >= TSI_FLAG_TIMES {
					//开始准备处理数据
					tsiRunStatus.IsRunning = true

					//发送日志 到浏览器
					ChanWeb.SendWebLog(WebKey.LOG_TYPE_TSI, "pm25已超过500，准备开始采集数据")
				} else {
					tsiRunStatus.AtRunningNum++
				}
			}
		}
	}

	TsiRunningStatusChan <- tsiRunStatus
}
