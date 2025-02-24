package sysExchange

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
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
			negotiationDetailList []model.NegotiationDetail
			node                  *snowflake.Node
			taskId                int64
			insertData            model.TaskData
			postData              system.SendDataReq
			postDataList          map[string][]model.TaskData
		)

		serviceID := int64(data["serviceID"].(float64))
		err = g.Model("negotiation").Where("service_id = ?", serviceID).Scan(&negotiationDetailList)
		g.Log().Info(ctx, "negotiationDetailList:", negotiationDetailList)
		postDataList = make(map[string][]model.TaskData)
		for _, negotiationDetail := range negotiationDetailList {
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
			insertData.HandleName = data["handleName"].(string)
			insertData.Format = data["format"].(string)
			insertData.Protocol = int(data["protocol"].(float64))
			providerIDString := gconv.String(negotiationDetail.ProviderID)
			if _, exists := postDataList[providerIDString]; exists {
				postDataList[providerIDString] = append(postDataList[providerIDString], insertData)
			} else {
				postDataList[providerIDString] = []model.TaskData{insertData}
			}
			_, err = g.Model("task").Data(insertData).Insert()
		}
		g.Log().Info(ctx, "postDataList:", postDataList)
		for providerID, itemList := range postDataList {
			client := g.Client()
			postData.TableList = itemList
			//postData.TaskID = itemList[providerID]
			//taskIDList := make([]string, 0)
			//for _, item := range itemList {
			//	taskIDList = append(taskIDList, gconv.String(item.TaskID))
			//}
			//postData.HandleID = int64(data["handleID"].(float64))
			g.Log().Info(ctx, "postData:", postData)
			providerCfg := g.Cfg().MustGet(ctx, "providerAddress."+providerID).Map()
			response, resErr := client.Post(ctx, providerCfg["address"].(string)+"/api/v1/system/exchange/sendData", postData)
			if resErr != nil {
				err = resErr
				g.Log().Info(ctx, "resErr:", resErr)
				return
			}
			responseString := response.ReadAllString()
			g.Log().Info(ctx, "response:", responseString)
		}
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
	g.Log().Info(ctx, "data:", data)
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
	return
}

func (s *sSysExchange) SendToMasking(ctx context.Context, data g.Map) (err error) {
	g.Log().Info(ctx, "data:", data)
	var reqData model.ProvideRawDataReq
	for _, table := range data["tableList"].([]interface{}) {
		g.Log().Info(ctx, "table:", table)
		reqData.TaskID = int64(table.(map[string]interface{})["task_id"].(float64))
		reqData.HandleID = int64(table.(map[string]interface{})["handle_id"].(float64))
		var tableDetail model.TaskTableDetail
		var result []map[string]interface{}
		var resultList [][]map[string]interface{}
		var dbName string
		var tableName string
		var tableData gdb.Result
		dbName = table.(map[string]interface{})["db_name"].(string)
		tableName = table.(map[string]interface{})["table_name"].(string)
		tableData, err = g.DB(dbName).Model(tableName).Ctx(ctx).All()
		err = json.Unmarshal([]byte(gconv.String(tableData)), &result)
		if err != nil {
			return
		}
		resultList, _ = libUtils.JsonFileSplit(ctx, result, 64*1024)
		var uploadList []g.Map
		paramCfg := g.Cfg().MustGet(ctx, "uploadAddress").Map()
		param := g.Map{
			"username": paramCfg["username"].(string),
			"password": paramCfg["password"].(string),
			"addr":     paramCfg["addr"].(string),
			"path":     "/root/" + table.(map[string]interface{})["securetable_name"].(string) + "_",
		}
		for i, item := range resultList {
			upload := g.Map{
				"tableData": item,
			}
			uploadList = append(uploadList, upload)
			reqData.DataAddress = append(reqData.DataAddress, param["path"].(string)+gconv.String(i)+".json")
			uploadByte, _ := json.Marshal(upload)
			h := md5.New()
			h.Write(uploadByte)
			reqData.HashCode = append(reqData.HashCode, hex.EncodeToString(h.Sum(nil)))
		}
		err = libUtils.Upload(ctx, param, uploadList)
		if err != nil {
			return
		}
		tableDetail.SecureTableName = table.(map[string]interface{})["securetable_name"].(string)
	}
	client := g.Client()
	baseCfg := g.Cfg().MustGet(ctx, "baseApi.default").Map()
	tmpData := gconv.String(reqData) //信工所要转成字符串才能接收，不然格式不是json，这是为什么
	response, resErr := client.Post(ctx, baseCfg["address"].(string)+"/data/provideRawData", tmpData)
	g.Log().Info(ctx, "reqData:", tmpData)
	if resErr != nil {
		err = resErr
	}
	defer response.Close()
	responseString := response.ReadAllString()
	g.Log().Info(ctx, "response:", responseString)
	//TODO task表状态升级为running
	return
}

// 查看handle列表
func (s *sSysExchange) Listhandle(ctx context.Context, data g.Map) (handleDataList []model.HandleList, err error) {
	g.Log().Info(ctx, "listData:", data)
	if data["user_type"].(string) == "owner" {
		ownerhandleData, _ := g.Model("handle_reg_log").Fields("handle_id,service_id,service_name,provider_id").Order("update_time DESC").All()
		g.Log().Info(ctx, "ownerhandleData:", ownerhandleData)
		for _, v := range ownerhandleData {
			handleData := model.HandleList{
				ServiceID:   v["service_id"].Int64(),
				HandleID:    v["handle_id"].Int64(),
				ProviderID:  v["provider_id"].Int64(),
				ServiceName: fmt.Sprintf("%v", v["service_name"]),
			}
			handleDataList = append(handleDataList, handleData)
		}
	}
	return
}
