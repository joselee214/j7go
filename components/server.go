package components

import (
	"fmt"
	"go.7yes.com/go/components/grace"
	"go.7yes.com/go/components/service_register"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
)

const RAND_PORT = 0

type ServerConfig []*NodeServerConfig

type NodeServerConfig struct {
	Ip       string
	Port     int
	Protocol string
	NodeId  string
	Version string
}

type Server interface {
	RegisterStreamInterceptors(...interface{})
	RegisterUnaryInterceptors(...interface{})
	RegisterCb(...interface{})
	NewServ() error
	StartServ() error
	Stop()
	GracefulStop()
	GetServicesInfo() map[string]service_register.ServerInfo
	GetAddress() *net.TCPAddr
	GetListener() *net.TCPListener
}

func NewServer(e *Engine, grpcOpts ...grpc.ServerOption) error {
	var server Server
	var err error

	for index, serverConfig := range e.Opts.ServerConfig {
		protocol := strings.ToUpper(serverConfig.Protocol)
		switch protocol {
		case "HTTP":
			server, err = NewHttpServ(serverConfig)
			if err != nil {
				return fmt.Errorf("new http server err: %s", err)
			}
			e.Server = append(e.Server, server)
			RegisterMiddleware(e)
		case "GRPC":
			server, err = NewGrpcServer(serverConfig, grpcOpts...)
			if err != nil {
				return fmt.Errorf("new grpc server err: %s", err)
			}
			e.Server = append(e.Server, server)
			RegisterInterceptors(e)
		default:
			return fmt.Errorf("server protocol set err")
		}

		e.Opts.ServerConfig[index].Protocol = protocol

		if serverConfig.Port == RAND_PORT {
			lis := server.GetListener()
			addr := lis.Addr().String()
			_, portString, err := net.SplitHostPort(addr)
			if err != nil {
				return err
			}
			port, err := strconv.Atoi(portString)
			if err != nil {
				return err
			}
			e.Opts.ServerConfig[index].Port = port
		}

		if serverConfig.NodeId  == "" {
			//TODO 这里需要随机生成 nodeId
			fmt.Println("====zzz===>", e.Opts.ServerConfig[index].NodeId )
		}


		e.GraceSrv = append(e.GraceSrv, grace.NewServer(server)) //grace use to hot reload
	}
	return nil
}
