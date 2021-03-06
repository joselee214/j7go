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

// GoodsCategorySpec represents a row from 'aypcddg.goods_category_spec'.
type GoodsCategorySpec struct {
	Gcsid uint           `json:"gcsid"` // gcsid
	Gcid  uint           `json:"gcid"`  // gcid
	Title sql.NullString `json:"title"` // title
	Seq   int16          `json:"seq"`   // seq
	Tp    sql.NullString `json:"tp"`    // tp
	Vals  sql.NullString `json:"vals"`  // vals

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GoodsCategorySpec exists in the database.
func (gcs *GoodsCategorySpec) Exists() bool { //goods_category_spec
	return gcs._exists
}

// Deleted provides information if the GoodsCategorySpec has been deleted from the database.
func (gcs *GoodsCategorySpec) Deleted() bool {
	return gcs._deleted
}

// Get table name
func GetGoodsCategorySpecTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "goods_category_spec", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the GoodsCategorySpec to the database.
func (gcs *GoodsCategorySpec) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if gcs._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategorySpecTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`gcid, title, seq, tp, vals` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcs.Gcid, gcs.Title, gcs.Seq, gcs.Tp, gcs.Vals)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, gcs.Gcid, gcs.Title, gcs.Seq, gcs.Tp, gcs.Vals)
	} else {
		res, err = dbConn.Exec(sqlstr, gcs.Gcid, gcs.Title, gcs.Seq, gcs.Tp, gcs.Vals)
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
	gcs.Gcsid = uint(id)
	gcs._exists = true

	return nil
}

// Update updates the GoodsCategorySpec in the database.
func (gcs *GoodsCategorySpec) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gcs._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategorySpecTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`gcid = ?, title = ?, seq = ?, tp = ?, vals = ?` +
		` WHERE gcsid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcs.Gcid, gcs.Title, gcs.Seq, gcs.Tp, gcs.Vals, gcs.Gcsid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gcs.Gcid, gcs.Title, gcs.Seq, gcs.Tp, gcs.Vals, gcs.Gcsid)
	} else {
		_, err = dbConn.Exec(sqlstr, gcs.Gcid, gcs.Title, gcs.Seq, gcs.Tp, gcs.Vals, gcs.Gcsid)
	}
	return err
}

// Save saves the GoodsCategorySpec to the database.
func (gcs *GoodsCategorySpec) Save(ctx context.Context) error {
	if gcs.Exists() {
		return gcs.Update(ctx)
	}

	return gcs.Insert(ctx)
}

// Delete deletes the GoodsCategorySpec from the database.
func (gcs *GoodsCategorySpec) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gcs._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategorySpecTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE gcsid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcs.Gcsid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gcs.Gcsid)
	} else {
		_, err = dbConn.Exec(sqlstr, gcs.Gcsid)
	}

	if err != nil {
		return err
	}

	// set deleted
	gcs._deleted = true

	return nil
}

// GoodsCategorySpecsByGcid retrieves a row from 'aypcddg.goods_category_spec' as a GoodsCategorySpec.
//
// Generated from index 'gcid'.
func GoodsCategorySpecsByGcid(ctx context.Context, gcid uint, key ...interface{}) ([]*GoodsCategorySpec, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategorySpecTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcsid, gcid, title, seq, tp, vals ` +
		`FROM ` + tableName +
		` WHERE gcid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, gcid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, gcid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*GoodsCategorySpec, 0)
	for queryData.Next() {
		gcs := GoodsCategorySpec{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&gcs.Gcsid, &gcs.Gcid, &gcs.Title, &gcs.Seq, &gcs.Tp, &gcs.Vals)
		if err != nil {
			return nil, err
		}

		res = append(res, &gcs)
	}

	return res, nil
}

// GoodsCategorySpecByGcsid retrieves a row from 'aypcddg.goods_category_spec' as a GoodsCategorySpec.
//
// Generated from index 'goods_category_spec_gcsid_pkey'.
func GoodsCategorySpecByGcsid(ctx context.Context, gcsid uint, key ...interface{}) (*GoodsCategorySpec, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategorySpecTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcsid, gcid, title, seq, tp, vals ` +
		`FROM ` + tableName +
		` WHERE gcsid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcsid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	gcs := GoodsCategorySpec{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, gcsid).Scan(&gcs.Gcsid, &gcs.Gcid, &gcs.Title, &gcs.Seq, &gcs.Tp, &gcs.Vals)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, gcsid).Scan(&gcs.Gcsid, &gcs.Gcid, &gcs.Title, &gcs.Seq, &gcs.Tp, &gcs.Vals)
		if err != nil {
			return nil, err
		}
	}

	return &gcs, nil
}

// GoodsCategorySpecsBySeq retrieves a row from 'aypcddg.goods_category_spec' as a GoodsCategorySpec.
//
// Generated from index 'seq'.
func GoodsCategorySpecsBySeq(ctx context.Context, seq int16, key ...interface{}) ([]*GoodsCategorySpec, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategorySpecTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcsid, gcid, title, seq, tp, vals ` +
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
	res := make([]*GoodsCategorySpec, 0)
	for queryData.Next() {
		gcs := GoodsCategorySpec{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&gcs.Gcsid, &gcs.Gcid, &gcs.Title, &gcs.Seq, &gcs.Tp, &gcs.Vals)
		if err != nil {
			return nil, err
		}

		res = append(res, &gcs)
	}

	return res, nil
}
