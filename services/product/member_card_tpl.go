package product

import (
	"context"
	"github.com/r3labs/diff"
	"github.com/joselee214/j7f/proto/common"
	"j7go/proto/product"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
	"j7go/components"
	"j7go/models/images"
	productModel "j7go/models/product"
	"j7go/models/shop"
	"j7go/services/images"
	"j7go/utils"
)

func AddMemberCardTpl(ctx context.Context, p *product.MemberCardTplRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	//新增会员卡
	c := &productModel.MemberCardTemplate{}
	c.BrandID = uint(p.BrandId)
	c.ShopID = uint(p.ShopId)
	c.CardType = int8(p.CardType)
	c.CardName = p.CardName
	c.PublishChannel = int8(p.PublishChannel)
	c.AdmissionRange = int8(p.AdmissionRange)
	c.PriceSetting = int8(p.PriceSetting)
	c.SupportSales = int8(p.SupportSales)
	c.IsTransfer = int8(p.IsTransfer)
	c.Unit = int8(p.TransferUnit)
	c.Num = uint(p.TransferNum)
	c.SellType = int8(p.SellType)
	c.CardContents = p.CardContents
	c.CardIntroduction = p.CardIntroduction
	c.StartTime = uint(p.StartTime)
	c.EndTime = uint(p.EndTime)
	c.SellStatus = utils.Available
	c.OperatorID = uint(p.OperatorId)
	c.SellStatus = utils.Available
	c.UpdatedTime = uint(time.Now().Unix())
	c.CreatedTime = uint(time.Now().Unix())

	//保存会员卡背景
	image := &imagesService.ImageInfo{
		ImageUrl:  p.CardBg.ImgUrl,
		CoverType: int8(imagesModel.COVER_TYPE_GENERAL),
	}
	albumId, err := imagesService.InsertImages(ctx, utils.IntZero, []*imagesService.ImageInfo{image})
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_card_bg", zap.Any("image", image), zap.Error(err))
		return err
	}

	c.AlbumID = int(albumId)
	err = c.Insert(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_card", zap.Any("card", c), zap.Error(err))
		return err
	}

	//支持使用场馆
	if p.AdmissionRange == product.AdmissionRange_ADMISSION_MUCH_SHOP {
		err = SaveSupportShop(ctx, uint(p.BrandId), uint(c.ID), int8(product.ProductType_PRODUCT_MEMBER_CARD), p.AdmissionShopList)
		if err != nil {
			utils.GetTraceLog(ctx).Error("add_support_shop", zap.Any("shop_ids", p.AdmissionShopList), zap.Error(err))
			return err
		}
	}

	//支持售卖场馆
	if p.SupportSales == product.SupportSales_SALES_PART {
		err = SaveSaleShop(ctx, uint(p.BrandId), uint(c.ID), int8(product.ProductType_PRODUCT_MEMBER_CARD), p.SellShopList)
		if err != nil {
			utils.GetTraceLog(ctx).Error("add_sale_shop", zap.Any("shop_ids", p.SellShopList), zap.Error(err))
			return err
		}
	}

	//价格阶梯
	err = saveCardPriceSetting(ctx, uint(c.ID), p.PriceGradient)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_price_setting", zap.Any("price", p.PriceGradient), zap.Error(err))
		return err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetMemberCardTplInfo(ctx context.Context, p *product.GetMemberCardTplInfoRequest) (*product.GetMemberCardTplInfoResponse, error) {
	cardInfo := &product.GetMemberCardTplInfoResponse{}
	c, err := productModel.MemberCardTemplateByID(ctx, uint(p.CardId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_card", zap.Uint32("card_id", p.CardId), zap.Error(err))
		return cardInfo, err
	}

	//入场门店
	if c.AdmissionRange == int8(product.AdmissionRange_ADMISSION_MUCH_SHOP) {
		cardInfo.AdmissionShopList, err = GetSupportShop(ctx, uint(p.CardId), int8(product.ProductType_PRODUCT_MEMBER_CARD))
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_support_shop", zap.Uint32("card_id", p.CardId), zap.Error(err))
			return cardInfo, err
		}
	}

	//售卖门店
	if c.SupportSales == int8(product.SupportSales_SALES_PART) {
		cardInfo.SellShopList, err = GetSaleShop(ctx, uint(p.CardId), int8(product.ProductType_PRODUCT_MEMBER_CARD))
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_sale_shop", zap.Uint32("card_id", p.CardId), zap.Error(err))
			return cardInfo, err
		}
		cardInfo.TotalShopNum = uint32(len(cardInfo.SellShopList))
	} else {
		shops, err := shopModel.ShopByBrandID(ctx, c.BrandID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_sale_shop", zap.Uint("brand_id", c.BrandID), zap.Error(err))
			return cardInfo, err
		}
		cardInfo.TotalShopNum = uint32(len(shops))
	}

	//价格梯度
	cardInfo.PriceGradient, err = getPriceGradient(ctx, uint(p.CardId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_price_setting", zap.Uint32("card_id", p.CardId), zap.Error(err))
		return cardInfo, err
	}

	//已上架门店数量
	shops, err := productModel.ShopShelfCardByProductId(ctx, uint(p.CardId), int8(product.ProductType_PRODUCT_MEMBER_CARD))
	for _, shop := range shops {
		if shop.ShelfStatus == int8(product.ShelfStatus_SHELFSTATUS_UP) {
			cardInfo.ShelfShopNum += 1
		}
	}

	//获取会员卡背景图
	img, err := imagesService.GetImages(ctx, uint(c.AlbumID), int8(imagesModel.COVER_TYPE_GENERAL))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_card_bg", zap.Int("album_id", c.AlbumID), zap.Error(err))
		return cardInfo, err
	}

	if len(img) != utils.IntZero {
		cardInfo.CardBg = &product.Image{
			ImageId: uint32(img[0].ID),
			ImgUrl:  img[0].ImageURL,
		}
	}
	cardInfo.CardId = uint32(c.ID)
	cardInfo.BrandId = uint32(c.BrandID)
	cardInfo.ShopId = uint32(c.ShopID)
	cardInfo.CardType = product.CardType(c.CardType)
	cardInfo.CardName = c.CardName
	cardInfo.AdmissionRange = product.AdmissionRange(c.AdmissionRange)
	cardInfo.PriceSetting = product.SetPriceType(c.PriceSetting)
	cardInfo.SupportSales = product.SupportSales(c.SupportSales)
	cardInfo.StartTime = uint32(c.StartTime)
	cardInfo.EndTime = uint32(c.EndTime)
	cardInfo.SellStatus = product.ProductSellStatus(c.SellStatus)
	cardInfo.IsTransfer = common.IsAvailable(c.IsTransfer)
	cardInfo.TransferUnit = product.TransferUnit(c.Unit)
	cardInfo.TransferNum = uint32(c.Num)
	cardInfo.SellType = product.CardSellType(c.SellType)
	cardInfo.CardContents = c.CardContents
	cardInfo.CardIntroduction = c.CardIntroduction
	cardInfo.PublishChannel = product.PublishChannel(c.PublishChannel)

	return cardInfo, nil
}

