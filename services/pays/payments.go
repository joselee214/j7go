package pays

import (
	"context"
	"go.7yes.com/go/components/errors"
	"go.7yes.com/go/errors"
	"go.7yes.com/go/proto/common"
	ps "go.7yes.com/go/proto/payments"
	"go.uber.org/zap"
	"time"
	"j7go/models/payment"
	"j7go/utils"
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
