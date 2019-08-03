package mesh

import "github.com/ghostec/goge/math"

type Geometry interface {
	Vertices() []math.Vec3
	Faces() []math.Vec3
}

type SimpleGeometry struct {
	vertices []math.Vec3
	faces    []math.Vec3
}

func (g SimpleGeometry) Vertices() []math.Vec3 {
	return g.vertices
}

func (g SimpleGeometry) Faces() []math.Vec3 {
	return g.faces
}