func EditMemberCardTpl(ctx context.Context, p *product.MemberCardTplRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	//更新会员卡
	c, err := productModel.MemberCardTemplateByID(ctx, uint(p.CardId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_card", zap.Uint32("card_id", p.CardId), zap.Error(err))
		return err
	}
	oldAdmissionRange := c.AdmissionRange
	oldSupportSales := c.SupportSales

	c.PublishChannel = int8(p.PublishChannel)
	c.AdmissionRange = int8(p.AdmissionRange)
	c.PriceSetting = int8(p.PriceSetting)
	c.SupportSales = int8(p.SupportSales)
	c.IsTransfer = int8(p.IsTransfer)
	c.Unit = int8(p.TransferUnit)
	c.Num = uint(p.TransferNum)
	c.SellType = int8(p.SellType)
	c.CardContents = p.CardContents
	c.CardIntroduction = p.CardIntroduction
	c.StartTime = uint(p.StartTime)
	c.EndTime = uint(p.EndTime)
	c.OperatorID = uint(p.OperatorId)
	c.UpdatedTime = uint(time.Now().Unix())

	//更新会员卡图片
	image := &imagesService.ImageInfo{
		ImageId:   p.CardBg.ImageId,
		ImageUrl:  p.CardBg.ImgUrl,
		CoverType: int8(imagesModel.COVER_TYPE_GENERAL),
	}
	err = imagesService.UpdateAlbumImages(ctx, uint(c.AlbumID), []*imagesService.ImageInfo{image})
	if err != nil {
		utils.GetTraceLog(ctx).Error("save_card_bg", zap.Int("album_id", c.AlbumID), zap.Any("image", image), zap.Error(err))
		return err
	}

	err = c.Update(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("save_card", zap.Any("card", c), zap.Error(err))
		return err
	}

	//更新支持使用场馆
	if oldAdmissionRange != c.AdmissionRange && product.AdmissionRange_ADMISSION_ALL_SHOP == p.AdmissionRange {
		err = productModel.DeleteAllSupportShopByProductId(ctx, c.ID, int8(product.ProductType_PRODUCT_MEMBER_CARD))
		if err != nil {
			utils.GetTraceLog(ctx).Error("del_support_shop", zap.Uint("card_id", c.ID), zap.Error(err))
			return err
		}
	}

	if product.AdmissionRange_ADMISSION_MUCH_SHOP == p.AdmissionRange {
		err = UpdateSupportShop(ctx, uint(c.BrandID), uint(c.ID), int8(product.ProductType_PRODUCT_MEMBER_CARD), p.AdmissionShopList)
		if err != nil {
			utils.GetTraceLog(ctx).Error("save_sale_shop", zap.Uint("card_id", c.ID),
				zap.Uint("brand_id", c.BrandID), zap.Any("shop_ids", p.AdmissionShopList), zap.Error(err))
			return err
		}
	}

	//更新支持售卖场馆
	if oldSupportSales != c.SupportSales && product.SupportSales_SALES_ALL == p.SupportSales {
		err = productModel.DeleteAllSupportShopByProductId(ctx, c.ID, int8(product.ProductType_PRODUCT_MEMBER_CARD))
		if err != nil {
			utils.GetTraceLog(ctx).Error("del_sale_shop", zap.Uint("card_id", c.ID), zap.Error(err))
			return err
		}
	}

	if p.SupportSales == product.SupportSales_SALES_PART {
		err = UpdateSaleShop(ctx, uint(c.BrandID), uint(c.ID), int8(product.ProductType_PRODUCT_MEMBER_CARD), p.SellShopList)
		if err != nil {
			utils.GetTraceLog(ctx).Error("save_sale_shop", zap.Uint("card_id", c.ID),
				zap.Uint("brand_id", c.BrandID), zap.Any("shop_ids", p.SellShopList), zap.Error(err))
			return err
		}
	}

	//更新价格阶梯
	err = productModel.CardPricesettingDeleteByCardId(ctx, c.ID)
	if err != nil {
		utils.GetTraceLog(ctx).Error("del_price_setting", zap.Uint("card_id", c.ID), zap.Error(err))
		return err
	}

	err = saveCardPriceSetting(ctx, uint(c.ID), p.PriceGradient)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_price_setting", zap.Uint("card_id", c.ID),
			zap.Any("price_setting", p.PriceGradient), zap.Error(err))
		return err
	}

	//todo 记录会员卡更新日志

	err = components.M.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func DelMemberCardTpl(ctx context.Context, p *product.DelMemberCardTplRequest) error {
	c, err := productModel.MemberCardTemplateByID(ctx, uint(p.CardId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_card", zap.Uint32("card_id", p.CardId), zap.Error(err))
		return err
	}

	c.IsDel = utils.Available
	c.OperatorID = uint(p.OperatorId)
	c.UpdatedTime = uint(time.Now().Unix())

	if err = c.Save(ctx); err != nil {
		utils.GetTraceLog(ctx).Error("del_card", zap.Any("card", c), zap.Error(err))
		return err
	}

	//todo 记录会员卡删除日志

	return nil
}

func StopSaleMemberCardTpl(ctx context.Context, p *product.StopSaleMemberCardTplRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	c, err := productModel.MemberCardTemplateByID(ctx, uint(p.CardId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_card", zap.Uint32("card_id", p.CardId), zap.Error(err))
		return err
	}

	c.SellStatus = int8(product.ProductSellStatus_PRODUCT_SELL_STATUS_UNAVAILABLE)
	c.OperatorID = uint(p.OperatorId)
	c.UpdatedTime = uint(time.Now().Unix())
	err = c.Update(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("stop_sale_card", zap.Any("card", c), zap.Error(err))
		return err
	}

	//todo 记录停售原因

	//下架门店会员卡
	err = productModel.ShopShelfCardBatchDown(ctx, uint(c.ID), int8(product.ProductType_PRODUCT_MEMBER_CARD), int8(product.ShelfStatus_SHELFSTATUS_DOWN))
	if err != nil {
		utils.GetTraceLog(ctx).Error("shelf_card_down", zap.Uint32("card_id", p.CardId), zap.Error(err))
		return err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func RecoverSaleMemberCardTpl(ctx context.Context, p *product.RecoverSaleMemberCardTplRequest) error {
	c, err := productModel.MemberCardTemplateByID(ctx, uint(p.CardId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_card", zap.Uint32("card_id", p.CardId), zap.Error(err))
		return err
	}

	c.StartTime = uint(p.StartTime)
	c.EndTime = uint(p.EndTime)
	c.SellStatus = int8(product.ProductSellStatus_PRODUCT_SELL_STATUS_AVAILABLE)
	c.OperatorID = uint(p.OperatorId)
	c.UpdatedTime = uint(time.Now().Unix())
	err = c.Update(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("recover_sale_card", zap.Any("card", c), zap.Error(err))
		return err
	}

	//todo 记录操作日志

	return nil
}

//维护阶梯价格
func saveCardPriceSetting(ctx context.Context, cardId uint, priceGradient []*product.PriceGradient) error {
	priceSettingList := make([]*productModel.CardPriceSetting, len(priceGradient))
	for index, priceInfo := range priceGradient {
		priceSetting := &productModel.CardPriceSetting{}
		priceSetting.CardID = uint(cardId)
		priceSetting.Unit = int8(priceInfo.Unit)
		priceSetting.Num = uint(priceInfo.Num)
		priceSetting.RallyPrice = uint(priceInfo.RallyPrice)
		priceSetting.MinPrice = uint(priceInfo.MinPrice)
		priceSetting.MaxPrice = uint(priceInfo.MaxPrice)
		priceSetting.FrozenDay = uint(priceInfo.FrozenDay)
		priceSetting.GiftUnit = uint(priceInfo.GiftUnit)
		priceSetting.UpdatedTime = uint(time.Now().Unix())
		priceSetting.CreatedTime = uint(time.Now().Unix())

		priceSettingList[index] = priceSetting
	}
	err := productModel.CardPriceSettingBatchInsert(ctx, priceSettingList)
	if err != nil {
		return err
	}
	return nil
}

//获取阶梯价格
func getPriceGradient(ctx context.Context, cardId uint) ([]*product.PriceGradient, error) {
	cardPriceList, err := productModel.CardPriceSettingByCardId(ctx, cardId)
	if err != nil {
		return nil, err
	}

	priceGradient := make([]*product.PriceGradient, len(cardPriceList))
	for index, cardPrice := range cardPriceList {
		price := &product.PriceGradient{}
		price.Id = uint32(cardPrice.ID)
		price.Unit = product.CardUnit(cardPrice.Unit)
		price.Num = uint32(cardPrice.Num)
		price.RallyPrice = uint32(cardPrice.RallyPrice)
		price.MinPrice = uint32(cardPrice.MinPrice)
		price.MaxPrice = uint32(cardPrice.MaxPrice)
		price.GiftUnit = uint32(cardPrice.GiftUnit)
		price.FrozenDay = uint32(cardPrice.FrozenDay)
		priceGradient[index] = price
	}

	return priceGradient, nil
}

//会员卡批量上架
func UpShelfCardTpl(ctx context.Context, p *product.UpShelfCardTplRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("begin_transaction", err.Error()))
		return err
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	cardShopList := p.CardShop

	shopShelfCardList := make([]*productModel.ShopShelfCard, len(p.CardShop))

	for index, cardShop := range cardShopList {
		//主表shop_shelf_card//需要判断唯一索引
		shopShelfCard := &productModel.ShopShelfCard{}
		shopShelfCard.BrandID = uint(p.ShelfBrandId)
		shopShelfCard.ShopID = uint(cardShop.ShelfShopId)
		shopShelfCard.ProductID = uint(cardShop.CardId)
		shopShelfCard.CourseInterests = int8(p.CourseInterests)
		shopShelfCard.ProductType = int8(product.ProductType_PRODUCT_MEMBER_CARD)
		shopShelfCard.ShelfStatus = int8(product.ShelfStatus_SHELFSTATUS_UP)
		shopShelfCard.AdmissionStatus = int8(p.ShopCardAdmissionTime)
		shopShelfCard.IsDel = utils.NOT_DELETED
		shopShelfCard.CreatedTime = uint(time.Now().Unix())
		shopShelfCard.UpdatedTime = uint(time.Now().Unix())
		shopShelfCardList[index] = shopShelfCard
		//关联关系表开卡方式

		err = batchInsertShopOpenTypeRelation(ctx, p.OpenTypeList, uint(p.AutomaticNum), uint(cardShop.CardId), uint(cardShop.ShelfShopId))
		if err != nil {
			utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("batch_insert_shop_card_open_type_relation", err.Error()))
			_ = components.M.Rollback(ctx)
			return err
		}

		if len(p.InoutTime) != utils.IntZero {

			err = batchInsertInoutTime(ctx, p.InoutTime, uint(p.ShelfBrandId), uint(cardShop.ShelfShopId), uint(cardShop.CardId))
			if err != nil {
				utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("batch_insert_inout_time", err.Error()))
				_ = components.M.Rollback(ctx)
				return err
			}
		}

		if len(p.CardPrice) != utils.IntZero {
			err = batchInsertShopCardPrice(ctx, p.CardPrice, uint(p.ShelfBrandId), uint(cardShop.ShelfShopId), uint(cardShop.CardId))
			if err != nil {
				utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("batch_insert_shop_card_price", err.Error()))
				_ = components.M.Rollback(ctx)
				return err
			}
		}

		if len(p.CourseList) != utils.IntZero {
			err = batchInsertShopCardCourseInterest(ctx, p.CourseList, uint(p.ShelfBrandId), uint(cardShop.ShelfShopId), uint(cardShop.CardId))

			if err != nil {
				utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("batch_insert_shop_card_course", err.Error()))
				_ = components.M.Rollback(ctx)
				return err
			}
		}

		if len(p.VipRegion) != utils.IntZero {
			RegionIds := make([]uint, len(p.VipRegion))

			for index, region := range p.VipRegion {
				RegionIds[index] = uint(region.ShopRegionId)
			}
			err = batchAddShopVipRegionRelation(ctx, RegionIds, uint(p.ShelfBrandId), uint(cardShop.ShelfShopId), uint(cardShop.CardId))

			if err != nil {
				utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("batch_insert_shop_vip_region_relation", err.Error()))
				_ = components.M.Rollback(ctx)
				return err
			}
		}
	}

	err = productModel.ShopShelfCardBatchInsert(ctx, shopShelfCardList)
	if err != nil {
		utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("batch_insert_shop_shelf_card", err.Error()))
		_ = components.M.Rollback(ctx)
		return err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("transaction_commit", err.Error()))
		return err
	}

	return nil
}

