package model

type DatabaseField struct {
	FieldName    string `json:"fieldName"`
	FieldNameNew string `json:"fieldNameNew"`
	FieldType    string `json:"fieldType"`
	IsKey        string `json:"isKey"`
	IsSecret     string `json:"isSecret"`
	IsNull       string `json:"isNull"`
}
