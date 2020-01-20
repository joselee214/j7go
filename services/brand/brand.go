package brandService

import (
	"context"
	"go.7yes.com/go/proto/brand"
	"j7go/models/shop"
)

func GetShopList(ctx context.Context, brandId uint) (*brand.GetShopListResponse, error) {
	res := &brand.GetShopListResponse{}

	//获取所有门店
	shops, err := shopModel.ShopByBrandID(ctx, brandId)
	if err != nil {
		return res, err
	}

	provinceMap := make(map[uint]*brand.ProvinceList)
	cityMap := make(map[uint]*brand.CityList)
	districtMap := make(map[uint]*brand.DistrictList)

	for _, shop := range shops {

		if _, ok := provinceMap[shop.ProvinceID]; !ok {
			province := &brand.ProvinceList{}
			province.ProvinceId = uint32(shop.ProvinceID)
			province.ProvinceName = shop.ProvinceName
			provinceMap[shop.ProvinceID] = province
		}

		if _, ok := cityMap[shop.CityID]; !ok {
			city := &brand.CityList{}
			city.CityId = uint32(shop.CityID)
			city.CityName = shop.CityName
			cityMap[shop.CityID] = city
			provinceMap[shop.ProvinceID].CityList = append(provinceMap[shop.ProvinceID].CityList, cityMap[shop.CityID])
		}

		if _, ok := districtMap[shop.DistrictID]; !ok {
			district := &brand.DistrictList{}
			district.DistrictId = uint32(shop.DistrictID)
			district.DistrictName = shop.DistrictName
			districtMap[shop.DistrictID] = district
			cityMap[shop.CityID].DistrictList = append(cityMap[shop.CityID].DistrictList, districtMap[shop.DistrictID])
		}

		shopInfo := &brand.Shops{
			ShopId:   uint32(shop.ID),
			ShopName: shop.ShopName,
		}
		districtMap[shop.DistrictID].Shops = append(districtMap[shop.DistrictID].Shops, shopInfo)
	}

	shopList := &brand.ShopList{}
	for _, province := range provinceMap {
		shopList.ProvinceList = append(shopList.ProvinceList, province)
	}

	res.ShopList = shopList

	return res, nil
}
