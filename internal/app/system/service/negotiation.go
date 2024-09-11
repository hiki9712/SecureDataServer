package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type (
	INegotiation interface {
		ResolveReq(ctx context.Context, req interface{}) (data g.Map, err error)
		SendNegotiationRequest(ctx context.Context, data g.Map) (serviceID int64, err error)
		SendNegotiationAgreeRequest(ctx context.Context, data g.Map) (err error)
		BuildMySQLDB(ctx context.Context, data g.Map) (err error)
		ListNegotiation(ctx context.Context, data g.Map) (negotiationData []model.NegotiationList, err error)
	}
)

var (
	localNegotiation INegotiation
)

func Negotiation() INegotiation {
	if localNegotiation == nil {
		panic("negotiation not initialized")
	}
	return localNegotiation
}

func RegisterNegotiation(i INegotiation) {
	localNegotiation = i
}