//批量插入入场时间
func batchInsertInoutTime(ctx context.Context, p []*product.InoutTime, brandId, shopId, cardId uint) error {
	shopCardAdmissionTimeList := make([]*productModel.ShopCardAdmissionTime, len(p))

	for indexTime, inoutTimeCard := range p {
		shopCardAdmissionTime := &productModel.ShopCardAdmissionTime{}
		shopCardAdmissionTime.BrandID = brandId
		shopCardAdmissionTime.ShopID = shopId
		shopCardAdmissionTime.CardID = cardId
		shopCardAdmissionTime.WeekDay = int8(inoutTimeCard.WeekDay)
		shopCardAdmissionTime.StartTime = uint(inoutTimeCard.StartTime)
		shopCardAdmissionTime.EndTime = uint(inoutTimeCard.EndTime)
		shopCardAdmissionTime.IsDel = utils.NOT_DELETED
		shopCardAdmissionTime.CreatedTime = uint(time.Now().Unix())
		shopCardAdmissionTime.UpdatedTime = uint(time.Now().Unix())
		shopCardAdmissionTimeList[indexTime] = shopCardAdmissionTime
	}

	err := productModel.ShopCardAdmissionTimeBatchInsert(ctx, shopCardAdmissionTimeList)
	if err != nil {
		return err
	}
	return nil
}

