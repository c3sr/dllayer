package layer

import (
  "github.com/rai-project/dllayer"
)

type Data struct {
  Base
  N int64
  C int64
  W int64
  H int64
}

func (Data) Type() string {
  return "Data"
}

func (Data) Aliases() []string {
  return []string{"Data"}
}

func (Data) Description() string {
  return ``
}

func (c *Data) LayerInformation(inputDimensions []int64) dllayer.LayerInfo {
  batchSize := c.N
  batchSize = 1
  return &Information{
    name:             c.name,
    typ:              c.Type(),
    inputDimensions:  []int64{batchSize, c.C, c.W, c.H},
    outputDimensions: []int64{batchSize, c.C, c.W, c.H},
  }
}

func init() {
  dllayer.Register(&Data{})
}
