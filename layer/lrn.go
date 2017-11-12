package layer

import (
	"github.com/rai-project/dllayer"
)

type LRN struct {
	Base
	Region string
	Size   uint32
}

func (LRN) Type() string {
	return "LRN"
}

func (LRN) Aliases() []string {
	return []string{"lrn"}
}

func (LRN) Description() string {
	return ``
}

func (c *LRN) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
	nIn := inputDimensions[0]
	cIn := inputDimensions[1]
	wIn := inputDimensions[2]
	hIn := inputDimensions[3]

	//  Each input value is divided by (1+(α/n)∑xi^2)^β
	numInputs := wIn * hIn * cIn * nIn
	size := int64(c.Size)
	if c.Region == "WITHIN_CHANNEL" {
		size = size * size
	}

	flops := dllayer.FlopsInformation{
		MultiplyAdds:    numInputs * size, // (∑xi^2)
		Additions:       numInputs,        //  (1+...)
		Exponentiations: numInputs,        // (...)^β
		Divisions:       numInputs * 2,    // (α/n)*... + divide by sum
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
	dllayer.Register(&LRN{})
}
