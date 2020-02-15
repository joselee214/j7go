package modules

import (
	"fmt"
	grpcServer "go.7yes.com/j7f/components/grpc/server"
	httpServer "go.7yes.com/j7f/components/http/server"
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
		fmt.Println( "config grpc...", s.Config )

		gs := s.GetEngine()
		brand.Init(gs)
		if os.Getenv("RUNTIME_ENV") != "prod" {
			reflection.Register(gs) //正式环境去掉//
		}
		return nil
	}))

	e.RegisterModules(httpServer.HttpCallback(func(s *httpServer.HttpServer) error {
		fmt.Println( "config http...", s.Config )
		r := s.GetEngine()
		tget.Init( r )
		return nil
	}))
}

func RegisterMqHandel() {
}
