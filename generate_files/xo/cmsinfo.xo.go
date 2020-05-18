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

// CmsInfo represents a row from 'aypcddg.cms_info'.
type CmsInfo struct {
	ID        int            `json:"id"`          // id
	Title     sql.NullString `json:"title"`       // title
	Content   sql.NullString `json:"content"`     // content
	AddTime   sql.NullInt64  `json:"add_time"`    // add_time
	Sort      sql.NullInt64  `json:"sort"`        // sort
	CmsTypeID int            `json:"cms_type_id"` // cms_type_id
	Tag       string         `json:"tag"`         // tag
	XcxScene  sql.NullString `json:"xcx_scene"`   // xcx_scene

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the CmsInfo exists in the database.
func (ci *CmsInfo) Exists() bool { //cms_info
	return ci._exists
}

// Deleted provides information if the CmsInfo has been deleted from the database.
func (ci *CmsInfo) Deleted() bool {
	return ci._deleted
}

// Get table name
func GetCmsInfoTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "cms_info", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the CmsInfo to the database.
func (ci *CmsInfo) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ci._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetCmsInfoTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`id, title, content, add_time, sort, cms_type_id, xcx_scene` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ci.ID, ci.Title, ci.Content, ci.AddTime, ci.Sort, ci.CmsTypeID, ci.XcxScene)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ci.ID, ci.Title, ci.Content, ci.AddTime, ci.Sort, ci.CmsTypeID, ci.XcxScene)
	} else {
		res, err = dbConn.Exec(sqlstr, ci.ID, ci.Title, ci.Content, ci.AddTime, ci.Sort, ci.CmsTypeID, ci.XcxScene)
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
	ci.Tag = string(id)
	ci._exists = true

	return nil
}

// Update updates the CmsInfo in the database.
func (ci *CmsInfo) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ci._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetCmsInfoTableName(key...)
	if err != nil {
		return err
	}

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`title = ?, content = ?, add_time = ?, sort = ?, xcx_scene = ?` +
		` WHERE id = ? AND cms_type_id = ? AND tag = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ci.Title, ci.Content, ci.AddTime, ci.Sort, ci.XcxScene, ci.ID, ci.CmsTypeID, ci.Tag)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ci.Title, ci.Content, ci.AddTime, ci.Sort, ci.XcxScene, ci.ID, ci.CmsTypeID, ci.Tag)
	} else {
		_, err = dbConn.Exec(sqlstr, ci.Title, ci.Content, ci.AddTime, ci.Sort, ci.XcxScene, ci.ID, ci.CmsTypeID, ci.Tag)
	}
	return err
}

// Save saves the CmsInfo to the database.
func (ci *CmsInfo) Save(ctx context.Context) error {
	if ci.Exists() {
		return ci.Update(ctx)
	}

	return ci.Insert(ctx)
}

// Delete deletes the CmsInfo from the database.
func (ci *CmsInfo) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ci._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetCmsInfoTableName(key...)
	if err != nil {
		return err
	}
	//3

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE tag = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ci.Tag)))

	if tx != nil {
		_, err = tx.Exec(sqlstr, ci.Tag)
	} else {
		_, err = dbConn.Exec(sqlstr, ci.Tag)
	}
	if err != nil {
		return err
	}

	// set deleted
	ci._deleted = true

	return nil
}

// CmsInfoByTag retrieves a row from 'aypcddg.cms_info' as a CmsInfo.
//
// Generated from index 'cms_info_tag_pkey'.
func CmsInfoByTag(ctx context.Context, tag string, key ...interface{}) (*CmsInfo, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetCmsInfoTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, title, content, add_time, sort, cms_type_id, tag, xcx_scene ` +
		`FROM ` + tableName +
		` WHERE tag = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, tag)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ci := CmsInfo{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, tag).Scan(&ci.ID, &ci.Title, &ci.Content, &ci.AddTime, &ci.Sort, &ci.CmsTypeID, &ci.Tag, &ci.XcxScene)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, tag).Scan(&ci.ID, &ci.Title, &ci.Content, &ci.AddTime, &ci.Sort, &ci.CmsTypeID, &ci.Tag, &ci.XcxScene)
		if err != nil {
			return nil, err
		}
	}

	return &ci, nil
}
