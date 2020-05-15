package modules

import (
	"github.com/gin-gonic/gin"
	grpcServer "github.com/joselee214/j7f/components/grpc/server"
	httpServer "github.com/joselee214/j7f/components/http/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"j7go/components"
	"j7go/modules/brand"
	"j7go/modules/tget"
	"os"
)

func RegisterModules(e *components.Engine) {
	e.RegisterModules(grpcServer.GrpcCallback(func(s *grpcServer.GrpcServer) error {
		//payment.Init(s)
		//shop.Init(s)
		//image.Init(s)
		//staff.Init(s)
		//product.Init(s)
		//member.Init(s)

		gs := s.GetEngine()
		brand.Init(gs)
		//fmt.Println( "s.GetServicesInfo()" )
		//fmt.Println( s.GetServicesInfo() )

		explainModules(s.Config,gs)
		if os.Getenv("RUNTIME_ENV") != "prod" {
			reflection.Register(gs) //正式环境去掉//
		}
		return nil
	}))

	e.RegisterModules(httpServer.HttpCallback(func(s *httpServer.HttpServer) error {
		r := s.GetEngine()
		//fmt.Println( "s.GetServicesInfo()" )
		//fmt.Println( s.GetServicesInfo() )
		explainModules(s.Config,r)
		return nil
	}))
}

func explainModules(sconfigs map[string]interface{},s interface{})  {
	var gs *grpc.Server
	var gr *gin.Engine
	if srv,ok := s.(*grpc.Server); ok{
		gs = srv
	}
	if srv,ok := s.(*gin.Engine); ok{
		gr = srv
	}
	for k, v := range sconfigs {
		if k == "modules" {
			if enableModules,ok := v.([]string); ok {
				for _,m := range enableModules {
					switch m {
						case "brand":
							println("GrpcServer enableModules : " + m)
							brand.Init(gs)
						case "tget":
							println("HttpServer enableModules : " + m)
							tget.Init(gr)
					}
				}
			}
		}
	}
}



func RegisterMqHandel() {
}