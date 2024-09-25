package compute

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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

func (s *sCompute) StoreComputeTaskToDB(ctx context.Context, data g.Map) (dataAlter g.Map, err error) {
	var (
		serviceID  int64
		insertData model.ComputeReg
		handleID   int64
	)
	serviceID = int64(data["serviceID"].(float64))
	g.Log().Info(ctx, serviceID, data)
	id := libUtils.GenUniqId(ctx)
	result, err := g.Model("handle_reg").Fields("handle_id").Where("service_id = ?", serviceID).One()
	handleID = result.Map()["handle_id"].(int64)
	dataAlter = data
	dataAlter["id"] = id
	dataAlter["handle_id"] = handleID
	insertData.ComputeTaskID = id
	insertData.ServiceID = serviceID
	insertData.ComputeType = int(data["computeType"].(float64))
	g.Log().Info(ctx, "insertData:", insertData)
	_, err = g.Model("compute_reg").Data(insertData).Insert()
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
		postData.Criteria.FieldName = data["criteria"].(map[string]interface{})["fieldName"].([]interface{})
		postData.Criteria.FieldValue = data["criteria"].(map[string]interface{})["fieldValue"].([]interface{})
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
		identifierData.Identifier.FieldName = data["identifier"].(map[string]interface{})["fieldName"].([]interface{})
		identifierData.Identifier.FieldValue = data["identifier"].(map[string]interface{})["fieldValue"].([]interface{})
		g.Log().Info(ctx, "identifierData", identifierData)
		response, resErr = client.Post(ctx, baseCfg["address"].(string)+"/search/provideIdentifier", identifierData)
		if resErr != nil {
			err = resErr
			return
		}
		defer response.Close()
	}
	return
}
