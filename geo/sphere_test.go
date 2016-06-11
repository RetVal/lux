package geo

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestTestSphereSphere(t *testing.T) {
	tests := []struct {
		a, b       Sphere
		intersects bool
	}{
		{
			a: Sphere{
				Center: glm.Vec3{},
				Radius: 1,
			},
			b: Sphere{
				Center: glm.Vec3{},
				Radius: 1,
			},
			intersects: true,
		},
		{
			a: Sphere{
				Center: glm.Vec3{},
				Radius: 1,
			},
			b: Sphere{
				Center: glm.Vec3{X: 0, Y: 4, Z: 4},
				Radius: 1,
			},
			intersects: false,
		},

		{
			a: Sphere{
				Center: glm.Vec3{},
				Radius: 1,
			},
			b: Sphere{
				Center: glm.Vec3{X: 0, Y: 2, Z: 0},
				Radius: 1,
			},
			intersects: true,
		},
	}

	for i, test := range tests {
		if TestSphereSphere(&test.a, &test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func TestAABBFromSphere(t *testing.T) {
	tests := []struct {
		a Sphere
		b AABB
	}{
		{
			a: Sphere{
				Center: glm.Vec3{X: 3, Y: 4, Z: 7},
				Radius: 1,
			},
			b: AABB{
				Center:     glm.Vec3{X: 3, Y: 4, Z: 7},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
		},
		{
			a: Sphere{
				Center: glm.Vec3{X: -4, Y: 2.3, Z: 9.9},
				Radius: 1,
			},
			b: AABB{
				Center:     glm.Vec3{X: -4, Y: 2.3, Z: 9.9},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
		},

		{
			a: Sphere{
				Center: glm.Vec3{X: 1, Y: 4, Z: 5},
				Radius: 1,
			},
			b: AABB{
				Center:     glm.Vec3{X: 1, Y: 4, Z: 5},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
		},
	}

	for i, test := range tests {
		aabb := AABBFromSphere(&test.a)
		if !aabb.Center.EqualThreshold(&test.b.Center, 1e-4) ||
			!aabb.HalfExtend.EqualThreshold(&test.b.HalfExtend, 1e-4) {
			t.Errorf("[%d] %v.AABB = %v, want %v", i, test.a, aabb, test.b)
		}
	}
}

func TestSphere_MergePoint(t *testing.T) {
	tests := []struct {
		sphere Sphere
		point  glm.Vec3
		next   Sphere
	}{

		{Sphere{glm.Vec3{}, 1},
			glm.Vec3{},
			Sphere{glm.Vec3{}, 1}},
		{Sphere{glm.Vec3{X: 0, Y: 2, Z: 0}, 1},
			glm.Vec3{},
			Sphere{glm.Vec3{X: 0, Y: 1.5, Z: 0}, 1.5}},
		{Sphere{glm.Vec3{X: 0, Y: 5, Z: 0}, 1},
			glm.Vec3{},
			Sphere{glm.Vec3{X: 0, Y: 3, Z: 0}, 3}},
		{Sphere{glm.Vec3{X: 0, Y: 0.5, Z: 0}, 1},
			glm.Vec3{},
			Sphere{glm.Vec3{X: 0, Y: 0.5, Z: 0}, 1}},
	}
	for i, test := range tests {
		s := test.sphere
		s.MergePoint(&test.point)
		if s != test.next {
			t.Errorf("[%d] sphere = %+v, want %+v", i, s, test.next)
		}
	}
}

func TestRitterEigenSphere(t *testing.T) {
	tests := []struct {
		points []glm.Vec3
		sphere Sphere
	}{
		{[]glm.Vec3{{}},
			Sphere{glm.Vec3{}, 0}},
		{[]glm.Vec3{{0, 1, 0}},
			Sphere{glm.Vec3{X: 0, Y: 1, Z: 0}, 0}},
		{[]glm.Vec3{{0, 1, 0}, {}},
			Sphere{glm.Vec3{X: 0, Y: 0.5, Z: 0}, 0.5}},
		{[]glm.Vec3{{0, 1, 0}, {}, {0, 2, 0}, {3, 1, 0}, {3, 0, 0}, {3, 2, 0}, {0, 1, 2}, {0, 0, 2}, {0, 2, 2}, {-2, 1, 0}, {-2, 0, 0}, {-2, 2, 0}, {0, 1, -2}, {0, 0, -2}, {0, 2, -2}},
			Sphere{glm.Vec3{X: 0.50364625, Y: 0.9490025, Z: 0}, 2.7152972}},
		{[]glm.Vec3{{0, 1, 0 + 10}, {0, 0, 0 + 10}, {0, 2, 0 + 10}, {3, 1, 0 + 10}, {3, 0, 0 + 10}, {3, 2, 0}, {0, 1, 2}, {0, 0, 2}, {0, 2, 2}, {-2, 1, 0}, {-2, 0, 0}, {-2, 2, 0}, {0, 1, -2}, {0, 0, -2}, {0, 2, -2}},
			Sphere{glm.Vec3{X: 1.5, Y: 1, Z: 4}, 6.264982}},
	}
	for i, test := range tests {
		s := RitterEigenSphere(test.points)
		if s != test.sphere {
			t.Errorf("[%d] sphere = %+v, want %+v", i, s, test.sphere)
		}

	}
}

func BenchmarkTestSphereSphere(tb *testing.B) {
	a := Sphere{
		Center: glm.Vec3{},
		Radius: 1,
	}
	b := Sphere{
		Center: glm.Vec3{},
		Radius: 1,
	}

	for n := 0; n < tb.N; n++ {
		TestSphereSphere(&a, &b)
	}
}
