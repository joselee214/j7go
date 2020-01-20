package staff

import (
	"context"
	"go.7yes.com/go/components/errors"
	"go.7yes.com/go/errors"
	"go.7yes.com/go/proto/staff"
	"go.uber.org/zap"
	"time"
	"j7go/components"
	countryModel "j7go/models/country"
	imageModel "j7go/models/images"
	"j7go/models/shop"
	staffModel "j7go/models/staff"
	"j7go/utils"
)

//新增员工返回id
func CreateStaffs(ctx context.Context, request *staff.StaffRequest) (uint32, error) {
	//开启事物
	ctx, err := components.M.BeginTransaction(ctx, 1 * time.Second)
	if err != nil {
		return utils.IntZero, errors.NewFromError(err)
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	//新增员工
	staffs, err := CreateStaff(ctx, request)
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff", zap.String("method", "create_staff"), zap.Any("staff", staffs), zap.Error(err))
		return utils.IntZero, errors.NewFromCode(business_errors.StaffError_CREATE_NEW_STAFF_ERROR)
	}

	//新增员工账户
	err = CreateStaffAccount(ctx, request, staffs.ID)
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff_account", zap.String("method", "create_staff_account"), zap.Error(err))
		return utils.IntZero, errors.NewFromCode(business_errors.StaffError_CREATE_NEW_STAFF_ERROR)
	}

	//绑定员工身份
	err = BindStaffIdentity(ctx, uint32(staffs.ID), request)
	if err != nil {
		return utils.IntZero, err
	}

	//绑定门店
	err = BindStaffShop(ctx, uint32(staffs.ID), request)
	if err != nil {
		return utils.IntZero, err
	}


	_ = components.M.Commit(ctx)
	return uint32(staffs.ID), nil

}

func CreateStaff(ctx context.Context, request *staff.StaffRequest) (*staffModel.Staff, error) {
	st := &staffModel.Staff{}
	st.BrandID = uint(request.BrandId)
	st.StaffName = request.StaffName
	st.Nickname = request.Nickname
	st.Mobile = string(request.Mobile)
	st.Mail = request.Email
	st.Sex = int8(request.Sex)
	st.IDType = int8(request.IdType)
	st.IDNumber = request.IdNum
	st.StaffNum = request.StaffNum
	st.CoachLevelID = uint(request.CoachLevelId)
	st.WorkingPost = request.WorkingPost
	st.NatureWork = int8(request.NatureWork)
	st.EntryDate = uint(request.EntryDate)
	st.IsPermission = int8(request.IsPermission)
	st.CountryCodeID = uint(request.CountryCodeId)
	st.CreatedTime = uint(time.Now().Unix())
	st.UpdatedTime = uint(time.Now().Unix())
	err := st.Insert(ctx)
	return st, err
}
//新增员工账户
func CreateStaffAccount(ctx context.Context, request *staff.StaffRequest, staffId uint) error {
	staffAccount := &staffModel.StaffAccount{}
	staffAccount.AccountType = staffModel.USER_ACCOUNT
	staffAccount.StaffID = staffId
	staffAccount.AccountName = request.Account
	staffAccount.AccountPwd = request.Password
	staffAccount.CreatedTime = uint(time.Now().Unix())
	staffAccount.UpdatedTime = uint(time.Now().Unix())
	err := staffAccount.Insert(ctx)
	return err
}

//更新员工信息
func UpdateStaffs(ctx context.Context, request *staff.StaffRequest) error {
	//开启事物
	ctx, err := components.M.BeginTransaction(ctx, 2 * time.Second)
	if err != nil {
		return errors.NewFromError(err)
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	err = UpdateStaff(ctx, request)
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff", zap.String("method", "update_staff"), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_CREATE_NEW_STAFF_ERROR)
	}

	//更新员工身份
	err = UpdateStaffIdentity(ctx, uint(request.BrandId), uint(request.StaffId), request)
	if err != nil {
		return err
	}

	//更新员工所在门店
	err = UpdateStaffBindShop(ctx, uint(request.StaffId), request)
	if err != nil {
		return err
	}

	_ = components.M.Commit(ctx)
	return nil
}

//更新员工基础信息
func UpdateStaff(ctx context.Context, request *staff.StaffRequest) error {
	st := &staffModel.Staff{}
	st.ID = uint(request.StaffId)
	st.BrandID = uint(request.BrandId)
	st.StaffName = request.StaffName
	st.Nickname = request.Nickname
	st.Mobile = string(request.Mobile)
	st.StaffNum = request.StaffNum
	st.Mail = request.Email
	st.Sex = int8(request.Sex)
	st.IDType = int8(request.IdType)
	st.IDNumber = request.IdNum
	st.WorkingPost = request.WorkingPost
	st.CoachLevelID = uint(request.CoachLevelId)
	st.NatureWork = int8(request.NatureWork)
	st.EntryDate = uint(request.EntryDate)
	st.CountryCodeID = uint(request.CountryCodeId)
	st.UpdatedTime = uint(time.Now().Unix())
	err := st.StaffUpdate(ctx, uint(request.StaffId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff", zap.String("method", "update_staff"), zap.Any("staff", st), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_UPDATE_STAFF_ERROR)
	}
	return nil
}

//更新员工详细信息
func UpdateStaffDetailed(ctx context.Context, request *staff.StaffRequest) error {
	//更新员工基础信息
	err := UpdateStaffDetailedInfo(ctx, request)
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff", zap.String("method", "update_staff_detailed"), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_UPDATE_STAFF_ERROR)
	}
	return nil
}

