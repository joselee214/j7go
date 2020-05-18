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

// DevelopAdminUser represents a row from 'aypcddg.develop_admin_user'.
type DevelopAdminUser struct {
	UID          uint           `json:"uid"`            // uid
	Name         sql.NullString `json:"name"`           // name
	Password     sql.NullString `json:"password"`       // password
	Extent       sql.NullString `json:"extent"`         // extent
	IsLock       sql.NullInt64  `json:"is_lock"`        // is_lock
	OaPositionID sql.NullInt64  `json:"oa_position_id"` // oa_position_id
	Truename     sql.NullString `json:"truename"`       // truename
	FirstEntry   int            `json:"first_entry"`    // first_entry

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DevelopAdminUser exists in the database.
func (dau *DevelopAdminUser) Exists() bool { //develop_admin_user
	return dau._exists
}

// Deleted provides information if the DevelopAdminUser has been deleted from the database.
func (dau *DevelopAdminUser) Deleted() bool {
	return dau._deleted
}

// Get table name
func GetDevelopAdminUserTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "develop_admin_user", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DevelopAdminUser to the database.
func (dau *DevelopAdminUser) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if dau._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDevelopAdminUserTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`name, password, extent, is_lock, oa_position_id, truename, first_entry` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dau.Name, dau.Password, dau.Extent, dau.IsLock, dau.OaPositionID, dau.Truename, dau.FirstEntry)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dau.Name, dau.Password, dau.Extent, dau.IsLock, dau.OaPositionID, dau.Truename, dau.FirstEntry)
	} else {
		res, err = dbConn.Exec(sqlstr, dau.Name, dau.Password, dau.Extent, dau.IsLock, dau.OaPositionID, dau.Truename, dau.FirstEntry)
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
	dau.UID = uint(id)
	dau._exists = true

	return nil
}

// Update updates the DevelopAdminUser in the database.
func (dau *DevelopAdminUser) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dau._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDevelopAdminUserTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`name = ?, password = ?, extent = ?, is_lock = ?, oa_position_id = ?, truename = ?, first_entry = ?` +
		` WHERE uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dau.Name, dau.Password, dau.Extent, dau.IsLock, dau.OaPositionID, dau.Truename, dau.FirstEntry, dau.UID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dau.Name, dau.Password, dau.Extent, dau.IsLock, dau.OaPositionID, dau.Truename, dau.FirstEntry, dau.UID)
	} else {
		_, err = dbConn.Exec(sqlstr, dau.Name, dau.Password, dau.Extent, dau.IsLock, dau.OaPositionID, dau.Truename, dau.FirstEntry, dau.UID)
	}
	return err
}

// Save saves the DevelopAdminUser to the database.
func (dau *DevelopAdminUser) Save(ctx context.Context) error {
	if dau.Exists() {
		return dau.Update(ctx)
	}

	return dau.Insert(ctx)
}

// Delete deletes the DevelopAdminUser from the database.
func (dau *DevelopAdminUser) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dau._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDevelopAdminUserTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dau.UID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dau.UID)
	} else {
		_, err = dbConn.Exec(sqlstr, dau.UID)
	}

	if err != nil {
		return err
	}

	// set deleted
	dau._deleted = true

	return nil
}

// DevelopAdminUserByUID retrieves a row from 'aypcddg.develop_admin_user' as a DevelopAdminUser.
//
// Generated from index 'develop_admin_user_uid_pkey'.
func DevelopAdminUserByUID(ctx context.Context, uid uint, key ...interface{}) (*DevelopAdminUser, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDevelopAdminUserTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uid, name, password, extent, is_lock, oa_position_id, truename, first_entry ` +
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
	dau := DevelopAdminUser{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, uid).Scan(&dau.UID, &dau.Name, &dau.Password, &dau.Extent, &dau.IsLock, &dau.OaPositionID, &dau.Truename, &dau.FirstEntry)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, uid).Scan(&dau.UID, &dau.Name, &dau.Password, &dau.Extent, &dau.IsLock, &dau.OaPositionID, &dau.Truename, &dau.FirstEntry)
		if err != nil {
			return nil, err
		}
	}

	return &dau, nil
}
