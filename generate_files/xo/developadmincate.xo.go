// Package xo contains the types for schema 'aypcddg'.
package xo

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

// DevelopAdminCate represents a row from 'aypcddg.develop_admin_cate'.
type DevelopAdminCate struct {
	ID       uint           `json:"id"`        // id
	ParentID sql.NullInt64  `json:"parent_id"` // parent_id
	Name     sql.NullString `json:"name"`      // name
	URL      sql.NullString `json:"url"`       // url
	Link     string         `json:"link"`      // link
	MenuID   int            `json:"menu_id"`   // menu_id
	IsShow   int8           `json:"is_show"`   // is_show
	Sort     sql.NullInt64  `json:"sort"`      // sort

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DevelopAdminCate exists in the database.
func (dac *DevelopAdminCate) Exists() bool { //develop_admin_cate
	return dac._exists
}

// Deleted provides information if the DevelopAdminCate has been deleted from the database.
func (dac *DevelopAdminCate) Deleted() bool {
	return dac._deleted
}

// Get table name
func GetDevelopAdminCateTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "develop_admin_cate", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DevelopAdminCate to the database.
func (dac *DevelopAdminCate) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if dac._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDevelopAdminCateTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`parent_id, name, url, link, menu_id, is_show, sort` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dac.ParentID, dac.Name, dac.URL, dac.Link, dac.MenuID, dac.IsShow, dac.Sort)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dac.ParentID, dac.Name, dac.URL, dac.Link, dac.MenuID, dac.IsShow, dac.Sort)
	} else {
		res, err = dbConn.Exec(sqlstr, dac.ParentID, dac.Name, dac.URL, dac.Link, dac.MenuID, dac.IsShow, dac.Sort)
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
	dac.ID = uint(id)
	dac._exists = true

	return nil
}

// Update updates the DevelopAdminCate in the database.
func (dac *DevelopAdminCate) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dac._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDevelopAdminCateTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`parent_id = ?, name = ?, url = ?, link = ?, menu_id = ?, is_show = ?, sort = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dac.ParentID, dac.Name, dac.URL, dac.Link, dac.MenuID, dac.IsShow, dac.Sort, dac.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dac.ParentID, dac.Name, dac.URL, dac.Link, dac.MenuID, dac.IsShow, dac.Sort, dac.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dac.ParentID, dac.Name, dac.URL, dac.Link, dac.MenuID, dac.IsShow, dac.Sort, dac.ID)
	}
	return err
}

// Save saves the DevelopAdminCate to the database.
func (dac *DevelopAdminCate) Save(ctx context.Context) error {
	if dac.Exists() {
		return dac.Update(ctx)
	}

	return dac.Insert(ctx)
}

// Delete deletes the DevelopAdminCate from the database.
func (dac *DevelopAdminCate) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dac._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDevelopAdminCateTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dac.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dac.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dac.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	dac._deleted = true

	return nil
}

// DevelopAdminCateByID retrieves a row from 'aypcddg.develop_admin_cate' as a DevelopAdminCate.
//
// Generated from index 'develop_admin_cate_id_pkey'.
func DevelopAdminCateByID(ctx context.Context, id uint, key ...interface{}) (*DevelopAdminCate, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDevelopAdminCateTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, parent_id, name, url, link, menu_id, is_show, sort ` +
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
	dac := DevelopAdminCate{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&dac.ID, &dac.ParentID, &dac.Name, &dac.URL, &dac.Link, &dac.MenuID, &dac.IsShow, &dac.Sort)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&dac.ID, &dac.ParentID, &dac.Name, &dac.URL, &dac.Link, &dac.MenuID, &dac.IsShow, &dac.Sort)
		if err != nil {
			return nil, err
		}
	}

	return &dac, nil
}
