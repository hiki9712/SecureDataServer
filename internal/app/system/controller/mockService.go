package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var (
	MockService = mockServiceController{}
)

type mockServiceController struct {
	BaseController
}

func (c *mockServiceController) MockService(ctx context.Context, req *system.MockServiceReq) (res *system.MockServiceRes, err error) {
	var (
		reqData   g.Map
		tableData *model.Table
	)
	reqData, err = service.MockService().ResolveReq(ctx, req)
	if err != nil {
		return
	}
	tableData, err = service.MockService().FetchTable(ctx, reqData)
	if err != nil {
		return
	}
	err = service.MockService().SendToKafka(ctx, tableData)
	if err != nil {
		return
	}
	res = &system.MockServiceRes{
		Status: "success",
		TaskID: 123,
	}
	return
}
