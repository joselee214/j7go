package product

import (
	"context"
	"github.com/joselee214/j7f/proto/common"
	"j7go/proto/images"
	"j7go/proto/product"
	"go.uber.org/zap"
	"time"
	"j7go/components"
	imageModel "j7go/models/images"
	productModel "j7go/models/product"
	"j7go/services/images"
	"j7go/utils"
)

func AddTeamCourseTpl(ctx context.Context, c *product.CourseTplRequest) (id uint, err error) {
	ctx, err = components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return utils.IntZero, err
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	imagesInfo := make([]*imagesService.ImageInfo, utils.ONE)
	imagesInfo[0] = &imagesService.ImageInfo{
		ImageId:   c.CourseImg.ImageId,
		ImageUrl:  c.CourseImg.ImageUrl,
		CoverType: int8(imageModel.COVER_TYPE_COVER),
	}
	albumId, err := imagesService.InsertImages(ctx, utils.IntZero, imagesInfo)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_course_image", zap.Any("image", imagesInfo), zap.Error(err))
		return utils.IntZero, err
	}

	teamCourseTpl := &productModel.TeamCourseTemplate{}
	teamCourseTpl.AlbumID = uint(albumId)
	teamCourseTpl.BrandID = uint(c.BrandId)
	teamCourseTpl.ShopID = uint(c.ShopId)
	teamCourseTpl.CourseName = c.CourseName
	teamCourseTpl.CategoryID = uint(c.CourseCategory)
	teamCourseTpl.Duration = uint(c.Duration)
	teamCourseTpl.TimeUnit = int8(c.TimeUnit)
	teamCourseTpl.Price = uint(c.Price)
	teamCourseTpl.StrengthLevel = int8(c.StrengthLevel)
	teamCourseTpl.Calories = uint(c.Calories)
	teamCourseTpl.Description = c.Description
	teamCourseTpl.PublishChannel = int8(c.PublishChannel)
	teamCourseTpl.OperatorID = uint(c.OperatorId)
	teamCourseTpl.IsAvailable = utils.Available
	teamCourseTpl.CreatedTime = uint(time.Now().Unix())
	teamCourseTpl.UpdatedTime = uint(time.Now().Unix())
	err = teamCourseTpl.Save(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_course", zap.Any("course", teamCourseTpl), zap.Error(err))
		return utils.IntZero, err
	}

	err = CourseSettingRelationBatchInsert(ctx, c.TrainAim, teamCourseTpl.ID, productModel.COURSE_TEAM)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_course_train_aim", zap.Uint("course_id", teamCourseTpl.ID),
			zap.Any("train_aim", c.TrainAim), zap.Error(err))
		return utils.IntZero, err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		return utils.IntZero, err
	}

	return teamCourseTpl.ID, nil
}

