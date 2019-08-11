package main

import (
	"math/rand"
	"time"

	"github.com/ghostec/goge/event"
	"github.com/ghostec/goge/game"
	"github.com/ghostec/goge/gameobject"
	"github.com/ghostec/goge/math"
	"github.com/ghostec/goge/mesh"
	"github.com/ghostec/goge/scene"
	"github.com/ghostec/goge/three"
	"github.com/gopherjs/gopherjs/js"
)

const (
	AddBoxEvent  = event.Key("add_box")
	ZoomInEvent  = event.Key("zoom_in")
	ZoomOutEvent = event.Key("zoom_out")
)

func main() {
	camera := three.NewCamera(60, 1, 1, 10000)
	camera.SetPosition(math.Vec3{0, 0, 5})
	screen := three.NewScreen(js.Global.Get("document").Call("getElementById", "screen"))
	renderer := three.NewRenderer()
	renderer.SetCamera(camera)
	renderer.SetScreen(screen)
	// TODO: scene should be in scene? (how would the game change it?)
	// scene := buildScene1()
	scene := scene.New() // editor
	c := game.Config{
		MaxFPS:   60,
		Renderer: renderer,
		Scene:    scene,
	}
	g := game.New(c)
	buildEditor(g)
	go g.Loop()
	js.Global.Set("goge", map[string]interface{}{
		"Ready":        three.Ready,
		"Dispatch":     func(key string) { g.Dispatcher().Dispatch(event.New(event.Key(key))) },
		"AddBoxEvent":  AddBoxEvent,
		"ZoomInEvent":  ZoomInEvent,
		"ZoomOutEvent": ZoomOutEvent,
	})
}

func buildEditor(game *game.Game) {
	editorRef := 1
	dispatcher := game.Dispatcher()
	scene := game.Scene()
	renderer := game.Renderer()
	camera := renderer.Camera()
	dispatcher.Subscribe(&editorRef, ZoomInEvent, func(*event.Event) {
		pos := camera.Position()
		at := camera.LookingAt()
		camera.SetPosition(math.Vec3{
			pos.X + at.X*0.5,
			pos.Y + at.Y*0.5,
			pos.Z + at.Z*0.5,
		})
	})
	dispatcher.Subscribe(&editorRef, ZoomOutEvent, func(*event.Event) {
		pos := camera.Position()
		at := camera.LookingAt()
		camera.SetPosition(math.Vec3{
			pos.X - at.X*0.5,
			pos.Y - at.Y*0.5,
			pos.Z - at.Z*0.5,
		})
	})
	dispatcher.Subscribe(&editorRef, AddBoxEvent, func(*event.Event) {
		box := gameobject.New()
		drawable := gameobject.NewDrawableComponent()
		m := mesh.New()
		m.Geometry = mesh.NewBox(1, 1, 1)
		drawable.Set(m)
		box.Set(drawable)
		scene.Graph.Root().NewChild().Value = box
	})
}

func buildScene1() *scene.Scene {
	scene := scene.New()
	box := gameobject.New()
	codeList := gameobject.NewCodeListComponent()
	code := gameobject.NewSimpleCodeComponent()
	someEventKey := event.Key("some_event")
	code.SetInit(func(ctx *gameobject.Context) {
		ctx.Dispatcher.Subscribe(ctx.GameObject, someEventKey, func(e *event.Event) {
			println(e.Key())
		})
		ctx.Dispatcher.Dispatch(event.New(someEventKey))
		ctx.GameObject.Transform.Scale.X = 1
		ctx.GameObject.Transform.Scale.Y = 2
		ctx.GameObject.Transform.Scale.Z = 0.5
	})
	code.SetUpdate(func(ctx *gameobject.Context) {
		ctx.GameObject.Transform.Rotate.X += 0.01
		ctx.GameObject.Transform.Rotate.Y += 0.01
	})
	codeList.Add(code)
	box.Set(codeList)
	drawable := gameobject.NewDrawableComponent()
	m := mesh.New()
	m.Geometry = mesh.NewBox(1, 1, 1)
	drawable.Set(m)
	box.Set(drawable)
	root := scene.Graph.Root()
	root.Value = box
	return scene
}

func buildScene2() *scene.Scene {
	rand.Seed(time.Now().UnixNano())
	scene := scene.New()
	root := scene.Graph.Root()
	for i := 0; i < 250; i++ {
		box := gameobject.New()
		box.Transform.Translate.X = rand.Float64()*30 - 10
		box.Transform.Translate.Y = rand.Float64()*30 - 10
		box.Transform.Translate.Z = rand.Float64()*30 - 10
		box.Transform.Rotate.X = rand.Float64()*2 - 3.1415
		box.Transform.Rotate.Y = rand.Float64()*2 - 3.1415
		drawable := gameobject.NewDrawableComponent()
		m := mesh.New()
		m.Geometry = mesh.NewBox(1, 1, 1)
		drawable.Set(m)
		box.Set(drawable)
		c := root.NewChild()
		c.Value = box
	}
	return scene
}
