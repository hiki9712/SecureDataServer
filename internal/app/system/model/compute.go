package model

import (
	"time"
)

type ComputeReg struct {
	ComputeTaskID  int64     `json:"compute_task_id"    comment:"计算任务的唯一ID"`
	ComputeType    int       `json:"compute_type"       comment:"计算类型"`
	ServiceID      int64     `json:"service_id"         comment:"服务的唯一ID"`
	ServiceOwnerID int64     `json:"service_owner_id"   comment:"服务注册者ID"`
	ServiceOwner   string    `json:"service_owner"      comment:"服务注册者"`
	ServiceName    string    `json:"service_name"       comment:"注册的服务名"`
	QueryName      string    `json:"query_name"         comment:"查询的联系人姓名"`
	QueryPhone     string    `json:"query_phone"        comment:"查询的手机号"`
	QueryStartTime time.Time `json:"query_start_time"   comment:"查询的开始时间"`
	QueryEndTime   time.Time `json:"query_end_time"     comment:"查询的结束时间"`
	ProviderIDList string    `json:"provider_id_list"   comment:"数据提供者的ID列表"`
	HandleList     string    `json:"handle_list"        comment:"该服务已经注册完成的handle列表"`
	DelFlag        int       `json:"del_flag"           comment:"删除标志（0存在1删除）"`
	CreateBy       string    `json:"create_by"          comment:"创建者"`
	CreateTime     time.Time `json:"create_time"        comment:"创建时间"`
	UpdateBy       string    `json:"update_by"          comment:"更新者"`
	UpdateTime     time.Time `json:"update_time"        comment:"更新时间"`
	Remark         string    `json:"remark"             comment:"备注"`
	QueryHH        string    `json:"query_hh"           comment:"查询的户号"`
}

type ComputeResult struct {
	ComputeTaskID int64  `json:"compute_task_id"    comment:"计算任务索引ID"`
	Result        string `json:"result"             comment:"计算结果"`
}

// type ComputeResult struct {
// 	ComputeResultID int64     `json:"compute_result_id"  comment:"计算结果的唯一ID"`
// 	ComputeType     int       `json:"compute_type"       comment:"计算类型"`
// 	ServiceID       int64     `json:"service_id"         comment:"服务的唯一ID"`
// 	ServiceOwnerID  int64     `json:"service_owner_id"   comment:"服务注册者ID"`
// 	ServiceOwner    string    `json:"service_owner"      comment:"服务注册者"`
// 	ServiceName     string    `json:"service_name"       comment:"注册的服务名"`
// 	QueryName       string    `json:"query_name"         comment:"查询的联系人姓名"`
// 	QueryPhone      string    `json:"query_phone"        comment:"查询的手机号"`
// 	QueryStartTime  time.Time `json:"query_start_time"   comment:"查询的开始时间"`
// 	QueryEndTime    time.Time `json:"query_end_time"     comment:"查询的结束时间"`
// 	ProviderIDList  string    `json:"provider_id_list"   comment:"数据提供者的ID列表"`
// 	HandleList      string    `json:"handle_list"        comment:"该服务已经注册完成的handle列表"`
// 	DelFlag         int       `json:"del_flag"           comment:"删除标志（0存在1删除）"`
// 	CreateBy        string    `json:"create_by"          comment:"创建者"`
// 	CreateTime      time.Time `json:"create_time"        comment:"创建时间"`
// 	UpdateBy        string    `json:"update_by"          comment:"更新者"`
// 	UpdateTime      time.Time `json:"update_time"        comment:"更新时间"`
// 	Remark          string    `json:"remark"             comment:"备注"`
// 	QueryHH         string    `json:"query_hh"           comment:"查询的户号"`
// 	Result          string    `json:"result"             comment:"计算结果"`
// }
