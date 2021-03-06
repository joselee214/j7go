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

// FactoryFaretp represents a row from 'aypcddg.factory_faretp'.
type FactoryFaretp struct {
	FareTid        int            `json:"fare_tid"`         // fare_tid
	FareTname      sql.NullString `json:"fare_tname"`       // fare_tname
	FareDesp       sql.NullString `json:"fare_desp"`        // fare_desp
	Fid            sql.NullInt64  `json:"fid"`              // fid
	Sid            sql.NullInt64  `json:"sid"`              // sid
	FareType       int            `json:"fare_type"`        // fare_type
	FareFirstNum   int            `json:"fare_first_num"`   // fare_first_num
	FareFirst      float64        `json:"fare_first"`       // fare_first
	FareAddNum     int            `json:"fare_add_num"`     // fare_add_num
	FareAdd        float64        `json:"fare_add"`         // fare_add
	FareDelivery   sql.NullString `json:"fare_delivery"`    // fare_delivery
	FareDeliveryTp int            `json:"fare_delivery_tp"` // fare_delivery_tp
	FareAreas      sql.NullString `json:"fare_areas"`       // fare_areas

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryFaretp exists in the database.
func (ff *FactoryFaretp) Exists() bool { //factory_faretp
	return ff._exists
}

// Deleted provides information if the FactoryFaretp has been deleted from the database.
func (ff *FactoryFaretp) Deleted() bool {
	return ff._deleted
}

// Get table name
func GetFactoryFaretpTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "factory_faretp", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryFaretp to the database.
func (ff *FactoryFaretp) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ff._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryFaretpTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`fare_tname, fare_desp, fid, sid, fare_type, fare_first_num, fare_first, fare_add_num, fare_add, fare_delivery, fare_delivery_tp, fare_areas` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ff.FareTname, ff.FareDesp, ff.Fid, ff.Sid, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareAreas)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ff.FareTname, ff.FareDesp, ff.Fid, ff.Sid, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareAreas)
	} else {
		res, err = dbConn.Exec(sqlstr, ff.FareTname, ff.FareDesp, ff.Fid, ff.Sid, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareAreas)
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
	ff.FareTid = int(id)
	ff._exists = true

	return nil
}

// Update updates the FactoryFaretp in the database.
func (ff *FactoryFaretp) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ff._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryFaretpTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`fare_tname = ?, fare_desp = ?, fid = ?, sid = ?, fare_type = ?, fare_first_num = ?, fare_first = ?, fare_add_num = ?, fare_add = ?, fare_delivery = ?, fare_delivery_tp = ?, fare_areas = ?` +
		` WHERE fare_tid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ff.FareTname, ff.FareDesp, ff.Fid, ff.Sid, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareAreas, ff.FareTid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ff.FareTname, ff.FareDesp, ff.Fid, ff.Sid, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareAreas, ff.FareTid)
	} else {
		_, err = dbConn.Exec(sqlstr, ff.FareTname, ff.FareDesp, ff.Fid, ff.Sid, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareAreas, ff.FareTid)
	}
	return err
}

// Save saves the FactoryFaretp to the database.
func (ff *FactoryFaretp) Save(ctx context.Context) error {
	if ff.Exists() {
		return ff.Update(ctx)
	}

	return ff.Insert(ctx)
}

// Delete deletes the FactoryFaretp from the database.
func (ff *FactoryFaretp) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ff._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryFaretpTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fare_tid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ff.FareTid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ff.FareTid)
	} else {
		_, err = dbConn.Exec(sqlstr, ff.FareTid)
	}

	if err != nil {
		return err
	}

	// set deleted
	ff._deleted = true

	return nil
}

// FactoryFaretpByFareTid retrieves a row from 'aypcddg.factory_faretp' as a FactoryFaretp.
//
// Generated from index 'factory_faretp_fare_tid_pkey'.
func FactoryFaretpByFareTid(ctx context.Context, fareTid int, key ...interface{}) (*FactoryFaretp, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryFaretpTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`fare_tid, fare_tname, fare_desp, fid, sid, fare_type, fare_first_num, fare_first, fare_add_num, fare_add, fare_delivery, fare_delivery_tp, fare_areas ` +
		`FROM ` + tableName +
		` WHERE fare_tid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fareTid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ff := FactoryFaretp{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fareTid).Scan(&ff.FareTid, &ff.FareTname, &ff.FareDesp, &ff.Fid, &ff.Sid, &ff.FareType, &ff.FareFirstNum, &ff.FareFirst, &ff.FareAddNum, &ff.FareAdd, &ff.FareDelivery, &ff.FareDeliveryTp, &ff.FareAreas)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fareTid).Scan(&ff.FareTid, &ff.FareTname, &ff.FareDesp, &ff.Fid, &ff.Sid, &ff.FareType, &ff.FareFirstNum, &ff.FareFirst, &ff.FareAddNum, &ff.FareAdd, &ff.FareDelivery, &ff.FareDeliveryTp, &ff.FareAreas)
		if err != nil {
			return nil, err
		}
	}

	return &ff, nil
}
