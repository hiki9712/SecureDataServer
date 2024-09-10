package system

import "github.com/gogf/gf/v2/frame/g"

type MockRetryReq struct {
	g.Meta  `mime:"application/json" path:"/supplyDataExchangeResults" tags:"数据重试" method:"post" summary:"重试kafka"`
	TaskID  int64  `json:"taskID"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type MockRetryRes struct {
	g.Meta   `mime:"application/json"`
	Status   string `json:"status"`
	TaskID   int64  `json:"taskID"`
	Message  string `json:"message"`
	TryAgain bool   `json:"tryAgain"`
}
