package dllayer

type Layer interface {
	Name() string
	Aliases() []string
}
