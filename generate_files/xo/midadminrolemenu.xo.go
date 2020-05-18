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

// MidAdminRoleMenu represents a row from 'aypcddg.mid_admin_role_menu'.
type MidAdminRoleMenu struct {
	RoleID    int            `json:"role_id"`    // role_id
	MenuID    int            `json:"menu_id"`    // menu_id
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at
}

// MidAdminRoleMenusByRoleIDMenuID retrieves a row from 'aypcddg.mid_admin_role_menu' as a MidAdminRoleMenu.
//
// Generated from index 'mid_admin_role_menu_role_id_menu_id_index'.
func MidAdminRoleMenusByRoleIDMenuID(ctx context.Context, roleID int, menuID int, key ...interface{}) ([]*MidAdminRoleMenu, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminRoleMenuTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`role_id, menu_id, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE role_id = ? AND menu_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, roleID, menuID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, roleID, menuID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, roleID, menuID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*MidAdminRoleMenu, 0)
	for queryData.Next() {
		marm := MidAdminRoleMenu{}

		// scan
		err = queryData.Scan(&marm.RoleID, &marm.MenuID, &marm.CreatedAt, &marm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &marm)
	}

	return res, nil
}