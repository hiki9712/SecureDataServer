package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type ExchangeReq struct {
	g.Meta     `mime:"application/json" path:"/exchange/sendExchangeRequest" tags:"数据交换" method:"post" summary:"前端向数据请求方发送数据交换请求"`
	ServiceID  int64  `json:"serviceID"`
	HandleID   int64  `json:"handleID"`
	HandleName string `json:"handleName"`
	Format     string `json:"format"`
	Protocol   int    `json:"protocol"`
}

type ExchangeRes struct {
	g.Meta  `mime:"application/json"`
	Status  string `json:"status"`
	TaskID  int64  `json:"taskID"`
	Message string `json:"message"`
}

type ListhandleReq struct {
	g.Meta   `mime:"application/json" path:"/exchange/listHandle" tags:"数据交换" method:"post" summary:"获取数据请求方的处理列表"`
	UserType string `json:"user_type"`
	OwnerID  int64  `json:"owner_id"`
}

type SendDataReq struct {
	g.Meta    `mime:"application/json" path:"/exchange/sendData" tags:"数据交换" method:"post" summary:"数据提供方发送业务数据"`
	TableList []model.TaskData `json:"tableList"`
}

type SendDataRes struct {
	g.Meta  `mime:"application/json"`
	Status  string `json:"status"`
	TaskID  int64  `json:"taskID"`
	Message string `json:"message"`
}

type HandleListRes struct {
	g.Meta
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Items   []model.HandleList `json:"items"`
}

type tableDetail struct {
	DB              string `json:"db"`
	Table           string `json:"table"`
	SecureTableName string `json:"secureTableName"`
}
