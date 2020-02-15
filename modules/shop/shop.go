package shop

import (
	"go.7yes.com/j7f/components/errors"
	"go.7yes.com/j7f/proto/common"
	"j7go/proto/shop"
	"go.uber.org/zap"
	"j7go/services/shop"
	"j7go/utils"
)

type ShopService struct{}

//获取门店信息
func (s *ShopService) GetShopInfo(server shop.ShopServer_GetShopInfoServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		shopInfo, err := shopService.GetShopInfo(server.Context(), uint(request.ShopId))
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("get_shop_info", zap.String("result_error", err.Error()))
		}

		res := &shop.GetShopInfoResponse{
			Status: errors.GetResHeader(err),
		}

		if err == nil {
			res.Id = uint32(shopInfo.ID)
			res.BrandId = uint32(shopInfo.BrandID)
			res.ShopName = shopInfo.ShopName
			res.Address = shopInfo.Address
			res.ShopStatus = uint32(shopInfo.ShopStatus)
			res.Description = shopInfo.Description
			res.IsAllday = common.IsAvailable(shopInfo.IsAllday)
			res.ProvinceId = uint32(shopInfo.ProvinceID)
			res.CityId = uint32(shopInfo.CityID)
			res.DistrictId = uint32(shopInfo.DistrictID)
			res.Lat = shopInfo.Lat
			res.Lng = shopInfo.Lng
			res.IsDel = common.DelStatus(shopInfo.IsDel)
			res.CreatedTime = uint32(shopInfo.CreatedTime)
			res.UpdatedTime = uint32(shopInfo.UpdatedTime)
			res.AlbumId = uint32(shopInfo.AlbumID)
			res.ProvinceName = shopInfo.ProvinceName
			res.CityName = shopInfo.CityName
			res.DistrictName = shopInfo.DistrictName
		}
		err = server.Send(res)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//获取门店联系方式
func (s *ShopService) GetShopContacts(server shop.ShopServer_GetShopContactsServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		shopContactWaysModel, err := shopService.GetShopContactWays(server.Context(), uint(request.ShopId))

		res := &shop.GetShopContactsResponse{
			Status: errors.GetResHeader(err),
		}

		shopContactWays := make([]*shop.ShopContact, len(shopContactWaysModel))
		if err == nil {
			for k, v := range shopContactWaysModel {
				var contactWay shop.ShopContact

				contactWay.Id = uint32(v.ID)
				contactWay.ShopId = uint32(v.ShopID)
				contactWay.ContactDetail = v.ContactDetail
				contactWay.ContactType = uint32(v.ContactType)
				contactWay.IsFavorite = common.IsAvailable(v.IsFavorite)
				contactWay.IsDel = common.DelStatus(v.IsDel)
				contactWay.CreatedTime = uint32(v.CreatedTime)
				contactWay.UpdatedTime = uint32(v.UpdatedTime)

				shopContactWays[k] = &contactWay
			}
			utils.GetTraceLog(server.Context()).Debug("test", zap.Any("contactWay", shopContactWays))
			res.ShopContacts = shopContactWays
		}

		err = server.Send(res)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//获取门店营业时间信息
func (s *ShopService) GetShopBusinessTime(server shop.ShopServer_GetShopBusinessTimeServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		//调用service获取营业时间
		shopBusinessTimeModel, err := shopService.GetShopBusinessTime(server.Context(), uint(request.ShopId))

		res := &shop.GetShopBusinessTimeResponse{
			Status: errors.GetResHeader(err),
		}

		if err == nil {
			var shopBusinessTime []*shop.ShopBusinessTime

			for _, v := range shopBusinessTimeModel {
				var timeInfo shop.ShopBusinessTime

				timeInfo.Id = uint32(v.ID)
				timeInfo.ShopId = uint32(v.ShopID)
				timeInfo.WeekDay = uint32(v.WeekDay)
				timeInfo.StartTime = v.StartTime
				timeInfo.EndTime = v.EndTime
				timeInfo.IsDel = common.DelStatus(v.IsDel)
				timeInfo.CreatedTime = uint32(v.CreatedTime)
				timeInfo.UpdatedTime = uint32(v.UpdatedTime)

				shopBusinessTime = append(shopBusinessTime, &timeInfo)
			}
			res.ShopBusinessTime = shopBusinessTime
		}

		err = server.Send(res)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//获取门店服务列表
func (s *ShopService) GetShopServices(server shop.ShopServer_GetShopServicesServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		//获取门店服务列表
		shopServicesModel, err := shopService.GetShopServices(server.Context(), uint(request.ShopId))

		res := &shop.GetShopServicesResponse{
			Status: errors.GetResHeader(err),
		}

		if err == nil {
			var shopServices []*shop.ShopService
			for _, v := range shopServicesModel {
				var shopServiceInfo shop.ShopService
				shopServiceInfo.Id = uint32(v.ID)
				shopServiceInfo.ServiceName = v.ServiceName
				shopServiceInfo.ImgUrl = v.ImgURL
				shopServiceInfo.IsDel = common.DelStatus(v.IsDel)
				shopServiceInfo.CreatedTime = uint32(v.CreatedTime)
				shopServiceInfo.UpdatedTime = uint32(v.UpdatedTime)
				shopServices = append(shopServices, &shopServiceInfo)
			}
			res.ShopServices = shopServices
		}
		utils.GetTraceLog(server.Context()).Debug("shop_services", zap.Any("return", res))
		err = server.Send(res)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//更新门店聚合接口
func (s *ShopService) UpdateShop(server shop.ShopServer_UpdateShopServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		//更新门店信息
		err = shopService.UpdateShop(server.Context(), request)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("update_shop", zap.String("result_error", err.Error()))
		}

		updateResponse := &shop.CommonResponse{
			Status: errors.GetResHeader(err),
		}
		err = server.Send(updateResponse)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//插入门店聚合接口
func (s *ShopService) InsertShop(server shop.ShopServer_InsertShopServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		//新增门店信息
		err = shopService.InsertShop(server.Context(), request)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("insert_shop", zap.String("result_error", err.Error()))
		}

		updateResponse := &shop.CommonResponse{
			Status: errors.GetResHeader(err),
		}
		err = server.Send(updateResponse)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}
