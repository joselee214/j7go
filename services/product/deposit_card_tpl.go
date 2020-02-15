package product

import (
	"context"
	"go.7yes.com/j7f/proto/product"
	"time"
	"j7go/components"
	productModel "j7go/models/product"
	"j7go/models/shop"
	"j7go/services/images"
	"j7go/utils"
)

func AddDepositCardTpl(ctx context.Context, request *product.AddAndUpdateDepositCardTplRequest) (Id uint, error error) {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return utils.IntZero, err
	}
	//创建相册获取ID,并保存卡背景
	imageList := make([]*imagesService.ImageInfo, len(request.BgImage))
	for index, image := range request.BgImage {
		imageInfo := &imagesService.ImageInfo{
			ImageId:   image.ImageId,
			ImageUrl:  image.ImgUrl,
			CoverType: int8(image.CoverType),
		}
		imageList[index] = imageInfo
	}

	albumId, err := imagesService.InsertImages(ctx, utils.IntZero, imageList)

	//新增储值卡信息
	cardId, err := AddDeposit(ctx, request, uint(albumId))
	if err != nil {
		_ = components.M.Rollback(ctx)
		return utils.IntZero, err
	}

	//新增卡支持售卖门店
	if product.SupportSales_SALES_PART == request.SupportSales {
		err = SaveSaleShop(ctx, uint(request.BrandId), cardId, int8(product.ProductType_PRODUCT_DEPOSIT_CARD), request.SellShopList)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return utils.IntZero, err
		}
	}

	//新增卡支持使用门店
	if product.ConsumptionRange_CONSUMPTION_MUCH_SHOP == request.ConsumptionRange {
		err = SaveSupportShop(ctx, uint(request.BrandId), cardId, int8(product.ProductType_PRODUCT_DEPOSIT_CARD), request.ConsumerShopList)

		if err != nil {
			_ = components.M.Rollback(ctx)
			return utils.IntZero, err
		}
	}

	//新增消费类目
	err = saveConsumerIds(ctx, uint(request.BrandId), cardId, request.CardConsumerId)
	if err != nil {
		_ = components.M.Rollback(ctx)
		return utils.IntZero, err
	}
	err = components.M.Commit(ctx)
	if err != nil {
		return utils.IntZero, err
	}
	return cardId, nil
}

//编辑储值卡
func UpdateDepositCardTpl(ctx context.Context, request *product.AddAndUpdateDepositCardTplRequest) (error error) {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return err
	}

	cardTpl, err := productModel.DepositCardTemplateByID(ctx, uint(request.CardId))
	if err != nil {
		return err
	}
	//更新储值卡信息
	cardTpl.CardPrice = uint(request.CardPrice)
	cardTpl.SellPrice = uint(request.SellPrice)
	cardTpl.Num = uint(request.Num)
	cardTpl.Unit = int8(request.Unit)
	cardTpl.StartTime = uint(request.StartTime)
	cardTpl.EndTime = uint(request.EndTime)
	cardTpl.SellType = int8(request.CardSellType)
	cardTpl.CardContents = request.CardContents
	cardTpl.Description = request.Description
	cardTpl.ConsumptionRange = int8(request.ConsumptionRange)
	cardTpl.SupportSales = int8(request.SupportSales)
	err = cardTpl.Update(ctx)
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}

	//更新支持售卖门店 如果是自定义门店则需要更新
	if product.ConsumptionRange_CONSUMPTION_MUCH_SHOP == request.ConsumptionRange {
		err = UpdateSaleShop(ctx, cardTpl.BrandID, cardTpl.ID, int8(product.ProductType_PRODUCT_DEPOSIT_CARD), request.SellShopList)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//更新支持使用门店
	if product.SupportSales_SALES_PART == request.SupportSales {
		err = UpdateSupportShop(ctx, cardTpl.BrandID, cardTpl.ID, int8(product.ProductType_PRODUCT_DEPOSIT_CARD), request.ConsumerShopList)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}
	//更新消费类目
	err = productModel.DeportConsumerDeleteByCardId(ctx, cardTpl.ID)
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}
	err = saveConsumerIds(ctx, cardTpl.BrandID, cardTpl.ID, request.CardConsumerId)
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}
	//更新卡背景
	imageList := make([]*imagesService.ImageInfo, len(request.BgImage))
	for index, image := range request.BgImage {
		imageInfo := &imagesService.ImageInfo{
			ImageId:   image.ImageId,
			ImageUrl:  image.ImgUrl,
			CoverType: int8(image.CoverType),
		}
		imageList[index] = imageInfo
	}
	err = imagesService.UpdateAlbumImages(ctx, utils.IntZero, imageList)
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}
	//卡操作日志 方法封装
	return components.M.Commit(ctx)
}

