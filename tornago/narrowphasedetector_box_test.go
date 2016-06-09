package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

var boxboxTests = []struct {
	b1          CollisionBox
	b2          CollisionBox
	normal      glm.Vec3
	point       glm.Vec3
	penetration float32
	contact     bool
}{
	{ // 0 nor colliding
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 3, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{},
		point:       glm.Vec3{},
		penetration: 0,
		contact:     false,
	},
	{ // 1 same position/orientation
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 2, 2, 2),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 2, 2, 2),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{1, 0, 0},
		point:       glm.Vec3{1, 1, 1},
		penetration: 2,
		contact:     true,
	},
	{ // 2 same position different orientation
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.Quat{0.877613, glm.Vec3{-0.33896565, 0, 0.33896565}},
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{0, -0.3602956, -0.9328382},
		point:       glm.Vec3{0.2062844, -0.5340896, 0.20628464},
		penetration: 2.5862675,
		contact:     true,
	},
	{ // 3 same position/orientation but not aligned with world axis
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.Quat{0.877613, glm.Vec3{-0.33896565, 0, 0.33896565}},
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.Quat{0.877613, glm.Vec3{-0.33896565, 0, 0.33896565}},
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{0.77020454, 0.59496135, -0.22979543},
		point:       glm.Vec3{1.1353705, 0.6495136, 1.1353705},
		penetration: 2,
		contact:     true,
	},
	{ // 4 diff pos same ori
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 1, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 2, 2, 2),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 2, 2, 2),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{0, 1, 0},
		point:       glm.Vec3{1, 1, 1},
		penetration: 1,
		contact:     true,
	},
	{ // 5 diff pos dif ori
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.Quat{0.877613, glm.Vec3{-0.33896565, 0, 0.33896565}},
				position:             glm.Vec3{0.25, 0.25, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{0.77020454, 0.59496135, -0.22979543},
		point:       glm.Vec3{1, 1, -1},
		penetration: 2.25367,
		contact:     true,
	},
	{ // 6 diff pos same ori
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, -1, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{0, -1, 0},
		point:       glm.Vec3{1, -1, 1},
		penetration: 1,
		contact:     true,
	},
	{ // 7 diff pos same ori
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{1, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{1, 0, 0},
		point:       glm.Vec3{1, 1, 1},
		penetration: 1,
		contact:     true,
	},
	{ // 8 diff pos same ori
		b1: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{-1, -1, -1},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		b2: CollisionBox{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: cuboidInertiaTensor(1, 1, 1, 1),
				linearDamping:        1,
				angularDamping:       1,
				restitution:          1,
			},
			halfSize: glm.Vec3{1, 1, 1},
		},
		normal:      glm.Vec3{-1, 0, 0},
		point:       glm.Vec3{-1, 1, 1},
		penetration: 1,
		contact:     true,
	},
}

func TestBoxBoxEarly(t *testing.T) {
	for i, test := range boxboxTests {
		test.b1.body.calculateDerivedData()
		test.b2.body.calculateDerivedData()

		if boxAndBoxEarly(&test.b1, &test.b2) != test.contact {
			t.Errorf("%d. Unexpected answer", i)
			continue
		}
	}
}

func TestBoxAndBox(t *testing.T) {
	for i, test := range boxboxTests {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%d. %s", i, r)
				}
			}()
			test.b1.body.calculateDerivedData()
			test.b2.body.calculateDerivedData()

			contacts := make([]Contact, 1)
			numcontacts := boxAndBox(&test.b1, &test.b2, contacts)

			if !test.contact {
				if numcontacts != 0 {
					t.Errorf("[%d] weird number of contacts generated when we expected 0, %d", i, numcontacts)
				}
				return
			}

			if numcontacts != 1 {
				t.Errorf("[%d] weird number of contacts generated when we expected 1, %d", i, numcontacts)
				return
			}
			contact := contacts[0]

			if contact.bodies[0] == nil || contact.bodies[1] == nil {
				t.Errorf("[%d] contact bodies should not be nil", i)
			}

			if !contact.normal.EqualThreshold(&test.normal, 1e-4) {
				t.Errorf("[%d] normal = %v, want %v", i, contact.normal, test.normal)
			}

			if !glm.FloatEqualThreshold(contact.penetration, test.penetration, 1e-4) {
				t.Errorf("[%d] penetration = %v, want %v", i, contact.penetration, test.penetration)
			}

			if !contact.point.EqualThreshold(&test.point, 1e-4) {
				t.Errorf("[%d] point = %v, want %v", i, contact.point, test.point)
			}
		}()
	}
}
