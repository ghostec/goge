package gameobject

import "time"

type Context struct {
	GameObject *GameObject
	Elapsed    time.Duration
}

func NewContext() *Context {
	return &Context{}
}
