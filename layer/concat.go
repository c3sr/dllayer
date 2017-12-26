package layer

import (
	"github.com/rai-project/dllayer"
)

type Concat struct {
	Base               `json:",inline,flatten",omitempty"`
	ParentsInformation []dllayer.LayerInfo `json:"parents_information,omitempty"`
}

func (Concat) Type() string {
	return "Concat"
}

func (Concat) Aliases() []string {
	return []string{"concat"}
}

func (Concat) Description() string {
	return ``
}

func (c *Concat) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
	nIn := inputDimensions[0]
	cIn := inputDimensions[1]
	wIn := inputDimensions[2]
	hIn := inputDimensions[3]

	wOut := wIn
	hOut := hIn
	cIn = 0
	for _, parent := range c.ParentsInformation {
		outputDimensions := parent.OutputDimensions()
		cIn += outputDimensions[1]
	}
	cOut := cIn

	flops := dllayer.FlopsInformation{}

	return &Information{
		name:             c.name,
		typ:              c.Type(),
		flops:            flops,
		inputDimensions:  inputDimensions,
		outputDimensions: []int64{nIn, cOut, hOut, wOut},
	}
}

func init() {
	dllayer.Register(&Concat{})
}
