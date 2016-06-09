package tornago

import (
	"testing"
)

func TestResolvePotentialContact_CollisionGrouping(t *testing.T) {
	rb1 := NewRigidBody()
	rb2 := NewRigidBody()
	rb1.SetCollisionShape(NewCollisionSphere(0.5))
	rb2.SetCollisionShape(NewCollisionSphere(0.5))

	rb1.SetGroup(Group(99))
	rb2.SetGroup(Group(99))

	rb1.SetPosition3f(0, 1, 0)

	pcontacts := []potentialContact{
		{bodies: [2]*RigidBody{rb1, rb2}},
	}
	contacts := make([]Contact, 1)
	if n := resolvePotentialContacts(pcontacts, contacts); n != 1 {
		t.Errorf("Group CollideAll & CollideAll should generate a contact. %d", n)
	}

	rb1.SetGroup(Group(-1))

	if n := resolvePotentialContacts(pcontacts, contacts); n != 0 {
		t.Errorf("Group CollideAll & CollideNone should not generate a contact. %d", n)
	}
}
