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

// MidAdminConfig represents a row from 'aypcddg.mid_admin_config'.
type MidAdminConfig struct {
	ID           uint           `json:"id"`             // id
	Name         string         `json:"name"`           // name
	Value        string         `json:"value"`          // value
	Description  sql.NullString `json:"description"`    // description
	CreatedAt    mysql.NullTime `json:"created_at"`     // created_at
	UpdatedAt    mysql.NullTime `json:"updated_at"`     // updated_at
	MidAdminName sql.NullString `json:"mid_admin_name"` // mid_admin_name
	MidAdminID   int            `json:"mid_admin_id"`   // mid_admin_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MidAdminConfig exists in the database.
func (mac *MidAdminConfig) Exists() bool { //mid_admin_config
	return mac._exists
}

// Deleted provides information if the MidAdminConfig has been deleted from the database.
func (mac *MidAdminConfig) Deleted() bool {
	return mac._deleted
}

// Get table name
func GetMidAdminConfigTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "mid_admin_config", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the MidAdminConfig to the database.
func (mac *MidAdminConfig) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if mac._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminConfigTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`name, value, description, created_at, updated_at, mid_admin_name, mid_admin_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mac.Name, mac.Value, mac.Description, mac.CreatedAt, mac.UpdatedAt, mac.MidAdminName, mac.MidAdminID)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, mac.Name, mac.Value, mac.Description, mac.CreatedAt, mac.UpdatedAt, mac.MidAdminName, mac.MidAdminID)
	} else {
		res, err = dbConn.Exec(sqlstr, mac.Name, mac.Value, mac.Description, mac.CreatedAt, mac.UpdatedAt, mac.MidAdminName, mac.MidAdminID)
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
	mac.ID = uint(id)
	mac._exists = true

	return nil
}

// Update updates the MidAdminConfig in the database.
func (mac *MidAdminConfig) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if mac._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminConfigTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`name = ?, value = ?, description = ?, created_at = ?, updated_at = ?, mid_admin_name = ?, mid_admin_id = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mac.Name, mac.Value, mac.Description, mac.CreatedAt, mac.UpdatedAt, mac.MidAdminName, mac.MidAdminID, mac.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, mac.Name, mac.Value, mac.Description, mac.CreatedAt, mac.UpdatedAt, mac.MidAdminName, mac.MidAdminID, mac.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, mac.Name, mac.Value, mac.Description, mac.CreatedAt, mac.UpdatedAt, mac.MidAdminName, mac.MidAdminID, mac.ID)
	}
	return err
}

// Save saves the MidAdminConfig to the database.
func (mac *MidAdminConfig) Save(ctx context.Context) error {
	if mac.Exists() {
		return mac.Update(ctx)
	}

	return mac.Insert(ctx)
}

// Delete deletes the MidAdminConfig from the database.
func (mac *MidAdminConfig) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if mac._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminConfigTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mac.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, mac.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, mac.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	mac._deleted = true

	return nil
}

// MidAdminConfigByID retrieves a row from 'aypcddg.mid_admin_config' as a MidAdminConfig.
//
// Generated from index 'mid_admin_config_id_pkey'.
func MidAdminConfigByID(ctx context.Context, id uint, key ...interface{}) (*MidAdminConfig, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminConfigTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, value, description, created_at, updated_at, mid_admin_name, mid_admin_id ` +
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
	mac := MidAdminConfig{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&mac.ID, &mac.Name, &mac.Value, &mac.Description, &mac.CreatedAt, &mac.UpdatedAt, &mac.MidAdminName, &mac.MidAdminID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&mac.ID, &mac.Name, &mac.Value, &mac.Description, &mac.CreatedAt, &mac.UpdatedAt, &mac.MidAdminName, &mac.MidAdminID)
		if err != nil {
			return nil, err
		}
	}

	return &mac, nil
}

// MidAdminConfigByName retrieves a row from 'aypcddg.mid_admin_config' as a MidAdminConfig.
//
// Generated from index 'mid_admin_config_name_unique'.
func MidAdminConfigByName(ctx context.Context, name string, key ...interface{}) (*MidAdminConfig, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminConfigTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, value, description, created_at, updated_at, mid_admin_name, mid_admin_id ` +
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
	mac := MidAdminConfig{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, name).Scan(&mac.ID, &mac.Name, &mac.Value, &mac.Description, &mac.CreatedAt, &mac.UpdatedAt, &mac.MidAdminName, &mac.MidAdminID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, name).Scan(&mac.ID, &mac.Name, &mac.Value, &mac.Description, &mac.CreatedAt, &mac.UpdatedAt, &mac.MidAdminName, &mac.MidAdminID)
		if err != nil {
			return nil, err
		}
	}

	return &mac, nil
}