func UpdateStaffDetailedInfo(ctx context.Context, request *staff.StaffRequest) error {
	st := &staffModel.Staff{}
	st.ID = uint(request.StaffId)
	st.GraduatedSchool = request.GraduatedSchool
	st.GraduationTime = uint(request.GraduationTime)
	st.Education = int8(request.Education)
	st.Profession = request.Profession
	st.Birthday = request.Birthday
	st.NativePlace = request.NativePlace
	st.MarryStatus = int8(request.MarryStatus)
	st.ChildrenStatus = int8(request.ChildrenStatus)
	st.ProvinceID = uint(request.ProvinceId)
	st.ProvinceName = request.ProvinceName
	st.CityID = uint(request.CityId)
	st.CityName = request.CityName
	st.DistrictID = uint(request.DistrictId)
	st.DistrictName = request.DistrictName
	st.Address = request.Address
	st.Description = request.Description
	st.UpdatedTime = uint(time.Now().Unix())
	err := st.StaffDetailedUpdate(ctx, uint(request.StaffId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff", zap.String("method", "update_staff_detailed"), zap.Any("staff", st), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_UPDATE_STAFF_ERROR)
	}
	return nil
}

//更新员工教练信息
func UpdateStaffCoach(ctx context.Context, request *staff.StaffRequest) error {
	//开启事物
	ctx, err := components.M.BeginTransaction(ctx, 1 * time.Second)
	if err != nil {
		return errors.NewFromError(err)
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	//更新员工教练信息
	err = UpdateStaffCoachInfo(ctx, request)
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff", zap.String("method", "update_staff_detailed"), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_UPDATE_STAFF_ERROR)
	}
	//更新员工擅长项目
	err = UpdateStaffSpecialty(ctx, request)
	if err != nil {
		return err
	}

	//更新员工证书
	err = UpdateStaffMienCertification(ctx, request)
	if err != nil {
		return err
	}
	_ = components.M.Commit(ctx)
	return nil

}

func UpdateStaffCoachInfo(ctx context.Context, request *staff.StaffRequest) error {
	st := &staffModel.Staff{}
	st.EmploymentTime = uint(request.EmploymentTime)
	st.Introduction = request.Introduction
	st.AlbumID = uint(request.AlbumId)
	st.IsShow = int8(request.IsShow)
	st.UpdatedTime = uint(time.Now().Unix())
	err := st.StaffCoachUpdate(ctx, uint(request.StaffId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff", zap.String("method", "update_staff_coach"), zap.Any("staff", st), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_UPDATE_STAFF_ERROR)
	}
	return nil
}

//更新员工身份
func UpdateStaffIdentity(ctx context.Context, brandId, staffId uint, request *staff.StaffRequest) error {
	//处理员工以前身份
	err := DelStaffIdentity(ctx, brandId, staffId)
	if err != nil {
		return err
	}
	//重新绑定员工身份
	err = BindStaffIdentity(ctx, uint32(staffId), request)
	return err
}

//员工绑定身份
func BindStaffIdentity(ctx context.Context, staffId uint32, request *staff.StaffRequest) error {
	for _, identity := range request.Identity {
		staffIdentity := staffModel.StaffIdentityRelation{}
		staffIdentity.StaffID = uint(staffId)
		staffIdentity.BrandID = uint(request.BrandId)
		staffIdentity.Identity = int8(identity)
		staffIdentity.CreatedTime = uint(time.Now().Unix())
		staffIdentity.UpdatedTime = uint(time.Now().Unix())
		err := staffIdentity.Insert(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("staff", zap.String("method", "bind_identity_relation"), zap.Any("identity", staffIdentity), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_BIND_STAFF_IDENTITY_ERROR)
		}
	}
	return nil
}

//员工身份
func DelStaffIdentity(ctx context.Context, brandId, staffId uint) error {
	staffs, err := staffModel.GetStaffIdentityRelation(ctx, brandId, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_identity", zap.Uint("brand_id", brandId),
			zap.Uint("staff_id", staffId),
			zap.Any("get_staff_identity", staffs), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_GET_STAFF_IDENTITY_ERROR)
	}

	//删除已经存在的员工身份
	for _, staffIdentity := range staffs {
		err = staffIdentity.Delete(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_staff_identity", zap.Any("staff_identity", staffIdentity), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_DELETE_STAFF_INFO_ERROR)
		}
	}
	return nil
}

//更新员工绑定门店
func UpdateStaffBindShop(ctx context.Context, staffId uint, request *staff.StaffRequest) error {
	//删除员工绑定场馆
	err := DelStaffBindShop(ctx, staffId, uint(request.BrandId))
	if err != nil {
		return err
	}
	//员工重新绑定场馆
	err = BindStaffShop(ctx, uint32(staffId), request)
	return err
}

//员工绑定场馆
func BindStaffShop(ctx context.Context, staffId uint32, request *staff.StaffRequest) error {
	if request.CurrentShop == utils.IntZero {
		shopIds := request.ShopId
		for _, shopId := range shopIds {
			shops := staffModel.ShopStaffRelation{}
			shops.StaffID = uint(staffId)
			shops.ShopID = uint(shopId)
			shops.BrandID = uint(request.BrandId)
			shops.CreatedTime = uint(time.Now().Unix())
			shops.UpdatedTime = uint(time.Now().Unix())
			err := shops.Insert(ctx)
			if err != nil {
				utils.GetTraceLog(ctx).Error("bind_shop_staff", zap.String("method", "bind_shop_staff"), zap.Any("shops", shops), zap.Error(err))
				return errors.NewFromCode(business_errors.StaffError_BIND_STAFF_SHOP_ERROR)
			}
		}
	}
	return nil
}

//处理员工绑定场馆
func DelStaffBindShop(ctx context.Context, staffId, brandId uint) error {
	staffShops, err := staffModel.GetShopStaffs(ctx, brandId, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_shop_staff", zap.Uint("brand_id", brandId),
			zap.Uint("staff_id", staffId),
			zap.Any("get_shop_staff", staffShops), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_GET_STAFF_SHOP_ERROR)
	}
	//TODO 品牌维度判断当前员工在场馆下的关系 如果有业务事物则不可以删除 --未开发
	//删除已存在的场馆员工关系
	for _, staffs := range staffShops {
		err = staffs.Delete(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("delete_staff_shop", zap.Any("staffs", staffs), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_DELETE_STAFF_SHOP_ERROR)
		}
	}
	return nil
}