//批量新增开卡方式
func batchInsertShopOpenTypeRelation(ctx context.Context, p []product.OpenType, automaticNum, cardId, shopId uint) error {
	shopCardOpenTypeRelationList := make([]*productModel.ShopCardOpenTypeRelation, len(p))

	for indexOpen, openType := range p {
		shopCardOpenTypeRelation := &productModel.ShopCardOpenTypeRelation{}
		shopCardOpenTypeRelation.ShopID = shopId
		shopCardOpenTypeRelation.CardID = cardId
		shopCardOpenTypeRelation.OpenType = int8(openType)
		shopCardOpenTypeRelation.IsDel = utils.NOT_DELETED
		shopCardOpenTypeRelation.AutomaticNum = automaticNum
		shopCardOpenTypeRelation.CreatedTime = uint(time.Now().Unix())
		shopCardOpenTypeRelation.UpdatedTime = uint(time.Now().Unix())
		shopCardOpenTypeRelationList[indexOpen] = shopCardOpenTypeRelation
	}
	err := productModel.ShopCardOpenTypeRelationBatchInsert(ctx, shopCardOpenTypeRelationList)
	if err != nil {
		return err
	}
	return nil
}

//批量插入卡价格设置
func batchInsertShopCardPrice(ctx context.Context, p []*product.CardPrice, brandId, shopId, cardId uint) error {
	shopCardPriceList := make([]*productModel.ShopCardPrice, len(p))

	utils.GetTraceLog(ctx).Debug("batchInsertShopCardPrice", zap.Any("cardPriceShop", p))
	for indexPrice, cardPriceShop := range p {
		shopCardPrice := &productModel.ShopCardPrice{}
		shopCardPrice.BrandID = brandId
		shopCardPrice.ShopID = shopId
		shopCardPrice.CardID = cardId
		shopCardPrice.CardTplPriceSettingID = uint(cardPriceShop.CardTplPriceSettingId)
		shopCardPrice.IsDel = utils.NOT_DELETED
		shopCardPrice.RallyPrice = uint(cardPriceShop.SetPrice)
		shopCardPrice.CreatedTime = uint(time.Now().Unix())
		shopCardPrice.UpdatedTime = uint(time.Now().Unix())
		shopCardPriceList[indexPrice] = shopCardPrice
	}
	err := productModel.ShopCardPriceBatchInsert(ctx, shopCardPriceList)

	if err != nil {
		return err
	}

	return nil
}

//批量添加课程权益
func batchInsertShopCardCourseInterest(ctx context.Context, courseList []*product.CourseList, brandId, shopId, cardId uint) error {
	shopCardCourseInterestList := make([]*productModel.ShopCardCourseInterest, len(courseList))

	for indexCourse, courseShop := range courseList {
		shopCardCourseInterest := &productModel.ShopCardCourseInterest{}
		shopCardCourseInterest.BrandID = brandId
		shopCardCourseInterest.ShopID = shopId
		shopCardCourseInterest.CardID = cardId
		shopCardCourseInterest.CourseID = uint(courseShop.CourseId)
		shopCardCourseInterest.CourseName = courseShop.CourseName
		shopCardCourseInterest.CourseType = int8(courseShop.CourseType)
		shopCardCourseInterest.IsDel = utils.NOT_DELETED
		shopCardCourseInterest.CreatedTime = uint(time.Now().Unix())
		shopCardCourseInterest.UpdatedTime = uint(time.Now().Unix())
		shopCardCourseInterestList[indexCourse] = shopCardCourseInterest
	}

	err := productModel.ShopCardCourseInterestBatchInsert(ctx, shopCardCourseInterestList)
	if err != nil {
		return err
	}

	return nil
}

