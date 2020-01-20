// Package product contains the types for schema 'saas'.
package product

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"j7go/components"
	"j7go/models/shop"
	"j7go/utils"
)

// ProductSupportShop represents a row from 'saas.product_support_shop'.
type ProductSupportShop struct {
	ID          uint `json:"id"`           // id
	BrandID     uint `json:"brand_id"`     // brand_id
	ShopID      uint `json:"shop_id"`      // shop_id
	ProductID   uint `json:"product_id"`   // product_id
	ProductType int8 `json:"product_type"` // product_type
	IsDel       int8 `json:"is_del"`       // is_del
	CreatedTime uint `json:"created_time"` // created_time
	UpdatedTime uint `json:"updated_time"` // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ProductSupportShop exists in the database.
func (pss *ProductSupportShop) Exists() bool { //product_support_shop
	return pss._exists
}

// Deleted provides information if the ProductSupportShop has been deleted from the database.
func (pss *ProductSupportShop) Deleted() bool {
	return pss._deleted
}

// Get table name
func GetProductSupportShopTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("saas", "product_support_shop", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ProductSupportShop to the database.
func (pss *ProductSupportShop) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if pss._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetProductSupportShopTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`brand_id, shop_id, product_id, product_type, is_del, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pss.BrandID, pss.ShopID, pss.ProductID, pss.ProductType, pss.IsDel, pss.CreatedTime, pss.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, pss.BrandID, pss.ShopID, pss.ProductID, pss.ProductType, pss.IsDel, pss.CreatedTime, pss.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, pss.BrandID, pss.ShopID, pss.ProductID, pss.ProductType, pss.IsDel, pss.CreatedTime, pss.UpdatedTime)
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	pss.ID = uint(id)
	pss._exists = true

	return nil
}

// BatchInsert inserts the CardPriceSetting to the database.
func ProductSupportShopBatchInsert(ctx context.Context, shopList []*ProductSupportShop, key ...interface{}) error {
	if utils.IntZero == len(shopList) {
		return nil
	}
	var err error
	var dbConn *sql.DB

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetProductSupportShopTableName(key...)

	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlBuilder := squirrel.
		Insert(tableName).
		Columns("brand_id", "shop_id", "product_id", "product_type", "is_del", "created_time", "updated_time")

	for _, shop := range shopList {
		sqlBuilder = sqlBuilder.Values(shop.BrandID, shop.ShopID, shop.ProductID,
			shop.ProductType, shop.IsDel, shop.CreatedTime, shop.UpdatedTime)
	}

	sqlstr, args, err := sqlBuilder.ToSql()
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

// Update updates the ProductSupportShop in the database.
func (pss *ProductSupportShop) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if pss._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetProductSupportShopTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`brand_id = ?, shop_id = ?, product_id = ?, product_type = ?, is_del = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pss.BrandID, pss.ShopID, pss.ProductID, pss.ProductType, pss.IsDel, pss.CreatedTime, pss.UpdatedTime, pss.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, pss.BrandID, pss.ShopID, pss.ProductID, pss.ProductType, pss.IsDel, pss.CreatedTime, pss.UpdatedTime, pss.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, pss.BrandID, pss.ShopID, pss.ProductID, pss.ProductType, pss.IsDel, pss.CreatedTime, pss.UpdatedTime, pss.ID)
	}
	return err
}

// Save saves the ProductSupportShop to the database.
func (pss *ProductSupportShop) Save(ctx context.Context) error {
	if pss.Exists() {
		return pss.Update(ctx)
	}

	return pss.Insert(ctx)
}

// Delete deletes the ProductSupportShop from the database.
func (pss *ProductSupportShop) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if pss._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetProductSupportShopTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pss.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, pss.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, pss.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	pss._deleted = true

	return nil
}

