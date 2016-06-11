package qhull

import (
	"testing"

	"github.com/luxengine/lux/glm"
)

func TestBuildInitialTetrahedron(t *testing.T) {
	tests := []struct {
		a, b, c, d int
		points     []glm.Vec3
	}{
		{0, 1, 2, 3,
			[]glm.Vec3{{0, 0, -0.0001}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
		{0, 2, 1, 3,
			[]glm.Vec3{{0, 0, -0.0001}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
		{0, 1, 3, 2,
			[]glm.Vec3{{0, 0, -0.0001}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
		{1, 0, 3, 2,
			[]glm.Vec3{{0, 0, -0.0001}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
	}
	for i, test := range tests {
		var center glm.Vec3
		center.AddWith(&test.points[test.a])
		center.AddWith(&test.points[test.b])
		center.AddWith(&test.points[test.c])
		center.AddWith(&test.points[test.d])
		center.MulWith(0.25)

		tetra := BuildInitialTetrahedron(test.a, test.b, test.c, test.d, test.points, &center)
		for j, face := range tetra {
			if j != 3 {
				continue
			}
			other := 6 - face.Vertices[0] - face.Vertices[1] - face.Vertices[2]

			if dist := DistToPlane(&face.Plane, &test.points[other]); dist > 0 {
				t.Errorf("[%d] vertices [%d, %d, %d]", i, face.Vertices[0], face.Vertices[1], face.Vertices[2])
				t.Errorf("[%d] f[%d] dist = %f", i, j, dist)
			}
		}
	}
}
