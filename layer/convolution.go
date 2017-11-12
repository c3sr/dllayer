package layer

import (
	"math"

	"github.com/rai-project/dllayer"
)

type Convolution struct {
	Base
	NumOutput uint32
	PadH      uint32
	PadW      uint32
	KernelH   uint32
	KernelW   uint32
	StrideH   uint32
	StrideW   uint32
	Dilation  uint32
	Group     uint32
}

func (Convolution) Type() string {
	return "Convolution"
}

func (Convolution) Aliases() []string {
	return []string{"conv", "SpatialConvolution"}
}

func (Convolution) Description() string {
	return ``
}

func (c *Convolution) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
	/*
	  ## Add Input/Output Dimensions + Channels to each Node / Layer
	  # shape.dim: (    N   x   K   x   W   x   H   )
	  #              batch   channel  width   height
	  #               nIn      cIn     wIn     wOut
	*/

	nIn := inputDimensions[0]
	cIn := inputDimensions[1]
	wIn := inputDimensions[2]
	hIn := inputDimensions[3]

	batchOut := nIn

	kernelW := c.Dilation*(c.KernelW-1) + 1
	wOut := int64(math.Floor(float64(wIn+2*int64(c.PadW)-int64(kernelW))/float64(c.StrideW))) + 1
	kernelH := c.Dilation*(c.KernelH-1) + 1
	hOut := int64(math.Floor(float64(hIn+2*int64(c.PadH)-int64(kernelH))/float64(c.StrideH))) + 1
	cOut := int64(c.NumOutput)

	flops := dllayer.FlopsInformation{
		MultiplyAdds: (int64(c.KernelW*c.KernelH) * (wIn * hIn) * cIn * cOut * batchOut) / int64(c.Group),
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
	dllayer.Register(&Convolution{})
}
