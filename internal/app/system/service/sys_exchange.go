package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type (
	ISysExchange interface {
		ResolveReq(ctx context.Context, req interface{}) (data g.Map, err error)
		StoreExchangeTaskToDB(ctx context.Context, data g.Map) (message string, err error)
		SendExchangeReqToKafka(ctx context.Context, data g.Map) error
		Listhandle(ctx context.Context, data g.Map) (handleData []model.HandleList, err error)
		FetchTable(ctx context.Context, data g.Map) (tableData gdb.Result, handleID int64, err error)
		SendToMasking(ctx context.Context, data g.Map) error
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
