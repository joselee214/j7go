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

// LogisticLog represents a row from 'aypcddg.logistic_log'.
type LogisticLog struct {
	Lid      uint           `json:"lid"`      // lid
	Oid      uint64         `json:"oid"`      // oid
	Aid      int8           `json:"aid"`      // aid
	Operator sql.NullString `json:"operator"` // operator
	Message  sql.NullString `json:"message"`  // message
	Time     uint           `json:"time"`     // time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the LogisticLog exists in the database.
func (ll *LogisticLog) Exists() bool { //logistic_log
	return ll._exists
}

// Deleted provides information if the LogisticLog has been deleted from the database.
func (ll *LogisticLog) Deleted() bool {
	return ll._deleted
}

// Get table name
func GetLogisticLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "logistic_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the LogisticLog to the database.
func (ll *LogisticLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ll._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetLogisticLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`oid, aid, operator, message, time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ll.Oid, ll.Aid, ll.Operator, ll.Message, ll.Time)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ll.Oid, ll.Aid, ll.Operator, ll.Message, ll.Time)
	} else {
		res, err = dbConn.Exec(sqlstr, ll.Oid, ll.Aid, ll.Operator, ll.Message, ll.Time)
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
	ll.Lid = uint(id)
	ll._exists = true

	return nil
}

// Update updates the LogisticLog in the database.
func (ll *LogisticLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ll._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetLogisticLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`oid = ?, aid = ?, operator = ?, message = ?, time = ?` +
		` WHERE lid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ll.Oid, ll.Aid, ll.Operator, ll.Message, ll.Time, ll.Lid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ll.Oid, ll.Aid, ll.Operator, ll.Message, ll.Time, ll.Lid)
	} else {
		_, err = dbConn.Exec(sqlstr, ll.Oid, ll.Aid, ll.Operator, ll.Message, ll.Time, ll.Lid)
	}
	return err
}

// Save saves the LogisticLog to the database.
func (ll *LogisticLog) Save(ctx context.Context) error {
	if ll.Exists() {
		return ll.Update(ctx)
	}

	return ll.Insert(ctx)
}

// Delete deletes the LogisticLog from the database.
func (ll *LogisticLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ll._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetLogisticLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE lid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ll.Lid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ll.Lid)
	} else {
		_, err = dbConn.Exec(sqlstr, ll.Lid)
	}

	if err != nil {
		return err
	}

	// set deleted
	ll._deleted = true

	return nil
}

// LogisticLogByLid retrieves a row from 'aypcddg.logistic_log' as a LogisticLog.
//
// Generated from index 'logistic_log_lid_pkey'.
func LogisticLogByLid(ctx context.Context, lid uint, key ...interface{}) (*LogisticLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetLogisticLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`lid, oid, aid, operator, message, time ` +
		`FROM ` + tableName +
		` WHERE lid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, lid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ll := LogisticLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, lid).Scan(&ll.Lid, &ll.Oid, &ll.Aid, &ll.Operator, &ll.Message, &ll.Time)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, lid).Scan(&ll.Lid, &ll.Oid, &ll.Aid, &ll.Operator, &ll.Message, &ll.Time)
		if err != nil {
			return nil, err
		}
	}

	return &ll, nil
}