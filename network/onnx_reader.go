package network

import (
	"github.com/rai-project/dllayer"
	"github.com/rai-project/onnx"
)

type Onnx struct {
	irVersion    string
	producerName string
	model        *onnx.ModelProto
	graph        *onnx.GraphProto
	nodes        map[string]*onnx.NodeProto
	initializers map[string]*onnx.TensorProto
	inputs       map[string]*onnx.ValueInfoProto
	outputs      map[string]*onnx.ValueInfoProto
}

func NewOnnx(protoFileName string) (*Onnx, error) {
	model, err := onnx.ReadModel(protoFileName)
	if err != nil {
		return nil, err
	}

	graph := model.GetGraph()

	nodes := map[string]*onnx.NodeProto{}
	for _, n := range graph.Node {
		nodes[n.name] = n
	}

	initializers := map[string]*onnx.TensorProto{}
	for _, t := range graph.Initializer {
		initializers[t.name] = t
	}

	inputs := map[string]*onnx.ValueInfoProto{}
	for _, i := range graph.Input {
		inputs[i.name] = i
	}

	outputs := map[string]*onnx.ValueInfoProto{}
	for _, o := range graph.Output {
		outputs[o.name] = o
	}

	return &Onnx{
		irVersion: model.GetIrVersion(),
		model,
		graph,
		nodes,
		initializers,
		inputs,
		outputs,
		layerInformation: make(map[string]dllayer.LayerInfo),
	}, nil
}
