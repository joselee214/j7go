package ddgadmin

import (
	"fmt"
	//ferrors "github.com/joselee214/j7f/components/errors"
	"google.golang.org/grpc"
	"j7go/models/ddg"
	"j7go/proto/ddgadmin"
	//"j7go/utils"
	//"go.uber.org/zap"
	context "golang.org/x/net/context"
)

type DdgAdminGrpcSrv struct {}

func Init(g *grpc.Server)  {
	s := &DdgAdminGrpcSrv{}
	ddgadmin.RegisterDdgAdminFrontServer(g,s)
}

func (s *DdgAdminGrpcSrv) ValidateUserAccountAndPwd(ctx context.Context, in *ddgadmin.ValidateRequest) (*ddgadmin.ValidateResponse, error) {

	res := &ddgadmin.ValidateResponse{}

	println("==================")
	println(in.Account)
	println(in.Password)

	data,err := ddg.DdgAdminUserGroupsByGid(ctx,1)

	if err != nil {
		fmt.Println("======x=>",err)
	}

	for k, v := range data {
		fmt.Println(k)
		fmt.Println(v)
		res.Msg = "获得数据"
		res.Success = 123
	}

	return  res,nil
}
