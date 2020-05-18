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

// MidAd represents a row from 'aypcddg.mid_ad'.
type MidAd struct {
	ID              uint64         `json:"id"`                 // id
	MidAdPositionID int            `json:"mid_ad_position_id"` // mid_ad_position_id
	Type            string         `json:"type"`               // type
	Title           sql.NullString `json:"title"`              // title
	JumpType        sql.NullString `json:"jump_type"`          // jump_type
	ParamsID        sql.NullInt64  `json:"params_id"`          // params_id
	JumpParams      sql.NullString `json:"jump_params"`        // jump_params
	MidAdImage      sql.NullString `json:"mid_ad_image"`       // mid_ad_image
	ImageColor      string         `json:"image_color"`        // image_color
	StartAt         mysql.NullTime `json:"start_at"`           // start_at
	EndAt           mysql.NullTime `json:"end_at"`             // end_at
	Sort            int            `json:"sort"`               // sort
	Status          int8           `json:"status"`             // status
	Logo            sql.NullString `json:"logo"`               // logo
	Desc            sql.NullString `json:"desc"`               // desc
	Tag             sql.NullString `json:"tag"`                // tag
	ShowType        int8           `json:"show_type"`          // show_type
	MidAdminID      int            `json:"mid_admin_id"`       // mid_admin_id
	MidAdminName    sql.NullString `json:"mid_admin_name"`     // mid_admin_name
	CreatedAt       mysql.NullTime `json:"created_at"`         // created_at
	UpdatedAt       mysql.NullTime `json:"updated_at"`         // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MidAd exists in the database.
func (ma *MidAd) Exists() bool { //mid_ad
	return ma._exists
}

// Deleted provides information if the MidAd has been deleted from the database.
func (ma *MidAd) Deleted() bool {
	return ma._deleted
}

// Get table name
func GetMidAdTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "mid_ad", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the MidAd to the database.
func (ma *MidAd) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ma._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`mid_ad_position_id, type, title, jump_type, params_id, jump_params, mid_ad_image, image_color, start_at, end_at, sort, status, logo, desc, tag, show_type, mid_admin_id, mid_admin_name, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ma.MidAdPositionID, ma.Type, ma.Title, ma.JumpType, ma.ParamsID, ma.JumpParams, ma.MidAdImage, ma.ImageColor, ma.StartAt, ma.EndAt, ma.Sort, ma.Status, ma.Logo, ma.Desc, ma.Tag, ma.ShowType, ma.MidAdminID, ma.MidAdminName, ma.CreatedAt, ma.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ma.MidAdPositionID, ma.Type, ma.Title, ma.JumpType, ma.ParamsID, ma.JumpParams, ma.MidAdImage, ma.ImageColor, ma.StartAt, ma.EndAt, ma.Sort, ma.Status, ma.Logo, ma.Desc, ma.Tag, ma.ShowType, ma.MidAdminID, ma.MidAdminName, ma.CreatedAt, ma.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, ma.MidAdPositionID, ma.Type, ma.Title, ma.JumpType, ma.ParamsID, ma.JumpParams, ma.MidAdImage, ma.ImageColor, ma.StartAt, ma.EndAt, ma.Sort, ma.Status, ma.Logo, ma.Desc, ma.Tag, ma.ShowType, ma.MidAdminID, ma.MidAdminName, ma.CreatedAt, ma.UpdatedAt)
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
	ma.ID = uint64(id)
	ma._exists = true

	return nil
}

// Update updates the MidAd in the database.
func (ma *MidAd) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ma._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`mid_ad_position_id = ?, type = ?, title = ?, jump_type = ?, params_id = ?, jump_params = ?, mid_ad_image = ?, image_color = ?, start_at = ?, end_at = ?, sort = ?, status = ?, logo = ?, desc = ?, tag = ?, show_type = ?, mid_admin_id = ?, mid_admin_name = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ma.MidAdPositionID, ma.Type, ma.Title, ma.JumpType, ma.ParamsID, ma.JumpParams, ma.MidAdImage, ma.ImageColor, ma.StartAt, ma.EndAt, ma.Sort, ma.Status, ma.Logo, ma.Desc, ma.Tag, ma.ShowType, ma.MidAdminID, ma.MidAdminName, ma.CreatedAt, ma.UpdatedAt, ma.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ma.MidAdPositionID, ma.Type, ma.Title, ma.JumpType, ma.ParamsID, ma.JumpParams, ma.MidAdImage, ma.ImageColor, ma.StartAt, ma.EndAt, ma.Sort, ma.Status, ma.Logo, ma.Desc, ma.Tag, ma.ShowType, ma.MidAdminID, ma.MidAdminName, ma.CreatedAt, ma.UpdatedAt, ma.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ma.MidAdPositionID, ma.Type, ma.Title, ma.JumpType, ma.ParamsID, ma.JumpParams, ma.MidAdImage, ma.ImageColor, ma.StartAt, ma.EndAt, ma.Sort, ma.Status, ma.Logo, ma.Desc, ma.Tag, ma.ShowType, ma.MidAdminID, ma.MidAdminName, ma.CreatedAt, ma.UpdatedAt, ma.ID)
	}
	return err
}

