package three

import "github.com/gopherjs/gopherjs/js"

type Scene struct {
	it   *js.Object
	cube *js.Object
}

func NewScene() *Scene {
	s := &Scene{
		it: THREE().Get("Scene").New(),
	}
	geometry := THREE().Get("BoxGeometry").New(1, 1, 1)
	material := THREE().Get("MeshBasicMaterial").New(map[string]interface{}{
		"color": 0x00ff00,
	})
	s.cube = THREE().Get("Mesh").New(geometry, material)
	s.it.Call("add", s.cube)
	return s
}
