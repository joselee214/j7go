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

	"go.uber.org/zap"
)

// SaleplanPlan represents a row from 'aypcddg.saleplan_plans'.
type SaleplanPlan struct {
	ID             uint64 `json:"id"`              // id
	Fid            int    `json:"fid"`             // fid
	CreateUID      int    `json:"create_uid"`      // create_uid
	CreateAt       int    `json:"create_at"`       // create_at
	PlanType       int8   `json:"plan_type"`       // plan_type
	Name           string `json:"name"`            // name
	Note           string `json:"note"`            // note
	StartTime      int    `json:"start_time"`      // start_time
	EndTime        int    `json:"end_time"`        // end_time
	Status         int8   `json:"status"`          // status
	GoodsNum       int    `json:"goods_num"`       // goods_num
	PlatformLock   int8   `json:"platform_lock"`   // platform_lock
	PlatformReport int8   `json:"platform_report"` // platform_report

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the SaleplanPlan exists in the database.
func (sp *SaleplanPlan) Exists() bool { //saleplan_plans
	return sp._exists
}

// Deleted provides information if the SaleplanPlan has been deleted from the database.
func (sp *SaleplanPlan) Deleted() bool {
	return sp._deleted
}

// Get table name
func GetSaleplanPlanTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "saleplan_plans", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the SaleplanPlan to the database.
func (sp *SaleplanPlan) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if sp._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSaleplanPlanTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`fid, create_uid, create_at, plan_type, name, note, start_time, end_time, status, goods_num, platform_lock, platform_report` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, sp.Fid, sp.CreateUID, sp.CreateAt, sp.PlanType, sp.Name, sp.Note, sp.StartTime, sp.EndTime, sp.Status, sp.GoodsNum, sp.PlatformLock, sp.PlatformReport)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, sp.Fid, sp.CreateUID, sp.CreateAt, sp.PlanType, sp.Name, sp.Note, sp.StartTime, sp.EndTime, sp.Status, sp.GoodsNum, sp.PlatformLock, sp.PlatformReport)
	} else {
		res, err = dbConn.Exec(sqlstr, sp.Fid, sp.CreateUID, sp.CreateAt, sp.PlanType, sp.Name, sp.Note, sp.StartTime, sp.EndTime, sp.Status, sp.GoodsNum, sp.PlatformLock, sp.PlatformReport)
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
	sp.ID = uint64(id)
	sp._exists = true

	return nil
}

