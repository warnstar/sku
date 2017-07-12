package skuServer

import (
	"github.com/leesper/tao"
	"github.com/leesper/holmes"
	"net"
	"os/signal"
	"syscall"
	"runtime"
	"os"
)

// EchoServer represents the echo server.
type EchoServer struct {
	*tao.Server
}

// NewEchoServer returns an EchoServer.
func NewEchoServer() *EchoServer {
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

	return &EchoServer{
		tao.NewServer(codec, onConnect, onClose, onError, onMessage),
	}
}

func Run() {
	defer holmes.Start().Stop()
	runtime.GOMAXPROCS(runtime.NumCPU())


	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		holmes.Fatalf("listen error %v", err)
	}
	echoServer := NewEchoServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		echoServer.Stop()
	}()

	echoServer.Start(l)
}