package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
)

type (
	ISysExchange interface {
		ResolveReq(ctx context.Context, req *system.ExchangeReq) (data g.Map, err error)
		StoreExchangeTaskToDB(ctx context.Context, data g.Map) error
		SendExchangeReqToKafka(ctx context.Context, data g.Map) error
	}
)

var (
	localSysExchange ISysExchange
)

func SysExchange() ISysExchange {
	if localSysExchange == nil {
		panic("implement not found for interface ISysExchange")
	}
	return localSysExchange
}

func ExchangeSysExchange(i ISysExchange) {
	localSysExchange = i
}
