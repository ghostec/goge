package three

import "github.com/gopherjs/gopherjs/js"

type Screen struct {
	onResizeCallbacks map[string]func()
}

func NewScreen() *Screen {
	return &Screen{
		onResizeCallbacks: map[string]func(){},
	}
}

func (s Screen) Width() float64 {
	return js.Global.Get("innerWidth").Float()
}

func (s Screen) Height() float64 {
	return js.Global.Get("innerHeight").Float()
}

func (s Screen) OnResize(callback func()) func() {
	return func() {}
}
