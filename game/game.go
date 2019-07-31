package game

import (
	"runtime"
	"time"

	"github.com/ghostec/goge/gameobject"
	"github.com/ghostec/goge/renderer"
	"github.com/ghostec/goge/scene"
)

type Game struct {
	renderer    renderer.Renderer
	scene       *scene.Scene
	minInterval time.Duration
}

func New(config Config) *Game {
	g := &Game{
		scene:       config.Scene,
		minInterval: time.Millisecond * time.Duration(float64(1000)/float64(config.GetMaxFPS())),
		renderer:    config.Renderer,
	}
	if g.renderer != nil {
		g.renderer.SetScene(g.scene)
	}
	return g
}

func (g *Game) Loop() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	last := time.Now()
	var now, after, minAfter time.Time
	var elapsed, diff time.Duration
	for {
		now = time.Now()
		elapsed = now.Sub(last)
		last = now
		g.updateState(elapsed)
		g.renderer.Update(elapsed)
		// guarantees MaxFPS
		after = time.Now()
		// TODO: measure after.Sub(now)
		// max min avg
		minAfter = last.Add(g.minInterval)
		if after.Before(minAfter) {
			diff = minAfter.Sub(after)
			last = minAfter
			time.Sleep(diff)
		}
	}
}

func (g Game) updateState(elapsed time.Duration) {
	bfs := g.scene.Graph.BFS()
	for _, node := range bfs {
		switch v := node.Value.(type) {
		case *gameobject.GameObject:
			codeList, ok := v.GetComponent(gameobject.CodeListComponentType)
			if !ok {
				continue
			}
			codeList.(*gameobject.CodeListComponent).Update(v, elapsed)
		}
	}
}

func (g *Game) SetRenderer(r renderer.Renderer) {
	// TODO: clean something from previous renderer?
	g.renderer = r
	g.renderer.SetScene(g.scene)
}

func (g Game) Renderer() renderer.Renderer {
	return g.renderer
}
