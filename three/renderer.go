package three

import (
	"time"

	"github.com/ghostec/goge/camera"
	"github.com/ghostec/goge/renderer"
	"github.com/ghostec/goge/scene"
	"github.com/gopherjs/gopherjs/js"
)

type Renderer struct {
	it     *js.Object
	camera *Camera
	screen *Screen
	tscene *Scene
	scene  *scene.Scene
}

func NewRenderer() *Renderer {
	it := THREE().Get("WebGLRenderer").New()
	// TODO: Attach instead of doing it on ctor
	js.Global.Get("document").Get("body").Call("appendChild", it.Get("domElement"))
	return &Renderer{
		it:     it,
		tscene: NewScene(),
	}
}

func (r *Renderer) Update(elapsed time.Duration) error {
	r.Render()
	return nil
}

func (r *Renderer) Render() error {
	rot := r.tscene.cube.Get("rotation")
	rot.Set("x", rot.Get("x").Float()+0.01)
	rot.Set("y", rot.Get("y").Float()+0.01)
	r.it.Call("render", r.tscene.it, r.camera.it)
	return nil
}

func (r *Renderer) SetScreen(screen renderer.Screen) {
	r.screen = screen.(*Screen)
	w := screen.Width()
	h := screen.Height()
	r.it.Call("setSize", w, h)
	if r.camera != nil {
		r.camera.SetAspectRatio(w / h)
	}
}

func (r *Renderer) SetCamera(camera camera.Camera) {
	r.camera = camera.(*Camera)
	if r.screen != nil {
		r.camera.SetAspectRatio(r.screen.Width() / r.screen.Height())
	}
}

func (r *Renderer) SetScene(scene *scene.Scene) {
	// TODO: create new scene
	r.scene = scene
}
