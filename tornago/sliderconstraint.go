package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// SliderToWorldConstraint represents a slider that attaches a point on a rigid
// body to a world point. A slider has a minimum and maximum length it can
// extend and will generate contacts when the distance between the body and the
// point goes out of bounds.
type SliderToWorldConstraint struct {
	worldPoint           glm.Vec3
	localPoint           glm.Vec3
	body                 *RigidBody
	minlength, maxlength float32
	restitution          float32
}

// NewSliderToWorldConstraint returns a slider-to-world constraint with the
// given arguments. minlength must be smaller then maxlength, restitution must
// be between 0 and 1.
func NewSliderToWorldConstraint(worldPoint, localPoint glm.Vec3, body *RigidBody, minlength, maxlength, restitution float32) *SliderToWorldConstraint {
	return &SliderToWorldConstraint{
		worldPoint:  worldPoint,
		localPoint:  localPoint,
		body:        body,
		minlength:   minlength,
		maxlength:   maxlength,
		restitution: restitution,
	}
}

// GenerateContacts is given a slice of contacts of size at least 1. Do not
// increase the slice of the slice. The reason that the size is limited is
// to better control memory allocation and time spent resolving contacts.
// Returns how many contacts we're generated.
func (s *SliderToWorldConstraint) GenerateContacts(contacts []Contact) int {
	// find the local point on the rigid body in world position.
	bodyPointInWorld := s.body.transformMatrix.Transform(&s.localPoint)
	// from that point, get the direction to the world static point
	dir := s.worldPoint.Sub(&bodyPointInWorld)
	l2 := dir.Len2()

	// if we're farther then a certain distance away.
	if l2 > s.maxlength*s.maxlength {
		// normal of the contact
		normal := dir.Normalized()

		contacts[0] = Contact{
			bodies:      [2]*RigidBody{s.body, nil},
			point:       bodyPointInWorld,
			normal:      normal,
			penetration: math.Sqrt(l2) - s.maxlength,
			friction:    0,
			restitution: (s.restitution + s.body.restitution) / 2,
		}
		return 1
	}

	// if we're farther then a certain distance away.
	if l2 < s.minlength*s.minlength {
		// normal of the contact
		normal := dir.Normalized()

		contacts[0] = Contact{
			bodies:      [2]*RigidBody{s.body, nil},
			point:       bodyPointInWorld,
			normal:      normal.Inverse(), // point AWAY from s.worldPoint
			penetration: math.Sqrt(l2) - s.minlength,
			friction:    0,
			restitution: (s.restitution + s.body.restitution) / 2,
		}
		return 1
	}
	return 0
}
