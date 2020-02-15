package components

import (
	"encoding/json"
	"github.com/joselee214/j7f/components/service_register"
	"github.com/joselee214/j7f/util"
	"strings"
	"time"
)

type ServiceConfig struct {
	Ttl       int
	Key       string
	Heartbeat int
}

type Register struct {
	R *service_register.RegisterOpts
}

func NewRegister() *Register {
	return &Register{}
}

func (r *Register) Register(e *Engine, index int) error {
	var nodeServiceInfo []*service_register.ServiceInfo

	node := service_register.NewNode(e.Opts.ServerConfig[index].NodeId,
		e.Opts.ServerConfig[index].Ip,
		e.Opts.ServerConfig[index].Port,
		map[string]string{"protocol": e.Opts.ServerConfig[index].Protocol,"ips": strings.Join(util.GetLocalIps(),",")})

	serviceInfo := e.Server[index].GetServicesInfo()
	for k, v := range serviceInfo {
		nodeService := new(service_register.ServiceInfo)
		for _, method := range v.Methods {
			nodeService.Methods = append(nodeService.Methods, method.Name)
			nodeService.Metadata = make(map[string]string)
		}
		nodeServiceInfo = append(nodeServiceInfo, nodeService)
		node.SetServices(k, nodeServiceInfo...)
	}

	jsonNode, err := json.Marshal(node)
	if err != nil {
		return err
	}

	ttlOption := service_register.NewTTLOption(time.Duration(e.Opts.ServiceConfig.Heartbeat)*time.Second,
		time.Duration(e.Opts.ServiceConfig.Ttl)*time.Second)

	key := e.Opts.ServiceConfig.Key + "_" + e.Opts.ServerConfig[index].Protocol + "/" + e.Opts.ServerConfig[index].NodeId

	data := &service_register.Service{
		Key:   key,
		Value: string(jsonNode),
		TTL:   ttlOption,
	}

	etcdCli, err := service_register.NewEtcd(e.Opts.EtcdConfig)
	if err != nil {
		return err
	}

	rr := *service_register.NewRegisterOpts(data, etcdCli)
	return rr.Register()
}

