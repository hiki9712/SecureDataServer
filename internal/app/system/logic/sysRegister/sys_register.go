package sysRegister

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"math/rand/v2"
	"time"
)

func init() {
	service.RegisterSysRegister(New())
}

type sSysRegister struct {
	casBinRegisterPrefix string
}

func New() *sSysRegister {
	return &sSysRegister{
		casBinRegisterPrefix: "r_",
	}
}

func (s *sSysRegister) SendToBaseApi(ctx context.Context, data g.Map) (res *system.BaseAPIRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		client := g.Client()
		baseCfg := g.Cfg().MustGet(ctx, "baseApi.default").Map()
		g.Log().Info(ctx, "ip:", "http://"+baseCfg["address"].(string)+"/handle/register")
		g.Log().Info(ctx, "data:", data)
		sendData := g.Map{
			"handleType":      data["handleType"],
			"keyValueContent": data["keyValueContent"],
			"keyValueCount":   data["keyValueCount"],
		}
		g.Log().Info(ctx, "senddata:", sendData)
		jsonData, _ := gjson.Encode(sendData)
		response, resErr := client.Post(ctx, "http://"+baseCfg["address"].(string)+"/handle/register", jsonData, g.Map{"Content-Type": "application/json"})
		if resErr != nil {
			err = resErr
		}
		defer response.Close()
		responseString := response.ReadAllString()
		gjson.New(responseString).Scan(&res)
		g.Log().Info(ctx, "response:", responseString)
		//mock底层api接口返回
		res = &system.BaseAPIRes{
			Status:   "success",
			HandleID: rand.Int64(),
			Message:  "",
		}
	})
	return
}

func (s *sSysRegister) ResolveReq(ctx context.Context, req *system.RegisterReq) (data g.Map, err error) {
	var (
	//resData g.Map
	)
	////底层api请求参数初始化，无需处理err
	//resJson, _ := json.Marshal(&system.BaseAPIReq{})
	//_ = json.Unmarshal(resJson, &resData)

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
	//fmt.Println(reqMap)
	//将请求转化为底层api需要的格式,原子句柄和组合句柄处理不同
	//if handleType, ok := reqMap["handleType"].(string); ok {
	//	resData["handleType"] = handleType
	//	if handleName, ok := reqMap["handleName"].(string); ok {
	//		resData["handleName"] = handleName
	//	}
	//	if databaseName, ok := reqMap["databaseName"].(string); ok {
	//		resData["databaseName"] = databaseName
	//	}
	//	//if handleType == "atomic" {
	//	//	resData["fieldNum"] = reqMap["fieldNum"]
	//	//	resData["atomicHandleContent"] = reqMap["atomicHandleContent"]
	//	//}
	//	//if handleType == "combined" {
	//	//	resData["atomicHandleNum"] = reqMap["atomicHandleNum"]
	//	//	resData["combinedHandleContent"] = reqMap["combinedHandleContent"]
	//	//}
	//	if keyValueContent, ok := reqMap["keyValueContent"].(string); ok {
	//		resData["keyValueContent"] = keyValueContent
	//	}
	//}
	//return resData, nil
}

func (s *sSysRegister) StoreToDB(ctx context.Context, data g.Map) (err error) {
	HandleNum, err := g.Model("handle_reg").Count()
	g.Log().Info(ctx, "datakeyvaluecontent:", data["keyValueContent"].([]interface{})[0], HandleNum)
	insertData := &model.AtomHandleReg{}
	insertData.HandleID = int64(HandleNum + 1)
	insertData.HandleName = data["handleName"].(string)
	insertData.HandleType = data["handleType"].(string)
	insertData.ServiceID = int64(data["serviceID"].(float64))
	insertData.ServiceName = data["serviceName"].(string)
	insertData.ProviderID = int64(data["providerID"].(float64))
	insertData.KeyValueCount = int(data["keyValueCount"].(float64))
	insertData.KeyValueContent = data["keyValueContent"].([]interface{})
	insertData.DelFlag = 0
	//insertData.CreateBy TODO
	insertData.CreateTime = time.Now()
	//insertData.UpdateBy
	insertData.UpdateTime = time.Now()
	//insertData.Remark
	_, err = g.Model("handle_reg").Data(insertData).Insert()
	//_, err = g.Model("handle_register").Data(handle).Insert()
	return
}
