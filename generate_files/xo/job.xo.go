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

// Job represents a row from 'aypcddg.jobs'.
type Job struct {
	ID          uint64        `json:"id"`           // id
	Queue       string        `json:"queue"`        // queue
	Payload     string        `json:"payload"`      // payload
	Attempts    int8          `json:"attempts"`     // attempts
	ReservedAt  sql.NullInt64 `json:"reserved_at"`  // reserved_at
	AvailableAt uint          `json:"available_at"` // available_at
	CreatedAt   uint          `json:"created_at"`   // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Job exists in the database.
func (j *Job) Exists() bool { //jobs
	return j._exists
}

// Deleted provides information if the Job has been deleted from the database.
func (j *Job) Deleted() bool {
	return j._deleted
}

// Get table name
func GetJobTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "jobs", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Job to the database.
func (j *Job) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if j._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetJobTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`queue, payload, attempts, reserved_at, available_at, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, j.Queue, j.Payload, j.Attempts, j.ReservedAt, j.AvailableAt, j.CreatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, j.Queue, j.Payload, j.Attempts, j.ReservedAt, j.AvailableAt, j.CreatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, j.Queue, j.Payload, j.Attempts, j.ReservedAt, j.AvailableAt, j.CreatedAt)
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
	j.ID = uint64(id)
	j._exists = true

	return nil
}

// Update updates the Job in the database.
func (j *Job) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if j._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetJobTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`queue = ?, payload = ?, attempts = ?, reserved_at = ?, available_at = ?, created_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, j.Queue, j.Payload, j.Attempts, j.ReservedAt, j.AvailableAt, j.CreatedAt, j.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, j.Queue, j.Payload, j.Attempts, j.ReservedAt, j.AvailableAt, j.CreatedAt, j.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, j.Queue, j.Payload, j.Attempts, j.ReservedAt, j.AvailableAt, j.CreatedAt, j.ID)
	}
	return err
}

// Save saves the Job to the database.
func (j *Job) Save(ctx context.Context) error {
	if j.Exists() {
		return j.Update(ctx)
	}

	return j.Insert(ctx)
}

// Delete deletes the Job from the database.
func (j *Job) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if j._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetJobTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, j.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, j.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, j.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	j._deleted = true

	return nil
}

// JobByID retrieves a row from 'aypcddg.jobs' as a Job.
//
// Generated from index 'jobs_id_pkey'.
func JobByID(ctx context.Context, id uint64, key ...interface{}) (*Job, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetJobTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, queue, payload, attempts, reserved_at, available_at, created_at ` +
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
	j := Job{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&j.ID, &j.Queue, &j.Payload, &j.Attempts, &j.ReservedAt, &j.AvailableAt, &j.CreatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&j.ID, &j.Queue, &j.Payload, &j.Attempts, &j.ReservedAt, &j.AvailableAt, &j.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &j, nil
}

// JobsByQueue retrieves a row from 'aypcddg.jobs' as a Job.
//
// Generated from index 'jobs_queue_index'.
func JobsByQueue(ctx context.Context, queue string, key ...interface{}) ([]*Job, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetJobTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, queue, payload, attempts, reserved_at, available_at, created_at ` +
		`FROM ` + tableName +
		` WHERE queue = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, queue)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, queue)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, queue)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*Job, 0)
	for queryData.Next() {
		j := Job{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&j.ID, &j.Queue, &j.Payload, &j.Attempts, &j.ReservedAt, &j.AvailableAt, &j.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &j)
	}

	return res, nil
}
