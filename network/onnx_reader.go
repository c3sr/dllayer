package network

import (
	"github.com/rai-project/dllayer"
	"github.com/rai-project/onnx"
)

type Onnx struct {
	irVersion        int64
	producerName     string
	model            *onnx.ModelProto
	graph            *onnx.GraphProto
	nodes            map[string]*onnx.NodeProto
	initializers     map[string]*onnx.TensorProto
	layerInformation map[string]dllayer.LayerInfo
}

func NewOnnx(protoFileName string) (*Onnx, error) {
	model, err := onnx.ReadModel(protoFileName)
	if err != nil {
		return nil, err
	}

	graph := model.GetGraph()

	nodes := map[string]*onnx.NodeProto{}
	for _, n := range graph.Node {
		nodes[n.Name] = n
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
