package product

import (
	"context"
	"go.7yes.com/j7f/proto/images"
	product2 "go.7yes.com/j7f/proto/product"
	"go.uber.org/zap"
	"time"
	"j7go/components"
	"j7go/models/product"
	imagesService "j7go/services/images"
	"j7go/utils"
)

func AddPackageCourse(ctx context.Context, img *images.CommonAlbumImageRequest, c *product.PackageCourseTemplate, t []uint, p []*product2.PackageCoursePersonalSetting) (uint, error) {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_package_course", zap.String("begin_transaction", err.Error()))
		return utils.IntZero, err
	}

	//新增图片
	image := imagesService.FormatImageInfo(img)
	albumId, err := imagesService.InsertImages(ctx, 0, image)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_package_course", zap.String("insert_images", err.Error()))
		_ = components.M.Rollback(ctx)
		return utils.IntZero, err
	}

	//新增课程包模板
	c.ImageID = uint(albumId)
	c.IsDel = utils.NOT_DELETED
	c.UpdatedTime = uint(time.Now().Unix())
	c.CreatedTime = uint(time.Now().Unix())
	err = c.Insert(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_package_course", zap.String("insert_package_course", err.Error()))
		_ = components.M.Rollback(ctx)
		return utils.IntZero, err
	}

	//新增课程包团体课
	if len(t) != utils.IntZero {
		err = AddPackageCourseTeams(ctx, t, c.BrandID, c.ShopID, c.ID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("add_package_course", zap.String("add_package_course_team", err.Error()))
			_ = components.M.Rollback(ctx)
			return utils.IntZero, err
		}
	}

	//新增课程包私教课
	if len(p) != utils.IntZero {
		err = AddPackageCoursePersonals(ctx, p, c.BrandID, c.ShopID, c.ID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("add_package_course", zap.String("add_package_course_personal", err.Error()))
			_ = components.M.Rollback(ctx)
			return utils.IntZero, err
		}
	}

	err = components.M.Commit(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("add_package_course", zap.String("commit", err.Error()))
		return utils.IntZero, err
	}
	return c.ID, nil
}

//转换课程包模板数据结构
func FormatPackageCourseTpl(t *product2.PackageCourseTplInfo, courseType int8) *product.PackageCourseTemplate {
	c := &product.PackageCourseTemplate{
		BrandID:           uint(t.BrandId),
		ShopID:            uint(t.ShopId),
		CourseName:        t.CourseName,
		Price:             uint(t.Price),
		PackageType:       courseType,
		TotalPrice:        uint(t.TotalPrice),
		TotalTimes:        uint(t.TotalTimes),
		TeamTimes:         uint(t.TeamTimes),
		TeamUnitPrice:     uint(t.TeamUnitPrice),
		Personal:          uint(t.Personal),
		PersonalUnitPrice: uint(t.PersonalUnitPrice),
		StartTime:         uint(t.StartTime),
		EndTime:           uint(t.EndTime),
		ValidDays:         uint(t.ValidDays),
		SaleMode:          int8(int32(t.SaleMode)),
		PublishChannel:    int8(int32(t.PublishChannel)),
		Intro:             t.Intro,
		Remarks:           t.Remarks,
	}

	return c
}

