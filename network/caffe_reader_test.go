package network

import (
	"path/filepath"
	"testing"

	"github.com/GeertJohan/go-sourcepath"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

func TestCaffeReader(t *testing.T) {
	cwd := sourcepath.MustAbsoluteDir()
	resnet101ProtoTxtPath := filepath.Join(cwd, "..", "assets", "model_graph", "squeezenet_1.1.prototxt")
	net, err := NewCaffe(resnet101ProtoTxtPath)
	assert.NoError(t, err)
	assert.NotEmpty(t, net)

	info := net.FlopsInformation()
	pp.Println(info)
	// pp.Println(net.LayerInformations())

	// net.Layer = nil
	// net.Layers = nil
	// pp.Println(net)
}
