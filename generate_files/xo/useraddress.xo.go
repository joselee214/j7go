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

// UserAddress represents a row from 'aypcddg.user_address'.
type UserAddress struct {
	Uaid       uint           `json:"uaid"`        // uaid
	UID        uint           `json:"uid"`         // uid
	ProvinceID int16          `json:"province_id"` // province_id
	CityID     int16          `json:"city_id"`     // city_id
	DistrictID int16          `json:"district_id"` // district_id
	Address    string         `json:"address"`     // address
	Street     sql.NullString `json:"street"`      // street
	Postcode   sql.NullString `json:"postcode"`    // postcode
	Contact    sql.NullString `json:"contact"`     // contact
	Phone      sql.NullString `json:"phone"`       // phone
	Mobile     sql.NullString `json:"mobile"`      // mobile
	Email      sql.NullString `json:"email"`       // email
	IsDefault  bool           `json:"is_default"`  // is_default
	Buid       int            `json:"buid"`        // buid

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserAddress exists in the database.
func (ua *UserAddress) Exists() bool { //user_address
	return ua._exists
}

// Deleted provides information if the UserAddress has been deleted from the database.
func (ua *UserAddress) Deleted() bool {
	return ua._deleted
}

// Get table name
func GetUserAddressTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("aypcddg", "user_address", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserAddress to the database.
func (ua *UserAddress) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ua._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserAddressTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, province_id, city_id, district_id, address, street, postcode, contact, phone, mobile, email, is_default, buid` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ua.UID, ua.ProvinceID, ua.CityID, ua.DistrictID, ua.Address, ua.Street, ua.Postcode, ua.Contact, ua.Phone, ua.Mobile, ua.Email, ua.IsDefault, ua.Buid)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ua.UID, ua.ProvinceID, ua.CityID, ua.DistrictID, ua.Address, ua.Street, ua.Postcode, ua.Contact, ua.Phone, ua.Mobile, ua.Email, ua.IsDefault, ua.Buid)
	} else {
		res, err = dbConn.Exec(sqlstr, ua.UID, ua.ProvinceID, ua.CityID, ua.DistrictID, ua.Address, ua.Street, ua.Postcode, ua.Contact, ua.Phone, ua.Mobile, ua.Email, ua.IsDefault, ua.Buid)
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
	ua.Uaid = uint(id)
	ua._exists = true

	return nil
}

// Update updates the UserAddress in the database.
func (ua *UserAddress) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ua._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserAddressTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, province_id = ?, city_id = ?, district_id = ?, address = ?, street = ?, postcode = ?, contact = ?, phone = ?, mobile = ?, email = ?, is_default = ?, buid = ?` +
		` WHERE uaid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ua.UID, ua.ProvinceID, ua.CityID, ua.DistrictID, ua.Address, ua.Street, ua.Postcode, ua.Contact, ua.Phone, ua.Mobile, ua.Email, ua.IsDefault, ua.Buid, ua.Uaid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ua.UID, ua.ProvinceID, ua.CityID, ua.DistrictID, ua.Address, ua.Street, ua.Postcode, ua.Contact, ua.Phone, ua.Mobile, ua.Email, ua.IsDefault, ua.Buid, ua.Uaid)
	} else {
		_, err = dbConn.Exec(sqlstr, ua.UID, ua.ProvinceID, ua.CityID, ua.DistrictID, ua.Address, ua.Street, ua.Postcode, ua.Contact, ua.Phone, ua.Mobile, ua.Email, ua.IsDefault, ua.Buid, ua.Uaid)
	}
	return err
}

// Save saves the UserAddress to the database.
func (ua *UserAddress) Save(ctx context.Context) error {
	if ua.Exists() {
		return ua.Update(ctx)
	}

	return ua.Insert(ctx)
}

// Delete deletes the UserAddress from the database.
func (ua *UserAddress) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ua._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserAddressTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE uaid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ua.Uaid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ua.Uaid)
	} else {
		_, err = dbConn.Exec(sqlstr, ua.Uaid)
	}

	if err != nil {
		return err
	}

	// set deleted
	ua._deleted = true

	return nil
}

// UserAddressesByCityID retrieves a row from 'aypcddg.user_address' as a UserAddress.
//
// Generated from index 'city_id'.
func UserAddressesByCityID(ctx context.Context, cityID int16, key ...interface{}) ([]*UserAddress, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserAddressTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uaid, uid, province_id, city_id, district_id, address, street, postcode, contact, phone, mobile, email, is_default, buid ` +
		`FROM ` + tableName +
		` WHERE city_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, cityID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, cityID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, cityID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserAddress, 0)
	for queryData.Next() {
		ua := UserAddress{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ua.Uaid, &ua.UID, &ua.ProvinceID, &ua.CityID, &ua.DistrictID, &ua.Address, &ua.Street, &ua.Postcode, &ua.Contact, &ua.Phone, &ua.Mobile, &ua.Email, &ua.IsDefault, &ua.Buid)
		if err != nil {
			return nil, err
		}

		res = append(res, &ua)
	}

	return res, nil
}

// UserAddressesByUIDBuid retrieves a row from 'aypcddg.user_address' as a UserAddress.
//
// Generated from index 'idx_uid_buid'.
func UserAddressesByUIDBuid(ctx context.Context, uid uint, buid int, key ...interface{}) ([]*UserAddress, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserAddressTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uaid, uid, province_id, city_id, district_id, address, street, postcode, contact, phone, mobile, email, is_default, buid ` +
		`FROM ` + tableName +
		` WHERE uid = ? AND buid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uid, buid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, uid, buid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, uid, buid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserAddress, 0)
	for queryData.Next() {
		ua := UserAddress{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ua.Uaid, &ua.UID, &ua.ProvinceID, &ua.CityID, &ua.DistrictID, &ua.Address, &ua.Street, &ua.Postcode, &ua.Contact, &ua.Phone, &ua.Mobile, &ua.Email, &ua.IsDefault, &ua.Buid)
		if err != nil {
			return nil, err
		}

		res = append(res, &ua)
	}

	return res, nil
}

// UserAddressesByUID retrieves a row from 'aypcddg.user_address' as a UserAddress.
//
// Generated from index 'uid'.
func UserAddressesByUID(ctx context.Context, uid uint, key ...interface{}) ([]*UserAddress, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserAddressTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uaid, uid, province_id, city_id, district_id, address, street, postcode, contact, phone, mobile, email, is_default, buid ` +
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
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserAddress, 0)
	for queryData.Next() {
		ua := UserAddress{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ua.Uaid, &ua.UID, &ua.ProvinceID, &ua.CityID, &ua.DistrictID, &ua.Address, &ua.Street, &ua.Postcode, &ua.Contact, &ua.Phone, &ua.Mobile, &ua.Email, &ua.IsDefault, &ua.Buid)
		if err != nil {
			return nil, err
		}

		res = append(res, &ua)
	}

	return res, nil
}

// UserAddressByUaid retrieves a row from 'aypcddg.user_address' as a UserAddress.
//
// Generated from index 'user_address_uaid_pkey'.
func UserAddressByUaid(ctx context.Context, uaid uint, key ...interface{}) (*UserAddress, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserAddressTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uaid, uid, province_id, city_id, district_id, address, street, postcode, contact, phone, mobile, email, is_default, buid ` +
		`FROM ` + tableName +
		` WHERE uaid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uaid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ua := UserAddress{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, uaid).Scan(&ua.Uaid, &ua.UID, &ua.ProvinceID, &ua.CityID, &ua.DistrictID, &ua.Address, &ua.Street, &ua.Postcode, &ua.Contact, &ua.Phone, &ua.Mobile, &ua.Email, &ua.IsDefault, &ua.Buid)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, uaid).Scan(&ua.Uaid, &ua.UID, &ua.ProvinceID, &ua.CityID, &ua.DistrictID, &ua.Address, &ua.Street, &ua.Postcode, &ua.Contact, &ua.Phone, &ua.Mobile, &ua.Email, &ua.IsDefault, &ua.Buid)
		if err != nil {
			return nil, err
		}
	}

	return &ua, nil
}
