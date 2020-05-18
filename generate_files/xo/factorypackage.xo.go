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

// FactoryPackage represents a row from 'aypcddg.factory_package'.
type FactoryPackage struct {
	Fpid    int16          `json:"fpid"`    // fpid
	Title   sql.NullString `json:"title"`   // title
	Desc    sql.NullString `json:"desc"`    // desc
	Seq     int16          `json:"seq"`     // seq
	Status  bool           `json:"status"`  // status
	Created uint           `json:"created"` // created

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryPackage exists in the database.
func (fp *FactoryPackage) Exists() bool { //factory_package
	return fp._exists
}

// Deleted provides information if the FactoryPackage has been deleted from the database.
func (fp *FactoryPackage) Deleted() bool {
	return fp._deleted
}

// Get table name
func GetFactoryPackageTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "factory_package", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryPackage to the database.
func (fp *FactoryPackage) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if fp._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryPackageTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`title, desc, seq, status, created` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fp.Title, fp.Desc, fp.Seq, fp.Status, fp.Created)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, fp.Title, fp.Desc, fp.Seq, fp.Status, fp.Created)
	} else {
		res, err = dbConn.Exec(sqlstr, fp.Title, fp.Desc, fp.Seq, fp.Status, fp.Created)
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
	fp.Fpid = int16(id)
	fp._exists = true

	return nil
}

// Update updates the FactoryPackage in the database.
func (fp *FactoryPackage) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryPackageTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`title = ?, desc = ?, seq = ?, status = ?, created = ?` +
		` WHERE fpid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fp.Title, fp.Desc, fp.Seq, fp.Status, fp.Created, fp.Fpid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fp.Title, fp.Desc, fp.Seq, fp.Status, fp.Created, fp.Fpid)
	} else {
		_, err = dbConn.Exec(sqlstr, fp.Title, fp.Desc, fp.Seq, fp.Status, fp.Created, fp.Fpid)
	}
	return err
}

// Save saves the FactoryPackage to the database.
func (fp *FactoryPackage) Save(ctx context.Context) error {
	if fp.Exists() {
		return fp.Update(ctx)
	}

	return fp.Insert(ctx)
}

// Delete deletes the FactoryPackage from the database.
func (fp *FactoryPackage) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fp._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryPackageTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fpid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fp.Fpid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fp.Fpid)
	} else {
		_, err = dbConn.Exec(sqlstr, fp.Fpid)
	}

	if err != nil {
		return err
	}

	// set deleted
	fp._deleted = true

	return nil
}

// FactoryPackageByFpid retrieves a row from 'aypcddg.factory_package' as a FactoryPackage.
//
// Generated from index 'factory_package_fpid_pkey'.
func FactoryPackageByFpid(ctx context.Context, fpid int16, key ...interface{}) (*FactoryPackage, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryPackageTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`fpid, title, desc, seq, status, created ` +
		`FROM ` + tableName +
		` WHERE fpid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fpid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	fp := FactoryPackage{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fpid).Scan(&fp.Fpid, &fp.Title, &fp.Desc, &fp.Seq, &fp.Status, &fp.Created)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fpid).Scan(&fp.Fpid, &fp.Title, &fp.Desc, &fp.Seq, &fp.Status, &fp.Created)
		if err != nil {
			return nil, err
		}
	}

	return &fp, nil
}
