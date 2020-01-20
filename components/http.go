package components

import (
	"go.7yes.com/go/components/http/middleware"
	"go.7yes.com/go/components/http/server"
	"net"
	"os"
)

func NewHttpServ(serverConfig *NodeServerConfig) (*server.HttpServer, error) {
	tcpAddr := new(net.TCPAddr)
	tcpAddr.IP = net.ParseIP(serverConfig.Ip)
	tcpAddr.Port = serverConfig.Port
	env := os.Getenv("RUNTIME_ENV")
	return server.NewHttpServer(tcpAddr, L, env)
}

func RegisterMiddleware(e *Engine) {
	e.RegisterUnaryInterceptors(
		middleware.Logger(L),
		middleware.JsonParam(),
		middleware.Recovery(L),
		middleware.Error(),
	)
}
