// Package test contains the types for schema 'product'.
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

// CourseCategory represents a row from 'saas.course_category'.
type CourseCategory struct {
	ID           uint   `json:"id"`            // id
	BrandID      uint   `json:"brand_id"`      // brand_id
	ShopID       uint   `json:"shop_id"`       // shop_id
	Name         string `json:"name"`          // name
	CourseType   int8   `json:"course_type"`   // course_type
	OperatorName string `json:"operator_name"` // operator_name
	OperatorID   uint   `json:"operator_id"`   // operator_id
	IsDel        int8   `json:"is_del"`        // is_del
	CreatedTime  uint   `json:"created_time"`  // created_time
	UpdatedTime  uint   `json:"updated_time"`  // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the CourseCategory exists in the database.
func (cc *CourseCategory) Exists() bool { //course_category
	return cc._exists
}

// Deleted provides information if the CourseCategory has been deleted from the database.
func (cc *CourseCategory) Deleted() bool {
	return cc._deleted
}

// Get table name
func GetCourseCategoryTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "course_category", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the CourseCategory to the database.
func (cc *CourseCategory) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if cc._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetCourseCategoryTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`brand_id, shop_id, name, course_type, operator_name, operator_id, is_del, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, cc.BrandID, cc.ShopID, cc.Name, cc.CourseType, cc.OperatorName, cc.OperatorID, cc.IsDel, cc.CreatedTime, cc.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, cc.BrandID, cc.ShopID, cc.Name, cc.CourseType, cc.OperatorName, cc.OperatorID, cc.IsDel, cc.CreatedTime, cc.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, cc.BrandID, cc.ShopID, cc.Name, cc.CourseType, cc.OperatorName, cc.OperatorID, cc.IsDel, cc.CreatedTime, cc.UpdatedTime)
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
	cc.ID = uint(id)
	cc._exists = true

	return nil
}

// Update updates the CourseCategory in the database.
func (cc *CourseCategory) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if cc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetCourseCategoryTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`brand_id = ?, shop_id = ?, name = ?, course_type = ?, operator_name = ?, operator_id = ?, is_del = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, cc.BrandID, cc.ShopID, cc.Name, cc.CourseType, cc.OperatorName, cc.OperatorID, cc.IsDel, cc.CreatedTime, cc.UpdatedTime, cc.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, cc.BrandID, cc.ShopID, cc.Name, cc.CourseType, cc.OperatorName, cc.OperatorID, cc.IsDel, cc.CreatedTime, cc.UpdatedTime, cc.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, cc.BrandID, cc.ShopID, cc.Name, cc.CourseType, cc.OperatorName, cc.OperatorID, cc.IsDel, cc.CreatedTime, cc.UpdatedTime, cc.ID)
	}
	return err
}

// Save saves the CourseCategory to the database.
func (cc *CourseCategory) Save(ctx context.Context) error {
	if cc.Exists() {
		return cc.Update(ctx)
	}

	return cc.Insert(ctx)
}

// Delete deletes the CourseCategory from the database.
func (cc *CourseCategory) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if cc._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetCourseCategoryTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, cc.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, cc.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, cc.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	cc._deleted = true

	return nil
}

// CourseCategoryByID retrieves a row from 'saas.course_category' as a CourseCategory.
//
// Generated from index 'course_category_id_pkey'.
func CourseCategoryByID(ctx context.Context, id uint, key ...interface{}) (*CourseCategory, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetCourseCategoryTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, brand_id, shop_id, name, course_type, operator_name, operator_id, is_del, created_time, updated_time ` +
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
	cc := CourseCategory{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&cc.ID, &cc.BrandID, &cc.ShopID, &cc.Name, &cc.CourseType, &cc.OperatorName, &cc.OperatorID, &cc.IsDel, &cc.CreatedTime, &cc.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&cc.ID, &cc.BrandID, &cc.ShopID, &cc.Name, &cc.CourseType, &cc.OperatorName, &cc.OperatorID, &cc.IsDel, &cc.CreatedTime, &cc.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &cc, nil
}
