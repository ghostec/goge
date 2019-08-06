package main

import (
	"math/rand"
	"time"

	"github.com/ghostec/goge/event"
	"github.com/ghostec/goge/game"
	"github.com/ghostec/goge/gameobject"
	"github.com/ghostec/goge/mesh"
	"github.com/ghostec/goge/scene"
	"github.com/ghostec/goge/three"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	camera := three.NewCamera(60, 1, 1, 10000)
	screen := three.NewScreen()
	renderer := three.NewRenderer()
	renderer.SetCamera(camera)
	renderer.SetScreen(screen)
	// TODO: scene should be in scene? (how would the game change it?)
	scene := buildScene1()
	c := game.Config{
		MaxFPS:   60,
		Renderer: renderer,
		Scene:    scene,
	}
	g := game.New(c)
	go g.Loop()
	js.Global.Set("goge", map[string]interface{}{
		"Ready": three.Ready,
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
