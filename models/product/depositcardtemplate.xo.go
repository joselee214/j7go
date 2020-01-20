// Package tpl contains the types for schema 'saas'.
package product

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"j7go/components"
	"j7go/utils"
)

// DepositCardTemplate represents a row from 'saas.deposit_card_template'.
type DepositCardTemplate struct {
	ID               uint   `json:"id"`                // id
	BrandID          uint   `json:"brand_id"`          // brand_id
	ShopID           uint   `json:"shop_id"`           // shop_id
	CardName         string `json:"card_name"`         // card_name
	CardPrice        uint   `json:"card_price"`        // card_price
	SellPrice        uint   `json:"sell_price"`        // sell_price
	Unit             int8   `json:"unit"`              // unit
	Num              uint   `json:"num"`               // num
	PublishChannel   int8   `json:"publish_channel"`   // publish_channel
	SellStatus       int8   `json:"sell_status"`       // sell_status
	ConsumptionRange int8   `json:"consumption_range"` // consumption_range
	SellType         int8   `json:"sell_type"`         // sell_type
	SupportSales     int8   `json:"support_sales"`     // support_sales
	StartTime        uint   `json:"start_time"`        // start_time
	EndTime          uint   `json:"end_time"`          // end_time
	CardContents     string `json:"card_contents"`     // card_contents
	IsDel            int8   `json:"is_del"`            // is_del
	AlbumID          uint   `json:"album_id"`          // album_id
	Description      string `json:"description"`       // description
	OperatorID       uint   `json:"operator_id"`       // operator_id
	CreatedTime      uint   `json:"created_time"`      // created_time
	UpdatedTime      uint   `json:"updated_time"`      // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DepositCardTemplate exists in the database.
func (dct *DepositCardTemplate) Exists() bool { //deposit_card_template
	return dct._exists
}

// Deleted provides information if the DepositCardTemplate has been deleted from the database.
func (dct *DepositCardTemplate) Deleted() bool {
	return dct._deleted
}

// Get table name
func GetDepositCardTemplateTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "deposit_card_template", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DepositCardTemplate to the database.
func (dct *DepositCardTemplate) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if dct._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDepositCardTemplateTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`brand_id, shop_id, card_name, card_price, sell_price, unit, num, publish_channel, sell_status, consumption_range, support_sales, sell_type, start_time, end_time, card_contents, is_del, album_id, description, operator_id, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dct.BrandID, dct.ShopID, dct.CardName, dct.CardPrice, dct.SellPrice, dct.Unit, dct.Num, dct.PublishChannel, dct.SellStatus, dct.ConsumptionRange, dct.SupportSales, dct.SellType, dct.StartTime, dct.EndTime, dct.CardContents, dct.IsDel, dct.AlbumID, dct.Description, dct.OperatorID, dct.CreatedTime, dct.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dct.BrandID, dct.ShopID, dct.CardName, dct.CardPrice, dct.SellPrice, dct.Unit, dct.Num, dct.PublishChannel, dct.SellStatus, dct.ConsumptionRange, dct.SupportSales, dct.SellType, dct.StartTime, dct.EndTime, dct.CardContents, dct.IsDel, dct.AlbumID, dct.Description, dct.OperatorID, dct.CreatedTime, dct.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, dct.BrandID, dct.ShopID, dct.CardName, dct.CardPrice, dct.SellPrice, dct.Unit, dct.Num, dct.PublishChannel, dct.SellStatus, dct.ConsumptionRange, dct.SupportSales, dct.SellType, dct.StartTime, dct.EndTime, dct.CardContents, dct.IsDel, dct.AlbumID, dct.Description, dct.OperatorID, dct.CreatedTime, dct.UpdatedTime)
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
	dct.ID = uint(id)
	dct._exists = true

	return nil
}

// Update updates the DepositCardTemplate in the database.
func (dct *DepositCardTemplate) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dct._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDepositCardTemplateTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`brand_id = ?, shop_id = ?, card_name = ?, card_price = ?, sell_price = ?, unit = ?, num = ?, publish_channel = ?,sell_status = ?, consumption_range = ?, support_sales = ?, start_time = ?, end_time = ?, card_contents = ?, is_del = ?, album_id = ?, description = ?, operator_id = ?, sell_type = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dct.BrandID, dct.ShopID, dct.CardName, dct.CardPrice, dct.SellPrice, dct.Unit, dct.Num, dct.PublishChannel, dct.SellStatus, dct.ConsumptionRange, dct.SupportSales, dct.StartTime, dct.EndTime, dct.CardContents, dct.IsDel, dct.AlbumID, dct.Description, dct.OperatorID, dct.SellType, dct.CreatedTime, dct.UpdatedTime, dct.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dct.BrandID, dct.ShopID, dct.CardName, dct.CardPrice, dct.SellPrice, dct.Unit, dct.Num, dct.PublishChannel, dct.SellStatus, dct.ConsumptionRange, dct.SupportSales, dct.StartTime, dct.EndTime, dct.CardContents, dct.IsDel, dct.AlbumID, dct.Description, dct.OperatorID, dct.SellType, dct.CreatedTime, dct.UpdatedTime, dct.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dct.BrandID, dct.ShopID, dct.CardName, dct.CardPrice, dct.SellPrice, dct.Unit, dct.Num, dct.PublishChannel, dct.SellStatus, dct.ConsumptionRange, dct.SupportSales, dct.StartTime, dct.EndTime, dct.CardContents, dct.IsDel, dct.AlbumID, dct.Description, dct.OperatorID, dct.SellType, dct.CreatedTime, dct.UpdatedTime, dct.ID)
	}
	return err
}

