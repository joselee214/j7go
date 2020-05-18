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

// OrderNotifyLog represents a row from 'aypcddg.order_notify_log'.
type OrderNotifyLog struct {
	ID           uint           `json:"id"`             // id
	Payid        sql.NullString `json:"payid"`          // payid
	ThirdOrderNo sql.NullString `json:"third_order_no"` // third_order_no
	OurResult    sql.NullString `json:"our_result"`     // our_result
	ThirdResult  sql.NullString `json:"third_result"`   // third_result
	CreatedAt    mysql.NullTime `json:"created_at"`     // created_at
	UpdatedAt    mysql.NullTime `json:"updated_at"`     // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OrderNotifyLog exists in the database.
func (onl *OrderNotifyLog) Exists() bool { //order_notify_log
	return onl._exists
}

// Deleted provides information if the OrderNotifyLog has been deleted from the database.
func (onl *OrderNotifyLog) Deleted() bool {
	return onl._deleted
}

// Get table name
func GetOrderNotifyLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "order_notify_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the OrderNotifyLog to the database.
func (onl *OrderNotifyLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if onl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderNotifyLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`payid, third_order_no, our_result, third_result, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, onl.Payid, onl.ThirdOrderNo, onl.OurResult, onl.ThirdResult, onl.CreatedAt, onl.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, onl.Payid, onl.ThirdOrderNo, onl.OurResult, onl.ThirdResult, onl.CreatedAt, onl.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, onl.Payid, onl.ThirdOrderNo, onl.OurResult, onl.ThirdResult, onl.CreatedAt, onl.UpdatedAt)
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
	onl.ID = uint(id)
	onl._exists = true

	return nil
}

// Update updates the OrderNotifyLog in the database.
func (onl *OrderNotifyLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if onl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderNotifyLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`payid = ?, third_order_no = ?, our_result = ?, third_result = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, onl.Payid, onl.ThirdOrderNo, onl.OurResult, onl.ThirdResult, onl.CreatedAt, onl.UpdatedAt, onl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, onl.Payid, onl.ThirdOrderNo, onl.OurResult, onl.ThirdResult, onl.CreatedAt, onl.UpdatedAt, onl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, onl.Payid, onl.ThirdOrderNo, onl.OurResult, onl.ThirdResult, onl.CreatedAt, onl.UpdatedAt, onl.ID)
	}
	return err
}

// Save saves the OrderNotifyLog to the database.
func (onl *OrderNotifyLog) Save(ctx context.Context) error {
	if onl.Exists() {
		return onl.Update(ctx)
	}

	return onl.Insert(ctx)
}

// Delete deletes the OrderNotifyLog from the database.
func (onl *OrderNotifyLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if onl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderNotifyLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, onl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, onl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, onl.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	onl._deleted = true

	return nil
}

// OrderNotifyLogByID retrieves a row from 'aypcddg.order_notify_log' as a OrderNotifyLog.
//
// Generated from index 'order_notify_log_id_pkey'.
func OrderNotifyLogByID(ctx context.Context, id uint, key ...interface{}) (*OrderNotifyLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetOrderNotifyLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, payid, third_order_no, our_result, third_result, created_at, updated_at ` +
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
	onl := OrderNotifyLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&onl.ID, &onl.Payid, &onl.ThirdOrderNo, &onl.OurResult, &onl.ThirdResult, &onl.CreatedAt, &onl.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&onl.ID, &onl.Payid, &onl.ThirdOrderNo, &onl.OurResult, &onl.ThirdResult, &onl.CreatedAt, &onl.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &onl, nil
}
