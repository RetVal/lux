package terrain

import (
	"errors"

	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/render"
)

// NewTerrain returns a (n-2)(n-2) heightmap. We need the 4 adjacent vertices to accuratelly calculate normals.
func NewTerrain(heightmap [][]float32, scale float32) (lux.Mesh, error) {
	width := len(heightmap) - 2
	height := len(heightmap[0]) - 2
	if width < 2 || height < 2 {
		return nil, errors.New("size too small, cannot calculate normals")
	}
	vertices := make([]glm.Vec3, width*height)
	normals := make([]glm.Vec3, width*height)
	uvs := make([]glm.Vec2, width*height)
	for x := 1; x < width+1; x++ {
		for y := 1; y < height+1; y++ { //technicly z axis
			vertices[(y-1)*width+(x-1)] = glm.Vec3{float32(x-1) * scale, heightmap[x][y], float32(y-1) * scale}
			uvs[(y-1)*width+(x-1)] = glm.Vec2{float32(x-1) / float32(width-1), float32(y-1) / float32(height-1)}
		}
	}
	for x := 1; x < width+1; x++ {
		for y := 1; y < height+1; y++ { //technicly z axis
			n := glm.Vec3{float32(x-1) * scale, heightmap[x][y], float32(y-1) * scale}
			nxm := glm.Vec3{float32(x-2) * scale, heightmap[x-1][y], float32(y-1) * scale}
			nxp := glm.Vec3{float32(x) * scale, heightmap[x+1][y], float32(y-1) * scale}
			nym := glm.Vec3{float32(x-1) * scale, heightmap[x][y-1], float32(y-2) * scale}
			nyp := glm.Vec3{float32(x-1) * scale, heightmap[x][y+1], float32(y) * scale}
			n1 := NormalToPlane(n, nyp, nxp)
			n1.Normalize()
			n2 := NormalToPlane(n, nxp, nym)
			n2.Normalize()
			n3 := NormalToPlane(n, nym, nxm)
			n3.Normalize()
			n4 := NormalToPlane(n, nxm, nyp)
			n4.Normalize()
			normals[(y-1)*width+(x-1)] = AverageVec(n1, n2, n3, n4)
		}
	}

	//indices, freaking complicated to figure out :\
	indices := make([]uint16, (width-1)*(height-1)*6) //*2 triangle per square, *3 indices per triangle
	for x := 0; x < len(indices); x += 3 {
		indices[x] = uint16(x/6 + x/(6*(width-1)) + x%2)
		indices[x+1] = indices[x] + uint16(width-x%2)
		indices[x+2] = indices[x+x%2] + 1
	}
	return lux.NewVUNModel(indices, vertices, uvs, normals), nil
}

// NormalToPlane returns the normal to the triangle defined by the 3 vertices
// given
//TODO: move to a math package
func NormalToPlane(v1, v2, v3 glm.Vec3) glm.Vec3 {
	u := v2.Sub(&v1)
	v := v3.Sub(&v1)
	return glm.Vec3{u.Y()*v.Z() - u.Z()*v.Y(), u.Z()*v.X() - u.X()*v.Z(), u.X()*v.Y() - u.Y()*v.X()}
}

// AverageVec returns the average of the vectors. they must be normalized.
func AverageVec(vecs ...glm.Vec3) glm.Vec3 {
	x, y, z := float32(0), float32(0), float32(0)
	for _, vec := range vecs {
		x += vec.X()
		y += vec.Y()
		z += vec.Z()
	}
	x /= float32(len(vecs))
	y /= float32(len(vecs))
	z /= float32(len(vecs))
	out := glm.Vec3{x, y, z}
	return out.Normalized()
}
