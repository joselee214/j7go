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

// OrderDeliver represents a row from 'aypcddg.order_deliver'.
type OrderDeliver struct {
	OrderDid           int            `json:"order_did"`            // order_did
	Orderid            int            `json:"orderid"`              // orderid
	Fid                sql.NullInt64  `json:"fid"`                  // fid
	Sid                sql.NullInt64  `json:"sid"`                  // sid
	UID                int            `json:"uid"`                  // uid
	Created            int            `json:"created"`              // created
	Note               sql.NullString `json:"note"`                 // note
	DeliverType        int            `json:"deliver_type"`         // deliver_type
	DeliverCompanyID   sql.NullInt64  `json:"deliver_company_id"`   // deliver_company_id
	DeliverCompanyName sql.NullString `json:"deliver_company_name"` // deliver_company_name
	DeliverSn          sql.NullString `json:"deliver_sn"`           // deliver_sn
	DeliverUserID      sql.NullInt64  `json:"deliver_user_id"`      // deliver_user_id
	DeliverUserName    sql.NullString `json:"deliver_user_name"`    // deliver_user_name
	DeliverUserPhone   sql.NullString `json:"deliver_user_phone"`   // deliver_user_phone
	DeliverSendTime    mysql.NullTime `json:"deliver_send_time"`    // deliver_send_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OrderDeliver exists in the database.
func (od *OrderDeliver) Exists() bool { //order_deliver
	return od._exists
}

// Deleted provides information if the OrderDeliver has been deleted from the database.
func (od *OrderDeliver) Deleted() bool {
	return od._deleted
}

// Get table name
func GetOrderDeliverTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "order_deliver", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the OrderDeliver to the database.
func (od *OrderDeliver) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if od._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderDeliverTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`orderid, fid, sid, uid, created, note, deliver_type, deliver_company_id, deliver_company_name, deliver_sn, deliver_user_id, deliver_user_name, deliver_user_phone, deliver_send_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, od.Orderid, od.Fid, od.Sid, od.UID, od.Created, od.Note, od.DeliverType, od.DeliverCompanyID, od.DeliverCompanyName, od.DeliverSn, od.DeliverUserID, od.DeliverUserName, od.DeliverUserPhone, od.DeliverSendTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, od.Orderid, od.Fid, od.Sid, od.UID, od.Created, od.Note, od.DeliverType, od.DeliverCompanyID, od.DeliverCompanyName, od.DeliverSn, od.DeliverUserID, od.DeliverUserName, od.DeliverUserPhone, od.DeliverSendTime)
	} else {
		res, err = dbConn.Exec(sqlstr, od.Orderid, od.Fid, od.Sid, od.UID, od.Created, od.Note, od.DeliverType, od.DeliverCompanyID, od.DeliverCompanyName, od.DeliverSn, od.DeliverUserID, od.DeliverUserName, od.DeliverUserPhone, od.DeliverSendTime)
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
	od.OrderDid = int(id)
	od._exists = true

	return nil
}

// Update updates the OrderDeliver in the database.
func (od *OrderDeliver) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if od._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderDeliverTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`orderid = ?, fid = ?, sid = ?, uid = ?, created = ?, note = ?, deliver_type = ?, deliver_company_id = ?, deliver_company_name = ?, deliver_sn = ?, deliver_user_id = ?, deliver_user_name = ?, deliver_user_phone = ?, deliver_send_time = ?` +
		` WHERE order_did = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, od.Orderid, od.Fid, od.Sid, od.UID, od.Created, od.Note, od.DeliverType, od.DeliverCompanyID, od.DeliverCompanyName, od.DeliverSn, od.DeliverUserID, od.DeliverUserName, od.DeliverUserPhone, od.DeliverSendTime, od.OrderDid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, od.Orderid, od.Fid, od.Sid, od.UID, od.Created, od.Note, od.DeliverType, od.DeliverCompanyID, od.DeliverCompanyName, od.DeliverSn, od.DeliverUserID, od.DeliverUserName, od.DeliverUserPhone, od.DeliverSendTime, od.OrderDid)
	} else {
		_, err = dbConn.Exec(sqlstr, od.Orderid, od.Fid, od.Sid, od.UID, od.Created, od.Note, od.DeliverType, od.DeliverCompanyID, od.DeliverCompanyName, od.DeliverSn, od.DeliverUserID, od.DeliverUserName, od.DeliverUserPhone, od.DeliverSendTime, od.OrderDid)
	}
	return err
}

// Save saves the OrderDeliver to the database.
func (od *OrderDeliver) Save(ctx context.Context) error {
	if od.Exists() {
		return od.Update(ctx)
	}

	return od.Insert(ctx)
}

// Delete deletes the OrderDeliver from the database.
func (od *OrderDeliver) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if od._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderDeliverTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE order_did = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, od.OrderDid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, od.OrderDid)
	} else {
		_, err = dbConn.Exec(sqlstr, od.OrderDid)
	}

	if err != nil {
		return err
	}

	// set deleted
	od._deleted = true

	return nil
}

// OrderDeliversByOrderid retrieves a row from 'aypcddg.order_deliver' as a OrderDeliver.
//
// Generated from index 'oid'.
func OrderDeliversByOrderid(ctx context.Context, orderid int, key ...interface{}) ([]*OrderDeliver, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetOrderDeliverTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`order_did, orderid, fid, sid, uid, created, note, deliver_type, deliver_company_id, deliver_company_name, deliver_sn, deliver_user_id, deliver_user_name, deliver_user_phone, deliver_send_time ` +
		`FROM ` + tableName +
		` WHERE orderid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, orderid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, orderid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, orderid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*OrderDeliver, 0)
	for queryData.Next() {
		od := OrderDeliver{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&od.OrderDid, &od.Orderid, &od.Fid, &od.Sid, &od.UID, &od.Created, &od.Note, &od.DeliverType, &od.DeliverCompanyID, &od.DeliverCompanyName, &od.DeliverSn, &od.DeliverUserID, &od.DeliverUserName, &od.DeliverUserPhone, &od.DeliverSendTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &od)
	}

	return res, nil
}

// OrderDeliverByOrderDid retrieves a row from 'aypcddg.order_deliver' as a OrderDeliver.
//
// Generated from index 'order_deliver_order_did_pkey'.
func OrderDeliverByOrderDid(ctx context.Context, orderDid int, key ...interface{}) (*OrderDeliver, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetOrderDeliverTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`order_did, orderid, fid, sid, uid, created, note, deliver_type, deliver_company_id, deliver_company_name, deliver_sn, deliver_user_id, deliver_user_name, deliver_user_phone, deliver_send_time ` +
		`FROM ` + tableName +
		` WHERE order_did = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, orderDid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	od := OrderDeliver{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, orderDid).Scan(&od.OrderDid, &od.Orderid, &od.Fid, &od.Sid, &od.UID, &od.Created, &od.Note, &od.DeliverType, &od.DeliverCompanyID, &od.DeliverCompanyName, &od.DeliverSn, &od.DeliverUserID, &od.DeliverUserName, &od.DeliverUserPhone, &od.DeliverSendTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, orderDid).Scan(&od.OrderDid, &od.Orderid, &od.Fid, &od.Sid, &od.UID, &od.Created, &od.Note, &od.DeliverType, &od.DeliverCompanyID, &od.DeliverCompanyName, &od.DeliverSn, &od.DeliverUserID, &od.DeliverUserName, &od.DeliverUserPhone, &od.DeliverSendTime)
		if err != nil {
			return nil, err
		}
	}

	return &od, nil
}
