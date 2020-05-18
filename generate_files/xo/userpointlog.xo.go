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

// UserPointLog represents a row from 'aypcddg.user_point_log'.
type UserPointLog struct {
	Uplid      uint64         `json:"uplid"`       // uplid
	UID        uint           `json:"uid"`         // uid
	ChangeType sql.NullString `json:"change_type"` // change_type
	RefID      uint64         `json:"ref_id"`      // ref_id
	Num        int            `json:"num"`         // num
	Note       sql.NullString `json:"note"`        // note
	Created    int            `json:"created"`     // created

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserPointLog exists in the database.
func (upl *UserPointLog) Exists() bool { //user_point_log
	return upl._exists
}

// Deleted provides information if the UserPointLog has been deleted from the database.
func (upl *UserPointLog) Deleted() bool {
	return upl._deleted
}

// Get table name
func GetUserPointLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "user_point_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserPointLog to the database.
func (upl *UserPointLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if upl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserPointLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, change_type, ref_id, num, note, created` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, upl.UID, upl.ChangeType, upl.RefID, upl.Num, upl.Note, upl.Created)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, upl.UID, upl.ChangeType, upl.RefID, upl.Num, upl.Note, upl.Created)
	} else {
		res, err = dbConn.Exec(sqlstr, upl.UID, upl.ChangeType, upl.RefID, upl.Num, upl.Note, upl.Created)
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
	upl.Uplid = uint64(id)
	upl._exists = true

	return nil
}

// Update updates the UserPointLog in the database.
func (upl *UserPointLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if upl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserPointLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, change_type = ?, ref_id = ?, num = ?, note = ?, created = ?` +
		` WHERE uplid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, upl.UID, upl.ChangeType, upl.RefID, upl.Num, upl.Note, upl.Created, upl.Uplid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, upl.UID, upl.ChangeType, upl.RefID, upl.Num, upl.Note, upl.Created, upl.Uplid)
	} else {
		_, err = dbConn.Exec(sqlstr, upl.UID, upl.ChangeType, upl.RefID, upl.Num, upl.Note, upl.Created, upl.Uplid)
	}
	return err
}

// Save saves the UserPointLog to the database.
func (upl *UserPointLog) Save(ctx context.Context) error {
	if upl.Exists() {
		return upl.Update(ctx)
	}

	return upl.Insert(ctx)
}

// Delete deletes the UserPointLog from the database.
func (upl *UserPointLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if upl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserPointLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE uplid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, upl.Uplid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, upl.Uplid)
	} else {
		_, err = dbConn.Exec(sqlstr, upl.Uplid)
	}

	if err != nil {
		return err
	}

	// set deleted
	upl._deleted = true

	return nil
}

// UserPointLogsByChangeType retrieves a row from 'aypcddg.user_point_log' as a UserPointLog.
//
// Generated from index 'change_type'.
func UserPointLogsByChangeType(ctx context.Context, changeType sql.NullString, key ...interface{}) ([]*UserPointLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserPointLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uplid, uid, change_type, ref_id, num, note, created ` +
		`FROM ` + tableName +
		` WHERE change_type = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, changeType)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, changeType)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, changeType)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserPointLog, 0)
	for queryData.Next() {
		upl := UserPointLog{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&upl.Uplid, &upl.UID, &upl.ChangeType, &upl.RefID, &upl.Num, &upl.Note, &upl.Created)
		if err != nil {
			return nil, err
		}

		res = append(res, &upl)
	}

	return res, nil
}

// UserPointLogsByRefID retrieves a row from 'aypcddg.user_point_log' as a UserPointLog.
//
// Generated from index 'ref_id'.
func UserPointLogsByRefID(ctx context.Context, refID uint64, key ...interface{}) ([]*UserPointLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserPointLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uplid, uid, change_type, ref_id, num, note, created ` +
		`FROM ` + tableName +
		` WHERE ref_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, refID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, refID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, refID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserPointLog, 0)
	for queryData.Next() {
		upl := UserPointLog{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&upl.Uplid, &upl.UID, &upl.ChangeType, &upl.RefID, &upl.Num, &upl.Note, &upl.Created)
		if err != nil {
			return nil, err
		}

		res = append(res, &upl)
	}

	return res, nil
}

// UserPointLogsByUID retrieves a row from 'aypcddg.user_point_log' as a UserPointLog.
//
// Generated from index 'uid'.
func UserPointLogsByUID(ctx context.Context, uid uint, key ...interface{}) ([]*UserPointLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserPointLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uplid, uid, change_type, ref_id, num, note, created ` +
		`FROM ` + tableName +
		` WHERE uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserPointLog, 0)
	for queryData.Next() {
		upl := UserPointLog{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&upl.Uplid, &upl.UID, &upl.ChangeType, &upl.RefID, &upl.Num, &upl.Note, &upl.Created)
		if err != nil {
			return nil, err
		}

		res = append(res, &upl)
	}

	return res, nil
}

// UserPointLogByUplid retrieves a row from 'aypcddg.user_point_log' as a UserPointLog.
//
// Generated from index 'user_point_log_uplid_pkey'.
func UserPointLogByUplid(ctx context.Context, uplid uint64, key ...interface{}) (*UserPointLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserPointLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uplid, uid, change_type, ref_id, num, note, created ` +
		`FROM ` + tableName +
		` WHERE uplid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uplid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	upl := UserPointLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, uplid).Scan(&upl.Uplid, &upl.UID, &upl.ChangeType, &upl.RefID, &upl.Num, &upl.Note, &upl.Created)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, uplid).Scan(&upl.Uplid, &upl.UID, &upl.ChangeType, &upl.RefID, &upl.Num, &upl.Note, &upl.Created)
		if err != nil {
			return nil, err
		}
	}

	return &upl, nil
}