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

// SaleUser represents a row from 'aypcddg.sale_user'.
type SaleUser struct {
	UID         uint           `json:"uid"`         // uid
	Sid         uint           `json:"sid"`         // sid
	Permissions sql.NullString `json:"permissions"` // permissions
	Status      int8           `json:"status"`      // status
	Created     time.Time      `json:"created"`     // created

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the SaleUser exists in the database.
func (su *SaleUser) Exists() bool { //sale_user
	return su._exists
}

// Deleted provides information if the SaleUser has been deleted from the database.
func (su *SaleUser) Deleted() bool {
	return su._deleted
}

// Get table name
func GetSaleUserTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "sale_user", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the SaleUser to the database.
func (su *SaleUser) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if su._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSaleUserTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, sid, permissions, status, created` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, su.UID, su.Sid, su.Permissions, su.Status, su.Created)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, su.UID, su.Sid, su.Permissions, su.Status, su.Created)
	} else {
		res, err = dbConn.Exec(sqlstr, su.UID, su.Sid, su.Permissions, su.Status, su.Created)
	}

	if err != nil {
		return err
	}

	// set existence
	su._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	su.Sid = uint(id)
	su._exists = true

	return nil
}

// Update updates the SaleUser in the database.
func (su *SaleUser) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if su._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSaleUserTableName(key...)
	if err != nil {
		return err
	}

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`permissions = ?, status = ?, created = ?` +
		` WHERE uid = ? AND sid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, su.Permissions, su.Status, su.Created, su.UID, su.Sid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, su.Permissions, su.Status, su.Created, su.UID, su.Sid)
	} else {
		_, err = dbConn.Exec(sqlstr, su.Permissions, su.Status, su.Created, su.UID, su.Sid)
	}
	return err
}

// Save saves the SaleUser to the database.
func (su *SaleUser) Save(ctx context.Context) error {
	if su.Exists() {
		return su.Update(ctx)
	}

	return su.Insert(ctx)
}

// Delete deletes the SaleUser from the database.
func (su *SaleUser) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if su._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSaleUserTableName(key...)
	if err != nil {
		return err
	}
	//2

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE sid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, su.Sid)))

	if tx != nil {
		_, err = tx.Exec(sqlstr, su.Sid)
	} else {
		_, err = dbConn.Exec(sqlstr, su.Sid)
	}
	if err != nil {
		return err
	}

	// set deleted
	su._deleted = true

	return nil
}

// SaleUserBySid retrieves a row from 'aypcddg.sale_user' as a SaleUser.
//
// Generated from index 'sale_user_sid_pkey'.
func SaleUserBySid(ctx context.Context, sid uint, key ...interface{}) (*SaleUser, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetSaleUserTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uid, sid, permissions, status, created ` +
		`FROM ` + tableName +
		` WHERE sid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, sid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	su := SaleUser{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, sid).Scan(&su.UID, &su.Sid, &su.Permissions, &su.Status, &su.Created)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, sid).Scan(&su.UID, &su.Sid, &su.Permissions, &su.Status, &su.Created)
		if err != nil {
			return nil, err
		}
	}

	return &su, nil
}
