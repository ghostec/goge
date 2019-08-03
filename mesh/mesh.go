package mesh

type Mesh struct {
	Geometry Geometry
}

func New() *Mesh {
	return &Mesh{}
}
