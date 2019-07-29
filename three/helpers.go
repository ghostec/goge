package three

import "github.com/gopherjs/gopherjs/js"

func Log(i interface{}) {
	js.Global.Get("console").Call("log", js.Global.Get("JSON").Call("stringify", i))
}
