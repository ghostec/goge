package mesh

import "github.com/ghostec/goge/math"

type Box struct {
	*SimpleGeometry
}

func NewBox() *Box {
	return &Box{
		SimpleGeometry: &SimpleGeometry{
			vertices: []math.Vec3{
				math.Vec3{X: 0.5, Y: 0.5, Z: 0.5},
				math.Vec3{X: 0.5, Y: 0.5, Z: -0.5},
				math.Vec3{X: 0.5, Y: -0.5, Z: 0.5},
				math.Vec3{X: 0.5, Y: -0.5, Z: -0.5},
				math.Vec3{X: -0.5, Y: 0.5, Z: -0.5},
				math.Vec3{X: -0.5, Y: 0.5, Z: 0.5},
				math.Vec3{X: -0.5, Y: -0.5, Z: -0.5},
				math.Vec3{X: -0.5, Y: -0.5, Z: 0.5},
			},
			faces: []math.Vec3{
				math.Vec3{X: 0, Y: 2, Z: 1},
				math.Vec3{X: 2, Y: 3, Z: 1},
				math.Vec3{X: 4, Y: 6, Z: 5},
				math.Vec3{X: 6, Y: 7, Z: 5},
				math.Vec3{X: 4, Y: 5, Z: 1},
				math.Vec3{X: 5, Y: 0, Z: 1},
				math.Vec3{X: 7, Y: 6, Z: 2},
				math.Vec3{X: 6, Y: 3, Z: 2},
				math.Vec3{X: 5, Y: 7, Z: 0},
				math.Vec3{X: 7, Y: 2, Z: 0},
				math.Vec3{X: 1, Y: 3, Z: 4},
				math.Vec3{X: 3, Y: 6, Z: 4},
			},
		},
	}
}
