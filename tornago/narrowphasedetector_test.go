package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestSphereAndSphere(t *testing.T) {
	var tests = []struct {
		s1          CollisionSphere
		s2          CollisionSphere
		normal      glm.Vec3
		point       glm.Vec3
		penetration float32
		contact     bool
	}{
		{
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 1, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, -0.5, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			normal:      glm.Vec3{0, 1, 0},
			point:       glm.Vec3{0, 0.25, 0},
			penetration: 0.5,
			contact:     true,
		},
		{
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 10, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 9.5,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			normal:      glm.Vec3{0, 1, 0},
			point:       glm.Vec3{0, 5, 0},
			penetration: 0.5,
			contact:     true,
		},
		{
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{1, 1, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 0.5,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			normal:      glm.Vec3{0.70710677, 0.70710677, 0},
			point:       glm.Vec3{0.5, 0.5, 0},
			penetration: 0.08578646,
			contact:     true,
		},
		{
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{1, 30, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 0.5,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			normal:      glm.Vec3{0, 0, 0},
			point:       glm.Vec3{0, 0, 0},
			penetration: 0,
			contact:     false,
		},
	}
	for i, test := range tests {

		contacts := make([]Contact, 1)
		numcontacts := sphereAndSphere(&test.s1, &test.s2, contacts)

		if !test.contact {
			if numcontacts != 0 {
				t.Errorf("[%d] weird number of contacts generated when we expected 0, %d", i, numcontacts)
			}
			continue
		}

		if numcontacts != 1 {
			t.Errorf("[%d] weird number of contacts generated, %d", i, numcontacts)
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
	}
}

func TestSphereAndBox(t *testing.T) {
	var tests = []struct {
		s           CollisionSphere
		b           CollisionBox
		normal      glm.Vec3
		point       glm.Vec3
		penetration float32
		contact     bool
	}{
		{ // face-face
			s: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 1.5, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			b: CollisionBox{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: cuboidInertiaTensor(1, 0.5, 0.5, 0.5),
					linearDamping:        1,
					angularDamping:       1,
				},
				halfSize: glm.Vec3{0.5, 0.5, 0.5},
			},
			normal:      glm.Vec3{0, 1, 0},
			point:       glm.Vec3{0, 0.5, 0},
			penetration: 0,
			contact:     true,
		},
		{ // face-face
			s: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 1.4, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			b: CollisionBox{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: cuboidInertiaTensor(1, 0.5, 0.5, 0.5),
					linearDamping:        1,
					angularDamping:       1,
				},
				halfSize: glm.Vec3{0.5, 0.5, 0.5},
			},
			normal:      glm.Vec3{0, 1, 0},
			point:       glm.Vec3{0, 0.5, 0},
			penetration: 0.1,
			contact:     true,
		},
		{ // face-edge
			s: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0.6, 1.3, 0.6},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			b: CollisionBox{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: cuboidInertiaTensor(1, 0.5, 0.5, 0.5),
					linearDamping:        1,
					angularDamping:       1,
				},
				halfSize: glm.Vec3{0.5, 0.5, 0.5},
			},
			normal:      glm.Vec3{0.12309153, 0.9847319, 0.12309153},
			point:       glm.Vec3{0.5, 0.5, 0.5},
			penetration: 0.1875962,
			contact:     true,
		},
		{ // face-edge
			s: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0.51, 1.1, 0.51},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			b: CollisionBox{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: cuboidInertiaTensor(1, 0.5, 0.5, 0.5),
					linearDamping:        1,
					angularDamping:       1,
				},
				halfSize: glm.Vec3{0.5, 0.5, 0.5},
			},
			normal:      glm.Vec3{0.016662024, 0.9997224, 0.016662024},
			point:       glm.Vec3{0.5, 0.5, 0.5},
			penetration: 0.39983338,
			contact:     true,
		},
		{ // face-edge
			s: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0.51, -0.51, 0.51},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			b: CollisionBox{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: cuboidInertiaTensor(1, 0.5, 0.5, 0.5),
					linearDamping:        1,
					angularDamping:       1,
				},
				halfSize: glm.Vec3{0.5, 0.5, 0.5},
			},
			normal:      glm.Vec3{0.57735026, -0.57735026, 0.57735026},
			point:       glm.Vec3{0.5, -0.5, 0.5},
			penetration: 0.9826795,
			contact:     true,
		},
	}
	for i, test := range tests {
		test.b.body.calculateDerivedData()
		test.s.body.calculateDerivedData()

		contacts := make([]Contact, 1)
		numcontacts := sphereAndBox(&test.s, &test.b, contacts)

		if !test.contact {
			if numcontacts != 0 {
				t.Errorf("%d. contacts generated when not expected, numcontacts = %d", i, numcontacts)
			}
			continue
		}
		if numcontacts != 1 {
			t.Errorf("%d. weird number of contacts generated, %d", i, numcontacts)
			continue
		}

		contact := contacts[0]

		if contact.bodies[0] == nil || contact.bodies[1] == nil {
			t.Errorf("%d. contact bodies should not be nil", i)
		}
		if !contact.normal.EqualThreshold(&test.normal, 1e-4) {
			t.Errorf("%d. normal = %v, want %v", i, contact.normal, test.normal)
		}
		if !glm.FloatEqualThreshold(contact.penetration, test.penetration, 1e-4) {
			t.Errorf("%d. penetration = %v, want %v", i, contact.penetration, test.penetration)
		}
		if !contact.point.EqualThreshold(&test.point, 1e-4) {
			t.Errorf("%d. point = %v, want %v", i, contact.point, test.point)
		}
	}
}

