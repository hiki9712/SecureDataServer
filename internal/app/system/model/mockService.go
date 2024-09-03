package model

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	_ "github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type Table struct {
	TableName    string                    `json:"table_name"`
	TaskID       int64                     `json:"task_id"`
	TableContent []*entity.ElectricityInfo `json:"table_content"`
}
