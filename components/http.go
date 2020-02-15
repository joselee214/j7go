package components

import (
	"github.com/joselee214/j7f/components/http/middleware"
	"github.com/joselee214/j7f/components/http/server"
	"net"
	"os"
)

func NewHttpServ(serverConfig *NodeServerConfig) (*server.HttpServer, error) {
	tcpAddr := new(net.TCPAddr)
	tcpAddr.IP = net.ParseIP(serverConfig.Ip)
	tcpAddr.Port = serverConfig.Port
	env := os.Getenv("RUNTIME_ENV")
	s,err := server.NewHttpServer(tcpAddr, L, env)
	if err==nil{
		s.Config = map[string]interface{}{"modules":serverConfig.EnableModules}
	}
	return s,err
}

func RegisterMiddleware(s Server) {
	s.RegisterUnaryInterceptors(
		middleware.Logger(L),
		middleware.JsonParam(),
		middleware.Recovery(L),
		middleware.Error(),
	)
}
