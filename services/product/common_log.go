package product

import (
	"context"
	"j7go/models/tests/product"
	"time"
)

//商品操作日志表
type PLog struct {
	TraceId string
	ShopId uint
	BrandId uint
	ProductId uint
	ProductType int8
	OperatorId uint
	OperateTable string
	OperateType int8
	ContentOld string
	ContentNew string
	Reason string
}

//LogInfo
func (P *PLog) LogInfo(ctx context.Context,logData *PLog) error {
	pL := &product.ProductOperateLog{}
	pL.TraceID = logData.TraceId
	pL.BrandID = logData.BrandId
	pL.ShopID = logData.ShopId
	pL.ProductType = logData.ProductType
	pL.ProductID = logData.ProductId
	pL.OperateTable = logData.OperateTable
	pL.OperateType = logData.OperateType
	pL.OperatorID = logData.OperatorId
	pL.ContentOld = logData.ContentOld
	pL.ContentNew = logData.ContentNew
	pL.Reason = logData.Reason
	pL.CreatedTime = uint(time.Now().Unix())
	pL.UpdatedTime = uint(time.Now().Unix())
	err := pL.Insert(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (P *PLog)LogInfoBatch(ctx context.Context,pLogs []*PLog) error  {
	list := make([]*product.ProductOperateLog,len(pLogs))
	for index,log := range pLogs{
		pL := &product.ProductOperateLog{}
		pL.TraceID = log.TraceId
		pL.BrandID = log.BrandId
		pL.ShopID = log.ShopId
		pL.ProductType = log.ProductType
		pL.ProductID = log.ProductId
		pL.OperateTable = log.OperateTable
		pL.OperateType = log.OperateType
		pL.OperatorID = log.OperatorId
		pL.ContentOld = log.ContentOld
		pL.ContentNew = log.ContentNew
		pL.Reason = log.Reason
		pL.CreatedTime = uint(time.Now().Unix())
		pL.UpdatedTime = uint(time.Now().Unix())

		list[index] = pL
	}
	err := product.OperateLogBatchInsert(ctx,list)
	if err != nil {
		return err
	}
	return nil
}