//更新员工擅长项目
func UpdateStaffSpecialty(ctx context.Context, request *staff.StaffRequest) error {
	//处理员工擅长项目
	err := DelStaffSpecialty(ctx, uint(request.BrandId), uint(request.StaffId))
	if err != nil {
		return err
	}
	//重新设置擅长项目
	err = BindStaffSpecialty(ctx, request.StaffId, request)

	return err
}

//员工擅长项目
func BindStaffSpecialty(ctx context.Context, staffId uint32, request *staff.StaffRequest) error {
	for _, sp := range request.SpecialtyId {
		specialty := staffModel.StaffSpecialtyRelation{}
		specialty.StaffID = uint(staffId)
		specialty.BrandID = uint(request.BrandId)
		specialty.SpecialtyID = uint(sp)
		specialty.CreatedTime = int(time.Now().Unix())
		specialty.UpdatedTime = int(time.Now().Unix())
		err := specialty.Insert(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("bind_staff_specialty", zap.String("method", "staff_specialty"), zap.Any("specialty", specialty), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_BIND_STAFF_SPECIALTY_ERROR)
		}
	}
	return nil
}

//删除员工擅长项目
func DelStaffSpecialty(ctx context.Context, brandId, staffId uint) error {
	staffSpecialty, err := staffModel.GetStaffSpecialtyRelation(ctx, brandId, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_specialty", zap.Uint("brand_id", brandId),
			zap.Uint("staff_id", staffId),
			zap.Any("get_staff_specialty", staffSpecialty), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_GET_STAFF_SPECIALTY_ERROR)
	}
	//删除已存在项目
	for _, staffs := range staffSpecialty {
		err = staffs.Delete(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("delete_staff_specialty", zap.Any("staffs", staffs), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_DELETE_STAFF_SPECIALTY_ERROR)
		}
	}
	return nil
}

//更新员工证书
func UpdateStaffMienCertification(ctx context.Context, request *staff.StaffRequest) error {
	//处理员工证书
	err := DelStaffMienCertification(ctx, uint(request.StaffId))
	if err != nil {
		return err
	}
	//重新设置证书
	err = StaffMienCertification(ctx, request.StaffId, request)
	return err
}

//员工添加专业证书
func StaffMienCertification(ctx context.Context, staffId uint32, request *staff.StaffRequest) error {
	for _, ce := range request.CertificationName {
		staffMien := staffModel.StaffMienCertification{}
		staffMien.StaffID = uint(staffId)
		staffMien.CertificationName = ce
		staffMien.CreatedTime = int(time.Now().Unix())
		staffMien.UpdatedTime = int(time.Now().Unix())
		err := staffMien.Insert(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("staff_mien_certification", zap.String("method", "staff_mien_certification"), zap.Any("staff_mien_certification", staffMien), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_BIND_STAFF_MIEN_CERTIFICATION_ERROR)
		}
	}
	return nil
}

