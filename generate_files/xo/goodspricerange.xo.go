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

// GoodsPriceRange represents a row from 'aypcddg.goods_price_range'.
type GoodsPriceRange struct {
	Gid      uint64          `json:"gid"`       // gid
	MinPrice sql.NullFloat64 `json:"min_price"` // min_price
	MaxPrice sql.NullFloat64 `json:"max_price"` // max_price

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GoodsPriceRange exists in the database.
func (gpr *GoodsPriceRange) Exists() bool { //goods_price_range
	return gpr._exists
}

// Deleted provides information if the GoodsPriceRange has been deleted from the database.
func (gpr *GoodsPriceRange) Deleted() bool {
	return gpr._deleted
}

// Get table name
func GetGoodsPriceRangeTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "goods_price_range", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the GoodsPriceRange to the database.
func (gpr *GoodsPriceRange) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if gpr._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsPriceRangeTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`gid, min_price, max_price` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gpr.Gid, gpr.MinPrice, gpr.MaxPrice)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, gpr.Gid, gpr.MinPrice, gpr.MaxPrice)
	} else {
		res, err = dbConn.Exec(sqlstr, gpr.Gid, gpr.MinPrice, gpr.MaxPrice)
	}

	if err != nil {
		return err
	}

	// set existence
	gpr._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	gpr.Gid = uint64(id)
	gpr._exists = true

	return nil
}

// Update updates the GoodsPriceRange in the database.
func (gpr *GoodsPriceRange) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gpr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsPriceRangeTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`min_price = ?, max_price = ?` +
		` WHERE gid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gpr.MinPrice, gpr.MaxPrice, gpr.Gid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gpr.MinPrice, gpr.MaxPrice, gpr.Gid)
	} else {
		_, err = dbConn.Exec(sqlstr, gpr.MinPrice, gpr.MaxPrice, gpr.Gid)
	}
	return err
}

// Save saves the GoodsPriceRange to the database.
func (gpr *GoodsPriceRange) Save(ctx context.Context) error {
	if gpr.Exists() {
		return gpr.Update(ctx)
	}

	return gpr.Insert(ctx)
}

// Delete deletes the GoodsPriceRange from the database.
func (gpr *GoodsPriceRange) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gpr._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsPriceRangeTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE gid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gpr.Gid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gpr.Gid)
	} else {
		_, err = dbConn.Exec(sqlstr, gpr.Gid)
	}

	if err != nil {
		return err
	}

	// set deleted
	gpr._deleted = true

	return nil
}

// GoodsPriceRangeByGid retrieves a row from 'aypcddg.goods_price_range' as a GoodsPriceRange.
//
// Generated from index 'goods_price_range_gid_pkey'.
func GoodsPriceRangeByGid(ctx context.Context, gid uint64, key ...interface{}) (*GoodsPriceRange, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsPriceRangeTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gid, min_price, max_price ` +
		`FROM ` + tableName +
		` WHERE gid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	gpr := GoodsPriceRange{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, gid).Scan(&gpr.Gid, &gpr.MinPrice, &gpr.MaxPrice)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, gid).Scan(&gpr.Gid, &gpr.MinPrice, &gpr.MaxPrice)
		if err != nil {
			return nil, err
		}
	}

	return &gpr, nil
}
