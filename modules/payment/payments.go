package payment

import (
	"go.7yes.com/j7f/components/errors"
	"j7go/proto/payments"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"j7go/services/pays"
	"j7go/utils"
)

func Init(g *grpc.Server) {
	s := &paymentService{}
	payments.RegisterPaymentsServerServer(g, s)

}

type paymentService struct{}

//支付方式列表
func (s *paymentService) GetPayments(srv payments.PaymentsServer_GetPaymentsServer) error {
	for {
		_, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		//processing service data
		pay, err := pays.GetPayments(srv.Context())
		utils.GetTraceLog(srv.Context()).Debug("pays", zap.Any("pays", pay))

		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("payments", zap.String("get_pays", err.Error()))
		}

		res := &payments.PaymentsResponse{
			Status:           errors.GetResHeader(err),
			PayMethodsSource: pay,
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}

	}

}

//修改支付方式
func (s *paymentService) GetEditPay(srv payments.PaymentsServer_GetEditPayServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		err = pays.EditPay(srv.Context(), uint(request.Id), uint(request.IsDel))

		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("services", zap.String("err", err.Error()))
		}

		res := &payments.EditPayResponse{
			Status: errors.GetResHeader(err),
		}
		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}
