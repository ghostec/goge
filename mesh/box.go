package mesh

import (
	"github.com/ghostec/goge/math"
)

type Box struct {
	Dimensions math.Vec3
	Rotate     math.Vec3
}

func NewBox(dimensions math.Vec3) *Box {
	return &Box{Dimensions: dimensions}
}