//批量新增场地vip
func batchAddShopVipRegionRelation(ctx context.Context, v []uint, brandId, shopId, cardId uint) error {
	shopVipRegionRelationList := make([]*productModel.ShopVipRegionRelation, len(v))

	for index, regionId := range v {
		regionAdd := &productModel.ShopVipRegionRelation{}
		regionAdd.BrandID = brandId
		regionAdd.ShopID = shopId
		regionAdd.ProductID = cardId
		regionAdd.ProductType = int8(product.ProductType_PRODUCT_MEMBER_CARD)
		regionAdd.ShopRegionID = regionId
		regionAdd.IsDel = utils.NOT_DELETED
		regionAdd.CreatedTime = uint(time.Now().Unix())
		regionAdd.UpdatedTime = uint(time.Now().Unix())
		shopVipRegionRelationList[index] = regionAdd
	}
	err := productModel.BatchAddShopVipRegionRelation(ctx, shopVipRegionRelationList)

	if err != nil {
		return err
	}
	return nil
}

//会员卡批量下架
func DownShelfCardTpl(ctx context.Context, p *product.DownShelfCardTplRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return err
	}

	cardShopList := p.CardShop

	downProductId := uint(utils.IntZero)

	downProductType := int8(product.ProductType_PRODUCT_MEMBER_CARD)

	downShelfStatus := int8(product.ShelfStatus_SHELFSTATUS_DOWN)

	downShopId := uint(utils.IntZero)

	for _, cardShop := range cardShopList {
		downProductId = uint(cardShop.CardId)
		downShopId = uint(cardShop.ShelfShopId)
		err = productModel.ShopShelfCardDown(ctx, downProductId, downProductType, downShopId, downShelfStatus)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	err = components.M.Commit(ctx)
	if err != nil {
		return err
	}
	return nil

}

//会员卡重新上架
func ReUpShelfCardTpl(ctx context.Context, p *product.UpShelfCardTplRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("batch_insert_shop_vip_region_relation", err.Error()))
		return err
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	reShelfCard, err := productModel.GetShopShelfCardByShopIdAndCardId(ctx, uint(p.CardShop[0].CardId), uint(p.CardShop[0].ShelfShopId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("get_shop_shelf_card", err.Error()))

		return err
	}

	reShelfCard.ShelfStatus = int8(product.ShelfStatus_SHELFSTATUS_UP)

	reShelfCard.CourseInterests = int8(p.CourseInterests)

	reShelfCard.AdmissionStatus = int8(p.ShopCardAdmissionTime)

	reShelfCard.ProductType = int8(product.ProductType_PRODUCT_MEMBER_CARD)

	reShelfCard.IsDel = utils.NOT_DELETED

	reShelfCard.UpdatedTime = uint(time.Now().Unix())

	err = reShelfCard.Save(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("update_shop_shelf_card", err.Error()))
		_ = components.M.Rollback(ctx)
		return err
	}
	//update shop_card_open_type_relation

	// 开卡方式
	err = UpdateOpenTypeRelation(ctx, uint(p.CardShop[0].ShelfShopId), uint(p.CardShop[0].CardId), p.OpenTypeList, uint(p.AutomaticNum))
	if err != nil {
		utils.GetTraceLog(ctx).Error("up_shelf_card", zap.Any("update_shop_card", err.Error()))
		return err
	}

	//入场限制
	if p.ShopCardAdmissionTime != product.ShopCardAdmissionTime_CARD_INOUT_TIME_UNLIMITED {
		err = updateCardAdmissionTime(ctx, uint(p.CardShop[0].CardId), uint(p.CardShop[0].ShelfShopId))
		if err != nil {
			return err
		}
	}

	// 关联关系课程列表
	if p.CourseInterests == product.CourseInterests_COURSE_INTERESTS_SPECIFIC_COURSE {
		err = updateCardCourse(ctx, uint(p.ShelfBrandId), uint(p.CardShop[0].CardId), uint(p.CardShop[0].ShelfShopId), p.CourseList)
		if err != nil {
			return err
		}
	}

	//更改卡价格
	if len(p.CardPrice) != utils.IntZero {
		err = updateCardPrice(ctx, uint(p.ShelfBrandId), uint(p.CardShop[0].CardId), uint(p.CardShop[0].ShelfShopId), p.CardPrice)
		if err != nil {
			return err
		}
	}
	//更改场地限制
	if len(p.VipRegion) != utils.IntZero {
		err = editVipRegion(ctx, uint(p.ShelfBrandId), uint(p.CardShop[0].CardId), uint(p.CardShop[0].ShelfShopId), p.VipRegion)
		if err != nil {
			return err
		}
	}

	err = components.M.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

//比对信息
func CheckDiffString(ctx context.Context, currentOpenTypes []string, newOpenTypes []string) ([]string, []string) {
	changelog, _ := diff.Diff(currentOpenTypes, newOpenTypes)

	utils.GetTraceLog(ctx).Info("diff card open type", zap.Any("current_open_type", currentOpenTypes),
		zap.Any("new_open_type", newOpenTypes), zap.Any("result", changelog))
	createTypes := make([]string, 0)
	deleteTypes := make([]string, 0)

	for _, log := range changelog {
		if log.Type == "delete" {
			if log.From != nil {
				deleteTypes = append(deleteTypes, log.From.(string))
			} else {
				deleteTypes = append(deleteTypes, log.To.(string))
			}
		} else if log.Type == "create" {
			if log.From != nil {
				createTypes = append(createTypes, log.From.(string))
			} else {
				createTypes = append(createTypes, log.To.(string))
			}
		} else if log.Type == "update" {
			if log.From != nil {
				deleteTypes = append(deleteTypes, log.From.(string))
				createTypes = append(createTypes, log.To.(string))
			} else {
				deleteTypes = append(deleteTypes, log.To.(string))
			}
		}
	}

	utils.GetTraceLog(ctx).Debug("diff result", zap.Any("create", createTypes), zap.Any("delete", deleteTypes))

	return createTypes, deleteTypes
}