//删除员工证书
func DelStaffMienCertification(ctx context.Context, staffId uint) error {
	mienCertification, err := staffModel.GetStaffMienCertification(ctx, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_mien_certification",
			zap.Uint("staff_id", staffId),
			zap.Any("get_staff_mien_certification", mienCertification), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_GET_STAFF_MIEN_CERTIFICATION_ERROR)
	}
	//删除已存在证书
	for _, mien := range mienCertification {
		err = mien.Delete(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("update_staff_mien_certification", zap.Any("mien_certification", mienCertification), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_DELETE_STAFF_MIEN_CERTIFICATION_ERROR)
		}
	}
	return nil
}

//员工详情
func StaffInfo(ctx context.Context, request *staff.StaffInfoRequest) (*staff.StaffInfoResponse, error) {
	staffs, err := staffModel.StaffByID(ctx, uint(request.StaffId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_info",
			zap.Uint32("staff_id", request.StaffId),
			zap.Any("staff_info", staffs), zap.Error(err))
		return nil, errors.NewFromCode(business_errors.StaffError_GET_STAFF_INFO_ERROR)
	}
	sr := &staff.StaffInfoResponse{}
	sr.Status = errors.GetResHeader(err)
	sr.StaffId = uint32(staffs.ID)
	sr.StaffName = staffs.StaffName
	sr.StaffNum = staffs.StaffNum
	sr.Nickname = staffs.Nickname
	sr.StaffSex = GetStaffSex(staffs.Sex)
	sr.IdType = GetIdType(staffs.IDType)
	sr.Mobile = staffs.Mobile
	coach, err := GetCoachLevel(ctx, staffs.CoachLevelID, staffs.BrandID)
	if err == nil {
		sr.CoachLevelId = coach
	}
	sr.IdNum = staffs.IDNumber
	ide, err := GetIdentity(ctx, staffs.BrandID, staffs.ID)
	if err == nil {
		sr.Identity = ide
	}
	sr.WorkingPost = staffs.WorkingPost
	shops, err := GetStaffShop(ctx, staffs.BrandID, uint(request.CurrentShop), staffs.ID)
	if err == nil {
		sr.ShopId = shops
	}
	sr.NatureWork = GetNatureWork(staffs.NatureWork)
	sr.EntryDate = int32(staffs.EntryDate)
	sr.IsPermission = GetPermission(staffs.IsPermission)
	sr.Birthday = staffs.Birthday
	sr.GraduatedSchool = staffs.GraduatedSchool
	sr.GraduationTime = int32(staffs.GraduationTime)
	sr.Education = GetEducation(staffs.Education)
	sr.Profession = staffs.Profession
	sr.NativePlace = staffs.NativePlace
	sr.MarryStatus = GetMarryStatus(staffs.MarryStatus)
	sr.ChildrenStatus = GetChildrenStatus(staffs.ChildrenStatus)
	sr.ProvinceId = uint32(staffs.ProvinceID)
	sr.ProvinceName = staffs.ProvinceName
	sr.CityId = uint32(staffs.CityID)
	sr.CityName = staffs.CityName
	sr.DistrictId = uint32(staffs.DistrictID)
	sr.DistrictName = staffs.DistrictName
	sr.Address = staffs.Address
	sr.Description = staffs.Description
	sr.EmploymentTime = int32(staffs.EmploymentTime)
	sp, err := GetSpecialty(ctx, staffs.BrandID, staffs.ID)
	if err == nil {
		sr.SpecialtyId = sp
	}
	ce, err := GetCertification(ctx, staffs.ID)
	if err == nil {
		sr.CertificationName = ce
	}
	sr.Introduction = staffs.Introduction
	sr.IsShow = GetIsShow(staffs.IsShow)

	images, err := GetImages(ctx, staffs.AlbumID)
	if err == nil {
		sr.ImagePersonal = images
	}
	country, err := GetCountryCode(ctx, staffs.CountryCodeID)
	if err == nil {
		sr.CountryCodeId = country
	}
	available, err := GetStaffAvatar(ctx, staffs.AlbumID)
	if err == nil {
		sr.ImageAvatar = available
	}
	faceImage, err := GetStaffFaceImages(ctx, staffs.AlbumID)
	if err == nil {
		sr.ImageFace = faceImage
	}
	sr.WorkStatus = GetWorkStatus(staffs.WorkStatus)
	sr.AlbumId = uint32(staffs.AlbumID)

	return sr, nil
}

//获取员工性别
func GetStaffSex(sex int8) *staff.StaffSex{
	s := &staff.StaffSex{}
	s.SexType = uint32(sex)
	switch sex {
	case staffModel.ENUM_ZERO:
		s.Name = staffModel.UNDEFINED
	case staffModel.ENUM_ONE:
		s.Name = staffModel.WOMAN
	case staffModel.ENUM_TWO:
		s.Name = staffModel.MAN
	}
	return s
}

