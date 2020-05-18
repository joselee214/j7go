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

// MidAdminRole represents a row from 'aypcddg.mid_admin_roles'.
type MidAdminRole struct {
	ID        uint           `json:"id"`         // id
	Name      string         `json:"name"`       // name
	Slug      string         `json:"slug"`       // slug
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MidAdminRole exists in the database.
func (mar *MidAdminRole) Exists() bool { //mid_admin_roles
	return mar._exists
}

// Deleted provides information if the MidAdminRole has been deleted from the database.
func (mar *MidAdminRole) Deleted() bool {
	return mar._deleted
}

// Get table name
func GetMidAdminRoleTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "mid_admin_roles", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the MidAdminRole to the database.
func (mar *MidAdminRole) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if mar._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminRoleTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`name, slug, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mar.Name, mar.Slug, mar.CreatedAt, mar.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, mar.Name, mar.Slug, mar.CreatedAt, mar.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, mar.Name, mar.Slug, mar.CreatedAt, mar.UpdatedAt)
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
	mar.ID = uint(id)
	mar._exists = true

	return nil
}

// Update updates the MidAdminRole in the database.
func (mar *MidAdminRole) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if mar._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminRoleTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`name = ?, slug = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mar.Name, mar.Slug, mar.CreatedAt, mar.UpdatedAt, mar.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, mar.Name, mar.Slug, mar.CreatedAt, mar.UpdatedAt, mar.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, mar.Name, mar.Slug, mar.CreatedAt, mar.UpdatedAt, mar.ID)
	}
	return err
}

// Save saves the MidAdminRole to the database.
func (mar *MidAdminRole) Save(ctx context.Context) error {
	if mar.Exists() {
		return mar.Update(ctx)
	}

	return mar.Insert(ctx)
}

// Delete deletes the MidAdminRole from the database.
func (mar *MidAdminRole) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if mar._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminRoleTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mar.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, mar.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, mar.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	mar._deleted = true

	return nil
}

// MidAdminRoleByID retrieves a row from 'aypcddg.mid_admin_roles' as a MidAdminRole.
//
// Generated from index 'mid_admin_roles_id_pkey'.
func MidAdminRoleByID(ctx context.Context, id uint, key ...interface{}) (*MidAdminRole, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminRoleTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, slug, created_at, updated_at ` +
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
	mar := MidAdminRole{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&mar.ID, &mar.Name, &mar.Slug, &mar.CreatedAt, &mar.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&mar.ID, &mar.Name, &mar.Slug, &mar.CreatedAt, &mar.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &mar, nil
}

// MidAdminRoleByName retrieves a row from 'aypcddg.mid_admin_roles' as a MidAdminRole.
//
// Generated from index 'mid_admin_roles_name_unique'.
func MidAdminRoleByName(ctx context.Context, name string, key ...interface{}) (*MidAdminRole, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminRoleTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, slug, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE name = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, name)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	mar := MidAdminRole{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, name).Scan(&mar.ID, &mar.Name, &mar.Slug, &mar.CreatedAt, &mar.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, name).Scan(&mar.ID, &mar.Name, &mar.Slug, &mar.CreatedAt, &mar.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &mar, nil
}

// MidAdminRoleBySlug retrieves a row from 'aypcddg.mid_admin_roles' as a MidAdminRole.
//
// Generated from index 'mid_admin_roles_slug_unique'.
func MidAdminRoleBySlug(ctx context.Context, slug string, key ...interface{}) (*MidAdminRole, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminRoleTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, slug, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE slug = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, slug)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	mar := MidAdminRole{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, slug).Scan(&mar.ID, &mar.Name, &mar.Slug, &mar.CreatedAt, &mar.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, slug).Scan(&mar.ID, &mar.Name, &mar.Slug, &mar.CreatedAt, &mar.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &mar, nil
}
