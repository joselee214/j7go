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

// Interviewer represents a row from 'aypcddg.interviewer'.
type Interviewer struct {
	DuID           uint          `json:"du_id"`           // du_id
	Interviewer    int           `json:"interviewer"`     // interviewer
	IsParticipated sql.NullInt64 `json:"is_participated"` // is_participated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Interviewer exists in the database.
func (i *Interviewer) Exists() bool { //interviewer
	return i._exists
}

// Deleted provides information if the Interviewer has been deleted from the database.
func (i *Interviewer) Deleted() bool {
	return i._deleted
}

// Get table name
func GetInterviewerTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "interviewer", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Interviewer to the database.
func (i *Interviewer) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if i._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetInterviewerTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`du_id, is_participated` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, i.DuID, i.IsParticipated)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, i.DuID, i.IsParticipated)
	} else {
		res, err = dbConn.Exec(sqlstr, i.DuID, i.IsParticipated)
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
	i.Interviewer = int(id)
	i._exists = true

	return nil
}

// Update updates the Interviewer in the database.
func (i *Interviewer) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if i._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetInterviewerTableName(key...)
	if err != nil {
		return err
	}

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`is_participated = ?` +
		` WHERE du_id = ? AND interviewer = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, i.IsParticipated, i.DuID, i.Interviewer)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, i.IsParticipated, i.DuID, i.Interviewer)
	} else {
		_, err = dbConn.Exec(sqlstr, i.IsParticipated, i.DuID, i.Interviewer)
	}
	return err
}

// Save saves the Interviewer to the database.
func (i *Interviewer) Save(ctx context.Context) error {
	if i.Exists() {
		return i.Update(ctx)
	}

	return i.Insert(ctx)
}

// Delete deletes the Interviewer from the database.
func (i *Interviewer) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if i._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetInterviewerTableName(key...)
	if err != nil {
		return err
	}
	//2

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE interviewer = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, i.Interviewer)))

	if tx != nil {
		_, err = tx.Exec(sqlstr, i.Interviewer)
	} else {
		_, err = dbConn.Exec(sqlstr, i.Interviewer)
	}
	if err != nil {
		return err
	}

	// set deleted
	i._deleted = true

	return nil
}

// InterviewerByInterviewer retrieves a row from 'aypcddg.interviewer' as a Interviewer.
//
// Generated from index 'interviewer_interviewer_pkey'.
func InterviewerByInterviewer(ctx context.Context, interviewer int, key ...interface{}) (*Interviewer, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetInterviewerTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`du_id, interviewer, is_participated ` +
		`FROM ` + tableName +
		` WHERE interviewer = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, interviewer)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	i := Interviewer{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, interviewer).Scan(&i.DuID, &i.Interviewer, &i.IsParticipated)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, interviewer).Scan(&i.DuID, &i.Interviewer, &i.IsParticipated)
		if err != nil {
			return nil, err
		}
	}

	return &i, nil
}