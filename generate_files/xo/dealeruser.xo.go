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

// DealerUser represents a row from 'aypcddg.dealer_user'.
type DealerUser struct {
	Duid         uint64         `json:"duid"`           // duid
	UID          int64          `json:"uid"`            // uid
	Status       int8           `json:"status"`         // status
	MidAdminName sql.NullString `json:"mid_admin_name"` // mid_admin_name
	MidAdminID   int            `json:"mid_admin_id"`   // mid_admin_id
	CreatedAt    mysql.NullTime `json:"created_at"`     // created_at
	UpdatedAt    mysql.NullTime `json:"updated_at"`     // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DealerUser exists in the database.
func (du *DealerUser) Exists() bool { //dealer_user
	return du._exists
}

// Deleted provides information if the DealerUser has been deleted from the database.
func (du *DealerUser) Deleted() bool {
	return du._deleted
}

// Get table name
func GetDealerUserTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "dealer_user", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DealerUser to the database.
func (du *DealerUser) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if du._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDealerUserTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, status, mid_admin_name, mid_admin_id, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, du.UID, du.Status, du.MidAdminName, du.MidAdminID, du.CreatedAt, du.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, du.UID, du.Status, du.MidAdminName, du.MidAdminID, du.CreatedAt, du.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, du.UID, du.Status, du.MidAdminName, du.MidAdminID, du.CreatedAt, du.UpdatedAt)
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
	du.Duid = uint64(id)
	du._exists = true

	return nil
}

// Update updates the DealerUser in the database.
func (du *DealerUser) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if du._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDealerUserTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, status = ?, mid_admin_name = ?, mid_admin_id = ?, created_at = ?, updated_at = ?` +
		` WHERE duid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, du.UID, du.Status, du.MidAdminName, du.MidAdminID, du.CreatedAt, du.UpdatedAt, du.Duid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, du.UID, du.Status, du.MidAdminName, du.MidAdminID, du.CreatedAt, du.UpdatedAt, du.Duid)
	} else {
		_, err = dbConn.Exec(sqlstr, du.UID, du.Status, du.MidAdminName, du.MidAdminID, du.CreatedAt, du.UpdatedAt, du.Duid)
	}
	return err
}

// Save saves the DealerUser to the database.
func (du *DealerUser) Save(ctx context.Context) error {
	if du.Exists() {
		return du.Update(ctx)
	}

	return du.Insert(ctx)
}

// Delete deletes the DealerUser from the database.
func (du *DealerUser) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if du._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDealerUserTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE duid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, du.Duid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, du.Duid)
	} else {
		_, err = dbConn.Exec(sqlstr, du.Duid)
	}

	if err != nil {
		return err
	}

	// set deleted
	du._deleted = true

	return nil
}

// DealerUserByDuid retrieves a row from 'aypcddg.dealer_user' as a DealerUser.
//
// Generated from index 'dealer_user_duid_pkey'.
func DealerUserByDuid(ctx context.Context, duid uint64, key ...interface{}) (*DealerUser, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDealerUserTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`duid, uid, status, mid_admin_name, mid_admin_id, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE duid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, duid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	du := DealerUser{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, duid).Scan(&du.Duid, &du.UID, &du.Status, &du.MidAdminName, &du.MidAdminID, &du.CreatedAt, &du.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, duid).Scan(&du.Duid, &du.UID, &du.Status, &du.MidAdminName, &du.MidAdminID, &du.CreatedAt, &du.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &du, nil
}