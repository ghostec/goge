package gameobject

import (
	"time"
)

type CodeComponent interface {
	Initialized() bool
	Init()
	Update(obj *GameObject, elapsed time.Duration)
}

type SimpleCodeComponent struct {
	name        string
	initialized bool
	initFunc    func()
	updateFunc  func(*GameObject, time.Duration)
}

func NewSimpleCodeComponent(name string) *SimpleCodeComponent {
	return &SimpleCodeComponent{name: name}
}

func (c *SimpleCodeComponent) Init() {
	if c.initFunc != nil {
		c.initFunc()
	}
	c.initialized = true
}

func (c SimpleCodeComponent) Initialized() bool {
	return c.initialized
}

func (c *SimpleCodeComponent) Update(obj *GameObject, elapsed time.Duration) {
	if c.updateFunc != nil {
		c.updateFunc(obj, elapsed)
	}
}

func (c *SimpleCodeComponent) SetUpdate(f func(*GameObject, time.Duration)) {
	c.updateFunc = f
}

func (c *SimpleCodeComponent) SetInit(f func()) {
	c.initFunc = f
}