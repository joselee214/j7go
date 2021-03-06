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

// GoodsComment represents a row from 'aypcddg.goods_comment'.
type GoodsComment struct {
	ID                   int            `json:"id"`                     // id
	Pid                  int            `json:"pid"`                    // pid
	Gid                  sql.NullInt64  `json:"gid"`                    // gid
	UID                  sql.NullInt64  `json:"uid"`                    // uid
	Content              sql.NullString `json:"content"`                // content
	Level                sql.NullInt64  `json:"level"`                  // level
	Pic1                 sql.NullString `json:"pic1"`                   // pic1
	Pic1content          sql.NullString `json:"pic1Content"`            // pic1Content
	Title                sql.NullString `json:"title"`                  // title
	Pic2                 sql.NullString `json:"pic2"`                   // pic2
	Pic2content          sql.NullString `json:"pic2Content"`            // pic2Content
	Pic3                 sql.NullString `json:"pic3"`                   // pic3
	Pic3content          sql.NullString `json:"pic3Content"`            // pic3Content
	Pic4                 sql.NullString `json:"pic4"`                   // pic4
	Pic4content          sql.NullString `json:"pic4Content"`            // pic4Content
	Pic5                 sql.NullString `json:"pic5"`                   // pic5
	Pic5content          sql.NullString `json:"pic5Content"`            // pic5Content
	Titlepageid          bool           `json:"titlePageId"`            // titlePageId
	Isshared             int8           `json:"isShared"`               // isShared
	Orderid              sql.NullInt64  `json:"orderid"`                // orderid
	Isactivity           bool           `json:"isActivity"`             // isActivity
	Isshow               bool           `json:"isShow"`                 // isShow
	Addtime              sql.NullInt64  `json:"addTime"`                // addTime
	Hits                 int            `json:"hits"`                   // hits
	ThrowColdWater       sql.NullInt64  `json:"throw_cold_water"`       // throw_cold_water
	ParticipatePeopleNum int            `json:"participate_people_num"` // participate_people_num

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GoodsComment exists in the database.
func (gc *GoodsComment) Exists() bool { //goods_comment
	return gc._exists
}

// Deleted provides information if the GoodsComment has been deleted from the database.
func (gc *GoodsComment) Deleted() bool {
	return gc._deleted
}

// Get table name
func GetGoodsCommentTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "goods_comment", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the GoodsComment to the database.
func (gc *GoodsComment) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if gc._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCommentTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`pid, gid, uid, content, level, pic1, pic1Content, title, pic2, pic2Content, pic3, pic3Content, pic4, pic4Content, pic5, pic5Content, titlePageId, isShared, orderid, isActivity, isShow, addTime, hits, throw_cold_water, participate_people_num` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gc.Pid, gc.Gid, gc.UID, gc.Content, gc.Level, gc.Pic1, gc.Pic1content, gc.Title, gc.Pic2, gc.Pic2content, gc.Pic3, gc.Pic3content, gc.Pic4, gc.Pic4content, gc.Pic5, gc.Pic5content, gc.Titlepageid, gc.Isshared, gc.Orderid, gc.Isactivity, gc.Isshow, gc.Addtime, gc.Hits, gc.ThrowColdWater, gc.ParticipatePeopleNum)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, gc.Pid, gc.Gid, gc.UID, gc.Content, gc.Level, gc.Pic1, gc.Pic1content, gc.Title, gc.Pic2, gc.Pic2content, gc.Pic3, gc.Pic3content, gc.Pic4, gc.Pic4content, gc.Pic5, gc.Pic5content, gc.Titlepageid, gc.Isshared, gc.Orderid, gc.Isactivity, gc.Isshow, gc.Addtime, gc.Hits, gc.ThrowColdWater, gc.ParticipatePeopleNum)
	} else {
		res, err = dbConn.Exec(sqlstr, gc.Pid, gc.Gid, gc.UID, gc.Content, gc.Level, gc.Pic1, gc.Pic1content, gc.Title, gc.Pic2, gc.Pic2content, gc.Pic3, gc.Pic3content, gc.Pic4, gc.Pic4content, gc.Pic5, gc.Pic5content, gc.Titlepageid, gc.Isshared, gc.Orderid, gc.Isactivity, gc.Isshow, gc.Addtime, gc.Hits, gc.ThrowColdWater, gc.ParticipatePeopleNum)
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
	gc.ID = int(id)
	gc._exists = true

	return nil
}

