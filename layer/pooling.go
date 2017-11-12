package layer

import (
	"math"
	"strings"

	"github.com/rai-project/dllayer"
)

type Pooling struct {
	Base
	Operator string
	PadH     uint32
	PadW     uint32
	KernelH  uint32
	KernelW  uint32
	StrideH  uint32
	StrideW  uint32
	Global   bool
}

func (Pooling) Type() string {
	return "Pooling"
}

func (Pooling) Aliases() []string {
	return []string{"pooling"}
}

func (Pooling) Description() string {
	return ``
}

func (c *Pooling) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
	nIn := inputDimensions[0]
	cIn := inputDimensions[1]
	wIn := inputDimensions[2]
	hIn := inputDimensions[3]

	batchOut := nIn

	wOut := int64(math.Ceil(float64(wIn+2*int64(c.PadW)-int64(c.KernelW))/float64(c.StrideW))) + 1
	hOut := int64(math.Ceil(float64(hIn+2*int64(c.PadH)-int64(c.KernelH))/float64(c.StrideH))) + 1
	cOut := cIn

	if c.Global {
		wOut = 1
		hOut = 1
	}

	var numOps int64
	if c.Global {
		numOps = wIn * hIn * cIn * batchOut
	} else {
		numOps = wOut * hOut * int64(c.KernelH) * int64(c.KernelW) * cIn * batchOut
	}

	flops := dllayer.FlopsInformation{}
	switch strings.ToLower(c.Operator) {
	case "max":
		flops.Comparisons = numOps
	case "ave":
		flops.Additions = numOps
	}

	return &Information{
		name:             c.name,
		typ:              c.Type(),
		flops:            flops,
		inputDimensions:  inputDimensions,
		outputDimensions: []int64{nIn, cOut, hOut, wOut},
	}
}

func init() {
	dllayer.Register(&Pooling{})
}