// Update updates the SaleplanPlan in the database.
func (sp *SaleplanPlan) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if sp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSaleplanPlanTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`fid = ?, create_uid = ?, create_at = ?, plan_type = ?, name = ?, note = ?, start_time = ?, end_time = ?, status = ?, goods_num = ?, platform_lock = ?, platform_report = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, sp.Fid, sp.CreateUID, sp.CreateAt, sp.PlanType, sp.Name, sp.Note, sp.StartTime, sp.EndTime, sp.Status, sp.GoodsNum, sp.PlatformLock, sp.PlatformReport, sp.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, sp.Fid, sp.CreateUID, sp.CreateAt, sp.PlanType, sp.Name, sp.Note, sp.StartTime, sp.EndTime, sp.Status, sp.GoodsNum, sp.PlatformLock, sp.PlatformReport, sp.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, sp.Fid, sp.CreateUID, sp.CreateAt, sp.PlanType, sp.Name, sp.Note, sp.StartTime, sp.EndTime, sp.Status, sp.GoodsNum, sp.PlatformLock, sp.PlatformReport, sp.ID)
	}
	return err
}

// Save saves the SaleplanPlan to the database.
func (sp *SaleplanPlan) Save(ctx context.Context) error {
	if sp.Exists() {
		return sp.Update(ctx)
	}

	return sp.Insert(ctx)
}

// Delete deletes the SaleplanPlan from the database.
func (sp *SaleplanPlan) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if sp._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetSaleplanPlanTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, sp.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, sp.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, sp.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	sp._deleted = true

	return nil
}

// SaleplanPlansByEndTime retrieves a row from 'aypcddg.saleplan_plans' as a SaleplanPlan.
//
// Generated from index 'saleplan_plans_end_time_index'.
func SaleplanPlansByEndTime(ctx context.Context, endTime int, key ...interface{}) ([]*SaleplanPlan, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetSaleplanPlanTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, fid, create_uid, create_at, plan_type, name, note, start_time, end_time, status, goods_num, platform_lock, platform_report ` +
		`FROM ` + tableName +
		` WHERE end_time = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, endTime)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, endTime)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, endTime)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*SaleplanPlan, 0)
	for queryData.Next() {
		sp := SaleplanPlan{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&sp.ID, &sp.Fid, &sp.CreateUID, &sp.CreateAt, &sp.PlanType, &sp.Name, &sp.Note, &sp.StartTime, &sp.EndTime, &sp.Status, &sp.GoodsNum, &sp.PlatformLock, &sp.PlatformReport)
		if err != nil {
			return nil, err
		}

		res = append(res, &sp)
	}

	return res, nil
}

// SaleplanPlansByFid retrieves a row from 'aypcddg.saleplan_plans' as a SaleplanPlan.
//
// Generated from index 'saleplan_plans_fid_index'.
func SaleplanPlansByFid(ctx context.Context, fid int, key ...interface{}) ([]*SaleplanPlan, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetSaleplanPlanTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, fid, create_uid, create_at, plan_type, name, note, start_time, end_time, status, goods_num, platform_lock, platform_report ` +
		`FROM ` + tableName +
		` WHERE fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, fid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, fid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*SaleplanPlan, 0)
	for queryData.Next() {
		sp := SaleplanPlan{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&sp.ID, &sp.Fid, &sp.CreateUID, &sp.CreateAt, &sp.PlanType, &sp.Name, &sp.Note, &sp.StartTime, &sp.EndTime, &sp.Status, &sp.GoodsNum, &sp.PlatformLock, &sp.PlatformReport)
		if err != nil {
			return nil, err
		}

		res = append(res, &sp)
	}

	return res, nil
}

// SaleplanPlanByID retrieves a row from 'aypcddg.saleplan_plans' as a SaleplanPlan.
//
// Generated from index 'saleplan_plans_id_pkey'.
func SaleplanPlanByID(ctx context.Context, id uint64, key ...interface{}) (*SaleplanPlan, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetSaleplanPlanTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, fid, create_uid, create_at, plan_type, name, note, start_time, end_time, status, goods_num, platform_lock, platform_report ` +
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
	sp := SaleplanPlan{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&sp.ID, &sp.Fid, &sp.CreateUID, &sp.CreateAt, &sp.PlanType, &sp.Name, &sp.Note, &sp.StartTime, &sp.EndTime, &sp.Status, &sp.GoodsNum, &sp.PlatformLock, &sp.PlatformReport)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&sp.ID, &sp.Fid, &sp.CreateUID, &sp.CreateAt, &sp.PlanType, &sp.Name, &sp.Note, &sp.StartTime, &sp.EndTime, &sp.Status, &sp.GoodsNum, &sp.PlatformLock, &sp.PlatformReport)
		if err != nil {
			return nil, err
		}
	}

	return &sp, nil
}

// SaleplanPlansByStartTime retrieves a row from 'aypcddg.saleplan_plans' as a SaleplanPlan.
//
// Generated from index 'saleplan_plans_start_time_index'.
func SaleplanPlansByStartTime(ctx context.Context, startTime int, key ...interface{}) ([]*SaleplanPlan, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetSaleplanPlanTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, fid, create_uid, create_at, plan_type, name, note, start_time, end_time, status, goods_num, platform_lock, platform_report ` +
		`FROM ` + tableName +
		` WHERE start_time = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, startTime)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, startTime)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, startTime)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*SaleplanPlan, 0)
	for queryData.Next() {
		sp := SaleplanPlan{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&sp.ID, &sp.Fid, &sp.CreateUID, &sp.CreateAt, &sp.PlanType, &sp.Name, &sp.Note, &sp.StartTime, &sp.EndTime, &sp.Status, &sp.GoodsNum, &sp.PlatformLock, &sp.PlatformReport)
		if err != nil {
			return nil, err
		}

		res = append(res, &sp)
	}

	return res, nil
}

// SaleplanPlansByStatus retrieves a row from 'aypcddg.saleplan_plans' as a SaleplanPlan.
//
// Generated from index 'saleplan_plans_status_index'.
func SaleplanPlansByStatus(ctx context.Context, status int8, key ...interface{}) ([]*SaleplanPlan, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetSaleplanPlanTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, fid, create_uid, create_at, plan_type, name, note, start_time, end_time, status, goods_num, platform_lock, platform_report ` +
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
	res := make([]*SaleplanPlan, 0)
	for queryData.Next() {
		sp := SaleplanPlan{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&sp.ID, &sp.Fid, &sp.CreateUID, &sp.CreateAt, &sp.PlanType, &sp.Name, &sp.Note, &sp.StartTime, &sp.EndTime, &sp.Status, &sp.GoodsNum, &sp.PlatformLock, &sp.PlatformReport)
		if err != nil {
			return nil, err
		}

		res = append(res, &sp)
	}

	return res, nil
}
