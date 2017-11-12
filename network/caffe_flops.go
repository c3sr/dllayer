package network

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/caffe"
	"github.com/rai-project/dllayer"
	"github.com/rai-project/dllayer/layer"
)

func kernelSizeOf(v uint32, size []uint32) uint32 {
	if v != 0 {
		return v
	}
	if v == 0 && len(size) == 0 {
		return 0
	}
	return size[0]
}

func padSizeOf(v uint32, size []uint32) uint32 {
	if v != 0 {
		return v
	}
	if v == 0 && len(size) == 0 {
		return 0
	}
	return size[0]
}

func strideSizeOf(v uint32, size []uint32) uint32 {
	if v != 0 {
		return v
	}
	if v == 0 && len(size) == 0 {
		return 0
	}
	return size[0]
}

func zeroIfNilUint32(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

func orSlice(slc []uint32, v uint32) uint32 {
	if len(slc) == 0 {
		return v
	}
	return slc[0]
}

func orPtr(ptr *uint32, v uint32) uint32 {
	if ptr != nil {
		return *ptr
	}
	return v
}

func zeroToOne(v uint32) uint32 {
	if v == 0 {
		return 1
	}
	return v
}

func (c Caffe) LayerInformation() []dllayer.LayerInfo {
	infos := []dllayer.LayerInfo{}

	inputDimensions := []int64{}
	if len(c.InputDim) != 0 {
		inputDimensions = toInt64Slice(c.InputDim)
	} else if len(c.InputShape) != 0 &&
		c.InputShape[0] != nil &&
		len(c.InputShape[0].Dim) != 0 {
		inputDimensions = toInt64Slice(c.InputShape[0].Dim)
	}

	for _, lyr := range c.Layer {
		switch lyr.Type {
		case "Input":
			input := lyr.InputParam
			inputDimensions = toInt64Slice(input.Shape[0].Dim)
		case "data":
			data := lyr.DataParam
			_ = data
			pp.Println(data)
		case "Convolution":
			c := mkConv(lyr.ConvolutionParam)
			info := c.LayerInformation(inputDimensions)
			pp.Println(info)
			return nil
		}
	}

	for _, lyr := range c.Layers {
		switch lyr.Type {
		case caffe.V1LayerParameter_DATA:
			println("data layer")
		case caffe.V1LayerParameter_CONVOLUTION:
			c := mkConv(lyr.ConvolutionParam)
			pp.Println(lyr.ConvolutionParam)
			pp.Println(c)
			info := c.LayerInformation(inputDimensions)
			pp.Println(info)
			return nil
		}
	}

	return infos
}

func mkConv(conv *caffe.ConvolutionParameter) dllayer.Layer {
	return layer.Convolution{
		NumOutput: conv.NumOutput,
		PadH:      padSizeOf(zeroIfNilUint32(conv.PadH), conv.Pad),
		PadW:      padSizeOf(zeroIfNilUint32(conv.PadW), conv.Pad),
		KernelH:   kernelSizeOf(conv.KernelH, conv.KernelSize),
		KernelW:   kernelSizeOf(conv.KernelW, conv.KernelSize),
		StrideH:   zeroToOne(strideSizeOf(conv.StrideH, conv.Stride)),
		StrideW:   zeroToOne(strideSizeOf(conv.StrideW, conv.Stride)),
		Dilation:  orSlice(conv.Dilation, 1),
		Group:     orPtr(conv.Group, 1),
	}
}
