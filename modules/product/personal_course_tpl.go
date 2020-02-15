package product

import (
	"go.7yes.com/j7f/components/errors"
	"j7go/proto/product"
	"go.uber.org/zap"
	individual "j7go/services/product"
	"j7go/utils"
)

type personalCourseSrv struct {
}

func (p *personalCourseSrv) AddCourse(srv product.PersonalCourseSrv_AddCourseServer) error {
	for {
		courseInfo, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_course_tpl", zap.String("receive_err", err.Error()))
			return err
		}

		courseId, err := individual.AddPersonalCourse(srv.Context(), courseInfo)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("add personal course fail", zap.String("err_msg", err.Error()))
		}

		res := &product.PersonalCourseResponse{}
		res.CourseId = courseId
		res.Status = errors.GetResHeader(err)

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("response add personal fail", zap.String("err_msg", err.Error()))
			return err
		}
	}
}

func (p *personalCourseSrv) EditCourse(srv product.PersonalCourseSrv_EditCourseServer) error {
	for {
		courseInfo, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_course_tpl", zap.String("receive_err", err.Error()))
			return err
		}

		courseId, err := individual.UpdatePersonalCourse(srv.Context(), courseInfo)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("update personal course fail", zap.String("err_msg", err.Error()))
		}

		res := &product.PersonalCourseResponse{}
		res.CourseId = courseId
		res.Status = errors.GetResHeader(err)

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("response update personal fail", zap.String("err_msg", err.Error()))
			return err
		}
	}
}

func (p *personalCourseSrv) SetCourseShops(srv product.PersonalCourseSrv_SetCourseShopsServer) error {
	for {
		shopsAndCoachs, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("set course shops fail", zap.String("receive_err", err.Error()))
			return err
		}

		err = individual.SetCourseShopsAndCoaches(srv.Context(), shopsAndCoachs)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("set course shops fail", zap.String("err_msg", err.Error()))
		}

		res := &product.PersonalCourseResponse{}
		res.CourseId = shopsAndCoachs.CourseId
		res.Status = errors.GetResHeader(err)

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("response set course shops fail", zap.String("err_msg", err.Error()))
			return err
		}
	}
}

func (p *personalCourseSrv) SetSalePrice(srv product.PersonalCourseSrv_SetSalePriceServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("set sale price fail", zap.String("recv_err", err.Error()))
			return err
		}

		err = individual.SetCourseSalePrice(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("set sale price fail", zap.String("err_msg", err.Error()))
		}

		res := &product.PersonalCourseResponse{}
		res.CourseId = request.CourseId
		res.Status = errors.GetResHeader(err)

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("response set course price fail", zap.String("err_msg", err.Error()))
			return err
		}
	}
}

func (p *personalCourseSrv) DelCourse(product.PersonalCourseSrv_DelCourseServer) error {
	panic("implement me")
}

func (p *personalCourseSrv) SetCourseInvalid(product.PersonalCourseSrv_SetCourseInvalidServer) error {
	panic("implement me")
}

func (p *personalCourseSrv) SetCourseValid(product.PersonalCourseSrv_SetCourseValidServer) error {
	panic("implement me")
}
