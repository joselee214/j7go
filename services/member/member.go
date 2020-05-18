package member

import (
	"context"
	"j7go/components"
	memberModel "j7go/models/tests/member"
	"j7go/models/tests/region"
	"j7go/proto/member"
	"j7go/utils"
	"time"
)

//新增会员 返回会员id
func AddMember(ctx context.Context, memberRequest *member.MemberInfo) (uint32, error) {
	//新增会员信息
	memberInfo, err := formatAddMemberInfo(ctx, memberRequest)
	if err != nil {
		return utils.IntZero, err
	}

	ctx, err = components.M.BeginTransaction(ctx, 1*time.Second)

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	err = memberInfo.Insert(ctx)
	if err != nil {
		return utils.IntZero, err
	}

	//新增门店会员信息
	err = AddShopMember(ctx, uint(memberRequest.BrandId), uint(memberRequest.ShopId), memberInfo.ID)
	if err != nil {
		return utils.IntZero, err
	}

	err = components.M.Commit(ctx)
	if err != nil {
		return utils.IntZero, err
	}

	return uint32(memberInfo.ID), nil
}

//编辑会员信息，部分更新
//func EditMember(ctx context.Context, memberRequest *member.MemberInfo) error {
//	editMemberInfo, err := formatEditMemberInfo(ctx, memberRequest)
//	if err != nil {
//		return err
//	}
//}

func formatEditMemberInfo(ctx context.Context, memberRequest *member.MemberInfo) (map[string]interface{}, error) {
	updateCondition := make(map[string]interface{})
	updateCondition["sex"] = memberRequest.Sex
	updateCondition["birthday"] = memberRequest.Birthday
	updateCondition["id_card_type"] = memberRequest.IdCardType
	updateCondition["id_card"] = memberRequest.IdCard
	updateCondition["married_type"] = memberRequest.MarriedType
	updateCondition["has_children"] = memberRequest.HasChildren
	updateCondition["country_id"] = memberRequest.CountryId
	updateCondition["nation_id"] = memberRequest.NationId
	updateCondition["education_level"] = memberRequest.EducationLevel
	updateCondition["jobs"] = memberRequest.Jobs
	updateCondition["income_level"] = memberRequest.IncomeLevel
	updateCondition["fitness_goal"] = memberRequest.FitnessGoal
	updateCondition["fitness_level"] = memberRequest.FitnessLevel
	updateCondition["email"] = memberRequest.Email
	updateCondition["province_id"] = memberRequest.ProvinceId
	updateCondition["city_id"] = memberRequest.CityId
	updateCondition["district_id"] = memberRequest.DistrictId
	updateCondition["living_address"] = memberRequest.LivingAddress

	//获取省市区信息
	provinceName, cityName, districtName, err := GetMemberRegionInfo(ctx, memberRequest.DistrictId, memberRequest.CityId, memberRequest.ProvinceId)
	if err != nil {
		return nil, err
	}
	updateCondition["province_name"] = provinceName
	updateCondition["province_name"] = cityName
	updateCondition["district_name"] = districtName
	return updateCondition, nil
}

