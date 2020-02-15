package product

import (
	"go.7yes.com/j7f/components/errors"
	"j7go/errors"
	"j7go/proto/product"
	"go.uber.org/zap"
	productService "j7go/services/product"
	"j7go/utils"
)

type DepositCardTplServer struct {}

func ( d *DepositCardTplServer ) AddDepositCard(srv product.DepositCardTplServer_AddDepositCardServer) error {
	for {
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("receive", err.Error()))
			return err
		}

		cardId,err := productService.AddDepositCardTpl(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.Any("add_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_SAVE_DEPOSIT_CARD_TPL_ERROR)
		}

		response := &product.AddAndUpdateDepositCardTplResponse{
			Status: errors.GetResHeader(err),
			Id:uint32(cardId),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("add_card", err.Error()))
			return err
		}

	}
}

func (d *DepositCardTplServer) UpdateDepositCard(srv product.DepositCardTplServer_UpdateDepositCardServer) error {
	for{
		request, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("receive", err.Error()))
			return err
		}
		err = productService.UpdateDepositCardTpl(srv.Context(), request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.Any("edit_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_SAVE_DEPOSIT_CARD_TPL_ERROR)
		}

		response := &product.AddAndUpdateDepositCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(response)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("add_card", err.Error()))
			return err
		}
	}
}

func (d *DepositCardTplServer) GetDepositCardDetail(srv product.DepositCardTplServer_GetDepositCardDetailServer) error {
	for {
		request,err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.String("receive",err.Error()))
			return err
		}
		res,err := productService.GetDepositDetailById(srv.Context(),request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.Any("get_deposit_detail",err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_GET_DEPOSIT_TPL_INFO_ERROR)
		}
		res.Status = errors.GetResHeader(err)

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("add_card", err.Error()))
			return err
		}
	}
}

func (d *DepositCardTplServer) StopSellDepositCard(srv product.DepositCardTplServer_StopSellDepositCardServer) error {
	for {
		request,err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.String("receive",err.Error()))
			return err
		}
		err = productService.StopSellDepositCard(srv.Context(),request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.Any("stop_sell_deposit_card",err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_STOP_SALE_DEPOSIT_CARD_ERROR)
		}

		err = srv.Send(&product.StopSellDepositCardResponse{
			Status: errors.GetResHeader(err),
		})
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("stop_sell_deposit_card", err.Error()))
			return err
		}
	}
}

func (d *DepositCardTplServer) RestartSellDepositCard(srv product.DepositCardTplServer_RestartSellDepositCardServer) error {
	for {
		request,err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.String("receive",err.Error()))
			return err
		}
		err = productService.RestartSellDepositCard(srv.Context(),request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.Any("restart_sell_deposit_card",err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_STOP_SALE_DEPOSIT_CARD_ERROR)
		}

		err = srv.Send(&product.RestartSellDepositCardResponse{
			Status: errors.GetResHeader(err),
		})
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("restart_sell_deposit_card", err.Error()))
			return err
		}
	}
}

func (d *DepositCardTplServer) OnlineSellDepositCard(srv product.DepositCardTplServer_OnlineSellDepositCardServer) error {
	for {
		request,err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.String("receive",err.Error()))
			return err
		}
		err = productService.OnlineSellDepositCard(srv.Context(),request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.Any("restart_sell_deposit_card",err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_STOP_SALE_DEPOSIT_CARD_ERROR)
		}

		err = srv.Send(&product.OnlineSellDepositCardResponse{
			Status: errors.GetResHeader(err),
		})
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("restart_sell_deposit_card", err.Error()))
			return err
		}
	}
}

func (d *DepositCardTplServer) OfflineSellDepositCard(srv product.DepositCardTplServer_OfflineSellDepositCardServer) error {
	for {
		request,err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.String("receive",err.Error()))
			return err
		}
		err = productService.OfflineSellDepositCard(srv.Context(),request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.Any("restart_sell_deposit_card",err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_STOP_SALE_DEPOSIT_CARD_ERROR)
		}

		err = srv.Send(&product.OfflineSellDepositCardResponse{
			Status: errors.GetResHeader(err),
		})
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("restart_sell_deposit_card", err.Error()))
			return err
		}
	}
}

func (d *DepositCardTplServer) DeleteSellDepositCard(srv product.DepositCardTplServer_DeleteSellDepositCardServer) error {
	for {
		request,err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.String("receive",err.Error()))
			return err
		}
		err = productService.DeleteSellDepositCard(srv.Context(),request)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl",zap.Any("restart_sell_deposit_card",err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_STOP_SALE_DEPOSIT_CARD_ERROR)
		}

		err = srv.Send(&product.DeleteDepositCardResponse{
			Status: errors.GetResHeader(err),
		})
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("deposit_card_tpl", zap.String("restart_sell_deposit_card", err.Error()))
			return err
		}
	}
}




