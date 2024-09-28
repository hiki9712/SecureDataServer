package sysExchange

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"strconv"
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
			postData          system.SendDataReq
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
		insertData.SecureTableName = negotiationDetail.SecureTableName
		insertData.HandleID = int64(data["handleID"].(float64))
		//TODO insertData.HandleID = negotiationDetail.HandleID
		_, err = g.Model("task").Data(insertData).Insert()
		client := g.Client()
		postData.TaskID = taskId
		postData.ProviderID = negotiationDetail.ProviderID
		g.Log().Info(ctx, "postData:", postData)
		providerCfg := g.Cfg().MustGet(ctx, "providerAddress."+strconv.FormatInt(negotiationDetail.ProviderID, 10)).Map()
		response, resErr := client.Post(ctx, providerCfg["address"].(string)+"/api/v1/system/exchange/sendData", postData)
		if resErr != nil {
			err = resErr
			g.Log().Info(ctx, "resErr:", resErr)
			return
		}
		responseString := response.ReadAllString()
		g.Log().Info(ctx, "response:", responseString)
	})
	return
}

func (s *sSysExchange) SendExchangeReqToKafka(ctx context.Context, data g.Map) (err error) {
	return
}

func (s *sSysExchange) FetchTable(ctx context.Context, data g.Map) (tableData gdb.Result, handleID int64, err error) {
	var (
		taskID     int64
		providerID int64
		taskData   model.TaskData
	)
	taskID = int64(data["taskID"].(float64))
	providerID = int64(data["providerID"].(float64))
	err = g.Model("task").Where("task_id = ?", taskID).Scan(&taskData)
	g.Log().Info(ctx, "result:", taskData)
	if providerID != taskData.ProviderID {
		err = errors.New("provider not correct")
		return
	}
	handleID = taskData.HandleID
	tableData, err = g.DB(taskData.DBName).Model(taskData.TableName).Ctx(ctx).All()
	g.Log().Info(ctx, "infos:", tableData)
	return
}

func (s *sSysExchange) SendToMasking(ctx context.Context, data g.Map, tableData gdb.Result) (err error) {
	var reqData model.ProvideRawDataReq
	g.Log().Info(ctx, "data:", data)
	reqData.TaskID = int64(data["taskID"].(float64))
	reqData.HandleID = data["handleID"].(int64)
	var tableDetail model.TaskTableDetail
	for _, v := range tableData {
		g.Log().Info(ctx, "tableData:", v)
		//tableDetail.TableData = append(tableDetail.TableData, strings.Trim(gconv.String(v), "{}"))
		tableDetail.TableData = append(tableDetail.TableData, v)
	}
	tableDetail.SecureTableName = "test"
	reqData.Data = append(reqData.Data, tableDetail)
	g.Log().Info(ctx, "reqData:", reqData)
	client := g.Client()
	baseCfg := g.Cfg().MustGet(ctx, "baseApi.default").Map()
	tmpData := gconv.String(reqData) //信工所要转成字符串才能接收，不然格式不是json，这是为什么
	response, resErr := client.Post(ctx, baseCfg["address"].(string)+"/data/provideRawData", tmpData)
	if resErr != nil {
		err = resErr
	}
	defer response.Close()
	responseString := response.ReadAllString()
	g.Log().Info(ctx, "response:", responseString)
	//TODO task表状态升级为running
	return
}
