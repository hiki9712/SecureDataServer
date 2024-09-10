package model

import "time"

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
