package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
)

type (
	ISysRegister interface {
		SendToBaseApi(ctx context.Context, data g.Map) (res *system.BaseAPIRes, err error)
		ResolveReq(ctx context.Context, req *system.RegisterReq) (data g.Map, err error)
		StoreToDB(ctx context.Context, data g.Map) (err error)
	}
)

var (
	localSysRegister ISysRegister
)

func SysRegister() ISysRegister {
	if localSysRegister == nil {
		panic("implement not found for interface ISysRegister")
	}
	return localSysRegister
}

func RegisterSysRegister(i ISysRegister) {
	localSysRegister = i
}