//执行添加储值卡
func AddDeposit(ctx context.Context, request *product.AddAndUpdateDepositCardTplRequest, albumId uint) (cardId uint, error error) {
	cardTpl := &productModel.DepositCardTemplate{}
	cardTpl.BrandID = uint(request.BrandId)
	cardTpl.ShopID = uint(request.ShopId)
	cardTpl.CardName = request.CardName
	cardTpl.SellPrice = uint(request.SellPrice)
	cardTpl.CardPrice = uint(request.CardPrice)
	cardTpl.Unit = int8(request.Unit)
	cardTpl.Num = uint(request.Num)
	cardTpl.SupportSales = int8(request.SupportSales)
	cardTpl.ConsumptionRange = int8(request.ConsumptionRange)
	cardTpl.StartTime = uint(request.StartTime)
	cardTpl.EndTime = uint(request.EndTime)
	cardTpl.CardContents = string(request.CardContents)
	cardTpl.AlbumID = uint(albumId)
	if utils.IntZero != request.ShopId {
		cardTpl.PublishChannel = int8(product.PublishChannel_CHANNEL_SHOP)
	} else {
		cardTpl.PublishChannel = int8(product.PublishChannel_CHANNEL_BRAND)
	}
	cardTpl.SellStatus = int8(product.ProductSellStatus_PRODUCT_SELL_STATUS_AVAILABLE)
	cardTpl.SellType = int8(request.CardSellType)
	cardTpl.OperatorID = uint(request.OperatorId)
	cardTpl.Description = string(request.Description)
	cardTpl.UpdatedTime = uint(time.Now().Unix())
	cardTpl.CreatedTime = uint(time.Now().Unix())
	err := cardTpl.Insert(ctx)
	if err != nil {
		return utils.IntZero, err
	}
	return cardTpl.ID, nil
}

//批量添加消费类目
func saveConsumerIds(ctx context.Context, brandId, cardId uint, consumerType []uint32) error {
	list := make([]*productModel.DepositCardConsumer, len(consumerType))
	for index, consumer := range consumerType {
		consumerModel := &productModel.DepositCardConsumer{}
		consumerModel.BrandID = brandId
		consumerModel.CardID = cardId
		consumerModel.ConsumerType = int8(consumer)
		consumerModel.CreatedTime = uint(time.Now().Unix())
		consumerModel.UpdatedTime = uint(time.Now().Unix())

		list[index] = consumerModel
	}
	err := productModel.DepositCardConsumerBatchInsert(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

//GetDepositDetailById
func GetDepositDetailById(ctx context.Context, request *product.GetDepositCardDetailRequest) (*product.GetDepositCardDetailResponse, error) {
	//查询储值卡主表
	cardObj, err := productModel.DepositCardTemplateByID(ctx, uint(request.CardId))
	if err != nil {
		return nil, err
	}

	detailResponse := &product.GetDepositCardDetailResponse{
		CardName:     cardObj.CardName,
		StartTime:    uint32(cardObj.StartTime),
		EndTime:      uint32(cardObj.EndTime),
		SellPrice:    uint32(cardObj.SellPrice),
		SellStatus:   product.ProductSellStatus(cardObj.SellStatus),
		SupportSales: product.SupportSales(cardObj.SupportSales),
		CardPrice:    uint32(cardObj.CardPrice),
		CardContents: cardObj.CardContents,
		Description:  cardObj.Description,
	}

	//查询背景图片
	imgLists, err := imagesService.GetImages(ctx, uint(cardObj.AlbumID), utils.IntZero)
	if err != nil {
		return nil, err
	}
	images := make([]*product.Image, len(imgLists))
	for index, img := range imgLists {
		i := &product.Image{
			ImageId:   uint32(img.ID),
			CoverType: product.ImageCoverType(img.CoverType),
			ImgUrl:    img.ImageURL,
		}
		images[index] = i
	}
	detailResponse.CardBg = images

	//查询支持售卖门店列表
	if int8(product.SupportSales_SALES_PART) == cardObj.SupportSales {
		detailResponse.SupportShop, err = GetSaleShop(ctx, cardObj.ID, int8(product.ProductType_PRODUCT_DEPOSIT_CARD))
		if err != nil {
			return nil, err
		}
	}

	//查询使用门店列表 单店和全店不查
	if int8(product.ConsumptionRange_CONSUMPTION_MUCH_SHOP) == cardObj.ConsumptionRange {
		detailResponse.CanUseShop, err = GetSupportShop(ctx, cardObj.ID, int8(product.ProductType_PRODUCT_DEPOSIT_CARD))
		if err != nil {
			return nil, err
		}
	}

	//获取上架门店列表
	detailResponse.OnlineShop, err = GetOnLineShop(ctx, cardObj.ID, int8(product.ProductType_PRODUCT_DEPOSIT_CARD))
	if err != nil {
		return nil, err
	}
	detailResponse.OnLineShopNum = uint32(len(detailResponse.OnlineShop))

	//获取品牌下所有门店数量
	shopLists, err := shopModel.ShopByBrandID(ctx, cardObj.BrandID)
	if err != nil {
		return nil, err
	}
	detailResponse.TotalShopNum = uint32(len(shopLists))
	//查询消费类目列表
	detailResponse.ConsumerType, err = GetConsumerType(ctx, uint(cardObj.ID))
	if err != nil {
		return nil, err
	}

	return detailResponse, nil
}

//获取消费类目
func GetConsumerType(ctx context.Context, cardId uint) ([]product.ConsumerType, error) {
	consumerList, err := productModel.DepositCardConsumerByCardId(ctx, cardId)
	if err != nil {
		return nil, err
	}
	consumerTypes := make([]product.ConsumerType, len(consumerList))
	for index, v := range consumerList {
		consumerTypes[index] = product.ConsumerType(v.ConsumerType)
	}
	return consumerTypes, nil
}

//停售
func StopSellDepositCard(ctx context.Context, request *product.StopSellDepositCardRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return err
	}

	//更改储值卡主表
	c, err := productModel.DepositCardTemplateByID(ctx, uint(request.CardId))
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}

	c.ID = uint(request.CardId)
	c.SellStatus = int8(product.ProductSellStatus_PRODUCT_SELL_STATUS_UNAVAILABLE)
	c.UpdatedTime = uint(time.Now().Unix())
	err = c.Update(ctx)
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}
	//todo 记录停售操作和原因

	//下架所有储值卡
	err = productModel.ShopShelfCardBatchDown(ctx, uint(c.ID), int8(product.ProductType_PRODUCT_DEPOSIT_CARD), int8(product.ShelfStatus_SHELFSTATUS_DOWN), int8(product.ShelfStatus_SHELFSTATUS_DOWN))
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}
	return components.M.Commit(ctx)
}

