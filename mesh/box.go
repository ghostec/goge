package mesh

import "github.com/ghostec/goge/math"

type Box struct {
	*SimpleGeometry
}

// NewBox creates a box with (x, y, z) dimensions (width, height, length)
func NewBox(width, height, length float64) *Box {
	halfWidth, halfHeight, halfLength := width*0.5, height*0.5, length*0.5
	vertices := []Vertex{
		// up plane
		Vertex{Position: math.Vec3{-halfWidth, halfHeight, -halfLength}},
		Vertex{Position: math.Vec3{-halfWidth, halfHeight, halfLength}},
		Vertex{Position: math.Vec3{halfWidth, halfHeight, halfLength}},
		Vertex{Position: math.Vec3{halfWidth, halfHeight, -halfLength}},
		// down plane
		Vertex{Position: math.Vec3{-halfWidth, -halfHeight, -halfLength}},
		Vertex{Position: math.Vec3{-halfWidth, -halfHeight, halfLength}},
		Vertex{Position: math.Vec3{halfWidth, -halfHeight, halfLength}},
		Vertex{Position: math.Vec3{halfWidth, -halfHeight, -halfLength}},
	}
	faces := []Face{
		// top
		Face{Vertices: [3]int64{0, 1, 2}},
		Face{Vertices: [3]int64{0, 2, 3}},
		// right
		Face{Vertices: [3]int64{2, 7, 3}},
		Face{Vertices: [3]int64{2, 6, 7}},
		// bottom
		Face{Vertices: [3]int64{4, 7, 5}},
		Face{Vertices: [3]int64{5, 7, 6}},
		// left
		Face{Vertices: [3]int64{4, 5, 0}},
		Face{Vertices: [3]int64{5, 1, 0}},
		// back
		Face{Vertices: [3]int64{5, 2, 1}},
		Face{Vertices: [3]int64{2, 5, 6}},
		// front
		Face{Vertices: [3]int64{0, 7, 3}},
		Face{Vertices: [3]int64{0, 4, 7}},
	}
	calculateFacesAndVerticesNormals(vertices, faces)
	return &Box{
		SimpleGeometry: &SimpleGeometry{
			vertices: vertices,
			faces:    faces,
		},
	}
}
