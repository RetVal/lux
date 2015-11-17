// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/luxengine/glm"
)

// Note: is_near is implemented in mathgl already as FloatEqual

type PackedVertex struct {
	Position glm.Vec3
	UV       glm.Vec2
	Norm     glm.Vec3
}

// Doesn't work like in C++ because Go maps use ==
/*func IndexVBO(vertices []glm.Vec3, uvs []glm.Vec2, normals []glm.Vec3) (outIndices []uint16, outVertices []glm.Vec3, outUVs []glm.Vec2, outNorms []glm.Vec3) {
	vertToOutIndex := make(map[PackedVertex]uint16, 0)

	for i := range vertices {
		packed := PackedVertex{vertices[i], uvs[i], normals[i]}

		index, ok := vertToOutIndex[packed]
		if ok {
			outIndices = append(outIndices, index)
		} else {
			outVertices = append(outVertices, vertices[i])
			outUVs = append(outUVs, uvs[i])
			outNorms = append(outNorms, normals[i])
			index = uint16(len(outVertices) - 1)
			outIndices = append(outIndices, index)
			vertToOutIndex[packed] = index
		}
	}

	return
}*/

func IndexVBOSlow(vertices []glm.Vec3, uvs []glm.Vec2, normals []glm.Vec3) (outIndices []uint16, outVertices []glm.Vec3, outUVs []glm.Vec2, outNorms []glm.Vec3) {

	for i := range vertices {

		index, ok := SimilarVertexIndexSlow(vertices[i], uvs[i], normals[i], outVertices, outUVs, outNorms)
		if ok {
			outIndices = append(outIndices, index)
		} else {
			outVertices = append(outVertices, vertices[i])
			outUVs = append(outUVs, uvs[i])
			outNorms = append(outNorms, normals[i])
			index = uint16(len(outVertices) - 1)
			outIndices = append(outIndices, index)
		}
	}

	return
}

func SimilarVertexIndexSlow(vertex glm.Vec3, uv glm.Vec2, normal glm.Vec3, vertices []glm.Vec3, uvs []glm.Vec2, normals []glm.Vec3) (index uint16, found bool) {
	// Lame linear search
	for i := range vertices {
		if glm.FloatEqualThreshold(vertex[0], vertices[i][0], .01) && glm.FloatEqualThreshold(vertex[1], vertices[i][1], .01) && glm.FloatEqualThreshold(vertex[2], vertices[i][2], .01) &&
			glm.FloatEqualThreshold(uv[0], uvs[i][0], .01) && glm.FloatEqualThreshold(uv[1], uvs[i][1], .01) &&
			glm.FloatEqualThreshold(normal[0], normals[i][0], .01) && glm.FloatEqualThreshold(normal[1], normals[i][1], .01) && glm.FloatEqualThreshold(normal[2], normals[i][2], .01) {
			return uint16(i), true
		}
	}
	// No other vertex could be used instead.
	// Looks like we'll have to add it to the VBO.
	return uint16(0), false
}
