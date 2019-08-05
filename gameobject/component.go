package gameobject

import (
	"github.com/ghostec/goge/store"
)

type ComponentType store.Key

type Component interface {
	Get() interface{}
	Set(interface{})
	Type() ComponentType
	Update(*Context)
}

type CustomComponent struct {
	value interface{}
	ct    ComponentType
}

func (c CustomComponent) Type() ComponentType {
	return c.ct
}

func (c CustomComponent) Get() interface{} {
	return c.value
}

func (c *CustomComponent) Set(value interface{}) {
	c.value = value
}

func (c *CustomComponent) Update(*Context) {
	return
}
