package system

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type ComputeSendReq struct {
	g.Meta `path:"/compute/sendRequest" method:"post" tags:"计算" summary:"发送数据"`
	//TODO 前端请求业务系统进行数据查询
	ServiceID   int64          `json:"serviceID"`
	ComputeType int            `json:"computeType"`
	HandleID    int64          `json:"handleID"`
	Criteria    CriteriaType   `json:"criteria"`
	Identifier  IdentifierType `json:"identifier"`
}

type AssignPowerConsumptionTask struct {
	TaskID      int64        `json:"task_id"`
	ComputeType int          `json:"compute_type"`
	HandleID    int64        `json:"handle_id"`
	Criteria    CriteriaType `json:"criteria"`
}

type ProviderIdentifier struct {
	TaskID     int64          `json:"taskID"`
	Identifier IdentifierType `json:"identifier"`
}

type IdentifierType struct {
	//FieldName  []interface{} `json:"fieldName"`
	//FieldValue []interface{} `json:"fieldValue"`
	FieldName  string        `json:"fieldName"`
	FieldValue []interface{} `json:"fieldValue"`
}

type CriteriaType struct {
	//FieldName  []interface{} `json:"fieldName"`
	//FieldValue []interface{} `json:"fieldValue"`
	FieldName  string `json:"fieldName"`
	FieldValue string `json:"fieldValue"`
}

type ComputeSendRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	TaskID  int64  `json:"taskID"`
}

type ComputeResultReq struct {
	g.Meta        `path:"/compute/getResult" method:"post" tags:"计算" summary:"接收数据"`
	TaskID        int64           `json:"taskID"`
	Status        string          `json:"status"`
	Message       string          `json:"message"`
	ComputeResult []computeResult `json:"computeResult"`
}

type computeResult struct {
	Identifier IdentifierType `json:"identifier"`
	Result     string         `json:"result"`
}

type ComputeResultRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	TaskID  int64  `json:"taskID"`
}

type ComputeTaskListReq struct {
	g.Meta     `path:"/compute/TaskList" method:"post" tags:"计算" summary:"展示数据"`
	UserType   string `json:"user_type"` //TODO 该字段含义
	OwnerID    int64  `json:"owner_id"`
	ProviderID int64  `json:"provider_id"` //TODO 该字段含义
}

type ComputeTaskListRes struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []ComputeTask `json:"data"`
}

type ComputeTask struct {
	ComputeTaskID  int64     `json:"computeTaskID"`
	ComputeType    int       `json:"computeType"`
	ServiceID      int64     `json:"serviceID"`
	HandleIDList   string    `json:"HandleIDList"`
	QueryStartTime time.Time `json:"queryStartTime"`
	QueryEndTime   time.Time `json:"queryEndTime"`
	ProviderID     int64     `json:"providerID"`
}

type ResultReq struct {
	g.Meta `path:"/compute/ShowResult" method:"get" tags:"计算" summary:"展示计算结果数据"`
	TaskID int64 `json:"owner_id"`
}

type Result struct {
	Result string `json:"result"`
}
