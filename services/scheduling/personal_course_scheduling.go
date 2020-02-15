package scheduling

import (
	"context"
	"go.7yes.com/j7f/proto/scheduling"
	"time"
	"j7go/components"
	"j7go/models/product"
	"j7go/utils"
)

//新增私教课排期
func AddPersonalSchedule(ctx context.Context, request *scheduling.PersonalScheduleRequest) error {
	scheduleLists := make([]*product.PersonalCourseSchedule, len(request.PersonalSchedules))
	for index, schedule := range request.PersonalSchedules {
		s := &product.PersonalCourseSchedule{}
		s.BrandID = uint(request.BrandId)
		s.ShopID = uint(request.ShopId)
		s.CoachID = uint(request.CoachId)
		s.ScheduleDate = uint(schedule.ScheduleDate)
		s.StartTime = uint(schedule.StartTime)
		s.EndTime = uint(schedule.EndTime)

		scheduleLists[index] = s
	}
	err := product.PersonalSchedulingBatchInsert(ctx, scheduleLists)
	if err != nil {
		return err
	}
	return nil
}

//编辑私教课排期
func EditPersonalSchedule(ctx context.Context, request *scheduling.PersonalScheduleRequest) error {
	ctx, err := components.M.BeginTransaction(ctx, 1*time.Second)
	if err != nil {
		return err
	}

	for _, schedule := range request.PersonalSchedules {
		s := &product.PersonalCourseSchedule{}
		s.ID = uint(schedule.ScheduleId)
		s.ScheduleDate = uint(schedule.ScheduleDate)
		s.StartTime = uint(schedule.StartTime)
		s.EndTime = uint(schedule.EndTime)
		err := product.EditPersonalScheduling(ctx, s)
		if err != nil {
			_ = components.M.Rollback(ctx)
			return err
		}
	}

	return components.M.Commit(ctx)
}

//复制私教课排期
func CopyPersonalSchedule(ctx context.Context, request *scheduling.CopyPersonalScheduleRequest) error {
	//查询出所要复制的排期

	//根据给定的应用日期，将排期批量插入
	return nil
}

func AddPersonalReserve(ctx context.Context, request *scheduling.PersonalReserveRequest) (uint32, error) {
	return utils.IntZero, nil
}
func EditPersonalReserve(ctx context.Context, request *scheduling.EditPersonalReserveRequest) error {
	return nil
}

func ViewPersonalReserve(ctx context.Context, request *scheduling.ViewPersonalReserveRequest) (*scheduling.ViewPersonalReserveResponse, error) {
	return nil, nil
}

func CancelPersonalReserve(ctx context.Context, request *scheduling.CancelPersonalReserveRequest) error {
	return nil
}

func CheckinPersonalReserve(ctx context.Context, request *scheduling.CheckinPersonalReserveRequest) error {
	return nil
}
