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

// DealerUserBuyUserMapping represents a row from 'aypcddg.dealer_user_buy_user_mapping'.
type DealerUserBuyUserMapping struct {
	ID             uint64         `json:"id"`              // id
	UID            int64          `json:"uid"`             // uid
	SubordinateUID int64          `json:"subordinate_uid"` // subordinate_uid
	Mobile         string         `json:"mobile"`          // mobile
	Name           string         `json:"name"`            // name
	Status         int8           `json:"status"`          // status
	Remark         sql.NullString `json:"remark"`          // remark
	MidAdminName   sql.NullString `json:"mid_admin_name"`  // mid_admin_name
	MidAdminID     int            `json:"mid_admin_id"`    // mid_admin_id
	CreatedAt      mysql.NullTime `json:"created_at"`      // created_at
	UpdatedAt      mysql.NullTime `json:"updated_at"`      // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DealerUserBuyUserMapping exists in the database.
func (dubum *DealerUserBuyUserMapping) Exists() bool { //dealer_user_buy_user_mapping
	return dubum._exists
}

// Deleted provides information if the DealerUserBuyUserMapping has been deleted from the database.
func (dubum *DealerUserBuyUserMapping) Deleted() bool {
	return dubum._deleted
}

// Get table name
func GetDealerUserBuyUserMappingTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "dealer_user_buy_user_mapping", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DealerUserBuyUserMapping to the database.
func (dubum *DealerUserBuyUserMapping) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if dubum._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDealerUserBuyUserMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, subordinate_uid, mobile, name, status, remark, mid_admin_name, mid_admin_id, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dubum.UID, dubum.SubordinateUID, dubum.Mobile, dubum.Name, dubum.Status, dubum.Remark, dubum.MidAdminName, dubum.MidAdminID, dubum.CreatedAt, dubum.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dubum.UID, dubum.SubordinateUID, dubum.Mobile, dubum.Name, dubum.Status, dubum.Remark, dubum.MidAdminName, dubum.MidAdminID, dubum.CreatedAt, dubum.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, dubum.UID, dubum.SubordinateUID, dubum.Mobile, dubum.Name, dubum.Status, dubum.Remark, dubum.MidAdminName, dubum.MidAdminID, dubum.CreatedAt, dubum.UpdatedAt)
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
	dubum.ID = uint64(id)
	dubum._exists = true

	return nil
}

// Update updates the DealerUserBuyUserMapping in the database.
func (dubum *DealerUserBuyUserMapping) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dubum._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDealerUserBuyUserMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, subordinate_uid = ?, mobile = ?, name = ?, status = ?, remark = ?, mid_admin_name = ?, mid_admin_id = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dubum.UID, dubum.SubordinateUID, dubum.Mobile, dubum.Name, dubum.Status, dubum.Remark, dubum.MidAdminName, dubum.MidAdminID, dubum.CreatedAt, dubum.UpdatedAt, dubum.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dubum.UID, dubum.SubordinateUID, dubum.Mobile, dubum.Name, dubum.Status, dubum.Remark, dubum.MidAdminName, dubum.MidAdminID, dubum.CreatedAt, dubum.UpdatedAt, dubum.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dubum.UID, dubum.SubordinateUID, dubum.Mobile, dubum.Name, dubum.Status, dubum.Remark, dubum.MidAdminName, dubum.MidAdminID, dubum.CreatedAt, dubum.UpdatedAt, dubum.ID)
	}
	return err
}

// Save saves the DealerUserBuyUserMapping to the database.
func (dubum *DealerUserBuyUserMapping) Save(ctx context.Context) error {
	if dubum.Exists() {
		return dubum.Update(ctx)
	}

	return dubum.Insert(ctx)
}

// Delete deletes the DealerUserBuyUserMapping from the database.
func (dubum *DealerUserBuyUserMapping) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dubum._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDealerUserBuyUserMappingTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dubum.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dubum.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dubum.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	dubum._deleted = true

	return nil
}

// DealerUserBuyUserMappingByID retrieves a row from 'aypcddg.dealer_user_buy_user_mapping' as a DealerUserBuyUserMapping.
//
// Generated from index 'dealer_user_buy_user_mapping_id_pkey'.
func DealerUserBuyUserMappingByID(ctx context.Context, id uint64, key ...interface{}) (*DealerUserBuyUserMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDealerUserBuyUserMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, subordinate_uid, mobile, name, status, remark, mid_admin_name, mid_admin_id, created_at, updated_at ` +
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
	dubum := DealerUserBuyUserMapping{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&dubum.ID, &dubum.UID, &dubum.SubordinateUID, &dubum.Mobile, &dubum.Name, &dubum.Status, &dubum.Remark, &dubum.MidAdminName, &dubum.MidAdminID, &dubum.CreatedAt, &dubum.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&dubum.ID, &dubum.UID, &dubum.SubordinateUID, &dubum.Mobile, &dubum.Name, &dubum.Status, &dubum.Remark, &dubum.MidAdminName, &dubum.MidAdminID, &dubum.CreatedAt, &dubum.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &dubum, nil
}
