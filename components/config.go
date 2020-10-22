package components

import (
	"github.com/fsnotify/fsnotify"
	"github.com/joselee214/j7f/components/config"
	"github.com/joselee214/j7f/components/dao"
	"github.com/joselee214/j7f/components/grpc/server"
	"github.com/joselee214/j7f/components/log"
	"github.com/joselee214/j7f/components/service_register"
	"go.uber.org/zap"
)

type Options struct {
	Config           *config.Configer
	HotResetChan     chan int
	LogConfig        *log.Config              `mapstructure:"log"`
	ServiceConfig    *ServiceConfig           `mapstructure:"service"`
	GrpcClientConfig *GrpcClientConfig        `mapstructure:"grpc_client"`
	EtcdConfig       *service_register.Config `mapstructure:"etcd"`
	ServerConfig     ServerConfig             `mapstructure:"server"`
	RedisConfig      *RedisConfig             `mapstructure:"redis"`
	DBConfig         *dao.DBConfig            `mapstructure:"db"`
	GrpcStreamConfig *server.Config           `mapstructure:"grpc_stream"`
	NsqConfig        *NsqConfig               `mapstructure:"nsq"`
	MQConfig         *MQConfig                `mapstructure:"mq"`
	MongoConfig         *MongoConfig          `mapstructure:"mongodb"`
}

var C *config.Configer

func InitConfig(cfgPath string) (*Options, error) {
	c, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	c.SetConfigFile(cfgPath)
	c.AutomaticEnv() // read in environment variables that match

	err = c.ReadInConfig()
	if err != nil {
		return nil, err
	}

	opts := &Options{
		Config:       c,
		HotResetChan: make(chan int, 1),
	}

	C = opts.Config

	err = c.Unmarshal(opts)
	if err != nil {
		return nil, err
	}

	c.WatchConfig()
	c.OnConfigChange(opts.hotReset)
	return opts, nil
}

func (o *Options) hotReset(e fsnotify.Event) {
	if e.Op == fsnotify.Write || e.Op == fsnotify.Create {
		o.HotResetChan <- 1
		err := o.Config.Unmarshal(o)
		if err != nil {
			L.Panic("faild unmarshal config", zap.Error(err))
		}

		err = ResetLog(o.LogConfig)
		if err != nil {
			L.Panic("faild init log config", zap.Error(err))
		}
	}
}
