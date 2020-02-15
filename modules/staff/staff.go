package staff

import (
	"github.com/joselee214/j7f/components/errors"
	"j7go/proto/staff"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	st "j7go/services/staff"
	"j7go/utils"
)

//初始化，商户员工服务
func Init(g *grpc.Server) {
	s := &StaffService{}
	staff.RegisterStaffServerServer(g, s)
}

type StaffService struct{}

//新增员工
func (s *StaffService) CreateStaffServices(srv staff.StaffServer_CreateStaffServicesServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		staffId, err := st.CreateStaffs(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("create_staff")
		}

		results := &staff.CreateStaffResponse{
			Status:errors.GetResHeader(err),
			StaffId: staffId,

		}
		err = srv.Send(results)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//更新员工信息
func (s *StaffService) UpdateStaffServices(srv staff.StaffServer_UpdateStaffServicesServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		err = st.UpdateStaffs(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("update_staff")
		}

		result := &staff.StaffResponse{
			Status:errors.GetResHeader(err),
		}
		err = srv.Send(result)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//更新员工详细信息
func (s *StaffService) UpdateStaffDetailedServices(srv staff.StaffServer_UpdateStaffDetailedServicesServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		err = st.UpdateStaffDetailed(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("update_staff_detailed")
		}

		result := &staff.StaffResponse{
			Status:errors.GetResHeader(err),
		}
		err = srv.Send(result)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
	return nil
}
//更新员工教练信息
func (s *StaffService) UpdateStaffCoachServices(srv staff.StaffServer_UpdateStaffCoachServicesServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		err = st.UpdateStaffCoach(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("update_staff_coach")
		}

		result := &staff.StaffResponse{
			Status:errors.GetResHeader(err),
		}
		err = srv.Send(result)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
	return nil
}
//获取商户详情
func (s *StaffService) GetStaffInfo(srv staff.StaffServer_GetStaffInfoServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		staffs, err := st.StaffInfo(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("get_staff_info")
		}
		err = srv.Send(staffs)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}

	}
}

//删除员工
//TODO 删除员工有进行到事物，不能删除 必须县close再进行删除操作 （相关联到商品和其他业务 暂时开发部分）
func (s *StaffService) DeleteStaffServices(srv staff.StaffServer_DeleteStaffServicesServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
	   	err = st.DelStaff(srv.Context(), uint(request.StaffId))
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("delete_staff")
		}
	   	results := &staff.StaffResponse{
	   		Status:errors.GetResHeader(err),
		}

		err = srv.Send(results)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//TODO 解绑员工与场馆之前的关系 场馆单方面操作  判断当前员工与当前场馆是否有业务耦合 有的话则不能删除

//员工账户信息
func (s *StaffService) GetStaffAccountService(srv staff.StaffServer_GetStaffAccountServiceServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		staffAccount, err := st.GetStaffAccount(srv.Context(), request.StaffId)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("get_staff_account_info")
		}
		err = srv.Send(staffAccount)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//员工账户修改
func (s *StaffService) UpdateStaffAccountService(srv staff.StaffServer_UpdateStaffAccountServiceServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		err = st.UpdateStaffBankAccount(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("update_staff_bank_account_info")
		}
		results := &staff.StaffResponse{
			Status: errors.GetResHeader(err),
		}
		err = srv.Send(results)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//员工职位信息
func (s *StaffService) GetStaffPositionService(srv staff.StaffServer_GetStaffPositionServiceServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		staffPosition, err := st.GetStaffPosition(srv.Context(), request.StaffId)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("get_staff_position")
		}
		err = srv.Send(staffPosition)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}

}

//新增员工职位信息
func (s *StaffService) StoreStaffPositionService(srv staff.StaffServer_StoreStaffPositionServiceServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		 err = st.StoreStaffPosition(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("store_staff_position")
		}
		results := &staff.StaffResponse{
			Status: errors.GetResHeader(err),
		}
		err = srv.Send(results)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}

}

