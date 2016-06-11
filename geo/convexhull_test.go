package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"testing"
)

func TestConvexhull_Support(t *testing.T) {
	tests := []struct {
		points    []glm.Vec3
		direction glm.Vec3
		support   glm.Vec3
	}{
		{
			[]glm.Vec3{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 0, Y: 1, Z: 0},
		},
		{
			[]glm.Vec3{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			glm.Vec3{X: -1, Y: -1, Z: -1},
			glm.Vec3{X: 0, Y: 0, Z: 0},
		},
		{
			[]glm.Vec3{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			glm.Vec3{X: 1, Y: 0, Z: 0},
			glm.Vec3{X: 1, Y: 0, Z: 0},
		},
		{
			[]glm.Vec3{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			glm.Vec3{X: 0, Y: 0, Z: 1},
			glm.Vec3{X: 0, Y: 0, Z: 1},
		},
	}
	for i, test := range tests {
		hull := Quickhull(test.points)
		for _, tri := range hull.Triangles {
			t.Log(tri.Adjacent)
		}
		supportslow := hull.supportSlow(&test.direction)
		if !glmtesting.Vec3Equal(supportslow, test.support) {
			t.Errorf("[%d] supportslow = %s, want %s", i, supportslow.String(), test.support.String())
		}
		if support, _ := hull.Support(&test.direction, nil); !glmtesting.Vec3Equal(support, supportslow) {
			t.Errorf("[%d] support = %s, want %s", i, support.String(), supportslow.String())
		}
	}
}

func TestTestConvexhullConvexhull(t *testing.T) {
	c0, c1 := Quickhull(suzannePointCloud), Quickhull(suzannePointCloud)
	if !TestConvexhullConvexhull(c0, c1) {
		t.Errorf("A convex hull does not intersect itself.")
	}

	tests := []struct {
		p0, p1    []glm.Vec3
		intersect bool
	}{
		{
			[]glm.Vec3{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[]glm.Vec3{{0 + 0.1, 0 + 0.1, 0 + 0.1}, {1 + 0.1, 0 + 0.1, 0 + 0.1}, {0 + 0.1, 1 + 0.1, 0 + 0.1}, {0 + 0.1, 0 + 0.1, 1 + 0.1}},
			true,
		},
		{
			[]glm.Vec3{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			[]glm.Vec3{{0 + 5, 0 + 5, 0 + 5}, {1 + 5, 0 + 5, 0 + 5}, {0 + 5, 1 + 5, 0 + 5}, {0 + 5, 0 + 5, 1 + 5}},
			false,
		},
	}
	for i, test := range tests {
		c0, c1 := Quickhull(test.p0), Quickhull(test.p1)
		if intersect := testConvexhullConvexhullSlow(c0, c1); intersect != test.intersect {
			t.Errorf("[%d] intersect slow = %t, want %t", i, intersect, test.intersect)
		}

		if intersect := TestConvexhullConvexhull(c0, c1); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestConvexhull_CalculateInternals(t *testing.T) {
	tests := []struct {
		points []glm.Vec3
		volume float32
		center glm.Vec3
	}{
		{[]glm.Vec3{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}, {1, 0, 0},
			{0, 1, 1}, {0, 1, 0}, {0, 0, 1}, {0, 0, 0}},
			1,
			glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
		{[]glm.Vec3{{2, 2, 2}, {2, 2, 0}, {2, 0, 2}, {2, 0, 0},
			{0, 2, 2}, {0, 2, 0}, {0, 0, 2}, {0, 0, 0}},
			8,
			glm.Vec3{X: 1, Y: 1, Z: 1}},
		{[]glm.Vec3{{2, 2, 2}, {2, 2, 1}, {2, 1, 2}, {2, 1, 1},
			{1, 2, 2}, {1, 2, 1}, {1, 1, 2}, {1, 1, 1}},
			1,
			glm.Vec3{X: 1.5, Y: 1.5, Z: 1.5}},
	}
	for i, test := range tests {
		hull := Quickhull(test.points)
		hull.CalculateInternals()
		if !glmtesting.FloatEqual(hull.Volume(), test.volume) {
			t.Errorf("[%d] volume = %f, want %f", i, hull.Volume(), test.volume)
		}
		if !glmtesting.Vec3Equal(hull.Center, test.center) {
			t.Errorf("[%d] center = %s, want %s", i, hull.Center.String(), test.center.String())
		}
	}
}

func TestConvexhull_Inertia(t *testing.T) {
	t.Skip("My interpretation of inertia tensors may be wrong")
	tests := []struct {
		points []glm.Vec3
	}{
		{[]glm.Vec3{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}, {1, 0, 0},
			{0, 1, 1}, {0, 1, 0}, {0, 0, 1}, {0, 0, 0}}},
		{[]glm.Vec3{{2, 2, 2}, {2, 2, 0}, {2, 0, 2}, {2, 0, 0},
			{0, 2, 2}, {0, 2, 0}, {0, 0, 2}, {0, 0, 0}}},
		{[]glm.Vec3{{2, 2, 2}, {2, 2, 1}, {2, 1, 2}, {2, 1, 1},
			{1, 2, 2}, {1, 2, 1}, {1, 1, 2}, {1, 1, 1}}},
	}
	for i, test := range tests {
		points := make([]glm.Vec3, len(test.points))
		for n, p := range test.points {
			points[n] = p
			points[n].X++
		}

		hull := Quickhull(test.points)
		hull.CalculateInternals()

		hull2 := Quickhull(points)
		hull2.CalculateInternals()

		if hull.inertia != hull2.inertia {
			t.Errorf("[%d] not equal \n%s\n%s", i, hull.inertia.String(), hull2.inertia.String())
		}
	}
}

func BenchmarkTestConvexhullConvexhull(b *testing.B) {
	c0, c1 := Quickhull(suzannePointCloud), Quickhull(suzannePointCloud)

	for n := 0; n < b.N; n++ {
		TestConvexhullConvexhull(c0, c1)
	}
}

func BenchmarkTestConvexhullConvexhullSlow(b *testing.B) {
	c0, c1 := Quickhull(suzannePointCloud), Quickhull(suzannePointCloud)

	for n := 0; n < b.N; n++ {
		testConvexhullConvexhullSlow(c0, c1)
	}
}
