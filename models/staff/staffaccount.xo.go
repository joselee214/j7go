// Package test contains the types for schema 'saas'.
package staff

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.7yes.com/j7f/proto/common"
	"j7go/components"
	"j7go/utils"

	"go.uber.org/zap"
)

// StaffAccount represents a row from 'saas.staff_account'.
type StaffAccount struct {
	ID            uint   `json:"id"`              // id
	AccountType   int8   `json:"account_type"`    // account_type
	StaffID       uint   `json:"staff_id"`        // staff_id
	AccountName   string `json:"account_name"`    // account_name
	AccountPwd    string `json:"account_pwd"`     // account_pwd
	Mail          string `json:"mail"`            // mail
	Phone         string `json:"phone"`           // phone
	CountryCodeID uint   `json:"country_code_id"` // country_code_id
	PhoneCode     int16  `json:"phone_code"`      // phone_code
	ShortName     string `json:"short_name"`      // short_name
	IsUpdateName  int8   `json:"is_update_name"`  // is_update_name
	IsDel         int8   `json:"is_del"`          // is_del
	UpdatedTime   uint   `json:"updated_time"`    // updated_time
	CreatedTime   uint   `json:"created_time"`    // created_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the StaffAccount exists in the database.
func (sa *StaffAccount) Exists() bool { //staff_account
	return sa._exists
}

// Deleted provides information if the StaffAccount has been deleted from the database.
func (sa *StaffAccount) Deleted() bool {
	return sa._deleted
}

// Get table name
func GetStaffAccountTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "staff_account", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the StaffAccount to the database.
func (sa *StaffAccount) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if sa._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetStaffAccountTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`account_type, staff_id, account_name, account_pwd, mail, phone, country_code_id, phone_code, short_name, is_update_name, is_del, updated_time, created_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, sa.AccountType, sa.StaffID, sa.AccountName, sa.AccountPwd, sa.Mail, sa.Phone, sa.CountryCodeID, sa.PhoneCode, sa.ShortName, sa.IsUpdateName, sa.IsDel, sa.UpdatedTime, sa.CreatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, sa.AccountType, sa.StaffID, sa.AccountName, sa.AccountPwd, sa.Mail, sa.Phone, sa.CountryCodeID, sa.PhoneCode, sa.ShortName, sa.IsUpdateName, sa.IsDel, sa.UpdatedTime, sa.CreatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, sa.AccountType, sa.StaffID, sa.AccountName, sa.AccountPwd, sa.Mail, sa.Phone, sa.CountryCodeID, sa.PhoneCode, sa.ShortName, sa.IsUpdateName, sa.IsDel, sa.UpdatedTime, sa.CreatedTime)
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
	sa.ID = uint(id)
	sa._exists = true

	return nil
}

// Update updates the StaffAccount in the database.
func (sa *StaffAccount) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if sa._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetStaffAccountTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`account_type = ?, staff_id = ?, account_name = ?, account_pwd = ?, mail = ?, phone = ?, country_code_id = ?, phone_code = ?, short_name = ?, is_update_name = ?, is_del = ?, updated_time = ?, created_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, sa.AccountType, sa.StaffID, sa.AccountName, sa.AccountPwd, sa.Mail, sa.Phone, sa.CountryCodeID, sa.PhoneCode, sa.ShortName, sa.IsUpdateName, sa.IsDel, sa.UpdatedTime, sa.CreatedTime, sa.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, sa.AccountType, sa.StaffID, sa.AccountName, sa.AccountPwd, sa.Mail, sa.Phone, sa.CountryCodeID, sa.PhoneCode, sa.ShortName, sa.IsUpdateName, sa.IsDel, sa.UpdatedTime, sa.CreatedTime, sa.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, sa.AccountType, sa.StaffID, sa.AccountName, sa.AccountPwd, sa.Mail, sa.Phone, sa.CountryCodeID, sa.PhoneCode, sa.ShortName, sa.IsUpdateName, sa.IsDel, sa.UpdatedTime, sa.CreatedTime, sa.ID)
	}
	return err
}

// Save saves the StaffAccount to the database.
func (sa *StaffAccount) Save(ctx context.Context) error {
	if sa.Exists() {
		return sa.Update(ctx)
	}

	return sa.Insert(ctx)
}

// Delete deletes the StaffAccount from the database.
func (sa *StaffAccount) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if sa._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetStaffAccountTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, sa.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, sa.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, sa.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	sa._deleted = true

	return nil
}

// StaffAccountByID retrieves a row from 'saas.staff_account' as a StaffAccount.
//
// Generated from index 'staff_account_id_pkey'.
func StaffAccountByID(ctx context.Context, id uint, key ...interface{}) (*StaffAccount, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetStaffAccountTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, account_type, staff_id, account_name, account_pwd, mail, phone, country_code_id, phone_code, short_name, is_update_name, is_del, updated_time, created_time ` +
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
	sa := StaffAccount{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&sa.ID, &sa.AccountType, &sa.StaffID, &sa.AccountName, &sa.AccountPwd, &sa.Mail, &sa.Phone, &sa.CountryCodeID, &sa.PhoneCode, &sa.ShortName, &sa.IsUpdateName, &sa.IsDel, &sa.UpdatedTime, &sa.CreatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&sa.ID, &sa.AccountType, &sa.StaffID, &sa.AccountName, &sa.AccountPwd, &sa.Mail, &sa.Phone, &sa.CountryCodeID, &sa.PhoneCode, &sa.ShortName, &sa.IsUpdateName, &sa.IsDel, &sa.UpdatedTime, &sa.CreatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &sa, nil
}

//获取员工账号信息
func GetStaffAccount(ctx context.Context, staffId uint) (*StaffAccount, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetStaffAccountTableName()
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, account_type, staff_id, account_name, account_pwd, mail, phone, country_code_id, phone_code, short_name, is_update_name, is_del, updated_time, created_time ` +
		`FROM ` + tableName +
		` WHERE staff_id = ? and is_del = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, staffId, common.DelStatus_NOT_DEL)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	sa := StaffAccount{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, staffId, common.DelStatus_NOT_DEL).Scan(&sa.ID, &sa.AccountType, &sa.StaffID, &sa.AccountName, &sa.AccountPwd, &sa.Mail, &sa.Phone, &sa.CountryCodeID, &sa.PhoneCode, &sa.ShortName, &sa.IsUpdateName, &sa.IsDel, &sa.UpdatedTime, &sa.CreatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, staffId, common.DelStatus_NOT_DEL).Scan(&sa.ID, &sa.AccountType, &sa.StaffID, &sa.AccountName, &sa.AccountPwd, &sa.Mail, &sa.Phone, &sa.CountryCodeID, &sa.PhoneCode, &sa.ShortName, &sa.IsUpdateName, &sa.IsDel, &sa.UpdatedTime, &sa.CreatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &sa, nil
}
