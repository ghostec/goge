package gameobject

import (
	"time"

	"github.com/ghostec/goge/event"
)

type Context struct {
	Dispatcher *event.Dispatcher
	Elapsed    time.Duration
	GameObject *GameObject
}

func NewContext() *Context {
	return &Context{}
}
