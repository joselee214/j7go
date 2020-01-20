package modules

import (
	"github.com/gin-gonic/gin"
	grpcServer "go.7yes.com/go/components/grpc/server"
	httpServer "go.7yes.com/go/components/http/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"j7go/components"
	"j7go/modules/brand"
	"j7go/modules/tget"
)

func RegisterModules(e *components.Engine) {
	e.RegisterModules(grpcServer.GrpcCallback(func(s *grpc.Server) error {
		//payment.Init(s)
		//shop.Init(s)
		//image.Init(s)
		//staff.Init(s)
		//product.Init(s)
		//member.Init(s)
		brand.Init(s)
		reflection.Register(s)
		return nil
	}))

	e.RegisterModules(httpServer.HttpCallback(func(s *gin.Engine) error {
		tget.Init(s)
		return nil
	}))
}

func RegisterMqHandel() {
}
