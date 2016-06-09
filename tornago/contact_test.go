package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
	"testing"
)

func TestContact_SetterGetter(t *testing.T) {
	const (
		penetration = 5.4
		restitution = 1.3
		friction    = 7.6
	)
	c := Contact{
		penetration: penetration,
		restitution: restitution,
		friction:    friction,
	}

	if c.Friction() != friction {
		t.Errorf("c.Friction() = %f, want %f", c.Friction(), friction)
	}

	if c.Restitution() != restitution {
		t.Errorf("c.Restitution() = %f, want %f", c.Restitution(), restitution)
	}

	if c.Penetration() != penetration {
		t.Errorf("c.Penetration() = %f, want %f", c.Penetration(), penetration)
	}
}

func TestContact_SwapBodies(t *testing.T) {
	var b1, b2 RigidBody
	c := Contact{
		bodies: [2]*RigidBody{&b1, &b2},
	}

	c.swapIfNeed()

	if c.bodies[0] != &b1 || c.bodies[1] != &b2 {
		t.Errorf("bodies = %v, want [%p, %p]", c.bodies, &b1, &b2)
	}

	c = Contact{
		bodies: [2]*RigidBody{&b1, nil},
	}

	c.swapIfNeed()

	if c.bodies[0] != &b1 || c.bodies[1] != nil {
		t.Errorf("bodies = %v, want [%p, %p]", c.bodies, &b1, (*RigidBody)(nil))
	}

	c = Contact{
		bodies: [2]*RigidBody{nil, &b1},
	}

	c.swapIfNeed()

	if c.bodies[0] != &b1 || c.bodies[1] != nil {
		t.Errorf("bodies = %v, want [%p, %p]", c.bodies, &b1, (*RigidBody)(nil))
	}
}

