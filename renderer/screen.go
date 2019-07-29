package renderer

type Screen interface {
	Width() float64
	Height() float64
	OnResize(callback func()) (deregister func())
}
