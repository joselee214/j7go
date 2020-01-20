// Package test contains the types for schema 'saas'.
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

// PersonalCourseTemplate represents a row from 'saas.personal_course_template'.
type PersonalCourseTemplate struct {
	ID             uint   `json:"id"`              // id
	BrandID        uint   `json:"brand_id"`        // brand_id
	ShopID         uint   `json:"shop_id"`         // shop_id
	CourseName     string `json:"course_name"`     // course_name
	CategoryID     uint   `json:"category_id"`     // category_id
	Duration       uint   `json:"duration"`        // duration
	EffectiveUnit  uint   `json:"effective_unit"`  // effective_unit
	TimeUnit       int8   `json:"time_unit"`       // time_unit
	Price          uint   `json:"price"`           // price
	TrainAim       string `json:"train_aim"`       // train_aim
	SuitablePeople string `json:"suitable_people"` // suitable_people
	AlbumID        uint   `json:"album_id"`        // album_id
	Description    string `json:"description"`     // description
	ShopSetting    int8   `json:"shop_setting"`    // shop_setting
	PriceSetting   int8   `json:"price_setting"`   // price_setting
	PublishChannel int8   `json:"publish_channel"` // publish_channel
	IsDel          int8   `json:"is_del"`          // is_del
	IsAvailable    int8   `json:"is_available"`    // is_available
	CreatedTime    uint   `json:"created_time"`    // created_time
	UpdatedTime    uint   `json:"updated_time"`    // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PersonalCourseTemplate exists in the database.
func (pct *PersonalCourseTemplate) Exists() bool { //personal_course_template
	return pct._exists
}

// Deleted provides information if the PersonalCourseTemplate has been deleted from the database.
func (pct *PersonalCourseTemplate) Deleted() bool {
	return pct._deleted
}

// Get table name
func GetPersonalCourseTemplateTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "personal_course_template", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the PersonalCourseTemplate to the database.
func (pct *PersonalCourseTemplate) Insert(ctx context.Context, key ...interface{}) error {
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

	tableName, err := GetPersonalCourseTemplateTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`brand_id, shop_id, course_name, category_id, duration, effective_unit, time_unit, price, train_aim, suitable_people, album_id, description, shop_setting, price_setting, publish_channel, is_del, is_available, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.CategoryID, pct.Duration, pct.EffectiveUnit, pct.TimeUnit, pct.Price, pct.TrainAim, pct.SuitablePeople, pct.AlbumID, pct.Description, pct.ShopSetting, pct.PriceSetting, pct.PublishChannel, pct.IsDel, pct.IsAvailable, pct.CreatedTime, pct.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.CategoryID, pct.Duration, pct.EffectiveUnit, pct.TimeUnit, pct.Price, pct.TrainAim, pct.SuitablePeople, pct.AlbumID, pct.Description, pct.ShopSetting, pct.PriceSetting, pct.PublishChannel, pct.IsDel, pct.IsAvailable, pct.CreatedTime, pct.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.CategoryID, pct.Duration, pct.EffectiveUnit, pct.TimeUnit, pct.Price, pct.TrainAim, pct.SuitablePeople, pct.AlbumID, pct.Description, pct.ShopSetting, pct.PriceSetting, pct.PublishChannel, pct.IsDel, pct.IsAvailable, pct.CreatedTime, pct.UpdatedTime)
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

// Update updates the PersonalCourseTemplate in the database.
func (pct *PersonalCourseTemplate) Update(ctx context.Context, key ...interface{}) error {
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

	tableName, err := GetPersonalCourseTemplateTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`brand_id = ?, shop_id = ?, course_name = ?, category_id = ?, duration = ?, effective_unit = ?, time_unit = ?, price = ?, train_aim = ?, suitable_people = ?, album_id = ?, description = ?, shop_setting = ?, price_setting = ?, publish_channel = ?, is_del = ?, is_available = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.CategoryID, pct.Duration, pct.EffectiveUnit, pct.TimeUnit, pct.Price, pct.TrainAim, pct.SuitablePeople, pct.AlbumID, pct.Description, pct.ShopSetting, pct.PriceSetting, pct.PublishChannel, pct.IsDel, pct.IsAvailable, pct.CreatedTime, pct.UpdatedTime, pct.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.CategoryID, pct.Duration, pct.EffectiveUnit, pct.TimeUnit, pct.Price, pct.TrainAim, pct.SuitablePeople, pct.AlbumID, pct.Description, pct.ShopSetting, pct.PriceSetting, pct.PublishChannel, pct.IsDel, pct.IsAvailable, pct.CreatedTime, pct.UpdatedTime, pct.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, pct.BrandID, pct.ShopID, pct.CourseName, pct.CategoryID, pct.Duration, pct.EffectiveUnit, pct.TimeUnit, pct.Price, pct.TrainAim, pct.SuitablePeople, pct.AlbumID, pct.Description, pct.ShopSetting, pct.PriceSetting, pct.PublishChannel, pct.IsDel, pct.IsAvailable, pct.CreatedTime, pct.UpdatedTime, pct.ID)
	}
	return err
}

// Save saves the PersonalCourseTemplate to the database.
func (pct *PersonalCourseTemplate) Save(ctx context.Context) error {
	if pct.Exists() {
		return pct.Update(ctx)
	}

	return pct.Insert(ctx)
}

// Delete deletes the PersonalCourseTemplate from the database.
func (pct *PersonalCourseTemplate) Delete(ctx context.Context, key ...interface{}) error {
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

	tableName, err := GetPersonalCourseTemplateTableName(key...)
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

// PersonalCourseTemplateByID retrieves a row from 'saas.personal_course_template' as a PersonalCourseTemplate.
//
// Generated from index 'personal_course_template_id_pkey'.
func PersonalCourseTemplateByID(ctx context.Context, id uint, key ...interface{}) (*PersonalCourseTemplate, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetPersonalCourseTemplateTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, brand_id, shop_id, course_name, category_id, duration, effective_unit, time_unit, price, train_aim, suitable_people, album_id, description, shop_setting, price_setting, publish_channel, is_del, is_available, created_time, updated_time ` +
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
	pct := PersonalCourseTemplate{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&pct.ID, &pct.BrandID, &pct.ShopID, &pct.CourseName, &pct.CategoryID, &pct.Duration, &pct.EffectiveUnit, &pct.TimeUnit, &pct.Price, &pct.TrainAim, &pct.SuitablePeople, &pct.AlbumID, &pct.Description, &pct.ShopSetting, &pct.PriceSetting, &pct.PublishChannel, &pct.IsDel, &pct.IsAvailable, &pct.CreatedTime, &pct.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&pct.ID, &pct.BrandID, &pct.ShopID, &pct.CourseName, &pct.CategoryID, &pct.Duration, &pct.EffectiveUnit, &pct.TimeUnit, &pct.Price, &pct.TrainAim, &pct.SuitablePeople, &pct.AlbumID, &pct.Description, &pct.ShopSetting, &pct.PriceSetting, &pct.PublishChannel, &pct.IsDel, &pct.IsAvailable, &pct.CreatedTime, &pct.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &pct, nil
}