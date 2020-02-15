package modules

import (
	"fmt"
	grpcServer "go.7yes.com/go/components/grpc/server"
	httpServer "go.7yes.com/go/components/http/server"
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
		fmt.Println( s.Config )

		e := s.GetEngine()
		brand.Init(e)
		if os.Getenv("RUNTIME_ENV") != "prod" {
			reflection.Register(e) //正式环境去掉//
		}
		return nil
	}))

	e.RegisterModules(httpServer.HttpCallback(func(s *httpServer.HttpServer) error {
		fmt.Println( s.Config )
		r := s.GetEngine()
		tget.Init( r )
		return nil
	}))
}

func RegisterMqHandel() {
}
