// Package xo contains the types for schema 'saas'.
package member

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"j7go/components"
	"j7go/utils"

	"go.uber.org/zap"
)

// Member represents a row from 'saas.member'.
type Member struct {
	ID                     uint   `json:"id"`                       // id
	BrandID                uint   `json:"brand_id"`                 // brand_id
	MemberName             string `json:"member_name"`              // member_name
	CountryPrefix          uint   `json:"country_prefix"`           // country_prefix
	Mobile                 string `json:"mobile"`                   // mobile
	RegisterType           int8   `json:"register_type"`            // register_type
	RegisterWay            uint   `json:"register_way"`             // register_way
	Sex                    int8   `json:"sex"`                      // sex
	Weight                 uint   `json:"weight"`                   // weight
	Height                 uint   `json:"height"`                   // height
	IDCardType             int8   `json:"id_card_type"`             // id_card_type
	IDCard                 string `json:"id_card"`                  // id_card
	Birthday               string `json:"birthday"`                 // birthday
	FirstRegTime           uint   `json:"first_reg_time"`           // first_reg_time
	Jobs                   string `json:"jobs"`                     // jobs
	EducationLevel         int8   `json:"education_level"`          // education_level
	MarriedType            int8   `json:"married_type"`             // married_type
	HasChildren            int8   `json:"has_children"`             // has_children
	MemberComment          string `json:"member_comment"`           // member_comment
	Email                  string `json:"email"`                    // email
	CountryID              uint   `json:"country_id"`               // country_id
	NationID               uint   `json:"nation_id"`                // nation_id
	EmergencyContactName   string `json:"emergency_contact_name"`   // emergency_contact_name
	EmergencyContactMobile string `json:"emergency_contact_mobile"` // emergency_contact_mobile
	ProvinceID             uint   `json:"province_id"`              // province_id
	ProvinceName           string `json:"province_name"`            // province_name
	CityID                 uint   `json:"city_id"`                  // city_id
	CityName               string `json:"city_name"`                // city_name
	DistrictID             uint   `json:"district_id"`              // district_id
	DistrictName           string `json:"district_name"`            // district_name
	LivingAddress          string `json:"living_address"`           // living_address
	FitnessGoal            string `json:"fitness_goal"`             // fitness_goal
	FitnessLevel           int8   `json:"fitness_level"`            // fitness_level
	IncomeLevel            string `json:"income_level"`             // income_level
	IsDel                  int8   `json:"is_del"`                   // is_del
	CreatedTime            uint   `json:"created_time"`             // created_time
	UpdatedTime            uint   `json:"updated_time"`             // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Member exists in the database.
func (m *Member) Exists() bool { //member
	return m._exists
}

// Deleted provides information if the Member has been deleted from the database.
func (m *Member) Deleted() bool {
	return m._deleted
}

// Get table name
func GetMemberTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "member", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Member to the database.
func (m *Member) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMemberTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`brand_id, member_name, country_prefix, mobile, register_type, register_way, sex, weight, height, id_card_type, id_card, birthday, first_reg_time, jobs, education_level, married_type, has_children, member_comment, email, country_id, nation_id, emergency_contact_name, emergency_contact_mobile, province_id, province_name, city_id, city_name, district_id, district_name, living_address, fitness_goal, fitness_level, income_level, is_del, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, m.BrandID, m.MemberName, m.CountryPrefix, m.Mobile, m.RegisterType, m.RegisterWay, m.Sex, m.Weight, m.Height, m.IDCardType, m.IDCard, m.Birthday, m.FirstRegTime, m.Jobs, m.EducationLevel, m.MarriedType, m.HasChildren, m.MemberComment, m.Email, m.CountryID, m.NationID, m.EmergencyContactName, m.EmergencyContactMobile, m.ProvinceID, m.ProvinceName, m.CityID, m.CityName, m.DistrictID, m.DistrictName, m.LivingAddress, m.FitnessGoal, m.FitnessLevel, m.IncomeLevel, m.IsDel, m.CreatedTime, m.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, m.BrandID, m.MemberName, m.CountryPrefix, m.Mobile, m.RegisterType, m.RegisterWay, m.Sex, m.Weight, m.Height, m.IDCardType, m.IDCard, m.Birthday, m.FirstRegTime, m.Jobs, m.EducationLevel, m.MarriedType, m.HasChildren, m.MemberComment, m.Email, m.CountryID, m.NationID, m.EmergencyContactName, m.EmergencyContactMobile, m.ProvinceID, m.ProvinceName, m.CityID, m.CityName, m.DistrictID, m.DistrictName, m.LivingAddress, m.FitnessGoal, m.FitnessLevel, m.IncomeLevel, m.IsDel, m.CreatedTime, m.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, m.BrandID, m.MemberName, m.CountryPrefix, m.Mobile, m.RegisterType, m.RegisterWay, m.Sex, m.Weight, m.Height, m.IDCardType, m.IDCard, m.Birthday, m.FirstRegTime, m.Jobs, m.EducationLevel, m.MarriedType, m.HasChildren, m.MemberComment, m.Email, m.CountryID, m.NationID, m.EmergencyContactName, m.EmergencyContactMobile, m.ProvinceID, m.ProvinceName, m.CityID, m.CityName, m.DistrictID, m.DistrictName, m.LivingAddress, m.FitnessGoal, m.FitnessLevel, m.IncomeLevel, m.IsDel, m.CreatedTime, m.UpdatedTime)
	}

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	m.ID = uint(id)
	m._exists = true

	return nil
}

// Update updates the Member in the database.
func (m *Member) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if m._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMemberTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`brand_id = ?, member_name = ?, country_prefix = ?, mobile = ?, register_type = ?, register_way = ?, sex = ?, weight = ?, height = ?, id_card_type = ?, id_card = ?, birthday = ?, first_reg_time = ?, jobs = ?, education_level = ?, married_type = ?, has_children = ?, member_comment = ?, email = ?, country_id = ?, nation_id = ?, emergency_contact_name = ?, emergency_contact_mobile = ?, province_id = ?, province_name = ?, city_id = ?, city_name = ?, district_id = ?, district_name = ?, living_address = ?, fitness_goal = ?, fitness_level = ?, income_level = ?, is_del = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, m.BrandID, m.MemberName, m.CountryPrefix, m.Mobile, m.RegisterType, m.RegisterWay, m.Sex, m.Weight, m.Height, m.IDCardType, m.IDCard, m.Birthday, m.FirstRegTime, m.Jobs, m.EducationLevel, m.MarriedType, m.HasChildren, m.MemberComment, m.Email, m.CountryID, m.NationID, m.EmergencyContactName, m.EmergencyContactMobile, m.ProvinceID, m.ProvinceName, m.CityID, m.CityName, m.DistrictID, m.DistrictName, m.LivingAddress, m.FitnessGoal, m.FitnessLevel, m.IncomeLevel, m.IsDel, m.CreatedTime, m.UpdatedTime, m.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, m.BrandID, m.MemberName, m.CountryPrefix, m.Mobile, m.RegisterType, m.RegisterWay, m.Sex, m.Weight, m.Height, m.IDCardType, m.IDCard, m.Birthday, m.FirstRegTime, m.Jobs, m.EducationLevel, m.MarriedType, m.HasChildren, m.MemberComment, m.Email, m.CountryID, m.NationID, m.EmergencyContactName, m.EmergencyContactMobile, m.ProvinceID, m.ProvinceName, m.CityID, m.CityName, m.DistrictID, m.DistrictName, m.LivingAddress, m.FitnessGoal, m.FitnessLevel, m.IncomeLevel, m.IsDel, m.CreatedTime, m.UpdatedTime, m.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, m.BrandID, m.MemberName, m.CountryPrefix, m.Mobile, m.RegisterType, m.RegisterWay, m.Sex, m.Weight, m.Height, m.IDCardType, m.IDCard, m.Birthday, m.FirstRegTime, m.Jobs, m.EducationLevel, m.MarriedType, m.HasChildren, m.MemberComment, m.Email, m.CountryID, m.NationID, m.EmergencyContactName, m.EmergencyContactMobile, m.ProvinceID, m.ProvinceName, m.CityID, m.CityName, m.DistrictID, m.DistrictName, m.LivingAddress, m.FitnessGoal, m.FitnessLevel, m.IncomeLevel, m.IsDel, m.CreatedTime, m.UpdatedTime, m.ID)
	}
	return err
}

