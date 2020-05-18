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

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// FactoryActivity represents a row from 'aypcddg.factory_activity'.
type FactoryActivity struct {
	ID               uint           `json:"id"`                 // id
	Name             string         `json:"name"`               // name
	ActivityTypeID   int            `json:"activity_type_id"`   // activity_type_id
	ActivityTypeCode string         `json:"activity_type_code"` // activity_type_code
	Content          string         `json:"content"`            // content
	ActivityImage    string         `json:"activity_image"`     // activity_image
	PicID            int            `json:"pic_id"`             // pic_id
	Status           int            `json:"status"`             // status
	ActivityStatus   int            `json:"activity_status"`    // activity_status
	ReplyContent     sql.NullString `json:"reply_content"`      // reply_content
	FactoryID        int            `json:"factory_id"`         // factory_id
	Sort             int            `json:"sort"`               // sort
	DisTag           int            `json:"dis_tag"`            // dis_tag
	FactoryName      string         `json:"factory_name"`       // factory_name
	StartTime        time.Time      `json:"start_time"`         // start_time
	EndTime          time.Time      `json:"end_time"`           // end_time
	CreatedAt        mysql.NullTime `json:"created_at"`         // created_at
	UpdatedAt        mysql.NullTime `json:"updated_at"`         // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryActivity exists in the database.
func (fa *FactoryActivity) Exists() bool { //factory_activity
	return fa._exists
}

// Deleted provides information if the FactoryActivity has been deleted from the database.
func (fa *FactoryActivity) Deleted() bool {
	return fa._deleted
}

// Get table name
func GetFactoryActivityTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "factory_activity", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryActivity to the database.
func (fa *FactoryActivity) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if fa._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryActivityTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`name, activity_type_id, activity_type_code, content, activity_image, pic_id, status, activity_status, reply_content, factory_id, sort, dis_tag, factory_name, start_time, end_time, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fa.Name, fa.ActivityTypeID, fa.ActivityTypeCode, fa.Content, fa.ActivityImage, fa.PicID, fa.Status, fa.ActivityStatus, fa.ReplyContent, fa.FactoryID, fa.Sort, fa.DisTag, fa.FactoryName, fa.StartTime, fa.EndTime, fa.CreatedAt, fa.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, fa.Name, fa.ActivityTypeID, fa.ActivityTypeCode, fa.Content, fa.ActivityImage, fa.PicID, fa.Status, fa.ActivityStatus, fa.ReplyContent, fa.FactoryID, fa.Sort, fa.DisTag, fa.FactoryName, fa.StartTime, fa.EndTime, fa.CreatedAt, fa.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, fa.Name, fa.ActivityTypeID, fa.ActivityTypeCode, fa.Content, fa.ActivityImage, fa.PicID, fa.Status, fa.ActivityStatus, fa.ReplyContent, fa.FactoryID, fa.Sort, fa.DisTag, fa.FactoryName, fa.StartTime, fa.EndTime, fa.CreatedAt, fa.UpdatedAt)
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
	fa.ID = uint(id)
	fa._exists = true

	return nil
}

// Update updates the FactoryActivity in the database.
func (fa *FactoryActivity) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fa._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryActivityTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`name = ?, activity_type_id = ?, activity_type_code = ?, content = ?, activity_image = ?, pic_id = ?, status = ?, activity_status = ?, reply_content = ?, factory_id = ?, sort = ?, dis_tag = ?, factory_name = ?, start_time = ?, end_time = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fa.Name, fa.ActivityTypeID, fa.ActivityTypeCode, fa.Content, fa.ActivityImage, fa.PicID, fa.Status, fa.ActivityStatus, fa.ReplyContent, fa.FactoryID, fa.Sort, fa.DisTag, fa.FactoryName, fa.StartTime, fa.EndTime, fa.CreatedAt, fa.UpdatedAt, fa.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fa.Name, fa.ActivityTypeID, fa.ActivityTypeCode, fa.Content, fa.ActivityImage, fa.PicID, fa.Status, fa.ActivityStatus, fa.ReplyContent, fa.FactoryID, fa.Sort, fa.DisTag, fa.FactoryName, fa.StartTime, fa.EndTime, fa.CreatedAt, fa.UpdatedAt, fa.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, fa.Name, fa.ActivityTypeID, fa.ActivityTypeCode, fa.Content, fa.ActivityImage, fa.PicID, fa.Status, fa.ActivityStatus, fa.ReplyContent, fa.FactoryID, fa.Sort, fa.DisTag, fa.FactoryName, fa.StartTime, fa.EndTime, fa.CreatedAt, fa.UpdatedAt, fa.ID)
	}
	return err
}

// Save saves the FactoryActivity to the database.
func (fa *FactoryActivity) Save(ctx context.Context) error {
	if fa.Exists() {
		return fa.Update(ctx)
	}

	return fa.Insert(ctx)
}

