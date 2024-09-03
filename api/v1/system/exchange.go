package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ExchangeReq struct {
	g.Meta          `mime:"application/json" path:"/exchange" tags:"数据交换" method:"post" summary:""`
	ServiceName     string         `json:"serviceName"`
	ProviderID      int64          `json:"providerNum"`
	ExchangeContent []providerData `json:"exchangeContent"`
}

type providerData struct {
	ProviderID int64 `json:"providerID"`
	HandleID   int64 `json:"handleID"`
}

type ExchangeRes struct {
	g.Meta  `mime:"application/json"`
	Status  string `json:"status"`
	TaskID  int64  `json:"taskID"`
	Message string `json:"message"`
}
