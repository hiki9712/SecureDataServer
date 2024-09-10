package controller

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
)

var (
	DesensitizeResult = desensitizeResultController{}
)

type desensitizeResultController struct {
	BaseController
}

func (c *desensitizeResultController) DesensitizeResult(ctx context.Context, req *system.DesensitizeResultReq) (res *system.DesensitizeResultRes, err error) {
	var (
		data   g.Map
		TaskID int64
	)
	reqJson, _ := json.Marshal(req)
	err = json.Unmarshal(reqJson, &data)
	g.Log().Info(ctx, "data:", data)
	//TaskID = data["task_id"]
	//_, err = g.DB("default").Model("task").Ctx(ctx).Data("status", "success").Where("task_id=?", TaskID).Update()
	res = &system.DesensitizeResultRes{
		Status:  "success",
		TaskID:  TaskID,
		Message: "",
	}
	return
}
