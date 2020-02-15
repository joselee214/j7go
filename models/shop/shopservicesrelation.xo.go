// Package xo contains the types for schema 'saas'.
package shopModel

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

// ShopServicesRelation represents a row from 'saas.shop_services_relation'.
type ShopServicesRelation struct {
	ID          uint `json:"id"`           // id
	ShopID      uint `json:"shop_id"`      // shop_id
	ServiceID   uint `json:"service_id"`   // service_id
	IsDel       int8 `json:"is_del"`       // is_del
	CreatedTime uint `json:"created_time"` // created_time
	UpdatedTime uint `json:"updated_time"` // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ShopServicesRelation exists in the database.
func (ssr *ShopServicesRelation) Exists() bool { //shop_services_relation
	return ssr._exists
}

// Deleted provides information if the ShopServicesRelation has been deleted from the database.
func (ssr *ShopServicesRelation) Deleted() bool {
	return ssr._deleted
}

// Get table name
func getShopServicesRelationTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "shop_services_relation", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ShopServicesRelation to the database.
func (ssr *ShopServicesRelation) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	// if already exist, bail
	if ssr._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := getShopServicesRelationTableName(key...)
	if err != nil {
		return err
	}

	var res sql.Result
	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`shop_id, service_id, is_del, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ssr.ShopID, ssr.ServiceID, ssr.IsDel, ssr.CreatedTime, ssr.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ssr.ShopID, ssr.ServiceID, ssr.IsDel, ssr.CreatedTime, ssr.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, ssr.ShopID, ssr.ServiceID, ssr.IsDel, ssr.CreatedTime, ssr.UpdatedTime)
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ssr.ID = uint(id)
	ssr._exists = true

	return nil
}

// Update updates the ShopServicesRelation in the database.
func (ssr *ShopServicesRelation) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ssr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := getShopServicesRelationTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`shop_id = ?, service_id = ?, is_del = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ssr.ShopID, ssr.ServiceID, ssr.IsDel, ssr.CreatedTime, ssr.UpdatedTime, ssr.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ssr.ShopID, ssr.ServiceID, ssr.IsDel, ssr.CreatedTime, ssr.UpdatedTime, ssr.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ssr.ShopID, ssr.ServiceID, ssr.IsDel, ssr.CreatedTime, ssr.UpdatedTime, ssr.ID)
	}
	return err
}

// Save saves the ShopServicesRelation to the database.
func (ssr *ShopServicesRelation) Save(ctx context.Context) error {
	if ssr.Exists() {
		return ssr.Update(ctx)
	}

	return ssr.Insert(ctx)
}

// Delete deletes the ShopServicesRelation from the database.
func (ssr *ShopServicesRelation) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ssr._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := getShopServicesRelationTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = ? WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ssr.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, common.DelStatus_DELED, ssr.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, common.DelStatus_DELED, ssr.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ssr._deleted = true

	return nil
}

// ShopServicesRelationByID retrieves a row from 'saas.shop_services_relation' as a ShopServicesRelation.
//
// Generated from index 'shop_services_relation_id_pkey'.
func ShopServicesRelationByID(ctx context.Context, id uint, key ...interface{}) (*ShopServicesRelation, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := getShopServicesRelationTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, shop_id, service_id, is_del, created_time, updated_time ` +
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
	ssr := ShopServicesRelation{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&ssr.ID, &ssr.ShopID, &ssr.ServiceID, &ssr.IsDel, &ssr.CreatedTime, &ssr.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&ssr.ID, &ssr.ShopID, &ssr.ServiceID, &ssr.IsDel, &ssr.CreatedTime, &ssr.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &ssr, nil
}

func ShopServicesRelationByShopId(ctx context.Context, shopId uint, key ...interface{}) ([]*ShopServicesRelation, error) {
	var rows *sql.Rows

	tableName, err := getShopServicesRelationTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, shop_id, service_id, is_del, created_time, updated_time ` +
		`FROM ` + tableName +
		` WHERE shop_id = ? AND is_del = ?`

	// run query
	components.L.Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, shopId)))

	var dbConn *sql.DB
	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}

	if tx != nil {
		rows, err = tx.Query(sqlstr, shopId, common.DelStatus_NOT_DEL)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = dbConn.Query(sqlstr, shopId, common.DelStatus_NOT_DEL)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	shopServicesRelation := make([]*ShopServicesRelation, 0)
	for rows.Next() {
		ssr := &ShopServicesRelation{
			_exists: true,
		}

		err = rows.Scan(&ssr.ID, &ssr.ShopID, &ssr.ServiceID, &ssr.IsDel, &ssr.CreatedTime, &ssr.UpdatedTime)
		if err != nil {
			return nil, err
		}
		shopServicesRelation = append(shopServicesRelation, ssr)
	}

	if err != nil {
		return nil, err
	}

	return shopServicesRelation, nil
}
