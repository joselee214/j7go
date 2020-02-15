package shop

import (
	"go.7yes.com/j7f/proto/shop"
	"google.golang.org/grpc"
)

//初始化，注册门店相关服务
func Init(g *grpc.Server) {
	s := &ShopService{}
	shop.RegisterShopServerServer(g, s)

	sa := &ShopAreaService{}
	shop.RegisterShopAreaServerServer(g, sa)
}
