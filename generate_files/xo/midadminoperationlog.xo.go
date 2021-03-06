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

// MidAdminOperationLog represents a row from 'aypcddg.mid_admin_operation_log'.
type MidAdminOperationLog struct {
	ID        uint           `json:"id"`         // id
	UserID    int            `json:"user_id"`    // user_id
	Path      string         `json:"path"`       // path
	Method    string         `json:"method"`     // method
	IP        string         `json:"ip"`         // ip
	Input     string         `json:"input"`      // input
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MidAdminOperationLog exists in the database.
func (maol *MidAdminOperationLog) Exists() bool { //mid_admin_operation_log
	return maol._exists
}

// Deleted provides information if the MidAdminOperationLog has been deleted from the database.
func (maol *MidAdminOperationLog) Deleted() bool {
	return maol._deleted
}

// Get table name
func GetMidAdminOperationLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "mid_admin_operation_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the MidAdminOperationLog to the database.
func (maol *MidAdminOperationLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if maol._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminOperationLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`user_id, path, method, ip, input, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, maol.UserID, maol.Path, maol.Method, maol.IP, maol.Input, maol.CreatedAt, maol.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, maol.UserID, maol.Path, maol.Method, maol.IP, maol.Input, maol.CreatedAt, maol.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, maol.UserID, maol.Path, maol.Method, maol.IP, maol.Input, maol.CreatedAt, maol.UpdatedAt)
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
	maol.ID = uint(id)
	maol._exists = true

	return nil
}

// Update updates the MidAdminOperationLog in the database.
func (maol *MidAdminOperationLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if maol._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminOperationLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`user_id = ?, path = ?, method = ?, ip = ?, input = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, maol.UserID, maol.Path, maol.Method, maol.IP, maol.Input, maol.CreatedAt, maol.UpdatedAt, maol.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, maol.UserID, maol.Path, maol.Method, maol.IP, maol.Input, maol.CreatedAt, maol.UpdatedAt, maol.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, maol.UserID, maol.Path, maol.Method, maol.IP, maol.Input, maol.CreatedAt, maol.UpdatedAt, maol.ID)
	}
	return err
}

// Save saves the MidAdminOperationLog to the database.
func (maol *MidAdminOperationLog) Save(ctx context.Context) error {
	if maol.Exists() {
		return maol.Update(ctx)
	}

	return maol.Insert(ctx)
}

// Delete deletes the MidAdminOperationLog from the database.
func (maol *MidAdminOperationLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if maol._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminOperationLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, maol.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, maol.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, maol.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	maol._deleted = true

	return nil
}

// MidAdminOperationLogByID retrieves a row from 'aypcddg.mid_admin_operation_log' as a MidAdminOperationLog.
//
// Generated from index 'mid_admin_operation_log_id_pkey'.
func MidAdminOperationLogByID(ctx context.Context, id uint, key ...interface{}) (*MidAdminOperationLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminOperationLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, user_id, path, method, ip, input, created_at, updated_at ` +
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
	maol := MidAdminOperationLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&maol.ID, &maol.UserID, &maol.Path, &maol.Method, &maol.IP, &maol.Input, &maol.CreatedAt, &maol.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&maol.ID, &maol.UserID, &maol.Path, &maol.Method, &maol.IP, &maol.Input, &maol.CreatedAt, &maol.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &maol, nil
}

// MidAdminOperationLogsByUserID retrieves a row from 'aypcddg.mid_admin_operation_log' as a MidAdminOperationLog.
//
// Generated from index 'mid_admin_operation_log_user_id_index'.
func MidAdminOperationLogsByUserID(ctx context.Context, userID int, key ...interface{}) ([]*MidAdminOperationLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminOperationLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, user_id, path, method, ip, input, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE user_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, userID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, userID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, userID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*MidAdminOperationLog, 0)
	for queryData.Next() {
		maol := MidAdminOperationLog{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&maol.ID, &maol.UserID, &maol.Path, &maol.Method, &maol.IP, &maol.Input, &maol.CreatedAt, &maol.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &maol)
	}

	return res, nil
}
