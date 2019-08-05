package graph

import (
	"fmt"

	"github.com/ghostec/goge/store"
	"github.com/google/uuid"
)

type Node struct {
	uuid     uuid.UUID
	children map[uuid.UUID]*Node
	Value    interface{}
	*store.Store
	// TODO: add transform matrix and update it when BFS runs
}

func NewNode() *Node {
	u, _ := uuid.NewRandom()
	return &Node{
		uuid:     u,
		children: map[uuid.UUID]*Node{},
		Store:    store.New(),
	}
}

func (n Node) UUID() uuid.UUID {
	return n.uuid
}

func (n *Node) NewChild() *Node {
	child := NewNode()
	n.children[child.UUID()] = child
	return child
}

func (n *Node) RemoveChild(child *Node) error {
	if _, ok := n.children[child.UUID()]; !ok {
		return fmt.Errorf("%s is not a child of %s", child.uuid.String(), n.uuid.String())
	}
	delete(n.children, child.UUID())
	return nil
}

func (n Node) Children() []*Node {
	children := make([]*Node, 0, len(n.children))
	for _, v := range n.children {
		children = append(children, v)
	}
	return children
}
