package layer

import (
	"github.com/rai-project/dllayer"
)

type Information struct {
	typ              string
	name             string
	flops            dllayer.FlopsInformation
	memory           dllayer.MemoryInformation
	inputDimensions  []int64
	outputDimensions []int64
}

func (layer *Information) Type() string {
	return layer.typ
}

func (layer *Information) Name() string {
	return layer.name
}

func (layer *Information) Flops() dllayer.FlopsInformation {
	return layer.flops
}
func (layer *Information) Memory() dllayer.MemoryInformation {
	return layer.memory
}
func (layer *Information) InputDimensions() []int64 {
	return layer.inputDimensions
}
func (layer *Information) OutputDimensions() []int64 {
	return layer.outputDimensions
}