// Delete deletes the FactoryActivity from the database.
func (fa *FactoryActivity) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fa._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryActivityTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fa.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fa.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, fa.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	fa._deleted = true

	return nil
}

// FactoryActivitiesByActivityStatus retrieves a row from 'aypcddg.factory_activity' as a FactoryActivity.
//
// Generated from index 'factory_activity_activity_status_index'.
func FactoryActivitiesByActivityStatus(ctx context.Context, activityStatus int, key ...interface{}) ([]*FactoryActivity, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryActivityTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, activity_type_id, activity_type_code, content, activity_image, pic_id, status, activity_status, reply_content, factory_id, sort, dis_tag, factory_name, start_time, end_time, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE activity_status = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, activityStatus)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, activityStatus)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, activityStatus)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*FactoryActivity, 0)
	for queryData.Next() {
		fa := FactoryActivity{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&fa.ID, &fa.Name, &fa.ActivityTypeID, &fa.ActivityTypeCode, &fa.Content, &fa.ActivityImage, &fa.PicID, &fa.Status, &fa.ActivityStatus, &fa.ReplyContent, &fa.FactoryID, &fa.Sort, &fa.DisTag, &fa.FactoryName, &fa.StartTime, &fa.EndTime, &fa.CreatedAt, &fa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &fa)
	}

	return res, nil
}

// FactoryActivitiesByFactoryID retrieves a row from 'aypcddg.factory_activity' as a FactoryActivity.
//
// Generated from index 'factory_activity_factory_id_index'.
func FactoryActivitiesByFactoryID(ctx context.Context, factoryID int, key ...interface{}) ([]*FactoryActivity, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryActivityTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, activity_type_id, activity_type_code, content, activity_image, pic_id, status, activity_status, reply_content, factory_id, sort, dis_tag, factory_name, start_time, end_time, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE factory_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, factoryID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, factoryID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, factoryID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*FactoryActivity, 0)
	for queryData.Next() {
		fa := FactoryActivity{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&fa.ID, &fa.Name, &fa.ActivityTypeID, &fa.ActivityTypeCode, &fa.Content, &fa.ActivityImage, &fa.PicID, &fa.Status, &fa.ActivityStatus, &fa.ReplyContent, &fa.FactoryID, &fa.Sort, &fa.DisTag, &fa.FactoryName, &fa.StartTime, &fa.EndTime, &fa.CreatedAt, &fa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &fa)
	}

	return res, nil
}

// FactoryActivityByID retrieves a row from 'aypcddg.factory_activity' as a FactoryActivity.
//
// Generated from index 'factory_activity_id_pkey'.
func FactoryActivityByID(ctx context.Context, id uint, key ...interface{}) (*FactoryActivity, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryActivityTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, activity_type_id, activity_type_code, content, activity_image, pic_id, status, activity_status, reply_content, factory_id, sort, dis_tag, factory_name, start_time, end_time, created_at, updated_at ` +
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
	fa := FactoryActivity{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&fa.ID, &fa.Name, &fa.ActivityTypeID, &fa.ActivityTypeCode, &fa.Content, &fa.ActivityImage, &fa.PicID, &fa.Status, &fa.ActivityStatus, &fa.ReplyContent, &fa.FactoryID, &fa.Sort, &fa.DisTag, &fa.FactoryName, &fa.StartTime, &fa.EndTime, &fa.CreatedAt, &fa.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&fa.ID, &fa.Name, &fa.ActivityTypeID, &fa.ActivityTypeCode, &fa.Content, &fa.ActivityImage, &fa.PicID, &fa.Status, &fa.ActivityStatus, &fa.ReplyContent, &fa.FactoryID, &fa.Sort, &fa.DisTag, &fa.FactoryName, &fa.StartTime, &fa.EndTime, &fa.CreatedAt, &fa.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &fa, nil
}

// FactoryActivitiesByStatus retrieves a row from 'aypcddg.factory_activity' as a FactoryActivity.
//
// Generated from index 'factory_activity_status_index'.
func FactoryActivitiesByStatus(ctx context.Context, status int, key ...interface{}) ([]*FactoryActivity, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryActivityTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, activity_type_id, activity_type_code, content, activity_image, pic_id, status, activity_status, reply_content, factory_id, sort, dis_tag, factory_name, start_time, end_time, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE status = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, status)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*FactoryActivity, 0)
	for queryData.Next() {
		fa := FactoryActivity{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&fa.ID, &fa.Name, &fa.ActivityTypeID, &fa.ActivityTypeCode, &fa.Content, &fa.ActivityImage, &fa.PicID, &fa.Status, &fa.ActivityStatus, &fa.ReplyContent, &fa.FactoryID, &fa.Sort, &fa.DisTag, &fa.FactoryName, &fa.StartTime, &fa.EndTime, &fa.CreatedAt, &fa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &fa)
	}

	return res, nil
}
