package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DesensitizeResultReq struct {
	g.Meta      `mime:"application/json" path:"/desensitizeResult" tags:"脱敏结果接收" method:"post" summary:"脱敏结果接收"`
	TaskID      int64  `json:"taskID"`
	ProviderNum int64  `json:"providerNum"`
	Status      string `json:"status"`
}

type DesensitizeResultRes struct {
	g.Meta  `mime:"application/json"`
	TaskID  int64  `json:"taskID"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
