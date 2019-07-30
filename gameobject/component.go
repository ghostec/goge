package gameobject

type Component interface {
	Get() interface{}
	Set(interface{})
	Type() ComponentType
	GameObject() *GameObject
	SetGameObject(*GameObject)
}

type CustomComponent struct {
	value interface{}
	ct    ComponentType
	obj   *GameObject
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

func (cc *CustomComponent) SetGameObject(obj *GameObject) {
	cc.obj = obj
}

func (cc CustomComponent) GameObject() *GameObject {
	return cc.obj
}
