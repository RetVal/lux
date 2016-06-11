package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestGravityForceGenerator_New(t *testing.T) {
	gravity := glm.Vec3{X: 0, Y: -10, Z: 0}
	g := NewGravityForceGenerator(&gravity)
	if g.gravity != gravity {
		t.Errorf("g.gravity = %v, want %v", g.gravity, gravity)
		return
	}
}

func TestGravityForceGenerator_SettersAndGetters(t *testing.T) {
	var gravity glm.Vec3
	g := NewGravityForceGenerator(&gravity)
	gravity = glm.Vec3{X: 1, Y: 2, Z: 3}
	g.SetGravityVec(&gravity)
	if g.gravity != gravity {
		t.Errorf("SetGravityVec = %v, want %v", g.gravity, gravity)
	}
	g.SetGravity3f(4, 5, 6)
	if g.gravity != (glm.Vec3{X: 4, Y: 5, Z: 6}) {
		t.Errorf("SetGravity3f = %v, want %v", g.gravity, (glm.Vec3{X: 4, Y: 5, Z: 6}))
	}
	if grav := g.Gravity(); grav != (glm.Vec3{X: 4, Y: 5, Z: 6}) {
		t.Errorf("Gravity = %v, want %v", grav, (glm.Vec3{X: 4, Y: 5, Z: 6}))
	}
}

func TestGravityForceGenerator_UpdateForce(t *testing.T) {
	const (
		duration   = 0.1
		iterations = 10
	)
	gravity := glm.Vec3{X: 0, Y: -10, Z: 0}
	g := NewGravityForceGenerator(&gravity)

	var b RigidBody

	b.SetLinearDamping(1)
	b.SetMass(1)
	it := sphereInertiaTensor(b.Mass(), 1)
	b.SetInertiaTensor(&it)

	for x := 0; x < iterations; x++ {
		g.UpdateForce(&b, duration)
		b.Integrate(duration)
	}

	expected := glm.Vec3{X: 0, Y: -5.5, Z: 0}
	qiden := glm.QuatIdent()
	if pos, ori := b.Position(), b.Orientation(); !pos.EqualThreshold(&expected, 1e-4) ||
		!qiden.EqualThreshold(&ori, 1e-4) {
		t.Errorf("Position, Orientation = %v, %v want %v, %v", pos, ori, expected, qiden)
		return
	}
}

func TestGravityForceGenerator_UpdateForce_InfiniteMass(t *testing.T) {
	const (
		duration   = 0.1
		iterations = 10
	)
	gravity := glm.Vec3{X: 0, Y: -10, Z: 0}
	g := NewGravityForceGenerator(&gravity)

	var b RigidBody

	b.SetLinearDamping(1)
	b.SetMass(0)
	it := sphereInertiaTensor(b.Mass(), 1)
	b.SetInertiaTensor(&it)

	for x := 0; x < iterations; x++ {
		g.UpdateForce(&b, duration)
		b.Integrate(duration)
	}

	var expected glm.Vec3
	qiden := glm.QuatIdent()
	if pos, ori := b.Position(), b.Orientation(); !pos.EqualThreshold(&expected, 1e-4) ||
		!qiden.EqualThreshold(&ori, 1e-4) {
		t.Errorf("Position, Orientation = %v, %v want %v, %v", pos, ori, expected, qiden)
		return
	}
}

func BenchmarkGravityForceGenerator_UpdateForce(b *testing.B) {
	const (
		duration = 0.1
	)
	gravity := glm.Vec3{X: 0, Y: -10, Z: 0}
	g := NewGravityForceGenerator(&gravity)

	var body RigidBody
	body.SetLinearDamping(1)
	body.SetMass(1)
	it := sphereInertiaTensor(body.Mass(), 1)
	body.SetInertiaTensor(&it)

	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		g.UpdateForce(&body, duration)
	}
}
