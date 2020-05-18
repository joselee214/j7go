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

// GoodsStockLog represents a row from 'aypcddg.goods_stock_log'.
type GoodsStockLog struct {
	ID        uint64         `json:"id"`         // id
	Gid       int            `json:"gid"`        // gid
	Stock     int            `json:"stock"`      // stock
	Type      int8           `json:"type"`       // type
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GoodsStockLog exists in the database.
func (gsl *GoodsStockLog) Exists() bool { //goods_stock_log
	return gsl._exists
}

// Deleted provides information if the GoodsStockLog has been deleted from the database.
func (gsl *GoodsStockLog) Deleted() bool {
	return gsl._deleted
}

// Get table name
func GetGoodsStockLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "goods_stock_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the GoodsStockLog to the database.
func (gsl *GoodsStockLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if gsl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsStockLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`gid, stock, type, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gsl.Gid, gsl.Stock, gsl.Type, gsl.CreatedAt, gsl.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, gsl.Gid, gsl.Stock, gsl.Type, gsl.CreatedAt, gsl.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, gsl.Gid, gsl.Stock, gsl.Type, gsl.CreatedAt, gsl.UpdatedAt)
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
	gsl.ID = uint64(id)
	gsl._exists = true

	return nil
}

// Update updates the GoodsStockLog in the database.
func (gsl *GoodsStockLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gsl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsStockLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`gid = ?, stock = ?, type = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gsl.Gid, gsl.Stock, gsl.Type, gsl.CreatedAt, gsl.UpdatedAt, gsl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gsl.Gid, gsl.Stock, gsl.Type, gsl.CreatedAt, gsl.UpdatedAt, gsl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, gsl.Gid, gsl.Stock, gsl.Type, gsl.CreatedAt, gsl.UpdatedAt, gsl.ID)
	}
	return err
}

// Save saves the GoodsStockLog to the database.
func (gsl *GoodsStockLog) Save(ctx context.Context) error {
	if gsl.Exists() {
		return gsl.Update(ctx)
	}

	return gsl.Insert(ctx)
}

// Delete deletes the GoodsStockLog from the database.
func (gsl *GoodsStockLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gsl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsStockLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gsl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gsl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, gsl.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	gsl._deleted = true

	return nil
}

// GoodsStockLogsByGid retrieves a row from 'aypcddg.goods_stock_log' as a GoodsStockLog.
//
// Generated from index 'goods_stock_log_gid_index'.
func GoodsStockLogsByGid(ctx context.Context, gid int, key ...interface{}) ([]*GoodsStockLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsStockLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, gid, stock, type, created_at, updated_at ` +
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
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, gid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, gid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*GoodsStockLog, 0)
	for queryData.Next() {
		gsl := GoodsStockLog{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&gsl.ID, &gsl.Gid, &gsl.Stock, &gsl.Type, &gsl.CreatedAt, &gsl.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &gsl)
	}

	return res, nil
}

// GoodsStockLogByID retrieves a row from 'aypcddg.goods_stock_log' as a GoodsStockLog.
//
// Generated from index 'goods_stock_log_id_pkey'.
func GoodsStockLogByID(ctx context.Context, id uint64, key ...interface{}) (*GoodsStockLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsStockLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, gid, stock, type, created_at, updated_at ` +
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
	gsl := GoodsStockLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&gsl.ID, &gsl.Gid, &gsl.Stock, &gsl.Type, &gsl.CreatedAt, &gsl.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&gsl.ID, &gsl.Gid, &gsl.Stock, &gsl.Type, &gsl.CreatedAt, &gsl.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &gsl, nil
}