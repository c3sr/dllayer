package network

import "github.com/rai-project/caffe"

type Caffe struct {
	*caffe.NetParameter
}

func NewCaffe(protoFileName string) (*Caffe, error) {
	net, err := caffe.ReadNet(protoFileName)
	if err != nil {
		return nil, err
	}
	return &Caffe{net}, nil
}
