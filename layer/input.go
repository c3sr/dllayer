package layer

import (
	"github.com/rai-project/dllayer"
)

type Input struct {
	Base
	N int64
	C int64
	W int64
	H int64
}

func (Input) Type() string {
	return "Input"
}

func (Input) Aliases() []string {
	return []string{"input"}
}

func (Input) Description() string {
	return ``
}

func (c *Input) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
	return &Information{
		name:             c.name,
		typ:              c.Type(),
		inputDimensions:  []int64{c.N, c.C, c.W, c.H},
		outputDimensions: []int64{c.N, c.C, c.W, c.H},
	}
}

func init() {
	dllayer.Register(&Input{})
}
