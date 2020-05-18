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

// Order represents a row from 'aypcddg.orders'.
type Order struct {
	Orderid          int             `json:"orderid"`           // orderid
	Q                sql.NullString  `json:"q"`                 // q
	UID              int             `json:"uid"`               // uid
	Name             sql.NullString  `json:"name"`              // name
	Mobile           sql.NullString  `json:"mobile"`            // mobile
	Email            sql.NullString  `json:"email"`             // email
	Fid              int             `json:"fid"`               // fid
	Sid              int             `json:"sid"`               // sid
	Status           int             `json:"status"`            // status
	StatusExt        sql.NullInt64   `json:"status_ext"`        // status_ext
	Created          uint            `json:"created"`           // created
	PayTime          sql.NullInt64   `json:"pay_time"`          // pay_time
	ShipTime         sql.NullInt64   `json:"ship_time"`         // ship_time
	TotalFee         float64         `json:"total_fee"`         // total_fee
	NeedFee          float64         `json:"need_fee"`          // need_fee
	SettlementPaid   float64         `json:"settlement_paid"`   // settlement_paid
	TotalPaid        float64         `json:"total_paid"`        // total_paid
	PayMethod        int8            `json:"pay_method"`        // pay_method
	Shipping         float64         `json:"shipping"`          // shipping
	GoodsFee         float64         `json:"goods_fee"`         // goods_fee
	DiscountFee      float64         `json:"discount_fee"`      // discount_fee
	LogisticType     int8            `json:"logistic_type"`     // logistic_type
	Pid              int             `json:"pid"`               // pid
	Cid              int             `json:"cid"`               // cid
	Aid              int             `json:"aid"`               // aid
	Address          sql.NullString  `json:"address"`           // address
	AddressInfo      sql.NullString  `json:"address_info"`      // address_info
	InvoiceInfo      sql.NullString  `json:"invoice_info"`      // invoice_info
	Slinfo           sql.NullString  `json:"slinfo"`            // slinfo
	Commission       float64         `json:"commission"`        // commission
	Prepay           sql.NullFloat64 `json:"prepay"`            // prepay
	AffirmTime       sql.NullInt64   `json:"affirm_time"`       // affirm_time
	NoteSale         sql.NullString  `json:"note_sale"`         // note_sale
	NoteUser         sql.NullString  `json:"note_user"`         // note_user
	RefundFee        float64         `json:"refund_fee"`        // refund_fee
	SettlementStatus int8            `json:"settlement_status"` // settlement_status
	SettlementTime   int             `json:"settlement_time"`   // settlement_time
	FromSite         int             `json:"from_site"`         // from_site
	FinishTime       int             `json:"finish_time"`       // finish_time
	BuyUserID        sql.NullInt64   `json:"buy_user_id"`       // buy_user_id
	ActivityID       sql.NullInt64   `json:"activity_id"`       // activity_id
	OrderNo          sql.NullString  `json:"order_no"`          // order_no
	Shareid          sql.NullInt64   `json:"shareid"`           // shareid
	FromUID          sql.NullInt64   `json:"from_uid"`          // from_uid
	UserAddressID    sql.NullInt64   `json:"user_address_id"`   // user_address_id
	GoodsOriginFee   float64         `json:"goods_origin_fee"`  // goods_origin_fee

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Order exists in the database.
func (o *Order) Exists() bool { //orders
	return o._exists
}

// Deleted provides information if the Order has been deleted from the database.
func (o *Order) Deleted() bool {
	return o._deleted
}

// Get table name
func GetOrderTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "orders", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Order to the database.
func (o *Order) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`q, uid, name, mobile, email, fid, sid, status, status_ext, created, pay_time, ship_time, total_fee, need_fee, settlement_paid, total_paid, pay_method, shipping, goods_fee, discount_fee, logistic_type, pid, cid, aid, address, address_info, invoice_info, slinfo, commission, prepay, affirm_time, note_sale, note_user, refund_fee, settlement_status, settlement_time, from_site, finish_time, buy_user_id, activity_id, order_no, shareid, from_uid, user_address_id, goods_origin_fee` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, o.Q, o.UID, o.Name, o.Mobile, o.Email, o.Fid, o.Sid, o.Status, o.StatusExt, o.Created, o.PayTime, o.ShipTime, o.TotalFee, o.NeedFee, o.SettlementPaid, o.TotalPaid, o.PayMethod, o.Shipping, o.GoodsFee, o.DiscountFee, o.LogisticType, o.Pid, o.Cid, o.Aid, o.Address, o.AddressInfo, o.InvoiceInfo, o.Slinfo, o.Commission, o.Prepay, o.AffirmTime, o.NoteSale, o.NoteUser, o.RefundFee, o.SettlementStatus, o.SettlementTime, o.FromSite, o.FinishTime, o.BuyUserID, o.ActivityID, o.OrderNo, o.Shareid, o.FromUID, o.UserAddressID, o.GoodsOriginFee)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, o.Q, o.UID, o.Name, o.Mobile, o.Email, o.Fid, o.Sid, o.Status, o.StatusExt, o.Created, o.PayTime, o.ShipTime, o.TotalFee, o.NeedFee, o.SettlementPaid, o.TotalPaid, o.PayMethod, o.Shipping, o.GoodsFee, o.DiscountFee, o.LogisticType, o.Pid, o.Cid, o.Aid, o.Address, o.AddressInfo, o.InvoiceInfo, o.Slinfo, o.Commission, o.Prepay, o.AffirmTime, o.NoteSale, o.NoteUser, o.RefundFee, o.SettlementStatus, o.SettlementTime, o.FromSite, o.FinishTime, o.BuyUserID, o.ActivityID, o.OrderNo, o.Shareid, o.FromUID, o.UserAddressID, o.GoodsOriginFee)
	} else {
		res, err = dbConn.Exec(sqlstr, o.Q, o.UID, o.Name, o.Mobile, o.Email, o.Fid, o.Sid, o.Status, o.StatusExt, o.Created, o.PayTime, o.ShipTime, o.TotalFee, o.NeedFee, o.SettlementPaid, o.TotalPaid, o.PayMethod, o.Shipping, o.GoodsFee, o.DiscountFee, o.LogisticType, o.Pid, o.Cid, o.Aid, o.Address, o.AddressInfo, o.InvoiceInfo, o.Slinfo, o.Commission, o.Prepay, o.AffirmTime, o.NoteSale, o.NoteUser, o.RefundFee, o.SettlementStatus, o.SettlementTime, o.FromSite, o.FinishTime, o.BuyUserID, o.ActivityID, o.OrderNo, o.Shareid, o.FromUID, o.UserAddressID, o.GoodsOriginFee)
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
	o.Orderid = int(id)
	o._exists = true

	return nil
}