//更新开卡方式
func UpdateOpenTypeRelation(ctx context.Context, shopId, cardId uint, p []product.OpenType, automaticNum uint) error {
	reShopCardOpenTypeRelationList, err := productModel.GetShopCardOpenTypeRelationByShopIdAndCardId(ctx, cardId, shopId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_open_type_relation", zap.String("get-shop_card_open_type_relation_by_shop_id_and_card_id", err.Error()))
		return err
	}

	originList := make([]string, len(reShopCardOpenTypeRelationList))

	newOpenTypeList := make([]string, len(p))

	//拼接字符串，拼接规则：开卡方式+延期天数
	for index, reShopCardOpenTypeRelation := range reShopCardOpenTypeRelationList {
		originList[index] = strconv.Itoa(int(reShopCardOpenTypeRelation.OpenType)) + "," + strconv.Itoa(int(reShopCardOpenTypeRelation.AutomaticNum))
	}

	for index, OpenType := range p {
		newOpenTypeList[index] = strconv.Itoa(int(OpenType)) + "," + strconv.Itoa(int(automaticNum))
	}

	insertOpenType, deleteOpenType := CheckDiffString(ctx, originList, newOpenTypeList)

	if len(insertOpenType) != utils.IntZero {
		openTypeList := make([]product.OpenType, len(insertOpenType))

		for index, openTypeInsert := range insertOpenType {
			openType, err := strconv.Atoi(strings.Split(openTypeInsert, ",")[0])
			if err != nil {
				utils.GetTraceLog(ctx).Error("update_open_type_relation", zap.String("insert_open_type", err.Error()))
				return err
			}

			openTypeList[index] = product.OpenType(openType)
		}

		err = batchInsertShopOpenTypeRelation(ctx, openTypeList, automaticNum, cardId, shopId)

		if err != nil {
			utils.GetTraceLog(ctx).Error("update_open_type_relation", zap.Any("openTypeList", openTypeList))
			return err
		}
	}

	if len(deleteOpenType) != utils.IntZero {

		deleteOpenTypeList, deleteAutomaticNumberList, err := PackDeleteList(deleteOpenType)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_open_type_relation", zap.String("delete_open_type", err.Error()))
		}

		err = productModel.BatchDeleteOpenType(ctx, deleteAutomaticNumberList, shopId, cardId, deleteOpenTypeList)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_open_type_relation", zap.String("delete some status", err.Error()))
			return err
		}
	}

	return nil
}

//c=["2,1","3,1","4,1"]
//a=[2,3,4]
//b=[1,1,1]

func PackDeleteList(s []string) (a []uint, b []uint, err error) {
	firstList := make([]uint, len(s))

	lastList := make([]uint, len(s))

	for index, c := range s {

		first, err := strconv.Atoi(strings.Split(c, ",")[0])
		if err != nil {
			return nil, nil, err
		}

		second, err := strconv.Atoi(strings.Split(c, ",")[1])
		if err != nil {
			return nil, nil, err
		}

		firstList[index] = uint(first)
		lastList[index] = uint(second)
	}
	return firstList, lastList, nil

}

//shop_card_admission_time
func updateCardAdmissionTime(ctx context.Context, cardId, shopId uint) error {

	reCardAdmissionTimeList, err := productModel.GetShopCardAdmissionTimeByShopIdAndCardId(ctx, shopId, cardId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_card_admission_time", zap.String("get_shop_card_admission_time", err.Error()))
		return err
	}

	for _, reCardAdmissionTime := range reCardAdmissionTimeList {
		reCardAdmissionTime.WeekDay = int8(reCardAdmissionTime.WeekDay)
		reCardAdmissionTime.StartTime = uint(reCardAdmissionTime.StartTime)
		reCardAdmissionTime.UpdatedTime = uint(reCardAdmissionTime.EndTime)
		err = reCardAdmissionTime.Save(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_card_admission_time", zap.String("update_shop_card_admission_time", err.Error()))
			return err
		}
	}

	return nil
}

//更新卡价格修改
func updateCardPrice(ctx context.Context, brandId, cardId, shopId uint, p []*product.CardPrice) error {
	reCardPrice, err := productModel.GetShopCardPriceByShopIdAndCardId(ctx, shopId, cardId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_card_price", zap.String("get_card_price_by_shop_id_and_card_id", err.Error()))
		return err
	}

	originSetIds := make([]string, len(reCardPrice))

	newSetIds := make([]string, len(p))
	//拼接规则：价格梯度+设置价格
	for index, origin := range reCardPrice {
		originSetIds[index] = strconv.Itoa(int(origin.CardTplPriceSettingID)) + "," + strconv.Itoa(int(origin.RallyPrice))
	}

	for indexNew, newId := range p {
		newSetIds[indexNew] = strconv.Itoa(int(newId.CardTplPriceSettingId)) + "," + strconv.Itoa(int(newId.SetPrice))

	}

	insertString, deleteString := CheckDiffString(ctx, originSetIds, newSetIds)

	if len(insertString) != utils.IntZero {

		insertSetInfo := make([]*product.CardPrice, len(insertString))

		for index, insertId := range insertString {
			cardPriceSetId, cardPriceSetPrice, err := StrToInt(ctx, insertId)
			if err != nil {
				utils.GetTraceLog(ctx).Error("update_card_price", zap.String("str_to_int", err.Error()))
				return err
			}

			insertSetInfo[index].SetPrice = uint32(cardPriceSetPrice)
			insertSetInfo[index].CardTplPriceSettingId = uint32(cardPriceSetId)
		}

		err = batchInsertShopCardPrice(ctx, insertSetInfo, brandId, cardId, shopId)

		if err != nil {
			return err
		}
	}
	//拆分规则：价格梯度+设置价格
	if len(deleteString) != utils.IntZero {
		settingIds, settingPrices, err := PackDeleteList(deleteString)
		err = productModel.BatchDeleteCardPrice(ctx, settingIds, shopId, cardId, settingPrices)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_card_price", zap.String("batch_delete_card-price", err.Error()))
			return err
		}
	}
	return nil
}

//strtoint
func StrToInt(ctx context.Context, s string) (a, b int, err error) {
	intA, err := strconv.Atoi(strings.Split(s, ",")[0])
	if err != nil {
		utils.GetTraceLog(ctx).Error("str_to_int", zap.String("first_part ", err.Error()))
		return utils.IntZero, utils.IntZero, err
	}

	intB, err := strconv.Atoi(strings.Split(s, ",")[1])
	if err != nil {
		utils.GetTraceLog(ctx).Error("str_to_int", zap.String("last_part", err.Error()))
		return utils.IntZero, utils.IntZero, err
	}
	return intA, intB, nil
}

