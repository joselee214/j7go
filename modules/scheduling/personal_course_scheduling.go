package scheduling

import (
	"go.7yes.com/go/components/errors"
	"go.7yes.com/go/errors"
	"go.7yes.com/go/proto/scheduling"
	"go.uber.org/zap"
	schedulingService "j7go/services/scheduling"
	"j7go/utils"
)

type PersonalCourseScheduling struct{}

func (PersonalCourseScheduling) AddPersonalSchedule(srv scheduling.PersonalCourseScheduleServer_AddPersonalScheduleServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("receive", err.Error()))
			return err
		}

		err = schedulingService.AddPersonalSchedule(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.Any("add_scheduling", err.Error()))
			err = errors.NewFromCode(business_errors.SchedulingError_ADD_PERSONAL_SCHEDULE_ERROR)
		}

		response := &scheduling.CommonScheduleResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("add_scheduling", err.Error()))
			return err
		}

	}
}

func (PersonalCourseScheduling) EditPersonalSchedule(srv scheduling.PersonalCourseScheduleServer_EditPersonalScheduleServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("receive", err.Error()))
			return err
		}

		err = schedulingService.EditPersonalSchedule(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.Any("edit_scheduling", err.Error()))
			err = errors.NewFromCode(business_errors.SchedulingError_EDIT_PERSONAL_SCHEDULE_ERROR)
		}

		response := &scheduling.CommonScheduleResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("edit_scheduling", err.Error()))
			return err
		}

	}
}

func (PersonalCourseScheduling) CopyPersonalSchedule(srv scheduling.PersonalCourseScheduleServer_CopyPersonalScheduleServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("receive", err.Error()))
			return err
		}

		err = schedulingService.CopyPersonalSchedule(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.Any("copy_scheduling", err.Error()))
			err = errors.NewFromCode(business_errors.SchedulingError_COPY_PERSONAL_SCHEDULE_ERROR)
		}

		response := &scheduling.CommonScheduleResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("copy_scheduling", err.Error()))
			return err
		}

	}
}

func (PersonalCourseScheduling) AddPersonalReserve(srv scheduling.PersonalCourseScheduleServer_AddPersonalReserveServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("receive", err.Error()))
			return err
		}

		cardId, err := schedulingService.AddPersonalReserve(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.Any("add_reserve", err.Error()))
			err = errors.NewFromCode(business_errors.SchedulingError_ADD_PERSONAL_SCHEDULE_ERROR)
		}

		response := &scheduling.AddPersonalReserveResponse{
			Status: errors.GetResHeader(err),
			Id:     uint32(cardId),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("add_reserve", err.Error()))
			return err
		}

	}
}

func (PersonalCourseScheduling) EditPersonalReserve(srv scheduling.PersonalCourseScheduleServer_EditPersonalReserveServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("receive", err.Error()))
			return err
		}

		err = schedulingService.EditPersonalReserve(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.Any("add_scheduling", err.Error()))
			err = errors.NewFromCode(business_errors.SchedulingError_EDIT_PERSONAL_RESERVE_ERROR)
		}

		response := &scheduling.CommonScheduleResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("add_scheduling", err.Error()))
			return err
		}

	}
}

func (PersonalCourseScheduling) ViewPersonalReserve(srv scheduling.PersonalCourseScheduleServer_ViewPersonalReserveServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("receive", err.Error()))
			return err
		}

		viewReserve, err := schedulingService.ViewPersonalReserve(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.Any("view_scheduling", err.Error()))
			err = errors.NewFromCode(business_errors.SchedulingError_ADD_PERSONAL_SCHEDULE_ERROR)
		}

		err = srv.Send(viewReserve)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("view_scheduling", err.Error()))
			return err
		}

	}
}

func (PersonalCourseScheduling) CancelPersonalReserve(srv scheduling.PersonalCourseScheduleServer_CancelPersonalReserveServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("receive", err.Error()))
			return err
		}

		err := schedulingService.CancelPersonalReserve(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.Any("cancel_scheduling", err.Error()))
			err = errors.NewFromCode(business_errors.SchedulingError_ADD_PERSONAL_SCHEDULE_ERROR)
		}

		response := &scheduling.CommonScheduleResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("cancel_scheduling", err.Error()))
			return err
		}

	}
}

func (PersonalCourseScheduling) CheckinPersonalReserve(srv scheduling.PersonalCourseScheduleServer_CheckinPersonalReserveServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("receive", err.Error()))
			return err
		}

		cardId, err := schedulingService.CheckinPersonalReserve(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.Any("add_scheduling", err.Error()))
			err = errors.NewFromCode(business_errors.SchedulingError_ADD_PERSONAL_SCHEDULE_ERROR)
		}

		response := &scheduling.AddPersonalScheduleResponse{
			Status: errors.GetResHeader(err),
			Id:     uint32(cardId),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("personal_scheduling", zap.String("add_scheduling", err.Error()))
			return err
		}

	}
}
