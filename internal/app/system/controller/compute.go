package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
)

var (
	Compute = computeController{}
)

type computeController struct {
	BaseController
}

func (c *computeController) TaskList(ctx context.Context, req *system.ComputeTaskListReq) (res *system.ComputeTaskListRes, err error) {
	var (
		data g.Map
	)
	res = &system.ComputeTaskListRes{
		Status:  "fail",
		Message: "",
	}
	data, err = libUtils.ResolveReq(ctx, req)
	if err != nil {
		return
	}
	dataList, err := service.Compute().ListCompute(ctx, data)
	if err != nil {
		return
	}
	res.Status = "success"
	res.Data = dataList
	return
}

func (c *computeController) SendRequest(ctx context.Context, req *system.ComputeSendReq) (res *system.ComputeSendRes, err error) {
	var (
		data g.Map
	)
	res = &system.ComputeSendRes{
		Status:  "fail",
		Message: "",
	}
	data, err = libUtils.ResolveReq(ctx, req)
	if err != nil {
		return
	}
	data, err = service.Compute().StoreComputeTaskToDB(ctx, data)
	if err != nil {
		return
	}
	g.Log().Info(ctx, "compute info:", data)
	_, err = service.Compute().SendReqByComputeType(ctx, data)
	if err != nil {
		return
	}
	res.Status = "success"
	return
}

func (c *computeController) GetResult(ctx context.Context, req *system.ComputeResultReq) (res *system.ComputeResultRes, err error) {
	var (
		data g.Map
	)
	res = &system.ComputeResultRes{
		Status:  "fail",
		Message: "",
	}
	data, err = libUtils.ResolveReq(ctx, req)
	if err != nil {
		return
	}
	g.Log().Info(ctx, "compute result:", data)
	return
}
