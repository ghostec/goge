package gameobject

func NewDrawableComponent() Component {
	return &CustomComponent{
		ct: DrawableComponentType,
	}
}
