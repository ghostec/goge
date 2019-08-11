package camera

import "github.com/ghostec/goge/math"

type Camera interface {
	LookingAt() math.Vec3
	Position() math.Vec3
	SetPosition(math.Vec3)
}
