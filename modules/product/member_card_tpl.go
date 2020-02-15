package product

import (
	"github.com/joselee214/j7f/components/errors"
	"j7go/components"
	"j7go/errors"
	"j7go/proto/product"
	"go.uber.org/zap"
	productService "j7go/services/product"
	"j7go/utils"
)

type MemberCardTplServer struct{}

//新增会员卡
func (s *MemberCardTplServer) AddCard(srv product.MemberCardTplServer_AddCardServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("add_receive", err.Error()))
			return err
		}

		err = productService.AddMemberCardTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.Any("add_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_INSERT_CARD_TPL_ERROR)
		}

		res := &product.MemberCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("add_send", err.Error()))
			return err
		}

	}
}

//获取会员卡详情
func (s *MemberCardTplServer) GetCardInfo(srv product.MemberCardTplServer_GetCardInfoServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("get_receive", err.Error()))
			return err
		}

		res, err := productService.GetMemberCardTplInfo(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.Any("get_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_GET_CARD_TPL_INFO_ERROR)
		}

		res.Status = errors.GetResHeader(err)

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("get_send", err.Error()))
			return err
		}

	}
}

//编辑会员卡
func (s *MemberCardTplServer) EditCard(srv product.MemberCardTplServer_EditCardServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("edit_receive", err.Error()))
			return err
		}

		err = productService.EditMemberCardTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.Any("edit_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_UPDATE_CARD_TPL_ERROR)
		}

		res := &product.MemberCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("edit_send", err.Error()))
			return err
		}

	}
}

//删除会员卡
func (s *MemberCardTplServer) DelCard(srv product.MemberCardTplServer_DelCardServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("del_receive", err.Error()))
			return err
		}

		err = productService.DelMemberCardTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.Any("del_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_DELETE_CARD_TPL_ERROR)
		}

		res := &product.MemberCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("del_send", err.Error()))
			return err
		}

	}
}

//停售会员卡
func (s *MemberCardTplServer) StopSaleCard(srv product.MemberCardTplServer_StopSaleCardServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("stop_sale_receive", err.Error()))
			return err
		}

		err = productService.StopSaleMemberCardTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("stop_sale_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_STOP_SALE_CARD_ERROR)
		}

		res := &product.MemberCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("stop_sale_send", err.Error()))
			return err
		}

	}

}

//恢复会员卡售卖
func (s *MemberCardTplServer) RecoverSaleCard(srv product.MemberCardTplServer_RecoverSaleCardServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("recover_sale_receive", err.Error()))
			return err
		}

		err = productService.RecoverSaleMemberCardTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("recover_sale_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_RECOVER_SALE_CARD_ERROR)
		}

		res := &product.MemberCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("recover_sale_send", err.Error()))
			return err
		}
	}

}

//会员卡上架
func (s *MemberCardTplServer) UpShelfCard(srv product.MemberCardTplServer_UpShelfCardServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("up_shelf_card", err.Error()))
			return err
		}

		err = productService.UpShelfCardTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("up_shelf_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_RECOVER_SALE_DEPOSIT_CARD_ERROR)
		}

		res := &product.MemberCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("up_shelf_card", err.Error()))
			return err
		}
	}
}

//会员卡下架
func (s *MemberCardTplServer) DownShelfCard(srv product.MemberCardTplServer_DownShelfCardServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("down_shelf_card", err.Error()))
			return err
		}

		err = productService.DownShelfCardTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("down_shelf_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_DOWN_SHELF_CARD_ERROR)
		}

		res := &product.MemberCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("down_shelf_card", err.Error()))
			return err
		}
	}
}

//上级会员卡详情
func (s *MemberCardTplServer) GetShelfCardInfo(srv product.MemberCardTplServer_GetShelfCardInfoServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("get_shelf_card", err.Error()))
			return err
		}

		res, err := productService.GetShelfCardInfo(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("get_shelf_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_GET_SHELF_CARD_INFO_ERROR)
		}

		res.Status = errors.GetResHeader(err)

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("get_shelf_card", err.Error()))
			return err
		}
	}
}

//重新上架
func (s *MemberCardTplServer) ReUpShelfCard(srv product.MemberCardTplServer_ReUpShelfCardServer) error {
	for {
		params, err := srv.Recv()
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("re_up_shelf_card", err.Error()))
			return err
		}

		err = productService.ReUpShelfCardTpl(srv.Context(), params)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("re_up_shelf_card", err.Error()))
			err = errors.NewFromCode(business_errors.ProductError_UP_SHELF_CARD_ERROR)
		}

		res := &product.MemberCardTplResponse{
			Status: errors.GetResHeader(err),
		}

		err = srv.Send(res)
		if err != nil {
			utils.GetTraceLog(srv.Context()).Error("member_card_tpl", zap.String("down_shelf_card", err.Error()))
			return err
		}
	}
}
