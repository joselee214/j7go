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

// BizCashLog represents a row from 'aypcddg.biz_cash_log'.
type BizCashLog struct {
	ID          int             `json:"id"`           // id
	Sid         int             `json:"sid"`          // sid
	Fid         sql.NullInt64   `json:"fid"`          // fid
	Fsid        sql.NullInt64   `json:"fsid"`         // fsid
	RelatedType sql.NullString  `json:"related_type"` // related_type
	RelatedID   sql.NullInt64   `json:"related_id"`   // related_id
	Action      sql.NullString  `json:"action"`       // action
	ActionCash  sql.NullFloat64 `json:"action_cash"`  // action_cash
	CashBefore  sql.NullFloat64 `json:"cash_before"`  // cash_before
	CashLeft    sql.NullFloat64 `json:"cash_left"`    // cash_left
	CreatedAt   mysql.NullTime  `json:"created_at"`   // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BizCashLog exists in the database.
func (bcl *BizCashLog) Exists() bool { //biz_cash_log
	return bcl._exists
}

// Deleted provides information if the BizCashLog has been deleted from the database.
func (bcl *BizCashLog) Deleted() bool {
	return bcl._deleted
}

// Get table name
func GetBizCashLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "biz_cash_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the BizCashLog to the database.
func (bcl *BizCashLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if bcl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBizCashLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`sid, fid, fsid, related_type, related_id, action, action_cash, cash_before, cash_left, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bcl.Sid, bcl.Fid, bcl.Fsid, bcl.RelatedType, bcl.RelatedID, bcl.Action, bcl.ActionCash, bcl.CashBefore, bcl.CashLeft, bcl.CreatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, bcl.Sid, bcl.Fid, bcl.Fsid, bcl.RelatedType, bcl.RelatedID, bcl.Action, bcl.ActionCash, bcl.CashBefore, bcl.CashLeft, bcl.CreatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, bcl.Sid, bcl.Fid, bcl.Fsid, bcl.RelatedType, bcl.RelatedID, bcl.Action, bcl.ActionCash, bcl.CashBefore, bcl.CashLeft, bcl.CreatedAt)
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
	bcl.ID = int(id)
	bcl._exists = true

	return nil
}

// Update updates the BizCashLog in the database.
func (bcl *BizCashLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if bcl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBizCashLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`sid = ?, fid = ?, fsid = ?, related_type = ?, related_id = ?, action = ?, action_cash = ?, cash_before = ?, cash_left = ?, created_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bcl.Sid, bcl.Fid, bcl.Fsid, bcl.RelatedType, bcl.RelatedID, bcl.Action, bcl.ActionCash, bcl.CashBefore, bcl.CashLeft, bcl.CreatedAt, bcl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, bcl.Sid, bcl.Fid, bcl.Fsid, bcl.RelatedType, bcl.RelatedID, bcl.Action, bcl.ActionCash, bcl.CashBefore, bcl.CashLeft, bcl.CreatedAt, bcl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, bcl.Sid, bcl.Fid, bcl.Fsid, bcl.RelatedType, bcl.RelatedID, bcl.Action, bcl.ActionCash, bcl.CashBefore, bcl.CashLeft, bcl.CreatedAt, bcl.ID)
	}
	return err
}

// Save saves the BizCashLog to the database.
func (bcl *BizCashLog) Save(ctx context.Context) error {
	if bcl.Exists() {
		return bcl.Update(ctx)
	}

	return bcl.Insert(ctx)
}

// Delete deletes the BizCashLog from the database.
func (bcl *BizCashLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if bcl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBizCashLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bcl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, bcl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, bcl.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	bcl._deleted = true

	return nil
}

// BizCashLogByID retrieves a row from 'aypcddg.biz_cash_log' as a BizCashLog.
//
// Generated from index 'biz_cash_log_id_pkey'.
func BizCashLogByID(ctx context.Context, id int, key ...interface{}) (*BizCashLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBizCashLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, sid, fid, fsid, related_type, related_id, action, action_cash, cash_before, cash_left, created_at ` +
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
	bcl := BizCashLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&bcl.ID, &bcl.Sid, &bcl.Fid, &bcl.Fsid, &bcl.RelatedType, &bcl.RelatedID, &bcl.Action, &bcl.ActionCash, &bcl.CashBefore, &bcl.CashLeft, &bcl.CreatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&bcl.ID, &bcl.Sid, &bcl.Fid, &bcl.Fsid, &bcl.RelatedType, &bcl.RelatedID, &bcl.Action, &bcl.ActionCash, &bcl.CashBefore, &bcl.CashLeft, &bcl.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &bcl, nil
}
