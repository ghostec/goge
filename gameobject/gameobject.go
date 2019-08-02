package gameobject

import (
	"fmt"

	"github.com/google/uuid"
)

type GameObject struct {
	uuid       uuid.UUID
	components map[ComponentType]Component
	Transform
}

func New() *GameObject {
	uuid, _ := uuid.NewRandom()
	return &GameObject{
		uuid:       uuid,
		components: map[ComponentType]Component{},
	}
}

func (o *GameObject) AddComponent(c Component) error {
	if _, ok := o.components[c.Type()]; ok {
		return fmt.Errorf("GameObject %s already has component with type %s", o.uuid, c.Type())
	}
	o.components[c.Type()] = c
	return nil
}

func (o *GameObject) RemoveComponent(c Component) error {
	if _, ok := o.components[c.Type()]; !ok {
		return fmt.Errorf("GameObject %s has no component with type %s", o.uuid, c.Type())
	}
	delete(o.components, c.Type())
	return nil
}

func (o GameObject) GetComponent(ct ComponentType) (Component, bool) {
	c, ok := o.components[ct]
	return c, ok
}
