package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"github.com/luxengine/lux/math"
	"testing"
)

var _ Constraint = &RodConstraintToWorld{}
var _ Constraint = &RodConstraintToBody{}

func TestRodConstraintToWorld_GenerateContacts(t *testing.T) {
	type Case struct {
		rod     *RodConstraintToWorld
		contact *Contact
	}
	tests := []Case{
		// no contact
		func() Case {
			b := NewRigidBody()
			b.calculateDerivedData()
			rod := &RodConstraintToWorld{
				Length:     10,
				Body:       b,
				LocalPoint: glm.Vec3{},
				WorldPoint: glm.Vec3{10, 0, 0},
			}
			return Case{
				rod:     rod,
				contact: nil,
			}
		}(),
		// too close
		func() Case {
			b := NewRigidBody()
			b.SetPosition3f(0.1, 0, 0)
			b.calculateDerivedData()
			return Case{
				rod: &RodConstraintToWorld{
					Length:     10,
					Body:       b,
					LocalPoint: glm.Vec3{},
					WorldPoint: glm.Vec3{10, 0, 0},
				},
				contact: &Contact{
					bodies:      [2]*RigidBody{b, nil},
					point:       glm.Vec3{0.1, 0, 0},
					normal:      glm.Vec3{1, 0, 0},
					penetration: -0.10000038146972656,
				},
			}
		}(),
		// too far
		func() Case {
			b := NewRigidBody()
			b.SetPosition3f(-0.1, 0, 0)
			b.calculateDerivedData()
			return Case{
				rod: &RodConstraintToWorld{
					Length:     10,
					Body:       b,
					LocalPoint: glm.Vec3{},
					WorldPoint: glm.Vec3{10, 0, 0},
				},
				contact: &Contact{
					bodies:      [2]*RigidBody{b, nil},
					point:       glm.Vec3{-0.1, 0, 0},
					normal:      glm.Vec3{0.99999994039535522, 0, 0},
					penetration: 0.10000038146972656,
				},
			}
		}(),
		// NaN
		func() Case {
			b := NewRigidBody()
			b.SetPosition3f(-0.1, 0, 0)
			b.calculateDerivedData()
			return Case{
				rod: &RodConstraintToWorld{
					Length:     10,
					Body:       b,
					LocalPoint: glm.Vec3{math.NaN(), math.NaN(), math.NaN()},
					WorldPoint: glm.Vec3{10, 0, 0},
				},
				contact: &Contact{
					bodies:      [2]*RigidBody{b, nil},
					point:       glm.Vec3{math.NaN(), math.NaN(), math.NaN()},
					normal:      glm.Vec3{math.NaN(), math.NaN(), math.NaN()},
					penetration: math.NaN(),
				},
			}
		}(),
		// NaN
		func() Case {
			b := NewRigidBody()
			b.SetPosition3f(-0.1, 0, 0)
			b.calculateDerivedData()
			return Case{
				rod: &RodConstraintToWorld{
					Length:     10,
					Body:       b,
					LocalPoint: glm.Vec3{0, 0, 0},
					WorldPoint: glm.Vec3{math.NaN(), math.NaN(), math.NaN()},
				},
				contact: &Contact{
					bodies:      [2]*RigidBody{b, nil},
					point:       glm.Vec3{-0.1, 0, 0},
					normal:      glm.Vec3{math.NaN(), math.NaN(), math.NaN()},
					penetration: math.NaN(),
				},
			}
		}(),
	}
	for i, test := range tests {
		contacts := make([]Contact, 1)
		n := test.rod.GenerateContacts(contacts)
		if test.contact == nil {
			if n != 0 {
				t.Errorf("[%d] generated %d contacts, want 0", i, n)
			}
			continue
		}
		contact := contacts[0]
		if test.contact.bodies != contact.bodies {
			t.Errorf("[%d] bodies = %v, want %v", i, contact.bodies, test.contact.bodies)
		}

		if !glmtesting.FloatEqual(test.contact.Penetration(), contact.Penetration()) {
			t.Errorf("[%d] penetration = %.17f, want %.17f", i, contact.Penetration(), test.contact.Penetration())
		}

		if !glmtesting.Vec3Equal(test.contact.point, contact.point) {
			t.Errorf("[%d] point = %s, want %s", i, contact.point.String(), test.contact.point.String())
		}
		if !glmtesting.Vec3Equal(test.contact.normal, contact.normal) {
			t.Errorf("[%d] normal = %s, want %s", i, contact.normal.String(), test.contact.normal.String())
		}
	}
}

