package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestInertiaTensors_sphereInertiaTensor(t *testing.T) {
	tensor := sphereInertiaTensor(5, 2)
	expected := glm.Mat3{8, 0, 0, 0, 8, 0, 0, 0, 8}

	if !tensor.EqualThreshold(&expected, 1e-4) {
		t.Errorf("Sphere tensor = %v, want %v", tensor, expected)
	}
}

func TestInertiaTensors_cuboidInertiaTensor(t *testing.T) {
	tensor := cuboidInertiaTensor(5, 3, 4, 5)
	expected := glm.Mat3{61.5, 0, 0, 0, 51, 0, 0, 0, 37.5}

	if !tensor.EqualThreshold(&expected, 1e-4) {
		t.Errorf("Cuboid tensor = %v, want %v", tensor, expected)
	}
}
