package three

import (
	"time"

	"github.com/ghostec/goge/camera"
	"github.com/ghostec/goge/gameobject"
	"github.com/ghostec/goge/mesh"
	"github.com/ghostec/goge/renderer"
	"github.com/ghostec/goge/scene"
	"github.com/ghostec/goge/scene/graph"
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
	it := THREE().Get("WebGLRenderer").New()
	// TODO: Attach instead of doing it on ctor
	js.Global.Get("document").Get("body").Call("appendChild", it.Get("domElement"))
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

func (r *Renderer) Render() error {
	bfs := r.scene.Graph.BFS()
	for _, node := range bfs {
		r.renderNode(node)
	}
	r.it.Call("render", r.tscene.it, r.camera.it)
	return nil
}

func (r *Renderer) SetScreen(screen renderer.Screen) {
	r.screen = screen.(*Screen)
	r.refreshScreen()
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
	value := node.Value.(*gameobject.GameObject)
	drawable, ok := value.GetComponent(gameobject.DrawableComponentType)
	if !ok {
		return
	}
	switch drawable.Get().(type) {
	case *mesh.Box:
		r.renderMeshBoxNode(node)
	default:
		// unknown drawable type
		return
	}
}

func (r *Renderer) renderMeshBoxNode(node *graph.Node) {
	if node.RendererValue == nil {
		geometry := THREE().Get("BoxGeometry").New(1, 1, 1)
		material := THREE().Get("MeshBasicMaterial").New(map[string]interface{}{
			"color": 0x00ff00,
		})
		cube := THREE().Get("Mesh").New(geometry, material)
		node.RendererValue = cube
		r.tscene.it.Call("add", cube)
	}
	obj := node.Value.(*gameobject.GameObject)
	mesh := node.RendererValue.(*js.Object)
	rotation := mesh.Get("rotation")
	rotation.Set("x", obj.Transform.Rotate.X)
	rotation.Set("y", obj.Transform.Rotate.Y)
	rotation.Set("z", obj.Transform.Rotate.Z)
	scale := mesh.Get("scale")
	scale.Set("x", obj.Transform.Scale.X)
	scale.Set("y", obj.Transform.Scale.Y)
	scale.Set("z", obj.Transform.Scale.Z)
}
