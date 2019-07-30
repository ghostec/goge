package three

import "github.com/gopherjs/gopherjs/js"

type Screen struct{}

func NewScreen() *Screen {
	return &Screen{}
}

func (s Screen) Width() float64 {
	return js.Global.Get("innerWidth").Float()
}

func (s Screen) Height() float64 {
	return js.Global.Get("innerHeight").Float()
}
