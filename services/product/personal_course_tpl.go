package product

import (
	"context"
	"github.com/joselee214/j7f/components/errors"
	"go.uber.org/zap"
	"j7go/components"
	"j7go/errors"
	"j7go/models/tests/images"
	productModel "j7go/models/tests/product"
	"j7go/proto/product"
	"j7go/services/images"
	"j7go/utils"
	"reflect"
	"time"
)

//新增私教课基础信息
func AddPersonalCourse(ctx context.Context, request *product.PersonalCourseRequest) (uint32,error) {
	ctx, err := components.M.BeginTransaction(ctx,time.Second * 2)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get transaction err",zap.String("scope","AddPersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	course, err := SavePersonalCourse(ctx,request)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("add course fail",zap.String("scope","AddPersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	err = SetCourseTrainAim(ctx,course.ID,request.TrainAim)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("add course train aim fail",zap.String("scope","AddPersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	err = SetCourseImages(ctx,course.AlbumID,request.CourseImg)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("add course image fail",zap.String("scope","AddPersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("commit add personal course transaction fail",zap.String("scope","AddPersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	return uint32(course.ID),nil
}

//更新私教课基础信息
func UpdatePersonalCourse(ctx context.Context, request *product.PersonalCourseRequest) (uint32,error) {
	ctx, err := components.M.BeginTransaction(ctx,time.Second * 2)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get transaction err",zap.String("scope","UpdatePersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	course, err := SavePersonalCourse(ctx,request)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("update course fail",zap.String("scope","UpdatePersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	err = SetCourseTrainAim(ctx,course.ID,request.TrainAim)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("update course train aim fail",zap.String("scope","UpdatePersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	err = SetCourseImages(ctx,course.AlbumID,request.CourseImg)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("update course image fail",zap.String("scope","UpdatePersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("commit update personal course transaction fail",zap.String("scope","UpdatePersonalCourse"),
			zap.String("err_msg",err.Error()))
		return utils.IntZero, err
	}

	return uint32(course.ID),nil
}

//保存私教课基础信息（新增或更新时调用）
func SavePersonalCourse(ctx context.Context,request *product.PersonalCourseRequest) (*productModel.PersonalCourseTemplate,error) {
	currentTime := utils.GetCurrentUnixTime()
	course := &productModel.PersonalCourseTemplate{}
	if utils.IntZero == request.CourseId {//新增课程
		//创建相册获取ID
		albumId, err := AddAlbum(ctx)
		if err != nil {
			_ = components.M.Rollback(ctx)
			utils.GetTraceLog(ctx).Error("new album fail",zap.String("scope","SavePersonalCourse"),
				zap.String("err_msg",err.Error()))
			return nil,errors.NewFromCode(business_errors.ImagesError_INSERT_ALBUM_ERROR)
		}

		course.BrandID = uint(request.BrandId)
		course.ShopID = uint(request.ShopId)
		course.CourseName = request.CourseName
		course.CategoryID = uint(request.CourseCategory)
		course.Duration = uint(request.Duration)
		course.Description = request.Description
		course.EffectiveUnit = uint(request.EffectiveUnit)
		course.PublishChannel = int8(request.PublishChannel)
		course.AlbumID = albumId
		course.IsDel = utils.NOT_DELETED
		course.IsAvailable = utils.Available
		course.CreatedTime = currentTime
		course.UpdatedTime = currentTime
	}else {//编辑课程
		course, err := productModel.PersonalCourseTemplateByID(ctx,uint(request.CourseId))
		if err != nil {
			utils.GetTraceLog(ctx).Error("get personal course fail",zap.String("scope","SavePersonalCourse"),
				zap.Uint("course_id",uint(request.CourseId)),zap.String("err_msg",err.Error()))
			return nil,errors.NewFromCode(business_errors.ProductError_UPDATE_PERSONAL_COURSE_ERROR)
		}
		course.Duration = uint(request.Duration)
		course.Description = request.Description
		course.EffectiveUnit = uint(request.EffectiveUnit)
		course.UpdatedTime = currentTime
	}

	err := course.Save(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("save personal course fail",zap.String("scope","SavePersonalCourse"),
			zap.String("err_msg",err.Error()))
		return nil,errors.NewFromCode(business_errors.ProductError_UPDATE_PERSONAL_COURSE_ERROR)
	}

	return course,nil
}

//设置课程图片
func SetCourseImages(ctx context.Context,albumId uint, img string) error {
	images := make([]*imagesService.ImageInfo,0)
	image := &imagesService.ImageInfo{}
	image.ImageUrl = img
	image.CoverType = imagesModel.IMAGE_TYPE_GENERAL
	images = append(images,image)
	_,err := imagesService.InsertImages(ctx,uint32(albumId),images)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add course image fail",zap.String("scope","AddPersonalCourse"),
			zap.Uint("album_id",albumId),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ImagesError_INSERT_IMAGES_ERROR)
	}

	return nil
}

//设置课程训练目标
func SetCourseTrainAim(ctx context.Context,courseId uint,newAimIds []uint32) error {
	currentTrainAims, err := productModel.GetCourseTrainAimsById(ctx,courseId, productModel.COURSE_PERSONAL)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get course train aims fail",zap.String("scope","SetCourseTrainAim"),
			zap.Uint("course_id",courseId),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_TRAIN_AIM_ERROR)
	}

	currentAimIds := GetCurrentIds(ctx,currentTrainAims)
	createIds,deleteIds := GetDiffIds(ctx,currentAimIds,newAimIds)
	formatCreateIds := make([]uint32,len(createIds))
	for index, createId := range createIds {
		formatCreateIds[index] = uint32(createId)
	}

	err = productModel.CourseSettingRelationBatchInsert(ctx,courseId, productModel.COURSE_PERSONAL,formatCreateIds)
	if err != err {
		utils.GetTraceLog(ctx).Error("batch insert train aim fail",zap.String("scope","SetCourseTrainAim"),
			zap.Uint("course_id",courseId),zap.Any("aim_ids",newAimIds),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_TRAIN_AIM_ERROR)
	}
	err = productModel.CourseSettingRelationBatchDelete(ctx,courseId, productModel.COURSE_PERSONAL,deleteIds)
	if err != err {
		utils.GetTraceLog(ctx).Error("batch delete train aim fail",zap.String("scope","SetCourseTrainAim"),
			zap.Uint("course_id",courseId),zap.Any("delete_ids",deleteIds),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_TRAIN_AIM_ERROR)
	}

	return nil
}

//设置上课门店及教练
func SetCourseShopsAndCoaches(ctx context.Context, request *product.SetCourseShopsRequest) error {
	ctx, err := components.M.BeginTransaction(ctx,time.Second * 2)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get transaction err",zap.String("scope","SetCourseShopsAndCoachs"),
			zap.String("err_msg",err.Error()))
		return err
	}

	course,currentShopSetting, err := UpdateCourseShopSetting(ctx,request)
	if err != nil {
		components.M.Rollback(ctx)
		return err
	}

	err = HandleSupportShops(ctx,course,currentShopSetting,request.ShopIds)
	if err != nil {
		components.M.Rollback(ctx)
		return err
	}

	err = HandleSupportCoaches(ctx,course,request.CoachIds)
	if err != nil {
		components.M.Rollback(ctx)
		return err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		components.M.Rollback(ctx)
		utils.GetTraceLog(ctx).Error("commit set course shops and coaches transaction fail",zap.String("scope","SetCourseShopsAndCoaches"),
			zap.Uint("course_id",uint(request.CourseId)),zap.String("err_msg",err.Error()))
		return err
	}

	return nil
}

