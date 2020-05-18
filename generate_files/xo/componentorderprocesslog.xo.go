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

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// ComponentOrderProcessLog represents a row from 'aypcddg.component_order_process_log'.
type ComponentOrderProcessLog struct {
	ID                      uint           `json:"id"`                         // id
	ComponentOrderProcessID int            `json:"component_order_process_id"` // component_order_process_id
	Title                   sql.NullString `json:"title"`                      // title
	AccessURL               string         `json:"access_url"`                 // access_url
	RequestParam            string         `json:"request_param"`              // request_param
	ResponseContent         string         `json:"response_content"`           // response_content
	Remark                  sql.NullString `json:"remark"`                     // remark
	CreatedAt               mysql.NullTime `json:"created_at"`                 // created_at
	UpdatedAt               mysql.NullTime `json:"updated_at"`                 // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ComponentOrderProcessLog exists in the database.
func (copl *ComponentOrderProcessLog) Exists() bool { //component_order_process_log
	return copl._exists
}

// Deleted provides information if the ComponentOrderProcessLog has been deleted from the database.
func (copl *ComponentOrderProcessLog) Deleted() bool {
	return copl._deleted
}

// Get table name
func GetComponentOrderProcessLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "component_order_process_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ComponentOrderProcessLog to the database.
func (copl *ComponentOrderProcessLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if copl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetComponentOrderProcessLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`component_order_process_id, title, access_url, request_param, response_content, remark, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, copl.ComponentOrderProcessID, copl.Title, copl.AccessURL, copl.RequestParam, copl.ResponseContent, copl.Remark, copl.CreatedAt, copl.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, copl.ComponentOrderProcessID, copl.Title, copl.AccessURL, copl.RequestParam, copl.ResponseContent, copl.Remark, copl.CreatedAt, copl.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, copl.ComponentOrderProcessID, copl.Title, copl.AccessURL, copl.RequestParam, copl.ResponseContent, copl.Remark, copl.CreatedAt, copl.UpdatedAt)
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
	copl.ID = uint(id)
	copl._exists = true

	return nil
}

// Update updates the ComponentOrderProcessLog in the database.
func (copl *ComponentOrderProcessLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if copl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetComponentOrderProcessLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`component_order_process_id = ?, title = ?, access_url = ?, request_param = ?, response_content = ?, remark = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, copl.ComponentOrderProcessID, copl.Title, copl.AccessURL, copl.RequestParam, copl.ResponseContent, copl.Remark, copl.CreatedAt, copl.UpdatedAt, copl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, copl.ComponentOrderProcessID, copl.Title, copl.AccessURL, copl.RequestParam, copl.ResponseContent, copl.Remark, copl.CreatedAt, copl.UpdatedAt, copl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, copl.ComponentOrderProcessID, copl.Title, copl.AccessURL, copl.RequestParam, copl.ResponseContent, copl.Remark, copl.CreatedAt, copl.UpdatedAt, copl.ID)
	}
	return err
}

// Save saves the ComponentOrderProcessLog to the database.
func (copl *ComponentOrderProcessLog) Save(ctx context.Context) error {
	if copl.Exists() {
		return copl.Update(ctx)
	}

	return copl.Insert(ctx)
}

// Delete deletes the ComponentOrderProcessLog from the database.
func (copl *ComponentOrderProcessLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if copl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetComponentOrderProcessLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, copl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, copl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, copl.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	copl._deleted = true

	return nil
}

// ComponentOrderProcessLogByID retrieves a row from 'aypcddg.component_order_process_log' as a ComponentOrderProcessLog.
//
// Generated from index 'component_order_process_log_id_pkey'.
func ComponentOrderProcessLogByID(ctx context.Context, id uint, key ...interface{}) (*ComponentOrderProcessLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetComponentOrderProcessLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, component_order_process_id, title, access_url, request_param, response_content, remark, created_at, updated_at ` +
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
	copl := ComponentOrderProcessLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&copl.ID, &copl.ComponentOrderProcessID, &copl.Title, &copl.AccessURL, &copl.RequestParam, &copl.ResponseContent, &copl.Remark, &copl.CreatedAt, &copl.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&copl.ID, &copl.ComponentOrderProcessID, &copl.Title, &copl.AccessURL, &copl.RequestParam, &copl.ResponseContent, &copl.Remark, &copl.CreatedAt, &copl.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &copl, nil
}