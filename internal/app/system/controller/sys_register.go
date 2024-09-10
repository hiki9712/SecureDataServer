package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	_ "github.com/tiger1103/gfast/v3/internal/app/system/logic"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var (
	Register = registerController{}
)

type registerController struct {
	BaseController
}

func (c *registerController) Register(ctx context.Context, req *system.RegisterReq) (res *system.RegisterRes, err error) {
	var (
		data g.Map
		//baseAPIres *system.BaseAPIRes
	)
	res = &system.RegisterRes{
		HandleName: req.HandleName,
		Status:     "success",
	}
	data, err = service.SysRegister().ResolveReq(ctx, req)
	g.Log().Info(ctx, "register success", data)
	//baseAPIres, err = service.SysRegister().SendToBaseApi(ctx, data)
	//handle := &system.Handle{
	//	HandleName: req.HandleName,
	//	HandleId:   baseAPIres.HandleID,
	//	HandleType: req.Type,
	//}
	//if req.Type == "atomic" {
	//	jsonStringAtomic, _ := gjson.New(data["atomicHandleContent"]).ToJsonString()
	//	handle.HandleContent = jsonStringAtomic
	//} else {
	//	jsonStringCombined, _ := gjson.New(data["combinedHandleContent"]).ToJsonString()
	//	handle.HandleContent = jsonStringCombined
	//}
	err = service.SysRegister().StoreToDB(ctx, data)
	//if err != nil {
	//	res.Status = "fail"
	//	return
	//}
	return
}

func (c *registerController) Negotiation(ctx context.Context, req *system.RegisterNegotiationReq) (res *system.RegisterNegotiationRes, err error) {
	data, err := service.Negotiation().ResolveReq(ctx, req)
	g.Log().Info(ctx, "negotiation success", data)
	res = &system.RegisterNegotiationRes{}
	return
}
