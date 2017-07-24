package main

import (
	"github.com/leesper/holmes"
	"runtime"
	"sku/SkuServer"
	"sku/WebServer"
)

func main() {
	defer holmes.Start().Stop()
	runtime.GOMAXPROCS(runtime.NumCPU())

	var endRunning = make(chan bool, 1)

	go func() {
		SkuServer.Run()
		println("end tcp")
		endRunning <- true
	}()

	go func() {
		WebServer.Run()
		println("end websocket")
		endRunning <- true
	}()

	<-endRunning
	return
}
