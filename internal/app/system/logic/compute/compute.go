package compute

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
)

func init() {
	service.ComputeRegister(New())
}

type sCompute struct {
}

func New() *sCompute {
	return &sCompute{}
}

func (s *sCompute) ListCompute(ctx context.Context, data g.Map) (computeData []system.ComputeTask, err error) {
	g.Log().Info(ctx, "listData:", data)
	if data["user_type"].(string) == "owner" {
		err := g.Model("compute_reg").Where("service_owner_id = ?", int64(data["owner_id"].(float64))).Order("update_time DESC").Scan(&computeData)
		if err != nil {
			return nil, err
		}
		g.Log().Info(ctx, "ownerData:", computeData)
	}
	return
}

func (s *sCompute) GetResult(ctx context.Context, data g.Map) (ResultData system.Result, err error) {
	g.Log().Info(ctx, "ResultData:", data)
	if data["user_type"].(string) == "owner" {
		err := g.Model("compute_result").Where("compute_task_id = ?", int64(data["taskid"].(float64))).Scan(&ResultData)
		if err != nil {
			return system.Result{}, err
		}
		g.Log().Info(ctx, "ResultData:", ResultData)
	}
	return
}

// 更新result表
func (s *sCompute) UpdateResultToDB(ctx context.Context, result_value string, TaskID int64) (err error) {
	var (
		updateData model.ComputeResult
	)
	updateData.Result = result_value
	// updateData.UpdateTime = time.Now()
	// updateData.QueryEndTime = time.Now()
	g.Log().Info(ctx, "updateData:", updateData)
	_, err = g.Model("compute_result").Where("compute_task_id = ?", TaskID).Data(updateData).Update()
	return
}

func (s *sCompute) StoreComputeTaskToDB(ctx context.Context, data g.Map) (dataAlter g.Map, err error) {
	var (
		serviceID     int64
		insertData    model.ComputeReg
		insertLogData model.ComputeRegLog
		// insertData_result model.ComputeResult
		handleID int64
	)
	serviceID = int64(data["serviceID"].(float64))
	g.Log().Info(ctx, serviceID, data)
	id := libUtils.GenUniqId(ctx)
	handleID = int64(data["handleID"].(float64))
	dataAlter = data
	dataAlter["id"] = id
	dataAlter["handle_id"] = handleID

	// 处理interfacedata
	critiera_json, _ := json.Marshal(data["criteria"])
	identifier_json, _ := json.Marshal(data["identifier"])

	insertData.ComputeTaskID = id
	insertData.ServiceID = serviceID
	insertData.ComputeType = int(data["computeType"].(float64))
	insertData.CreateTime = time.Now()
	insertData.UpdateTime = time.Now()
	insertData.QueryStartTime = time.Now()
	//insertData.QueryEndTime =
	insertData.Criteria = string(critiera_json)
	insertData.Identifier = string(identifier_json)
	insertData.ServiceOwnerID = 12
	insertData.ProviderIDList = "1234"
	insertData.Status = "commited"
	insertData.HandleList = gconv.String(handleID)

	g.Log().Info(ctx, "insertData:", insertData)
	_, err = g.Model("compute_reg").Data(insertData).Insert()

	if err != nil {
		g.Log().Error(ctx, err)
	}

	// insert log
	insertLogData.ComputeTaskLogID = 0 // 自增id
	insertLogData.ComputeTaskID = id
	insertLogData.ServiceID = serviceID
	insertLogData.ComputeType = int(data["computeType"].(float64))
	insertLogData.Criteria = string(critiera_json)
	insertLogData.Identifier = string(identifier_json)
	insertLogData.ServiceOwnerID = 12
	insertLogData.ProviderIDList = "1234"
	insertLogData.Status = "committed"
	insertLogData.HandleList = gconv.String(handleID)
	insertLogData.Time = time.Now()

	g.Log().Info(ctx, "insertLogData:", insertLogData)
	_, err = g.Model("compute_reg_log").Data(insertLogData).Insert()
	if err != nil {
		g.Log().Error(ctx, err)
	}

	return
}

