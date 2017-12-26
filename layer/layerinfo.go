package layer

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
	"github.com/rai-project/dllayer"
)

type Information struct {
	typ              string                    `json:"typ,omitempty"`
	name             string                    `json:"name,omitempty"`
	flops            dllayer.FlopsInformation  `json:"flops,omitempty"`
	memory           dllayer.MemoryInformation `json:"memory,omitempty"`
	inputDimensions  []int64                   `json:"input_dimensions,omitempty"`
	outputDimensions []int64                   `json:"output_dimensions,omitempty"`
}

func (layer *Information) Type() string {
	return layer.typ
}

func (layer *Information) Name() string {
	return layer.name
}

func (layer *Information) Flops() dllayer.FlopsInformation {
	return layer.flops
}
func (layer *Information) Memory() dllayer.MemoryInformation {
	return layer.memory
}
func (layer *Information) InputDimensions() []int64 {
	return layer.inputDimensions
}
func (layer *Information) OutputDimensions() []int64 {
	return layer.outputDimensions
}

func (layer Information) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"type":              layer.Type(),
		"name":              layer.Name(),
		"flops":             layer.Flops(),
		"memory":            layer.Memory(),
		"input_dimensions":  layer.InputDimensions(),
		"output_dimensions": layer.OutputDimensions(),
	}
	return json.Marshal(data)
}

func (layer *Information) UnmarshalJSON(b []byte) error {
	data := map[string]interface{}{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		TagName:  "json",
		Result:   layer,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(data)
}
