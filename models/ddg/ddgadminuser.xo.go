// Package ddg contains the types for schema 'aypcddg'.
package ddg

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

// DdgAdminUser represents a row from 'aypcddg.ddg_admin_user'.
type DdgAdminUser struct {
	ID            uint64         `json:"id"`             // id
	Name          string         `json:"name"`           // name
	Account       string         `json:"account"`        // account
	Password      string         `json:"password"`       // password
	PermissionIds sql.NullString `json:"permission_ids"` // permission_ids
	Status        int8           `json:"status"`         // status

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DdgAdminUser exists in the database.
func (dau *DdgAdminUser) Exists() bool { //ddg_admin_user
	return dau._exists
}

// Deleted provides information if the DdgAdminUser has been deleted from the database.
func (dau *DdgAdminUser) Deleted() bool {
	return dau._deleted
}

// Get table name
func GetDdgAdminUserTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "ddg_admin_user", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DdgAdminUser to the database.
func (dau *DdgAdminUser) Insert(ctx context.Context, key ...interface{}) error {
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

	tableName, err := GetDdgAdminUserTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`name, account, password, permission_ids, status` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dau.Name, dau.Account, dau.Password, dau.PermissionIds, dau.Status)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dau.Name, dau.Account, dau.Password, dau.PermissionIds, dau.Status)
	} else {
		res, err = dbConn.Exec(sqlstr, dau.Name, dau.Account, dau.Password, dau.PermissionIds, dau.Status)
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
	dau.ID = uint64(id)
	dau._exists = true

	return nil
}

// Update updates the DdgAdminUser in the database.
func (dau *DdgAdminUser) Update(ctx context.Context, key ...interface{}) error {
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

	tableName, err := GetDdgAdminUserTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`name = ?, account = ?, password = ?, permission_ids = ?, status = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dau.Name, dau.Account, dau.Password, dau.PermissionIds, dau.Status, dau.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dau.Name, dau.Account, dau.Password, dau.PermissionIds, dau.Status, dau.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dau.Name, dau.Account, dau.Password, dau.PermissionIds, dau.Status, dau.ID)
	}
	return err
}

// Save saves the DdgAdminUser to the database.
func (dau *DdgAdminUser) Save(ctx context.Context) error {
	if dau.Exists() {
		return dau.Update(ctx)
	}

	return dau.Insert(ctx)
}

// Delete deletes the DdgAdminUser from the database.
func (dau *DdgAdminUser) Delete(ctx context.Context, key ...interface{}) error {
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

	tableName, err := GetDdgAdminUserTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dau.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dau.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dau.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	dau._deleted = true

	return nil
}

// DdgAdminUserByAccount retrieves a row from 'aypcddg.ddg_admin_user' as a DdgAdminUser.
//
// Generated from index 'ddg_admin_user_account_unique'.
func DdgAdminUserByAccount(ctx context.Context, account string, key ...interface{}) (*DdgAdminUser, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDdgAdminUserTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, account, password, permission_ids, status ` +
		`FROM ` + tableName +
		` WHERE account = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, account)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	dau := DdgAdminUser{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, account).Scan(&dau.ID, &dau.Name, &dau.Account, &dau.Password, &dau.PermissionIds, &dau.Status)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, account).Scan(&dau.ID, &dau.Name, &dau.Account, &dau.Password, &dau.PermissionIds, &dau.Status)
		if err != nil {
			return nil, err
		}
	}

	return &dau, nil
}

// DdgAdminUserByID retrieves a row from 'aypcddg.ddg_admin_user' as a DdgAdminUser.
//
// Generated from index 'ddg_admin_user_id_pkey'.
func DdgAdminUserByID(ctx context.Context, id uint64, key ...interface{}) (*DdgAdminUser, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDdgAdminUserTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, account, password, permission_ids, status ` +
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
	dau := DdgAdminUser{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&dau.ID, &dau.Name, &dau.Account, &dau.Password, &dau.PermissionIds, &dau.Status)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&dau.ID, &dau.Name, &dau.Account, &dau.Password, &dau.PermissionIds, &dau.Status)
		if err != nil {
			return nil, err
		}
	}

	return &dau, nil
}
