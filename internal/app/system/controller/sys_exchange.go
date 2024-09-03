package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	_ "github.com/tiger1103/gfast/v3/internal/app/system/logic"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"math/rand/v2"
)

var (
	Exchange = exchangeController{}
)

type exchangeController struct {
	BaseController
}

func (c *exchangeController) Exchange(ctx context.Context, req *system.ExchangeReq) (res *system.ExchangeRes, err error) {
	var (
		data g.Map
	)
	res = &system.ExchangeRes{
		Status:  "fail",
		Message: "",
	}
	data, err = service.SysExchange().ResolveReq(ctx, req)
	g.Log().Info(ctx, "exchange info", data)
	err = service.SysExchange().StoreExchangeTaskToDB(ctx, data)
	if err != nil {
		g.Log().Error(ctx, err)
		res.Message = err.Error()
		return
	}
	err = service.SysExchange().SendExchangeReqToKafka(ctx, data)
	if err != nil {
		g.Log().Error(ctx, err)
		res.Message = err.Error()
		return
	}
	res.Status = "success"
	res.TaskID = rand.Int64()
	return
}
