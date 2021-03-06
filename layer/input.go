package layer

import (
	"github.com/c3sr/dllayer"
)

type Input struct {
	Base `json:",inline,flatten",omitempty"`
	N    int64 `json:"n,omitempty"`
	C    int64 `json:"c,omitempty"`
	W    int64 `json:"w,omitempty"`
	H    int64 `json:"h,omitempty"`
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
	batchSize := c.N
	batchSize = 1
	return &Information{
		name:             c.name,
		typ:              c.Type(),
		inputDimensions:  []int64{batchSize, c.C, c.W, c.H},
		outputDimensions: []int64{batchSize, c.C, c.W, c.H},
	}
}

func init() {
	dllayer.Register(&Input{})
}