// Save saves the Member to the database.
func (m *Member) Save(ctx context.Context) error {
	if m.Exists() {
		return m.Update(ctx)
	}

	return m.Insert(ctx)
}

// Delete deletes the Member from the database.
func (m *Member) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if m._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMemberTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, m.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, m.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, m.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	m._deleted = true

	return nil
}

// MemberByID retrieves a row from 'saas.member' as a Member.
//
// Generated from index 'member_id_pkey'.
func MemberByID(ctx context.Context, id uint, key ...interface{}) (*Member, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMemberTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, brand_id, member_name, country_prefix, mobile, register_type, register_way, sex, weight, height, id_card_type, id_card, birthday, first_reg_time, jobs, education_level, married_type, has_children, member_comment, email, country_id, nation_id, emergency_contact_name, emergency_contact_mobile, province_id, province_name, city_id, city_name, district_id, district_name, living_address, fitness_goal, fitness_level, income_level, is_del, created_time, updated_time ` +
		`FROM ` + tableName +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, id)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	m := Member{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&m.ID, &m.BrandID, &m.MemberName, &m.CountryPrefix, &m.Mobile, &m.RegisterType, &m.RegisterWay, &m.Sex, &m.Weight, &m.Height, &m.IDCardType, &m.IDCard, &m.Birthday, &m.FirstRegTime, &m.Jobs, &m.EducationLevel, &m.MarriedType, &m.HasChildren, &m.MemberComment, &m.Email, &m.CountryID, &m.NationID, &m.EmergencyContactName, &m.EmergencyContactMobile, &m.ProvinceID, &m.ProvinceName, &m.CityID, &m.CityName, &m.DistrictID, &m.DistrictName, &m.LivingAddress, &m.FitnessGoal, &m.FitnessLevel, &m.IncomeLevel, &m.IsDel, &m.CreatedTime, &m.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&m.ID, &m.BrandID, &m.MemberName, &m.CountryPrefix, &m.Mobile, &m.RegisterType, &m.RegisterWay, &m.Sex, &m.Weight, &m.Height, &m.IDCardType, &m.IDCard, &m.Birthday, &m.FirstRegTime, &m.Jobs, &m.EducationLevel, &m.MarriedType, &m.HasChildren, &m.MemberComment, &m.Email, &m.CountryID, &m.NationID, &m.EmergencyContactName, &m.EmergencyContactMobile, &m.ProvinceID, &m.ProvinceName, &m.CityID, &m.CityName, &m.DistrictID, &m.DistrictName, &m.LivingAddress, &m.FitnessGoal, &m.FitnessLevel, &m.IncomeLevel, &m.IsDel, &m.CreatedTime, &m.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}

//todo 没写完
func UpdateMemberInfoByConditionMap(ctx context.Context, memberId uint, condition map[string]interface{}) error {
	tableName, err := GetMemberTableName()
	if err != nil {
		return err
	}

	sqlstr, args, err := squirrel.
		Update(tableName).
		SetMap(condition).
		Where(squirrel.Eq{"id": memberId}).
		ToSql()
	if err != nil {
		utils.GetTraceLog(ctx).Error("sql_builder_error", zap.Uint("member_id", memberId), zap.Error(err))
	}

	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL,update_member_info_by_condition_map", fmt.Sprint(sqlstr)), zap.Any("args", args))

	var dbConn *sql.DB
	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	if tx != nil {
		_, err := tx.Exec(sqlstr, args...)
		if err != nil {
			return err
		}
	} else {
		_, err := dbConn.Exec(sqlstr, args...)
		if err != nil {
			return err
		}
	}

	return nil

}