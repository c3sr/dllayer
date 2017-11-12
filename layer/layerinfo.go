package layer

type Information struct {
	typ              string
	name             string
	flops            int64
	memory           int64
	inputDimensions  []int64
	outputDimensions []int64
}

func (layer *Information) Type() string {
	return layer.typ
}

func (layer *Information) Name() string {
	return layer.name
}

func (layer *Information) Flops() int64 {
	return layer.flops
}
func (layer *Information) InputDimensions() []int64 {
	return layer.inputDimensions
}
func (layer *Information) OutputDimensions() []int64 {
	return layer.outputDimensions
}
func (layer *Information) Memory() int64 {
	return layer.memory
}
