package shopService

import (
	"context"
	"go.7yes.com/j7f/components/errors"
	"j7go/errors"
	"j7go/proto/shop"
	"time"
	"j7go/components"
	"j7go/models/region"
	"j7go/models/shop"
	"j7go/services/images"
	"j7go/utils"
)

//调用model获取门店信息
func GetShopInfo(ctx context.Context, shopId uint) (*shopModel.Shop, error) {
	shopInfo, err := shopModel.ShopByID(ctx, shopId)
	if err != nil {
		return nil, errors.NewFromCode(business_errors.ShopError_SHOP_NOT_FOUND)
	}

	return shopInfo, nil
}

//更新门店信息
func UpdateShopInfo(ctx context.Context, request *shop.ShopInfoRequest) error {
	if request.ShopName != utils.StringInit {
		shopInfoSameName, err := shopModel.GetShopInfoByName(ctx, request.ShopName, request.BrandId)
		if err != nil {
			return err
		}
		//不允许存在同品牌下名称相同的门店
		if shopInfoSameName.ID != utils.IntZero && shopInfoSameName.ID != uint(request.Id) {
			return errors.NewFromCode(business_errors.ShopError_SHOP_NAME_ALREADY_EXISTS)
		}
	}

	//获取省市区信息
	regionInfo, err := region.RegionsByDistrictID(ctx, uint(request.DistrictId))
	if err != nil {
		return err
	}
	if len(regionInfo) == 0 {
		return errors.NewFromCode(business_errors.ShopError_DISTRICT_ID_INCORRECT)
	}

	shopInfo := shopModel.Shop{}
	shopInfo.ID = uint(request.Id)
	shopInfo.BrandID = uint(request.BrandId)
	shopInfo.ShopName = request.ShopName
	shopInfo.Address = request.Address
	shopInfo.ShopStatus = int8(request.ShopStatus)
	shopInfo.Description = request.Description
	shopInfo.IsAllday = int8(request.IsAllday)
	shopInfo.ProvinceID = uint(request.ProvinceId)
	shopInfo.CityID = uint(request.CityId)
	shopInfo.DistrictID = uint(request.DistrictId)
	shopInfo.Lat = request.Lat
	shopInfo.Lng = request.Lng
	shopInfo.UpdatedTime = uint(time.Now().Unix())
	shopInfo.ProvinceName = regionInfo[0].ProvinceName
	shopInfo.CityName = regionInfo[0].CityName
	shopInfo.DistrictName = regionInfo[0].DistrictName

	err = shopInfo.UpdateShopById(ctx, shopInfo.ID)
	if err != nil {
		return err
	}

	return nil
}

//插入门店信息 返回自增id
func InsertShopInfo(ctx context.Context, request *shop.ShopInfoRequest) (uint32, error) {
	if request.ShopName != utils.StringInit {
		shopInfoSameName, err := shopModel.GetShopInfoByName(ctx, request.ShopName, request.BrandId)
		if err != nil {
			return utils.IntZero, err
		}
		if shopInfoSameName != nil && shopInfoSameName.ID != utils.IntZero {
			return utils.IntZero, errors.NewFromCode(business_errors.ShopError_SHOP_NAME_ALREADY_EXISTS)
		}
	}

	//获取省市区信息
	regionInfo, err := region.RegionsByDistrictID(ctx, uint(request.DistrictId))
	if err != nil {
		return utils.IntZero, err
	}
	if len(regionInfo) == 0 {
		return utils.IntZero, errors.NewFromCode(business_errors.ShopError_DISTRICT_ID_INCORRECT)
	}

	shopInfo := &shopModel.Shop{}
	shopInfo.BrandID = uint(request.BrandId)
	shopInfo.ShopName = request.ShopName
	shopInfo.Address = request.Address
	shopInfo.ShopStatus = int8(request.ShopStatus)
	shopInfo.Description = request.Description
	shopInfo.IsAllday = int8(request.IsAllday)
	shopInfo.ProvinceID = uint(request.ProvinceId)
	shopInfo.CityID = uint(request.CityId)
	shopInfo.DistrictID = uint(request.DistrictId)
	shopInfo.Lat = request.Lat
	shopInfo.Lng = request.Lng
	shopInfo.AlbumID = uint(request.AlbumId)
	shopInfo.IsDel = int8(utils.NOT_DELETED)
	shopInfo.UpdatedTime = uint(time.Now().Unix())
	shopInfo.CreatedTime = uint(time.Now().Unix())
	shopInfo.ProvinceName = regionInfo[0].ProvinceName
	shopInfo.CityName = regionInfo[0].CityName
	shopInfo.DistrictName = regionInfo[0].DistrictName

	err = shopInfo.Insert(ctx)
	if err != nil {
		return utils.IntZero, errors.NewFromCode(business_errors.ShopError_UPDATE_SHOP_FAILURE)
	}

	return uint32(shopInfo.ID), nil
}

