package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"github.com/luxengine/lux/math"
	"testing"
)

func TestPlaneFromPoints(t *testing.T) {
	tests := []struct {
		a, b, c glm.Vec3
		plane   Plane
	}{
		{glmtesting.NaN3, glm.Vec3{}, glm.Vec3{},
			Plane{glmtesting.NaN3, math.NaN()}},
		{glm.Vec3{}, glmtesting.NaN3, glm.Vec3{},
			Plane{glmtesting.NaN3, math.NaN()}},
		{glm.Vec3{}, glm.Vec3{}, glmtesting.NaN3,
			Plane{glmtesting.NaN3, math.NaN()}},
		{glm.Vec3{0, 1, 0}, glm.Vec3{0, 1, 1}, glm.Vec3{1, 1, 0},
			Plane{glm.Vec3{0, 1, 0}, 1}},
		{glm.Vec3{0, 1, 0}, glm.Vec3{1, 1, 0}, glm.Vec3{0, 1, 1},
			Plane{glm.Vec3{0, -1, 0}, -1}},
		{glm.Vec3{0, -1, 0}, glm.Vec3{1, -1, 0}, glm.Vec3{0, -1, 1},
			Plane{glm.Vec3{0, -1, 0}, 1}},
		{glm.Vec3{0, -1, 0}, glm.Vec3{0, -1, 1}, glm.Vec3{1, -1, 0},
			Plane{glm.Vec3{0, 1, 0}, -1}},
		{glm.Vec3{1, 0, 0}, glm.Vec3{0, 1, 0}, glm.Vec3{0, 0, 1},
			Plane{glm.Vec3{0.577350, 0.577350, 0.577350}, 0.577350}},
	}
	for i, test := range tests {
		plane := PlaneFromPoints(&test.a, &test.b, &test.c)
		if !glmtesting.FloatEqual(plane.Offset, test.plane.Offset) {
			t.Errorf("[%d] plane.Offset = %f, want %f", i, plane.Offset, test.plane.Offset)
		}
		if !glmtesting.Vec3Equal(plane.Normal, test.plane.Normal) {
			t.Errorf("[%d] plane.Normal = %s, want %s", i, plane.Normal.String(), test.plane.Normal.String())
		}
	}
}

func TestDistanceToPlane(t *testing.T) {
	tests := []struct {
		plane Plane
		point glm.Vec3
		dist  float32
	}{
		{PlaneFromPoints(&glm.Vec3{0, 1, 0}, &glm.Vec3{0, 1, 1}, &glm.Vec3{1, 1, 0}),
			glm.Vec3{0, 2, 0},
			1},
		{Plane{glm.Vec3{0, -1, 0}, 1},
			glm.Vec3{0, 0, 0},
			-1},
	}
	for i, test := range tests {
		if dist := DistanceToPlane(&test.plane, &test.point); !glmtesting.FloatEqual(dist, test.dist) {
			t.Errorf("[%d] plane = %+v point = %+v", i, test.plane, test.point)
			t.Errorf("[%d] dist = %f, want %f", i, dist, test.dist)
		}
	}
}
