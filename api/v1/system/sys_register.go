package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type RegisterReq struct {
	g.Meta          `path:"/handle/register" tags:"handle注册" method:"post" summary:"请求注册handleID"`
	ServiceName     string               `json:"serviceName"`
	ServiceID       int64                `json:"serviceID"`
	HandleName      string               `json:"handleName"`
	HandleType      string               `json:"handleType"`
	ProviderID      int64                `json:"providerID"`
	DatabaseName    string               `json:"databaseName"`
	KeyValueCount   int                  `json:"keyValueCount"`
	KeyValueContent []*model.HandleField `json:"keyValueContent"`
}

type RegisterNegotiationReq struct {
	g.Meta       `path:"/handle/negotiation" tags:"handle注册" method:"post" summary:"数据协商"`
	ServiceName  string `json:"serviceName"`
	ProviderID   int64  `json:"providerID"`
	DatabaseName string `json:"databaseName"`
	TableName    string `json:"tableName"`
}

type RegisterNegotiationRes struct {
	g.Meta
	Status    string `json:"status"`
	Message   string `json:"message"`
	ServiceID string `json:"serviceID"`
}

type RegisterNegotiationAgreeReq struct {
	g.Meta           `path:"/handle/negotiationAgree" tags:"handle注册" method:"post" summary:""`
	Agree            bool        `json:"agree"`
	ServiceID        string      `json:"serviceID"`
	ServiceName      string      `json:"serviceName"`
	SecureTableField interface{} `json:"secureTableField"`
}

type ShowHandleInfoReq struct {
	g.Meta   `path:"/handle/registerShow" tags:"handle注册" method:"post" summary:"查看注册handle"`
	HandleID int `json:"handleID"`
}

type ShowHandleInfoRes struct {
	g.Meta   `mime:"application/json"`
	HandleID int `json:"handleID"`
}

type BaseAPIReq struct {
	g.Meta          `mime:"application/json"`
	ServiceName     string               `json:"serviceName"`
	HandleName      string               `json:"handleName"`
	HandleType      string               `json:"handleType"`
	Provider        int64                `json:"provider"`
	DatabaseName    string               `json:"databaseName"`
	KeyValueCount   int                  `json:"keyValueCount"`
	KeyValueContent []*model.HandleField `json:"keyValueContent"`
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
