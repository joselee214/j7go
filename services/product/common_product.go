package product

import (
	"context"
	"fmt"
	"github.com/r3labs/diff"
	"j7go/proto/product"
	"go.uber.org/zap"
	"time"
	"j7go/models/images"
	productModel "j7go/models/product"
	"j7go/utils"
)

//添加图片
func AddCardImages(ctx context.Context, imgs []string, albumId uint) error {
	//批量添加
	imgList := make([]*imagesModel.Image, len(imgs))
	for i, k := range imgs {
		images := &imagesModel.Image{}
		images.AlbumID = albumId
		images.ImageURL = k
		images.CoverType = 1
		images.CreatedTime = uint(time.Now().Unix())
		images.UpdatedTime = uint(time.Now().Unix())
		imgList[i] = images
	}
	err := imagesModel.ImagesBatchInsert(ctx, imgList)
	if err != nil {
		return err
	}
	return nil
}

//获取相册ID
func AddAlbum(ctx context.Context) (albumId uint, error error) {
	album := &imagesModel.Album{}
	album.CreatedTime = uint(time.Now().Unix())
	album.UpdatedTime = uint(time.Now().Unix())
	err := album.Insert(ctx)
	if err != nil {
		return utils.IntZero, err
	}
	return album.ID, nil
}

