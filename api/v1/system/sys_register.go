package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type RegisterReq struct {
	g.Meta          `mime:"application/json" path:"/register" tags:"handle注册" method:"post" summary:"请求注册handleID"`
	ServiceName     string               `json:"serviceName"`
	HandleName      string               `json:"handleName"`
	HandleType      string               `json:"handleType"`
	Provider        int64                `json:"provider"`
	DatabaseName    string               `json:"databaseName"`
	KeyValueCount   int                  `json:"keyValueCount"`
	KeyValueContent []*model.HandleField `json:"keyValueContent"`
}

type BaseAPIReq struct {
	g.Meta                `mime:"application/json"`
	Type                  string               `json:"type"`
	FieldNum              int                  `json:"fieldNum"`
	AtomicHandleNum       int                  `json:"atomicHandleNum"`
	HandleID              int64                `json:"handleID"`
	AtomicHandleContent   []*model.HandleField `json:"atomicHandleContent"`
	CombinedHandleContent []AtomRegisterReq    `json:"combinedHandleContent"`
}

type AtomRegisterReq struct {
	g.Meta              `path:"/register" tags:"注册" method:"post" summary:"handle注册"`
	HandleName          string               `json:"handleName"`
	Type                string               `json:"type"`
	FieldNum            int                  `json:"fieldNum"`
	AtomicHandleContent []*model.HandleField `json:"atomicHandleContent"`
}

type CombinedRegisterReq struct {
	g.Meta        `mime:"application/json"`
	HandleName    string
	HandleType    string
	AtomHandleNum int
	HandleContent []AtomRegisterReq
}

type RegisterRes struct {
	g.Meta     `mime:"application/json"`
	HandleName string `json:"handleName"`
	Status     string `json:"status"`
}

type BaseAPIRes struct {
	g.Meta   `mime:"application/json"`
	Status   string `json:"status"`
	HandleID int64  `json:"handleID"`
	Message  string `json:"message"`
}
