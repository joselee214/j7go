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

// WeixinBindMina represents a row from 'aypcddg.weixin_bind_mina'.
type WeixinBindMina struct {
	ID              uint           `json:"id"`                // id
	UID             int            `json:"uid"`               // uid
	Type            int8           `json:"type"`              // type
	NickName        string         `json:"nick_name"`         // nick_name
	UserName        string         `json:"user_name"`         // user_name
	AliasName       string         `json:"alias_name"`        // alias_name
	HeadImg         string         `json:"head_img"`          // head_img
	QrcodeURL       string         `json:"qrcode_url"`        // qrcode_url
	ServiceTypeInfo int8           `json:"service_type_info"` // service_type_info
	VerifyTypeInfo  int8           `json:"verify_type_info"`  // verify_type_info
	AccountExtra    string         `json:"account_extra"`     // account_extra
	MenuMediaID     string         `json:"menu_media_id"`     // menu_media_id
	Appid           string         `json:"appid"`             // appid
	AuthCode        string         `json:"auth_code"`         // auth_code
	AccessToken     string         `json:"access_token"`      // access_token
	RefreshToken    string         `json:"refresh_token"`     // refresh_token
	ExpiresIn       int            `json:"expires_in"`        // expires_in
	Extra           string         `json:"extra"`             // extra
	Status          int8           `json:"status"`            // status
	UpdatedTime     mysql.NullTime `json:"updated_time"`      // updated_time
	CreatedTime     mysql.NullTime `json:"created_time"`      // created_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WeixinBindMina exists in the database.
func (wbm *WeixinBindMina) Exists() bool { //weixin_bind_mina
	return wbm._exists
}

// Deleted provides information if the WeixinBindMina has been deleted from the database.
func (wbm *WeixinBindMina) Deleted() bool {
	return wbm._deleted
}

// Get table name
func GetWeixinBindMinaTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "weixin_bind_mina", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the WeixinBindMina to the database.
func (wbm *WeixinBindMina) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if wbm._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWeixinBindMinaTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, type, nick_name, user_name, alias_name, head_img, qrcode_url, service_type_info, verify_type_info, account_extra, menu_media_id, appid, auth_code, access_token, refresh_token, expires_in, extra, status, updated_time, created_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wbm.UID, wbm.Type, wbm.NickName, wbm.UserName, wbm.AliasName, wbm.HeadImg, wbm.QrcodeURL, wbm.ServiceTypeInfo, wbm.VerifyTypeInfo, wbm.AccountExtra, wbm.MenuMediaID, wbm.Appid, wbm.AuthCode, wbm.AccessToken, wbm.RefreshToken, wbm.ExpiresIn, wbm.Extra, wbm.Status, wbm.UpdatedTime, wbm.CreatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, wbm.UID, wbm.Type, wbm.NickName, wbm.UserName, wbm.AliasName, wbm.HeadImg, wbm.QrcodeURL, wbm.ServiceTypeInfo, wbm.VerifyTypeInfo, wbm.AccountExtra, wbm.MenuMediaID, wbm.Appid, wbm.AuthCode, wbm.AccessToken, wbm.RefreshToken, wbm.ExpiresIn, wbm.Extra, wbm.Status, wbm.UpdatedTime, wbm.CreatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, wbm.UID, wbm.Type, wbm.NickName, wbm.UserName, wbm.AliasName, wbm.HeadImg, wbm.QrcodeURL, wbm.ServiceTypeInfo, wbm.VerifyTypeInfo, wbm.AccountExtra, wbm.MenuMediaID, wbm.Appid, wbm.AuthCode, wbm.AccessToken, wbm.RefreshToken, wbm.ExpiresIn, wbm.Extra, wbm.Status, wbm.UpdatedTime, wbm.CreatedTime)
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
	wbm.ID = uint(id)
	wbm._exists = true

	return nil
}