func EditTeamCourseTpl(ctx context.Context, c *product.CourseTplRequest) (id uint, err error) {
	ctx, err = components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return utils.IntZero, err
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	teamCourseTpl, err := productModel.TeamCourseTemplateByID(ctx, uint(c.CourseId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_course", zap.Uint32("course_id", c.CourseId), zap.Error(err))
		return utils.IntZero, err
	}
	imagesInfo := make([]*imagesService.ImageInfo, utils.Available)
	imagesInfo[0] = &imagesService.ImageInfo{
		ImageId:   c.CourseImg.ImageId,
		ImageUrl:  c.CourseImg.ImageUrl,
		CoverType: int8(imageModel.COVER_TYPE_COVER),
	}
	err = imagesService.UpdateAlbumImages(ctx, teamCourseTpl.AlbumID, imagesInfo)
	if err != nil {
		utils.GetTraceLog(ctx).Error("edit_course_image", zap.Uint("album_id", teamCourseTpl.AlbumID),
			zap.Any("image", imagesInfo), zap.Error(err))
		return utils.IntZero, err
	}

	teamCourseTpl.CategoryID = uint(c.CourseCategory)
	teamCourseTpl.Duration = uint(c.Duration)
	teamCourseTpl.TimeUnit = int8(c.TimeUnit)
	teamCourseTpl.Price = uint(c.Price)
	teamCourseTpl.StrengthLevel = int8(c.StrengthLevel)
	teamCourseTpl.Calories = uint(c.Calories)
	teamCourseTpl.Description = c.Description
	teamCourseTpl.OperatorID = uint(c.OperatorId)
	teamCourseTpl.IsAvailable = int8(c.IsAvailable)
	teamCourseTpl.UpdatedTime = uint(time.Now().Unix())
	err = teamCourseTpl.Update(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("edit_course", zap.Any("course", teamCourseTpl), zap.Error(err))
		return utils.IntZero, err
	}

	courseSettings, err := productModel.GetCourseTrainAimsById(ctx, uint(c.CourseId), productModel.COURSE_TEAM)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_course_train_aim", zap.Uint32("course_id", c.CourseId),
			zap.Error(err))
		return utils.IntZero, err
	}

	trainAim := make([]uint32, len(courseSettings))
	for index, courseSetting := range courseSettings {
		trainAim[index] = uint32(courseSetting.SettingID)
	}

	delAim, insertAim := CompareUpdate(trainAim, c.TrainAim)

	err = CourseSettingRelationBatchInsert(ctx, insertAim, teamCourseTpl.ID, productModel.COURSE_TEAM)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_course_train_aim", zap.Uint("course_id", teamCourseTpl.ID),
			zap.Any("train_aim", insertAim), zap.Error(err))
		return utils.IntZero, err
	}

	err = CourseSettingRelationBatchDelete(ctx, delAim, teamCourseTpl.ID, productModel.COURSE_TEAM)
	if err != nil {
		utils.GetTraceLog(ctx).Error("del_course_train_aim", zap.Uint("course_id", teamCourseTpl.ID),
			zap.Any("train_aim", delAim), zap.Error(err))
		return utils.IntZero, err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		return utils.IntZero, err
	}

	return teamCourseTpl.ID, nil
}