// Update updates the Order in the database.
func (o *Order) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if o._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`q = ?, uid = ?, name = ?, mobile = ?, email = ?, fid = ?, sid = ?, status = ?, status_ext = ?, created = ?, pay_time = ?, ship_time = ?, total_fee = ?, need_fee = ?, settlement_paid = ?, total_paid = ?, pay_method = ?, shipping = ?, goods_fee = ?, discount_fee = ?, logistic_type = ?, pid = ?, cid = ?, aid = ?, address = ?, address_info = ?, invoice_info = ?, slinfo = ?, commission = ?, prepay = ?, affirm_time = ?, note_sale = ?, note_user = ?, refund_fee = ?, settlement_status = ?, settlement_time = ?, from_site = ?, finish_time = ?, buy_user_id = ?, activity_id = ?, order_no = ?, shareid = ?, from_uid = ?, user_address_id = ?, goods_origin_fee = ?` +
		` WHERE orderid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, o.Q, o.UID, o.Name, o.Mobile, o.Email, o.Fid, o.Sid, o.Status, o.StatusExt, o.Created, o.PayTime, o.ShipTime, o.TotalFee, o.NeedFee, o.SettlementPaid, o.TotalPaid, o.PayMethod, o.Shipping, o.GoodsFee, o.DiscountFee, o.LogisticType, o.Pid, o.Cid, o.Aid, o.Address, o.AddressInfo, o.InvoiceInfo, o.Slinfo, o.Commission, o.Prepay, o.AffirmTime, o.NoteSale, o.NoteUser, o.RefundFee, o.SettlementStatus, o.SettlementTime, o.FromSite, o.FinishTime, o.BuyUserID, o.ActivityID, o.OrderNo, o.Shareid, o.FromUID, o.UserAddressID, o.GoodsOriginFee, o.Orderid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, o.Q, o.UID, o.Name, o.Mobile, o.Email, o.Fid, o.Sid, o.Status, o.StatusExt, o.Created, o.PayTime, o.ShipTime, o.TotalFee, o.NeedFee, o.SettlementPaid, o.TotalPaid, o.PayMethod, o.Shipping, o.GoodsFee, o.DiscountFee, o.LogisticType, o.Pid, o.Cid, o.Aid, o.Address, o.AddressInfo, o.InvoiceInfo, o.Slinfo, o.Commission, o.Prepay, o.AffirmTime, o.NoteSale, o.NoteUser, o.RefundFee, o.SettlementStatus, o.SettlementTime, o.FromSite, o.FinishTime, o.BuyUserID, o.ActivityID, o.OrderNo, o.Shareid, o.FromUID, o.UserAddressID, o.GoodsOriginFee, o.Orderid)
	} else {
		_, err = dbConn.Exec(sqlstr, o.Q, o.UID, o.Name, o.Mobile, o.Email, o.Fid, o.Sid, o.Status, o.StatusExt, o.Created, o.PayTime, o.ShipTime, o.TotalFee, o.NeedFee, o.SettlementPaid, o.TotalPaid, o.PayMethod, o.Shipping, o.GoodsFee, o.DiscountFee, o.LogisticType, o.Pid, o.Cid, o.Aid, o.Address, o.AddressInfo, o.InvoiceInfo, o.Slinfo, o.Commission, o.Prepay, o.AffirmTime, o.NoteSale, o.NoteUser, o.RefundFee, o.SettlementStatus, o.SettlementTime, o.FromSite, o.FinishTime, o.BuyUserID, o.ActivityID, o.OrderNo, o.Shareid, o.FromUID, o.UserAddressID, o.GoodsOriginFee, o.Orderid)
	}
	return err
}

// Save saves the Order to the database.
func (o *Order) Save(ctx context.Context) error {
	if o.Exists() {
		return o.Update(ctx)
	}

	return o.Insert(ctx)
}

// Delete deletes the Order from the database.
func (o *Order) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if o._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetOrderTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE orderid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, o.Orderid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, o.Orderid)
	} else {
		_, err = dbConn.Exec(sqlstr, o.Orderid)
	}

	if err != nil {
		return err
	}

	// set deleted
	o._deleted = true

	return nil
}

// OrderByOrderNo retrieves a row from 'aypcddg.orders' as a Order.
//
// Generated from index 'orders_order_no_unique'.
func OrderByOrderNo(ctx context.Context, orderNo sql.NullString, key ...interface{}) (*Order, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`orderid, q, uid, name, mobile, email, fid, sid, status, status_ext, created, pay_time, ship_time, total_fee, need_fee, settlement_paid, total_paid, pay_method, shipping, goods_fee, discount_fee, logistic_type, pid, cid, aid, address, address_info, invoice_info, slinfo, commission, prepay, affirm_time, note_sale, note_user, refund_fee, settlement_status, settlement_time, from_site, finish_time, buy_user_id, activity_id, order_no, shareid, from_uid, user_address_id, goods_origin_fee ` +
		`FROM ` + tableName +
		` WHERE order_no = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, orderNo)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	o := Order{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, orderNo).Scan(&o.Orderid, &o.Q, &o.UID, &o.Name, &o.Mobile, &o.Email, &o.Fid, &o.Sid, &o.Status, &o.StatusExt, &o.Created, &o.PayTime, &o.ShipTime, &o.TotalFee, &o.NeedFee, &o.SettlementPaid, &o.TotalPaid, &o.PayMethod, &o.Shipping, &o.GoodsFee, &o.DiscountFee, &o.LogisticType, &o.Pid, &o.Cid, &o.Aid, &o.Address, &o.AddressInfo, &o.InvoiceInfo, &o.Slinfo, &o.Commission, &o.Prepay, &o.AffirmTime, &o.NoteSale, &o.NoteUser, &o.RefundFee, &o.SettlementStatus, &o.SettlementTime, &o.FromSite, &o.FinishTime, &o.BuyUserID, &o.ActivityID, &o.OrderNo, &o.Shareid, &o.FromUID, &o.UserAddressID, &o.GoodsOriginFee)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, orderNo).Scan(&o.Orderid, &o.Q, &o.UID, &o.Name, &o.Mobile, &o.Email, &o.Fid, &o.Sid, &o.Status, &o.StatusExt, &o.Created, &o.PayTime, &o.ShipTime, &o.TotalFee, &o.NeedFee, &o.SettlementPaid, &o.TotalPaid, &o.PayMethod, &o.Shipping, &o.GoodsFee, &o.DiscountFee, &o.LogisticType, &o.Pid, &o.Cid, &o.Aid, &o.Address, &o.AddressInfo, &o.InvoiceInfo, &o.Slinfo, &o.Commission, &o.Prepay, &o.AffirmTime, &o.NoteSale, &o.NoteUser, &o.RefundFee, &o.SettlementStatus, &o.SettlementTime, &o.FromSite, &o.FinishTime, &o.BuyUserID, &o.ActivityID, &o.OrderNo, &o.Shareid, &o.FromUID, &o.UserAddressID, &o.GoodsOriginFee)
		if err != nil {
			return nil, err
		}
	}

	return &o, nil
}

