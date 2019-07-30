package gameobject

type GameObject struct {
	Drawable interface{}
}

func New() *GameObject {
	return &GameObject{}
}