func (s *sCompute) SendReqByComputeType(ctx context.Context, data g.Map) (res interface{}, err error) {
	if int(data["computeType"].(float64)) == 1 {
		var (
			postData       system.AssignPowerConsumptionTask
			identifierData system.ProviderIdentifier
		)
		postData.TaskID = data["id"].(int64)
		postData.HandleID = data["handle_id"].(int64)
		postData.ComputeType = int(data["computeType"].(float64))
		//postData.Criteria.FieldName = "STARTTIME,ENDTIME"
		postData.Criteria.FieldName = data["criteria"].(map[string]interface{})["fieldName"].(string)
		//postData.Criteria.FieldValue = "2024-07-01,2024-09-01"
		postData.Criteria.FieldValue = data["criteria"].(map[string]interface{})["fieldValue"].(string)
		g.Log().Info(ctx, "postData", postData)
		client := g.Client()
		baseCfg := g.Cfg().MustGet(ctx, "baseApi.default").Map()
		tempData := gconv.String(postData)
		response, resErr := client.Post(ctx, baseCfg["address"].(string)+"/compute/assignPowerConsumptionTask", tempData)
		if resErr != nil {
			err = resErr
			return
		}
		// responseString := response.ReadAllString()
		// g.Log().Info(ctx, "response:", responseString)
		identifierData.TaskID = postData.TaskID
		//identifierData.Identifier.FieldName = "CUSTOMERID"
		identifierData.Identifier.FieldName = data["identifier"].(map[string]interface{})["fieldName"].(string)
		//identifierData.Identifier.FieldValue = [][]string{{"17"}}
		identifierData.Identifier.FieldValue = data["identifier"].(map[string]interface{})["fieldValue"].([]interface{})
		g.Log().Info(ctx, "identifierData", identifierData)
		tempData = gconv.String(identifierData)
		response, resErr = client.Post(ctx, baseCfg["address"].(string)+"/search/provideIdentifier", tempData)
		if resErr != nil {
			err = resErr
			return
		}
		responseString := response.ReadAllString()
		g.Log().Info(ctx, "response:", responseString)
		defer func(response *gclient.Response) {
			err := response.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(response)
	}

	// 更新compute_reg表
	var updateData model.ComputeReg
	updateData.ComputeTaskID = data["id"].(int64)
	updateData.Status = "waiting"
	updateData.UpdateTime = time.Now()
	_, err = g.Model("compute_reg").Where("compute_task_id = ?", updateData.ComputeTaskID).Data(updateData).OmitEmpty().Update()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	// 写入compute_reg_log表
	// 处理interfacedata
	critiera_json, _ := json.Marshal(data["criteria"])
	identifier_json, _ := json.Marshal(data["identifier"])
	var insertLogData model.ComputeRegLog
	insertLogData.ComputeTaskID = data["id"].(int64)
	insertLogData.ServiceID = int64(data["serviceID"].(float64))
	insertLogData.ComputeType = int(data["computeType"].(float64))
	insertLogData.Criteria = string(critiera_json)
	insertLogData.Identifier = string(identifier_json)
	insertLogData.ServiceOwnerID = 12
	insertLogData.ProviderIDList = "1234"
	insertLogData.Status = "waiting"
	insertLogData.HandleList = gconv.String(data["handle_id"].(int64))
	insertLogData.Time = time.Now()
	_, err = g.Model("compute_reg_log").Data(insertLogData).Insert()
	if err != nil {
		g.Log().Error(ctx, err)
	}

	return
}

// 任务结束，更新compute_reg表，写入日志
func (s *sCompute) UpdateComputeRegToDB(ctx context.Context, data g.Map) (err error) {
	var (
		updateData    model.ComputeReg
		insertLogData model.ComputeRegLog
	)
	updateData.ComputeTaskID = data["id"].(int64)
	updateData.Status = "finished"
	updateData.UpdateTime = time.Now()
	_, err = g.Model("compute_reg").Where("compute_task_id = ?", updateData.ComputeTaskID).Data(updateData).OmitEmpty().Update()
	if err != nil {
		g.Log().Error(ctx, err)
	}

	// 处理interfacedata
	critiera_json, _ := json.Marshal(data["criteria"])
	identifier_json, _ := json.Marshal(data["identifier"])

	// 写入compute_reg_log表
	insertLogData.ComputeTaskLogID = 0 // 自增id
	insertLogData.ComputeTaskID = data["id"].(int64)
	insertLogData.ServiceID = data["serviceID"].(int64)
	insertLogData.ComputeType = int(data["computeType"].(float64))
	insertLogData.Criteria = string(critiera_json)
	insertLogData.Identifier = string(identifier_json)
	insertLogData.ServiceOwnerID = 12
	insertLogData.ProviderIDList = "1234"
	insertLogData.Status = "finished"
	insertLogData.HandleList = gconv.String(data["handle_id"].(int64))
	insertLogData.Time = time.Now()
	_, err = g.Model("compute_reg_log").Data(insertLogData).Insert()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}
