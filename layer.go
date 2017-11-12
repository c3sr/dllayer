package dllayer

type Layer interface {
	Type() string
	Aliases() []string
	Name() string
	SetName(string)
	LayerInformation(inputDimensions []int64) LayerInfo
}

type LayerInfo interface {
	Flops() FlopsInformation
	Memory() MemoryInformation
	InputDimensions() []int64
	OutputDimensions() []int64
}

type FlopsInformation struct {
	MultiplyAdds    int64
	Additions       int64
	Divisions       int64
	Exponentiations int64
	Comparisons     int64
	General         int64
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
	Weights    int64
	Activation int64
}

func (this MemoryInformation) Add(other MemoryInformation) MemoryInformation {
	return MemoryInformation{
		Weights:    this.Weights + other.Weights,
		Activation: this.Activation + other.Activation,
	}
}
