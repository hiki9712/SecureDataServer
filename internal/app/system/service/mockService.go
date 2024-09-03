package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type (
	IMockService interface {
		ResolveReq(ctx context.Context, req *system.MockServiceReq) (data g.Map, err error)
		FetchTable(ctx context.Context, data g.Map) (tableData *model.Table, err error)
		SendToKafka(ctx context.Context, tableData *model.Table) error
	}
)

var (
	localMockService IMockService
)

func MockService() IMockService {
	if localMockService == nil {
		panic("implement not found for interface IMockService")
	}
	return localMockService
}

func RegisterMockService(service IMockService) {
	localMockService = service
}