// Update updates the WeixinBindMina in the database.
func (wbm *WeixinBindMina) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if wbm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWeixinBindMinaTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, type = ?, nick_name = ?, user_name = ?, alias_name = ?, head_img = ?, qrcode_url = ?, service_type_info = ?, verify_type_info = ?, account_extra = ?, menu_media_id = ?, appid = ?, auth_code = ?, access_token = ?, refresh_token = ?, expires_in = ?, extra = ?, status = ?, updated_time = ?, created_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wbm.UID, wbm.Type, wbm.NickName, wbm.UserName, wbm.AliasName, wbm.HeadImg, wbm.QrcodeURL, wbm.ServiceTypeInfo, wbm.VerifyTypeInfo, wbm.AccountExtra, wbm.MenuMediaID, wbm.Appid, wbm.AuthCode, wbm.AccessToken, wbm.RefreshToken, wbm.ExpiresIn, wbm.Extra, wbm.Status, wbm.UpdatedTime, wbm.CreatedTime, wbm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, wbm.UID, wbm.Type, wbm.NickName, wbm.UserName, wbm.AliasName, wbm.HeadImg, wbm.QrcodeURL, wbm.ServiceTypeInfo, wbm.VerifyTypeInfo, wbm.AccountExtra, wbm.MenuMediaID, wbm.Appid, wbm.AuthCode, wbm.AccessToken, wbm.RefreshToken, wbm.ExpiresIn, wbm.Extra, wbm.Status, wbm.UpdatedTime, wbm.CreatedTime, wbm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, wbm.UID, wbm.Type, wbm.NickName, wbm.UserName, wbm.AliasName, wbm.HeadImg, wbm.QrcodeURL, wbm.ServiceTypeInfo, wbm.VerifyTypeInfo, wbm.AccountExtra, wbm.MenuMediaID, wbm.Appid, wbm.AuthCode, wbm.AccessToken, wbm.RefreshToken, wbm.ExpiresIn, wbm.Extra, wbm.Status, wbm.UpdatedTime, wbm.CreatedTime, wbm.ID)
	}
	return err
}

// Save saves the WeixinBindMina to the database.
func (wbm *WeixinBindMina) Save(ctx context.Context) error {
	if wbm.Exists() {
		return wbm.Update(ctx)
	}

	return wbm.Insert(ctx)
}

// Delete deletes the WeixinBindMina from the database.
func (wbm *WeixinBindMina) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if wbm._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWeixinBindMinaTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wbm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, wbm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, wbm.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	wbm._deleted = true

	return nil
}

// WeixinBindMinaByID retrieves a row from 'aypcddg.weixin_bind_mina' as a WeixinBindMina.
//
// Generated from index 'weixin_bind_mina_id_pkey'.
func WeixinBindMinaByID(ctx context.Context, id uint, key ...interface{}) (*WeixinBindMina, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetWeixinBindMinaTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, type, nick_name, user_name, alias_name, head_img, qrcode_url, service_type_info, verify_type_info, account_extra, menu_media_id, appid, auth_code, access_token, refresh_token, expires_in, extra, status, updated_time, created_time ` +
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
	wbm := WeixinBindMina{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&wbm.ID, &wbm.UID, &wbm.Type, &wbm.NickName, &wbm.UserName, &wbm.AliasName, &wbm.HeadImg, &wbm.QrcodeURL, &wbm.ServiceTypeInfo, &wbm.VerifyTypeInfo, &wbm.AccountExtra, &wbm.MenuMediaID, &wbm.Appid, &wbm.AuthCode, &wbm.AccessToken, &wbm.RefreshToken, &wbm.ExpiresIn, &wbm.Extra, &wbm.Status, &wbm.UpdatedTime, &wbm.CreatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&wbm.ID, &wbm.UID, &wbm.Type, &wbm.NickName, &wbm.UserName, &wbm.AliasName, &wbm.HeadImg, &wbm.QrcodeURL, &wbm.ServiceTypeInfo, &wbm.VerifyTypeInfo, &wbm.AccountExtra, &wbm.MenuMediaID, &wbm.Appid, &wbm.AuthCode, &wbm.AccessToken, &wbm.RefreshToken, &wbm.ExpiresIn, &wbm.Extra, &wbm.Status, &wbm.UpdatedTime, &wbm.CreatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &wbm, nil
}
