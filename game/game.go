package game

import (
	"runtime"
	"time"

	"github.com/ghostec/goge/event"
	"github.com/ghostec/goge/gameobject"
	"github.com/ghostec/goge/renderer"
	"github.com/ghostec/goge/scene"
)

type Game struct {
	dispatcher  *event.Dispatcher
	minInterval time.Duration
	renderer    renderer.Renderer
	scene       *scene.Scene
}

func New(config Config) *Game {
	g := &Game{
		dispatcher:  event.NewDispatcher(),
		minInterval: time.Millisecond * time.Duration(float64(1000)/float64(config.GetMaxFPS())),
		renderer:    config.Renderer,
		scene:       config.Scene,
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
		// TODO: update scene graph -> apply transform matrix BFS order
		g.update(elapsed)
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

func (g Game) update(elapsed time.Duration) {
	g.dispatcher.Process()
	bfs := g.scene.Graph.BFS()
	ctx := gameobject.NewContext()
	ctx.Elapsed = elapsed
	ctx.Dispatcher = g.dispatcher
	for _, node := range bfs {
		switch v := node.Value.(type) {
		case *gameobject.GameObject:
			v.Update(ctx)
		}
	}
}

func (g Game) Dispatcher() *event.Dispatcher {
	return g.dispatcher
}

func (g Game) Renderer() renderer.Renderer {
	return g.renderer
}

func (g Game) Scene() *scene.Scene {
	return g.scene
}

func (g *Game) SetScene(s *scene.Scene) {
	g.scene = s
	if g.renderer != nil {
		g.renderer.SetScene(s)
	}
}

func (g *Game) SetRenderer(r renderer.Renderer) {
	// TODO: clean something from previous renderer?
	g.renderer = r
	g.renderer.SetScene(g.scene)
}
