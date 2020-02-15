package utils

import (
	"context"
	"go.7yes.com/j7f/components/grpc/server"
	"go.uber.org/zap"
	"j7go/components"
	"time"
)

func GetTraceLog(ctx context.Context) *zap.Logger {
	l := ctx.Value(server.UTO_CONTEXT_LOG_KEY)
	if l != nil {
		return l.(*zap.Logger)
	}

	return components.L.Logger
}

// InSliceIface checks given interface in interface slice.
func InSliceIface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceIntersect returns slice that are present in all the slice1 and slice2.
func SliceIntersect(slice1, slice2 []interface{}) (sameSlice []interface{}) {
	for _, v := range slice1 {
		if InSliceIface(v, slice2) {
			sameSlice = append(sameSlice, v)
		}
	}
	return sameSlice
}

// SliceDiff returns diff slice of slice1 - slice2.
func SliceDiff(slice1, slice2 []interface{}) (diffSlice []interface{}) {
	for _, v := range slice1 {
		if !InSliceIface(v, slice2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return diffSlice
}

func GetCurrentUnixTime() uint {
	return uint(time.Now().Unix())
}

func Bool2Int(v bool) int {
	if v {
		return 1
	} else {
		return 0
	}
}