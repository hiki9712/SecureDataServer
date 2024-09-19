package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	ICompute interface {
		StoreComputeTaskToDB(ctx context.Context, data g.Map) (dataAlter g.Map, err error)
		SendReqByComputeType(ctx context.Context, data g.Map) (res interface{}, err error)
	}
)

var (
	localCompute ICompute
)

func Compute() ICompute {
	if localCompute == nil {
		panic("c")
	}
	return localCompute
}

func ComputeRegister(i ICompute) {
	localCompute = i
}
