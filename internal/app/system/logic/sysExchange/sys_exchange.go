package sysExchange

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func init() {
	service.ExchangeSysExchange(New())
}

type sSysExchange struct {
}

func New() *sSysExchange {
	return &sSysExchange{}
}

func (s *sSysExchange) ResolveReq(ctx context.Context, req interface{}) (data g.Map, err error) {
	//将req序列化为JSON
	reqJson, err := json.Marshal(req)
	if err != nil {
		return
	}
	//将JSON解析为map[string]interface{}
	err = json.Unmarshal(reqJson, &data)
	if err != nil {
		return
	}
	return
}

func (s *sSysExchange) StoreExchangeTaskToDB(ctx context.Context, data g.Map) (message string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			negotiationDetail model.NegotiationDetail
			node              *snowflake.Node
			taskId            int64
			insertData        model.TaskData
		)

		serviceID := int64(data["serviceID"].(float64))
		err = g.Model("negotiation").Where("service_id = ?", serviceID).Scan(&negotiationDetail)
		g.Log().Info(ctx, "negotiationDetail:", negotiationDetail)
		if negotiationDetail.Status != "success" || negotiationDetail.DelFlag != 0 {
			g.Log().Info(ctx, "service not ready or deleted, service id:", serviceID)
			message = "service not ready or deleted"
			return
		}
		node, err = snowflake.NewNode(1)
		taskId = node.Generate().Int64()
		insertData.TaskID = taskId
		insertData.ServiceID = serviceID
		insertData.Status = "pending"
		insertData.ServiceName = negotiationDetail.ServiceName
		insertData.ServiceOwnerID = negotiationDetail.ServiceOwnerID
		insertData.ProviderID = negotiationDetail.ProviderID
		insertData.DBName = negotiationDetail.ProviderDB
		insertData.TableName = negotiationDetail.ProviderTable
		//TODO insertData.HandleID = negotiationDetail.HandleID
		_, err = g.Model("task").Data(insertData).Insert()
	})
	return
}

func (s *sSysExchange) SendExchangeReqToKafka(ctx context.Context, data g.Map) (err error) {
	return
}

func (s *sSysExchange) FetchTable(ctx context.Context, data g.Map) (tableData *model.Table, err error) {
	var (
		taskID     int64
		providerID int64
		taskData   model.TaskData
	)
	tableData = &model.Table{}
	taskID = int64(data["taskID"].(float64))
	providerID = int64(data["providerID"].(float64))
	err = g.Model("task").Where("task_id = ?", taskID).Scan(&taskData)
	g.Log().Info(ctx, "result:", taskData)
	if providerID != taskData.ProviderID {
		err = errors.New("provider not correct")
		return
	}
	result, err := g.DB(taskData.DBName).Model(taskData.TableName).Ctx(ctx).All()
	g.Log().Info(ctx, "infos:", result)
	return
}

func (s *sSysExchange) SendToMasking(ctx context.Context, tableData *model.Table) (err error) {
	return
}
