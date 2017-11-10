package layer

import "github.com/rai-project/dllayer"

type Convolution struct {
	Dimensions []int
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

func init() {
	dllayer.Register(Convolution{})
}
