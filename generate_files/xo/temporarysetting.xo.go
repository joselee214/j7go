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
	"time"

	"go.uber.org/zap"
)

// TemporarySetting represents a row from 'aypcddg.temporary_setting'.
type TemporarySetting struct {
	Keyid   string    `json:"keyid"`   // keyid
	Created time.Time `json:"created"` // created
	Data    string    `json:"data"`    // data
	Type    string    `json:"type"`    // type

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TemporarySetting exists in the database.
func (ts *TemporarySetting) Exists() bool { //temporary_setting
	return ts._exists
}

// Deleted provides information if the TemporarySetting has been deleted from the database.
func (ts *TemporarySetting) Deleted() bool {
	return ts._deleted
}

// Get table name
func GetTemporarySettingTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "temporary_setting", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the TemporarySetting to the database.
func (ts *TemporarySetting) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ts._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTemporarySettingTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`keyid, created, data, type` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ts.Keyid, ts.Created, ts.Data, ts.Type)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, ts.Keyid, ts.Created, ts.Data, ts.Type)
	} else {
		res, err = dbConn.Exec(sqlstr, ts.Keyid, ts.Created, ts.Data, ts.Type)
	}

	if err != nil {
		return err
	}

	// set existence
	ts._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ts.Keyid = string(id)
	ts._exists = true

	return nil
}

// Update updates the TemporarySetting in the database.
func (ts *TemporarySetting) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ts._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTemporarySettingTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`created = ?, data = ?, type = ?` +
		` WHERE keyid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ts.Created, ts.Data, ts.Type, ts.Keyid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ts.Created, ts.Data, ts.Type, ts.Keyid)
	} else {
		_, err = dbConn.Exec(sqlstr, ts.Created, ts.Data, ts.Type, ts.Keyid)
	}
	return err
}

// Save saves the TemporarySetting to the database.
func (ts *TemporarySetting) Save(ctx context.Context) error {
	if ts.Exists() {
		return ts.Update(ctx)
	}

	return ts.Insert(ctx)
}

// Delete deletes the TemporarySetting from the database.
func (ts *TemporarySetting) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ts._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTemporarySettingTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE keyid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ts.Keyid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ts.Keyid)
	} else {
		_, err = dbConn.Exec(sqlstr, ts.Keyid)
	}

	if err != nil {
		return err
	}

	// set deleted
	ts._deleted = true

	return nil
}

// TemporarySettingByKeyid retrieves a row from 'aypcddg.temporary_setting' as a TemporarySetting.
//
// Generated from index 'keyid'.
func TemporarySettingByKeyid(ctx context.Context, keyid string, key ...interface{}) (*TemporarySetting, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetTemporarySettingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`keyid, created, data, type ` +
		`FROM ` + tableName +
		` WHERE keyid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, keyid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ts := TemporarySetting{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, keyid).Scan(&ts.Keyid, &ts.Created, &ts.Data, &ts.Type)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, keyid).Scan(&ts.Keyid, &ts.Created, &ts.Data, &ts.Type)
		if err != nil {
			return nil, err
		}
	}

	return &ts, nil
}

// TemporarySettingByKeyid retrieves a row from 'aypcddg.temporary_setting' as a TemporarySetting.
//
// Generated from index 'temporary_setting_keyid_pkey'.
func TemporarySettingByKeyid(ctx context.Context, keyid string, key ...interface{}) (*TemporarySetting, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetTemporarySettingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`keyid, created, data, type ` +
		`FROM ` + tableName +
		` WHERE keyid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, keyid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ts := TemporarySetting{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, keyid).Scan(&ts.Keyid, &ts.Created, &ts.Data, &ts.Type)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, keyid).Scan(&ts.Keyid, &ts.Created, &ts.Data, &ts.Type)
		if err != nil {
			return nil, err
		}
	}

	return &ts, nil
}
