package product

import (
	"go.7yes.com/j7f/components/errors"
	business_errors "go.7yes.com/j7f/errors"
	"go.7yes.com/j7f/proto/product"
	"go.uber.org/zap"
	packageCourse "j7go/models/product"
	product2 "j7go/services/product"
	"j7go/utils"
)

type packageCourseSrv struct {
}

//新增不限课程课程包
func (p *packageCourseSrv) AddUnlimitedPackageCourse(srv product.PackageCourseTplServer_AddUnlimitedPackageCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("add_unlimited_receive", err.Error()))
			return err
		}

		//转换课程包请求数据结构为model
		c := product2.FormatPackageCourseTpl(params.CourseTplInfo, packageCourse.UNLIMITED_COURSE)
		c.IsTeam = utils.Available
		c.IsPersonal = utils.Available

		//新增课程包
		packageCourseId, err := product2.AddPackageCourse(srv.Context(), params.CourseTplInfo.CourseImg, c, nil, nil)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.Any("add_package_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_ADD_PACKAGE_COURSE_ERROR)
		}

		res := &product.PackageCourseTplResponse{
			Status:          errors.GetResHeader(err),
			PackageCourseId: uint32(packageCourseId),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("add_unlimited_send", err.Error()))
			return err
		}
	}
}

//新增范围内课程课程包
func (p *packageCourseSrv) AddScopePackageCourse(srv product.PackageCourseTplServer_AddScopePackageCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("add_scope_receive", err.Error()))
			return err
		}

		//转换课程包请求数据结构为model
		c := product2.FormatPackageCourseTpl(params.CourseTplInfo, packageCourse.SCOPE_COURSE)

		t := make([]uint, len(params.CourseTeamInfo.CourseIds))
		if params.CourseTeamInfo != nil {
			c.IsTeam = utils.Available
			for index, courseId:= range params.CourseTeamInfo.CourseIds{
				t[index] = uint(courseId)
			}
		}

		p := make([]*product.PackageCoursePersonalSetting, 0)
		if params.CoursePersonalInfo != nil {
			c.IsPersonal = utils.Available
			p = params.CoursePersonalInfo.Courses
		}

		//新增课程包
		packageCourseId, err := product2.AddPackageCourse(srv.Context(), params.CourseTplInfo.CourseImg, c, t, p)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.Any("add_scope_package_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_ADD_PACKAGE_COURSE_ERROR)
		}

		res := &product.PackageCourseTplResponse{
			Status:          errors.GetResHeader(err),
			PackageCourseId: uint32(packageCourseId),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("add_scope_send", err.Error()))
			return err
		}
	}
}

//新增固定课程课程包
func (p *packageCourseSrv) AddFixedPackageCourse(srv product.PackageCourseTplServer_AddFixedPackageCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("add_fixed_receive", err.Error()))
			return err
		}

		//转换课程包请求数据结构为model
		c := product2.FormatPackageCourseTpl(params.CourseTplInfo, packageCourse.FIXED_COURSE)

		t := make([]uint, len(params.CourseTeamInfo.CourseIds))
		if params.CourseTeamInfo != nil {
			c.IsTeam = utils.Available
			for index, courseId:= range params.CourseTeamInfo.CourseIds{
				t[index] = uint(courseId)
			}
		}

		p := make([]*product.PackageCoursePersonalSetting, 0)
		if params.CoursePersonalInfo != nil {
			c.IsPersonal = utils.Available
			p = params.CoursePersonalInfo.Courses
		}

		//新增课程包
		packageCourseId, err := product2.AddPackageCourse(srv.Context(), params.CourseTplInfo.CourseImg, c, t, p)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.Any("add_fixed_package_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_ADD_PACKAGE_COURSE_ERROR)
		}

		res := &product.PackageCourseTplResponse{
			Status:          errors.GetResHeader(err),
			PackageCourseId: uint32(packageCourseId),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("add_fixed_send", err.Error()))
			return err
		}
	}
}

//修改不限课程课程包
func (p *packageCourseSrv) EditUnlimitedPackageCourse(srv product.PackageCourseTplServer_EditUnlimitedPackageCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("edit_unlimited_receive", err.Error()))
			return err
		}

		//转换课程包请求数据结构为model
		c := product2.FormatPackageCourseTpl(params.CourseTplInfo, packageCourse.UNLIMITED_COURSE)
		c.ID = uint(params.PackageCourseId)
		c.IsTeam = utils.Available
		c.IsPersonal = utils.Available

		//修改课程包
		err = product2.EditPackageCourse(srv.Context(), params.CourseTplInfo.CourseImg, c, nil, nil)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.Any("edit_unlimited_package_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_EDIT_PACKAGE_COURSE_ERROR)
		}

		res := &product.PackageCourseTplResponse{
			Status:          errors.GetResHeader(err),
			PackageCourseId: params.PackageCourseId,
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("edit_unlimited_send", err.Error()))
			return err
		}
	}
}

//修改限定课程课程包
func (p *packageCourseSrv) EditScopePackageCourse(srv product.PackageCourseTplServer_EditScopePackageCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("edit_scope_receive", err.Error()))
			return err
		}

		//转换课程包请求数据结构为model
		c := product2.FormatPackageCourseTpl(params.CourseTplInfo, packageCourse.SCOPE_COURSE)

		t := make([]uint, len(params.CourseTeamInfo.CourseIds))
		if params.CourseTeamInfo != nil {
			c.IsTeam = utils.Available
			for index, courseId:= range params.CourseTeamInfo.CourseIds{
				t[index] = uint(courseId)
			}
		}

		p := make([]*product.PackageCoursePersonalSetting, 0)
		if params.CoursePersonalInfo != nil {
			c.IsPersonal = utils.Available
			p = params.CoursePersonalInfo.Courses
		}

		//修改课程包
		err = product2.EditPackageCourse(srv.Context(), params.CourseTplInfo.CourseImg, c, t, p)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.Any("edit_scope_package_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_EDIT_PACKAGE_COURSE_ERROR)
		}

		res := &product.PackageCourseTplResponse{
			Status:          errors.GetResHeader(err),
			PackageCourseId: params.PackageCourseId,
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("edit_scope_send", err.Error()))
			return err
		}
	}
}

//修改指定课程课程包
func (p *packageCourseSrv) EditFixedPackageCourse(srv product.PackageCourseTplServer_EditFixedPackageCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("edit_fixed_receive", err.Error()))
			return err
		}

		//转换课程包请求数据结构为model
		c := product2.FormatPackageCourseTpl(params.CourseTplInfo, packageCourse.FIXED_COURSE)

		t := make([]uint, len(params.CourseTeamInfo.CourseIds))
		if params.CourseTeamInfo != nil {
			c.IsTeam = utils.Available
			for index, courseId:= range params.CourseTeamInfo.CourseIds{
				t[index] = uint(courseId)
			}
		}

		p := make([]*product.PackageCoursePersonalSetting, 0)
		if params.CoursePersonalInfo != nil {
			c.IsPersonal = utils.Available
			p = params.CoursePersonalInfo.Courses
		}

		//修改课程包
		err = product2.EditPackageCourse(srv.Context(), params.CourseTplInfo.CourseImg, c, t, p)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.Any("edit_fixed_package_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_EDIT_PACKAGE_COURSE_ERROR)
		}

		res := &product.PackageCourseTplResponse{
			Status:          errors.GetResHeader(err),
			PackageCourseId: params.PackageCourseId,
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("package_course_tpl", zap.String("edit_fixed_send", err.Error()))
			return err
		}
	}
}