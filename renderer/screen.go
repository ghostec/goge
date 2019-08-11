package renderer

type Screen interface {
	Width() float64
	Height() float64
	// Input returns where to send images to
	// this is usually used by renderer.Renderer implementations
	Input() interface{}
}
