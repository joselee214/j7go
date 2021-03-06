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

// GoodsAttr represents a row from 'aypcddg.goods_attr'.
type GoodsAttr struct {
	Gid    uint64 `json:"gid"`    // gid
	Gcaid  uint   `json:"gcaid"`  // gcaid
	Gcavid uint   `json:"gcavid"` // gcavid
	Vals   string `json:"vals"`   // vals

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GoodsAttr exists in the database.
func (ga *GoodsAttr) Exists() bool { //goods_attr
	return ga._exists
}

// Deleted provides information if the GoodsAttr has been deleted from the database.
func (ga *GoodsAttr) Deleted() bool {
	return ga._deleted
}

// Get table name
func GetGoodsAttrTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "goods_attr", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the GoodsAttr to the database.
func (ga *GoodsAttr) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ga._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsAttrTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`gid, gcaid, gcavid, vals` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ga.Gid, ga.Gcaid, ga.Gcavid, ga.Vals)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, ga.Gid, ga.Gcaid, ga.Gcavid, ga.Vals)
	} else {
		res, err = dbConn.Exec(sqlstr, ga.Gid, ga.Gcaid, ga.Gcavid, ga.Vals)
	}

	if err != nil {
		return err
	}

	// set existence
	ga._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ga.Vals = string(id)
	ga._exists = true

	return nil
}

// Update updates the GoodsAttr in the database.
func (ga *GoodsAttr) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ga._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsAttrTableName(key...)
	if err != nil {
		return err
	}

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`gcavid = ?` +
		` WHERE gid = ? AND gcaid = ? AND vals = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ga.Gcavid, ga.Gid, ga.Gcaid, ga.Vals)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ga.Gcavid, ga.Gid, ga.Gcaid, ga.Vals)
	} else {
		_, err = dbConn.Exec(sqlstr, ga.Gcavid, ga.Gid, ga.Gcaid, ga.Vals)
	}
	return err
}

// Save saves the GoodsAttr to the database.
func (ga *GoodsAttr) Save(ctx context.Context) error {
	if ga.Exists() {
		return ga.Update(ctx)
	}

	return ga.Insert(ctx)
}

// Delete deletes the GoodsAttr from the database.
func (ga *GoodsAttr) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ga._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsAttrTableName(key...)
	if err != nil {
		return err
	}
	//3

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE vals = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ga.Vals)))

	if tx != nil {
		_, err = tx.Exec(sqlstr, ga.Vals)
	} else {
		_, err = dbConn.Exec(sqlstr, ga.Vals)
	}
	if err != nil {
		return err
	}

	// set deleted
	ga._deleted = true

	return nil
}

// GoodsAttrByVals retrieves a row from 'aypcddg.goods_attr' as a GoodsAttr.
//
// Generated from index 'goods_attr_vals_pkey'.
func GoodsAttrByVals(ctx context.Context, vals string, key ...interface{}) (*GoodsAttr, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsAttrTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gid, gcaid, gcavid, vals ` +
		`FROM ` + tableName +
		` WHERE vals = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, vals)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ga := GoodsAttr{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, vals).Scan(&ga.Gid, &ga.Gcaid, &ga.Gcavid, &ga.Vals)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, vals).Scan(&ga.Gid, &ga.Gcaid, &ga.Gcavid, &ga.Vals)
		if err != nil {
			return nil, err
		}
	}

	return &ga, nil
}
