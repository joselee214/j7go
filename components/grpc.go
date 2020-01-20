package components

import (
	"go.7yes.com/go/components/grpc/interceptor"
	"go.7yes.com/go/components/grpc/server"
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
	return server.NewGrpcServer(tcpAddr, grpcOpts...)
}

func NewGrpcClient(grpcClientConfig *GrpcClientConfig) error {

	return nil
}

func RegisterInterceptors(e *Engine) {
	e.RegisterStreamInterceptors(
		interceptor.StreamServerErrorInterceptor(L),
		interceptor.StreamServerTraceInterceptor(L, e.Opts.GrpcStreamConfig),
	)
	e.RegisterUnaryInterceptors(
		interceptor.UnaryServerErrorInterceptor(L),
		//TODO:: unaryCall trace_id 拦截器
	)
}
