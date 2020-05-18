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

// Deliveryman represents a row from 'aypcddg.deliveryman'.
type Deliveryman struct {
	Fsdid   uint           `json:"fsdid"`   // fsdid
	Fsid    uint           `json:"fsid"`    // fsid
	Contact sql.NullString `json:"contact"` // contact
	Phone   sql.NullString `json:"phone"`   // phone
	Created uint           `json:"created"` // created

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Deliveryman exists in the database.
func (d *Deliveryman) Exists() bool { //deliveryman
	return d._exists
}

// Deleted provides information if the Deliveryman has been deleted from the database.
func (d *Deliveryman) Deleted() bool {
	return d._deleted
}

// Get table name
func GetDeliverymanTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "deliveryman", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Deliveryman to the database.
func (d *Deliveryman) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if d._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDeliverymanTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`fsid, contact, phone, created` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, d.Fsid, d.Contact, d.Phone, d.Created)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, d.Fsid, d.Contact, d.Phone, d.Created)
	} else {
		res, err = dbConn.Exec(sqlstr, d.Fsid, d.Contact, d.Phone, d.Created)
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
	d.Fsdid = uint(id)
	d._exists = true

	return nil
}

// Update updates the Deliveryman in the database.
func (d *Deliveryman) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if d._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDeliverymanTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`fsid = ?, contact = ?, phone = ?, created = ?` +
		` WHERE fsdid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, d.Fsid, d.Contact, d.Phone, d.Created, d.Fsdid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, d.Fsid, d.Contact, d.Phone, d.Created, d.Fsdid)
	} else {
		_, err = dbConn.Exec(sqlstr, d.Fsid, d.Contact, d.Phone, d.Created, d.Fsdid)
	}
	return err
}

// Save saves the Deliveryman to the database.
func (d *Deliveryman) Save(ctx context.Context) error {
	if d.Exists() {
		return d.Update(ctx)
	}

	return d.Insert(ctx)
}

// Delete deletes the Deliveryman from the database.
func (d *Deliveryman) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if d._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDeliverymanTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fsdid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, d.Fsdid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, d.Fsdid)
	} else {
		_, err = dbConn.Exec(sqlstr, d.Fsdid)
	}

	if err != nil {
		return err
	}

	// set deleted
	d._deleted = true

	return nil
}

// DeliverymanByFsdid retrieves a row from 'aypcddg.deliveryman' as a Deliveryman.
//
// Generated from index 'deliveryman_fsdid_pkey'.
func DeliverymanByFsdid(ctx context.Context, fsdid uint, key ...interface{}) (*Deliveryman, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDeliverymanTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`fsdid, fsid, contact, phone, created ` +
		`FROM ` + tableName +
		` WHERE fsdid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fsdid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	d := Deliveryman{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fsdid).Scan(&d.Fsdid, &d.Fsid, &d.Contact, &d.Phone, &d.Created)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fsdid).Scan(&d.Fsdid, &d.Fsid, &d.Contact, &d.Phone, &d.Created)
		if err != nil {
			return nil, err
		}
	}

	return &d, nil
}
