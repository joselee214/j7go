package components

import (
	"github.com/joselee214/j7f/components/grpc/interceptor"
	"github.com/joselee214/j7f/components/grpc/server"
	"google.golang.org/grpc"
	"net"
	"time"
)

type GrpcClientConfig struct {
	Addr          string
	ServiceName   string
	ServiceMethod string
	TimeOut       time.Duration
}

func NewGrpcServer(serverConfig *NodeServerConfig, grpcOpts ...grpc.ServerOption) (*server.GrpcServer, error) {
	tcpAddr := new(net.TCPAddr)
	tcpAddr.IP = net.ParseIP(serverConfig.Ip)
	tcpAddr.Port = serverConfig.Port
	s,err := server.NewGrpcServer(tcpAddr, grpcOpts...)
	if err==nil{
		s.Config = map[string]interface{}{"modules":serverConfig.EnableModules}
	}
	return s,err
}

func NewGrpcClient(grpcClientConfig *GrpcClientConfig) error {

	return nil
}

func RegisterInterceptors(s Server,e *Engine) {
	s.RegisterStreamInterceptors(
		interceptor.StreamServerErrorInterceptor(L),
		interceptor.StreamServerTraceInterceptor(L, e.Opts.GrpcStreamConfig),
	)
	s.RegisterUnaryInterceptors(
		interceptor.UnaryServerErrorInterceptor(L),
		//TODO:: unaryCall trace_id 拦截器
	)
}