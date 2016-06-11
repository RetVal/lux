package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"testing"
)

func TestClosestPointPointRect(t *testing.T) {
	tests := []struct {
		point glm.Vec3
		rect  Rect
		out   glm.Vec3
	}{
		{glm.Vec3{X: 0, Y: 2, Z: 0},
			Rect{glm.Vec3{X: 0, Y: 0, Z: 0},
				[2]glm.Vec3{{0, 1, 0}, {1, 0, 0}},
				glm.Vec2{X: 0.5, Y: 0.5}},
			glm.Vec3{X: 0, Y: 0.5, Z: 0}},
		{glm.Vec3{X: 2, Y: 0, Z: 0},
			Rect{glm.Vec3{X: 0, Y: 0, Z: 0},
				[2]glm.Vec3{{0, 1, 0}, {1, 0, 0}},
				glm.Vec2{X: 0.5, Y: 0.5}},
			glm.Vec3{X: 0.5, Y: 0, Z: 0}},
		{glm.Vec3{X: 2, Y: 0, Z: 0},
			Rect{glm.Vec3{X: 0, Y: 0, Z: 0},
				[2]glm.Vec3{{1, 0, 0}, {0, 1, 0}},
				glm.Vec2{X: 0.5, Y: 0.5}},
			glm.Vec3{X: 0.5, Y: 0, Z: 0}},
		{glm.Vec3{X: -100, Y: -100, Z: 0},
			Rect{glm.Vec3{X: 50, Y: 50, Z: 0},
				[2]glm.Vec3{{1, 0, 0}, {0, 1, 0}},
				glm.Vec2{X: 0.5, Y: 0.5}},
			glm.Vec3{X: 49.5, Y: 49.5, Z: 0}},
	}
	for i, test := range tests {
		out := ClosestPointPointRect(&test.point, &test.rect)
		if !glmtesting.Vec3Equal(out, test.out) {
			t.Errorf("[%d] out = %s, want %s", i, out.String(), test.out.String())
		}
	}
}

func TestSqDistPointRect(t *testing.T) {
	tests := []struct {
		point  glm.Vec3
		rect   Rect
		sqdist float32
	}{
		{glm.Vec3{X: 0, Y: 2, Z: 0},
			Rect{glm.Vec3{X: 0, Y: 0, Z: 0},
				[2]glm.Vec3{{1, 0, 0}, {0, 1, 0}},
				glm.Vec2{X: 0.5, Y: 0.5}},
			1.5 * 1.5},
		{glm.Vec3{X: 2, Y: 0, Z: 0},
			Rect{glm.Vec3{X: 0, Y: 0, Z: 0},
				[2]glm.Vec3{{1, 0, 0}, {0, 1, 0}},
				glm.Vec2{X: 0.5, Y: 0.5}},
			1.5 * 1.5},
		{glm.Vec3{X: 1, Y: 1.5, Z: 1},
			Rect{glm.Vec3{X: 0, Y: 1, Z: 0},
				[2]glm.Vec3{{1, 0, 0}, {0, 1, 0}},
				glm.Vec2{X: 0.5, Y: 0.5}},
			0.5 * 0.5},
		{glm.Vec3{X: 24, Y: 24, Z: 0},
			Rect{glm.Vec3{X: 25, Y: 25, Z: 0},
				[2]glm.Vec3{{1, 0, 0}, {0, 1, 0}},
				glm.Vec2{X: 0.5, Y: 0.5}},
			0.5},
	}
	for i, test := range tests {
		sqdist := SqDistPointRect(&test.point, &test.rect)
		if !glmtesting.FloatEqual(sqdist, test.sqdist) {
			t.Errorf("[%d] sqdist = %f, want %f", i, sqdist, test.sqdist)
		}
	}
}