//设置课程售价
func SetCourseSalePrice(ctx context.Context, request *product.SetSalePriceRequest) error {
	ctx, err := components.M.BeginTransaction(ctx,2 * time.Second)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get transaction err",zap.String("scope","SetCourseSalePrice"),
			zap.String("err_msg",err.Error()))
		return err
	}

	course, currentPriceSetting,err := UpdateCourseSalePriceSetting(ctx,uint(request.CourseId),int8(request.PriceSetting))
	if err != nil {
		components.M.Rollback(ctx)
		return err
	}

	currentPrices,err := productModel.PersonalCoursePriceByCourseId(ctx,uint(request.CourseId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get personal course price fail",zap.String("scope","SetCourseSalePrice"),
			zap.Uint("course_id",uint(request.CourseId)),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_SALE_PRICE_ERROR)
	}

	//发布渠道为品牌
	if productModel.PUBLISH_CHANNEL_BRAND == course.PublishChannel {
		//由统一定价切换为门店定价
		if productModel.PRICE_SETTING_ALL == currentPriceSetting && currentPriceSetting != int8(request.PriceSetting) {
			err = productModel.DeleteAllPriceByCourseId(ctx,course.ID)
			if err != nil {
				utils.GetTraceLog(ctx).Error("delete all course price fail",zap.String("scope","SetCourseSalePrice"),
					zap.Uint("course_id",uint(request.CourseId)),zap.String("err_msg",err.Error()))
				return errors.NewFromCode(business_errors.ProductError_SET_SALE_PRICE_ERROR)
			}
			return nil
		}
	}

	prices := request.CoursePrices
	createPrices := make([]*productModel.PersonalCoursePrice,0)
	deleteIds := make([]uint,0)
	currentUnixTime := utils.GetCurrentUnixTime()
	for _,price := range prices {
		if utils.IntZero == price.Id {
			priceModel := &productModel.PersonalCoursePrice{}
			priceModel.BrandID = course.BrandID
			priceModel.CourseID = course.ID
			if productModel.PUBLISH_CHANNEL_SHOP == course.PublishChannel {
				priceModel.ShopID = uint(request.ShopId)
			}
			priceModel.CoachLevelID = uint(price.CoachLevel)
			priceModel.SellStepStart = uint(price.SaleMin)
			priceModel.SellStepEnd = uint(price.SaleMax)
			priceModel.SellPrice = uint(price.Price)
			priceModel.TransferType = int8(price.TransferType)
			priceModel.TransferNum = uint(price.TransferNum)
			priceModel.IsOnlineSale = int8(price.IsOnlineSale)
			priceModel.IsDel = utils.NOT_DELETED
			priceModel.CreatedTime = currentUnixTime
			priceModel.UpdatedTime = currentUnixTime
			createPrices = append(createPrices,priceModel)
		} else {
			for _,currentPrice := range currentPrices {
				if uint(price.Id) == currentPrice.ID {
					currentPrice.CoachLevelID = uint(price.CoachLevel)
					currentPrice.SellStepStart = uint(price.SaleMin)
					currentPrice.SellStepEnd = uint(price.SaleMax)
					currentPrice.SellPrice = uint(price.Price)
					currentPrice.TransferType = int8(price.TransferType)
					currentPrice.TransferNum = uint(price.TransferNum)
					currentPrice.IsOnlineSale = int8(price.IsOnlineSale)
					currentPrice.UpdatedTime = currentUnixTime
					err = currentPrice.Save(ctx)
					if err != nil {
						utils.GetTraceLog(ctx).Error("update course price fail",zap.String("scope","SetCourseSalePrice"),
							zap.Any("course_price",currentPrice),zap.String("err_msg",err.Error()))
						return errors.NewFromCode(business_errors.ProductError_SET_SALE_PRICE_ERROR)
					}
					break
				}
			}
		}
	}

	for _,currentPrice := range currentPrices {
		needDelete := true
		for _,price := range prices {
			if currentPrice.ID == uint(price.Id) {
				needDelete = false
				break
			}
		}
		if needDelete {
			deleteIds = append(deleteIds,currentPrice.ID)
		}
	}

	err = productModel.PersonalCoursePriceBatchDelete(ctx,deleteIds)
	if err != nil {
		utils.GetTraceLog(ctx).Error("batch delete course price fail",zap.String("scope","SetCourseSalePrice"),
			zap.Uint("course_id",uint(request.CourseId)),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_SALE_PRICE_ERROR)
	}

	err = productModel.PersonalCoursePriceBatchInsert(ctx,createPrices)
	if err != nil {
		utils.GetTraceLog(ctx).Error("batch insert course price fail",zap.String("scope","SetCourseSalePrice"),
			zap.Uint("course_id",uint(request.CourseId)),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_SALE_PRICE_ERROR)
	}

	return nil
}

