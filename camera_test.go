package lux

import (
	"errors"
	"fmt"
	glm "github.com/go-gl/mathgl/mgl32"
	"testing"
)

var _ = fmt.Printf

func TestCameraSetPerspective(t *testing.T) {
	cam := Camera{}
	cam.SetPerspective(60, 19.0/9.0, 0.1, 100)
	expected := glm.Mat4{-0.07395156, 0, 0, 0, 0, -0.15611996, 0, 0, 0, 0, -1.002002, -1, 0, 0, -0.2002002, 0}
	if cam.Projection != expected {
		t.Error(errors.New("didnt get expected matrix from SetPerspective"))
	}
}

func TestCameraLookAtVec(t *testing.T) {
	eye, center, up := glm.Vec3{5, 5, 5}, glm.Vec3{0, 0, 0}, glm.Vec3{0, 1, 0}
	cam := Camera{}
	cam.LookAtVec(&eye, &center, &up)
	expected := glm.Mat4{0.70710677, -0.40824825, 0.5773502, 0, 0, 0.8164965, 0.5773502, 0, -0.70710677, -0.40824825, 0.5773502, 0, 0, 0, -8.660253, 1}
	if cam.View != expected {
		t.Error(errors.New("did not get expected value from LookAtVec"))
	}
}

func TestCameraLookAtVal(t *testing.T) {
	eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ := float32(5), float32(5), float32(5), float32(0), float32(0), float32(0), float32(0), float32(1), float32(0)
	cam := Camera{}
	cam.LookAtval(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ)
	expected := glm.Mat4{0.70710677, -0.40824825, 0.5773502, 0, 0, 0.8164965, 0.5773502, 0, -0.70710677, -0.40824825, 0.5773502, 0, 0, 0, -8.660253, 1}
	if cam.View != expected {
		t.Error(errors.New("did not get expected value from LookAtVec"))
	}
}

func TestCameraOrtho(t *testing.T) {
	cam := Camera{}
	cam.SetOrtho(-1, 1, -1, 1, 0.1, 100)
	expected := glm.Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -0.02002002, 0, -0, -0, -1.002002, 1}
	if cam.Projection != expected {
		t.Error(errors.New("unexpected results from SetOrtho"))
	}
}
