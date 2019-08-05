package gameobject

import (
	"github.com/ghostec/goge/math"
	"github.com/ghostec/goge/store"
	"github.com/google/uuid"
)

type GameObject struct {
	uuid       uuid.UUID
	components *store.Store
	Transform
}

func New() *GameObject {
	uuid, _ := uuid.NewRandom()
	return &GameObject{
		uuid:       uuid,
		components: store.New(),
		Transform:  Transform{Scale: math.Vec3{1, 1, 1}},
	}
}

func (o *GameObject) Update(ctx *Context) {
	ctx.GameObject = o
	for _, c := range o.components.Values() {
		c.(Component).Update(ctx)
	}
}

func (o GameObject) Get(ct ComponentType) (Component, bool) {
	c := o.components.Get(store.Key(ct))
	if c == nil {
		return nil, false
	}
	return c.(Component), true
}

func (o *GameObject) Set(c Component) {
	o.components.Set(store.Key(c.Type()), c)
}

func (o *GameObject) Unset(ct ComponentType) {
	o.components.Unset(store.Key(ct))
}
