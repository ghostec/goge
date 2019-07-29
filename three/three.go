package three

import (
	"github.com/gopherjs/gopherjs/js"
)

var three *js.Object

/*
All types that reference a *js.Object have an `it` member that reference said object
Example:
type Renderer struct {
  // `it` is a reference to a three.WebGLRenderer instance
  it *js.Object
  // other fields
}
*/

// THREE returns the cached three reference to a previously loaded three.js
// present in the html file you're running goge
func THREE() *js.Object {
	if three != nil {
		return three
	}
	three = js.Global.Get("THREE")
	return three
}

// Ready checks if three.js is loaded in the html file goge is running from
func Ready() bool {
	if THREE() == js.Undefined {
		return false
	}
	return true
}