//更新课程门店设置方式 1全店 2指定门店
func UpdateCourseShopSetting(ctx context.Context, request *product.SetCourseShopsRequest) (*productModel.PersonalCourseTemplate, int8, error) {
	//查询当前课程详情
	course, err := productModel.PersonalCourseTemplateByID(ctx,uint(request.CourseId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get personal course fail",zap.String("scope","UpdateCourseShopSetting"),
			zap.Uint("course_id",uint(request.CourseId)),zap.String("err_msg",err.Error()))
		return nil,utils.IntZero,errors.NewFromCode(business_errors.ProductError_UPDATE_PERSONAL_COURSE_ERROR)
	}

	currentShopSetting := course.ShopSetting
	if currentShopSetting != int8(request.ShopSetting) {
		course.ShopSetting = int8(request.ShopSetting)
		course.UpdatedTime = utils.GetCurrentUnixTime()
		err = course.Update(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update shop setting fail",zap.String("scope","UpdateCourseShopSetting"),
				zap.Uint("course_id",uint(request.CourseId)),zap.String("err_msg",err.Error()))
			return nil,utils.IntZero,errors.NewFromCode(business_errors.ProductError_UPDATE_PERSONAL_COURSE_ERROR)
		}
	}

	return course,currentShopSetting,nil
}

