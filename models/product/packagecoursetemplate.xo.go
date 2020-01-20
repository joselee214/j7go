// Package models contains the types for schema 'saas'.
package product

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"j7go/components"
	"j7go/utils"

	"go.uber.org/zap"
)

const (
	UNLIMITED_COURSE = 1
	SCOPE_COURSE     = 2
	FIXED_COURSE     = 3
)

// PackageCourseTemplate represents a row from 'saas.package_course_template'.
type PackageCourseTemplate struct {
	ID                uint   `json:"id"`                  // id
	BrandID           uint   `json:"brand_id"`            // brand_id
	ShopID            uint   `json:"shop_id"`             // shop_id
	CourseName        string `json:"course_name"`         // course_name
	Price             uint   `json:"price"`               // price
	PackageType       int8   `json:"package_type"`        // package_type
	TotalPrice        uint   `json:"total_price"`         // total_price
	TotalTimes        uint   `json:"total_times"`         // total_times
	IsTeam            int8   `json:"is_team"`             // is_team
	IsPersonal        int8   `json:"is_personal"`         // is_personal
	TeamTimes         uint   `json:"team_times"`          // team_times
	TeamUnitPrice     uint   `json:"team_unit_price"`     // team_unit_price
	Personal          uint   `json:"personal"`            // personal
	PersonalUnitPrice uint   `json:"personal_unit_price"` // personal_unit_price
	StartTime         uint   `json:"start_time"`          // start_time
	EndTime           uint   `json:"end_time"`            // end_time
	ValidDays         uint   `json:"valid_days"`          // valid_days
	SaleMode          int8   `json:"sale_mode"`           // sale_mode
	ImageID           uint   `json:"image_id"`            // image_id
	PublishChannel    int8   `json:"publish_channel"`     // publish_channel
	IsDel             int8   `json:"is_del"`              // is_del
	Intro             string `json:"intro"`               // intro
	Remarks           string `json:"remarks"`             // remarks
	UpdatedTime       uint   `json:"updated_time"`        // updated_time
	CreatedTime       uint   `json:"created_time"`        // created_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PackageCourseTemplate exists in the database.
func (pct *PackageCourseTemplate) Exists() bool { //package_course_template
	return pct._exists
}

// Deleted provides information if the PackageCourseTemplate has been deleted from the database.
func (pct *PackageCourseTemplate) Deleted() bool {
	return pct._deleted
}

// Get table name
func GetPackageCourseTemplateTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "package_course_template", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the PackageCourseTemplate to the database.
func (pct *PackageCourseTemplate) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if pct._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPackageCourseTemplateTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`brand_id, shop_id, course_name, price, package_type, total_price, total_times, is_team, is_personal, team_times, team_unit_price, personal, personal_unit_price, start_time, end_time, valid_days, sale_mode, image_id, publish_channel, is_del, intro, remarks, updated_time, created_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.Price, pct.PackageType, pct.TotalPrice, pct.TotalTimes, pct.IsTeam, pct.IsPersonal, pct.TeamTimes, pct.TeamUnitPrice, pct.Personal, pct.PersonalUnitPrice, pct.StartTime, pct.EndTime, pct.ValidDays, pct.SaleMode, pct.ImageID, pct.PublishChannel, pct.IsDel, pct.Intro, pct.Remarks, pct.UpdatedTime, pct.CreatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.Price, pct.PackageType, pct.TotalPrice, pct.TotalTimes, pct.IsTeam, pct.IsPersonal, pct.TeamTimes, pct.TeamUnitPrice, pct.Personal, pct.PersonalUnitPrice, pct.StartTime, pct.EndTime, pct.ValidDays, pct.SaleMode, pct.ImageID, pct.PublishChannel, pct.IsDel, pct.Intro, pct.Remarks, pct.UpdatedTime, pct.CreatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.Price, pct.PackageType, pct.TotalPrice, pct.TotalTimes, pct.IsTeam, pct.IsPersonal, pct.TeamTimes, pct.TeamUnitPrice, pct.Personal, pct.PersonalUnitPrice, pct.StartTime, pct.EndTime, pct.ValidDays, pct.SaleMode, pct.ImageID, pct.PublishChannel, pct.IsDel, pct.Intro, pct.Remarks, pct.UpdatedTime, pct.CreatedTime)
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
	pct.ID = uint(id)
	pct._exists = true

	return nil
}

// Update updates the PackageCourseTemplate in the database.
func (pct *PackageCourseTemplate) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if pct._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPackageCourseTemplateTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`brand_id = ?, shop_id = ?, course_name = ?, price = ?, package_type = ?, total_price = ?, total_times = ?, is_team = ?, is_personal = ?, team_times = ?, team_unit_price = ?, personal = ?, personal_unit_price = ?, start_time = ?, end_time = ?, valid_days = ?, sale_mode = ?, image_id = ?, publish_channel = ?, is_del = ?, intro = ?, remarks = ?, updated_time = ?, created_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.Price, pct.PackageType, pct.TotalPrice, pct.TotalTimes, pct.IsTeam, pct.IsPersonal, pct.TeamTimes, pct.TeamUnitPrice, pct.Personal, pct.PersonalUnitPrice, pct.StartTime, pct.EndTime, pct.ValidDays, pct.SaleMode, pct.ImageID, pct.PublishChannel, pct.IsDel, pct.Intro, pct.Remarks, pct.UpdatedTime, pct.CreatedTime, pct.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.Price, pct.PackageType, pct.TotalPrice, pct.TotalTimes, pct.IsTeam, pct.IsPersonal, pct.TeamTimes, pct.TeamUnitPrice, pct.Personal, pct.PersonalUnitPrice, pct.StartTime, pct.EndTime, pct.ValidDays, pct.SaleMode, pct.ImageID, pct.PublishChannel, pct.IsDel, pct.Intro, pct.Remarks, pct.UpdatedTime, pct.CreatedTime, pct.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.Price, pct.PackageType, pct.TotalPrice, pct.TotalTimes, pct.IsTeam, pct.IsPersonal, pct.TeamTimes, pct.TeamUnitPrice, pct.Personal, pct.PersonalUnitPrice, pct.StartTime, pct.EndTime, pct.ValidDays, pct.SaleMode, pct.ImageID, pct.PublishChannel, pct.IsDel, pct.Intro, pct.Remarks, pct.UpdatedTime, pct.CreatedTime, pct.ID)
	}
	return err
}

