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

// UserFeedBack represents a row from 'aypcddg.user_feed_back'.
type UserFeedBack struct {
	ID           uint64         `json:"id"`             // id
	UID          int            `json:"uid"`            // uid
	IssueID      int            `json:"issue_id"`       // issue_id
	Content      string         `json:"content"`        // content
	Contact      string         `json:"contact"`        // contact
	Device       string         `json:"device"`         // device
	Imgs         string         `json:"imgs"`           // imgs
	Status       int8           `json:"status"`         // status
	Remark       string         `json:"remark"`         // remark
	Platform     string         `json:"platform"`       // platform
	MidAdminName sql.NullString `json:"mid_admin_name"` // mid_admin_name
	MidAdminID   int            `json:"mid_admin_id"`   // mid_admin_id
	CreatedAt    mysql.NullTime `json:"created_at"`     // created_at
	UpdatedAt    mysql.NullTime `json:"updated_at"`     // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserFeedBack exists in the database.
func (ufb *UserFeedBack) Exists() bool { //user_feed_back
	return ufb._exists
}

// Deleted provides information if the UserFeedBack has been deleted from the database.
func (ufb *UserFeedBack) Deleted() bool {
	return ufb._deleted
}

// Get table name
func GetUserFeedBackTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "user_feed_back", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserFeedBack to the database.
func (ufb *UserFeedBack) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ufb._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserFeedBackTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, issue_id, content, contact, device, imgs, status, remark, platform, mid_admin_name, mid_admin_id, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ufb.UID, ufb.IssueID, ufb.Content, ufb.Contact, ufb.Device, ufb.Imgs, ufb.Status, ufb.Remark, ufb.Platform, ufb.MidAdminName, ufb.MidAdminID, ufb.CreatedAt, ufb.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ufb.UID, ufb.IssueID, ufb.Content, ufb.Contact, ufb.Device, ufb.Imgs, ufb.Status, ufb.Remark, ufb.Platform, ufb.MidAdminName, ufb.MidAdminID, ufb.CreatedAt, ufb.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, ufb.UID, ufb.IssueID, ufb.Content, ufb.Contact, ufb.Device, ufb.Imgs, ufb.Status, ufb.Remark, ufb.Platform, ufb.MidAdminName, ufb.MidAdminID, ufb.CreatedAt, ufb.UpdatedAt)
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
	ufb.ID = uint64(id)
	ufb._exists = true

	return nil
}

// Update updates the UserFeedBack in the database.
func (ufb *UserFeedBack) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ufb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserFeedBackTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, issue_id = ?, content = ?, contact = ?, device = ?, imgs = ?, status = ?, remark = ?, platform = ?, mid_admin_name = ?, mid_admin_id = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ufb.UID, ufb.IssueID, ufb.Content, ufb.Contact, ufb.Device, ufb.Imgs, ufb.Status, ufb.Remark, ufb.Platform, ufb.MidAdminName, ufb.MidAdminID, ufb.CreatedAt, ufb.UpdatedAt, ufb.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ufb.UID, ufb.IssueID, ufb.Content, ufb.Contact, ufb.Device, ufb.Imgs, ufb.Status, ufb.Remark, ufb.Platform, ufb.MidAdminName, ufb.MidAdminID, ufb.CreatedAt, ufb.UpdatedAt, ufb.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ufb.UID, ufb.IssueID, ufb.Content, ufb.Contact, ufb.Device, ufb.Imgs, ufb.Status, ufb.Remark, ufb.Platform, ufb.MidAdminName, ufb.MidAdminID, ufb.CreatedAt, ufb.UpdatedAt, ufb.ID)
	}
	return err
}

// Save saves the UserFeedBack to the database.
func (ufb *UserFeedBack) Save(ctx context.Context) error {
	if ufb.Exists() {
		return ufb.Update(ctx)
	}

	return ufb.Insert(ctx)
}

// Delete deletes the UserFeedBack from the database.
func (ufb *UserFeedBack) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ufb._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserFeedBackTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ufb.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ufb.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ufb.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ufb._deleted = true

	return nil
}

// UserFeedBacksByPlatformUID retrieves a row from 'aypcddg.user_feed_back' as a UserFeedBack.
//
// Generated from index 'idx_platform_uid'.
func UserFeedBacksByPlatformUID(ctx context.Context, platform string, uid int, key ...interface{}) ([]*UserFeedBack, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserFeedBackTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, issue_id, content, contact, device, imgs, status, remark, platform, mid_admin_name, mid_admin_id, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE platform = ? AND uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, platform, uid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, platform, uid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, platform, uid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserFeedBack, 0)
	for queryData.Next() {
		ufb := UserFeedBack{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ufb.ID, &ufb.UID, &ufb.IssueID, &ufb.Content, &ufb.Contact, &ufb.Device, &ufb.Imgs, &ufb.Status, &ufb.Remark, &ufb.Platform, &ufb.MidAdminName, &ufb.MidAdminID, &ufb.CreatedAt, &ufb.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &ufb)
	}

	return res, nil
}

// UserFeedBackByID retrieves a row from 'aypcddg.user_feed_back' as a UserFeedBack.
//
// Generated from index 'user_feed_back_id_pkey'.
func UserFeedBackByID(ctx context.Context, id uint64, key ...interface{}) (*UserFeedBack, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserFeedBackTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, issue_id, content, contact, device, imgs, status, remark, platform, mid_admin_name, mid_admin_id, created_at, updated_at ` +
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
	ufb := UserFeedBack{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&ufb.ID, &ufb.UID, &ufb.IssueID, &ufb.Content, &ufb.Contact, &ufb.Device, &ufb.Imgs, &ufb.Status, &ufb.Remark, &ufb.Platform, &ufb.MidAdminName, &ufb.MidAdminID, &ufb.CreatedAt, &ufb.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&ufb.ID, &ufb.UID, &ufb.IssueID, &ufb.Content, &ufb.Contact, &ufb.Device, &ufb.Imgs, &ufb.Status, &ufb.Remark, &ufb.Platform, &ufb.MidAdminName, &ufb.MidAdminID, &ufb.CreatedAt, &ufb.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &ufb, nil
}