//获取证件类别
func GetIdType(idType int8) *staff.IdTypeResult {
	i := &staff.IdTypeResult{}
	i.Type = uint32(idType)
	switch idType {
	case staffModel.ENUM_ONE:
		i.Name = staffModel.ID_CARD
	case staffModel.ENUM_TWO:
		i.Name = staffModel.PASSPORT
	}
	return i
}

//获取教练等级
func GetCoachLevel(ctx context.Context, id, brandId uint) (*staff.CoachLevel, error) {
	coachLevel := &staff.CoachLevel{}
	coach, err := staffModel.GetBrandCoachLevel(ctx, id, brandId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_coach_level",
			zap.Uint("coach_level_id", id),
			zap.Uint("brand_id", brandId),
			zap.Any("get_coach_level", coach), zap.Error(err))
		return nil, err
	}
	coachLevel.CoachLevelType = uint32(coach.ID)
	coachLevel.Name = coach.SettingName
	return coachLevel, nil
}

//获取员工身份
func GetIdentity(ctx context.Context, brandId, staffId uint) ([]*staff.IdentityResult, error) {
	staffs, err := staffModel.GetStaffIdentityRelation(ctx, brandId, staffId)
	identity := make([]*staff.IdentityResult, 0)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_identity", zap.Uint("brand_id", brandId),
			zap.Uint("staff_id", staffId),
			zap.Any("get_staff_identity", staffs), zap.Error(err))
		return nil, err
	}
	for _, staffIdentity := range staffs {
		identity = append(identity, EnumIdentity(staffIdentity.Identity))
	}
	return identity, nil
}

//枚举员工身份
func EnumIdentity(identity int8) *staff.IdentityResult {
	i := &staff.IdentityResult{}
	i.IdentityType = uint32(identity)
	switch identity {
	case staffModel.ENUM_ONE:
		i.Name = staffModel.ORDINARY_IDENTITY
	case staffModel.ENUM_TWO:
		i.Name = staffModel.MEMBERSHIP_SALES
	case staffModel.ENUM_THREE:
		i.Name = staffModel.TEAM_COACH
	case staffModel.ENUM_FOUR:
		i.Name = staffModel.PERSONAL_COACH
	case staffModel.ENUM_FIVE:
		i.Name = staffModel.SWIMMING_COACH
	}
	return i
}

//获取员工所属门店
func GetStaffShop(ctx context.Context, brandId, shopId, staffId uint) ([]*staff.ShopResult, error) {
	//当shopId等于0,表示在品牌维度
	if shopId == staffModel.SHOP_ID {
		shops, err := staffModel.GetShopStaffs(ctx, brandId, staffId)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_staff_shop", zap.Uint("brand_id", brandId),
				zap.Uint("staff_id", staffId),
				zap.Any("get_staff_shop", shops), zap.Error(err))
			return nil, err
		}
		sp, err := CheckStaffShop(shops, ctx, shopId)
		return sp, nil
	} else {
		shops, err := staffModel.GetShopStaffRelation(ctx, brandId, shopId, staffId)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_staff_shop", zap.Uint("brand_id", brandId),
				zap.Uint("shop_id", shopId),
				zap.Uint("staff_id", staffId),
				zap.Any("get_staff_shop", shops), zap.Error(err))
			return nil, err
		}
		sp, err := CheckStaffShop(shops, ctx, shopId)
		return sp, nil
	}
}

func CheckStaffShop(shops []*staffModel.ShopStaffRelation, ctx context.Context, shopId uint) ([]*staff.ShopResult, error) {
	sp := make([]*staff.ShopResult, len(shops))
	for index, staffShop := range shops {
		s := &staff.ShopResult{}
		sts, err := shopModel.ShopByID(ctx, staffShop.ShopID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_shop",
				zap.Uint("shop_id", shopId),
				zap.Any("get_shop", shops), zap.Error(err))
			return nil, err
		}
		s.ShopId = uint32(sts.ID)
		s.ShopName = sts.ShopName
		sp[index] = s
	}
	return sp, nil
}

//获取工作性质
func GetNatureWork(natureWork int8) *staff.NatureWorkResult {
	n := &staff.NatureWorkResult{}
	n.WorkId = uint32(natureWork)
	switch natureWork {
	case staffModel.ENUM_ONE:
		n.WorkName = staffModel.FULL_TIME
	case staffModel.ENUM_TWO:
		n.WorkName = staffModel.PART_TIME
	case staffModel.ENUM_THREE:
		n.WorkName = staffModel.PRACTICE
	}
	return n
}

//获取系统权限
func GetPermission(pm int8) *staff.PermissionResult {
	s := &staff.PermissionResult{}
	s.PermissionType = uint32(pm)
	switch pm {
	case staffModel.ENUM_ZERO:
		s.PermissionName = staffModel.UNAVAILABLE
	case staffModel.ENUM_ONE:
		s.PermissionName = staffModel.AVAILABLE
	}
	return s
}

