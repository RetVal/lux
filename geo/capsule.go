package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/tensors"
	"github.com/luxengine/lux/math"
)

// Capsule is a cylinder with round end or can be used as a swept sphere.
type Capsule struct {
	A, B   glm.Vec3
	Radius float32
}

// ShapeType returns the shape type for aabbs.
func (*Capsule) ShapeType() int {
	return capsuleShapeType
}

// Volume returns the volume of the capsule.
func (c *Capsule) Volume() float32 {
	d := c.B.Sub(&c.A)
	return math.Pi * c.Radius * c.Radius * (4.0/3.0*c.Radius + d.Len())
}

// Mass returns the mass of the capsule, in mass unit, given the density in
// (mass unit/distance unit^3).
func (c *Capsule) Mass(density float32) float32 {
	return density * c.Volume()
}

// InertiaTensor returns the inertia tensor of the capsule.
func (c *Capsule) InertiaTensor(density float32) glm.Mat3 {
	d := c.B.Sub(&c.A)
	return tensors.Capsule(c.Mass(density), c.Radius, d.Len())
}

// TestCapsuleCapsule returns true if these Capsules overlap.
func TestCapsuleCapsule(a, b *Capsule) bool {
	_, _, u, _, _ := ClosestPointSegmentSegment(&a.A, &a.B, &b.A, &b.B)
	// If squared distance is smaller than squared sum of radii, they collide
	r := a.Radius + b.Radius
	return u <= r*r
}

// TestCapsuleSphere returns true if the capsule and the sphere overlap.
func TestCapsuleSphere(c *Capsule, s *Sphere) bool {
	dist2 := SqDistPointSegment(&c.A, &c.B, &s.Center)
	r := s.Radius + c.Radius
	return dist2 <= r*r
}