//维护支持场馆
func SaveSupportShop(ctx context.Context, brandId uint, productId uint, productType int8, shopList []uint32) error {
	list := make([]*productModel.ProductSupportShop, len(shopList))
	for index, shopId := range shopList {
		shop := &productModel.ProductSupportShop{}
		shop.BrandID = brandId
		shop.ShopID = uint(shopId)
		shop.ProductID = productId
		shop.ProductType = productType
		shop.UpdatedTime = uint(time.Now().Unix())
		shop.CreatedTime = uint(time.Now().Unix())

		list[index] = shop
	}

	err := productModel.ProductSupportShopBatchInsert(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

//维护售卖场馆
func SaveSaleShop(ctx context.Context, brandId, productId uint, productType int8, shopList []uint32) error {
	list := make([]*productModel.ShopProductSale, len(shopList))
	for index, shopId := range shopList {
		shop := &productModel.ShopProductSale{}
		shop.BrandID = brandId
		shop.ShopID = uint(shopId)
		shop.ProductID = productId
		shop.ProductType = productType
		shop.UpdatedTime = uint(time.Now().Unix())
		shop.CreatedTime = uint(time.Now().Unix())

		list[index] = shop
	}
	err := productModel.ShopProductSaleBatchInsert(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

//获取支持场馆
func GetSupportShop(ctx context.Context, productId uint, productType int8) ([]*product.SupportShop, error) {
	supportShopList, err := productModel.ProductSupportShopByProductId(ctx, productId, productType)
	if err != nil {
		return nil, err
	}
	list := make([]*product.SupportShop, len(supportShopList))
	for index, supportShop := range supportShopList {
		shop := &product.SupportShop{}
		shop.ShopId = uint32(supportShop.ID)
		shop.BrandId = uint32(supportShop.BrandID)
		shop.ShopName = supportShop.ShopName
		shop.ProvinceId = uint32(supportShop.ProvinceID)
		shop.CityId = uint32(supportShop.CityID)
		shop.DistrictId = uint32(supportShop.DistrictID)
		shop.ProvinceName = supportShop.ProvinceName
		shop.CityName = supportShop.CityName
		shop.DistrictName = supportShop.DistrictName
		list[index] = shop
	}

	return list, nil
}

//获取上架门店列表
func GetOnLineShop(ctx context.Context, productId uint, productType int8) ([]*product.SupportShop, error) {
	supportShopList, err := productModel.ProductOnLineShopByProductId(ctx, productId, productType)
	if err != nil {
		return nil, err
	}
	list := make([]*product.SupportShop, len(supportShopList))
	for index, supportShop := range supportShopList {
		shop := &product.SupportShop{}
		shop.ShopId = uint32(supportShop.ID)
		shop.BrandId = uint32(supportShop.BrandID)
		shop.ShopName = supportShop.ShopName
		shop.ProvinceId = uint32(supportShop.ProvinceID)
		shop.CityId = uint32(supportShop.CityID)
		shop.DistrictId = uint32(supportShop.DistrictID)
		shop.ProvinceName = supportShop.ProvinceName
		shop.CityName = supportShop.CityName
		shop.DistrictName = supportShop.DistrictName
		list[index] = shop
	}

	return list, nil
}

//获取售卖场馆
func GetSaleShop(ctx context.Context, productId uint, productType int8) ([]*product.SupportShop, error) {
	saleShopList, err := productModel.ShopProductSaleByProductId(ctx, productId, productType)
	if err != nil {
		return nil, err
	}

	list := make([]*product.SupportShop, len(saleShopList))
	for index, saleShop := range saleShopList {
		shop := &product.SupportShop{}
		shop.ShopId = uint32(saleShop.ID)
		shop.BrandId = uint32(saleShop.BrandID)
		shop.ShopName = saleShop.ShopName
		shop.ProvinceId = uint32(saleShop.ProvinceID)
		shop.CityId = uint32(saleShop.CityID)
		shop.DistrictId = uint32(saleShop.DistrictID)
		shop.ProvinceName = saleShop.ProvinceName
		shop.CityName = saleShop.CityName
		shop.DistrictName = saleShop.DistrictName
		list[index] = shop
	}
	fmt.Printf("sup shop new is %+v\n", list)

	return list, nil
}

//更新支持场馆
func UpdateSupportShop(ctx context.Context, brandId, productId uint, productType int8, shopList []uint32) error {
	supportShopList, err := GetSupportShop(ctx, productId, productType)
	if err != nil {
		return err
	}

	origShopList := make([]uint32, len(supportShopList))
	for index, shop := range supportShopList {
		origShopList[index] = shop.ShopId
	}

	delSlice, insertSlice := CompareUpdate(origShopList, shopList)

	if len(delSlice) != utils.IntZero {
		err = productModel.ProductSupportShopBatchDelete(ctx, productId, productType, delSlice)
		if err != nil {
			return err
		}
	}

	if len(insertSlice) != utils.IntZero {
		err = SaveSupportShop(ctx, brandId, productId, productType, insertSlice)
		if err != nil {
			return err
		}
	}

	return nil
}

//更新售卖场馆
func UpdateSaleShop(ctx context.Context, brandId, productId uint, productType int8, shopList []uint32) error {
	saleShopList, err := GetSaleShop(ctx, productId, productType)
	if err != nil {
		return err
	}

	origShopList := make([]uint32, len(saleShopList))
	for index, shop := range saleShopList {
		origShopList[index] = shop.ShopId
	}

	delSlice, insertSlice := CompareUpdate(origShopList, shopList)
	if len(delSlice) != utils.IntZero {
		err = productModel.ShopProductSaleBatchDelete(ctx, productId, productType, delSlice)
		if err != nil {
			return err
		}
	}

	if len(insertSlice) != utils.IntZero {
		err = SaveSaleShop(ctx, brandId, productId, productType, insertSlice)
		if err != nil {
			return err
		}
	}

	return nil
}

//比较更新
func CompareUpdate(origSlice, newSlice []uint32) (delSlice []uint, insertSlice []uint32) {
	slice1 := make([]interface{}, len(origSlice))
	slice2 := make([]interface{}, len(newSlice))

	for index, v := range origSlice {
		slice1[index] = v
	}

	for index, v := range newSlice {
		slice2[index] = v
	}

	sameSlice := utils.SliceIntersect(slice1, slice2)
	diffSlice := utils.SliceDiff(slice1, sameSlice)
	for _, v := range diffSlice {
		if v, ok := v.(uint32); ok {
			delSlice = append(delSlice, uint(v))
		}
	}

	diffSlice = utils.SliceDiff(slice2, sameSlice)
	for _, v := range diffSlice {
		if v, ok := v.(uint32); ok {
			insertSlice = append(insertSlice, v)
		}
	}

	return delSlice, insertSlice
}

//根据当前支持门店及前端传入最新支持门店
//比对出需新增及删除的门店
func GetDiffIds(ctx context.Context, currentIds []uint32, newIds []uint32) ([]uint, []uint) {
	changelog, _ := diff.Diff(currentIds, newIds)
	utils.GetTraceLog(ctx).Info("diff ids", zap.Any("current_ids", currentIds),
		zap.Any("new_ids", newIds), zap.Any("result", changelog))
	createIds := make([]uint, 0)
	deleteIds := make([]uint, 0)
	for _, log := range changelog {
		if log.Type == "delete" {
			if log.From != nil {
				deleteIds = append(deleteIds, uint(log.From.(uint32)))
			} else {
				deleteIds = append(deleteIds, uint(log.To.(uint32)))
			}
		} else if log.Type == "create" {
			if log.From != nil {
				createIds = append(createIds, uint(log.From.(uint32)))
			} else {
				createIds = append(createIds, uint(log.To.(uint32)))
			}
		} else {
			deleteIds = append(deleteIds, uint(log.From.(uint32)))
			createIds = append(createIds, uint(log.To.(uint32)))
		}
	}
	utils.GetTraceLog(ctx).Debug("diff result", zap.Any("create", createIds), zap.Any("delete", deleteIds))
	return createIds, deleteIds
}

//获取标签
func GetCourseTags(ctx context.Context, courseId uint, courseType int8) ([]*product.Tag, error) {
	ctas, err := productModel.GetCourseTrainAimsById(ctx, courseId, courseType)
	if err != nil {
		return nil, err
	}

	tagIds := make([]uint, len(ctas))
	for index, c := range ctas {
		tagIds[index] = c.SettingID
	}

	tagList, err := productModel.TagByIDs(ctx, tagIds)
	if err != nil {
		return nil, err
	}

	list := make([]*product.Tag, len(tagList))
	for index, t := range tagList {
		tag := &product.Tag{}
		tag.TagName = t.TagName
		tag.TagId = uint32(t.ID)
		list[index] = tag
	}

	return list, nil
}
