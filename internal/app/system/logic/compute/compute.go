package compute

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"time"
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

func (s *sCompute) StoreComputeTaskToDB(ctx context.Context, data g.Map) (dataAlter g.Map, err error) {
	var (
		serviceID  int64
		insertData model.ComputeReg
		handleID   int64
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
	insertData.ServiceOwnerID = 12
	insertData.ProviderIDList = "1234"
	//insertData.QueryEndTime =
	insertData.HandleList = gconv.String(handleID)
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
		responseString := response.ReadAllString()
		g.Log().Info(ctx, "response:", responseString)
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
		responseString = response.ReadAllString()
		g.Log().Info(ctx, "response:", responseString)
		defer func(response *gclient.Response) {
			err := response.Close()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(response)
	}
	return
}
