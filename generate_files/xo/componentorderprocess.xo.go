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

// ComponentOrderProcess represents a row from 'aypcddg.component_order_process'.
type ComponentOrderProcess struct {
	ID        uint           `json:"id"`         // id
	Orderid   int            `json:"orderid"`    // orderid
	Event     int            `json:"event"`      // event
	DoStatus  int16          `json:"do_status"`  // do_status
	AccessURL string         `json:"access_url"` // access_url
	Remark    sql.NullString `json:"remark"`     // remark
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ComponentOrderProcess exists in the database.
func (cop *ComponentOrderProcess) Exists() bool { //component_order_process
	return cop._exists
}

// Deleted provides information if the ComponentOrderProcess has been deleted from the database.
func (cop *ComponentOrderProcess) Deleted() bool {
	return cop._deleted
}

// Get table name
func GetComponentOrderProcessTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "component_order_process", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ComponentOrderProcess to the database.
func (cop *ComponentOrderProcess) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if cop._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetComponentOrderProcessTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`orderid, event, do_status, access_url, remark, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, cop.Orderid, cop.Event, cop.DoStatus, cop.AccessURL, cop.Remark, cop.CreatedAt, cop.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, cop.Orderid, cop.Event, cop.DoStatus, cop.AccessURL, cop.Remark, cop.CreatedAt, cop.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, cop.Orderid, cop.Event, cop.DoStatus, cop.AccessURL, cop.Remark, cop.CreatedAt, cop.UpdatedAt)
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
	cop.ID = uint(id)
	cop._exists = true

	return nil
}

// Update updates the ComponentOrderProcess in the database.
func (cop *ComponentOrderProcess) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if cop._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetComponentOrderProcessTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`orderid = ?, event = ?, do_status = ?, access_url = ?, remark = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, cop.Orderid, cop.Event, cop.DoStatus, cop.AccessURL, cop.Remark, cop.CreatedAt, cop.UpdatedAt, cop.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, cop.Orderid, cop.Event, cop.DoStatus, cop.AccessURL, cop.Remark, cop.CreatedAt, cop.UpdatedAt, cop.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, cop.Orderid, cop.Event, cop.DoStatus, cop.AccessURL, cop.Remark, cop.CreatedAt, cop.UpdatedAt, cop.ID)
	}
	return err
}

// Save saves the ComponentOrderProcess to the database.
func (cop *ComponentOrderProcess) Save(ctx context.Context) error {
	if cop.Exists() {
		return cop.Update(ctx)
	}

	return cop.Insert(ctx)
}

// Delete deletes the ComponentOrderProcess from the database.
func (cop *ComponentOrderProcess) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if cop._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetComponentOrderProcessTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, cop.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, cop.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, cop.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	cop._deleted = true

	return nil
}

// ComponentOrderProcessByID retrieves a row from 'aypcddg.component_order_process' as a ComponentOrderProcess.
//
// Generated from index 'component_order_process_id_pkey'.
func ComponentOrderProcessByID(ctx context.Context, id uint, key ...interface{}) (*ComponentOrderProcess, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetComponentOrderProcessTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, orderid, event, do_status, access_url, remark, created_at, updated_at ` +
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
	cop := ComponentOrderProcess{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&cop.ID, &cop.Orderid, &cop.Event, &cop.DoStatus, &cop.AccessURL, &cop.Remark, &cop.CreatedAt, &cop.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&cop.ID, &cop.Orderid, &cop.Event, &cop.DoStatus, &cop.AccessURL, &cop.Remark, &cop.CreatedAt, &cop.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &cop, nil
}