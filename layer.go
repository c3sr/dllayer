package dllayer

type Layer interface {
	Name() string
	Aliases() []string
	LayerInformation(inputDimensions []int64) LayerInfo
}

type LayerInfo interface {
	Flops() int64
	Memory() int64
	InputDimensions() []int64
	OutputDimensions() []int64
}
