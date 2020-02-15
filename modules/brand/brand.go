package brand

import (
	"fmt"
	ferrors "go.7yes.com/j7f/components/errors"
	"j7go/errors"
	"j7go/proto/brand"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"j7go/services/brand"
	"j7go/utils"
)

type BrandService struct{}

func Init(g *grpc.Server) {
	s := &BrandService{}
	fmt.Println("=========>brand.go :: ",g)
	brand.RegisterBrandServerServer(g, s)
}

func (s *BrandService) GetShopList(srv brand.BrandServer_GetShopListServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("brand", zap.String("get_shop_list_receive", err.Error()))
			return err
		}

		res, err := brandService.GetShopList(srv.Context(), uint(params.BrandId))
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("brand", zap.Any("get_shop_list", err.Error()))
			err = ferrors.NewFromCode(business_errors.BrandError_GET_SHOP_LIST_ERROR)
		}

		res.Status = ferrors.GetResHeader(err)

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("brand", zap.String("get_shop_list_send", err.Error()))
			return err
		}

	}
}
