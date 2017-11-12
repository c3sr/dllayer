package network

import (
	"strings"

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

func kernelSizeOfPtr(v uint32, size *uint32) uint32 {
	if v != 0 {
		return v
	}
	if v == 0 && size == nil {
		return 0
	}
	return *size
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

func padSizeOfPtr(v uint32, size *uint32) uint32 {
	if v != 0 {
		return v
	}
	if v == 0 && size == nil {
		return 0
	}
	return *size
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

func strideSizeOfPtr(v uint32, size *uint32) uint32 {
	if v != 0 {
		return v
	}
	if v == 0 && size == nil {
		return 0
	}
	return *size
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

func (c Caffe) Information() (dllayer.FlopsInformation, dllayer.MemoryInformation) {
	infos := c.LayerInformations()
	flops := dllayer.FlopsInformation{}
	memory := dllayer.MemoryInformation{}
	for _, info := range infos {
		flops = flops.Add(info.Flops())
		memory = memory.Add(info.Memory())
	}
	return flops, memory
}

func (c Caffe) FlopsInformation() dllayer.FlopsInformation {
	flops, _ := c.Information()
	return flops
}

func (c Caffe) MemoryInformation() dllayer.MemoryInformation {
	_, mem := c.Information()
	return mem
}

func (c Caffe) LayerInformations() []dllayer.LayerInfo {
	infos := []dllayer.LayerInfo{}

	var inputDimensions []int64 = nil
	if len(c.InputDim) != 0 {
		inputDimensions = toInt64Slice(c.InputDim)
	} else if len(c.InputShape) != 0 &&
		c.InputShape[0] != nil &&
		len(c.InputShape[0].Dim) != 0 {
		inputDimensions = toInt64Slice(c.InputShape[0].Dim)
	}

	for _, lyr := range c.Layer {
		var layer dllayer.Layer
		switch strings.ToLower(lyr.Type) {
		case "input":
			layer = mkInput(lyr.InputParam)
		case "data":
			data := lyr.DataParam
			pp.Println(data)
		case "convolution":
			layer = mkConv(lyr.ConvolutionParam)
		case "relu":
			layer = mkReLU(lyr.ReluParam)
		case "dropout":
			layer = mkDropout(lyr.DropoutParam)
		case "innerproduct":
			layer = mkInnerProduct(lyr.InnerProductParam)
		case "pooling":
			layer = mkPooling(lyr.PoolingParam)
		case "batchnorm", "bn":
		case "lrn":
			layer = mkLRN(lyr.LrnParam)
		case "normalize":
		case "concat":
		case "softmax", "softmaxwithloss", "softmax_loss":
			layer = mkSoftMax(lyr.SoftmaxParam)
		case "flatten":
		case "eltwise":
		case "power":
		case "deconvolution":
		case "crop":
		case "scale":
		case "implicit":
		case "accuracy":
		case "permute":

		}
		if layer == nil {
			pp.Println(lyr)
			return nil
		}
		layer.SetName(lyr.Name)
		info := layer.LayerInformation(inputDimensions)
		infos = append(infos, info)
		inputDimensions = info.OutputDimensions()
	}

	for _, lyr := range c.Layers {
		switch lyr.Type {
		case caffe.V1LayerParameter_DATA:
			println("data layer")
		case caffe.V1LayerParameter_CONVOLUTION:
			c := mkConv(lyr.ConvolutionParam)
			info := c.LayerInformation(inputDimensions)
			inputDimensions = info.OutputDimensions()
			return nil
		}
	}

	return infos
}

func mkInput(param *caffe.InputParameter) dllayer.Layer {
	inputDimensions := toInt64Slice(param.Shape[0].Dim)
	return &layer.Input{
		N: inputDimensions[0],
		C: inputDimensions[1],
		W: inputDimensions[2],
		H: inputDimensions[3],
	}
}

func mkConv(param *caffe.ConvolutionParameter) dllayer.Layer {
	return &layer.Convolution{
		NumOutput: param.NumOutput,
		PadH:      padSizeOf(zeroIfNilUint32(param.PadH), param.Pad),
		PadW:      padSizeOf(zeroIfNilUint32(param.PadW), param.Pad),
		KernelH:   kernelSizeOf(param.KernelH, param.KernelSize),
		KernelW:   kernelSizeOf(param.KernelW, param.KernelSize),
		StrideH:   zeroToOne(strideSizeOf(param.StrideH, param.Stride)),
		StrideW:   zeroToOne(strideSizeOf(param.StrideW, param.Stride)),
		Dilation:  orSlice(param.Dilation, 1),
		Group:     orPtr(param.Group, 1),
	}
}

func mkReLU(param *caffe.ReLUParameter) dllayer.Layer {
	return &layer.ReLU{}
}

func mkDropout(param *caffe.DropoutParameter) dllayer.Layer {
	return &layer.Dropout{}
}

func mkSoftMax(param *caffe.SoftmaxParameter) dllayer.Layer {
	return &layer.SoftMax{}
}

func mkLRN(param *caffe.LRNParameter) dllayer.Layer {
	size := uint32(1)
	if param != nil && param.LocalSize != nil && *param.LocalSize != 0 {
		size = *param.LocalSize
	}
	region := "ACROSS_CHANNELS"
	if param != nil && param.NormRegion != nil && *param.NormRegion != 0 {
		region = "WITHIN_CHANNEL"
	}

	return &layer.LRN{
		Region: region,
		Size:   size,
	}
}

func mkPooling(param *caffe.PoolingParameter) dllayer.Layer {
	operator := "max"
	if param.Pool != nil && *param.Pool != 0 {
		operator = param.Pool.String()
	}
	global := false
	if param.GlobalPooling != nil {
		global = *param.GlobalPooling
	}
	return &layer.Pooling{
		Operator: strings.ToUpper(operator),
		PadH:     padSizeOfPtr(zeroIfNilUint32(param.PadH), param.Pad),
		PadW:     padSizeOfPtr(zeroIfNilUint32(param.PadW), param.Pad),
		KernelH:  kernelSizeOfPtr(param.KernelH, &param.KernelSize),
		KernelW:  kernelSizeOfPtr(param.KernelW, &param.KernelSize),
		StrideH:  zeroToOne(strideSizeOfPtr(param.StrideH, param.Stride)),
		StrideW:  zeroToOne(strideSizeOfPtr(param.StrideW, param.Stride)),
		Global:   global,
	}
}

func mkInnerProduct(param *caffe.InnerProductParameter) dllayer.Layer {
	return &layer.InnerProduct{
		NumOutput: param.NumOutput,
	}
}
