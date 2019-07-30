package main

import (
	"github.com/ghostec/goge/game"
	"github.com/ghostec/goge/math"
	"github.com/ghostec/goge/mesh"
	"github.com/ghostec/goge/scene"
	"github.com/ghostec/goge/three"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	camera := three.NewCamera(50, 1, 0.1, 1000)
	screen := three.NewScreen()
	scene := scene.New()
	root := scene.Graph.Root()
	root.Value = mesh.NewBox(math.Vec3{1, 2, 1})
	renderer := three.NewRenderer()
	renderer.SetCamera(camera)
	renderer.SetScreen(screen)
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
