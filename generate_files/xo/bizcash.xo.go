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
	"time"

	"go.uber.org/zap"
)

// BizCash represents a row from 'aypcddg.biz_cash'.
type BizCash struct {
	Sid             int             `json:"sid"`              // sid
	Fid             int             `json:"fid"`              // fid
	CashLeft        sql.NullFloat64 `json:"cash_left"`        // cash_left
	TotalCommission sql.NullFloat64 `json:"total_commission"` // total_commission
	TotalCash       sql.NullFloat64 `json:"total_cash"`       // total_cash
	UpdateAt        time.Time       `json:"update_at"`        // update_at
	TotalWithdraw   sql.NullFloat64 `json:"total_withdraw"`   // total_withdraw
	ProcessWithdraw sql.NullFloat64 `json:"process_withdraw"` // process_withdraw

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BizCash exists in the database.
func (bc *BizCash) Exists() bool { //biz_cash
	return bc._exists
}

// Deleted provides information if the BizCash has been deleted from the database.
func (bc *BizCash) Deleted() bool {
	return bc._deleted
}

// Get table name
func GetBizCashTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "biz_cash", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the BizCash to the database.
func (bc *BizCash) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if bc._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBizCashTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`sid, fid, cash_left, total_commission, total_cash, update_at, total_withdraw, process_withdraw` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bc.Sid, bc.Fid, bc.CashLeft, bc.TotalCommission, bc.TotalCash, bc.UpdateAt, bc.TotalWithdraw, bc.ProcessWithdraw)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, bc.Sid, bc.Fid, bc.CashLeft, bc.TotalCommission, bc.TotalCash, bc.UpdateAt, bc.TotalWithdraw, bc.ProcessWithdraw)
	} else {
		res, err = dbConn.Exec(sqlstr, bc.Sid, bc.Fid, bc.CashLeft, bc.TotalCommission, bc.TotalCash, bc.UpdateAt, bc.TotalWithdraw, bc.ProcessWithdraw)
	}

	if err != nil {
		return err
	}

	// set existence
	bc._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	bc.Fid = int(id)
	bc._exists = true

	return nil
}

// Update updates the BizCash in the database.
func (bc *BizCash) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if bc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBizCashTableName(key...)
	if err != nil {
		return err
	}

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`cash_left = ?, total_commission = ?, total_cash = ?, update_at = ?, total_withdraw = ?, process_withdraw = ?` +
		` WHERE sid = ? AND fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bc.CashLeft, bc.TotalCommission, bc.TotalCash, bc.UpdateAt, bc.TotalWithdraw, bc.ProcessWithdraw, bc.Sid, bc.Fid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, bc.CashLeft, bc.TotalCommission, bc.TotalCash, bc.UpdateAt, bc.TotalWithdraw, bc.ProcessWithdraw, bc.Sid, bc.Fid)
	} else {
		_, err = dbConn.Exec(sqlstr, bc.CashLeft, bc.TotalCommission, bc.TotalCash, bc.UpdateAt, bc.TotalWithdraw, bc.ProcessWithdraw, bc.Sid, bc.Fid)
	}
	return err
}

// Save saves the BizCash to the database.
func (bc *BizCash) Save(ctx context.Context) error {
	if bc.Exists() {
		return bc.Update(ctx)
	}

	return bc.Insert(ctx)
}

// Delete deletes the BizCash from the database.
func (bc *BizCash) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if bc._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBizCashTableName(key...)
	if err != nil {
		return err
	}
	//2

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bc.Fid)))

	if tx != nil {
		_, err = tx.Exec(sqlstr, bc.Fid)
	} else {
		_, err = dbConn.Exec(sqlstr, bc.Fid)
	}
	if err != nil {
		return err
	}

	// set deleted
	bc._deleted = true

	return nil
}

// BizCashByFid retrieves a row from 'aypcddg.biz_cash' as a BizCash.
//
// Generated from index 'biz_cash_fid_pkey'.
func BizCashByFid(ctx context.Context, fid int, key ...interface{}) (*BizCash, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBizCashTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`sid, fid, cash_left, total_commission, total_cash, update_at, total_withdraw, process_withdraw ` +
		`FROM ` + tableName +
		` WHERE fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	bc := BizCash{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fid).Scan(&bc.Sid, &bc.Fid, &bc.CashLeft, &bc.TotalCommission, &bc.TotalCash, &bc.UpdateAt, &bc.TotalWithdraw, &bc.ProcessWithdraw)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fid).Scan(&bc.Sid, &bc.Fid, &bc.CashLeft, &bc.TotalCommission, &bc.TotalCash, &bc.UpdateAt, &bc.TotalWithdraw, &bc.ProcessWithdraw)
		if err != nil {
			return nil, err
		}
	}

	return &bc, nil
}

// BizCashBySidFid retrieves a row from 'aypcddg.biz_cash' as a BizCash.
//
// Generated from index 'fidsid'.
func BizCashBySidFid(ctx context.Context, sid int, fid int, key ...interface{}) (*BizCash, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBizCashTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`sid, fid, cash_left, total_commission, total_cash, update_at, total_withdraw, process_withdraw ` +
		`FROM ` + tableName +
		` WHERE sid = ? AND fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, sid, fid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	bc := BizCash{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, sid, fid).Scan(&bc.Sid, &bc.Fid, &bc.CashLeft, &bc.TotalCommission, &bc.TotalCash, &bc.UpdateAt, &bc.TotalWithdraw, &bc.ProcessWithdraw)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, sid, fid).Scan(&bc.Sid, &bc.Fid, &bc.CashLeft, &bc.TotalCommission, &bc.TotalCash, &bc.UpdateAt, &bc.TotalWithdraw, &bc.ProcessWithdraw)
		if err != nil {
			return nil, err
		}
	}

	return &bc, nil
}