//获取学历
func GetEducation(ed int8) *staff.EducationResult {
	e := &staff.EducationResult{}
	e.EducationType = uint32(ed)
	switch ed {
	case staffModel.ENUM_ZERO:
		e.EducationName = staffModel.UNFILLED
	case staffModel.ENUM_ONE:
		e.EducationName = staffModel.PRIMARY_SCHOOL
	case staffModel.ENUM_TWO:
		e.EducationName = staffModel.JUNIOR_HIGH_SCHOOL
	case staffModel.ENUM_THREE:
		e.EducationName = staffModel.HIGH_SCHOOL
	case staffModel.ENUM_FOUR:
		e.EducationName = staffModel.SECONDARY_SCHOOL
	case staffModel.ENUM_FIVE:
		e.EducationName = staffModel.COLLEGE
	case staffModel.ENUM_SIX:
		e.EducationName = staffModel.BACHELOR
	case staffModel.ENUM_SEVEN:
		e.EducationName = staffModel.MASTER
	case staffModel.ENUM_EIGHT:
		e.EducationName = staffModel.DOCTOR
	}
	return e
}

//获取婚姻
func GetMarryStatus(mr int8) *staff.MarryStatusResult {
	m := &staff.MarryStatusResult{}
	m.MarryType = uint32(mr)
	switch mr {
	case staffModel.ENUM_ZERO:
		m.MarryName = staffModel.UNFILLED
	case staffModel.ENUM_ONE:
		m.MarryName = staffModel.MARRIED
	case staffModel.ENUM_TWO:
		m.MarryName = staffModel.UNMARRIED
	}

	return m
}

//获取子女
func GetChildrenStatus(cs int8) *staff.ChildrenStatusResult {
	c := &staff.ChildrenStatusResult{}
	c.ChildrenType = uint32(cs)
	switch cs {
	case staffModel.ENUM_ZERO:
		c.ChildrenName = staffModel.UNFILLED
	case staffModel.ENUM_ONE:
		c.ChildrenName = staffModel.HAVE
	case staffModel.ENUM_TWO:
		c.ChildrenName = staffModel.HAVE_NOT
	}
	return c
}

//获取员工擅长项目
func GetSpecialty(ctx context.Context, brandId, staffId uint) ([]*staff.SpecialtyIdResult, error) {
	sp := make([]*staff.SpecialtyIdResult, 0)
	staffs, err := staffModel.GetStaffSpecialtyRelation(ctx, brandId, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_specialty", zap.Uint("brand_id", brandId),
		 zap.Uint("staff_id", staffId),
			zap.Any("get_staff_specialty", sp), zap.Error(err))
		return nil, err
	}
	for _, st := range staffs {
		s := &staff.SpecialtyIdResult{}
		sps, err := staffModel.BrandSpecialtySettingByID(ctx, st.SpecialtyID)
		if err != nil {
			utils.GetTraceLog(ctx).Error("get_brand_specialty",
				zap.Uint("specialty_id", st.SpecialtyID),
				zap.Any("get_brand_specialty", sps), zap.Error(err))
			return nil, err
		}
		s.SpecialtyType = uint32(sps.ID)
		s.SpecialtyName = sps.SpecialtyName
		sp = append(sp, s)
	}
	return sp, nil
}

//获取员工专业证书
func GetCertification(ctx context.Context, staffId uint) ([]*staff.CertificationNameResult, error) {
	cer := make([]*staff.CertificationNameResult, 0)
	staffMien, err := staffModel.GetStaffMienCertification(ctx, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_certification", zap.Uint("staff_id", staffId),
			zap.Any("get_certification", staffMien), zap.Error(err))
		return nil, err
	}
	for _, st := range staffMien {
		c := &staff.CertificationNameResult{}
		c.CertificationType = uint32(st.ID)
		c.CertificationName = st.CertificationName
		cer = append(cer, c)
	}
	return cer, nil
}

//获取对外展示
func GetIsShow(isShow int8) *staff.IsShow {
	sw := &staff.IsShow{}
	sw.ShowType = uint32(isShow)
	switch isShow {
	case staffModel.ENUM_ZERO:
		sw.ShowName = staffModel.NOT_SHOWING
	case staffModel.ENUM_ONE:
		sw.ShowName = staffModel.SHOW
	}
	return sw
}

//图片url
func GetImages(ctx context.Context, albumId uint) ([]*staff.ImagesResult, error) {
	images := make([]*staff.ImagesResult, 0)
	albums, err :=  imageModel.GetAlbumImagesByCoverType(ctx, albumId, staffModel.ORDINARY_IMAGE)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_album_images",
			zap.Uint("album_id", albumId),
			zap.Any("get_album_images", albums), zap.Error(err))
			return nil, err
	}
	for _, img := range albums {
		i := &staff.ImagesResult{}
		i.ImageId = uint32(img.ID)
		i.ImageUrl = string(img.ImageURL)
		images = append(images, i)
	}
	return images, nil
}

