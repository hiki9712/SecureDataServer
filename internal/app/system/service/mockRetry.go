package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type (
	IMockRetry interface {
		ResolveReq(ctx context.Context, req *system.MockRetryReq) (data g.Map, err error)
		FetchTableByTaskID(ctx context.Context, data g.Map) (tableData *model.Table, err error)
		SendToKafka(ctx context.Context, tableData *model.Table) error
	}
)

var (
	localMockRetry IMockRetry
)

func MockRetry() IMockRetry {
	if localMockRetry == nil {
		panic("implement not found for interface IMockRetry")
	}
	return localMockRetry
}

func RegisterMockRetry(service IMockRetry) {
	localMockRetry = service
}
