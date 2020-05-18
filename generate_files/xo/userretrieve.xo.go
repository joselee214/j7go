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

// UserRetrieve represents a row from 'aypcddg.user_retrieve'.
type UserRetrieve struct {
	UID          uint           `json:"uid"`           // uid
	Code         sql.NullString `json:"code"`          // code
	Created      uint           `json:"created"`       // created
	RetrieveType sql.NullString `json:"retrieve_type"` // retrieve_type

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserRetrieve exists in the database.
func (ur *UserRetrieve) Exists() bool { //user_retrieve
	return ur._exists
}

// Deleted provides information if the UserRetrieve has been deleted from the database.
func (ur *UserRetrieve) Deleted() bool {
	return ur._deleted
}

// Get table name
func GetUserRetrieveTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "user_retrieve", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserRetrieve to the database.
func (ur *UserRetrieve) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ur._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserRetrieveTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, code, created, retrieve_type` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ur.UID, ur.Code, ur.Created, ur.RetrieveType)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, ur.UID, ur.Code, ur.Created, ur.RetrieveType)
	} else {
		res, err = dbConn.Exec(sqlstr, ur.UID, ur.Code, ur.Created, ur.RetrieveType)
	}

	if err != nil {
		return err
	}

	// set existence
	ur._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ur.UID = uint(id)
	ur._exists = true

	return nil
}

// Update updates the UserRetrieve in the database.
func (ur *UserRetrieve) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ur._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserRetrieveTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`code = ?, created = ?, retrieve_type = ?` +
		` WHERE uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ur.Code, ur.Created, ur.RetrieveType, ur.UID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ur.Code, ur.Created, ur.RetrieveType, ur.UID)
	} else {
		_, err = dbConn.Exec(sqlstr, ur.Code, ur.Created, ur.RetrieveType, ur.UID)
	}
	return err
}

// Save saves the UserRetrieve to the database.
func (ur *UserRetrieve) Save(ctx context.Context) error {
	if ur.Exists() {
		return ur.Update(ctx)
	}

	return ur.Insert(ctx)
}

// Delete deletes the UserRetrieve from the database.
func (ur *UserRetrieve) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ur._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserRetrieveTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ur.UID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ur.UID)
	} else {
		_, err = dbConn.Exec(sqlstr, ur.UID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ur._deleted = true

	return nil
}

// UserRetrievesByCode retrieves a row from 'aypcddg.user_retrieve' as a UserRetrieve.
//
// Generated from index 'code'.
func UserRetrievesByCode(ctx context.Context, code sql.NullString, key ...interface{}) ([]*UserRetrieve, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserRetrieveTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uid, code, created, retrieve_type ` +
		`FROM ` + tableName +
		` WHERE code = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, code)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, code)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, code)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserRetrieve, 0)
	for queryData.Next() {
		ur := UserRetrieve{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ur.UID, &ur.Code, &ur.Created, &ur.RetrieveType)
		if err != nil {
			return nil, err
		}

		res = append(res, &ur)
	}

	return res, nil
}

// UserRetrievesByCreated retrieves a row from 'aypcddg.user_retrieve' as a UserRetrieve.
//
// Generated from index 'created'.
func UserRetrievesByCreated(ctx context.Context, created uint, key ...interface{}) ([]*UserRetrieve, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserRetrieveTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uid, code, created, retrieve_type ` +
		`FROM ` + tableName +
		` WHERE created = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, created)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, created)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, created)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserRetrieve, 0)
	for queryData.Next() {
		ur := UserRetrieve{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ur.UID, &ur.Code, &ur.Created, &ur.RetrieveType)
		if err != nil {
			return nil, err
		}

		res = append(res, &ur)
	}

	return res, nil
}

// UserRetrieveByUID retrieves a row from 'aypcddg.user_retrieve' as a UserRetrieve.
//
// Generated from index 'user_retrieve_uid_pkey'.
func UserRetrieveByUID(ctx context.Context, uid uint, key ...interface{}) (*UserRetrieve, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserRetrieveTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uid, code, created, retrieve_type ` +
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
	ur := UserRetrieve{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, uid).Scan(&ur.UID, &ur.Code, &ur.Created, &ur.RetrieveType)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, uid).Scan(&ur.UID, &ur.Code, &ur.Created, &ur.RetrieveType)
		if err != nil {
			return nil, err
		}
	}

	return &ur, nil
}