//更新课程门店设置方式 1全店 2指定门店
func UpdateCourseSalePriceSetting(ctx context.Context, courseId uint, priceSetting int8) (*productModel.PersonalCourseTemplate,int8,error) {
	//查询当前课程详情
	course, err := productModel.PersonalCourseTemplateByID(ctx,courseId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get personal course fail",zap.String("scope","UpdateCourseSalePriceSetting"),
			zap.Uint("course_id",courseId),zap.String("err_msg",err.Error()))
		return nil,utils.IntZero,errors.NewFromCode(business_errors.ProductError_UPDATE_PERSONAL_COURSE_ERROR)
	}
	currentPriceSetting := course.PriceSetting
	course.PriceSetting = priceSetting
	course.UpdatedTime = utils.GetCurrentUnixTime()
	err = course.Update(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update shop setting fail",zap.String("scope","UpdateCourseShopSetting"),
			zap.Uint("course_id",courseId),zap.Int8("price_setting",priceSetting),zap.String("err_msg",err.Error()))
		return nil,utils.IntZero,errors.NewFromCode(business_errors.ProductError_UPDATE_PERSONAL_COURSE_ERROR)
	}

	return course,currentPriceSetting,errors.NewFromCode(business_errors.ProductError_UPDATE_PERSONAL_COURSE_ERROR)
}

//处理课程支持门店逻辑
func HandleSupportShops(ctx context.Context, course *productModel.PersonalCourseTemplate, currentShopSetting int8, shopIds []uint32) error {
	//查询当前已存在关联关系
	currentSupportShops,err := productModel.ProductSupportShopByProductId(ctx,course.ID, productModel.PRODUCT_PERSONAL_COURSE)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get current support shop fail",zap.String("scope","HandleSupportShops"),
			zap.Uint("course_id",course.ID),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_COURSE_SHOPS_ERROR)
	}

	if productModel.SHOP_SETTING_ALL == currentShopSetting {
		err = productModel.DeleteAllSupportShopByProductId(ctx,course.ID, productModel.PRODUCT_PERSONAL_COURSE)
		if err != nil {
			utils.GetTraceLog(ctx).Error("delete all support shops fail",zap.String("scope","HandleSupportShops"),
				zap.Uint("course_id",course.ID),zap.String("err_msg",err.Error()))
			return errors.NewFromCode(business_errors.ProductError_UPDATE_PERSONAL_COURSE_ERROR)
		}
	}else {
		currentShopIds := GetCurrentIds(ctx,currentSupportShops)
		GetDiffIds(ctx,currentShopIds,shopIds)
		createIds,deleteIds := GetDiffIds(ctx,currentShopIds,shopIds)
		err = InsertNewSupportShops(ctx,createIds,uint(course.BrandID),course.ID, productModel.COURSE_PERSONAL)
		if err != nil {
			utils.GetTraceLog(ctx).Error("insert new support shops fail",zap.String("scope","HandleSupportShops"),
				zap.Uint("course_id",course.ID),zap.String("err_msg",err.Error()))
			return errors.NewFromCode(business_errors.ProductError_SET_COURSE_SHOPS_ERROR)
		}
		err = DeleteUnsupportShops(ctx,deleteIds,course.ID, productModel.COURSE_PERSONAL)
		if err != nil {
			utils.GetTraceLog(ctx).Error("delete unsupport shops fail",zap.String("scope","HandleSupportShops"),
				zap.Uint("course_id",course.ID),zap.String("err_msg",err.Error()))
			return errors.NewFromCode(business_errors.ProductError_SET_COURSE_SHOPS_ERROR)
		}
	}
	return nil
}

