package scene

import "github.com/ghostec/goge/scene/graph"

type Scene struct {
	Graph *graph.Graph
}

func New() *Scene {
	return &Scene{
		Graph: graph.NewGraph(),
	}
}
