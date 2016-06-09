package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// StringToWorldConstraint is a constraint that generates a contact when a
// specific point in the rigid body gets too far away from a point in the world.
type StringToWorldConstraint struct {
	worldPoint  glm.Vec3
	localPoint  glm.Vec3
	body        *RigidBody
	length      float32
	restitution float32
}

// NewStringToWorldConstraint returns a new StringToWorldConstraint from the
// given arguments.
func NewStringToWorldConstraint(worldPoint glm.Vec3, localPoint glm.Vec3, body *RigidBody, length, restitution float32) Constraint {
	return &StringToWorldConstraint{
		worldPoint:  worldPoint,
		body:        body,
		localPoint:  localPoint,
		length:      length,
		restitution: restitution,
	}
}

// WorldPoint returns the point in world of this constraint.
func (c *StringToWorldConstraint) WorldPoint() glm.Vec3 {
	return c.worldPoint
}

// SetWorldPoint sets the world point to use as target.
func (c *StringToWorldConstraint) SetWorldPoint(point glm.Vec3) {
	c.worldPoint = point
}

// LocalPoint returns the local point in the body of this constraint.
func (c *StringToWorldConstraint) LocalPoint() glm.Vec3 {
	return c.localPoint
}

// SetLocalPoint sets the local point in the rigid body to use as target.
func (c *StringToWorldConstraint) SetLocalPoint(point glm.Vec3) {
	c.localPoint = point
}

// Body returns the rigid body this string is attached to.
func (c *StringToWorldConstraint) Body() *RigidBody {
	return c.body
}

// SetBody sets the rigid body that this string is attached to.
func (c *StringToWorldConstraint) SetBody(body *RigidBody) {
	c.body = body
}

// Len returns the length of the string.
func (c *StringToWorldConstraint) Len() float32 {
	return c.length
}

// SetLen sets the length of the string.
func (c *StringToWorldConstraint) SetLen(length float32) {
	c.length = length
}

// Restitution returns the restitution of this string.
func (c *StringToWorldConstraint) Restitution() float32 { return c.restitution }

// SetRestitution sets the restitution of this string.
func (c *StringToWorldConstraint) SetRestitution(restitution float32) {
	c.restitution = restitution
}

// GenerateContacts will generate maximum 1 contact if the rigid body's point in
// world coordinates is too far from the set world point.
func (c *StringToWorldConstraint) GenerateContacts(contacts []Contact) int {
	bodyPointInWorld := c.body.transformMatrix.Transform(&c.localPoint)
	dir := c.worldPoint.Sub(&bodyPointInWorld)

	// if we're farther then a certain distance away.
	if l2 := dir.Len2(); l2 > c.length*c.length {
		normal := dir.Normalized()

		contacts[0] = Contact{
			bodies:      [2]*RigidBody{c.body, nil},
			point:       bodyPointInWorld,
			normal:      normal,
			penetration: math.Sqrt(l2) - c.length,
			friction:    0,
			restitution: (c.Restitution() + c.body.restitution) / 2,
		}
		return 1
	}
	return 0
}

// StringToBodyConstraint is the same as StringToWorldConstraint excepts it
// attaches to another rigid body.
type StringToBodyConstraint struct {
	localPoints [2]glm.Vec3
	bodies      [2]*RigidBody
	length      float32
	restitution float32
}

// NewStringToBodyConstraint returns a new StringToBodyConstraint
func NewStringToBodyConstraint(localPoints [2]glm.Vec3, bodies [2]*RigidBody, length, restitution float32) Constraint {
	return &StringToBodyConstraint{
		localPoints: localPoints,
		bodies:      bodies,
		length:      length,
		restitution: restitution,
	}
}

// LocalPoint returns the local point in the body of this constraint. Use
// either 0 or 1 for the first or second rigid body else it panics.
func (c *StringToBodyConstraint) LocalPoint(index int) glm.Vec3 {
	return c.localPoints[index]
}

// SetLocalPoint sets the local point in the rigid body to use as target. Use
// either 0 or 1 for the first or second rigid body else it panics.
func (c *StringToBodyConstraint) SetLocalPoint(index int, point glm.Vec3) {
	c.localPoints[index] = point
}

// Body returns the rigid body this string is attached to. Use either 0 or 1 for
// the first or second rigid body else it panics.
func (c *StringToBodyConstraint) Body(index int) *RigidBody {
	return c.bodies[index]
}

// SetBody sets the rigid body that this string is attached to. Use either 0 or
// 1 for the first or second rigid body else it panics.
func (c *StringToBodyConstraint) SetBody(index int, body *RigidBody) {
	c.bodies[index] = body
}

// Len returns the length of the string.
func (c *StringToBodyConstraint) Len() float32 {
	return c.length
}

// SetLen sets the length of the string.
func (c *StringToBodyConstraint) SetLen(length float32) {
	c.length = length
}

// Restitution returns the restitution of this string.
func (c *StringToBodyConstraint) Restitution() float32 { return c.restitution }

// SetRestitution sets the restitution of this string.
func (c *StringToBodyConstraint) SetRestitution(restitution float32) {
	c.restitution = restitution
}

// GenerateContacts will generate maximum 1 contact if the rigid body's point in
// world coordinates is too far from the other point in the other rigid body.
func (c *StringToBodyConstraint) GenerateContacts(contacts []Contact) int {
	bodyPointsInWorld := [2]glm.Vec3{c.bodies[0].transformMatrix.Transform(&c.localPoints[0]),
		c.bodies[1].transformMatrix.Transform(&c.localPoints[1])}
	dir := bodyPointsInWorld[1].Sub(&bodyPointsInWorld[0])

	// if we're farther then a certain distance away.
	if l2 := dir.Len2(); l2 > c.length*c.length {
		contacts[0] = Contact{
			bodies:      c.bodies,
			point:       bodyPointsInWorld[1],
			normal:      dir.Normalized(),
			penetration: math.Sqrt(l2) - c.length,
			friction:    0,
			restitution: (c.Restitution() + c.bodies[0].restitution + c.bodies[1].restitution) / 3,
		}
		return 1
	}
	return 0
}