//获取门店联系方式
func GetShopContactWays(ctx context.Context, shopId uint) ([]*shopModel.ShopContactWay, error) {
	shopContactWays, err := shopModel.GetShopContactWaysByShopId(ctx, shopId)
	if err != nil {
		return nil, errors.NewFromCode(business_errors.ShopError_SHOP_CONTACT_WAYS_NOT_FOUND)
	}

	return shopContactWays, nil
}

//更新门店联系方式
func UpdateShopContacts(ctx context.Context, request *shop.ShopContactsRequest) error {
	oldContactWays, err := shopModel.GetShopContactWaysByShopId(ctx, uint(request.ShopId))
	if err != nil {
		return errors.NewFromCode(business_errors.ShopError_GET_SHOP_CONTACT_WAYS_ERROR)
	}

	//删除老数据
	for _, contactWay := range oldContactWays {
		err = contactWay.Delete(ctx)
		if err != nil {
			return errors.NewFromCode(business_errors.ShopError_UPDATE_SHOP_SERVICES_FAILURE)
		}
	}

	//插入数据库中不存在的联系方式数据
	err = InsertShopContacts(ctx, request)
	return err
}

//插入门店联系方式
func InsertShopContacts(ctx context.Context, request *shop.ShopContactsRequest) error {
	for _, contactRequest := range request.ShopContacts {
		contact := shopModel.ShopContactWay{}
		contact.ShopID = uint(request.ShopId)
		contact.ContactDetail = contactRequest.ContactDetail
		contact.ContactType = int8(contactRequest.ContactType)
		contact.IsFavorite = int8(contactRequest.IsFavorite)
		contact.IsDel = int8(utils.NOT_DELETED)
		contact.CreatedTime = uint(time.Now().Unix())
		contact.UpdatedTime = uint(time.Now().Unix())
		err := contact.Insert(ctx)
		if err != nil {
			return errors.NewFromCode(business_errors.ShopError_INSERT_SHOP_CONTACT_WATS_ERROR)
		}
	}
	return nil
}

//获取门店营业时间
func GetShopBusinessTime(ctx context.Context, shopId uint) ([]*shopModel.ShopBusinessTime, error) {
	shopBusinessTime, err := shopModel.ShopBusinessTimeByShopId(ctx, shopId)
	if err != nil {
		return nil, errors.NewFromCode(business_errors.ShopError_SHOP_BUSINESS_TIME_NOT_FOUND)
	}

	return shopBusinessTime, nil
}

//更新门店营业时间
func UpdateShopBusinessTime(ctx context.Context, request *shop.ShopBussinessTimeRequest) error {
	shopBusinessTimeInDb, err := shopModel.ShopBusinessTimeByShopId(ctx, uint(request.ShopId))
	if err != nil {
		return errors.NewFromCode(business_errors.ShopError_GET_SHOP_CONTACT_WAYS_ERROR)
	}

	for _, businessTime := range shopBusinessTimeInDb {
		err := businessTime.Delete(ctx)
		if err != nil {
			return errors.NewFromCode(business_errors.ShopError_DELETE_SHOP_BUSINESS_TIME_ERROR)
		}
	}

	for _, businessTime := range request.BusinessTime {
		err := InsertShopBusinessTime(ctx, uint(request.ShopId), businessTime)
		if err != nil {
			return err
		}
	}

	return nil
}

//插入多条门店营业时间
func InsertShopBusinessTimes(ctx context.Context, request *shop.ShopBussinessTimeRequest) error {
	for _, businessTime := range request.BusinessTime {
		err := InsertShopBusinessTime(ctx, uint(request.ShopId), businessTime)
		if err != nil {
			return err
		}
	}
	return nil
}

//插入单条门店营业时间
func InsertShopBusinessTime(ctx context.Context, shopId uint, businessTime *shop.ShopBusinessTime) error {
	newBussinessTime := shopModel.ShopBusinessTime{}
	newBussinessTime.ShopID = shopId
	newBussinessTime.WeekDay = int8(businessTime.WeekDay)
	newBussinessTime.StartTime = businessTime.StartTime
	newBussinessTime.EndTime = businessTime.EndTime
	newBussinessTime.IsDel = int8(utils.NOT_DELETED)
	newBussinessTime.CreatedTime = uint(time.Now().Unix())
	newBussinessTime.UpdatedTime = uint(time.Now().Unix())
	err := newBussinessTime.Insert(ctx)
	if err != nil {
		return errors.NewFromCode(business_errors.ShopError_INSERT_SHOP_BUSINESS_TIME_ERROR)
	}
	return nil
}

//获取门店服务
func GetShopServices(ctx context.Context, shopId uint) ([]*shopModel.ShopService, error) {
	shopServices, err := shopModel.GetShopServicesByShopId(ctx, shopId)
	if err != nil {
		return nil, errors.NewFromCode(business_errors.ShopError_SHOP_SERVICES_NOT_FOUND)
	}

	return shopServices, nil
}

