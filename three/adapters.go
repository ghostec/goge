package three

import (
	"github.com/ghostec/goge/math"
	"github.com/ghostec/goge/mesh"
	"github.com/gopherjs/gopherjs/js"
)

func fromMesh(m *mesh.Mesh) *js.Object {
	mvs := m.Geometry.Vertices()
	gvs := make([]interface{}, len(mvs))
	for i := range mvs {
		gvs[i] = fromVertex(mvs[i])
	}
	mfs := m.Geometry.Faces()
	gfs := make([]interface{}, len(mfs))
	for i := range mfs {
		gfs[i] = fromFace(mfs[i])
	}
	geometry := THREE().Get("Geometry").New()
	geometry.Get("vertices").Call("push", gvs...)
	geometry.Get("faces").Call("push", gfs...)
	material := THREE().Get("MeshNormalMaterial").New()
	mjs := THREE().Get("Mesh").New(geometry, material)
	return mjs
}

func fromVertex(v mesh.Vertex) *js.Object {
	p := v.Position
	return THREE().Get("Vector3").New(p.X, p.Y, p.Z)
}

func fromFace(f mesh.Face) *js.Object {
	vs := f.Vertices
	return THREE().Get("Face3").New(vs[0], vs[1], vs[2], fromVec3(f.Normal))
}

func fromVec3(v math.Vec3) *js.Object {
	return THREE().Get("Vector3").New(v.X, v.Y, v.Z)
}
