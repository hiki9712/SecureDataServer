package system

import "github.com/gogf/gf/v2/frame/g"

type ComputeSendReq struct {
	g.Meta `path:"/compute/sendRequest" method:"get" tags:"计算" summary:"发送数据"`
	//TODO 前端请求业务系统进行数据查询
}

type ComputeSendRes struct {
}