//更新门店服务
func UpdateShopServices(ctx context.Context, request *shop.ShopServicesRequest) error {
	serviceRelationsInDb, err := shopModel.ShopServicesRelationByShopId(ctx, uint(request.ShopId))
	if err != nil {
		return errors.NewFromCode(business_errors.ShopError_GET_SHOP_SERVICES_RELATION_ERROR)
	}

	//删除旧的关系
	for _, relation := range serviceRelationsInDb {
		err = relation.Delete(ctx)
		if err != nil {
			return errors.NewFromCode(business_errors.ShopError_DELETE_SHOP_SERVICE_ERROR)
		}
	}

	for _, serviceId := range request.ServiceIds {
		//插入新数据
		serviceRelation := shopModel.ShopServicesRelation{}
		serviceRelation.ShopID = uint(request.ShopId)
		serviceRelation.ServiceID = uint(serviceId)
		serviceRelation.IsDel = int8(utils.NOT_DELETED)
		serviceRelation.CreatedTime = uint(time.Now().Unix())
		serviceRelation.UpdatedTime = uint(time.Now().Unix())
		err = serviceRelation.Insert(ctx)
		if err != nil {
			return errors.NewFromCode(business_errors.ShopError_UPDATE_SHOP_SERVICES_FAILURE)
		}
	}

	return nil
}

//插入门店服务
func InsertShopServices(ctx context.Context, request *shop.ShopServicesRequest) error {
	for _, serviceId := range request.ServiceIds {
		serviceRelation := shopModel.ShopServicesRelation{}
		serviceRelation.ShopID = uint(request.ShopId)
		serviceRelation.ServiceID = uint(serviceId)
		serviceRelation.IsDel = int8(utils.NOT_DELETED)
		serviceRelation.CreatedTime = uint(time.Now().Unix())
		serviceRelation.UpdatedTime = uint(time.Now().Unix())
		err := serviceRelation.Insert(ctx)
		if err != nil {
			return errors.NewFromCode(business_errors.ShopError_INSERT_SHOP_SERVICES_ERROR)
		}
	}
	return nil
}

//更新门店聚合接口
func UpdateShop(ctx context.Context, request *shop.ShopRequest) error {

	ctx, err := components.M.BeginTransaction(ctx, 2*time.Second)
	if err != nil {
		return err
	}

	//更新门店基本信息
	if request.ShopInfo != nil {
		err = UpdateShopInfo(ctx, request.ShopInfo)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//更新门店联系方式
	if request.ShopContacts != nil {
		err = UpdateShopContacts(ctx, request.ShopContacts)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//更新门店营业时间
	if request.ShopBussinessTimes != nil {
		err = UpdateShopBusinessTime(ctx, request.ShopBussinessTimes)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//更新门店服务
	if request.ShopServices != nil {
		err = UpdateShopServices(ctx, request.ShopServices)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//更新门店图片
	if request.ImageInfos != nil {
		shopInfo, err := GetShopInfo(ctx, uint(request.ShopId))
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}

		imageInfos := make([]*imagesService.ImageInfo, len(request.ImageInfos.Images))
		for index, image := range request.ImageInfos.Images {
			imageInfo := &imagesService.ImageInfo{}
			imageInfo.ImageUrl = image.ImageUrl
			imageInfo.ImageId = image.ImageId
			imageInfo.CoverType = int8(image.CoverType)
			imageInfos[index] = imageInfo
		}

		err = imagesService.UpdateAlbumImages(ctx, shopInfo.AlbumID, imageInfos)
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

//新增门店聚合接口
func InsertShop(ctx context.Context, request *shop.ShopRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 2*time.Second)
	if err != nil {
		return err
	}

	//新增门店图片
	var albumId uint32
	if request.ImageInfos != nil {
		imageInfos := make([]*imagesService.ImageInfo, len(request.ImageInfos.Images))
		for index, image := range request.ImageInfos.Images {
			imageInfo := &imagesService.ImageInfo{}
			imageInfo.ImageUrl = image.ImageUrl
			imageInfo.ImageId = image.ImageId
			imageInfo.CoverType = int8(image.CoverType)
			imageInfos[index] = imageInfo
		}

		albumId, err = imagesService.InsertImages(ctx, utils.IntZero, imageInfos)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//新增门店基本信息
	var shopId uint32
	request.ShopInfo.AlbumId = albumId
	shopId, err = InsertShopInfo(ctx, request.ShopInfo)
	if err != nil {
		_ = components.M.Rollback(ctx)
		return err
	}

	//新增门店联系方式
	if request.ShopContacts != nil {
		request.ShopContacts.ShopId = shopId
		err = InsertShopContacts(ctx, request.ShopContacts)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//新增门店营业时间
	if request.ShopBussinessTimes != nil {
		request.ShopBussinessTimes.ShopId = shopId
		err = InsertShopBusinessTimes(ctx, request.ShopBussinessTimes)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//新增门店服务
	if request.ShopServices != nil {
		request.ShopServices.ShopId = shopId
		err = InsertShopServices(ctx, request.ShopServices)
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
