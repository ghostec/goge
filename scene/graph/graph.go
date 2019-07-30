package graph

import (
	"github.com/google/uuid"
)

type Graph struct {
	root *Node
}

func NewGraph() *Graph {
	return &Graph{root: NewNode()}
}

func (g Graph) Root() *Node {
	return g.root
}

func (g Graph) BFS() []*Node {
	result := make([]*Node, 0)
	visited := map[uuid.UUID]bool{}
	queue := make([]*Node, 1)
	queue[0] = g.root
	visited[g.root.UUID()] = true
	for len(queue) > 0 {
		s := queue[0]
		result = append(result, s)
		queue = queue[1:]
		for _, adj := range s.Children() {
			if visited[adj.UUID()] {
				continue
			}
			visited[adj.UUID()] = true
			queue = append(queue, adj)
		}
	}
	return result
}
