package layer

import (
	"github.com/c3sr/dllayer"
)

type Dropout struct {
	Base `json:",inline,flatten""`
}

func (Dropout) Type() string {
	return "Dropout"
}

func (Dropout) Aliases() []string {
	return []string{"dropout"}
}

func (Dropout) Description() string {
	return ``
}

func (c *Dropout) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
	nIn := inputDimensions[0]
	cIn := inputDimensions[1]
	wIn := inputDimensions[2]
	hIn := inputDimensions[3]

	flops := dllayer.FlopsInformation{
		Comparisons: wIn * hIn * cIn * nIn,
	}

	return &Information{
		name:             c.name,
		typ:              c.Type(),
		flops:            flops,
		inputDimensions:  inputDimensions,
		outputDimensions: inputDimensions,
	}
}

func init() {
	dllayer.Register(&Dropout{})
}
