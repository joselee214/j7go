// Package xo contains the types for schema 'aypcddg'.
package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"fmt"
	"j7go/components"
	"j7go/utils"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// MidAdminUserPermission represents a row from 'aypcddg.mid_admin_user_permissions'.
type MidAdminUserPermission struct {
	UserID       int            `json:"user_id"`       // user_id
	PermissionID int            `json:"permission_id"` // permission_id
	CreatedAt    mysql.NullTime `json:"created_at"`    // created_at
	UpdatedAt    mysql.NullTime `json:"updated_at"`    // updated_at
}

// MidAdminUserPermissionsByUserIDPermissionID retrieves a row from 'aypcddg.mid_admin_user_permissions' as a MidAdminUserPermission.
//
// Generated from index 'mid_admin_user_permissions_user_id_permission_id_index'.
func MidAdminUserPermissionsByUserIDPermissionID(ctx context.Context, userID int, permissionID int, key ...interface{}) ([]*MidAdminUserPermission, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminUserPermissionTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`user_id, permission_id, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE user_id = ? AND permission_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, userID, permissionID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, userID, permissionID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, userID, permissionID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*MidAdminUserPermission, 0)
	for queryData.Next() {
		maup := MidAdminUserPermission{}

		// scan
		err = queryData.Scan(&maup.UserID, &maup.PermissionID, &maup.CreatedAt, &maup.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &maup)
	}

	return res, nil
}
