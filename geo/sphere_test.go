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
				Center: glm.Vec3{0, 0, 0},
				Radius: 1,
			},
			b: Sphere{
				Center: glm.Vec3{0, 0, 0},
				Radius: 1,
			},
			intersects: true,
		},
		{
			a: Sphere{
				Center: glm.Vec3{0, 0, 0},
				Radius: 1,
			},
			b: Sphere{
				Center: glm.Vec3{0, 4, 4},
				Radius: 1,
			},
			intersects: false,
		},

		{
			a: Sphere{
				Center: glm.Vec3{0, 0, 0},
				Radius: 1,
			},
			b: Sphere{
				Center: glm.Vec3{0, 2, 0},
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
				Center: glm.Vec3{3, 4, 7},
				Radius: 1,
			},
			b: AABB{
				Center:     glm.Vec3{3, 4, 7},
				HalfExtend: glm.Vec3{1, 1, 1},
			},
		},
		{
			a: Sphere{
				Center: glm.Vec3{-4, 2.3, 9.9},
				Radius: 1,
			},
			b: AABB{
				Center:     glm.Vec3{-4, 2.3, 9.9},
				HalfExtend: glm.Vec3{1, 1, 1},
			},
		},

		{
			a: Sphere{
				Center: glm.Vec3{1, 4, 5},
				Radius: 1,
			},
			b: AABB{
				Center:     glm.Vec3{1, 4, 5},
				HalfExtend: glm.Vec3{1, 1, 1},
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

		{Sphere{glm.Vec3{0, 0, 0}, 1},
			glm.Vec3{0, 0, 0},
			Sphere{glm.Vec3{}, 1}},
		{Sphere{glm.Vec3{0, 2, 0}, 1},
			glm.Vec3{0, 0, 0},
			Sphere{glm.Vec3{0, 1.5, 0}, 1.5}},
		{Sphere{glm.Vec3{0, 5, 0}, 1},
			glm.Vec3{0, 0, 0},
			Sphere{glm.Vec3{0, 3, 0}, 3}},
		{Sphere{glm.Vec3{0, 0.5, 0}, 1},
			glm.Vec3{0, 0, 0},
			Sphere{glm.Vec3{0, 0.5, 0}, 1}},
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
		{[]glm.Vec3{{0, 0, 0}},
			Sphere{glm.Vec3{}, 0}},
		{[]glm.Vec3{{0, 1, 0}},
			Sphere{glm.Vec3{0, 1, 0}, 0}},
		{[]glm.Vec3{{0, 1, 0}, {0, 0, 0}},
			Sphere{glm.Vec3{0, 0.5, 0}, 0.5}},
		{[]glm.Vec3{{0, 1, 0}, {0, 0, 0}, {0, 2, 0}, {3, 1, 0}, {3, 0, 0}, {3, 2, 0}, {0, 1, 2}, {0, 0, 2}, {0, 2, 2}, {-2, 1, 0}, {-2, 0, 0}, {-2, 2, 0}, {0, 1, -2}, {0, 0, -2}, {0, 2, -2}},
			Sphere{glm.Vec3{0.50364625, 0.9490025, 0}, 2.7152972}},
		{[]glm.Vec3{{0, 1, 0 + 10}, {0, 0, 0 + 10}, {0, 2, 0 + 10}, {3, 1, 0 + 10}, {3, 0, 0 + 10}, {3, 2, 0}, {0, 1, 2}, {0, 0, 2}, {0, 2, 2}, {-2, 1, 0}, {-2, 0, 0}, {-2, 2, 0}, {0, 1, -2}, {0, 0, -2}, {0, 2, -2}},
			Sphere{glm.Vec3{1.5, 1, 4}, 6.264982}},
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
		Center: glm.Vec3{0, 0, 0},
		Radius: 1,
	}
	b := Sphere{
		Center: glm.Vec3{0, 0, 0},
		Radius: 1,
	}

	for n := 0; n < tb.N; n++ {
		TestSphereSphere(&a, &b)
	}
}
