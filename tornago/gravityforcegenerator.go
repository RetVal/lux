package tornago

import (
	"github.com/luxengine/lux/glm"
)

// GravityForceGenerator is a generator used to apply gravity on rigid bodies.
type GravityForceGenerator struct {
	gravity glm.Vec3
}

// NewGravityForceGenerator returns a new GravityForceGenerator with the given
// gravity.
func NewGravityForceGenerator(gravity *glm.Vec3) *GravityForceGenerator {
	var g GravityForceGenerator
	g.New(gravity)
	return &g
}

// New initializes this force generator with the given arguments. Used for
// memory management.
func (g *GravityForceGenerator) New(gravity *glm.Vec3) {
	g.gravity = *gravity
}

// SetGravity3f sets gravity.
func (g *GravityForceGenerator) SetGravity3f(x, y, z float32) {
	g.gravity = glm.Vec3{X: x, Y: y, Z: z}
}

// SetGravityVec sets gravity. will panic if vector is nil.
func (g *GravityForceGenerator) SetGravityVec(gravity *glm.Vec3) {
	g.gravity = *gravity
}

// Gravity returns the gravity that this generator applies.
func (g *GravityForceGenerator) Gravity() glm.Vec3 {
	return g.gravity
}

// UpdateForce applies gravity to the given RigidBody.
func (g *GravityForceGenerator) UpdateForce(b *RigidBody, _ float32) {
	if !b.HasFiniteMass() {
		return
	}

	var force glm.Vec3
	force.MulOf(b.Mass(), &g.gravity)
	b.AddForce(&force)
}
