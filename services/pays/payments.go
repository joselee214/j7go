package pays

import (
	"context"
	"github.com/joselee214/j7f/components/errors"
	"github.com/joselee214/j7f/proto/common"
	"go.uber.org/zap"
	"j7go/errors"
	"j7go/models/tests/payment"
	ps "j7go/proto/payments"
	"j7go/utils"
	"time"
)

//获取支付列表
func GetPayments(ctx context.Context) ([]*ps.PaymentResult, error) {

	payments, err := payment.PaymentMethods(ctx)
	pays := make([]*ps.PaymentResult, len(payments))
	for index, value := range payments {
		results := &ps.PaymentResult{}
		results.Id = uint32(value.ID)
		results.PaymentName = value.PaymentName
		results.IsDel = common.DelStatus(value.IsDel)
		results.CreatingMode = ps.WayCreation(value.CreatingMode)
		results.IsOnlinePay = ps.OnlinePayment(value.IsOnlinePay)
		results.CreatedTime = uint32(value.CreatedTime)
		results.UpdatedTime = uint32(value.UpdatedTime)
		pays[index] = results
	}
	if err != nil {
		utils.GetTraceLog(ctx).Warn("payment", zap.Error(err))
		return nil, errors.NewFromCode(business_errors.PaymentsError_PAYMENT_METHODS_NOT_FOUND)
	}

	return pays, nil
}

//修改支付方式
func EditPay(ctx context.Context, id, isDel uint) error {
	payments := &payment.PaymentMethod{}
	payments.ID = int(id)
	payments.IsDel = int8(isDel)
	payments.UpdatedTime = int(time.Now().Unix())
	payments.Update(ctx)
	return nil
}
