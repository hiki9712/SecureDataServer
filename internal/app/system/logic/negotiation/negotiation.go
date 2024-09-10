package negotiation

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
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
