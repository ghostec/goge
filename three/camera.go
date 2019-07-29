package three

import "github.com/gopherjs/gopherjs/js"

type Camera struct {
	it *js.Object
}

// NewCamera ctor
// fov: Camera frustum vertical field of view, from bottom to top of view, in degrees
// ar: Camera frustum aspect ratio, usually the canvas width / canvas height
// near: Camera frustum near plane
// far: Camera frustum far plane
func NewCamera(fov, ar, near, far float64) *Camera {
	c := &Camera{
		it: THREE().Get("PerspectiveCamera").New(fov, ar, near, far),
	}
	c.it.Get("position").Set("z", 5)
	return c
}

func (c *Camera) SetAspectRatio(ar float64) {
	c.it.Set("aspect", ar)
}
