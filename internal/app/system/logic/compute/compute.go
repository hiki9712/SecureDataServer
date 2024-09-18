package compute

import "github.com/tiger1103/gfast/v3/internal/app/system/service"

func init() {
	service.ComputeRegister(New())
}

type sCompute struct {
}

func New() *sCompute {
	return &sCompute{}
}

//TODO
