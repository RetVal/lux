package tornago

import (
	"github.com/luxengine/lux/glm"
)

// BuoyancyForceGenerator helps simulate water forces, with particle however,
// since they have no size, we have to fix a limit at which it is considered
// "fully submerged". This version only works for the XZ plane but feel free to
// make one that works in the plane you want.
type BuoyancyForceGenerator struct {
	Height   float32
	MaxDepth float32
	Volume   float32

	//in kg/m^3, water is 1000
	Density float32
}

// NewBuoyancyForceGenerator returns a buoyancy force generator from the given
// arguments.
func NewBuoyancyForceGenerator(height, maxDepth, volume, density float32) *BuoyancyForceGenerator {
	var g BuoyancyForceGenerator
	g.New(height, maxDepth, volume, density)
	return &g
}

// New initialises this BuoyancyForceGenerator with the given arguments. This is
// used for memory management.
func (g *BuoyancyForceGenerator) New(height, maxDepth, volume, density float32) {
	g.Height = height
	g.MaxDepth = maxDepth
	g.Volume = volume
	g.Density = density
}

// UpdateForce applies buoyancy force to the given particle.
func (g *BuoyancyForceGenerator) UpdateForce(b *RigidBody, _ float32) {
	if !b.HasFiniteMass() {
		return
	}

	depth := b.Position().Y

	// Verify is we even need to apply a force.
	if depth >= g.Height+g.MaxDepth {
		return
	}

	var force glm.Vec3

	// We're at maximum depth, the force is constant at that point.
	if depth <= g.Height-g.MaxDepth {
		force.Y = g.Volume * g.Density
	} else {
		force.Y = -g.Volume * g.Density * ((depth - g.MaxDepth - g.Height) / (g.MaxDepth * 2))
	}
	force.MulWith(b.Mass())
	b.AddForce(&force)
}
