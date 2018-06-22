package onnx

import (
	"strings"

	"github.com/k0kubun/pp"
	"github.com/rai-project/dllayer"
	"github.com/rai-project/onnx"
)

func (o Onnx) Information() (dllayer.FlopsInformation, dllayer.MemoryInformation) {
	infos := o.LayerInformations()
	flops := dllayer.FlopsInformation{}
	memory := dllayer.MemoryInformation{}
	for _, info := range infos {
		flops = flops.Add(info.Flops())
		memory = memory.Add(info.Memory())
	}
	return flops, memory
}

func (o Onnx) FlopsInformation() dllayer.FlopsInformation {
	flops, _ := o.Information()
	return flops
}

func (o Onnx) MemoryInformation() dllayer.MemoryInformation {
	_, mem := o.Information()
	return mem
}

func getDimValues(dims []*onnx.TensorShapeProto_Dimension) []int64 {
	ret := []int64{}
	for _, dim := range dims {
		ret = append(ret, dim.GetDimValue())
	}

	return ret
}

func (o Onnx) GetNodeInputs(node *onnx.NodeProto) []*onnx.ValueInfoProto {
	ret := []*onnx.ValueInfoProto{}
	for _, name := range node.GetInput() {
		val, ok := o.values[name]
		if ok {
			ret = append(ret, val)
		}
	}

	return ret
}

func (o Onnx) GetNodeInputDimensions(node *onnx.NodeProto) []int64 {
	ret := []int64{}
	inputs := o.GetNodeInputs(node)
	for _, input := range o.GetNodeInputs(node) {
		name := input.GetName()

	}

	return ret
}

// LayerInformations uses the first input
func (o Onnx) LayerInformations() []dllayer.LayerInfo {
	inputs := o.Graph.GetInput()
	inputDims := inputs[0].GetType().GetTensorType().GetShape().GetDim()

	if len(inputs) == 0 || inputs[0] == nil || len(inputDims) == 0 {
		log.Info("no input info for graph", o.Graph.GetName())
		return nil
	}

	ret := []dllayer.LayerInfo{}

	// The nodes in the graph are sorted topologically
	for _, node := range o.Graph.GetNode() {
		name := node.GetName()
		layer := o.mkLayer(node)

		if layer == nil {
			pp.Println("failed to create ", name)
			return nil
		}

		inputDimensions := o.GetNodeInputDimensions(node)
		info := layer.LayerInformation(inputDimensions)
		o.layerInformation[name] = info
		ret = append(ret, info)
	}

	return ret
}

func (o Onnx) mkLayer(node *onnx.NodeProto) dllayer.Layer {
	var layer dllayer.Layer
	layerType := strings.ToLower(node.GetOpType())
	// switch layerType {
	// case "Constant":
	// 	layer = mkInput(node.GetInput())
	// case "data":
	// 	layer = mkData(node.InputParam)
	// case "convolution":
	// 	layer = mkConv(node.ConvolutionParam)
	// case "relu":
	// 	layer = mkReLU(node.ReluParam)
	// case "dropout":
	// 	layer = mkDropout(node.DropoutParam)
	// case "innerproduct", "inner_product":
	// 	layer = mkInnerProduct(node.InnerProductParam)
	// case "pooling":
	// 	layer = mkPooling(node.PoolingParam)
	// case "batchnorm", "bn":
	// 	layer = mkBatchNorm(node.BatchNormParam)
	// case "lrn":
	// 	layer = mkLRN(node.LrnParam)
	// case "normalize":
	// case "concat":
	// 	parentsInfo := c.getParentsInfo(node)
	// 	layer = mkConcat(parentsInfo, node.ConcatParam)
	// case "eltwise":
	// 	parentsInfo := c.getParentsInfo(node)
	// 	layer = mkElementWise(parentsInfo, node.EltwiseParam)
	// case "softmax", "softmaxwithloss", "softmax_loss":
	// 	layer = mkSoftMax(node.SoftmaxParam)
	// case "flatten":
	// 	pp.Println("unhandeled", layerType)
	// case "power":
	// 	pp.Println("unhandeled", layerType)
	// case "deconvolution":
	// 	pp.Println("unhandeled", layerType)
	// case "crop":
	// 	pp.Println("unhandeled", layerType)
	// case "scale":
	// 	layer = mkScale(node.ScaleParam)
	// case "implicit":
	// 	pp.Println("unhandeled", layerType)
	// case "accuracy":
	// 	pp.Println("unhandeled", layerType)
	// case "permute":
	// default:
	// 	pp.Println("unhandeled", layerType)

	if layer == nil {
		pp.Println(node)
		return nil
	}
	// layer.SetName(node.Name)

	return layer
}

func mkInput(inputs []string) dllayer.Layer {
	// inputDimensions := dllayer.ToInt64Slice(param.Shape[0].Dim)
	// return &layer.Input{
	// 	N: inputDimensions[0],
	// 	C: inputDimensions[1],
	// 	W: inputDimensions[2],
	// 	H: inputDimensions[3],
	// }
	return nil
}