// OrderByOrderid retrieves a row from 'aypcddg.orders' as a Order.
//
// Generated from index 'orders_orderid_pkey'.
func OrderByOrderid(ctx context.Context, orderid int, key ...interface{}) (*Order, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`orderid, q, uid, name, mobile, email, fid, sid, status, status_ext, created, pay_time, ship_time, total_fee, need_fee, settlement_paid, total_paid, pay_method, shipping, goods_fee, discount_fee, logistic_type, pid, cid, aid, address, address_info, invoice_info, slinfo, commission, prepay, affirm_time, note_sale, note_user, refund_fee, settlement_status, settlement_time, from_site, finish_time, buy_user_id, activity_id, order_no, shareid, from_uid, user_address_id, goods_origin_fee ` +
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
	o := Order{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, orderid).Scan(&o.Orderid, &o.Q, &o.UID, &o.Name, &o.Mobile, &o.Email, &o.Fid, &o.Sid, &o.Status, &o.StatusExt, &o.Created, &o.PayTime, &o.ShipTime, &o.TotalFee, &o.NeedFee, &o.SettlementPaid, &o.TotalPaid, &o.PayMethod, &o.Shipping, &o.GoodsFee, &o.DiscountFee, &o.LogisticType, &o.Pid, &o.Cid, &o.Aid, &o.Address, &o.AddressInfo, &o.InvoiceInfo, &o.Slinfo, &o.Commission, &o.Prepay, &o.AffirmTime, &o.NoteSale, &o.NoteUser, &o.RefundFee, &o.SettlementStatus, &o.SettlementTime, &o.FromSite, &o.FinishTime, &o.BuyUserID, &o.ActivityID, &o.OrderNo, &o.Shareid, &o.FromUID, &o.UserAddressID, &o.GoodsOriginFee)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, orderid).Scan(&o.Orderid, &o.Q, &o.UID, &o.Name, &o.Mobile, &o.Email, &o.Fid, &o.Sid, &o.Status, &o.StatusExt, &o.Created, &o.PayTime, &o.ShipTime, &o.TotalFee, &o.NeedFee, &o.SettlementPaid, &o.TotalPaid, &o.PayMethod, &o.Shipping, &o.GoodsFee, &o.DiscountFee, &o.LogisticType, &o.Pid, &o.Cid, &o.Aid, &o.Address, &o.AddressInfo, &o.InvoiceInfo, &o.Slinfo, &o.Commission, &o.Prepay, &o.AffirmTime, &o.NoteSale, &o.NoteUser, &o.RefundFee, &o.SettlementStatus, &o.SettlementTime, &o.FromSite, &o.FinishTime, &o.BuyUserID, &o.ActivityID, &o.OrderNo, &o.Shareid, &o.FromUID, &o.UserAddressID, &o.GoodsOriginFee)
		if err != nil {
			return nil, err
		}
	}

	return &o, nil
}