//国家编码
func GetCountryCode(ctx context.Context, countId uint) (*staff.CountryCode, error) {
	c := &staff.CountryCode{}
	country, err := countryModel.CountryCodeByID(ctx, countId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_country_code",
			zap.Uint("country_id", countId),
			zap.Any("get_country_code", country), zap.Error(err))
		return nil, err
	}
	c.CountryId = uint32(country.ID)
	c.CountryCode = uint32(country.PhoneCode)
	return c, nil

}

//员工头像
func GetStaffAvatar(ctx context.Context, avatarId uint) (*staff.AvatarIdResult, error) {
	at := &staff.AvatarIdResult{}
	avatar, err := imageModel.GetImagesByID(ctx, avatarId, staffModel.AVATAR_IMAGE)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_specialty", zap.Uint("avatar_id", avatarId),
			zap.Uint("cover_type", staffModel.AVATAR_IMAGE),
			zap.Any("get_staff_avatar", avatar), zap.Error(err))
		return nil, err
	}
	at.ImageId = uint32(avatar.ID)
	at.ImageUrl = avatar.ImageURL
	return at, nil
}

//员工人脸
func GetStaffFaceImages(ctx context.Context, faceId uint) (*staff.FaceIdResult, error) {
	at := &staff.FaceIdResult{}
	faces, err := imageModel.GetImagesByID(ctx, faceId, staffModel.FACE_IMAGE)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_specialty", zap.Uint("face_id", faceId),
			zap.Uint("cover_type", staffModel.FACE_IMAGE),
			zap.Any("get_staff_face", faces), zap.Error(err))
		return nil, err
	}
	at.ImageId = uint32(faces.ID)
	at.ImageUrl = faces.ImageURL
	return at, nil
}

//在职状态
func GetWorkStatus(work int8) *staff.WorkStatusResult {
	w := &staff.WorkStatusResult{}
	w.WorkId = uint32(work)
	switch work {
	case staffModel.ENUM_ONE:
		w.WorkName = staffModel.IN_SERVICE
	case staffModel.ENUM_TWO:
		w.WorkName = staffModel.STAFF_INACTIVE
	}
	return w
}

//删除员工账号
func DelStaffAccount(ctx context.Context, staffId uint) error {
	staffAccount, err := staffModel.GetStaffAccount(ctx, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_account",
			zap.Uint("staff_id", staffId),
			zap.Any("staff_account", staffAccount), zap.Error(err))
		return nil
	}
	err = staffAccount.Delete(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("delete_staff_account", zap.Any("staff_account", staffAccount), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_DELETE_STAFF_ACCOUNT_ERROR)
	}
	return nil
}

//品牌删除员工
func DelStaff(ctx context.Context, staffId uint) error {
	//开启事物
	ctx, err := components.M.BeginTransaction(ctx, 2 * time.Second)
	if err != nil {
		return errors.NewFromError(err)
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	staffs, err := staffModel.StaffByID(ctx, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_info",
			zap.Uint("staff_id", staffId),
			zap.Any("staff_info", staffs), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_GET_STAFF_INFO_ERROR)
	}
	//TODO 员工若存在绑定事物时，必须close掉才能删除 业务暂无
	//删除员工身份关系
	err = DelStaffIdentity(ctx,staffs.BrandID, staffs.ID)
	if err != nil {
		return err
	}
	//删除员工绑定场馆关系
	err = DelStaffBindShop(ctx, staffs.ID, staffs.BrandID)
	if err != nil {
		return err
	}
	//删除员工擅长项目关系
	err = DelStaffSpecialty(ctx, staffs.BrandID, staffs.ID)
	if err != nil {
		return err
	}
	//删除员工证书关系
	err = DelStaffMienCertification(ctx, staffs.ID)
	if err != nil {
		return err
	}
	//删除员工账号
	err = DelStaffAccount(ctx, staffId)
	if err != nil {
		return err
	}
	err = staffs.Delete(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("delete_staff",
			zap.Uint("staff_id", staffId),
			zap.Any("staff_info", staffs), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_DELETE_STAFF_INFO_ERROR)
	}

	_ = components.M.Commit(ctx)

	return nil
}

//员工账户信息
func GetStaffAccount(ctx context.Context, staffId uint32) (*staff.StaffAccountResponse, error) {
	staffAccount, err := staffModel.GetStaffBankCard(ctx, staffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_account_info",
			zap.Uint32("staff_id", staffId),
			zap.Any("staff_account", staffAccount), zap.Error(err))
		return nil, errors.NewFromCode(business_errors.StaffError_GET_STAFF_ACCOUNT_ERROR)
	}
	sts, err := staffModel.StaffByID(ctx, uint(staffId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_info",
			zap.Uint32("staff_id", staffId),
			zap.Any("staff_info", sts), zap.Error(err))
		return nil, errors.NewFromCode(business_errors.StaffError_GET_STAFF_INFO_ERROR)
	}
	staffs := &staff.StaffAccountResponse{}
	staffs.Status = errors.GetResHeader(err)
	staffs.StaffName = sts.StaffName
	staffs.AccountName = staffAccount.AccountName
	staffs.CardNumber = staffAccount.CardNumber
	staffs.BankName = staffAccount.BankName
	staffs.AccountId = uint32(staffAccount.ID)
	return staffs, nil
}

