package shopService

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"go.7yes.com/j7f/components/dao"
	"j7go/components"
	"go.7yes.com/j7f/components/log"
	"go.uber.org/zap"
	"context"
	"go.7yes.com/j7f/components/grpc/server"
	"go.7yes.com/j7f/proto/shop"
	"go.7yes.com/j7f/proto/common"
)

func TestSaveShopArea(t *testing.T) {
	Convey("新增/编辑场地信息", t, func(c C) {
		InitDbTest()
		InitLogTest()
		ctx := context.Background()
		components.L.With(zap.String("trace_id", "123456"))
		context.WithValue(ctx, server.UTO_CONTEXT_LOG_KEY, components.L)
		Convey("新增场地信息", func() {
			shopAreaId, err := SaveShopArea(ctx, &shop.ShopAreaRequest{
				Header: &common.CommonHeader{
					TraceId: "123456",
				},
				BrandId:       1,
				ShopId:        1,
				AreaName:      "test_area",
				ContainNumber: 1,
				IsVip:         1,
			})
			c.So(err, ShouldBeNil)
			c.So(shopAreaId, ShouldNotEqual, 0)

		})
		Convey("编辑场地信息", func() {
			shopAreaId, err := SaveShopArea(ctx, &shop.ShopAreaRequest{
				Header: &common.CommonHeader{
					TraceId: "123456",
				},
				ShopAreaId:    1,
				BrandId:       1,
				ShopId:        1,
				AreaName:      "test_area",
				ContainNumber: 1,
				IsVip:         1,
			})
			c.So(err, ShouldBeNil)
			c.So(shopAreaId, ShouldNotEqual, 0)
		})
	})
}

func TestDelShopArea(t *testing.T) {
	Convey("删除场地信息", t, func(c C) {
		InitDbTest()
		InitLogTest()
		ctx := context.Background()
		components.L.With(zap.String("trace_id", "123456"))
		context.WithValue(ctx, server.UTO_CONTEXT_LOG_KEY, components.L)
		shopAreaId, err := DelShopArea(ctx, &shop.ShopAreaRequest{
			Header: &common.CommonHeader{
				TraceId: "123456",
			},
			ShopAreaId: 1,
		})
		c.So(err, ShouldBeNil)
		c.So(shopAreaId, ShouldNotEqual, 0)
	})
}

func TestGetShopArea(t *testing.T) {
	Convey("获取场地信息", t, func(c C) {
		InitDbTest()
		InitLogTest()
		ctx := context.Background()
		components.L.With(zap.String("trace_id", "123456"))
		context.WithValue(ctx, server.UTO_CONTEXT_LOG_KEY, components.L)
		shopAreaId, err := GetShopArea(ctx, &shop.ShopAreaRequest{
			Header: &common.CommonHeader{
				TraceId: "123456",
			},
			ShopAreaId: 1,
		})
		c.So(err, ShouldBeNil)
		c.So(shopAreaId, ShouldNotEqual, 0)
	})
}

func InitDbTest() {
	var DbConfig = &dao.DBConfig{
		Name:         "",
		MaxIdleConns: 5,
		MaxConnNum:   10,
		Master: &dao.NodeConfig{
			Addr:     "127.0.0.1:3307",
			User:     "root",
			Password: "123456",
			Weight:   1,
		},
		Slave: []*dao.NodeConfig{
			{
				Addr:     "127.0.0.1:3307",
				User:     "root",
				Password: "123456",
				Weight:   1,
			},
		},
	}

	components.InitDB(DbConfig)
}

func InitLogTest() {
	var logConfig = &log.Config{
		Level:    "debug",
		Encoding: "console",
		EncoderConfig: log.EncoderConfig{
			MessageKey: "",
			LevelKey:   "",
			NameKey:    "",
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	components.InitLog(logConfig)
}
