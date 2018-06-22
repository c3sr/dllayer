package onnx

import (
	"github.com/rai-project/dllayer"
	"github.com/rai-project/onnx"
)

type Onnx struct {
	*onnx.ModelProto
	*onnx.GraphProto
	nodesInputs      map[string]*onnx.ValueInfoProto
	nodesOutputs     map[string]*onnx.ValueInfoProto
	values           map[string]*onnx.ValueInfoProto
	initializers     map[string]*onnx.TensorProto
	layerInformation map[string]dllayer.LayerInfo
}

func NewOnnx(protoFileName string) (*Onnx, error) {
	model, err := onnx.ReadModel(protoFileName)
	if err != nil {
		return nil, err
	}

	graph := model.GetGraph()

	nodesInputs := map[string]*onnx.ValueInfoProto{}
	nodesOutputs := map[string]*onnx.ValueInfoProto{}
	for _, n := range graph.Node {
		nodes[n.Name] = n
	}

	values := map[string]*onnx.ValueInfoProto{}
	for _, v := range graph.ValueInfo {
		values[v.Name] = v
	}

	initializers := map[string]*onnx.TensorProto{}
	for _, t := range graph.Initializer {
		initializers[t.Name] = t
	}

	return &Onnx{
		irVersion:        model.GetIrVersion(),
		model:            model,
		graph:            graph,
		nodes:            nodes,
		initializers:     initializers,
		layerInformation: make(map[string]dllayer.LayerInfo),
	}, nil
}
