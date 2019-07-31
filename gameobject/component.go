package gameobject

import (
	"time"
)

type Component interface {
	Get() interface{}
	Set(interface{})
	Type() ComponentType
	Update(obj *GameObject, elapsed time.Duration)
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

func (c *CustomComponent) Update(obj *GameObject, elapsed time.Duration) {
	return
}
