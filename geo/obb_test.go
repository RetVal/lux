package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"testing"
)

func TestClosestPointOBBPoint(t *testing.T) {
	tests := []struct {
		obb     OBB
		point   glm.Vec3
		closest glm.Vec3
	}{
		{OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
			glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
			glm.Vec3{X: 1, Y: 1, Z: 1}},
			glm.Vec3{X: 0.3, Y: 0.5, Z: -0.6},
			glm.Vec3{X: 0.3, Y: 0.5, Z: -0.6}},
		{OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
			glm.Mat3{0, 1, 0, 1, 0, 0, 0, 0, 1},
			glm.Vec3{X: 1, Y: 1, Z: 1}},
			glm.Vec3{X: 0.3, Y: 0.5, Z: -0.6},
			glm.Vec3{X: 0.3, Y: 0.5, Z: -0.6}},
		{OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
			glm.Mat3{1, 0, 0, 0, 0, 1, 0, 1, 0},
			glm.Vec3{X: 1, Y: 1, Z: 1}},
			glm.Vec3{X: 0.3, Y: 0.5, Z: -0.6},
			glm.Vec3{X: 0.3, Y: 0.5, Z: -0.6}},
		{OBB{glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Mat3{1, 0, 0, 0, 0, 1, 0, 1, 0},
			glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			glm.Vec3{X: 0, Y: 5, Z: 0},
			glm.Vec3{X: 0, Y: 1.5, Z: 0}},
	}
	for i, test := range tests {
		closest := ClosestPointOBBPoint(&test.obb, &test.point)
		if !glmtesting.Vec3Equal(closest, test.closest) {
			t.Errorf("[%d] closest = %s, want %s", i, closest.String(), test.closest.String())
		}
	}
}

func TestSqDistOBBPoint(t *testing.T) {
	tests := []struct {
		obb   OBB
		point glm.Vec3
		dist2 float32
	}{
		{OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
			glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
			glm.Vec3{X: 1, Y: 1, Z: 1}},
			glm.Vec3{X: 0.3, Y: 0.5, Z: -0.6},
			0},
		{OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
			glm.Mat3{0, 1, 0, 1, 0, 0, 0, 0, 1},
			glm.Vec3{X: 1, Y: 1, Z: 1}},
			glm.Vec3{X: 0.3, Y: 0.5, Z: -0.6},
			0},
		{OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
			glm.Mat3{0, 1, 0, 1, 0, 0, 0, 0, 1},
			glm.Vec3{X: 1, Y: 1, Z: 1}},
			glm.Vec3{X: 0.3, Y: 5.5, Z: -0.6},
			20.25},
	}
	for i, test := range tests {
		dist2 := SqDistOBBPoint(&test.obb, &test.point)
		if !glmtesting.FloatEqual(dist2, test.dist2) {
			t.Errorf("[%d] dist2 = %f, want %f", i, dist2, test.dist2)
		}
	}
}

func TestTestOBBOBB(t *testing.T) {
	tests := []struct {
		a, b      OBB
		intersect bool
	}{
		{
			OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
				glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
				glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true,
		},
		{
			OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
				glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
				glm.Mat3{0, 1, 0, 1, 0, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true,
		},
		{
			OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
				glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
				glm.Mat3{1, 0, 0, 0, 0, 1, 0, 1, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true,
		},
		{
			OBB{glm.Vec3{X: 0, Y: 0, Z: 0},
				glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			OBB{glm.Vec3{X: 3, Y: 0, Z: 0},
				glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			false,
		},
	}
	for i, test := range tests {
		if intersect := TestOBBOBB(&test.a, &test.b); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}
