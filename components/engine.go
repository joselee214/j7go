package components

import (
	"fmt"
	"github.com/joselee214/j7f/components/grace"
	"github.com/joselee214/j7f/components/lock"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"sync"
)

var E *Engine

type Engine struct {
	Opts      *Options
	Server    []Server
	GraceSrv  []*grace.Server
	Register  []*Register
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

	if e.Opts.DBConfig !=nil {
		err = InitDB(e.Opts.DBConfig)
		if err != nil {
			return nil, err
		}
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

	if e.Opts.RedisConfig != nil {
		err = NewRedis(e.Opts.RedisConfig)
		if err != nil {
			return nil, err
		}
	}


	if e.Opts.MongoConfig != nil {
		NewMongoPool(e.Opts.MongoConfig)
	}

	if e.Opts.GrpcClientConfig != nil {
		err = NewGrpcClient(e.Opts.GrpcClientConfig)
		if err != nil {
			return nil, err
		}
	}

	err = NewServer(e, grpcOpts...)
	if err != nil {
		return nil, err
	}

	return e, nil
}

//func (e *Engine) StartServ() error {
//	for _, gsrv := range e.GraceSrv {
//		err := gsrv.ListenAndServe()
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (e *Engine) RegisterModules(f ...interface{}) {
	for _, srv := range e.Server {
		srv.RegisterCb(f...)
	}
}
//
//func (e *Engine) RegisterStreamInterceptors(f ...interface{}) {
//	for _, srv := range e.Server {
//		srv.RegisterStreamInterceptors(f...)
//	}
//}
//
//func (e *Engine) RegisterUnaryInterceptors(f ...interface{}) {
//	for _, srv := range e.Server {
//		srv.RegisterUnaryInterceptors(f...)
//	}
//}

func (e *Engine) Run(wg *sync.WaitGroup) {
	for index, server := range e.Server {
		go func(index int, server Server) {

			//fmt.Println("==================RUN")
			//fmt.Println(os.Getpid())

			err := server.NewServ()
			if err != nil {
				panic(fmt.Errorf("new %s server err %s", e.Opts.ServerConfig[index].Protocol, err))
			}

			//fmt.Println("==================e.Opts.ServerConfig")
			//fmt.Println(e.Opts.ServerConfig)

			if e.Opts.ServiceConfig != nil {

				err = e.Register[index].Register(e, index)
				if err != nil {
					panic(fmt.Errorf("register service err %s", err))
				}

				L.Info(e.Opts.ServerConfig[index].Protocol+" SERVICE START",
					zap.Int("local_pid", os.Getpid()),
					zap.String("addr", e.Opts.ServerConfig[index].Ip),
					zap.Int("port", e.Opts.ServerConfig[index].Port),
					zap.String("service_key", e.Opts.ServiceConfig.Key+"_"+e.Opts.ServerConfig[index].Protocol),
					zap.String("service_node_id", e.Opts.ServerConfig[index].NodeId),
					zap.String("service_node_version", e.Opts.ServerConfig[index].Version))
			}

			//fmt.Println("SERVICE START", e.Opts.ServerConfig[index].Ip,e.Opts.ServerConfig[index].Port )

			//defer func() {
			//	if r:= recover();r!=nil{
			//		fmt.Println( r )
			//	}
			//}()
			err = e.GraceSrv[index].ListenAndServe()

			wg.Done()

			//fmt.Println("==================STOP")
			//fmt.Println(err)
			//fmt.Println(wg)
			//fmt.Println("==================STOP")


			L.Info(e.Opts.ServerConfig[index].Protocol+" SERVICE STOP",
				zap.Int("local_pid", os.Getpid()),
				zap.String("addr", e.Opts.ServerConfig[index].Ip),
				zap.Int("port", e.Opts.ServerConfig[index].Port),
				zap.String("service_key", e.Opts.ServiceConfig.Key+"_"+e.Opts.ServerConfig[index].Protocol),
				zap.String("service_node_id", e.Opts.ServerConfig[index].NodeId),
				zap.String("service_node_version", e.Opts.ServerConfig[index].Version))

			//if err != nil {
			//	fmt.Println("start service / end service err %s", err)
			//	//panic(fmt.Errorf("start service / end service err %s", err))
			//}
		}(index, server)
	}
}
