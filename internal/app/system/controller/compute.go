package controller

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/internal/websocket"
	"github.com/tiger1103/gfast/v3/library/libUtils"
)

var (
	Compute = computeController{}
)

type computeController struct {
	BaseController
}

func (c *computeController) TaskList(ctx context.Context, req *system.ComputeTaskListReq) (res *system.ComputeTaskListRes, err error) {
	var (
		data g.Map
	)
	res = &system.ComputeTaskListRes{
		Status:  "fail",
		Message: "",
	}
	data, err = libUtils.ResolveReq(ctx, req)
	if err != nil {
		return
	}
	dataList, err := service.Compute().ListCompute(ctx, data)
	if err != nil {
		return
	}
	res.Status = "success"
	res.Data = dataList
	return
}

func (c *computeController) SendRequest(ctx context.Context, req *system.ComputeSendReq) (res *system.ComputeSendRes, err error) {
	var (
		data g.Map
	)
	res = &system.ComputeSendRes{
		Status:  "fail",
		Message: "",
	}
	data, err = libUtils.ResolveReq(ctx, req)
	if err != nil {
		return
	}
	data, err = service.Compute().StoreComputeTaskToDB(ctx, data)
	if err != nil {
		return
	}
	g.Log().Info(ctx, "compute info:", data)
	_, err = service.Compute().SendReqByComputeType(ctx, data) // 开始查询
	if err != nil {
		return
	}

	res.Status = "success"
	res.TaskID = gconv.Int64(data["id"])
	return
}

func (c *computeController) GetResult(ctx context.Context, req *system.ComputeResultReq) (res *system.ComputeResultRes, err error) {
	var (
		data g.Map
	)
	res = &system.ComputeResultRes{
		Status:  "fail",
		Message: "",
	}
	data, err = libUtils.ResolveReq(ctx, req)
	if err != nil {
		return
	}

	// 使用 gconv.Int64 进行类型转换
	taskIDValue, taskIDExists := data["taskID"]
	if !taskIDExists {
		err = errors.New("TaskID does not exist")
		return
	}

	taskID := gconv.Int64(taskIDValue)
	if taskID == 0 {
		err = errors.New("TaskID is not of type int64")
		return
	}
	// 打印转换后的 TaskID 值
	g.Log().Info(ctx, "Converted TaskID value:", taskID)

	result_res := &system.ComputeResultReq{
		Status:  "fail",
		Message: "",
	}

	jsonStr := gconv.String(data)
	err = json.Unmarshal([]byte(jsonStr), result_res)
	if err != nil {
		return
	}

	if len(result_res.ComputeResult) > 0 {

		// 如果有多个计算结果，拼接成字符串发送给前端
		resultValue := ""
		for _, v := range result_res.ComputeResult {

			identifier_name := v.Identifier.FieldName
			identifier_value := v.Identifier.FieldValue[0].(string)
			identifier := identifier_name + ":" + identifier_value + "\n"
			result := v.Result

			resultValue += identifier + " --> " + result + "\n"
			//resultValue += v.Result + "\n"
			// g.Log().Info(ctx, "compute result:", v)
		}
		g.Log().Info(ctx, "compute result:", resultValue)
		// resultValue := result_res.ComputeResult[0].Result
		// g.Log().Info(ctx, "compute result:", resultValue)

		// 通过 WebSocket 发送消息
		websocket.SendWebSocketMessage(ctx, taskID, resultValue)
	}

	err = service.Compute().UpdateComputeRegToDB(ctx, taskID)
	if err != nil {
		g.Log().Warning(ctx, "compute Reg update failed:", err)
	}

	res.Status = "success"
	return
}
