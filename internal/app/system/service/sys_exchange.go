package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type (
	ISysExchange interface {
		ResolveReq(ctx context.Context, req interface{}) (data g.Map, err error)
		StoreExchangeTaskToDB(ctx context.Context, data g.Map) (message string, err error)
		SendExchangeReqToKafka(ctx context.Context, data g.Map) error
		FetchTable(ctx context.Context, data g.Map) (tableData *model.Table, err error)
		SendToMasking(ctx context.Context, tableData *model.Table) error
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
