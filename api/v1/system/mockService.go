package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MockServiceReq struct {
	g.Meta `mime:"application/json" path:"/mockService" tags:"模拟业务系统kafka" method:"post" summary:"发送未脱敏数据"`
	TaskID int64  `json:"taskID"`
	DB     string `json:"db"`
	Table  string `json:"table"`
}

type MockServiceRes struct {
	g.Meta `mime:"application/json"`
	Status string `json:"status"`
	TaskID int64  `json:"taskID"`
}
