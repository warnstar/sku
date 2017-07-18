package main

import (
	"sku/SkuServer"
	"sku/WebServer"
	"github.com/leesper/holmes"
	"runtime"
)

func main() {
	defer holmes.Start().Stop()
	runtime.GOMAXPROCS(runtime.NumCPU())

	var endRunning = make(chan bool, 1)

	go func(){
		SkuServer.Run()
		endRunning <- true
	}()

	go func(){
		WebServer.Run()
		endRunning <- true
	}()

	<-endRunning
	return
}