//重新售卖
func RestartSellDepositCard(ctx context.Context, request *product.RestartSellDepositCardRequest) error {
	c, err := productModel.DepositCardTemplateByID(ctx, uint(request.CardId))
	if err != nil {
		return err
	}

	if request.EndTime != 0 && request.StartTime != 0 {
		c.StartTime = uint(request.StartTime)
		c.EndTime = uint(request.EndTime)
	}
	c.SellStatus = int8(product.ProductSellStatus_PRODUCT_SELL_STATUS_AVAILABLE)
	c.UpdatedTime = uint(time.Now().Unix())
	err = c.Update(ctx)
	if err != nil {
		return err
	}
	//todo 记录停售操作和原因

	return nil
}

//上架储值卡
func OnlineSellDepositCard(ctx context.Context, request *product.OnlineSellDepositCardRequest) error {
	if utils.IntZero == request.ShopId { //批量上架
		err := productModel.ShopShelfCardBatchDown(ctx, uint(request.CardId), int8(product.ProductType_PRODUCT_DEPOSIT_CARD), int8(product.ShelfStatus_SHELFSTATUS_UP))
		if err != nil {
			return err
		}
	} else { //单个门店上架
		err := productModel.ShopShelfCardDown(ctx, uint(request.CardId), int8(product.ProductType_PRODUCT_DEPOSIT_CARD), uint(request.ShopId), int8(product.ShelfStatus_SHELFSTATUS_UP))
		if err != nil {
			return err
		}
	}
	return nil
}

//下架储值卡(品牌+门店)
func OfflineSellDepositCard(ctx context.Context, request *product.OfflineSellDepositCardRequest) error {
	if utils.IntZero == request.ShopId { //批量下架
		err := productModel.ShopShelfCardBatchDown(ctx, uint(request.CardId), int8(product.ProductType_PRODUCT_DEPOSIT_CARD), int8(product.ShelfStatus_SHELFSTATUS_DOWN))
		if err != nil {
			return err
		}
	} else { //单个门店下架
		err := productModel.ShopShelfCardDown(ctx, uint(request.CardId), int8(product.ProductType_PRODUCT_DEPOSIT_CARD), uint(request.ShopId), int8(product.ShelfStatus_SHELFSTATUS_DOWN))
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteSellDepositCard(ctx context.Context, request *product.DeleteDepositCardRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return err
	}
	//删除更新当前卡
	c, err := productModel.DepositCardTemplateByID(ctx, uint(request.CardId))
	if err != nil {
		return err
	}
	c.IsDel = utils.DELETED
	c.SellStatus = int8(product.ShelfStatus_SHELFSTATUS_DOWN)
	c.UpdatedTime = uint(time.Now().Unix())
	err = c.Update(ctx)
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}

	//下架所有在售门店
	err = productModel.ShopShelfCardBatchDown(ctx, c.ID, int8(product.ProductType_PRODUCT_DEPOSIT_CARD), int8(product.ShelfStatus_SHELFSTATUS_DOWN))
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}

	return components.M.Commit(ctx)
}
