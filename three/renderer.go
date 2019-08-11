package three

import (
	"time"

	"github.com/ghostec/goge/camera"
	"github.com/ghostec/goge/gameobject"
	"github.com/ghostec/goge/mesh"
	"github.com/ghostec/goge/renderer"
	"github.com/ghostec/goge/scene"
	"github.com/ghostec/goge/scene/graph"
	"github.com/ghostec/goge/store"
	"github.com/gopherjs/gopherjs/js"
)

type Renderer struct {
	it                   *js.Object
	camera               *Camera
	screen               *Screen
	tscene               *Scene
	scene                *scene.Scene
	previousScreenWidth  float64
	previousScreenHeight float64
}

func NewRenderer() *Renderer {
	it := THREE().Get("WebGLRenderer").New(map[string]interface{}{
		"antialias": true,
	})
	return &Renderer{
		it:     it,
		tscene: NewScene(),
	}
}

func (r *Renderer) Update(elapsed time.Duration) error {
	r.refreshScreen()
	r.Render()
	return nil
}

func (r Renderer) Camera() camera.Camera {
	return r.camera
}

func (r *Renderer) Render() error {
	bfs := r.scene.Graph.BFS()
	for _, node := range bfs {
		r.renderNode(node)
	}
	r.it.Call("render", r.tscene.it, r.camera.it)
	return nil
}

func (r *Renderer) SetScreen(screen renderer.Screen) {
	switch screen := screen.(type) {
	case *Screen:
		// TODO: dettach from previous screen
		attachTo := screen.Input().(*js.Object) // DOM element
		attachTo.Call("appendChild", r.it.Get("domElement"))
		r.screen = screen
		r.refreshScreen()
	}
}

func (r *Renderer) refreshScreen() {
	w := r.screen.Width()
	h := r.screen.Height()
	if w == r.previousScreenWidth && h == r.previousScreenHeight {
		return
	}
	r.it.Call("setSize", w, h)
	if r.camera != nil {
		r.camera.SetAspectRatio(w / h)
	}
	r.previousScreenWidth = w
	r.previousScreenHeight = h
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

func (r *Renderer) renderNode(node *graph.Node) {
	switch node.Value.(type) {
	case *gameobject.GameObject:
		r.renderGameObjectNode(node)
	default:
		// unknown value type
		return
	}
}

func (r *Renderer) renderGameObjectNode(node *graph.Node) {
	obj := node.Value.(*gameobject.GameObject)
	drawable, ok := obj.Get(gameobject.DrawableComponentType)
	if !ok {
		return
	}
	switch drawable.Get().(type) {
	case *mesh.Mesh:
		r.renderMeshNode(node)
	default:
		// unknown drawable type
		return
	}
}

const RendererValueKey = store.Key("three.RendererValue")

func (r *Renderer) renderMeshNode(node *graph.Node) {
	obj := node.Value.(*gameobject.GameObject)
	rv := node.Get(RendererValueKey)
	if rv == nil {
		drawable, _ := obj.Get(gameobject.DrawableComponentType)
		m := fromMesh(drawable.Get().(*mesh.Mesh))
		node.Set(RendererValueKey, m)
		r.tscene.it.Call("add", m)
		rv = m
	}
	mesh := rv.(*js.Object)
	rotation := mesh.Get("rotation")
	rotation.Set("x", obj.Transform.Rotate.X)
	rotation.Set("y", obj.Transform.Rotate.Y)
	rotation.Set("z", obj.Transform.Rotate.Z)
	scale := mesh.Get("scale")
	scale.Set("x", obj.Transform.Scale.X)
	scale.Set("y", obj.Transform.Scale.Y)
	scale.Set("z", obj.Transform.Scale.Z)
	position := mesh.Get("position")
	position.Set("x", obj.Transform.Translate.X)
	position.Set("y", obj.Transform.Translate.Y)
	position.Set("z", obj.Transform.Translate.Z)
}
