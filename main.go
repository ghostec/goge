package main

import (
	"github.com/ghostec/goge/game"
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
	c := game.Config{
		MaxFPS:   30,
		Renderer: renderer,
		Scene:    scene.New(),
	}
	g := game.New(c)
	go g.Loop()
	js.Global.Set("goge", map[string]interface{}{
		"Ready":  three.Ready,
		"Render": renderer.Update,
	})
}
