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

// Express represents a row from 'aypcddg.express'.
type Express struct {
	ID        uint           `json:"id"`         // id
	Title     sql.NullString `json:"title"`      // title
	Code      sql.NullString `json:"code"`       // code
	Remark    sql.NullString `json:"remark"`     // remark
	Status    sql.NullInt64  `json:"status"`     // status
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Express exists in the database.
func (e *Express) Exists() bool { //express
	return e._exists
}

// Deleted provides information if the Express has been deleted from the database.
func (e *Express) Deleted() bool {
	return e._deleted
}

// Get table name
func GetExpressTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "express", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Express to the database.
func (e *Express) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetExpressTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`title, code, remark, status, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, e.Title, e.Code, e.Remark, e.Status, e.CreatedAt, e.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, e.Title, e.Code, e.Remark, e.Status, e.CreatedAt, e.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, e.Title, e.Code, e.Remark, e.Status, e.CreatedAt, e.UpdatedAt)
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
	e.ID = uint(id)
	e._exists = true

	return nil
}

// Update updates the Express in the database.
func (e *Express) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if e._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetExpressTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`title = ?, code = ?, remark = ?, status = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, e.Title, e.Code, e.Remark, e.Status, e.CreatedAt, e.UpdatedAt, e.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, e.Title, e.Code, e.Remark, e.Status, e.CreatedAt, e.UpdatedAt, e.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, e.Title, e.Code, e.Remark, e.Status, e.CreatedAt, e.UpdatedAt, e.ID)
	}
	return err
}

// Save saves the Express to the database.
func (e *Express) Save(ctx context.Context) error {
	if e.Exists() {
		return e.Update(ctx)
	}

	return e.Insert(ctx)
}

// Delete deletes the Express from the database.
func (e *Express) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if e._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetExpressTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, e.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, e.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, e.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	e._deleted = true

	return nil
}

// ExpressByID retrieves a row from 'aypcddg.express' as a Express.
//
// Generated from index 'express_id_pkey'.
func ExpressByID(ctx context.Context, id uint, key ...interface{}) (*Express, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetExpressTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, title, code, remark, status, created_at, updated_at ` +
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
	e := Express{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&e.ID, &e.Title, &e.Code, &e.Remark, &e.Status, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&e.ID, &e.Title, &e.Code, &e.Remark, &e.Status, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &e, nil
}

// ExpressesByTitle retrieves a row from 'aypcddg.express' as a Express.
//
// Generated from index 'express_title_index'.
func ExpressesByTitle(ctx context.Context, title sql.NullString, key ...interface{}) ([]*Express, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetExpressTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, title, code, remark, status, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE title = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, title)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, title)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, title)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*Express, 0)
	for queryData.Next() {
		e := Express{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&e.ID, &e.Title, &e.Code, &e.Remark, &e.Status, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &e)
	}

	return res, nil
}
