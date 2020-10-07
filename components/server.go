package components

import (
	"fmt"
	"github.com/joselee214/j7f/components/grace"
	"github.com/joselee214/j7f/components/service_register"
	"github.com/joselee214/j7f/util"
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
	EnableModules []string
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
		//case "APPLICATION":
		//	server, err = NewApplication(serverConfig)
		//	if err != nil {
		//		return fmt.Errorf("new application err: %s", err)
		//	}
		//	RegisterAppModules(server)
		//	e.Server = append(e.Server, server)
		case "HTTP":
			server, err = NewHttpServ(serverConfig)
			if err != nil {
				return fmt.Errorf("new http server err: %s", err)
			}
			RegisterMiddleware(server)
			e.Server = append(e.Server, server)
		case "GRPC":
			server, err = NewGrpcServer(serverConfig, grpcOpts...)
			if err != nil {
				return fmt.Errorf("new grpc server err: %s", err)
			}
			RegisterInterceptors(server,e)
			e.Server = append(e.Server, server)
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
			//e.Opts.ServerConfig[index].NodeId = strings.Replace(s, `\`, `\\\`, -1)
			e.Opts.ServerConfig[index].NodeId =  strings.Join(util.GetLocalIps(),",") + ":" + strconv.Itoa(e.Opts.ServerConfig[index].Port)        //NewNid()
		}

		e.GraceSrv = append(e.GraceSrv, grace.NewServer(server)) //grace use to hot reload
		e.Register = append(e.Register, NewRegister(e) )
	}
	return nil
}

//func NewNid() string {
//	t := time.Now().String()
//	timeStr := t[2:26]
//	tStr := strings.Replace(strings.Replace(strings.Replace(strings.Replace(timeStr, `:`, ``, -1), `.`, ``, -1), ` `, ``, -1), `-`, ``, -1)
//	uuidbyte,_ := uuid.NewUUID()
//	return tStr+"-"+uuidbyte.String()
//}