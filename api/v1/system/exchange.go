package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ExchangeReq struct {
	g.Meta    `mime:"application/json" path:"/exchange/sendExchangeRequest" tags:"数据交换" method:"post" summary:"前端向数据请求方发送数据交换请求"`
	ServiceID int64 `json:"serviceID"`
	HandleID  int64 `json:"handleID"`
}

type ExchangeRes struct {
	g.Meta  `mime:"application/json"`
	Status  string `json:"status"`
	TaskID  int64  `json:"taskID"`
	Message string `json:"message"`
}

type SendDataReq struct {
	g.Meta     `mime:"application/json" path:"/exchange/sendData" tags:"数据交换" method:"post" summary:"数据提供方发送业务数据"`
	TaskID     int64 `json:"taskID"`
	ProviderID int64 `json:"providerID"`
}

type SendDataRes struct {
	g.Meta  `mime:"application/json"`
	Status  string `json:"status"`
	TaskID  int64  `json:"taskID"`
	Message string `json:"message"`
}
