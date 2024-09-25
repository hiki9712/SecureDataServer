package libUtils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

// 获取数据库字段信息
func GetDBField(ctx context.Context, db string, table string) (databaseField []model.DatabaseField, err error) {
	result, err := g.DB(db).GetAll(ctx, fmt.Sprintf("SHOW COLUMNS FROM %s", table))
	if err != nil {
		g.Log().Error(ctx, err)
		return databaseField, err
	}
	for _, row := range result.List() {
		databaseRow := model.DatabaseField{
			FieldName: gconv.String(row["Field"]),
			FieldType: gconv.String(row["Type"]),
			IsKey:     gconv.String(row["Key"]),
			IsNull:    gconv.String(row["Null"]),
		}
		databaseField = append(databaseField, databaseRow)
	}
	return databaseField, nil
}
