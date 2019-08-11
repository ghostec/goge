package three

import (
	"github.com/ghostec/goge/math"
	"github.com/gopherjs/gopherjs/js"
)

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
	return c
}

func (c *Camera) SetAspectRatio(ar float64) {
	c.it.Set("aspect", ar)
}

func (c *Camera) SetPosition(p math.Vec3) {
	c.it.Get("position").Set("x", p.X)
	c.it.Get("position").Set("y", p.Y)
	c.it.Get("position").Set("z", p.Z)
}

func (c Camera) Position() math.Vec3 {
	p := c.it.Get("position")
	return math.Vec3{p.Get("x").Float(), p.Get("y").Float(), p.Get("z").Float()}
}

func (c Camera) LookingAt() math.Vec3 {
	at := THREE().Get("Vector3").New()
	c.it.Call("getWorldDirection", at)
	return math.Vec3{at.Get("x").Float(), at.Get("y").Float(), at.Get("z").Float()}
}