// ProductSupportShopByID retrieves a row from 'saas.product_support_shop' as a ProductSupportShop.
//
// Generated from index 'product_support_shop_id_pkey'.
func ProductSupportShopByID(ctx context.Context, id uint, key ...interface{}) (*ProductSupportShop, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetProductSupportShopTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, brand_id, shop_id, product_id, product_type, is_del, created_time, updated_time ` +
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
	pss := ProductSupportShop{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&pss.ID, &pss.BrandID, &pss.ShopID, &pss.ProductID, &pss.ProductType, &pss.IsDel, &pss.CreatedTime, &pss.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&pss.ID, &pss.BrandID, &pss.ShopID, &pss.ProductID, &pss.ProductType, &pss.IsDel, &pss.CreatedTime, &pss.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &pss, nil
}

//获取商品支持使用门店
func ProductSupportShopByProductId(ctx context.Context, productId uint, productType int8, key ...interface{}) ([]*shopModel.Shop, error) {
	var err error
	var dbConn *sql.DB
	var list = make([]*shopModel.Shop, 0)
	var rows *sql.Rows

	shopTableName, err := shopModel.GetShopTableName(key...)
	if err != nil {
		return nil, err
	}

	supportShopTableName, err := GetProductSupportShopTableName(key...)
	if err != nil {
		return nil, err
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}

	sqlBuilder :=
		squirrel.Select(
			"s.id",
			"s.brand_id",
			"s.shop_name",
			"s.address",
			"s.province_id",
			"s.city_id",
			"s.district_id",
			"s.province_name",
			"s.city_name",
			"s.district_name").
			From(shopTableName + " s").
			Join(supportShopTableName + " ss on s.id = ss.shop_id").
			Where(squirrel.Eq{"ss.product_type": productType,
				"ss.is_del": utils.NOT_DELETED, "ss.product_id": productId, "s.is_del": utils.NOT_DELETED})

	sqlStr, arg, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlStr, arg)))
	if tx != nil {
		rows, err = tx.Query(sqlStr, arg...)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		rows, err = dbConn.Query(sqlStr, arg...)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
	}

	defer rows.Close()

	for rows.Next() {
		shopInfo := shopModel.Shop{}
		err := rows.Scan(&shopInfo.ID,
			&shopInfo.BrandID,
			&shopInfo.ShopName,
			&shopInfo.Address,
			&shopInfo.ProvinceID,
			&shopInfo.CityID,
			&shopInfo.DistrictID,
			&shopInfo.ProvinceName,
			&shopInfo.CityName,
			&shopInfo.DistrictName,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &shopInfo)
	}

	return list, nil
}

func ProductSupportShopBatchDelete(ctx context.Context, productId uint, productType int8, shopList []uint, key ...interface{}) error {
	if utils.IntZero == len(shopList) {
		return nil
	}
	var err error
	var dbConn *sql.DB

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetProductSupportShopTableName(key...)
	if err != nil {
		return err
	}

	sqlBuilder := squirrel.
		Update(tableName).
		Set("is_del", utils.Available).
		Where(squirrel.Eq{"product_id": productId, "product_type": productType, "shop_id": shopList, "is_del": utils.Unavailable})

	sqlStr, args, err := sqlBuilder.ToSql()
	if err != nil {
		return err
	}

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlStr, args)))
	if tx != nil {
		_, err = tx.Exec(sqlStr, args...)
	} else {
		_, err = dbConn.Exec(sqlStr, args...)
	}

	if err != nil {
		return err
	}

	return nil
}

//删除所有支持的门店
func DeleteAllSupportShopByProductId(ctx context.Context, productId uint, productType int8, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetProductSupportShopTableName(key...)
	if err != nil {
		return err
	}

	sqlBuilder := squirrel.
		Update(tableName).
		Set("is_del", utils.DELETED).
		Where(squirrel.Eq{"product_id": productId, "product_type": productType, "is_del": utils.NOT_DELETED})

	sqlStr, args, err := sqlBuilder.ToSql()
	if err != nil {
		return err
	}

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlStr, args)))
	if tx != nil {
		_, err = tx.Exec(sqlStr, args...)
	} else {
		_, err = dbConn.Exec(sqlStr, args...)
	}

	if err != nil {
		return err
	}

	return nil
}