func SetTeamCourseTplShops(ctx context.Context, p *product.SetCourseShopsRequest, isIntoBrand bool) error {
	ctx, err := components.M.BeginTransaction(ctx, time.Second*2)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	course, err := productModel.TeamCourseTemplateByID(ctx, uint(p.CourseId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_course", zap.Uint32("course_id", p.CourseId), zap.Error(err))
		return err
	}

	if isIntoBrand {
		course.PublishChannel = int8(product.PublishChannel_CHANNEL_BRAND)
		course.ShopID = utils.IntZero
	}

	if course.ShopSetting != int8(p.ShopSetting) {
		course.ShopSetting = int8(p.ShopSetting)
	}

	course.OperatorID = uint(p.OperatorId)
	course.UpdatedTime = utils.GetCurrentUnixTime()
	err = course.Update(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("save_course", zap.Any("course", course), zap.Error(err))
		return err
	}

	shops, err := productModel.ProductSupportShopByProductId(ctx, course.ID, productModel.PRODUCT_TEAM_COURSE)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_support_shop", zap.Uint("course_id", course.ID), zap.Error(err))
		return err
	}

	if productModel.SHOP_SETTING_ALL == course.ShopSetting {
		err = productModel.DeleteAllSupportShopByProductId(ctx, course.ID, productModel.PRODUCT_TEAM_COURSE)
		if err != nil {
			utils.GetTraceLog(ctx).Error("del_support_shop", zap.Uint("course_id", course.ID), zap.Error(err))
			return err
		}
	} else {
		currentShopIds := GetCurrentIds(ctx, shops)
		createIds, deleteIds := GetDiffIds(ctx, currentShopIds, p.ShopIds)
		err = InsertNewSupportShops(ctx, createIds, uint(course.BrandID), course.ID, productModel.PRODUCT_TEAM_COURSE)
		if err != nil {
			utils.GetTraceLog(ctx).Error("add_support_shop", zap.Uint("course_id", course.ID),
				zap.Uint("brand_id", course.BrandID), zap.Any("shop_ids", createIds), zap.Error(err))
			return err
		}
		err = DeleteUnsupportShops(ctx, deleteIds, course.ID, productModel.PRODUCT_TEAM_COURSE)
		if err != nil {
			utils.GetTraceLog(ctx).Error("del_support_shop", zap.Uint("course_id", course.ID),
				zap.Any("shop_ids", deleteIds), zap.Error(err))
			return err
		}
	}

	err = components.M.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetTeamCourseTplInfo(ctx context.Context, p *product.SetCourseStateRequest) (*product.GetCourseTplResponse, error) {
	info := &product.GetCourseTplResponse{}

	c, err := productModel.TeamCourseTemplateByID(ctx, uint(p.CourseId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_course_info", zap.Uint32("course_id", p.CourseId), zap.Error(err))
		return info, err
	}

	//获取支持门店
	if c.ShopSetting == productModel.SHOP_SETTING_ALL {
		shopList, err := GetSupportShop(ctx, c.ID, productModel.PRODUCT_TEAM_COURSE)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_support_shop", zap.Uint32("course_id", p.CourseId), zap.Error(err))
			return info, err
		}
		info.SupportShopList = shopList
	}

	//获取分类名称
	cc, err := productModel.CourseCategoryByID(ctx, c.CategoryID)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_course_category", zap.Uint("category_id", c.CategoryID), zap.Error(err))
		return info, err
	}

	//获取训练目的
	tagList, err := GetCourseTags(ctx, c.ID, int8(productModel.PRODUCT_TEAM_COURSE))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_course_tag", zap.Uint32("course_id", p.CourseId), zap.Error(err))
		return info, err
	}

	//获取课程图片
	img, err := imagesService.GetImages(ctx, uint(c.AlbumID), int8(imageModel.COVER_TYPE_COVER))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_course_img", zap.Uint32("album_id", p.CourseId), zap.Error(err))
		return info, err
	}

	if len(img) != utils.IntZero {
		info.CourseImg = &images.Image{
			ImageId:   uint32(img[0].ID),
			ImageUrl:  img[0].ImageURL,
			CoverType: uint32(img[0].CoverType),
		}
	}
	info.CourseId = uint32(c.ID)
	info.CourseName = c.CourseName
	info.CourseCategoryId = uint32(c.CategoryID)
	info.CourseCategoryName = cc.Name
	info.StrengthLevel = uint32(c.StrengthLevel)
	info.Calories = uint32(c.Calories)
	info.TimeUnit = uint32(c.TimeUnit)
	info.Duration = uint32(c.Duration)
	info.Price = uint32(c.Price)
	info.Description = c.Description
	info.TrainAimList = tagList
	info.IsAvailable = common.IsAvailable(c.IsAvailable)

	return info, nil
}

func DelTeamCourseTpl(ctx context.Context, p *product.SetCourseStateRequest) error {
	course := &productModel.TeamCourseTemplate{
		ID: uint(p.CourseId),
	}

	err := course.Delete(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("del_course", zap.Uint32("course_id", p.CourseId), zap.Error(err))
		return err
	}

	return nil
}

func SetTeamCourseTplState(ctx context.Context, p *product.SetCourseStateRequest, isAvailable bool) error {
	course, err := productModel.TeamCourseTemplateByID(ctx, uint(p.CourseId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_course", zap.Uint32("course_id", p.CourseId), zap.Error(err))
		return err
	}

	course.IsAvailable = int8(utils.Bool2Int(isAvailable))
	course.OperatorID = uint(p.OperatorId)
	course.UpdatedTime = utils.GetCurrentUnixTime()
	err = course.Update(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("set_course_state", zap.Any("course", course), zap.Error(err))
		return err
	}

	return nil
}
