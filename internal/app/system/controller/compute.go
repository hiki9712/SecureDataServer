package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
)

var (
	Compute = computeController{}
)

type computeController struct {
	BaseController
}

func (c *computeController) SendRequest(ctx context.Context, req *system.ComputeSendReq) (res *system.ComputeSendRes, err error) {
	res = &system.ComputeSendRes{}
	//TODO
	return
}
