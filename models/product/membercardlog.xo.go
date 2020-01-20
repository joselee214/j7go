// Package product contains the types for schema 'saas'.
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

// MemberCardLog represents a row from 'saas.member_card_log'.
type MemberCardLog struct {
	ID           uint   `json:"id"`            // id
	TraceID      string `json:"trace_id"`      // trace_id
	ShopID       uint   `json:"shop_id"`       // shop_id
	BrandID      uint   `json:"brand_id"`      // brand_id
	CardID       uint   `json:"card_id"`       // card_id
	OperatorID   uint   `json:"operator_id"`   // operator_id
	OperateTable int8   `json:"operate_table"` // operate_table
	OperatorType int8   `json:"operator_type"` // operator_type
	ContentOld   string `json:"content_old"`   // content_old
	ContentNew   string `json:"content_new"`   // content_new
	IsDel        int8   `json:"is_del"`        // is_del
	CreatedTime  uint   `json:"created_time"`  // created_time
	UpdatedTime  uint   `json:"updated_time"`  // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MemberCardLog exists in the database.
func (mcl *MemberCardLog) Exists() bool { //member_card_log
	return mcl._exists
}

// Deleted provides information if the MemberCardLog has been deleted from the database.
func (mcl *MemberCardLog) Deleted() bool {
	return mcl._deleted
}

// Get table name
func getMemberCardLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "member_card_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the MemberCardLog to the database.
func (mcl *MemberCardLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if mcl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := getMemberCardLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`trace_id, shop_id, brand_id, card_id, operator_id, operate_table, operator_type, content_old, content_new, is_del, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mcl.TraceID, mcl.ShopID, mcl.BrandID, mcl.CardID, mcl.OperatorID, mcl.OperateTable, mcl.OperatorType, mcl.ContentOld, mcl.ContentNew, mcl.IsDel, mcl.CreatedTime, mcl.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, mcl.TraceID, mcl.ShopID, mcl.BrandID, mcl.CardID, mcl.OperatorID, mcl.OperateTable, mcl.OperatorType, mcl.ContentOld, mcl.ContentNew, mcl.IsDel, mcl.CreatedTime, mcl.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, mcl.TraceID, mcl.ShopID, mcl.BrandID, mcl.CardID, mcl.OperatorID, mcl.OperateTable, mcl.OperatorType, mcl.ContentOld, mcl.ContentNew, mcl.IsDel, mcl.CreatedTime, mcl.UpdatedTime)
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	mcl.ID = uint(id)
	mcl._exists = true

	return nil
}

// Update updates the MemberCardLog in the database.
func (mcl *MemberCardLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if mcl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := getMemberCardLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`trace_id = ?, shop_id = ?, brand_id = ?, card_id = ?, operator_id = ?, operate_table = ?, operator_type = ?, content_old = ?, content_new = ?, is_del = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mcl.TraceID, mcl.ShopID, mcl.BrandID, mcl.CardID, mcl.OperatorID, mcl.OperateTable, mcl.OperatorType, mcl.ContentOld, mcl.ContentNew, mcl.IsDel, mcl.CreatedTime, mcl.UpdatedTime, mcl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, mcl.TraceID, mcl.ShopID, mcl.BrandID, mcl.CardID, mcl.OperatorID, mcl.OperateTable, mcl.OperatorType, mcl.ContentOld, mcl.ContentNew, mcl.IsDel, mcl.CreatedTime, mcl.UpdatedTime, mcl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, mcl.TraceID, mcl.ShopID, mcl.BrandID, mcl.CardID, mcl.OperatorID, mcl.OperateTable, mcl.OperatorType, mcl.ContentOld, mcl.ContentNew, mcl.IsDel, mcl.CreatedTime, mcl.UpdatedTime, mcl.ID)
	}
	return err
}

// Save saves the MemberCardLog to the database.
func (mcl *MemberCardLog) Save(ctx context.Context) error {
	if mcl.Exists() {
		return mcl.Update(ctx)
	}

	return mcl.Insert(ctx)
}

// Delete deletes the MemberCardLog from the database.
func (mcl *MemberCardLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if mcl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := getMemberCardLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mcl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, mcl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, mcl.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	mcl._deleted = true

	return nil
}

// MemberCardLogByID retrieves a row from 'saas.member_card_log' as a MemberCardLog.
//
// Generated from index 'member_card_log_id_pkey'.
func MemberCardLogByID(ctx context.Context, id uint, key ...interface{}) (*MemberCardLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := getMemberCardLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, trace_id, shop_id, brand_id, card_id, operator_id, operate_table, operator_type, content_old, content_new, is_del, created_time, updated_time ` +
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
	mcl := MemberCardLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&mcl.ID, &mcl.TraceID, &mcl.ShopID, &mcl.BrandID, &mcl.CardID, &mcl.OperatorID, &mcl.OperateTable, &mcl.OperatorType, &mcl.ContentOld, &mcl.ContentNew, &mcl.IsDel, &mcl.CreatedTime, &mcl.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&mcl.ID, &mcl.TraceID, &mcl.ShopID, &mcl.BrandID, &mcl.CardID, &mcl.OperatorID, &mcl.OperateTable, &mcl.OperatorType, &mcl.ContentOld, &mcl.ContentNew, &mcl.IsDel, &mcl.CreatedTime, &mcl.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &mcl, nil
}
