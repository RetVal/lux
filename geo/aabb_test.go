package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"testing"
)

func TestAABB_Intersects(t *testing.T) {
	tests := []struct {
		a, b       AABB
		intersects bool
	}{
		{ // 0
			a: AABB{
				Center:     glm.Vec3{X: 0, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			b: AABB{
				Center:     glm.Vec3{X: 0.5, Y: 0.5, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			intersects: true,
		},
		{ // 1
			a: AABB{
				Center:     glm.Vec3{X: 0, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			b: AABB{
				Center:     glm.Vec3{X: 5, Y: 5, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			intersects: false,
		},
		{ // 2
			a: AABB{
				Center:     glm.Vec3{X: 5, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			b: AABB{
				Center:     glm.Vec3{X: 7, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 0, Z: 1},
			},
			intersects: true,
		},
		{ // 3
			a: AABB{
				Center:     glm.Vec3{X: 0, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			b: AABB{
				Center:     glm.Vec3{X: 0, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			intersects: true,
		},
		{ // 4
			a: AABB{
				Center:     glm.Vec3{X: 0, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			b: AABB{
				Center:     glm.Vec3{X: 2, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			intersects: true,
		},

		{ // 5
			a: AABB{
				Center:     glm.Vec3{X: 0, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			b: AABB{
				Center:     glm.Vec3{X: 0, Y: 6, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			intersects: false,
		},

		{ // 6
			a: AABB{
				Center:     glm.Vec3{X: 0, Y: 0, Z: 1},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			b: AABB{
				Center:     glm.Vec3{X: 0, Y: 0, Z: 7},
				HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			intersects: false,
		},
	}

	for i, test := range tests {
		if TestAABBAABB(&test.a, &test.b) != test.intersects {
			t.Errorf("[%d] Intersection test failed %v %v", i, test.a, test.b)
		}
	}
}

func TestClosestPointPointAABB(t *testing.T) {
	tests := []struct {
		aabb    AABB
		point   glm.Vec3
		closest glm.Vec3
	}{
		{AABB{glm.Vec3{},
			glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 0, Y: 0.5, Z: 0}},
	}
	for i, test := range tests {
		closest := ClosestPointPointAABB(&test.point, &test.aabb)
		if !glmtesting.Vec3Equal(closest, test.closest) {
			t.Errorf("[%d] closest = %s, want %s", i, closest.String(), test.closest.String())
		}
	}
}

func TestUpdateAABB(t *testing.T) {
	tests := []struct {
		aabb      AABB
		transform glm.Mat3x4
		fill      AABB
	}{
		{AABB{glm.Vec3{},
			glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			glm.Mat3x4{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
				0, 0, 0,
			},
			AABB{glm.Vec3{},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}}},
		{AABB{glm.Vec3{},
			glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			glm.Mat3x4{
				0, 1, 0,
				1, 0, 0,
				0, 0, 1,
				0, 0, 0,
			},
			AABB{glm.Vec3{},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}}},
		{AABB{glm.Vec3{},
			glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			glm.Mat3x4{
				1, 0, 0,
				0, 0, 1,
				0, 1, 0,
				0, 0, 0,
			},
			AABB{glm.Vec3{},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}}},
		{AABB{glm.Vec3{},
			glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			glm.Mat3x4{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
				0, 0, 0,
			},
			AABB{glm.Vec3{},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}}},
		{AABB{glm.Vec3{},
			glm.Vec3{X: 0.5, Y: 1, Z: 0.5}},
			glm.Mat3x4{
				0, 1, 0,
				1, 0, 0,
				0, 0, 1,
				0, 0, 0,
			},
			AABB{glm.Vec3{},
				glm.Vec3{X: 1, Y: 0.5, Z: 0.5}}},
		{AABB{glm.Vec3{},
			glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			glm.Mat3x4{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
				0, 0, 5,
			},
			AABB{glm.Vec3{X: 0, Y: 0, Z: 5},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}}},
	}
	for i, test := range tests {
		var fill AABB
		UpdateAABB(&test.aabb, &fill, &test.transform)
		if fill != test.fill {
			t.Errorf("[%d] fill = %+v, want %+v", i, fill, test.fill)
		}
	}
}

func TestSqDistPointAABB(t *testing.T) {
	tests := []struct {
		point  glm.Vec3
		aabb   AABB
		sqdist float32
	}{
		{glm.Vec3{X: 0, Y: 2, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			1.5 * 1.5},
		{glm.Vec3{X: 0, Y: -2, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			1.5 * 1.5},
		{glm.Vec3{X: -2, Y: 0, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			1.5 * 1.5},
		{glm.Vec3{X: 2, Y: 0, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			1.5 * 1.5},
		{glm.Vec3{X: 0, Y: 0, Z: -2},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			1.5 * 1.5},
		{glm.Vec3{X: 0, Y: 0, Z: 2},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			1.5 * 1.5},

		{glm.Vec3{X: 0, Y: 0, Z: 0.4},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			0},
		{glm.Vec3{X: 0, Y: 0.4, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			0},
		{glm.Vec3{X: 0.4, Y: 0, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			0},
	}
	for i, test := range tests {
		sqdist := SqDistPointAABB(&test.point, &test.aabb)
		if !glmtesting.FloatEqual(sqdist, test.sqdist) {
			t.Errorf("[%d] sqdist = %f, want %f", i, sqdist, test.sqdist)
		}

	}
}

func BenchmarkTestAABBAABB(tb *testing.B) {
	a := AABB{
		Center:     glm.Vec3{},
		HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
	}
	b := AABB{
		Center:     glm.Vec3{},
		HalfExtend: glm.Vec3{X: 1, Y: 1, Z: 1},
	}
	for n := 0; n < tb.N; n++ {
		TestAABBAABB(&a, &b)
	}
}
