package sysExchange

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func init() {
	service.ExchangeSysExchange(New())
}

type sSysExchange struct {
}

func New() *sSysExchange {
	return &sSysExchange{}
}

func (s *sSysExchange) ResolveReq(ctx context.Context, req *system.ExchangeReq) (data g.Map, err error) {
	g.Log().Debug(ctx, "req", req)
	//将req序列化为JSON
	reqJson, err := json.Marshal(req)
	if err != nil {
		return
	}
	//将JSON解析为map[string]interface{}
	err = json.Unmarshal(reqJson, &data)
	if err != nil {
		return
	}
	return
}

func (s *sSysExchange) StoreExchangeTaskToDB(ctx context.Context, data g.Map) (err error) {
	return
}

func (s *sSysExchange) SendExchangeReqToKafka(ctx context.Context, data g.Map) (err error) {
	return
}
