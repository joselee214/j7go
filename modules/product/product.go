package product

import (
	"go.7yes.com/go/proto/product"
	"google.golang.org/grpc"
)

func Init(g *grpc.Server) {
	s := &MemberCardTplServer{}
	product.RegisterMemberCardTplServerServer(g, s)

	d := &DepositCardTplServer{}
	product.RegisterDepositCardTplServerServer(g, d)

	p := &packageCourseSrv{}
	product.RegisterPackageCourseTplServerServer(g, p)

	pcs := &personalCourseSrv{}
	product.RegisterPersonalCourseSrvServer(g,pcs)

	tcs := &TeamCourseTplServer{}
	product.RegisterTeamCourseSrvServer(g,tcs)
}
