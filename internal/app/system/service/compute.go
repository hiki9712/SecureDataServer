package service

type (
	ICompute interface {
		//TODO
	}
)

var (
	localCompute ICompute
)

func Compute() ICompute {
	if localCompute == nil {
		panic("c")
	}
	return localCompute
}

func ComputeRegister(i ICompute) {
	localCompute = i
}
