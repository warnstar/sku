package WebServer

import (
	"fmt"
	"github.com/leesper/holmes"
	"golang.org/x/net/websocket"
	"net/http"
	"sku/WebServer/WebMessages"
	"sku/base/config"
)

func Run() {
	http.Handle("/", websocket.Handler(WebMessages.Register))

	host := config.Ini.String("webSocket.host")
	port := config.Ini.String("webSocket.port")
	addr := fmt.Sprintf("%s:%s", host, port)

	holmes.Infof("WebSocket Server Start, net tcp addr %s \n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		holmes.Errorf("ListenAndServe %s\n", err.Error())
	}
}
