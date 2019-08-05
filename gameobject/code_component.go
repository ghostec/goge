package gameobject

import (
	"github.com/ghostec/goge/store"
)

type CodeComponentKey store.Key

type CodeComponent interface {
	Initialized() bool
	Init(*Context)
	Update(*Context)
}

type SimpleCodeComponent struct {
	initialized bool
	initFunc    func(*Context)
	updateFunc  func(*Context)
}

func NewSimpleCodeComponent() *SimpleCodeComponent {
	return &SimpleCodeComponent{}
}

func (c *SimpleCodeComponent) Init(ctx *Context) {
	if c.initFunc != nil {
		c.initFunc(ctx)
	}
	c.initialized = true
}

func (c SimpleCodeComponent) Initialized() bool {
	return c.initialized
}

func (c *SimpleCodeComponent) Update(ctx *Context) {
	if c.updateFunc != nil {
		c.updateFunc(ctx)
	}
}

func (c *SimpleCodeComponent) SetUpdate(f func(*Context)) {
	c.updateFunc = f
}

func (c *SimpleCodeComponent) SetInit(f func(*Context)) {
	c.initFunc = f
}
