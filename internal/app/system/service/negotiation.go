package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	INegotiation interface {
		ResolveReq(ctx context.Context, req interface{}) (data g.Map, err error)
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