//新增员工账户信息
func CreateStaffBankAccount(ctx context.Context, request *staff.StaffBankRequest) error {
	staffAccount := &staffModel.StaffBankCard{}
	staffAccount.BankName = request.BankName
	staffAccount.CardNumber = request.CardNumber
	staffAccount.AccountName = request.AccountName
	staffAccount.StaffID = uint(request.StaffId)
	staffAccount.CreatedTime = uint(time.Now().Unix())
	staffAccount.UpdatedTime = uint(time.Now().Unix())
	err := staffAccount.Insert(ctx)
	if err != nil {
		utils.GetTraceLog(ctx).Error("staff", zap.String("method", "create_staff_bank_account"), zap.Any("staff_account", staffAccount), zap.Error(err))
		return errors.NewFromCode(business_errors.StaffError_CREATE_STAFF_ACCOUNT_ERROR)
	}
	return nil
}

//修改员工账户信息
func UpdateStaffBankAccount(ctx context.Context, request *staff.StaffBankRequest) error {
	staffs, err := staffModel.GetStaffBankCard(ctx, request.StaffId)
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_account",
			zap.Uint32("staff_id", request.StaffId),
			zap.Any("staff_account", staffs), zap.Error(err))
		return nil
	}
	if staffs != nil {
		staffs.BankName = request.BankName
		staffs.CardNumber = request.CardNumber
		staffs.AccountName = request.AccountName
		staffs.BrandID = uint(request.BrandId)
		staffs.UpdatedTime = uint(time.Now().Unix())
		err = staffs.UpdateStaffBank(ctx, request.StaffId)
		if err != nil {
			utils.GetTraceLog(ctx).Error("staff", zap.String("method", "update_staff_bank_account"), zap.Any("staff_account", staffs), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_UPDATE_STAFF_ACCOUNT_ERROR)
		}
	} else {
		err = CreateStaffBankAccount(ctx, request)
		if err != nil {
			return err
		}
	}
	return nil
}

//员工职位信息
func GetStaffPosition(ctx context.Context, staffId uint32) (*staff.StaffPositionResponse, error) {
	//TODO 暂无待处理员工事件和薪资模版
	staffs := &staff.StaffPositionResponse{}
	stf, err := staffModel.StaffByID(ctx, uint(staffId))
	if err != nil {
		utils.GetTraceLog(ctx).Error("get_staff_info",
			zap.Uint32("staff_id", staffId),
			zap.Any("staff_info", stf), zap.Error(err))
		return nil, errors.NewFromCode(business_errors.StaffError_GET_STAFF_INFO_ERROR)
	}
	staffs.Status = errors.GetResHeader(err)
	staffs.StaffName = stf.StaffName
	staffs.NatureWork = GetNatureWork(stf.NatureWork)
	staffIdentity, err := GetIdentity(ctx, uint(stf.BrandID), uint(stf.ID))
	if err == nil {
		staffs.Identity = staffIdentity
	}
	coachLevel, err := GetCoachLevel(ctx, stf.CoachLevelID, stf.BrandID)
	if err == nil {
		staffs.CoachLevelId = coachLevel
	}
	return staffs, nil
}

//保存员工职位
func StoreStaffPosition(ctx context.Context, request *staff.StaffPositionRequest) error {
	//开启事物
	ctx, err := components.M.BeginTransaction(ctx, 2 * time.Second)
	if err != nil {
		return errors.NewFromError(err)
	}

	defer func() {
		if err != nil {
			_ = components.M.Rollback(ctx)
		}
	}()

	staffs := staffModel.Staff{}
	staffs.ID = uint(request.StaffId)
	staffs.NatureWork = int8(request.NatureWork)
	staffs.CoachLevelID = uint(request.CoachLevelId)
	staffs.UpdatedTime = uint(time.Now().Unix())
	staffs.Update(ctx)

	//删除员工职位关系
	err = DelStaffIdentity(ctx, staffs.BrandID, staffs.ID)
	if err != nil {
		return err
	}
	//重新绑定员工身份
	for _, identity := range request.Identity {
		staffIdentity := staffModel.StaffIdentityRelation{}
		staffIdentity.StaffID = uint(staffs.ID)
		staffIdentity.BrandID = uint(staffs.BrandID)
		staffIdentity.Identity = int8(identity)
		staffIdentity.CreatedTime = uint(time.Now().Unix())
		staffIdentity.UpdatedTime = uint(time.Now().Unix())
		err := staffIdentity.Insert(ctx)
		if err != nil {
			utils.GetTraceLog(ctx).Error("staff", zap.String("method", "bind_identity_relation"), zap.Any("identity", staffIdentity), zap.Error(err))
			return errors.NewFromCode(business_errors.StaffError_BIND_STAFF_IDENTITY_ERROR)
		}
	}
	//TODO 暂无待处理员工事件和薪资模版
	_ = components.M.Commit(ctx)
	return nil
}