//处理课程支持教练逻辑
func HandleSupportCoaches(ctx context.Context,course *productModel.PersonalCourseTemplate,coachIds []uint32) error {
	currentSupportCoaches, err := productModel.CourseCoachRelationByCourseID(ctx,course.ID)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get current support coaches fail",zap.String("scope","HandleSupportCoaches"),
			zap.Uint("course_id",course.ID),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_COURSE_COACH_ERROR)
	}

	currentSupportCoacheIds := GetCurrentIds(ctx,currentSupportCoaches)
	createIds,deleteIds := GetDiffIds(ctx,currentSupportCoacheIds,coachIds)

	err = InsertNewSupportCoaches(ctx,createIds,uint(course.BrandID),uint(course.ID), productModel.COURSE_PERSONAL)
	if err != nil {
		utils.GetTraceLog(ctx).Error("insert new support coaches fail",zap.String("scope","HandleSupportCoaches"),
			zap.Uint("course_id",course.ID),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_COURSE_COACH_ERROR)
	}

	err = DeleteUnsupportCoaches(ctx,deleteIds,uint(course.ID), productModel.COURSE_PERSONAL)
	if err != nil {
		utils.GetTraceLog(ctx).Error("delete unsupport coaches fail",zap.String("scope","HandleSupportCoaches"),
			zap.Uint("course_id",course.ID),zap.String("err_msg",err.Error()))
		return errors.NewFromCode(business_errors.ProductError_SET_COURSE_COACH_ERROR)
	}

	return  nil
}

//插入比对后需新增的教练
func InsertNewSupportCoaches(ctx context.Context,createIds []uint,brandId uint,courseId uint,courseType int8) error {
	createSupportCoaches := make([]*productModel.CourseCoachRelation,len(createIds))
	currentUinxTime := utils.GetCurrentUnixTime()
	for index, createCoachId := range createIds {
		ccr := &productModel.CourseCoachRelation{}
		ccr.BrandID = brandId
		ccr.CourseID = courseId
		ccr.CoachID = createCoachId
		ccr.CourseType = courseType
		ccr.IsDel = utils.NOT_DELETED
		ccr.CreatedTime = currentUinxTime
		ccr.UpdatedTime = currentUinxTime
		createSupportCoaches[index] = ccr
	}

	err := productModel.CourseCoachRelationBatchInsert(ctx,createSupportCoaches)
	if err != nil {
		return errors.NewFromCode(business_errors.ProductError_SET_COURSE_COACH_ERROR)
	}

	return nil
}

//删除不再支持的教练
func DeleteUnsupportCoaches(ctx context.Context,deleteIds []uint,courseId uint,courseType int8) error {
	err := productModel.CourseCoachRelationBatchDelete(ctx,courseId,courseType,deleteIds)
	if err != nil {
		return errors.NewFromCode(business_errors.ProductError_SET_COURSE_COACH_ERROR)
	}

	return nil
}

//插入比对后需新增的门店
func InsertNewSupportShops(ctx context.Context,createIds []uint,brandId uint,productId uint,courseType int8) error {
	createSupportShops := make([]*productModel.ProductSupportShop,len(createIds))
	currentUinxTime := utils.GetCurrentUnixTime()
	for index, createShopId := range createIds {
		pss := &productModel.ProductSupportShop{}
		pss.BrandID = brandId
		pss.ShopID = createShopId
		pss.ProductID = productId
		pss.ProductType = courseType
		pss.IsDel = utils.NOT_DELETED
		pss.CreatedTime = currentUinxTime
		pss.UpdatedTime = currentUinxTime
		createSupportShops[index] = pss
	}

	err := productModel.ProductSupportShopBatchInsert(ctx,createSupportShops)
	if err != nil {
		return errors.NewFromCode(business_errors.ProductError_SET_COURSE_SHOPS_ERROR)
	}

	return nil
}

//删除不再支持的门店
func DeleteUnsupportShops(ctx context.Context,deleteIds []uint,productId uint,courseType int8) error {
	err := productModel.ProductSupportShopBatchDelete(ctx,productId,courseType,deleteIds)
	if err != nil {
		return errors.NewFromCode(business_errors.ProductError_SET_COURSE_SHOPS_ERROR)
	}

	return nil
}

//获取当前支持的ID列表
func GetCurrentIds(ctx context.Context,m interface{}) []uint32 {
	ids := make([]uint32,0)
	models,ok := m.([]interface{})
	if !ok {
		utils.GetTraceLog(ctx).Error("assert interface to slice interface fail",zap.Any("object",m))
		return ids
	}

	for _, model := range models {
		id := reflect.ValueOf(model).FieldByName("ID").Uint()
		ids = append(ids,uint32(id))
	}

	return ids
}