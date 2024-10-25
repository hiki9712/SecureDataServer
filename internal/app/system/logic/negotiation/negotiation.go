package negotiation

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
)

func init() {
	service.RegisterNegotiation(New())
}

type sNegotiation struct {
}

func New() *sNegotiation {
	return &sNegotiation{}
}

func (s *sNegotiation) ResolveReq(ctx context.Context, req interface{}) (data g.Map, err error) {
	reqJson, err := json.Marshal(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(reqJson, &data)
	if err != nil {
		return
	}
	return
}

func (s *sNegotiation) SendNegotiationRequest(ctx context.Context, data g.Map) (serviceID int64, err error) {
	/*
		serviceNum, err := g.Model("negotiation").Count()
		g.Log().Info(ctx, "negotiation:ServiceNum:", serviceNum)
		negotiationData := &model.Negotiation{}
		negotiationData.ServiceID = int64(serviceNum + 1)
		negotiationData.ServiceName = data["serviceName"].(string)
		negotiationData.ServiceOwnerID = int64(data["serviceOwnerID"].(float64))
		negotiationData.ProviderID = int64(data["providerID"].(float64))
		//negotiationData.ProviderTable = data["tableName"].(string)
		//negotiationData.ProviderDB = data["databaseName"].(string)
		//negotiationData.SecureTableName = "secure_" + data["tableName"].(string)
		negotiationData.Status = consts.NegotiationStart
		negotiationData.DelFlag = 0
		negotiationData.CreateTime = time.Now()
		negotiationData.UpdateTime = time.Now()
		_, err = g.Model("negotiation").Insert(negotiationData)
	*/
	// var negotiationData *model.Negotiation
	// var negotiationLogData *model.NegotiationLog

	g.Log().Info(ctx, "data:", data)
	// 提取单条记录
	serviceName := data["serviceName"].(string)
	providerID := int64(data["providerID"].(float64))

	// 获取当前服务编号，使用相同的 serviceNum
	serviceID = libUtils.GenUniqId(ctx)
	g.Log().Info(ctx, "negotiation:ServiceID:", serviceID)

	// 假设 fieldContent 是一个包含多条记录的数组
	fieldContent := data["fieldContent"].([]interface{})

	// 遍历 fieldContent 提取 databaseName 和 tableName
	for _, field := range fieldContent {
		fieldMap := field.(map[string]interface{})
		databaseName := fieldMap["databaseName"].(string)
		tableName := fieldMap["tableName"].(string)

		// 创建negotiationID
		negotiationID := libUtils.GenUniqId(ctx)
		g.Log().Info(ctx, "negotiation:NegotiationID:", negotiationID)
		fieldMap["negotiationID"] = negotiationID

		// 创建新的 Negotiation 实例
		negotiationData := &model.Negotiation{}
		negotiationData.ServiceID = serviceID         // 使用相同的 serviceID
		negotiationData.NegotiationID = negotiationID // 使用新的 negotiationID
		negotiationData.ServiceName = serviceName
		negotiationData.ServiceOwnerID = int64(data["serviceOwnerID"].(float64))
		negotiationData.ProviderID = providerID
		negotiationData.ProviderTable = tableName // 记录当前表名
		negotiationData.ProviderDB = databaseName // 记录当前数据库名
		negotiationData.Status = consts.NegotiationStart
		negotiationData.DelFlag = 0
		negotiationData.CreateTime = time.Now()
		negotiationData.UpdateTime = time.Now()

		// // 创建新的 NegotiationLog 实例
		// negotiationLogData = &model.NegotiationLog{}
		// negotiationLogData.NegotiationLogID = 0          // 自增ID
		// negotiationLogData.ServiceID = serviceID         // 使用相同的 serviceID
		// negotiationLogData.NegotiationID = negotiationID // 使用新的 negotiationID
		// negotiationLogData.ServiceName = serviceName
		// negotiationLogData.ServiceOwnerID = int64(data["serviceOwnerID"].(float64))
		// negotiationLogData.ProviderID = providerID
		// negotiationLogData.ProviderTable = tableName // 记录当前表名
		// negotiationLogData.ProviderDB = databaseName // 记录当前数据库名
		// negotiationLogData.Status = consts.NegotiationStart
		// negotiationLogData.DelFlag = 0
		// negotiationLogData.CreateTime = time.Now()
		// negotiationLogData.UpdateTime = time.Now()

		// 插入到数据库
		_, err = g.Model("negotiation").Insert(negotiationData)
		if err != nil {
			return 0, err // 如果插入失败，返回错误
		}
		// _, log_err := g.Model("negotiation_pro_log").Insert(negotiationLogData)
		// if log_err != nil {
		// 	return 0, log_err // 如果插入失败，返回错误
		// }

	}

	// 发送协商信息给服务提供方
	client := g.Client()
	postData := data
	postData["serviceID"] = serviceID // 发送协商信息时，将 serviceID 也发送过去
	// 获取提供者的配置并发送 POST 请求
	providerCfg := g.Cfg().MustGet(ctx, "providerAddress."+strconv.FormatInt(providerID, 10)).Map()
	g.Log().Info(ctx, "providerCfg:", providerCfg["address"])
	g.Log().Info(ctx, "postData:", postData)
	response, resErr := client.Post(ctx, providerCfg["address"].(string)+"/api/v1/system/handle/negotiationToPro", postData)
	g.Log().Info(ctx, "postData:", postData)
	if resErr != nil {
		g.Log().Error(ctx, "发送协商信息失败:", resErr)
		return 0, resErr // 返回错误
	}
	// 在这里可以处理成功的响应，例如解析响应内容
	g.Log().Info(ctx, "协商信息发送成功，响应:", response)
	return serviceID, err
}

// 提供方处理接收到的协商
func (s *sNegotiation) SendNegotiationToProvider(ctx context.Context, data g.Map) (serviceID int64, err error) {
	/*
		serviceNum, err := g.Model("negotiation_pro").Count()
		g.Log().Info(ctx, "negotiation_pro:ServiceNum:", serviceNum)
		negotiationData := &model.Negotiation{}
		negotiationData.ServiceID = int64(serviceNum + 1)
		negotiationData.ServiceName = data["serviceName"].(string)
		negotiationData.ServiceOwnerID = int64(data["serviceOwnerID"].(float64))
		negotiationData.ProviderID = int64(data["providerID"].(float64))
		negotiationData.ProviderTable = data["tableName"].(string)
		negotiationData.ProviderDB = data["databaseName"].(string)
		negotiationData.SecureTableName = "secure_" + data["tableName"].(string)
		negotiationData.Status = consts.NegotiationStart
		negotiationData.DelFlag = 0
		negotiationData.CreateTime = time.Now()
		negotiationData.UpdateTime = time.Now()
		_, err = g.Model("negotiation_pro").Insert(negotiationData)
		return negotiationData.ServiceID, err
	*/
	var negotiationData *model.Negotiation
	var negotiationLogData *model.NegotiationLog
	g.Log().Info(ctx, "data:", data)
	// 提取单条记录
	serviceName := data["serviceName"].(string)
	providerID := int64(data["providerID"].(float64))

	// 获取当前服务编号
	serviceID = int64(data["serviceID"].(float64))
	g.Log().Info(ctx, "negotiation:ServiceID:", serviceID)

	// 假设 fieldContent 是一个包含多条记录的数组
	var fields []interface{}
	err = json.Unmarshal([]byte(data["fieldContent"].(string)), &fields)
	if err != nil {
		// 处理错误
		fmt.Println("Error parsing JSON:", err)
	}

	// 遍历 fieldContent 提取 databaseName 和 tableName
	for _, field := range fields {

		//fieldContent := data["fieldContent"].([]interface{})
		//for _, field := range fieldContent {
		fieldMap := field.(map[string]interface{})
		databaseName := fieldMap["databaseName"].(string)
		tableName := fieldMap["tableName"].(string)
		negotiationID := int64(fieldMap["negotiationID"].(float64))

		// 创建新的 Negotiation 实例
		negotiationData = &model.Negotiation{}
		negotiationData.ServiceID = serviceID         // 使用相同的 serviceID
		negotiationData.NegotiationID = negotiationID // 使用新的 negotiationID
		negotiationData.ServiceName = serviceName
		negotiationData.ServiceOwnerID = int64(data["serviceOwnerID"].(float64))
		negotiationData.ProviderID = providerID
		negotiationData.ProviderTable = tableName // 记录当前表名
		negotiationData.ProviderDB = databaseName // 记录当前数据库名
		negotiationData.Status = consts.NegotiationStart
		negotiationData.DelFlag = 0
		negotiationData.CreateTime = time.Now()
		negotiationData.UpdateTime = time.Now()

		// 创建新的 NegotiationLog 实例
		negotiationLogData = &model.NegotiationLog{}
		negotiationLogData.NegotiationLogID = 0          // 自增ID
		negotiationLogData.ServiceID = serviceID         // 使用相同的 serviceID
		negotiationLogData.NegotiationID = negotiationID // 使用新的 negotiationID
		negotiationLogData.ServiceName = serviceName
		negotiationLogData.ServiceOwnerID = int64(data["serviceOwnerID"].(float64))
		negotiationLogData.ProviderID = providerID
		negotiationLogData.ProviderTable = tableName // 记录当前表名
		negotiationLogData.ProviderDB = databaseName // 记录当前数据库名
		negotiationLogData.Status = consts.NegotiationStart
		negotiationLogData.DelFlag = 0
		negotiationLogData.CreateTime = time.Now()
		negotiationLogData.UpdateTime = time.Now()

		// 插入到数据库
		_, err = g.Model("negotiation_pro").Insert(negotiationData)
		if err != nil {
			return 0, err // 如果插入失败，返回错误
		}
		_, log_err := g.Model("negotiation_pro_log").Insert(negotiationLogData)
		if log_err != nil {
			return 0, log_err // 如果插入失败，返回错误
		}

	}
	return negotiationData.ServiceID, err
}

// 建表要求/状态存入提供方数据库
func (s *sNegotiation) SendNegotiationAgreeRequest(ctx context.Context, data g.Map) (err error) {
	negotiationID := int64(data["negotiationID"].(float64))
	//serviceID := int64(data["serviceID"].(float64))
	Agree := data["agree"].(bool)
	requestorIDVar, _ := g.Model("negotiation_pro").Where("negotiation_id=?", negotiationID).Value("service_owner_id")
	requestorID := requestorIDVar.Int64()
	g.Log().Info(ctx, "data:", data)

	if Agree {
		result, _ := g.Model("negotiation_pro").Where("negotiation_id=?", negotiationID).Fields("provider_db,provider_table").One()
		db := result["provider_db"]
		table := result["provider_table"]
		g.Log().Info(ctx, "result:", result)
		rawField, _ := libUtils.GetDBField(ctx, gconv.String(db), gconv.String(table))
		for _, field := range data["secureTableField"].([]interface{}) {
			g.Log().Info(ctx, "field:", field)
			for i, rawRow := range rawField {
				if rawRow.FieldName == field.(map[string]interface{})["fieldName"] {
					rawField[i].FieldNameNew = field.(map[string]interface{})["fieldNameNew"].(string)
					rawField[i].IsSecret = field.(map[string]interface{})["isSecret"].(string)
				}
			}
		}
		message := g.Map{
			"status":            consts.NegotiationAgree,
			"securetable_field": rawField,
			"securetable_name":  data["secureTableName"].(string),
		}
		_, err = g.Model("negotiation_pro").Data(message).Where("negotiation_id = ?", negotiationID).Update()
	} else {
		message := g.Map{
			"status":  consts.NegotiationReject,
			"message": gconv.String(data["message"].(string)),
		}
		_, err = g.Model("negotiation_pro").Data(message).Where("negotiation_id = ?", negotiationID).Update()
	}

	// 发送建表要求给需求方
	client := g.Client()
	postData := data
	postData["negotiationID"] = negotiationID // 发送协商信息时，将 negotiationID 也发送过去
	// 获取需求方的配置并发送 POST 请求
	requestorCfg := g.Cfg().MustGet(ctx, "requestorAddress."+strconv.FormatInt(requestorID, 10)).Map()
	g.Log().Info(ctx, "postData:", postData)
	response, resErr := client.Post(ctx, requestorCfg["address"].(string)+"/api/v1/system/handle/negotiationAgreeToReq", postData)
	if resErr != nil {
		g.Log().Error(ctx, "发送建表要求失败:", resErr)
		return resErr // 返回错误
	}
	// 在这里可以处理成功的响应，例如解析响应内容
	g.Log().Info(ctx, "建表要求发送成功，响应:", response)

	return
}

// 建表要求存入需求方数据库
func (s *sNegotiation) SendNegotiationAgreeToRequestor(ctx context.Context, data g.Map) (err error) {
	negotiationID := int64(data["negotiationID"].(float64))
	//serviceID := int64(data["serviceID"].(float64))
	Agree := data["agree"].(bool)
	g.Log().Info(ctx, "data:", data)
	if Agree {
		result, _ := g.Model("negotiation").Where("negotiation_id=?", negotiationID).Fields("provider_db,provider_table").One()
		db := result["provider_db"]
		table := result["provider_table"]
		g.Log().Info(ctx, "result:", result)
		rawField, _ := libUtils.GetDBField(ctx, gconv.String(db), gconv.String(table))

		var fields []interface{}
		err := json.Unmarshal([]byte(data["secureTableField"].(string)), &fields)
		if err != nil {
			// 处理错误
			fmt.Println("Error parsing JSON:", err)
		}
		for _, field := range fields {
			g.Log().Info(ctx, "field:", field)
			for i, rawRow := range rawField {
				if rawRow.FieldName == field.(map[string]interface{})["fieldName"] {
					rawField[i].FieldNameNew = field.(map[string]interface{})["fieldNameNew"].(string)
					rawField[i].IsSecret = field.(map[string]interface{})["isSecret"].(string)
				}
			}
		}
		message := g.Map{
			"status":            consts.NegotiationAgree,
			"securetable_field": rawField,
			"securetable_name":  data["secureTableName"].(string),
		}
		_, err = g.Model("negotiation").Data(message).Where("negotiation_id = ?", negotiationID).Update()
	} else {
		message := g.Map{
			"status":  consts.NegotiationReject,
			"message": gconv.String(data["message"].(string)),
		}
		_, err = g.Model("negotiation").Data(message).Where("negotiation_id = ?", negotiationID).Update()
	}
	return
}

func (s *sNegotiation) ListNegotiation(ctx context.Context, data g.Map) (negotiationDataList []model.NegotiationList, err error) {
	g.Log().Info(ctx, "listData:", data)
	if data["user_type"].(string) == "provider" {
		providerData, _ := g.Model("negotiation_pro").Fields("status,service_id,service_name,provider_db,provider_table,negotiation_id").Where("provider_id = ? AND status = ?", int64(data["provider_id"].(float64)), "start").Order("update_time DESC").All()
		g.Log().Info(ctx, "providerData:", providerData)
		for _, v := range providerData {
			negotiationData := model.NegotiationList{
				Status:        fmt.Sprintf("%v", v["status"]),
				ServiceID:     v["service_id"].Int64(),
				NegotiationID: v["negotiation_id"].Int64(), // 用于区分不同的建表信息
				TableName:     fmt.Sprintf("%v", v["provider_table"]),
				DBName:        fmt.Sprintf("%v", v["provider_db"]),
				ServiceName:   fmt.Sprintf("%v", v["service_name"]),
			}
			negotiationDataList = append(negotiationDataList, negotiationData)
		}
	}
	if data["user_type"].(string) == "owner" {
		ownerData, _ := g.Model("negotiation").Fields("status,service_id,service_name,provider_db,provider_table").Where("service_owner_id = ?", int64(data["owner_id"].(float64))).Order("update_time DESC").All()
		g.Log().Info(ctx, "ownerData:", ownerData)
		for _, v := range ownerData {
			negotiationData := model.NegotiationList{
				Status:      fmt.Sprintf("%v", v["status"]),
				ServiceID:   v["service_id"].Int64(),
				TableName:   fmt.Sprintf("%v", v["provider_table"]),
				DBName:      fmt.Sprintf("%v", v["provider_db"]),
				ServiceName: fmt.Sprintf("%v", v["service_name"]),
			}
			negotiationDataList = append(negotiationDataList, negotiationData)
		}
	}
	return
}

func (s *sNegotiation) BuildMySQLDB(ctx context.Context, data g.Map) (err error) {
	var (
		secureTableInfo g.Map
		tableFieldList  []g.Map
	)

	serviceID := int64(data["serviceID"].(float64))

	/*
		secureTableData, err := g.Model("negotiation").Fields("securetable_name,securetable_field").Where("service_id = ?", serviceID).One()
		if err != nil {
			return
		}
		err = gjson.DecodeTo(secureTableData, &secureTableInfo)
		if err != nil {
			return
		}
		err = gjson.DecodeTo(secureTableData["securetable_field"], &tableFieldList)
		if err != nil {
			return
		}
		sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (", secureTableInfo["securetable_name"])
		for _, tableField := range tableFieldList {
			sql += fmt.Sprintf("`%s` %s,", tableField["fieldName"], tableField["fieldType"])
			if tableField["isSecret"].(string) == "True" {
				sql += fmt.Sprintf("`%s` %s,", tableField["fieldNameNew"], tableField["fieldType"])
			}
			//TODO 主键该配哪个
			if tableField["isKey"].(string) == "PRI" {
				PrimaryKey = append(PrimaryKey, gconv.String(tableField["fieldName"]))
			}
		}
		g.Log().Info(ctx, "primaryKey:", PrimaryKey)
		sql += fmt.Sprintf("PRIMARY KEY (`%s`)", strings.Join(PrimaryKey, ","))
		sql = strings.TrimRight(sql, ",")
		sql += fmt.Sprintf(") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
		g.Log().Info(ctx, "sql:", sql)
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Model("negotiation").Data("status", consts.NegotiationFail, "update_time", time.Now()).Where("service_id = ?", serviceID).Update()
			return
		}
		g.Model("negotiation").Data("status", consts.NegotiationSuccess, "update_time", time.Now()).Where("service_id = ?", serviceID).Update()
	*/

	// 查询所有与 serviceID 相关的记录
	secureTableDataList, err := g.Model("negotiation").Fields("securetable_name,securetable_field").Where("service_id = ?", serviceID).All()
	g.Log().Info(ctx, "secureTableDataList:", secureTableDataList)

	if err != nil {
		return
	}

	// 遍历每个 secureTableData
	for _, secureTableData := range secureTableDataList {
		g.Log().Info(ctx, "secureTableData:", secureTableData)
		err = gjson.DecodeTo(secureTableData, &secureTableInfo)
		if err != nil {
			return
		}
		err = gjson.DecodeTo(secureTableData["securetable_field"], &tableFieldList)
		if err != nil {
			return
		}
		var PrimaryKey []string
		// 构建 SQL 语句
		sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (", secureTableInfo["securetable_name"])
		for _, tableField := range tableFieldList {
			sql += fmt.Sprintf("`%s` %s,", tableField["fieldName"], tableField["fieldType"])
			if tableField["isSecret"].(string) == "True" {
				sql += fmt.Sprintf("`%s` %s,", tableField["fieldNameNew"], tableField["fieldType"])
			}
			if tableField["isKey"].(string) == "PRI" {
				PrimaryKey = append(PrimaryKey, gconv.String(tableField["fieldName"]))
			}
		}
		g.Log().Info(ctx, "primaryKey:", PrimaryKey)
		sql += fmt.Sprintf("PRIMARY KEY (`%s`)", strings.Join(PrimaryKey, "`,`"))
		sql = strings.TrimRight(sql, ",")
		sql += fmt.Sprintf(") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
		g.Log().Info(ctx, "sql:", sql)

		// 执行创建表
		_, err = g.DB("biz_dms").Exec(ctx, sql)
		if err != nil {
			g.Model("negotiation").Data("status", consts.NegotiationFail, "update_time", time.Now()).Where("service_id = ?", serviceID).Update()
			return
		}
	}
	g.Model("negotiation").Data("status", consts.NegotiationSuccess, "update_time", time.Now()).Where("service_id = ?", serviceID).Update()

	return
}
