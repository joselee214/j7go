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

// ShopCourseRelation represents a row from 'saas.shop_course_relation'.
type ShopCourseRelation struct {
	ID          int  `json:"id"`           // id
	BrandID     uint `json:"brand_id"`     // brand_id
	ShopID      uint `json:"shop_id"`      // shop_id
	CourseType  int8 `json:"course_type"`  // course_type
	CourseID    uint `json:"course_id"`    // course_id
	IsDel       int8 `json:"is_del"`       // is_del
	CreatedTime uint `json:"created_time"` // created_time
	UpdatedTime uint `json:"updated_time"` // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ShopCourseRelation exists in the database.
func (scr *ShopCourseRelation) Exists() bool { //shop_course_relation
	return scr._exists
}

// Deleted provides information if the ShopCourseRelation has been deleted from the database.
func (scr *ShopCourseRelation) Deleted() bool {
	return scr._deleted
}

// Get table name
func GetShopCourseRelationTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "shop_course_relation", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ShopCourseRelation to the database.
func (scr *ShopCourseRelation) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if scr._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetShopCourseRelationTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`id, brand_id, shop_id, course_type, course_id, is_del, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, scr.ID, scr.BrandID, scr.ShopID, scr.CourseType, scr.CourseID, scr.IsDel, scr.CreatedTime, scr.UpdatedTime)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, scr.ID, scr.BrandID, scr.ShopID, scr.CourseType, scr.CourseID, scr.IsDel, scr.CreatedTime, scr.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, scr.ID, scr.BrandID, scr.ShopID, scr.CourseType, scr.CourseID, scr.IsDel, scr.CreatedTime, scr.UpdatedTime)
	}

	if err != nil {
		return err
	}

	// set existence
	scr._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	scr.ID = int(id)
	scr._exists = true

	return nil
}

// Update updates the ShopCourseRelation in the database.
func (scr *ShopCourseRelation) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if scr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetShopCourseRelationTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`brand_id = ?, shop_id = ?, course_type = ?, course_id = ?, is_del = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, scr.BrandID, scr.ShopID, scr.CourseType, scr.CourseID, scr.IsDel, scr.CreatedTime, scr.UpdatedTime, scr.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, scr.BrandID, scr.ShopID, scr.CourseType, scr.CourseID, scr.IsDel, scr.CreatedTime, scr.UpdatedTime, scr.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, scr.BrandID, scr.ShopID, scr.CourseType, scr.CourseID, scr.IsDel, scr.CreatedTime, scr.UpdatedTime, scr.ID)
	}
	return err
}

// Save saves the ShopCourseRelation to the database.
func (scr *ShopCourseRelation) Save(ctx context.Context) error {
	if scr.Exists() {
		return scr.Update(ctx)
	}

	return scr.Insert(ctx)
}

// Delete deletes the ShopCourseRelation from the database.
func (scr *ShopCourseRelation) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if scr._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetShopCourseRelationTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, scr.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, scr.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, scr.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	scr._deleted = true

	return nil
}

// ShopCourseRelationByID retrieves a row from 'saas.shop_course_relation' as a ShopCourseRelation.
//
// Generated from index 'shop_course_relation_id_pkey'.
func ShopCourseRelationByID(ctx context.Context, id int, key ...interface{}) (*ShopCourseRelation, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetShopCourseRelationTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, brand_id, shop_id, course_type, course_id, is_del, created_time, updated_time ` +
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
	scr := ShopCourseRelation{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&scr.ID, &scr.BrandID, &scr.ShopID, &scr.CourseType, &scr.CourseID, &scr.IsDel, &scr.CreatedTime, &scr.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&scr.ID, &scr.BrandID, &scr.ShopID, &scr.CourseType, &scr.CourseID, &scr.IsDel, &scr.CreatedTime, &scr.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &scr, nil
}