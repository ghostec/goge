package gameobject

type Component interface {
	Get() interface{}
	Set(interface{})
	Type() ComponentType
}

type CustomComponent struct {
	value interface{}
	ct    ComponentType
}

func (cc CustomComponent) Type() ComponentType {
	return cc.ct
}

func (cc CustomComponent) Get() interface{} {
	return cc.value
}

func (cc *CustomComponent) Set(value interface{}) {
	cc.value = value
}

func NewDrawableComponent() Component {
	return &CustomComponent{
		ct: DrawableComponentType,
	}
}
