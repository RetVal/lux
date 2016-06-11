package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestRayResultAny_AddResult(t *testing.T) {
	var rr RayResultAny
	b := NewRigidBody()
	hit := glm.Vec3{X: 1, Y: 2, Z: 3}

	if rr.AddResult(b, hit) != false {
		t.Error("RayResultAny does not stop after 1 result")
	}

	if rr.Hit != hit {
		t.Error("Incorrect hit stored in RayResultAny")
	}

	if rr.Body != b {
		t.Error("Incorrect body stored in RayResultAny")
	}
}

func TestRayResultClosest_AddResult(t *testing.T) {
	rr := RayResultClosest{
		Origin: glm.Vec3{},
	}

	b := []*RigidBody{
		NewRigidBody(),
		NewRigidBody(),
		NewRigidBody(),
	}
	hit := []glm.Vec3{
		{X: 1, Y: 2, Z: 3},
		{X: 0, Y: 1, Z: 0},
		{X: 4, Y: 5, Z: 6},
	}

	for n := range b {
		rr.AddResult(b[n], hit[n])
	}

	const expected = 1

	if rr.Body != b[expected] {
		t.Errorf("rr.Body = %p, want %p", rr.Body, b[expected])
	}

	if rr.Hit != hit[expected] {
		t.Errorf("rr.Hit = %s, want %s", rr.Hit.String(), hit[expected].String())
	}
}

func TestRayClosestAll_AddResult(t *testing.T) {
	rr := RayResultAll{
		Origin: glm.Vec3{},
	}

	b := []*RigidBody{
		NewRigidBody(),
		NewRigidBody(),
		NewRigidBody(),
	}
	hit := []glm.Vec3{
		{X: 1, Y: 2, Z: 3},
		{X: 0, Y: 1, Z: 0},
		{X: 4, Y: 5, Z: 6},
	}

	for n := range b {
		rr.AddResult(b[n], hit[n])
	}

	if len(rr.Len2s) != 3 || len(rr.Bodies) != 3 || len(rr.Points) != 3 {
		t.Error("Lengths of slices not as expected, cannot continue")
		t.Errorf("Len2s: %d, Bodies: %d, Points: %d", len(rr.Len2s), len(rr.Bodies), len(rr.Points))
		return
	}

	if rr.Len2s[0] > rr.Len2s[1] || rr.Len2s[1] > rr.Len2s[2] {
		t.Errorf("Len2s not in increasing order %v", rr.Len2s)
	}

	if rr.Points[0] != hit[1] || rr.Points[1] != hit[0] || rr.Points[2] != hit[2] {
		t.Errorf("Points not in expected order %v", rr.Points)
	}

	if rr.Bodies[0] != b[1] || rr.Bodies[1] != b[0] || rr.Bodies[2] != b[2] {
		t.Error("Bodies not in expected order")
	}
}