//添加课程包团体课
func AddPackageCourseTeams(ctx context.Context, courseIds []uint, brandId, shopId, packageCourseId uint) error {
	for _, courseId := range courseIds {
		t := &product.PackageCourseTeam{
			BrandID:         brandId,
			ShopID:          shopId,
			CoursePackageID: packageCourseId,
			CourseID:        courseId,
			IsDel:           utils.NOT_DELETED,
			UpdatedTime:     uint(time.Now().Unix()),
			CreatedTime:     uint(time.Now().Unix()),
		}
		err := t.Insert(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

//添加课程包私教课
func AddPackageCoursePersonals(ctx context.Context, courses []*product2.PackageCoursePersonalSetting, brandId, shopId, packageCourseId uint) error {
	for _, course := range courses {
		p := &product.PackageCoursePersonal{
			BrandID:         brandId,
			ShopID:          shopId,
			CoursePackageID: packageCourseId,
			CourseID:        uint(course.CourseId),
			IsDel:           utils.NOT_DELETED,
			UpdatedTime:     uint(time.Now().Unix()),
			CreatedTime:     uint(time.Now().Unix()),
		}
		err := p.Insert(ctx)
		if err != nil {
			return err
		}

		err = AddCoachLevels(ctx, course.CoachLevelIds, uint(course.CourseId), packageCourseId)
		if err != nil {
			utils.GetTraceLog(ctx).Error("add_package_course", zap.String("add_package_course_personal_level", err.Error()))
			return err
		}
	}

	return nil
}

//添加课程包私教课私教等级配置
func AddCoachLevels(ctx context.Context, coachLevelIds []uint32, courseId, packageCourseId uint) error {
	for _, coachLevelId := range coachLevelIds {
		l := &product.PackageCourseCoachLevel{
			CoursePackageID: packageCourseId,
			CoachLevelID:    uint(coachLevelId),
			CourseID:        uint(courseId),
			IsDel:           utils.NOT_DELETED,
			UpdatedTime:     uint(time.Now().Unix()),
			CreatedTime:     uint(time.Now().Unix()),
		}
		err := l.Insert(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func EditPackageCourse(ctx context.Context, img *images.CommonAlbumImageRequest, c *product.PackageCourseTemplate, t []uint, p []*product2.PackageCoursePersonalSetting) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("begin_transaction", err.Error()))
		return err
	}

	if img != nil {
		//修改图片
		image := imagesService.FormatImageInfo(img)
		err := imagesService.UpdateAlbumImages(ctx, uint(img.AlbumId), image)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_package_course", zap.String("update_images", err.Error()))
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//修改课程包模板
	c.IsDel = utils.NOT_DELETED
	c.UpdatedTime = uint(time.Now().Unix())
	err = c.Update(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("update_package_course", err.Error()))
		_ = components.M.Rollback(ctx)
		return err
	}

	//修改课程包团体课
	if len(t) != utils.IntZero {
		err = EditPackageCourseTeams(ctx, t, c.BrandID, c.ShopID, c.ID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_package_course", zap.String("update_package_course_team", err.Error()))
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	//修改课程包私教课
	if len(p) != utils.IntZero {
		err = EditPackageCoursePersonals(ctx, p, c.BrandID, c.ShopID, c.ID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_package_course", zap.String("update_package_course_personal", err.Error()))
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	err = components.M.Commit(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("commit", err.Error()))
		return err
	}
	return nil
}

//修改课程包团体课
func EditPackageCourseTeams(ctx context.Context, newCourseIds []uint, brandId, shopId, packageCourseId uint) error {
	courseTeams, err := product.PackageCourseTeamByPackageCourseIDs(ctx, packageCourseId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("get_course_teams", err.Error()))
		return err
	}

	courseIds := make([]uint32, len(courseTeams))
	for index, course := range courseTeams {
		courseIds[index] = uint32(course.ID)
	}

	newCourseIdsUint32 := make([]uint32, len(newCourseIds))
	for index, courseId := range newCourseIds {
		newCourseIdsUint32[index] = uint32(courseId)
	}

	createIds, deleteIds := GetDiffIds(ctx, courseIds, newCourseIdsUint32)

	err = DelCourseTeam(ctx, deleteIds, packageCourseId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("del_course_teams", err.Error()))
		return err
	}

	err = AddPackageCourseTeams(ctx, createIds, brandId, shopId, packageCourseId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("add_course_teams", err.Error()))
		return err
	}

	return nil
}

//删除课程包团体课
func DelCourseTeam(ctx context.Context, courseIds []uint, packageCourseId uint) error {
	return product.DelPackageCourseTeamByPackageCourseIDAndCourseId(ctx, packageCourseId, courseIds)
}

//修改课程包私教课
func EditPackageCoursePersonals(ctx context.Context, courses []*product2.PackageCoursePersonalSetting, brandId, shopId, packageCourseId uint) error {
	coursePersonals, err := product.PackageCoursePersonalByPackageCourseIDs(ctx, packageCourseId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("get_course_personals", err.Error()))
		return err
	}

	courseIds := make([]uint32, len(coursePersonals))
	for index, course := range coursePersonals {
		courseIds[index] = uint32(course.ID)
	}

	newCourseIdsUint32 := make([]uint32, len(courses))
	for index, course := range courses {
		newCourseIdsUint32[index] = uint32(course.CourseId)
	}

	createIds, deleteIds := GetDiffIds(ctx, courseIds, newCourseIdsUint32)

	err = DelCoursePersonals(ctx, deleteIds, packageCourseId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("del_course_personal", err.Error()))
		return err
	}

	err = DelCoachLevels(ctx, packageCourseId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("del_course_personal_coach_level", err.Error()))
		return err
	}

	createCourses := make([]*product2.PackageCoursePersonalSetting, len(createIds))
	for index, courseId := range createIds {
		for _, course := range courses {
			if uint(course.CourseId) == courseId {
				createCourses[index] = course
			}
		}
	}

	err = AddPackageCoursePersonals(ctx, createCourses, brandId, shopId, packageCourseId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("update_package_course", zap.String("add_course_teams", err.Error()))
		return err
	}
	return nil
}

//添加课程包私教课私教等级配置
func DelCoursePersonals(ctx context.Context, courseIds []uint, packageCourseId uint) error {
	return product.DelPackageCoursePersonalByPackageCourseIDAndCourseId(ctx, packageCourseId, courseIds)
}

//删除私教课教练等级
func DelCoachLevels (ctx context.Context, packageCourseId uint) error {
	return product.DelPackageCourseCoachLevelByPackageCourseID(ctx, packageCourseId)
}