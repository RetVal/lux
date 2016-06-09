package geo

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestTestCapsuleCapsule(t *testing.T) {
	tests := []struct {
		c0, c1    Capsule
		intersect bool
	}{
		{Capsule{glm.Vec3{-10, 0, 0}, glm.Vec3{10, 0, 0}, 1},
			Capsule{glm.Vec3{0, -10, 0}, glm.Vec3{0, 10, 0}, 1},
			true},
		{Capsule{glm.Vec3{-10, 0, 5}, glm.Vec3{10, 0, 5}, 1},
			Capsule{glm.Vec3{0, -10, 0}, glm.Vec3{0, 10, 0}, 1},
			false},
	}
	for i, test := range tests {
		intersect := TestCapsuleCapsule(&test.c0, &test.c1)
		if intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestTestCapsuleSphere(t *testing.T) {
	tests := []struct {
		c         Capsule
		s         Sphere
		intersect bool
	}{
		{Capsule{glm.Vec3{-10, 0, 0}, glm.Vec3{10, 0, 0}, 1},
			Sphere{glm.Vec3{-5, 0, 0}, 1},
			true},
		{Capsule{glm.Vec3{-10, 0, 5}, glm.Vec3{10, 0, 5}, 1},
			Sphere{glm.Vec3{-5, 0, 0}, 1},
			false},
	}
	for i, test := range tests {
		intersect := TestCapsuleSphere(&test.c, &test.s)
		if intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}
