package main

import (
	"time"

	"github.com/ghostec/goge/game"
	"github.com/ghostec/goge/gameobject"
	"github.com/ghostec/goge/mesh"
	"github.com/ghostec/goge/scene"
	"github.com/ghostec/goge/three"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	camera := three.NewCamera(50, 1, 0.1, 1000)
	screen := three.NewScreen()
	renderer := three.NewRenderer()
	renderer.SetCamera(camera)
	renderer.SetScreen(screen)
	// TODO: scene should be in scene? (how would the game change it?)
	scene := buildScene()
	c := game.Config{
		MaxFPS:   30,
		Renderer: renderer,
		Scene:    scene,
	}
	g := game.New(c)
	go g.Loop()
	js.Global.Set("goge", map[string]interface{}{
		"Ready": three.Ready,
	})
}

func buildScene() *scene.Scene {
	scene := scene.New()
	box := gameobject.New()
	codeList := gameobject.NewCodeListComponent()
	code := gameobject.NewSimpleCodeComponent("rotate_cube")
	code.SetInit(func(obj *gameobject.GameObject) {
		obj.Transform.Scale.X = 1
		obj.Transform.Scale.Y = 2
		obj.Transform.Scale.Z = 0.5
	})
	code.SetUpdate(func(obj *gameobject.GameObject, elapsed time.Duration) {
		obj.Transform.Rotate.X += 0.01
		obj.Transform.Rotate.Y += 0.01
	})
	codeList.Add(code)
	box.AddComponent(codeList)
	drawable := gameobject.NewDrawableComponent()
	m := mesh.New()
	m.Geometry = mesh.NewBox()
	drawable.Set(m)
	box.AddComponent(drawable)
	root := scene.Graph.Root()
	root.Value = box
	return scene
}
