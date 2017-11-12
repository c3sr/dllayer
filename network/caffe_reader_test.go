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
	resnet101ProtoTxtPath := filepath.Join(cwd, "..", "assets", "model_graph", "alexnet.prototxt")
	net, err := NewCaffe(resnet101ProtoTxtPath)
	assert.NoError(t, err)
	assert.NotEmpty(t, net)

	pp.Println(net.LayerInformation())

	// net.Layer = nil
	// net.Layers = nil
	// pp.Println(net)
}
