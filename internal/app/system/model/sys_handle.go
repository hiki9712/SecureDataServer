package model

import (
	"time"
)

type HandleField struct {
	KeyValueName      string            `json:"keyValueName"`
	OriginalTableName string            `json:"originalTableName"`
	SecureTableName   string            `json:"secureTableName"`
	FieldCount        int               `json:"fieldCount"`
	FieldContent      []FieldContent    `json:"fieldContent"`
	LibName           string            `json:"libName"`
	FuncName          string            `json:"funcName"`
	GarbleBitCount    int               `json:"garbleBitCount"`
	GarbleCoverType   int               `json:"garbleCoverType"`
	GarbleSaveField   []GarbleSaveField `json:"garbleSaveField"`
	GarbleAlgorithm   string            `json:"garbleAlgorithm"`
	OriginTableName   string            `json:"originTableName"` //2025.2.24更新
	OriginFieldName   []string          `json:"originFieldName"`
	DesenTableName    string            `json:"desenTableName"`
	DesenFieldName    []string          `json:"desenFieldName"`
	Format            string            `json:"format"`
	Protocol          int               `json:"protocol"`
}

type FieldContent struct {
	FieldName string `json:"fieldName"`
}

type GarbleSaveField struct {
	FieldName string `json:"fieldName"`
	StartByte int    `json:"startByte"`
	CoverByte int    `json:"coverByte"`
}

type AtomHandleReg struct {
	HandleID        int64         `json:"handle_id"`
	HandleName      string        `json:"handle_name"`
	HandleType      string        `json:"handle_type"`
	CreateTime      time.Time     `json:"create_time"`
	ServiceID       int64         `json:"service_id"`
	ServiceName     string        `json:"service_name"`
	ProviderID      int64         `json:"provider_id"`
	KeyValueCount   int           `json:"keyValueCount"`
	KeyValueContent []interface{} `json:"keyValueContent"`
	DelFlag         int           `json:"del_flag"`
	CreateBy        string        `json:"create_by"`
	UpdateBy        string        `json:"update_by"`
	UpdateTime      time.Time     `json:"update_time"`
	Remark          string        `json:"remark"`
}

type AtomHandleRegLog struct {
	HandleLogID     int64         `json:"handle_log_id"`
	HandleID        int64         `json:"handle_id"`
	HandleName      string        `json:"handle_name"`
	HandleType      string        `json:"handle_type"`
	CreateTime      time.Time     `json:"create_time"`
	ServiceID       int64         `json:"service_id"`
	ServiceName     string        `json:"service_name"`
	ProviderID      int64         `json:"provider_id"`
	KeyValueCount   int           `json:"keyValueCount"`
	KeyValueContent []interface{} `json:"keyValueContent"`
	DelFlag         int           `json:"del_flag"`
	CreateBy        string        `json:"create_by"`
	UpdateBy        string        `json:"update_by"`
	UpdateTime      time.Time     `json:"update_time"`
	Remark          string        `json:"remark"`
}

type Negotiation struct {
	NegotiationID    int64     `json:"negotiation_id"`
	ServiceID        int64     `json:"service_id"`
	ServiceName      string    `json:"service_name"`
	ServiceOwnerID   int64     `json:"service_owner_id"`
	ServiceOwnerName string    `json:"service_owner_name"`
	ProviderID       int64     `json:"provider_id"`
	ProviderName     string    `json:"provider_name"`
	ProviderTable    string    `json:"provider_table"`
	ProviderDB       string    `json:"provider_db"`
	SecureTableName  string    `json:"securetable_name"`
	SecureTableField string    `json:"securetable_field"`
	Status           string    `json:"status"`
	Message          string    `json:"message"`
	DelFlag          int       `json:"del_flag"`
	CreateBy         string    `json:"create_by"`
	CreateTime       time.Time `json:"create_time"`
	UpdateBy         string    `json:"update_by"`
	UpdateTime       time.Time `json:"update_time"`
	Remark           string    `json:"remark"`
}

type NegotiationLog struct {
	NegotiationLogID int64     `json:"negotiation_log_id"`
	NegotiationID    int64     `json:"negotiation_id"`
	ServiceID        int64     `json:"service_id"`
	ServiceName      string    `json:"service_name"`
	ServiceOwnerID   int64     `json:"service_owner_id"`
	ServiceOwnerName string    `json:"service_owner_name"`
	ProviderID       int64     `json:"provider_id"`
	ProviderName     string    `json:"provider_name"`
	ProviderTable    string    `json:"provider_table"`
	ProviderDB       string    `json:"provider_db"`
	SecureTableName  string    `json:"securetable_name"`
	SecureTableField string    `json:"securetable_field"`
	Status           string    `json:"status"`
	Message          string    `json:"message"`
	DelFlag          int       `json:"del_flag"`
	CreateBy         string    `json:"create_by"`
	CreateTime       time.Time `json:"create_time"`
	UpdateBy         string    `json:"update_by"`
	UpdateTime       time.Time `json:"update_time"`
	Remark           string    `json:"remark"`
}

type NegotiationList struct {
	NegotiationID int64  `json:"negotiation_id"`
	ServiceID     int64  `json:"service_id"`
	ServiceName   string `json:"service_name"`
	TableName     string `json:"table_name"`
	DBName        string `json:"db_name"`
	Status        string `json:"status"`
}

type HandleList struct {
	HandleID    int64  `json:"handle_id"`
	ServiceID   int64  `json:"service_id"`
	ServiceName string `json:"service_name"`
	ProviderID  int64  `json:"provider_id"`
}
