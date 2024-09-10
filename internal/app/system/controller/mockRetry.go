package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var (
	MockRetry = mockRetryController{}
)

type mockRetryController struct {
	BaseController
}

func (c *mockRetryController) MockRetry(ctx context.Context, req *system.MockRetryReq) (res *system.MockRetryRes, err error) {
	var (
		reqData   g.Map
		tableData *model.Table
		TaskID    int64
	)
	reqData, err = service.MockRetry().ResolveReq(ctx, req)
	g.Log().Info(ctx, "mock retry req:", reqData)
	if err != nil {
		return
	}
	TaskID = reqData["task_id"].(int64)
	res = &system.MockRetryRes{
		Status:   "success",
		TaskID:   TaskID,
		Message:  "",
		TryAgain: true,
	}
	tableData, err = service.MockRetry().FetchTableByTaskID(ctx, reqData)
	if err != nil {
		return
	}
	g.Log().Info(ctx, "mock retry table Data:", tableData)
	//err = service.MockRetry().SendToKafka(ctx, tableData)
	//if err != nil {
	//	res.TryAgain = false
	//	return
	//}
	return
}
