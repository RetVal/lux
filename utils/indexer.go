// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/luxengine/glm"
)

// Is this even needed
type packedVertex struct {
	Position glm.Vec3
	UV       glm.Vec2
	Norm     glm.Vec3
}

// IndexVBOSlow builds the index of the given attributes.
func IndexVBOSlow(vertices []glm.Vec3, uvs []glm.Vec2, normals []glm.Vec3) (outIndices []uint16, outVertices []glm.Vec3, outUVs []glm.Vec2, outNorms []glm.Vec3) {
	for i := range vertices {

		index, ok := similarVertexIndexSlow(vertices[i], uvs[i], normals[i], outVertices, outUVs, outNorms)
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

func similarVertexIndexSlow(vertex glm.Vec3, uv glm.Vec2, normal glm.Vec3, vertices []glm.Vec3, uvs []glm.Vec2, normals []glm.Vec3) (index uint16, found bool) {
	// Lame linear search
	for i := range vertices {
		if glm.FloatEqualThreshold(vertex.X, vertices[i].X, .01) && glm.FloatEqualThreshold(vertex.Y, vertices[i].Y, .01) && glm.FloatEqualThreshold(vertex.Z, vertices[i].Z, .01) &&
			glm.FloatEqualThreshold(uv.X, uvs[i].X, .01) && glm.FloatEqualThreshold(uv.Y, uvs[i].Y, .01) &&
			glm.FloatEqualThreshold(normal.X, normals[i].X, .01) && glm.FloatEqualThreshold(normal.Y, normals[i].Y, .01) && glm.FloatEqualThreshold(normal.Z, normals[i].Z, .01) {
			return uint16(i), true
		}
	}
	// No other vertex could be used instead.
	// Looks like we'll have to add it to the VBO.
	return uint16(0), false
}
