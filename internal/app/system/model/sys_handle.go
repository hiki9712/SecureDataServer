package model

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
