package product

import (
	"go.7yes.com/j7f/components/errors"
	"j7go/errors"
	"j7go/proto/product"
	"go.uber.org/zap"
	productService "j7go/services/product"
	"j7go/utils"
)

type TeamCourseTplServer struct {
}

//新增团体课
func (t *TeamCourseTplServer) AddCourse(srv product.TeamCourseSrv_AddCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("add_receive", err.Error()))
			return err
		}

		courseId, err := productService.AddTeamCourseTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.Any("add_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_ADD_TEAM_COURSE_ERROR)
		}

		res := &product.CourseTplResponse{
			Status:   errors.GetResHeader(err),
			CourseId: uint32(courseId),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("add_send", err.Error()))
			return err
		}
	}
}

//编辑团体课
func (t *TeamCourseTplServer) EditCourse(srv product.TeamCourseSrv_EditCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("edit_receive", err.Error()))
			return err
		}

		courseId, err := productService.EditTeamCourseTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.Any("edit_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_EDIT_TEAM_COURSE_ERROR)
		}

		res := &product.CourseTplResponse{
			Status:   errors.GetResHeader(err),
			CourseId: uint32(courseId),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("edit_send", err.Error()))
			return err
		}
	}
}

//设置团体课支持门店
func (t *TeamCourseTplServer) SetCourseShops(srv product.TeamCourseSrv_SetCourseShopsServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("set_course_receive", err.Error()))
			return err
		}

		err = productService.SetTeamCourseTplShops(srv.Context(), params, false)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.Any("set_course_shops", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_SET_TEAM_COURSE_SHOPS_ERROR)
		}

		res := &product.CourseTplResponse{
			Status: errors.GetResHeader(err),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("set_course_send", err.Error()))
			return err
		}
	}
}

//获取团体课详情
func (t *TeamCourseTplServer) GetCourseInfo(srv product.TeamCourseSrv_GetCourseInfoServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("get_course_receive", err.Error()))
			return err
		}

		res, err := productService.GetTeamCourseTplInfo(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.Any("get_course_info", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_GET_TEAM_COURSE_INFO_ERROR)
		}

		res.Status = errors.GetResHeader(err)
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("get_course_send", err.Error()))
			return err
		}
	}
}

//将门店团体课转入品牌
func (t *TeamCourseTplServer) IntoBrandCourse(srv product.TeamCourseSrv_IntoBrandCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("into_brand_course_receive", err.Error()))
			return err
		}

		err = productService.SetTeamCourseTplShops(srv.Context(), params, true)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.Any("into_brand_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_INTO_BRAND_COURSE_ERROR)
		}

		res := &product.CourseTplResponse{
			Status: errors.GetResHeader(err),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("into_brand_course_send", err.Error()))
			return err
		}
	}
}

//删除团体课
func (t *TeamCourseTplServer) DelCourse(srv product.TeamCourseSrv_DelCourseServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("del_course_receive", err.Error()))
			return err
		}

		err = productService.DelTeamCourseTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.Any("del_course", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_DEL_TEAM_COURSE_ERROR)
		}

		res := &product.CourseTplResponse{
			Status: errors.GetResHeader(err),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("del_course_send", err.Error()))
			return err
		}
	}
}

//将团体课置为无效
func (t *TeamCourseTplServer) SetCourseInvalid(srv product.TeamCourseSrv_SetCourseInvalidServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("set_course_invalid_receive", err.Error()))
			return err
		}

		err = productService.SetTeamCourseTplState(srv.Context(), params, false)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.Any("set_course_invalid", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_DEL_TEAM_COURSE_ERROR)
		}

		res := &product.CourseTplResponse{
			Status: errors.GetResHeader(err),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("set_course_invalid_send", err.Error()))
			return err
		}
	}
}

//将团体课置为有效
func (t *TeamCourseTplServer) SetCourseValid(srv product.TeamCourseSrv_SetCourseValidServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("set_course_valid_receive", err.Error()))
			return err
		}

		err = productService.SetTeamCourseTplState(srv.Context(), params, true)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.Any("set_course_valid", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_DEL_TEAM_COURSE_ERROR)
		}

		res := &product.CourseTplResponse{
			Status: errors.GetResHeader(err),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("team_course_tpl", zap.String("set_course_valid_send", err.Error()))
			return err
		}
	}
}
