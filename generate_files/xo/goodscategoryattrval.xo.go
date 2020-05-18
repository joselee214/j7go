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

// GoodsCategoryAttrVal represents a row from 'aypcddg.goods_category_attr_val'.
type GoodsCategoryAttrVal struct {
	Gcavid uint64         `json:"gcavid"` // gcavid
	Gcaid  uint64         `json:"gcaid"`  // gcaid
	Fid    uint64         `json:"fid"`    // fid
	Status bool           `json:"status"` // status
	Val    sql.NullString `json:"val"`    // val
	Seq    uint           `json:"seq"`    // seq

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GoodsCategoryAttrVal exists in the database.
func (gcav *GoodsCategoryAttrVal) Exists() bool { //goods_category_attr_val
	return gcav._exists
}

// Deleted provides information if the GoodsCategoryAttrVal has been deleted from the database.
func (gcav *GoodsCategoryAttrVal) Deleted() bool {
	return gcav._deleted
}

// Get table name
func GetGoodsCategoryAttrValTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "goods_category_attr_val", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the GoodsCategoryAttrVal to the database.
func (gcav *GoodsCategoryAttrVal) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if gcav._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategoryAttrValTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`gcaid, fid, status, val, seq` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcav.Gcaid, gcav.Fid, gcav.Status, gcav.Val, gcav.Seq)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, gcav.Gcaid, gcav.Fid, gcav.Status, gcav.Val, gcav.Seq)
	} else {
		res, err = dbConn.Exec(sqlstr, gcav.Gcaid, gcav.Fid, gcav.Status, gcav.Val, gcav.Seq)
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
	gcav.Gcavid = uint64(id)
	gcav._exists = true

	return nil
}

// Update updates the GoodsCategoryAttrVal in the database.
func (gcav *GoodsCategoryAttrVal) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gcav._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategoryAttrValTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`gcaid = ?, fid = ?, status = ?, val = ?, seq = ?` +
		` WHERE gcavid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcav.Gcaid, gcav.Fid, gcav.Status, gcav.Val, gcav.Seq, gcav.Gcavid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gcav.Gcaid, gcav.Fid, gcav.Status, gcav.Val, gcav.Seq, gcav.Gcavid)
	} else {
		_, err = dbConn.Exec(sqlstr, gcav.Gcaid, gcav.Fid, gcav.Status, gcav.Val, gcav.Seq, gcav.Gcavid)
	}
	return err
}

// Save saves the GoodsCategoryAttrVal to the database.
func (gcav *GoodsCategoryAttrVal) Save(ctx context.Context) error {
	if gcav.Exists() {
		return gcav.Update(ctx)
	}

	return gcav.Insert(ctx)
}

// Delete deletes the GoodsCategoryAttrVal from the database.
func (gcav *GoodsCategoryAttrVal) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gcav._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategoryAttrValTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE gcavid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcav.Gcavid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gcav.Gcavid)
	} else {
		_, err = dbConn.Exec(sqlstr, gcav.Gcavid)
	}

	if err != nil {
		return err
	}

	// set deleted
	gcav._deleted = true

	return nil
}

// GoodsCategoryAttrValsByFid retrieves a row from 'aypcddg.goods_category_attr_val' as a GoodsCategoryAttrVal.
//
// Generated from index 'fid'.
func GoodsCategoryAttrValsByFid(ctx context.Context, fid uint64, key ...interface{}) ([]*GoodsCategoryAttrVal, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategoryAttrValTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcavid, gcaid, fid, status, val, seq ` +
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
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, fid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, fid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*GoodsCategoryAttrVal, 0)
	for queryData.Next() {
		gcav := GoodsCategoryAttrVal{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&gcav.Gcavid, &gcav.Gcaid, &gcav.Fid, &gcav.Status, &gcav.Val, &gcav.Seq)
		if err != nil {
			return nil, err
		}

		res = append(res, &gcav)
	}

	return res, nil
}

// GoodsCategoryAttrValsByGcaid retrieves a row from 'aypcddg.goods_category_attr_val' as a GoodsCategoryAttrVal.
//
// Generated from index 'gcaid'.
func GoodsCategoryAttrValsByGcaid(ctx context.Context, gcaid uint64, key ...interface{}) ([]*GoodsCategoryAttrVal, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategoryAttrValTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcavid, gcaid, fid, status, val, seq ` +
		`FROM ` + tableName +
		` WHERE gcaid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcaid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, gcaid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, gcaid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*GoodsCategoryAttrVal, 0)
	for queryData.Next() {
		gcav := GoodsCategoryAttrVal{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&gcav.Gcavid, &gcav.Gcaid, &gcav.Fid, &gcav.Status, &gcav.Val, &gcav.Seq)
		if err != nil {
			return nil, err
		}

		res = append(res, &gcav)
	}

	return res, nil
}

// GoodsCategoryAttrValByGcavid retrieves a row from 'aypcddg.goods_category_attr_val' as a GoodsCategoryAttrVal.
//
// Generated from index 'goods_category_attr_val_gcavid_pkey'.
func GoodsCategoryAttrValByGcavid(ctx context.Context, gcavid uint64, key ...interface{}) (*GoodsCategoryAttrVal, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategoryAttrValTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcavid, gcaid, fid, status, val, seq ` +
		`FROM ` + tableName +
		` WHERE gcavid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcavid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	gcav := GoodsCategoryAttrVal{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, gcavid).Scan(&gcav.Gcavid, &gcav.Gcaid, &gcav.Fid, &gcav.Status, &gcav.Val, &gcav.Seq)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, gcavid).Scan(&gcav.Gcavid, &gcav.Gcaid, &gcav.Fid, &gcav.Status, &gcav.Val, &gcav.Seq)
		if err != nil {
			return nil, err
		}
	}

	return &gcav, nil
}

// GoodsCategoryAttrValsBySeq retrieves a row from 'aypcddg.goods_category_attr_val' as a GoodsCategoryAttrVal.
//
// Generated from index 'seq'.
func GoodsCategoryAttrValsBySeq(ctx context.Context, seq uint, key ...interface{}) ([]*GoodsCategoryAttrVal, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategoryAttrValTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcavid, gcaid, fid, status, val, seq ` +
		`FROM ` + tableName +
		` WHERE seq = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, seq)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, seq)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, seq)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*GoodsCategoryAttrVal, 0)
	for queryData.Next() {
		gcav := GoodsCategoryAttrVal{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&gcav.Gcavid, &gcav.Gcaid, &gcav.Fid, &gcav.Status, &gcav.Val, &gcav.Seq)
		if err != nil {
			return nil, err
		}

		res = append(res, &gcav)
	}

	return res, nil
}