//更新卡课程
func updateCardCourse(ctx context.Context, brandId, cardId, shopId uint, p []*product.CourseList) error {
	reCardCourseList, err := productModel.ShopCourseByCardId(ctx, cardId, shopId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_card_course", zap.String("shop_course_by_card_id", err.Error()))
		return err
	}

	originList := make([]string, len(reCardCourseList))

	newCardCourseList := make([]string, len(p))

	for index, reCardCourse := range reCardCourseList {
		originList[index] = strconv.Itoa(int(reCardCourse.CourseID)) + "," + strconv.Itoa(int(reCardCourse.CourseType))
	}

	for index, course := range p {
		newCardCourseList[index] = strconv.Itoa(int(course.CourseId)) + "," + strconv.Itoa(int(course.CourseType))
	}

	insertCourseList, deleteCourseIds := CheckDiffString(ctx, originList, newCardCourseList)

	if len(insertCourseList) != utils.IntZero {
		reCourseList := make([]*product.CourseList, len(insertCourseList))

		for index, courseList := range insertCourseList {

			for _, newCourse := range p {
				if courseList == strconv.Itoa(int(newCourse.CourseId))+strconv.Itoa(int(newCourse.CourseType)) {
					reCourseList[index] = newCourse
				}
			}
		}

		err = batchInsertShopCardCourseInterest(ctx, reCourseList, brandId, shopId, cardId)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_card_course", zap.String("batch__shop_course_interest", err.Error()))
			return err
		}
	}

	if len(deleteCourseIds) != utils.IntZero {
		reDeleteCourseIds := make([]uint, len(deleteCourseIds))

		reDeleteCourseTypes := make([]uint, len(deleteCourseIds))

		for indexDelete, deleteCourseId := range deleteCourseIds {
			courseId, err := strconv.Atoi(strings.Split(deleteCourseId, ",")[0])
			if err != nil {
				utils.GetTraceLog(ctx).Error("update_card_course", zap.String("course_id_str_to_int_failure", err.Error()))
				return err
			}

			typeId, err := strconv.Atoi(strings.Split(deleteCourseId, ",")[1])
			if err != nil {
				utils.GetTraceLog(ctx).Error("update_card_course", zap.String("course_type_str_to_int_failure", err.Error()))
				return err
			}
			reDeleteCourseIds[indexDelete] = uint(courseId)

			reDeleteCourseTypes[indexDelete] = uint(typeId)
		}

		err = productModel.BatchDeletecourseList(ctx, reDeleteCourseIds, reDeleteCourseTypes, shopId, cardId)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_card_course", zap.String("batch_delete_course", err.Error()))
			return err
		}
	}
	return nil
}

//更新vip场地通行
func editVipRegion(ctx context.Context, brandId, cardId, shopId uint, p []*product.VipRegion) error {
	vipRegionList, err := productModel.GetShopVipRegionRelationByShopIdAndCardId(ctx, cardId, shopId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("edit_vip_region", zap.String("get_shop_vip_region_relation_by_shop_id_and_card_id", err.Error()))
		return err
	}
	originRegionIds := make([]uint32, len(vipRegionList))

	newRegionIds := make([]uint32, len(p))

	for index, reRegion := range vipRegionList {
		originRegionIds[index] = uint32(reRegion.ShopRegionID)
	}

	for index, newRegion := range p {
		newRegionIds[index] = uint32(newRegion.ShopRegionId)
	}

	insertRegionIds, deleteRegionIds := GetDiffIds(ctx, originRegionIds, newRegionIds)

	if len(insertRegionIds) != utils.IntZero {
		err = batchAddShopVipRegionRelation(ctx, insertRegionIds, brandId, shopId, cardId)
		if err != nil {
			utils.GetTraceLog(ctx).Error("edit_vip_region", zap.String("batch_add_shop_vip_region_relation", err.Error()))
			return err
		}

	}

	if len(deleteRegionIds) != utils.IntZero {
		err = productModel.BatchDelVipRegion(ctx, shopId, cardId, deleteRegionIds)
		if err != nil {
			utils.GetTraceLog(ctx).Error("edit_vip_region", zap.String("batch_del_shop_vip_region", err.Error()))
			return err
		}

	}

	return nil
}

