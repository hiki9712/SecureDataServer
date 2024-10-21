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

type RegisterListReq struct {
}

type RegisterNegotiationReq struct {
	g.Meta         `path:"/handle/negotiation" tags:"handle注册" method:"post" summary:"数据协商"`
	ServiceName    string `json:"serviceName"`
	ProviderID     int64  `json:"providerID"`
	ServiceOwnerID int64  `json:"serviceOwnerID"`
	FieldContent   interface{}          `json:"fieldContent"`
}

type RegisterNegotiationToProReq struct {
	g.Meta         `path:"/handle/negotiationToPro" tags:"handle注册" method:"post" summary:"数据协商传给提供方"`
	ServiceName    string `json:"serviceName"`
	ProviderID     int64  `json:"providerID"`
	ServiceID     int64  `json:"serviceID"`
	ServiceOwnerID int64  `json:"serviceOwnerID"`
	FieldContent   interface{}          `json:"fieldContent"`
}

type RegisterNegotiationRes struct {
	g.Meta
	Status    string `json:"status"`
	Message   string `json:"message"`
	ServiceID int64  `json:"serviceID"`
}

type RegisterNegotiationAgreeReq struct {
	g.Meta           `path:"/handle/negotiationAgree" tags:"handle注册" method:"post" summary:"数据提供方审核"`
	Agree            bool        `json:"agree"`
	NegotiationID    int64       `json:"negotiationID"`
	ServiceID        int64       `json:"serviceID"`
	ServiceName      string      `json:"serviceName"`
	SecureTableName  string      `json:"secureTableName"`
	SecureTableField interface{} `json:"secureTableField"`
	Message          string      `json:"message"`
}

type RegisterNegotiationAgreeToReqReq struct {
	g.Meta           `path:"/handle/negotiationAgreeToReq" tags:"handle注册" method:"post" summary:"协商信息发送给需求方"`
	Agree            bool        `json:"agree"`
	NegotiationID    int64       `json:"negotiationID"`
	ServiceID        int64       `json:"serviceID"`
	ServiceName      string      `json:"serviceName"`
	SecureTableName  string      `json:"secureTableName"`
	SecureTableField interface{} `json:"secureTableField"`
	Message          string      `json:"message"`
}

type RegisterNegotiationAgreeRes struct {
	g.Meta
	Status    string `json:"status"`
	Message   string `json:"message"`
	ServiceID int64  `json:"serviceID"`
}

type NegotiationNotifyReq struct {
	g.Meta    `path:"/handle/notify" tags:"handle注册" method:"post" summary:"建表通知"`
	ServiceID int64 `json:"serviceID"`
}

type NegotiationNotifyRes struct {
	g.Meta
	ServiceID int64  `json:"serviceID"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

type NegotiationListReq struct {
	g.Meta     `path:"/handle/negotiationList" tags:"handle注册" method:"post" summary:"查看协商"`
	UserType   string `json:"user_type"`
	ProviderID int64  `json:"provider_id"`
	OwnerID    int64  `json:"owner_id"`
}

type NegotiationListRes struct {
	g.Meta
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Items   []model.NegotiationList `json:"items"`
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
