package layer

import (
	"github.com/rai-project/dllayer"
)

type BatchNorm struct {
	Base
}

func (BatchNorm) Type() string {
	return "BatchNorm"
}

func (BatchNorm) Aliases() []string {
	return []string{"batchnorm", "bn"}
}

func (BatchNorm) Description() string {
	return ``
}

func (c *BatchNorm) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
	nIn := inputDimensions[0]
	cIn := inputDimensions[1]
	wIn := inputDimensions[2]
	hIn := inputDimensions[3]

	numOps := wIn * hIn * cIn * nIn
	flops := dllayer.FlopsInformation{
		Additions: numOps,
		Divisions: numOps,
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
	dllayer.Register(&BatchNorm{})
}
