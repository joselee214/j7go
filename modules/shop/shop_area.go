package shop

import (
	"j7go/proto/shop"
	"go.uber.org/zap"
	"j7go/utils"
	"j7go/services/shop"
	"github.com/joselee214/j7f/components/errors"
	"github.com/joselee214/j7f/proto/common"
)

type ShopAreaService struct {
}

func (s *ShopAreaService) AddShopArea(srv shop.ShopAreaServer_AddShopAreaServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.String("add_receive", err.Error()))
			return err
		}

		shopAreaId, err := shopService.SaveShopArea(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.Any("add_shop_area", err.Error()))
		}

		res := &shop.ShopAreaIdResponse{
			Status: errors.GetResHeader(err),
			ShopAreaId: uint32(shopAreaId),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.String("add_send", err.Error()))
			return err
		}
	}
}

func (s *ShopAreaService) EditShopArea(srv shop.ShopAreaServer_EditShopAreaServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.String("add_shop_area", err.Error()))
			return err
		}

		shopAreaId, err := shopService.SaveShopArea(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.Any("add_shop_area", err.Error()))
		}

		res := &shop.ShopAreaIdResponse{
			Status: errors.GetResHeader(err),
			ShopAreaId: uint32(shopAreaId),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.String("add_send", err.Error()))
			return err
		}
	}
}

func (s *ShopAreaService) DelShopArea(srv shop.ShopAreaServer_DelShopAreaServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.String("add_shop_area", err.Error()))
			return err
		}

		shopAreaId, err := shopService.DelShopArea(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.Any("del_shop_area", err.Error()))
		}

		res := &shop.ShopAreaIdResponse{
			Status: errors.GetResHeader(err),
			ShopAreaId: uint32(shopAreaId),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.String("add_send", err.Error()))
			return err
		}
	}
}

func (s *ShopAreaService) GetShopArea(srv shop.ShopAreaServer_GetShopAreaServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.String("add_shop_area", err.Error()))
			return err
		}

		shopArea, err := shopService.GetShopArea(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.Any("del_shop_area", err.Error()))
		}

		res := &shop.ShopAreaResponse{
			Status: errors.GetResHeader(err),
			ShopAreaId: uint32(shopArea.ID),
			BrandId:uint32(shopArea.BrandID),
			ShopId:uint32(shopArea.ShopID),
			AreaName:shopArea.AreaName,
			ContainNumber:uint32(shopArea.ContainNumber),
			IsVip:uint32(shopArea.IsVip),
			IsDel:common.DelStatus(shopArea.IsDel),
			CreatedTime:uint32(shopArea.CreatedTime),
			UpdatedTime:uint32(shopArea.UpdatedTime),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("shop_area", zap.String("add_send", err.Error()))
			return err
		}
	}
}
