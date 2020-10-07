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
	E *service_register.EtcdCli
}

func NewRegister(e *Engine) *Register {
	r := &Register{}
	if e.Opts.EtcdConfig != nil {
		r.E , _ = service_register.NewEtcd(e.Opts.EtcdConfig)
	}
	return r
}

func (r *Register) Register(e *Engine, index int) error {
	if e.Opts.EtcdConfig == nil {
		return nil
	}

	var nodeServiceInfo []*service_register.ServiceInfo

	node := service_register.NewNode(e.Opts.ServerConfig[index].NodeId,
		e.Opts.ServerConfig[index].Ip,
		strings.Join(util.GetLocalIps(),","),
		e.Opts.ServerConfig[index].Version,
		e.Opts.ServerConfig[index].Port,
		map[string]string{} )  //"protocol": e.Opts.ServerConfig[index].Protocol,"ip": strings.Join(util.GetLocalIps(),",")

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

	ttlOption := service_register.NewTTLOption(
		time.Duration(e.Opts.ServiceConfig.Heartbeat)*time.Second,
		time.Duration(e.Opts.ServiceConfig.Ttl)*time.Second)

	key := e.Opts.ServiceConfig.Key + "_" + e.Opts.ServerConfig[index].Protocol + "/" + e.Opts.ServerConfig[index].Version + "/" + e.Opts.ServerConfig[index].NodeId

	data := &service_register.Service{
		Key:   key,
		Value: string(jsonNode),
		TTL:   ttlOption,
	}

	e.GraceSrv[index].Rr = *service_register.NewRegisterOpts(data, r.E)

	return e.GraceSrv[index].Rr.Register()
}

