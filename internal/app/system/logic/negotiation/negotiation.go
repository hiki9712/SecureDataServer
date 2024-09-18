package negotiation

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"strings"
	"time"
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
	serviceNum, err := g.Model("negotiation").Count()
	g.Log().Info(ctx, "negotiation:ServiceNum:", serviceNum)
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
	_, err = g.Model("negotiation").Insert(negotiationData)
	return negotiationData.ServiceID, err
}

func (s *sNegotiation) SendNegotiationAgreeRequest(ctx context.Context, data g.Map) (err error) {
	serviceID := int64(data["serviceID"].(float64))
	Agree := data["agree"].(bool)
	if Agree {
		message := g.Map{
			"status":            consts.NegotiationAgree,
			"securetable_field": data["secureTableField"].(interface{}),
			"securetable_name":  data["secureTableName"].(string),
		}
		_, err = g.Model("negotiation").Data(message).Where("service_id = ?", serviceID).Update()
	} else {
		message := g.Map{
			"status":  consts.NegotiationReject,
			"message": gconv.String(data["message"].(string)),
		}
		_, err = g.Model("negotiation").Data(message).Where("service_id = ?", serviceID).Update()
	}
	return
}

func (s *sNegotiation) ListNegotiation(ctx context.Context, data g.Map) (negotiationDataList []model.NegotiationList, err error) {
	g.Log().Info(ctx, "listData:", data)
	if data["user_type"].(string) == "provider" {
		providerData, _ := g.Model("negotiation").Fields("status,service_id,service_name,provider_db,provider_table").Where("provider_id = ?", int64(data["provider_id"].(float64))).All()
		g.Log().Info(ctx, "providerData:", providerData)
		for _, v := range providerData {
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
	if data["user_type"].(string) == "owner" {
		ownerData, _ := g.Model("negotiation").Fields("status,service_id,service_name,provider_db,provider_table").Where("service_owner_id = ?", int64(data["owner_id"].(float64))).All()
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
	}
	sql = strings.TrimRight(sql, ",")
	sql += fmt.Sprintf(") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
	g.Log().Info(ctx, "sql:", sql)
	_, err = g.DB().Exec(ctx, sql)
	if err != nil {
		g.Model("negotiation").Where("service_id = ?", serviceID).Update("status", consts.NegotiationFail)
		return
	}
	g.Model("negotiation").Where("service_id = ?", serviceID).Update("status", consts.NegotiationSuccess)
	return
}
