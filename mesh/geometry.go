package mesh

import "github.com/ghostec/goge/math"

type Geometry interface {
	Vertices() []Vertex
	Faces() []Face
}

type Vertex struct {
	Position math.Vec3
	Normal   math.Vec3
}

type Face struct {
	Vertices [3]int64
	Normal   math.Vec3
}

type SimpleGeometry struct {
	vertices []Vertex
	faces    []Face
}

func (g SimpleGeometry) Vertices() []Vertex {
	return g.vertices
}

func (g SimpleGeometry) Faces() []Face {
	return g.faces
}

func calculateFacesAndVerticesNormals(vs []Vertex, fs []Face) {
	vfMap := map[int64][]int64{}
	for i := range vs {
		vfMap[int64(i)] = []int64{}
	}
	for i := range fs {
		fvs := fs[i].Vertices
		vfMap[fvs[0]] = append(vfMap[fvs[0]], int64(i))
		vfMap[fvs[1]] = append(vfMap[fvs[1]], int64(i))
		vfMap[fvs[2]] = append(vfMap[fvs[2]], int64(i))
		fs[i].Normal = math.Normal(vs[fvs[0]].Position, vs[fvs[1]].Position, vs[fvs[2]].Position)
	}
	// TODO: calculate vertices normals with weighted vfMap normals
}
