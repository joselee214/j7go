package product

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"go.7yes.com/j7f/components/dao"
	"go.7yes.com/j7f/components/grpc/server"
	"go.7yes.com/j7f/components/log"
	"go.7yes.com/j7f/proto/common"
	"go.7yes.com/j7f/proto/images"
	"go.7yes.com/j7f/proto/product"
	"go.uber.org/zap"
	"testing"
	"j7go/components"
)

func TestAddTeamCourseTpl(t *testing.T) {
	InitDbTest()
	InitLogTest()
	Convey("添加团体课", t, func(c C) {
		ctx := context.Background()
		components.L.With(zap.String("trace_id", "123456"))
		context.WithValue(ctx, server.UTO_CONTEXT_LOG_KEY, components.L)
		id, err := AddTeamCourseTpl(ctx, &product.CourseTplRequest{
			Header: &common.CommonHeader{
				TraceId: "123456",
			},
			BrandId:        1,
			CourseName:     "test_course",
			CourseCategory: 1,
			TrainAim:       []uint32{1, 2},
			Duration:       1,
			TimeUnit:       1,
			Description:    "",
			Price:          1,
			StrengthLevel:  1,
			Calories:       1,
			CourseImg: &images.Image{
				ImageId:   0,
				ImageUrl:  "test_url",
				CoverType: 1,
			},
			PublishChannel: 2,
			ShopId:         1,
			IsAvailable:    1,
			OperatorId:     1,
		})
		c.So(err, ShouldBeNil)
		c.So(id, ShouldNotEqual, 0)

	})
}

func TestEditTeamCourseTpl(t *testing.T) {
	Convey("修改团体课", t, func(c C) {
		InitDbTest()
		InitLogTest()
		ctx := context.Background()
		components.L.With(zap.String("trace_id", "123456"))
		context.WithValue(ctx, server.UTO_CONTEXT_LOG_KEY, components.L)
		id, err := EditTeamCourseTpl(ctx, &product.CourseTplRequest{
			Header: &common.CommonHeader{
				TraceId: "123456",
			},
			CourseId:       1,
			BrandId:        1,
			CourseName:     "test_course",
			CourseCategory: 1,
			TrainAim:       []uint32{3, 2},
			Duration:       1,
			TimeUnit:       1,
			Description:    "",
			Price:          1,
			StrengthLevel:  1,
			Calories:       1,
			CourseImg: &images.Image{
				ImageId:   0,
				ImageUrl:  "test_url",
				CoverType: 1,
			},
		})
		c.So(err, ShouldBeNil)
		c.So(id, ShouldNotEqual, 0)
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
