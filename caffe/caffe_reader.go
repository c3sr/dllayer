package caffe

import (
	"github.com/rai-project/caffe"
	"github.com/rai-project/dllayer"
)

type Caffe struct {
	*caffe.NetParameter
	layers           map[string]*caffe.LayerParameter
	v1layers         map[string]*caffe.V1LayerParameter
	layerInformation map[string]dllayer.LayerInfo
}

func NewCaffe(protoFileName string) (*Caffe, error) {
	net, err := caffe.ReadNet(protoFileName)
	if err != nil {
		return nil, err
	}

	layers := map[string]*caffe.LayerParameter{}
	for _, lyr := range net.Layer {
		layers[lyr.Name] = lyr
	}
	v1layers := map[string]*caffe.V1LayerParameter{}
	for _, lyr := range net.Layers {
		v1layers[lyr.Name] = lyr
	}
	return &Caffe{
		NetParameter:     net,
		layers:           layers,
		v1layers:         v1layers,
		layerInformation: make(map[string]dllayer.LayerInfo),
	}, nil
}
