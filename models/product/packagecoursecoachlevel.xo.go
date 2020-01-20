// Package models contains the types for schema 'saas'.
package product

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"j7go/components"
	"j7go/utils"

	"go.uber.org/zap"
)

// PackageCourseCoachLevel represents a row from 'saas.package_course_coach_level'.
type PackageCourseCoachLevel struct {
	ID              uint `json:"id"`                // id
	CoursePackageID uint `json:"course_package_id"` // course_package_id
	CoachLevelID    uint `json:"coach_level_id"`    // coach_level_id
	CourseID        uint `json:"course_id"`         // course_id
	IsDel           int8 `json:"is_del"`            // is_del
	CreatedTime     uint `json:"created_time"`      // created_time
	UpdatedTime     uint `json:"updated_time"`      // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PackageCourseCoachLevel exists in the database.
func (pccl *PackageCourseCoachLevel) Exists() bool { //package_course_coach_level
	return pccl._exists
}

// Deleted provides information if the PackageCourseCoachLevel has been deleted from the database.
func (pccl *PackageCourseCoachLevel) Deleted() bool {
	return pccl._deleted
}

// Get table name
func GetPackageCourseCoachLevelTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "package_course_coach_level", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the PackageCourseCoachLevel to the database.
func (pccl *PackageCourseCoachLevel) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if pccl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPackageCourseCoachLevelTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`course_package_id, coach_level_id, course_id, is_del, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pccl.CoursePackageID, pccl.CoachLevelID, pccl.CourseID, pccl.IsDel, pccl.CreatedTime, pccl.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, pccl.CoursePackageID, pccl.CoachLevelID, pccl.CourseID, pccl.IsDel, pccl.CreatedTime, pccl.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, pccl.CoursePackageID, pccl.CoachLevelID, pccl.CourseID, pccl.IsDel, pccl.CreatedTime, pccl.UpdatedTime)
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
	pccl.ID = uint(id)
	pccl._exists = true

	return nil
}

// Update updates the PackageCourseCoachLevel in the database.
func (pccl *PackageCourseCoachLevel) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if pccl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPackageCourseCoachLevelTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`course_package_id = ?, coach_level_id = ?, course_id = ?, is_del = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pccl.CoursePackageID, pccl.CoachLevelID, pccl.CourseID, pccl.IsDel, pccl.CreatedTime, pccl.UpdatedTime, pccl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, pccl.CoursePackageID, pccl.CoachLevelID, pccl.CourseID, pccl.IsDel, pccl.CreatedTime, pccl.UpdatedTime, pccl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, pccl.CoursePackageID, pccl.CoachLevelID, pccl.CourseID, pccl.IsDel, pccl.CreatedTime, pccl.UpdatedTime, pccl.ID)
	}
	return err
}

// Save saves the PackageCourseCoachLevel to the database.
func (pccl *PackageCourseCoachLevel) Save(ctx context.Context) error {
	if pccl.Exists() {
		return pccl.Update(ctx)
	}

	return pccl.Insert(ctx)
}

// Delete deletes the PackageCourseCoachLevel from the database.
func (pccl *PackageCourseCoachLevel) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if pccl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPackageCourseCoachLevelTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pccl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, pccl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, pccl.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	pccl._deleted = true

	return nil
}

// PackageCourseCoachLevelByID retrieves a row from 'saas.package_course_coach_level' as a PackageCourseCoachLevel.
//
// Generated from index 'package_course_coach_level_id_pkey'.
func PackageCourseCoachLevelByID(ctx context.Context, id uint, key ...interface{}) (*PackageCourseCoachLevel, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetPackageCourseCoachLevelTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, course_package_id, coach_level_id, course_id, is_del, created_time, updated_time ` +
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
	pccl := PackageCourseCoachLevel{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&pccl.ID, &pccl.CoursePackageID, &pccl.CoachLevelID, &pccl.CourseID, &pccl.IsDel, &pccl.CreatedTime, &pccl.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&pccl.ID, &pccl.CoursePackageID, &pccl.CoachLevelID, &pccl.CourseID, &pccl.IsDel, &pccl.CreatedTime, &pccl.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &pccl, nil
}

//根据课程包id以及课程id删除课程包内的课程
func DelPackageCourseCoachLevelByPackageCourseID (ctx context.Context, packageCourseId uint, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPackageCourseCoachLevelTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr, args, err := squirrel.Update(tableName).
		Set("is_del", utils.DELETED).
		Where(squirrel.Eq{"course_package_id": packageCourseId}).
		ToSql()
	if err != nil {
		return err
	}

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, args)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, args...)
	} else {
		_, err = dbConn.Exec(sqlstr, args...)
	}

	if err != nil {
		return err
	}

	return nil
}