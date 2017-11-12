package layer

import (
	"math"

	"github.com/rai-project/dllayer"
)

type Convolution struct {
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

func (Convolution) Name() string {
	return "Convolution"
}

func (Convolution) Aliases() []string {
	return []string{"conv", "SpatialConvolution"}
}

func (Convolution) Description() string {
	return ``
}

func (c Convolution) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
	/*
	  ## Add Input/Output Dimensions + Channels to each Node / Layer
	  # shape.dim: (    N   x   K   x   W   x   H   )
	  #              batch   channel  width   height
	  #               chIn    chOut   wIn     wOut
	*/

	nIn := inputDimensions[0]
	cIn := inputDimensions[1]
	hIn := inputDimensions[2]
	wIn := inputDimensions[3]

	_, _, _, _ = nIn, cIn, hIn, wIn

	batchOut := nIn

	kernelW := c.Dilation*(c.KernelW-1) + 1
	wOut := int64(math.Floor(float64(wIn+2*int64(c.PadW)-int64(kernelW))/float64(c.StrideW))) + 1
	kernelH := c.Dilation*(c.KernelH-1) + 1
	hOut := int64(math.Floor(float64(hIn+2*int64(c.PadH)-int64(kernelH))/float64(c.StrideH))) + 1
	chOut := int64(c.NumOutput)

	flops := (int64(c.KernelW*c.KernelH) * (wOut * hOut) * nIn * chOut * batchOut) / int64(c.Group)

	return &Information{
		flops:            flops,
		inputDimensions:  inputDimensions,
		outputDimensions: []int64{nIn, chOut, hOut, wOut},
	}
}

func init() {
	dllayer.Register(Convolution{})
}
