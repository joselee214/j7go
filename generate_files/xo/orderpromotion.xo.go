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

// OrderPromotion represents a row from 'aypcddg.order_promotions'.
type OrderPromotion struct {
	Opid     int             `json:"opid"`     // opid
	Orderid  sql.NullInt64   `json:"orderid"`  // orderid
	Type     sql.NullInt64   `json:"type"`     // type
	ID       sql.NullInt64   `json:"id"`       // id
	Discount sql.NullFloat64 `json:"discount"` // discount

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OrderPromotion exists in the database.
func (op *OrderPromotion) Exists() bool { //order_promotions
	return op._exists
}

// Deleted provides information if the OrderPromotion has been deleted from the database.
func (op *OrderPromotion) Deleted() bool {
	return op._deleted
}

// Get table name
func GetOrderPromotionTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "order_promotions", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the OrderPromotion to the database.
func (op *OrderPromotion) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if op._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderPromotionTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`orderid, type, id, discount` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, op.Orderid, op.Type, op.ID, op.Discount)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, op.Orderid, op.Type, op.ID, op.Discount)
	} else {
		res, err = dbConn.Exec(sqlstr, op.Orderid, op.Type, op.ID, op.Discount)
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
	op.Opid = int(id)
	op._exists = true

	return nil
}

// Update updates the OrderPromotion in the database.
func (op *OrderPromotion) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if op._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderPromotionTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`orderid = ?, type = ?, id = ?, discount = ?` +
		` WHERE opid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, op.Orderid, op.Type, op.ID, op.Discount, op.Opid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, op.Orderid, op.Type, op.ID, op.Discount, op.Opid)
	} else {
		_, err = dbConn.Exec(sqlstr, op.Orderid, op.Type, op.ID, op.Discount, op.Opid)
	}
	return err
}

// Save saves the OrderPromotion to the database.
func (op *OrderPromotion) Save(ctx context.Context) error {
	if op.Exists() {
		return op.Update(ctx)
	}

	return op.Insert(ctx)
}

// Delete deletes the OrderPromotion from the database.
func (op *OrderPromotion) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if op._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderPromotionTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE opid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, op.Opid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, op.Opid)
	} else {
		_, err = dbConn.Exec(sqlstr, op.Opid)
	}

	if err != nil {
		return err
	}

	// set deleted
	op._deleted = true

	return nil
}

// OrderPromotionByOpid retrieves a row from 'aypcddg.order_promotions' as a OrderPromotion.
//
// Generated from index 'order_promotions_opid_pkey'.
func OrderPromotionByOpid(ctx context.Context, opid int, key ...interface{}) (*OrderPromotion, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetOrderPromotionTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`opid, orderid, type, id, discount ` +
		`FROM ` + tableName +
		` WHERE opid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, opid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	op := OrderPromotion{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, opid).Scan(&op.Opid, &op.Orderid, &op.Type, &op.ID, &op.Discount)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, opid).Scan(&op.Opid, &op.Orderid, &op.Type, &op.ID, &op.Discount)
		if err != nil {
			return nil, err
		}
	}

	return &op, nil
}
