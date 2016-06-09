package tornago

import (
	"github.com/luxengine/lux/glm"
)

// AttractionSphere is a force generator that attract/repulse objects in a
// sphere shape. (This does not use the collision
// detector)
type AttractionSphere struct {
	// The strength of the force to apply.
	Force float32
	// The center of the sphere.
	Center glm.Vec3
}

// NewAttractionSphere returns an attraction sphere with the given parameters.
func NewAttractionSphere(force float32, center *glm.Vec3) *AttractionSphere {
	var s AttractionSphere
	s.New(force, center)
	return &s
}

// New initialises this AttractionSphere with the given arguments. This is used
// for memory management.
func (s *AttractionSphere) New(force float32, center *glm.Vec3) {
	s.Force = force
	s.Center = *center
}

//UpdateForce calculates and update the force applied to the given rigid body.
func (s *AttractionSphere) UpdateForce(b *RigidBody, _ float32) {
	pos := b.Position()
	force := s.Center.Sub(&pos)
	var zero glm.Vec3
	if force.Equal(&zero) {
		return
	}
	force.Normalize()
	force.MulWith(s.Force)
	force.MulWith(b.Mass())
	b.AddForce(&force)
}

// AttractionCylinder is a force generator that attract/repulse objects that
// have their center inside the cylinder. (This does not use the collision
// detector)
type AttractionCylinder struct {
	// The strength of the force to apply.
	Force float32
	// The radius of the cylinder
	Radius float32
	// The height of the cylinder, 0 for infinite
	Height float32
	// the middle of the circle at the base of the cylinder.
	Base glm.Vec3
	// The direction in which to apply the force. If Force is positive it will
	// act against this vector. If negative it acts in the same direction as
	// this vector. This vector must be normalized.
	Direction glm.Vec3
}

// NewAttractionCylinder returns an attraction cylinder with the given
// arguments. The direction must be normalized, the radius must be positive.
func NewAttractionCylinder(force, radius, height float32, base, direction *glm.Vec3) *AttractionCylinder {
	var c AttractionCylinder
	c.New(force, radius, height, base, direction)
	return &c
}

// New initialises this AttractionCylinder with the given arguments. This is
// used for memory management.
func (c *AttractionCylinder) New(force, radius, height float32, base, direction *glm.Vec3) {
	c.Force = force
	c.Radius = radius
	c.Height = height
	c.Base = *base
	c.Direction = *direction
}

//UpdateForce calculates and update the force applied to the given rigid body.
func (c *AttractionCylinder) UpdateForce(b *RigidBody, _ float32) {
	pos := b.Position()
	dir := c.Base.Sub(&pos)
	tmp := glm.Vec3{X: 1, Y: 0, Z: 0}
	zero := glm.Vec3{}
	plane := dir.Cross(&tmp)
	if plane.Equal(&zero) {
		// bad luck it was aligned.
		tmp = glm.Vec3{X: 0, Y: 1, Z: 0}
		plane = dir.Cross(&tmp)
	}
	// find the distance on the cylinder plane to the base
	distsq := dir.Dot(&plane)
	if distsq > c.Radius*c.Radius {
		return
	}
	force := c.Direction.Mul(-c.Force)
	force.MulWith(b.Mass())
	b.AddForce(&force)
}
