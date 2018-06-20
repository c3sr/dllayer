package network

import (
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

// LayerInformations uses the first input
func (o Onnx) LayerInformations() []dllayer.LayerInfo {
	inputDimensions := []int64{}
	inputs := o.graph.GetInput()
	inputDims := inputs[0].GetType().GetTensorType().GetShape().GetDim()

	if len(inputs) != 0 && inputs[0] != nil && len(inputDims) != 0 {
		inputDimensions = getDimValues(inputDims)
		if len(o.nodes) != 0 {
			return o.layerInformations(inputDimensions)
		}
	} else {
		log.Info("no input info for graph", o.graph.GetName())
	}

	return nil
}

func (o Onnx) layerInformations(inputDimensions []int64) []dllayer.LayerInfo {
	res := []dllayer.LayerInfo{}

	// The nodes in the graph are sorted topologically
	for _, node := range o.graph.GetNode() {
		name := node.GetName()
		layer := o.mkLayer(node)
		if layer == nil {
			pp.Println("failed to create ", node.GetName())
			return nil
		}
		info := layer.LayerInformation(inputDimensions)
		o.layerInformation[name] = info
		res = append(res, info)
	}

	return res
}

func (o Onnx) mkLayer(lyr *onnx.NodeProto) dllayer.Layer {
	var layer dllayer.Layer
	return layer
}
