package SkuServer

import (
	"github.com/leesper/tao"
	"github.com/leesper/holmes"
	"net"
	"os/signal"
	"syscall"
	"os"
	_ "sku/SkuServer/TcpMessages"
	"fmt"
	"sku/base/config"
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
	})

	onError := tao.OnErrorOption(func(conn tao.WriteCloser) {
		holmes.Infoln("on error")
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
		holmes.Fatalf("listen error %v", err)
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