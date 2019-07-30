package mesh

import (
	"github.com/ghostec/goge/math"
	"github.com/ghostec/goge/types"
)

const BoxType = types.Type("Mesh.Box")

type Box struct {
	Dimensions math.Vec3
}

func NewBox(dimensions math.Vec3) *Box {
	return &Box{dimensions}
}

func (b Box) Type() types.Type {
	return BoxType
}
