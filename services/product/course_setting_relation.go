package product

import (
	"context"
	"j7go/utils"
	"j7go/models/product"
	"math"
	"math/big"
)

func CourseSettingRelationBatchInsert(ctx context.Context, settingIds []uint32, courseId uint, courseType int8) error {
	idLen := len(settingIds)
	f, exact := big.NewRat(int64(idLen),utils.BATCH_INSERT_MAX).Float64()
	if exact == false {
		f = math.Ceil(f)
	}
	for i := 0; i < int(f); i++ {
		ids := settingIds[i * utils.BATCH_INSERT_MAX:]
		if len(ids) > utils.BATCH_INSERT_MAX {
			ids = settingIds[i * utils.BATCH_INSERT_MAX : (i+1) * utils.BATCH_INSERT_MAX]
		}
		err := product.CourseSettingRelationBatchInsert(ctx,courseId, courseType, settingIds)
		if err != nil {
			return err
		}
	}

	return nil
}

func CourseSettingRelationBatchDelete(ctx context.Context, settingIds []uint, courseId uint, courseType int8) error {
	return product.CourseSettingRelationBatchDelete(ctx, courseId, courseType, settingIds)
}