func TestContact_ResolveContact(t *testing.T) {
	type Result struct {
		pos, vel [2]glm.Vec3
	}
	var tests = []struct {
		s1     CollisionSphere
		s2     CollisionSphere
		result Result
	}{
		{ // 0. no velocity, no penetration
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 2, 0},
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
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 2, 0}, {0, 0, 0}},
				vel: [2]glm.Vec3{{0, 0, 0}, {0, 0, 0}},
			},
		},
		{ // 1. no penetration, just equal perpendicular velocity
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 2, 0},
					velocity:             glm.Vec3{0, -1, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{0, 1, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 2, 0}, {0, 0, 0}},
				vel: [2]glm.Vec3{{0, 1, 0}, {0, -1, 0}},
			},
		},
		{ // 2. no penetration, just non equal perpendicular velocity
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 2, 0},
					velocity:             glm.Vec3{0, -1, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{0, 0.5, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 2, 0}, {0, 0, 0}},
				vel: [2]glm.Vec3{{0, 0.5, 0}, {0, -1, 0}},
			},
		},
		{ // 3. no penetration, just non-perpendicular velocity
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 2, 0},
					velocity:             glm.Vec3{1, 0, 1},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{-1, 0, -1},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 2, 0}, {0, 0, 0}},
				vel: [2]glm.Vec3{{1, 0, 1}, {-1, 0, -1}},
			},
		},
		{ // 4. velocity with non-0 perpendicular velocity
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 2, 0},
					velocity:             glm.Vec3{-1, -1, 1},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{1, 1, -1},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 2, 0}, {0, 0, 0}},
				vel: [2]glm.Vec3{{-1, 1, 1}, {1, -1, -1}},
			},
		},
		{ // 5. velocity with pi/4 angle between 2 spheres.
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{math.Sqrt2, math.Sqrt2, 0},
					velocity:             glm.Vec3{-1, -1, 1},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{1, 1, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{math.Sqrt2, math.Sqrt2, 0}, {0, 0, 0}},
				vel: [2]glm.Vec3{{1, 1, 1}, {-1, -1, 0}},
			},
		},
		{ // 6. no velocity, some penetration
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 10, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 9.5),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
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
					restitution:          1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 10.25, 0}, {0, -0.25, 0}},
				vel: [2]glm.Vec3{{0, 0, 0}, {0, 0, 0}},
			},
		},
		{ // 7. no velocity, some penetration, different weight
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          0.25,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 10, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(4, 9.5),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
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
					restitution:          1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 10.1, 0}, {0, -0.4, 0}},
				vel: [2]glm.Vec3{{0, 0, 0}, {0, 0, 0}},
			},
		},
		{ // 8. no velocity, some penetration, infinite weight on one
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(4, 9.5),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          0,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, -5, 0},
					velocity:             glm.Vec3{0, 0, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 5,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 1, 0}, {0, -5, 0}},
				vel: [2]glm.Vec3{{0, 0, 0}, {0, 0, 0}},
			},
		},
		{ // 9. no penetration, just equal perpendicular velocity
			s1: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 2, 0},
					velocity:             glm.Vec3{0, -1, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			s2: CollisionSphere{
				body: &RigidBody{
					inverseMass:          1,
					orientation:          glm.QuatIdent(),
					position:             glm.Vec3{0, 0, 0},
					velocity:             glm.Vec3{0, 1, 0},
					inverseInertiaTensor: sphereInertiaTensor(1, 1),
					linearDamping:        1,
					angularDamping:       1,
					restitution:          1,
				},
				radius: 1,
			},
			result: Result{
				pos: [2]glm.Vec3{{0, 2, 0}, {0, 0, 0}},
				vel: [2]glm.Vec3{{0, 1, 0}, {0, -1, 0}},
			},
		},
	}
	var data contactDerivateData
	for i, test := range tests {

		test.s1.body.calculateDerivedData()
		test.s2.body.calculateDerivedData()

		contacts := make([]Contact, 1)
		if sphereAndSphere(&test.s1, &test.s2, contacts) != 1 {
			t.Errorf("[%d] did not generate any contact, can't resolve it.", i)
			continue
		}
		contact := contacts[0]

		prepos := [2]glm.Vec3{test.s1.body.Position(), test.s2.body.Position()}
		prevel := [2]glm.Vec3{test.s1.body.Velocity(), test.s2.body.Velocity()}

		var velch, rotch [2]glm.Vec3
		var linch, angch [2]glm.Vec3

		contact.calculateDerivateData(&data, 0.16)
		contact.resolveVelocity(&data, &velch, &rotch)
		contact.resolvePenetration(&data, &linch, &angch)

		if p1, p2 := test.s1.body.Position(), test.s2.body.Position(); !p1.EqualThreshold(&test.result.pos[0], 1e-2) || !p2.EqualThreshold(&test.result.pos[1], 1e-2) {
			t.Errorf("%d. pos[1] %v->%v, want %v, pos[2] %v->%v, want %v", i, prepos[0], p1, test.result.pos[0], prepos[1], p2, test.result.pos[1])
		}

		if v1, v2 := test.s1.body.Velocity(), test.s2.body.Velocity(); !v1.EqualThreshold(&test.result.vel[0], 1e-2) || !v2.EqualThreshold(&test.result.vel[1], 1e-2) {
			t.Errorf("%d. vel[1] %v->%v, want %v, vel[2] %v->%v, want %v", i, prevel[0], v1, test.result.vel[0], prevel[1], v2, test.result.vel[1])
		}
	}
}

func genContact() Contact {
	var test = struct {
		s1 CollisionSphere
		s2 CollisionSphere
	}{
		s1: CollisionSphere{
			body: &RigidBody{
				inverseMass:          1,
				orientation:          glm.QuatIdent(),
				position:             glm.Vec3{0, 2, 0},
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
				position:             glm.Vec3{0, 0, 0},
				velocity:             glm.Vec3{0, 0, 0},
				inverseInertiaTensor: sphereInertiaTensor(1, 1),
				linearDamping:        1,
				angularDamping:       1,
			},
			radius: 1,
		},
	}
	test.s1.body.calculateDerivedData()
	test.s2.body.calculateDerivedData()
	contacts := make([]Contact, 1)
	sphereAndSphere(&test.s1, &test.s2, contacts)
	return contacts[0]
}

func BenchmarkContat_resolvePenetration(b *testing.B) {
	contact := genContact()
	var data contactDerivateData
	contact.calculateDerivateData(&data, 0.16)
	var a, c [2]glm.Vec3
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		contact.resolvePenetration(&data, &a, &c)
	}
}

func BenchmarkContact_resolveVelocity(b *testing.B) {
	contact := genContact()
	var data contactDerivateData
	contact.calculateDerivateData(&data, 0.16)
	var a, c [2]glm.Vec3
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		contact.resolveVelocity(&data, &a, &c)
	}
}

func BenchmarkContact_calculateDerivedData(b *testing.B) {
	contact := genContact()
	var data contactDerivateData
	contact.calculateDerivateData(&data, 0.16)
	var a, c [2]glm.Vec3
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		contact.resolveVelocity(&data, &a, &c)
	}
}
