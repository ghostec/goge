package three

import "github.com/gopherjs/gopherjs/js"

type Scene struct {
	it *js.Object
}

func NewScene() *Scene {
	s := &Scene{
		it: THREE().Get("Scene").New(),
	}
	return s
}