//上架会员卡详情
func GetShelfCardInfo(ctx context.Context, p *product.GetMemberCardTplInfoRequest) (*product.GetShelfCardTplInfoResponse, error) {
	cardInfo := &product.GetShelfCardTplInfoResponse{}

	shelfCard, err := productModel.ShopShelfCardByID(ctx, uint(p.CardId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_shelf_card", zap.Uint32("card_id", p.CardId), zap.Error(err))
		return cardInfo, err
	}

	tplCard, err := productModel.MemberCardTemplateByID(ctx, shelfCard.ProductID)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_tpl_card", zap.Uint("card_id", shelfCard.ProductID), zap.Error(err))
		return cardInfo, err
	}

	//todo 获取停售原因

	//入场门店
	if tplCard.AdmissionRange == int8(product.AdmissionRange_ADMISSION_MUCH_SHOP) {
		cardInfo.AdmissionShopList, err = GetSupportShop(ctx, tplCard.ID, int8(product.ProductType_PRODUCT_MEMBER_CARD))
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_support_shop", zap.Uint("tpl_card_id", tplCard.ID), zap.Error(err))
			return cardInfo, err
		}
	}

	//售卖门店
	if tplCard.SupportSales == int8(product.SupportSales_SALES_PART) {
		cardInfo.SellShopList, err = GetSaleShop(ctx, uint(tplCard.ID), int8(product.ProductType_PRODUCT_MEMBER_CARD))
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_sale_shop", zap.Uint("tpl_card_id", tplCard.ID), zap.Error(err))
			return cardInfo, err
		}
		cardInfo.TotalShopNum = uint32(len(cardInfo.SellShopList))
	} else {
		shops, err := shopModel.ShopByBrandID(ctx, tplCard.BrandID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_sale_shop", zap.Uint("brand_id", tplCard.BrandID), zap.Error(err))
			return cardInfo, err
		}
		cardInfo.TotalShopNum = uint32(len(shops))
	}

	//已上架门店数量
	shops, err := productModel.ShopShelfCardByProductId(ctx, uint(tplCard.ID), int8(product.ProductType_PRODUCT_MEMBER_CARD))
	for _, shop := range shops {
		if shop.ShelfStatus == int8(product.ShelfStatus_SHELFSTATUS_UP) {
			cardInfo.ShelfShopNum += 1
		}
	}

	//价格梯度
	cardInfo.PriceGradient, err = getPriceGradient(ctx, uint(tplCard.ID))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_price_gradient", zap.Uint("tpl_card_id", tplCard.ID), zap.Error(err))
		return cardInfo, err
	}

	if tplCard.PriceSetting == int8(product.SetPriceType_SET_PRICE_TYPE_SHOP) {
		shelfCardPrices, err := productModel.GetShopCardPriceByShopIdAndCardId(ctx, shelfCard.ProductID, shelfCard.ShopID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_shelf_price", zap.Uint("tpl_card_id", shelfCard.ProductID),
				zap.Uint("shop_id", shelfCard.ShopID), zap.Error(err))
			return cardInfo, err
		}

		shelfCardPriceList := make([]*product.CardPrice, len(shelfCardPrices))
		for index, scp := range shelfCardPrices {
			price := &product.CardPrice{
				CardTplPriceSettingId: uint32(scp.CardTplPriceSettingID),
				SetPrice:              uint32(scp.RallyPrice),
			}
			shelfCardPriceList[index] = price
		}
		cardInfo.ShopPriceGradient = shelfCardPriceList
	}

	//开卡方式
	openTypes, err := productModel.GetShopCardOpenTypeRelationByShopIdAndCardId(ctx, shelfCard.ProductID, shelfCard.ShopID)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_open_type", zap.Uint("tpl_card_id", shelfCard.ProductID),
			zap.Uint("shop_id", shelfCard.ShopID), zap.Error(err))
		return cardInfo, err
	}
	openTypeList := make([]product.OpenType, len(openTypes))
	for index, ot := range openTypes {
		openType := product.OpenType(ot.OpenType)
		if openType == product.OpenType_USE_BY_SPECIAL_DATE {
			cardInfo.AutomaticNum = uint32(ot.AutomaticNum)
		}
		openTypeList[index] = openType
	}
	cardInfo.OpenTypeList = openTypeList

	//约课权益
	if shelfCard.CourseInterests == int8(product.CourseInterests_COURSE_INTERESTS_SPECIFIC_COURSE) {
		courses, err := productModel.ShopCourseByCardId(ctx, shelfCard.ProductID, shelfCard.ShopID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_course_interests", zap.Uint("tpl_card_id", shelfCard.ProductID),
				zap.Uint("shop_id", shelfCard.ShopID), zap.Error(err))
			return cardInfo, err
		}
		courseList := make([]*product.CourseList, len(courses))
		for index, c := range courses {
			course := &product.CourseList{}
			course.CourseName = c.CourseName
			course.CourseType = uint32(c.CourseType)
			course.CourseId = uint32(c.CourseID)
			courseList[index] = course
		}
		cardInfo.CourseList = courseList
	}

	//入场时段
	inOutTimes, err := productModel.GetShopCardAdmissionTimeByShopIdAndCardId(ctx, shelfCard.ProductID, shelfCard.ShopID)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_inout_time", zap.Uint("tpl_card_id", shelfCard.ProductID),
			zap.Uint("shop_id", shelfCard.ShopID), zap.Error(err))
		return cardInfo, err
	}
	inOutTimeList := make([]*product.InoutTime, len(inOutTimes))
	for index, iot := range inOutTimes {
		inoutTime := &product.InoutTime{
			WeekDay:   uint32(iot.WeekDay),
			StartTime: uint32(iot.StartTime),
			EndTime:   uint32(iot.EndTime),
		}
		inOutTimeList[index] = inoutTime
	}
	cardInfo.InoutTime = inOutTimeList

	// vip场地通行
	vipRegion, err := productModel.GetShopVipRegionRelationByShopIdAndCardId(ctx, shelfCard.ProductID, shelfCard.ShopID)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_region_ids", zap.Uint("tpl_card_id", shelfCard.ProductID),
			zap.Uint("shop_id", shelfCard.ShopID), zap.Error(err))
		return cardInfo, err
	}

	if len(vipRegion) != utils.IntZero {
		regionIds := make([]uint, len(vipRegion))

		for index, region := range vipRegion {
			regionIds[index] = region.ShopRegionID
		}
		regionInfo, err := shopModel.GetShopAreaByIds(ctx, regionIds)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_region_info", zap.Any("region_ids", regionIds), zap.Error(err))
			return cardInfo, err
		}

		for index, regionDetail := range regionInfo {
			cardInfo.VipRegionInfo[index].RegionName = regionDetail.AreaName
			cardInfo.VipRegionInfo[index].RegionId = uint32(regionDetail.ID)
		}
	}

	//获取会员卡背景图
	img, err := imagesService.GetImages(ctx, uint(tplCard.AlbumID), int8(imagesModel.COVER_TYPE_GENERAL))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_card_bg", zap.Int("album_id", tplCard.AlbumID), zap.Error(err))
		return cardInfo, err
	}
	if len(img) != utils.IntZero {
		cardInfo.CardBg = &product.Image{
			ImageId: uint32(img[0].ID),
			ImgUrl:  img[0].ImageURL,
		}
	}

	cardInfo.BrandId = uint32(shelfCard.BrandID)
	cardInfo.CardId = uint32(shelfCard.ID)
	cardInfo.CardTplId = uint32(tplCard.ID)
	cardInfo.CardName = tplCard.CardName
	cardInfo.CardType = product.CardType(tplCard.CardType)
	cardInfo.SellStatus = product.ProductSellStatus(tplCard.SellStatus)
	cardInfo.ShelfStatus = product.ShelfStatus(shelfCard.ShelfStatus)
	cardInfo.StartTime = uint32(tplCard.StartTime)
	cardInfo.EndTime = uint32(tplCard.EndTime)
	cardInfo.SupportSales = product.SupportSales(tplCard.SupportSales)
	cardInfo.AdmissionRange = product.AdmissionRange(tplCard.AdmissionRange)
	cardInfo.PriceSetting = product.SetPriceType(tplCard.PriceSetting)
	cardInfo.SellType = product.CardSellType(tplCard.SellType)
	cardInfo.CourseInterests = product.CourseInterests(shelfCard.CourseInterests)
	cardInfo.ShopCardAdmissionTime = product.ShopCardAdmissionTime(shelfCard.AdmissionStatus)
	cardInfo.CardContents = tplCard.CardContents
	cardInfo.CardIntroduction = tplCard.CardIntroduction

	return cardInfo, nil
}
