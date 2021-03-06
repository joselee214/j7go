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

// ActivityType represents a row from 'aypcddg.activity_type'.
type ActivityType struct {
	ID               uint           `json:"id"`                 // id
	ActivityTypeCode string         `json:"activity_type_code"` // activity_type_code
	ActivityType     string         `json:"activity_type"`      // activity_type
	Status           int8           `json:"status"`             // status
	CreatedAt        mysql.NullTime `json:"created_at"`         // created_at
	UpdatedAt        mysql.NullTime `json:"updated_at"`         // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ActivityType exists in the database.
func (at *ActivityType) Exists() bool { //activity_type
	return at._exists
}

// Deleted provides information if the ActivityType has been deleted from the database.
func (at *ActivityType) Deleted() bool {
	return at._deleted
}

// Get table name
func GetActivityTypeTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "activity_type", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ActivityType to the database.
func (at *ActivityType) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if at._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityTypeTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`activity_type_code, activity_type, status, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, at.ActivityTypeCode, at.ActivityType, at.Status, at.CreatedAt, at.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, at.ActivityTypeCode, at.ActivityType, at.Status, at.CreatedAt, at.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, at.ActivityTypeCode, at.ActivityType, at.Status, at.CreatedAt, at.UpdatedAt)
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
	at.ID = uint(id)
	at._exists = true

	return nil
}

// Update updates the ActivityType in the database.
func (at *ActivityType) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if at._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityTypeTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`activity_type_code = ?, activity_type = ?, status = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, at.ActivityTypeCode, at.ActivityType, at.Status, at.CreatedAt, at.UpdatedAt, at.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, at.ActivityTypeCode, at.ActivityType, at.Status, at.CreatedAt, at.UpdatedAt, at.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, at.ActivityTypeCode, at.ActivityType, at.Status, at.CreatedAt, at.UpdatedAt, at.ID)
	}
	return err
}

// Save saves the ActivityType to the database.
func (at *ActivityType) Save(ctx context.Context) error {
	if at.Exists() {
		return at.Update(ctx)
	}

	return at.Insert(ctx)
}

// Delete deletes the ActivityType from the database.
func (at *ActivityType) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if at._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityTypeTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, at.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, at.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, at.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	at._deleted = true

	return nil
}

// ActivityTypeByID retrieves a row from 'aypcddg.activity_type' as a ActivityType.
//
// Generated from index 'activity_type_id_pkey'.
func ActivityTypeByID(ctx context.Context, id uint, key ...interface{}) (*ActivityType, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityTypeTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_type_code, activity_type, status, created_at, updated_at ` +
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
	at := ActivityType{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&at.ID, &at.ActivityTypeCode, &at.ActivityType, &at.Status, &at.CreatedAt, &at.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&at.ID, &at.ActivityTypeCode, &at.ActivityType, &at.Status, &at.CreatedAt, &at.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &at, nil
}
