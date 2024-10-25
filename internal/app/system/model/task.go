/*
* @desc:定时任务
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/1/13 17:47
 */

package model

import "context"

type TimeTask struct {
	FuncName string
	Param    []string
	Run      func(ctx context.Context)
}

type NegotiationDetail struct {
	ServiceID       int64  `json:"service_id"`
	ServiceOwnerID  int64  `json:"service_owner_id"`
	ServiceName     string `json:"service_name"`
	ProviderID      int64  `json:"provider_id"`
	ProviderTable   string `json:"provider_table"`
	ProviderDB      string `json:"provider_db"`
	Status          string `json:"status"`
	DelFlag         int    `json:"del_flag"`
	SecureTableName string `json:"securetable_name"`
}

type TaskData struct {
	TaskID          int64  `json:"task_id"`
	ServiceID       int64  `json:"service_id"`
	ServiceName     string `json:"service_name"`
	ServiceOwnerID  int64  `json:"service_owner_id"`
	ProviderID      int64  `json:"provider_id"`
	HandleID        int64  `json:"handle_id"`
	HandleName      string `json:"handle_name"`
	DBName          string `json:"db_name"`
	TableName       string `json:"table_name"`
	Status          string `json:"status"`
	SecureTableName string `json:"securetable_name"`
}

type TaskDataLog struct {
	TaskLogID       int64  `json:"task_log_id"`
	TaskID          int64  `json:"task_id"`
	ServiceID       int64  `json:"service_id"`
	ServiceName     string `json:"service_name"`
	ServiceOwnerID  int64  `json:"service_owner_id"`
	ProviderID      int64  `json:"provider_id"`
	HandleID        int64  `json:"handle_id"`
	HandleName      string `json:"handle_name"`
	DBName          string `json:"db_name"`
	TableName       string `json:"table_name"`
	Status          string `json:"status"`
	SecureTableName string `json:"securetable_name"`
}

type ProvideRawDataReq struct {
	TaskID      int32    `json:"taskID"`
	HandleID    int64    `json:"handleID"`
	DataAddress []string `json:"dataAddress"`
	HashCode    []string `json:"hashCode"`
}

type TaskTableDetail struct {
	SecureTableName string        `json:"secureTableName"`
	TableData       []interface{} `json:"tableData"`
}
