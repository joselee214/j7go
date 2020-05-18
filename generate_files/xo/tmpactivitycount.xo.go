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

// TmpActivityCount represents a row from 'aypcddg.tmp_activity_count'.
type TmpActivityCount struct {
	ID       int64          `json:"id"`        // id
	ReferURI sql.NullString `json:"refer_uri"` // refer_uri
	Time     int            `json:"time"`      // time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TmpActivityCount exists in the database.
func (tac *TmpActivityCount) Exists() bool { //tmp_activity_count
	return tac._exists
}

// Deleted provides information if the TmpActivityCount has been deleted from the database.
func (tac *TmpActivityCount) Deleted() bool {
	return tac._deleted
}

// Get table name
func GetTmpActivityCountTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "tmp_activity_count", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the TmpActivityCount to the database.
func (tac *TmpActivityCount) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if tac._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTmpActivityCountTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`refer_uri, time` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, tac.ReferURI, tac.Time)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, tac.ReferURI, tac.Time)
	} else {
		res, err = dbConn.Exec(sqlstr, tac.ReferURI, tac.Time)
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
	tac.ID = int64(id)
	tac._exists = true

	return nil
}

// Update updates the TmpActivityCount in the database.
func (tac *TmpActivityCount) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if tac._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTmpActivityCountTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`refer_uri = ?, time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, tac.ReferURI, tac.Time, tac.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, tac.ReferURI, tac.Time, tac.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, tac.ReferURI, tac.Time, tac.ID)
	}
	return err
}

// Save saves the TmpActivityCount to the database.
func (tac *TmpActivityCount) Save(ctx context.Context) error {
	if tac.Exists() {
		return tac.Update(ctx)
	}

	return tac.Insert(ctx)
}

// Delete deletes the TmpActivityCount from the database.
func (tac *TmpActivityCount) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if tac._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTmpActivityCountTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, tac.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, tac.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, tac.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	tac._deleted = true

	return nil
}

// TmpActivityCountByID retrieves a row from 'aypcddg.tmp_activity_count' as a TmpActivityCount.
//
// Generated from index 'tmp_activity_count_id_pkey'.
func TmpActivityCountByID(ctx context.Context, id int64, key ...interface{}) (*TmpActivityCount, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetTmpActivityCountTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refer_uri, time ` +
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
	tac := TmpActivityCount{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&tac.ID, &tac.ReferURI, &tac.Time)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&tac.ID, &tac.ReferURI, &tac.Time)
		if err != nil {
			return nil, err
		}
	}

	return &tac, nil
}