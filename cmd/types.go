package cmd

import (
	"fmt"

	"github.com/rai-project/dllayer"
	"github.com/rai-project/utils"
)

type layer struct {
	Name                     string `json:"name"`
	Type                     string `json:"type"`
	dllayer.FlopsInformation `json:",inline"`
}

func (layer) Header() []string {
	base := dllayer.FlopsInformation{}.Header()
	return append([]string{"LayerName", "LayerType"}, base...)
}

func (l layer) Row(humanFlops bool) []string {
	base := l.FlopsInformation.Row(humanFlops)
	return append([]string{l.Name, l.Type}, base...)
}

type netSummary struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

func (netSummary) Header() []string {
	return []string{"Flop Type", "#"}
}

func (l netSummary) Row(humanFlops bool) []string {
	flopsToString := func(e int64) string {
		return fmt.Sprintf("%v", e)
	}
	if humanFlops {
		flopsToString = func(e int64) string {
			return utils.Flops(uint64(e))
		}
	}

	return []string{l.Name, flopsToString(l.Value)}
}