// Save saves the MidAd to the database.
func (ma *MidAd) Save(ctx context.Context) error {
	if ma.Exists() {
		return ma.Update(ctx)
	}

	return ma.Insert(ctx)
}

// Delete deletes the MidAd from the database.
func (ma *MidAd) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ma._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ma.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ma.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ma.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ma._deleted = true

	return nil
}

// MidAdsByMidAdPositionID retrieves a row from 'aypcddg.mid_ad' as a MidAd.
//
// Generated from index 'idx_position_id'.
func MidAdsByMidAdPositionID(ctx context.Context, midAdPositionID int, key ...interface{}) ([]*MidAd, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, mid_ad_position_id, type, title, jump_type, params_id, jump_params, mid_ad_image, image_color, start_at, end_at, sort, status, logo, desc, tag, show_type, mid_admin_id, mid_admin_name, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE mid_ad_position_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, midAdPositionID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, midAdPositionID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, midAdPositionID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*MidAd, 0)
	for queryData.Next() {
		ma := MidAd{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ma.ID, &ma.MidAdPositionID, &ma.Type, &ma.Title, &ma.JumpType, &ma.ParamsID, &ma.JumpParams, &ma.MidAdImage, &ma.ImageColor, &ma.StartAt, &ma.EndAt, &ma.Sort, &ma.Status, &ma.Logo, &ma.Desc, &ma.Tag, &ma.ShowType, &ma.MidAdminID, &ma.MidAdminName, &ma.CreatedAt, &ma.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &ma)
	}

	return res, nil
}

// MidAdByID retrieves a row from 'aypcddg.mid_ad' as a MidAd.
//
// Generated from index 'mid_ad_id_pkey'.
func MidAdByID(ctx context.Context, id uint64, key ...interface{}) (*MidAd, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, mid_ad_position_id, type, title, jump_type, params_id, jump_params, mid_ad_image, image_color, start_at, end_at, sort, status, logo, desc, tag, show_type, mid_admin_id, mid_admin_name, created_at, updated_at ` +
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
	ma := MidAd{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&ma.ID, &ma.MidAdPositionID, &ma.Type, &ma.Title, &ma.JumpType, &ma.ParamsID, &ma.JumpParams, &ma.MidAdImage, &ma.ImageColor, &ma.StartAt, &ma.EndAt, &ma.Sort, &ma.Status, &ma.Logo, &ma.Desc, &ma.Tag, &ma.ShowType, &ma.MidAdminID, &ma.MidAdminName, &ma.CreatedAt, &ma.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&ma.ID, &ma.MidAdPositionID, &ma.Type, &ma.Title, &ma.JumpType, &ma.ParamsID, &ma.JumpParams, &ma.MidAdImage, &ma.ImageColor, &ma.StartAt, &ma.EndAt, &ma.Sort, &ma.Status, &ma.Logo, &ma.Desc, &ma.Tag, &ma.ShowType, &ma.MidAdminID, &ma.MidAdminName, &ma.CreatedAt, &ma.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &ma, nil
}