// OrdersByUID retrieves a row from 'aypcddg.orders' as a Order.
//
// Generated from index 'uid'.
func OrdersByUID(ctx context.Context, uid int, key ...interface{}) ([]*Order, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`orderid, q, uid, name, mobile, email, fid, sid, status, status_ext, created, pay_time, ship_time, total_fee, need_fee, settlement_paid, total_paid, pay_method, shipping, goods_fee, discount_fee, logistic_type, pid, cid, aid, address, address_info, invoice_info, slinfo, commission, prepay, affirm_time, note_sale, note_user, refund_fee, settlement_status, settlement_time, from_site, finish_time, buy_user_id, activity_id, order_no, shareid, from_uid, user_address_id, goods_origin_fee ` +
		`FROM ` + tableName +
		` WHERE uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*Order, 0)
	for queryData.Next() {
		o := Order{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&o.Orderid, &o.Q, &o.UID, &o.Name, &o.Mobile, &o.Email, &o.Fid, &o.Sid, &o.Status, &o.StatusExt, &o.Created, &o.PayTime, &o.ShipTime, &o.TotalFee, &o.NeedFee, &o.SettlementPaid, &o.TotalPaid, &o.PayMethod, &o.Shipping, &o.GoodsFee, &o.DiscountFee, &o.LogisticType, &o.Pid, &o.Cid, &o.Aid, &o.Address, &o.AddressInfo, &o.InvoiceInfo, &o.Slinfo, &o.Commission, &o.Prepay, &o.AffirmTime, &o.NoteSale, &o.NoteUser, &o.RefundFee, &o.SettlementStatus, &o.SettlementTime, &o.FromSite, &o.FinishTime, &o.BuyUserID, &o.ActivityID, &o.OrderNo, &o.Shareid, &o.FromUID, &o.UserAddressID, &o.GoodsOriginFee)
		if err != nil {
			return nil, err
		}

		res = append(res, &o)
	}

	return res, nil
}
