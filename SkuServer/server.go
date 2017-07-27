package SkuServer

import (
	"fmt"
	"github.com/warnstar/holmes"
	"github.com/warnstar/tao"
	"net"
	"os"
	"os/signal"
	_ "sku/SkuServer/TcpMessages"
	"sku/base/config"
	"syscall"
	"sku/Channel/ChanWeb"
	"sku/WebServer/WebKey"
	"sku/SkuServer/SkuRun"
)

// SkuServer represents the Sku server.
type SkuServer struct {
	*tao.Server
}

// NewSkuServer returns an SkuServer.
func NewSkuServer() *SkuServer {
	onConnect := tao.OnConnectOption(func(conn tao.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onClose := tao.OnCloseOption(func(conn tao.WriteCloser) {
		holmes.Infoln("closing client")

		//发送pi树形数据到浏览器
		ChanWeb.SendWeb(WebKey.WEB_CLIENT_TREE_DATA,SkuRun.GetClientTree())
	})

	onError := tao.OnErrorOption(func(conn tao.WriteCloser) {
		holmes.Infoln("on error")

		//发送pi树形数据到浏览器
		ChanWeb.SendWeb(WebKey.WEB_CLIENT_TREE_DATA,SkuRun.GetClientTree())
	})

	onMessage := tao.OnMessageOption(func(msg tao.Message, conn tao.WriteCloser) {
		holmes.Infoln("receving message")
	})

	codec := tao.CustomCodecOption(StringValue{})

	return &SkuServer{
		tao.NewServer(codec, onConnect, onClose, onError, onMessage),
	}
}

func Run() {
	addr := fmt.Sprintf(":%s", config.Ini.String("tcp.port"))
	l, err := net.Listen("tcp", addr)
	if err != nil {
		holmes.Fatalf("listen error %v %v", err, addr)
	}
	skuServer := NewSkuServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		skuServer.Stop()
	}()

	skuServer.Start(l)
}