// Update updates the GoodsComment in the database.
func (gc *GoodsComment) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCommentTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`pid = ?, gid = ?, uid = ?, content = ?, level = ?, pic1 = ?, pic1Content = ?, title = ?, pic2 = ?, pic2Content = ?, pic3 = ?, pic3Content = ?, pic4 = ?, pic4Content = ?, pic5 = ?, pic5Content = ?, titlePageId = ?, isShared = ?, orderid = ?, isActivity = ?, isShow = ?, addTime = ?, hits = ?, throw_cold_water = ?, participate_people_num = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gc.Pid, gc.Gid, gc.UID, gc.Content, gc.Level, gc.Pic1, gc.Pic1content, gc.Title, gc.Pic2, gc.Pic2content, gc.Pic3, gc.Pic3content, gc.Pic4, gc.Pic4content, gc.Pic5, gc.Pic5content, gc.Titlepageid, gc.Isshared, gc.Orderid, gc.Isactivity, gc.Isshow, gc.Addtime, gc.Hits, gc.ThrowColdWater, gc.ParticipatePeopleNum, gc.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gc.Pid, gc.Gid, gc.UID, gc.Content, gc.Level, gc.Pic1, gc.Pic1content, gc.Title, gc.Pic2, gc.Pic2content, gc.Pic3, gc.Pic3content, gc.Pic4, gc.Pic4content, gc.Pic5, gc.Pic5content, gc.Titlepageid, gc.Isshared, gc.Orderid, gc.Isactivity, gc.Isshow, gc.Addtime, gc.Hits, gc.ThrowColdWater, gc.ParticipatePeopleNum, gc.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, gc.Pid, gc.Gid, gc.UID, gc.Content, gc.Level, gc.Pic1, gc.Pic1content, gc.Title, gc.Pic2, gc.Pic2content, gc.Pic3, gc.Pic3content, gc.Pic4, gc.Pic4content, gc.Pic5, gc.Pic5content, gc.Titlepageid, gc.Isshared, gc.Orderid, gc.Isactivity, gc.Isshow, gc.Addtime, gc.Hits, gc.ThrowColdWater, gc.ParticipatePeopleNum, gc.ID)
	}
	return err
}

// Save saves the GoodsComment to the database.
func (gc *GoodsComment) Save(ctx context.Context) error {
	if gc.Exists() {
		return gc.Update(ctx)
	}

	return gc.Insert(ctx)
}

// Delete deletes the GoodsComment from the database.
func (gc *GoodsComment) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gc._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCommentTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gc.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gc.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, gc.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	gc._deleted = true

	return nil
}

// GoodsCommentByID retrieves a row from 'aypcddg.goods_comment' as a GoodsComment.
//
// Generated from index 'goods_comment_id_pkey'.
func GoodsCommentByID(ctx context.Context, id int, key ...interface{}) (*GoodsComment, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCommentTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, pid, gid, uid, content, level, pic1, pic1Content, title, pic2, pic2Content, pic3, pic3Content, pic4, pic4Content, pic5, pic5Content, titlePageId, isShared, orderid, isActivity, isShow, addTime, hits, throw_cold_water, participate_people_num ` +
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
	gc := GoodsComment{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&gc.ID, &gc.Pid, &gc.Gid, &gc.UID, &gc.Content, &gc.Level, &gc.Pic1, &gc.Pic1content, &gc.Title, &gc.Pic2, &gc.Pic2content, &gc.Pic3, &gc.Pic3content, &gc.Pic4, &gc.Pic4content, &gc.Pic5, &gc.Pic5content, &gc.Titlepageid, &gc.Isshared, &gc.Orderid, &gc.Isactivity, &gc.Isshow, &gc.Addtime, &gc.Hits, &gc.ThrowColdWater, &gc.ParticipatePeopleNum)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&gc.ID, &gc.Pid, &gc.Gid, &gc.UID, &gc.Content, &gc.Level, &gc.Pic1, &gc.Pic1content, &gc.Title, &gc.Pic2, &gc.Pic2content, &gc.Pic3, &gc.Pic3content, &gc.Pic4, &gc.Pic4content, &gc.Pic5, &gc.Pic5content, &gc.Titlepageid, &gc.Isshared, &gc.Orderid, &gc.Isactivity, &gc.Isshow, &gc.Addtime, &gc.Hits, &gc.ThrowColdWater, &gc.ParticipatePeopleNum)
		if err != nil {
			return nil, err
		}
	}

	return &gc, nil
}

// GoodsCommentsByOrderid retrieves a row from 'aypcddg.goods_comment' as a GoodsComment.
//
// Generated from index 'orderid'.
func GoodsCommentsByOrderid(ctx context.Context, orderid sql.NullInt64, key ...interface{}) ([]*GoodsComment, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCommentTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, pid, gid, uid, content, level, pic1, pic1Content, title, pic2, pic2Content, pic3, pic3Content, pic4, pic4Content, pic5, pic5Content, titlePageId, isShared, orderid, isActivity, isShow, addTime, hits, throw_cold_water, participate_people_num ` +
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
	res := make([]*GoodsComment, 0)
	for queryData.Next() {
		gc := GoodsComment{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&gc.ID, &gc.Pid, &gc.Gid, &gc.UID, &gc.Content, &gc.Level, &gc.Pic1, &gc.Pic1content, &gc.Title, &gc.Pic2, &gc.Pic2content, &gc.Pic3, &gc.Pic3content, &gc.Pic4, &gc.Pic4content, &gc.Pic5, &gc.Pic5content, &gc.Titlepageid, &gc.Isshared, &gc.Orderid, &gc.Isactivity, &gc.Isshow, &gc.Addtime, &gc.Hits, &gc.ThrowColdWater, &gc.ParticipatePeopleNum)
		if err != nil {
			return nil, err
		}

		res = append(res, &gc)
	}

	return res, nil
}