// Save saves the DepositCardTemplate to the database.
func (dct *DepositCardTemplate) Save(ctx context.Context) error {
	if dct.Exists() {
		return dct.Update(ctx)
	}

	return dct.Insert(ctx)
}

// Delete deletes the DepositCardTemplate from the database.
func (dct *DepositCardTemplate) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dct._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDepositCardTemplateTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dct.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dct.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dct.ID)
	}

	if err != nil {
		return err
	}
	return nil
}

// DepositCardTemplateByID retrieves a row from 'saas.deposit_card_template' as a DepositCardTemplate.
//
// Generated from index 'deposit_card_template_id_pkey'.
func DepositCardTemplateByID(ctx context.Context, id uint, key ...interface{}) (*DepositCardTemplate, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDepositCardTemplateTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, brand_id, shop_id, card_name, card_price, sell_type,sell_price, unit, num, publish_channel, sell_status, consumption_range, support_sales, start_time, end_time, card_contents, is_del, album_id, description, operator_id, created_time, updated_time ` +
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
	dct := DepositCardTemplate{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&dct.ID, &dct.BrandID, &dct.ShopID, &dct.CardName, &dct.CardPrice, &dct.SellType, &dct.SellPrice, &dct.Unit, &dct.Num, &dct.PublishChannel, &dct.SellStatus, &dct.ConsumptionRange, &dct.SupportSales, &dct.StartTime, &dct.EndTime, &dct.CardContents, &dct.IsDel, &dct.AlbumID, &dct.Description, &dct.OperatorID, &dct.CreatedTime, &dct.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&dct.ID, &dct.BrandID, &dct.ShopID, &dct.CardName, &dct.CardPrice, &dct.SellType, &dct.SellPrice, &dct.Unit, &dct.Num, &dct.PublishChannel, &dct.SellStatus, &dct.ConsumptionRange, &dct.SupportSales, &dct.StartTime, &dct.EndTime, &dct.CardContents, &dct.IsDel, &dct.AlbumID, &dct.Description, &dct.OperatorID, &dct.CreatedTime, &dct.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &dct, nil
}