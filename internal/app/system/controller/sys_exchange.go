package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	_ "github.com/tiger1103/gfast/v3/internal/app/system/logic"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var (
	Exchange = exchangeController{}
)

type exchangeController struct {
	BaseController
}

func (c *exchangeController) SendExchangeRequest(ctx context.Context, req *system.ExchangeReq) (res *system.ExchangeRes, err error) {
	var (
		data    g.Map
		message string
	)
	res = &system.ExchangeRes{
		Status:  "fail",
		Message: "",
	}
	data, err = service.SysExchange().ResolveReq(ctx, req)
	g.Log().Info(ctx, "exchange info", data)
	message, err = service.SysExchange().StoreExchangeTaskToDB(ctx, data)
	if err != nil {
		g.Log().Error(ctx, err)
		res.Message = err.Error()
		return
	}
	if message != "" {
		res.Message = message
		return
	}
	res.Status = "success"
	return
}

func (c *exchangeController) SendData(ctx context.Context, req *system.SendDataReq) (res *system.SendDataRes, err error) {
	res = &system.SendDataRes{
		Status:  "fail",
		Message: "",
	}
	data, err := service.SysExchange().ResolveReq(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	tableData, err := service.SysExchange().FetchTable(ctx, data)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = service.SysExchange().SendToMasking(ctx, tableData)
	if err != nil {
		return
	}
	res.Status = "success"
	return
}
