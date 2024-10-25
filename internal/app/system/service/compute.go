package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
)

type (
	ICompute interface {
		StoreComputeTaskToDB(ctx context.Context, data g.Map) (dataAlter g.Map, err error)
		SendReqByComputeType(ctx context.Context, data g.Map) (res interface{}, err error)
		ListCompute(ctx context.Context, data g.Map) (computeData []system.ComputeTask, err error)
		UpdateComputeRegToDB(ctx context.Context, data g.Map) (err error)
		// UpdateResultToDB(ctx context.Context, result_value string, resultID int64) (err error)
		// GetResult(ctx context.Context, TaskID int64) (resultData system.Result, err error)
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
