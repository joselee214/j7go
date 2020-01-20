package shopService

import (
	"context"
	"go.7yes.com/go/proto/shop"
	"time"
	"j7go/models/shop"
	"j7go/utils"
)

func SaveShopArea(ctx context.Context, s *shop.ShopAreaRequest) (id uint, err error) {
	shopArea := &shopModel.ShopArea{}

	if s.ShopAreaId != utils.IntZero {
		shopArea, err = shopModel.ShopAreaByID(ctx, uint(s.ShopAreaId))
		if err != nil {
			return utils.IntZero, err
		}
	}

	shopArea.BrandID = uint(s.BrandId)
	shopArea.ShopID = uint(s.ShopId)
	shopArea.AreaName = s.AreaName
	shopArea.ContainNumber = uint(s.ContainNumber)
	shopArea.IsVip = int8(s.IsVip)
	shopArea.CreatedTime = uint(time.Now().Unix())
	shopArea.UpdatedTime = uint(time.Now().Unix())
	err = shopArea.Save(ctx)
	if err != nil {
		return utils.IntZero, err
	}

	return shopArea.ID, nil
}

func DelShopArea(ctx context.Context, s *shop.ShopAreaRequest) (id uint, err error) {
	shopArea, err := shopModel.ShopAreaByID(ctx, uint(s.ShopAreaId))
	if err != nil {
		return utils.IntZero, err
	}

	shopArea.IsDel = utils.DELETED
	err = shopArea.Save(ctx)
	if err != nil {
		return utils.IntZero, err
	}

	return shopArea.ID, nil
}

func GetShopArea(ctx context.Context, s *shop.ShopAreaRequest) (shopArea *shopModel.ShopArea, err error) {
	shopArea, err = shopModel.ShopAreaByID(ctx, uint(s.ShopAreaId))
	if err != nil {
		return nil, err
	}

	return shopArea, nil
}

func SaveShopSeat(ctx context.Context, s *shop.ShopAreaRequest) (id uint, err error) {
	return 0, nil
}
