package components

import (
	"fmt"
	"go.7yes.com/go/components/grace"
	"go.7yes.com/go/components/lock"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var E *Engine

type Engine struct {
	Opts      *Options
	Server    []Server
	GraceSrv  []*grace.Server
	Register  *Register
	RedisLock *lock.RedisLock
}

type RegisterModules func(e *Engine)

func NewEngine(cfgPath string, grpcOpts ...grpc.ServerOption) (*Engine, error) {
	var err error

	e := &Engine{}
	e.Opts, err = InitConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	err = InitLog(e.Opts.LogConfig)
	if err != nil {
		return nil, err
	}

	err = InitDB(e.Opts.DBConfig)
	if err != nil {
		return nil, err
	}

	//err = InitNsq(e.Opts.NsqConfig)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = InitMq(e.Opts.MQConfig)
	//if err != nil {
	//	return nil, err
	//}

	err = NewRedis(e.Opts.RedisConfig)
	if err != nil {
		return nil, err
	}

	err = NewServer(e, grpcOpts...)
	if err != nil {
		return nil, err
	}

	e.Register = NewRegister()


	err = NewGrpcClient(e.Opts.GrpcClientConfig)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Engine) StartServ() error {
	for _, gsrv := range e.GraceSrv {
		err := gsrv.ListenAndServe()
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) RegisterModules(f ...interface{}) {
	for _, srv := range e.Server {
		srv.RegisterCb(f...)
	}
}

func (e *Engine) RegisterStreamInterceptors(f ...interface{}) {
	for _, srv := range e.Server {
		srv.RegisterStreamInterceptors(f...)
	}
}

func (e *Engine) RegisterUnaryInterceptors(f ...interface{}) {
	for _, srv := range e.Server {
		srv.RegisterUnaryInterceptors(f...)
	}
}

func (e *Engine) Run() {
	for index, server := range e.Server {
		go func(index int, server Server) {
			err := server.NewServ()
			if err != nil {
				panic(fmt.Errorf("new %s server err %s", e.Opts.ServerConfig[index].Protocol, err))
			}

			err = e.Register.Register(e, index)
			if err != nil {
				panic(fmt.Errorf("register service err %s", err))
			}

			L.Info(e.Opts.ServerConfig[index].Protocol+" SERVICE START",
				zap.String("addr", e.Opts.ServerConfig[index].Ip),
				zap.Int("port", e.Opts.ServerConfig[index].Port),
				zap.String("service_key", e.Opts.ServiceConfig.Key+"_"+e.Opts.ServerConfig[index].Protocol),
				zap.String("service_node_id", e.Opts.ServerConfig[index].NodeId),
				zap.String("service_node_version", e.Opts.ServerConfig[index].Version))

			err = e.GraceSrv[index].ListenAndServe()
			if err != nil {
				panic(fmt.Errorf("start service err %s", err))
			}
		}(index, server)
	}
}
