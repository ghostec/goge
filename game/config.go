package game

import (
	"github.com/ghostec/goge/renderer"
	"github.com/ghostec/goge/scene"
)

type Config struct {
	MaxFPS   int
	Renderer renderer.Renderer
	Scene    *scene.Scene
}

func (c Config) GetMaxFPS() int {
	if c.MaxFPS <= 0 {
		return 120
	}
	return c.MaxFPS
}
