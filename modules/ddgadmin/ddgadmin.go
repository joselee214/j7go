package ddgadmin

import (
	//ferrors "github.com/joselee214/j7f/components/errors"
	"google.golang.org/grpc"
	"j7go/proto/ddgadmin"
	//"j7go/utils"
	//"go.uber.org/zap"
	context "golang.org/x/net/context"
)

type DdgAdminService struct {}

func Init(g *grpc.Server)  {
	s := &DdgAdminService{}
	ddgadmin.RegisterDdgAdminFrontServer(g,s)
}

func (s *DdgAdminService) ValidateUserAccountAndPwd(ctx context.Context, in *ddgadmin.ValidateRequest) (*ddgadmin.ValidateResponse, error) {

	res := &ddgadmin.ValidateResponse{}

	println("==================")
	println(in.Account)
	println(in.Password)

	return  res,nil
}
