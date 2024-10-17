package compute

import (
	"context"
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
		err := g.Model("compute_reg").Where("service_owner_id = ?", int64(data["owner_id"].(float64))).Scan(&computeData)
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
		serviceID  int64
		insertData model.ComputeReg
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
	insertData.ComputeTaskID = id
	insertData.ServiceID = serviceID
	insertData.ComputeType = int(data["computeType"].(float64))
	insertData.CreateTime = time.Now()
	insertData.UpdateTime = time.Now()
	insertData.QueryStartTime = time.Now()
	//insertData.QueryEndTime =
	insertData.HandleList = gconv.String(handleID)
	g.Log().Info(ctx, "insertData:", insertData)
	_, err = g.Model("compute_reg").Data(insertData).Insert()

	// insert result
	// insertData_result.ComputeResultID = id
	// insertData_result.ServiceID = serviceID
	// insertData_result.ComputeType = int(data["computeType"].(float64))
	// insertData_result.CreateTime = time.Now()
	// insertData_result.UpdateTime = time.Now()
	// insertData_result.QueryStartTime = time.Now()
	// //insertData_result.QueryEndTime =
	// insertData_result.HandleList = gconv.String(handleID)
	// //g.Log().Info(ctx, "insertData_result:", insertData_result)
	// _, err = g.Model("compute_result").Data(insertData_result).Insert()

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
		postData.Criteria.FieldName = data["criteria"].(map[string]interface{})["fieldName"].(string)
		postData.Criteria.FieldValue = data["criteria"].(map[string]interface{})["fieldValue"].(string)
		//postData.Criteria.FieldName = data["criteria"].(map[string]interface{})["fieldName"].([]interface{})
		//postData.Criteria.FieldValue = data["criteria"].(map[string]interface{})["fieldValue"].([]interface{})
		g.Log().Info(ctx, "postData", postData)
		client := g.Client()
		baseCfg := g.Cfg().MustGet(ctx, "baseApi.default").Map()
		response, resErr := client.Post(ctx, baseCfg["address"].(string)+"/compute/assignPowerConsumptionTask", postData)
		if resErr != nil {
			err = resErr
			return
		}
		responseString := response.ReadAllString()
		g.Log().Info(ctx, "response:", responseString)
		identifierData.TaskID = postData.TaskID
		identifierData.Identifier.FieldName = data["identifier"].(map[string]interface{})["fieldName"].(string)
		identifierData.Identifier.FieldValue = data["identifier"].(map[string]interface{})["fieldValue"].(string)
		//identifierData.Identifier.FieldName = data["identifier"].(map[string]interface{})["fieldName"].([]interface{})
		//identifierData.Identifier.FieldValue = data["identifier"].(map[string]interface{})["fieldValue"].([]interface{})
		g.Log().Info(ctx, "identifierData", identifierData)
		response, resErr = client.Post(ctx, baseCfg["address"].(string)+"/search/provideIdentifier", identifierData)
		if resErr != nil {
			err = resErr
			return
		}
		defer func(response *gclient.Response) {
			err := response.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(response)
	}
	return
}