// Save saves the PackageCourseTemplate to the database.
func (pct *PackageCourseTemplate) Save(ctx context.Context) error {
	if pct.Exists() {
		return pct.Update(ctx)
	}

	return pct.Insert(ctx)
}

// Delete deletes the PackageCourseTemplate from the database.
func (pct *PackageCourseTemplate) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if pct._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPackageCourseTemplateTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pct.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, pct.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, pct.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	pct._deleted = true

	return nil
}

// PackageCourseTemplateByID retrieves a row from 'saas.package_course_template' as a PackageCourseTemplate.
//
// Generated from index 'package_course_template_id_pkey'.
func PackageCourseTemplateByID(ctx context.Context, id uint, key ...interface{}) (*PackageCourseTemplate, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetPackageCourseTemplateTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, brand_id, shop_id, course_name, price, package_type, total_price, total_times, is_team, is_personal, team_times, team_unit_price, personal, personal_unit_price, start_time, end_time, valid_days, sale_mode, image_id, publish_channel, is_del, intro, remarks, updated_time, created_time ` +
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
	pct := PackageCourseTemplate{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&pct.ID, &pct.BrandID, &pct.ShopID, &pct.CourseName, &pct.Price, &pct.PackageType, &pct.TotalPrice, &pct.TotalTimes, &pct.IsTeam, &pct.IsPersonal, &pct.TeamTimes, &pct.TeamUnitPrice, &pct.Personal, &pct.PersonalUnitPrice, &pct.StartTime, &pct.EndTime, &pct.ValidDays, &pct.SaleMode, &pct.ImageID, &pct.PublishChannel, &pct.IsDel, &pct.Intro, &pct.Remarks, &pct.UpdatedTime, &pct.CreatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&pct.ID, &pct.BrandID, &pct.ShopID, &pct.CourseName, &pct.Price, &pct.PackageType, &pct.TotalPrice, &pct.TotalTimes, &pct.IsTeam, &pct.IsPersonal, &pct.TeamTimes, &pct.TeamUnitPrice, &pct.Personal, &pct.PersonalUnitPrice, &pct.StartTime, &pct.EndTime, &pct.ValidDays, &pct.SaleMode, &pct.ImageID, &pct.PublishChannel, &pct.IsDel, &pct.Intro, &pct.Remarks, &pct.UpdatedTime, &pct.CreatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &pct, nil
}