func TestRodConstraintToBody_GenerateContacts(t *testing.T) {
	type Case struct {
		rod     *RodConstraintToBody
		contact *Contact
	}
	tests := []Case{
		// no contact
		func() Case {
			b0 := NewRigidBody()
			b0.calculateDerivedData()
			b1 := NewRigidBody()
			b1.SetPosition3f(0, 10, 0)
			b1.calculateDerivedData()
			rod := &RodConstraintToBody{
				Length:      10,
				Bodies:      [2]*RigidBody{b0, b1},
				LocalPoints: [2]glm.Vec3{},
			}
			return Case{
				rod:     rod,
				contact: nil,
			}
		}(),
		// too close
		func() Case {
			b0 := NewRigidBody()
			b0.calculateDerivedData()
			b1 := NewRigidBody()
			b1.SetPosition3f(0, 9.9, 0)
			b1.calculateDerivedData()
			rod := &RodConstraintToBody{
				Length:      10,
				Bodies:      [2]*RigidBody{b0, b1},
				LocalPoints: [2]glm.Vec3{},
			}
			return Case{
				rod: rod,
				contact: &Contact{
					bodies:      [2]*RigidBody{b0, b1},
					point:       glm.Vec3{0, 0, 0},
					normal:      glm.Vec3{0, -1, 0},
					penetration: -0.10000038146972656,
				},
			}
		}(),
		// too far
		func() Case {
			b0 := NewRigidBody()
			b0.calculateDerivedData()
			b1 := NewRigidBody()
			b1.SetPosition3f(0, 10.1, 0)
			b1.calculateDerivedData()
			rod := &RodConstraintToBody{
				Length:      10,
				Bodies:      [2]*RigidBody{b0, b1},
				LocalPoints: [2]glm.Vec3{},
			}
			return Case{
				rod: rod,
				contact: &Contact{
					bodies:      [2]*RigidBody{b0, b1},
					point:       glm.Vec3{0, 0, 0},
					normal:      glm.Vec3{0, -0.99999994039535522, 0},
					penetration: 0.10000038146972656,
				},
			}
		}(),
		// NaN
		func() Case {
			b0 := NewRigidBody()
			b0.calculateDerivedData()
			b1 := NewRigidBody()
			b1.SetPosition3f(0, 10, 0)
			b1.calculateDerivedData()
			rod := &RodConstraintToBody{
				Length:      10,
				Bodies:      [2]*RigidBody{b0, b1},
				LocalPoints: [2]glm.Vec3{{math.NaN(), math.NaN(), math.NaN()}, {}},
			}
			return Case{
				rod: rod,
				contact: &Contact{
					bodies:      [2]*RigidBody{b0, b1},
					point:       glm.Vec3{math.NaN(), math.NaN(), math.NaN()},
					normal:      glm.Vec3{math.NaN(), math.NaN(), math.NaN()},
					penetration: math.NaN(),
				},
			}
		}(),
		// NaN
		func() Case {
			b0 := NewRigidBody()
			b0.calculateDerivedData()
			b1 := NewRigidBody()
			b1.SetPosition3f(0, 10, 0)
			b1.calculateDerivedData()
			rod := &RodConstraintToBody{
				Length:      10,
				Bodies:      [2]*RigidBody{b0, b1},
				LocalPoints: [2]glm.Vec3{{}, {math.NaN(), math.NaN(), math.NaN()}},
			}
			return Case{
				rod: rod,
				contact: &Contact{
					bodies:      [2]*RigidBody{b0, b1},
					point:       glm.Vec3{0, 0, 0},
					normal:      glm.Vec3{math.NaN(), math.NaN(), math.NaN()},
					penetration: math.NaN(),
				},
			}
		}(),
		// NaN
		func() Case {
			b0 := NewRigidBody()
			b0.calculateDerivedData()
			b1 := NewRigidBody()
			b1.SetPosition3f(0, 10, 0)
			b1.calculateDerivedData()
			rod := &RodConstraintToBody{
				Length:      math.NaN(),
				Bodies:      [2]*RigidBody{b0, b1},
				LocalPoints: [2]glm.Vec3{},
			}
			return Case{
				rod: rod,
				contact: &Contact{
					bodies:      [2]*RigidBody{b0, b1},
					point:       glm.Vec3{0, 0, 0},
					normal:      glm.Vec3{0, -1, 0},
					penetration: math.NaN(),
				},
			}
		}(),
	}
	for i, test := range tests {
		contacts := make([]Contact, 1)
		n := test.rod.GenerateContacts(contacts)
		if test.contact == nil {
			if n != 0 {
				t.Errorf("[%d] generated %d contacts, want 0", i, n)
			}
			continue
		}
		contact := contacts[0]
		if test.contact.bodies != contact.bodies {
			t.Errorf("[%d] bodies = %v, want %v", i, contact.bodies, test.contact.bodies)
		}

		if !glmtesting.FloatEqual(test.contact.Penetration(), contact.Penetration()) {
			t.Errorf("[%d] penetration = %.17f, want %.17f", i, contact.Penetration(), test.contact.Penetration())
		}

		if !glmtesting.Vec3Equal(test.contact.point, contact.point) {
			t.Errorf("[%d] point = %s, want %s", i, contact.point.String(), test.contact.point.String())
		}
		if !glmtesting.Vec3Equal(test.contact.normal, contact.normal) {
			t.Errorf("[%d] normal = %s, want %s", i, contact.normal.String(), test.contact.normal.String())
		}
	}
}
