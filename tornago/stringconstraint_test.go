package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestStringToWorldConstraint_GenerateContacts(t *testing.T) {
	body := &RigidBody{
		inverseMass:          1,
		orientation:          glm.QuatIdent(),
		position:             glm.Vec3{0, -0.5, 0},
		velocity:             glm.Vec3{0, 0, 0},
		inverseInertiaTensor: sphereInertiaTensor(1, 1),
		linearDamping:        1,
		angularDamping:       1,
	}
	body.calculateDerivedData()

	c := NewStringToWorldConstraint(glm.Vec3{0, 5, 0}, glm.Vec3{0, 0, 0}, body, 9, 1)
	contacts := make([]Contact, 1)
	n := c.GenerateContacts(contacts)
	t.Logf("contacts generated: %d", n)
	t.Logf("%+v", contacts[0])
}
