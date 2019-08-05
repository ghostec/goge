package gameobject

const DrawableComponentType = ComponentType("gameobject.component.Drawable")

func NewDrawableComponent() Component {
	return &CustomComponent{
		ct: DrawableComponentType,
	}
}