func BenchmarkCollision_SphereSphere(b *testing.B) {
	s1 := CollisionSphere{
		body: &RigidBody{
			inverseMass:          1,
			orientation:          glm.QuatIdent(),
			position:             glm.Vec3{1, 1, 0},
			velocity:             glm.Vec3{0, 0, 0},
			inverseInertiaTensor: sphereInertiaTensor(1, 1),
			linearDamping:        1,
			angularDamping:       1,
		},
		radius: 0.5,
	}
	s2 := CollisionSphere{
		body: &RigidBody{
			inverseMass:          1,
			orientation:          glm.QuatIdent(),
			position:             glm.Vec3{0, 0, 0},
			velocity:             glm.Vec3{0, 0, 0},
			inverseInertiaTensor: sphereInertiaTensor(1, 1),
			linearDamping:        1,
			angularDamping:       1,
		},
		radius: 1,
	}

	contacts := make([]Contact, 1)
	for x := 0; x < b.N; x++ {
		sphereAndSphere(&s1, &s2, contacts)
	}
}

func BenchmarkCollision_SphereBox(b *testing.B) {
	s := CollisionSphere{
		body: &RigidBody{
			inverseMass:          1,
			orientation:          glm.QuatIdent(),
			position:             glm.Vec3{0, 1.5, 0},
			velocity:             glm.Vec3{0, 0, 0},
			inverseInertiaTensor: sphereInertiaTensor(1, 1),
			linearDamping:        1,
			angularDamping:       1,
		},
		radius: 1,
	}
	box := CollisionBox{
		body: &RigidBody{
			inverseMass:          1,
			orientation:          glm.QuatIdent(),
			position:             glm.Vec3{0, 0, 0},
			velocity:             glm.Vec3{0, 0, 0},
			inverseInertiaTensor: sphereInertiaTensor(1, 1),
			linearDamping:        1,
			angularDamping:       1,
		},
		halfSize: glm.Vec3{0.5, 0.5, 0.5},
	}
	box.body.calculateDerivedData()
	contacts := make([]Contact, 1)
	for x := 0; x < b.N; x++ {
		sphereAndBox(&s, &box, contacts)
	}
}
