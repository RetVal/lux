package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/flops"
	"github.com/luxengine/lux/math"
)

// RodConstraintToWorld represents a rod constraint, meaning it must always be
// at the same length. This is the version that attaches a rigidbody to the
// world.
type RodConstraintToWorld struct {
	// The length of the rod.
	Length float32
	// The body this rod is attached to.
	Body *RigidBody
	// The point in local space inside the rigidbody.
	LocalPoint glm.Vec3
	// the point in world coordinate that this constraint is attached to.
	WorldPoint glm.Vec3
}

// RodConstraintToBody represents a rod constraint, meaning it must always be
// at the same length. This is the version that attaches 2 rigidbody togheter.
type RodConstraintToBody struct {
	// The length of the rod.
	Length float32
	// the bodies involved in the constraint.
	Bodies [2]*RigidBody
	// the local points for each bodies.
	LocalPoints [2]glm.Vec3
}

// NewRodConstraintToWorld returns a new RodConstraintToWorld with the given
// parameters.
func NewRodConstraintToWorld(length float32, body *RigidBody, localPoint, worldPoint *glm.Vec3) *RodConstraintToWorld {
	return &RodConstraintToWorld{
		Length:     length,
		Body:       body,
		LocalPoint: *localPoint,
		WorldPoint: *worldPoint,
	}
}

// NewRodConstraintToBody returns a new RodConstraintToBody with the given
// parameters.
func NewRodConstraintToBody(length float32, body0, body1 *RigidBody, localPoint0, localPoint1 *glm.Vec3) *RodConstraintToBody {
	return &RodConstraintToBody{
		Length:      length,
		Bodies:      [2]*RigidBody{body0, body1},
		LocalPoints: [2]glm.Vec3{*localPoint0, *localPoint1},
	}
}

// GenerateContacts is given a slice of contacts of size at least 1. Do not
// increase the slice of the slice. The reason that the size is limited is
// to better control memory allocation and time spent resolving contacts.
// Returns how many contacts we're generated.
func (r *RodConstraintToWorld) GenerateContacts(contacts []Contact) int {

	bodyPointInWorld := r.Body.transformMatrix.Transform(&r.LocalPoint)
	// from that point, get the direction to the world static point
	dir := r.WorldPoint.Sub(&bodyPointInWorld)

	// if we're farther then a certain distance away.
	if l2 := dir.Len2(); !flops.Eq(l2, r.Length*r.Length) {
		// normal of the contact
		normal := dir.Normalized()

		contacts[0] = Contact{
			bodies:      [2]*RigidBody{r.Body, nil},
			point:       bodyPointInWorld,
			normal:      normal,
			penetration: math.Sqrt(l2) - r.Length,
			friction:    0,
			restitution: 0,
		}
		return 1
	}
	return 0
}

// GenerateContacts is given a slice of contacts of size at least 1. Do not
// increase the slice of the slice. The reason that the size is limited is
// to better control memory allocation and time spent resolving contacts.
// Returns how many contacts we're generated.
func (r *RodConstraintToBody) GenerateContacts(contacts []Contact) int {
	worldPoints := [2]glm.Vec3{
		r.Bodies[0].transformMatrix.Transform(&r.LocalPoints[0]),
		r.Bodies[1].transformMatrix.Transform(&r.LocalPoints[1]),
	}

	dir := worldPoints[0].Sub(&worldPoints[1])

	// if we're farther then a certain distance away.
	if l2 := dir.Len2(); !flops.Eq(l2, r.Length*r.Length) {
		// normal of the contact
		normal := dir.Normalized()

		contacts[0] = Contact{
			bodies:      r.Bodies,
			point:       worldPoints[0],
			normal:      normal,
			penetration: math.Sqrt(l2) - r.Length,
			friction:    0,
			restitution: 0,
		}
		return 1
	}
	return 0
}
