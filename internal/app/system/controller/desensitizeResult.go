package controller

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"time"
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
	TaskID = int64(data["taskID"].(float64))
	if data["status"].(string) == "success" {
		_, err = g.Model("task").Ctx(ctx).Data("status", consts.TaskSuccess, "update_time", time.Now()).Where("task_id=?", TaskID).Update()
	} else {
		_, err = g.Model("task").Ctx(ctx).Data("status", consts.TaskFailed, "update_time", time.Now()).Where("task_id=?", TaskID).Update()
	}
	res = &system.DesensitizeResultRes{
		Status:  "success",
		TaskID:  TaskID,
		Message: "",
	}
	return
}
