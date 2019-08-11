package three

import "github.com/gopherjs/gopherjs/js"

type Screen struct {
	domElement *js.Object
}

// NewScreen creates a *three.Screen from *js.Object that holds
// reference for a DOM element
// domElement = js.Global.Get("document").Call("getElementById", "some-id")
func NewScreen(domElement *js.Object) *Screen {
	return &Screen{domElement: domElement}
}

func (s Screen) Width() float64 {
	return s.domElement.Get("offsetWidth").Float()
}

func (s Screen) Height() float64 {
	return s.domElement.Get("offsetHeight").Float()
}

func (s Screen) Input() interface{} {
	return s.domElement
}
