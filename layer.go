package dllayer

type Layer interface {
	Type() string
	Aliases() []string
	Name() string
	SetName(string)
	LayerInformation(inputDimensions []int64) LayerInfo
}

type LayerInfo interface {
	Name() string
	Type() string
	Flops() FlopsInformation
	Memory() MemoryInformation
	InputDimensions() []int64
	OutputDimensions() []int64
}

type FlopsInformation struct {
	MultiplyAdds    int64 `json:"multiply_adds,omitempty"`
	Additions       int64 `json:"additions,omitempty"`
	Divisions       int64 `json:"divisions,omitempty"`
	Exponentiations int64 `json:"exponentiations,omitempty"`
	Comparisons     int64 `json:"comparisons,omitempty"`
	General         int64 `json:"general,omitempty"`
}

func (FlopsInformation) Header() []string {
	return []string{"MultiplyAdds", "Additions", "Divisions", "Exponentiations", "Comparisons", "General"}
}

func (this FlopsInformation) Add(other FlopsInformation) FlopsInformation {
	return FlopsInformation{
		MultiplyAdds:    this.MultiplyAdds + other.MultiplyAdds,
		Additions:       this.Additions + other.Additions,
		Divisions:       this.Divisions + other.Divisions,
		Exponentiations: this.Exponentiations + other.Exponentiations,
		Comparisons:     this.Comparisons + other.Comparisons,
		General:         this.General + other.General,
	}
}

type MemoryInformation struct {
	Weights    int64 `json:"weights,omitempty"`
	Activation int64 `json:"activation,omitempty"`
}

func (this MemoryInformation) Add(other MemoryInformation) MemoryInformation {
	return MemoryInformation{
		Weights:    this.Weights + other.Weights,
		Activation: this.Activation + other.Activation,
	}
}
