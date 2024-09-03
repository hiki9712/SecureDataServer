package sysRegister

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"math/rand/v2"
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
		//client := g.Client()
		//baseCfg := g.Cfg().MustGet(ctx, "baseApi.default").Map()
		//response, resErr := client.Post(ctx, baseCfg["address"].(string)+"/handle/register", data)
		//if resErr != nil {
		//	err = resErr
		//}
		//defer response.Close()
		//responseString := response.ReadAllString()
		//gjson.New(responseString).Scan(&res)

		//mock底层api接口返回
		res = &system.BaseAPIRes{
			Status:   "success",
			HandleID: rand.Int64(),
			Message:  "hello world",
		}
	})
	return
}

func (s *sSysRegister) ResolveReq(ctx context.Context, req *system.RegisterReq) (data g.Map, err error) {
	var (
		resData g.Map
		reqMap  g.Map
	)
	//请求参数初始化，无需处理err
	resJson, _ := json.Marshal(&system.BaseAPIReq{})
	_ = json.Unmarshal(resJson, &resData)

	//将req序列化为JSON
	reqJson, err := json.Marshal(req)
	if err != nil {
		return resData, err
	}

	//将JSON解析为map[string]interface{}
	err = json.Unmarshal(reqJson, &reqMap)
	if err != nil {
		return resData, err
	}
	//fmt.Println(reqMap)
	//将请求转化为底层api需要的格式,原子句柄和组合句柄处理不同
	if handleType, ok := reqMap["handleType"].(string); ok {
		resData["handleType"] = handleType
		if handleName, ok := reqMap["handleName"].(string); ok {
			resData["handleName"] = handleName
		}
		if databaseName, ok := reqMap["databaseName"].(string); ok {
			resData["databaseName"] = databaseName
		}
		//if handleType == "atomic" {
		//	resData["fieldNum"] = reqMap["fieldNum"]
		//	resData["atomicHandleContent"] = reqMap["atomicHandleContent"]
		//}
		//if handleType == "combined" {
		//	resData["atomicHandleNum"] = reqMap["atomicHandleNum"]
		//	resData["combinedHandleContent"] = reqMap["combinedHandleContent"]
		//}
		if keyValueContent, ok := reqMap["keyValueContent"].(string); ok {
			resData["keyValueContent"] = keyValueContent
		}
	}
	return resData, nil
}

func (s *sSysRegister) StoreToDB(ctx context.Context, handle *system.Handle) (err error) {
	_, err = g.Model("handle_register").Data(handle).Insert()
	return
}