//格式化会员信息
func formatAddMemberInfo(ctx context.Context, memberRequest *member.MemberInfo) (*memberModel.Member, error) {
	memberInfo := &memberModel.Member{}
	memberInfo.BrandID = uint(memberRequest.BrandId)
	memberInfo.MemberName = memberRequest.MemberName
	memberInfo.CountryPrefix = uint(memberRequest.CountryPrefix)
	memberInfo.Mobile = memberRequest.Mobile
	memberInfo.RegisterType = int8(memberRequest.RegisterType)
	memberInfo.RegisterWay = uint(memberRequest.RegisterWay)
	memberInfo.Sex = int8(memberRequest.Sex)
	memberInfo.Weight = uint(memberRequest.Weight)
	memberInfo.Height = uint(memberRequest.Height)
	memberInfo.IDCardType = int8(memberRequest.IdCardType)
	memberInfo.IDCard = memberRequest.IdCard
	memberInfo.Birthday = memberRequest.Birthday
	memberInfo.FirstRegTime = uint(time.Now().Unix())
	memberInfo.Jobs = memberRequest.Jobs
	memberInfo.EducationLevel = int8(memberRequest.EducationLevel)
	memberInfo.MarriedType = int8(memberRequest.MarriedType)
	memberInfo.HasChildren = int8(memberRequest.HasChildren)
	memberInfo.MemberComment = memberRequest.MemberComment
	memberInfo.Email = memberRequest.Email
	memberInfo.NationID = uint(memberRequest.NationId)
	memberInfo.EmergencyContactName = memberRequest.EmergencyContactName
	memberInfo.EmergencyContactMobile = memberRequest.EmergencyContactMobile
	memberInfo.ProvinceID = uint(memberRequest.ProvinceId)
	memberInfo.CityID = uint(memberRequest.CityId)
	memberInfo.DistrictID = uint(memberRequest.DistrictId)
	memberInfo.LivingAddress = memberRequest.LivingAddress
	memberInfo.FitnessGoal = memberRequest.FitnessGoal
	memberInfo.FitnessLevel = int8(memberRequest.FitnessLevel)
	memberInfo.IncomeLevel = memberRequest.IncomeLevel
	memberInfo.IsDel = utils.NOT_DELETED
	memberInfo.CreatedTime = uint(time.Now().Unix())
	memberInfo.UpdatedTime = uint(time.Now().Unix())

	//获取省市区信息
	provinceName, cityName, districtName, err := GetMemberRegionInfo(ctx, memberRequest.DistrictId, memberRequest.CityId, memberRequest.ProvinceId)
	if err != nil {
		return nil, err
	}
	memberInfo.ProvinceName = provinceName
	memberInfo.CityName = cityName
	memberInfo.DistrictName = districtName

	return memberInfo, nil
}

//获取会员的地区信息
func GetMemberRegionInfo(ctx context.Context, districtId uint32, cityId uint32, provinceId uint32) (provinceName string, cityName string, districtName string, err error) {
	condition := make(map[string]interface{})
	if utils.IntZero != districtId {
		condition["district_id"] = districtId
	}
	if utils.IntZero != cityId {
		condition["city_id"] = cityId
	}
	if utils.IntZero != provinceId {
		condition["province_id"] = provinceId
	}
	//条件为空
	if len(condition) == utils.IntZero {
		return utils.StringInit, utils.StringInit, utils.StringInit, nil
	}
	//获取地区信息
	regionInfo, err := region.GetRegionInfoByCondition(ctx, condition)
	if err != nil {
		return utils.StringInit, utils.StringInit, utils.StringInit, err
	}
	if regionInfo == nil {
		provinceName = utils.StringInit
		cityName = utils.StringInit
		districtName = utils.StringInit
	} else {
		provinceName = regionInfo.ProvinceName
		cityName = regionInfo.CityName
		districtName = regionInfo.DistrictName
	}
	return provinceName, cityName, districtName, nil
}

//新增门店会员
func AddShopMember(ctx context.Context, brandId uint, shopId uint, memberId uint) (err error) {
	shopMemberInfo := &memberModel.ShopMember{}
	shopMemberInfo.ShopID = shopId
	shopMemberInfo.BrandID = brandId
	shopMemberInfo.MemberID = memberId
	shopMemberInfo.MemberStatus = memberModel.MEMBER_STATUS_NORMAL
	shopMemberInfo.MemberLevel = memberModel.MEMBER_LEVEL_LATENT
	shopMemberInfo.FollowCoachID = utils.IntZero
	shopMemberInfo.FollowSalesmanID = utils.IntZero
	shopMemberInfo.BeMemberTime = utils.IntZero
	shopMemberInfo.IsDel = utils.NOT_DELETED
	shopMemberInfo.CreatedTime = uint(time.Now().Unix())
	shopMemberInfo.UpdatedTime = uint(time.Now().Unix())
	err = shopMemberInfo.Insert(ctx)
	if err != nil {
		return err
	}
	return nil
}
