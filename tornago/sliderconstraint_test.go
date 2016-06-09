package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/flops"
	"testing"
)

// check for interface satisfiability at compile time.
var _ Constraint = &SliderToWorldConstraint{}

func TestSliderToWorldConstraint_GenerateContacts(t *testing.T) {
	t.Skip("re-enable after floating point robustness is complete")
	const friction = 0
	tests := []struct {
		slider      *SliderToWorldConstraint
		normal      glm.Vec3
		point       glm.Vec3
		restitution float32
		penetration float32
		contact     bool
	}{
		{
			// ok here we're gonna use a cheat when doing table testing :)
			// anon function calling
			slider: func() *SliderToWorldConstraint {
				const (
					restitution = 0.5
				)
				sphereBody := NewRigidBody()
				sphereBody.SetPosition3f(10.1, 0, 0)
				sphereBody.SetVelocity3f(0, 0, 0)
				sphereBody.SetRestitution(restitution)
				sphereBody.calculateDerivedData()

				return NewSliderToWorldConstraint(
					glm.Vec3{0, 0, 0},
					glm.Vec3{0, 0, 0},
					sphereBody,
					0, 10, restitution)
			}(),
			normal:      glm.Vec3{-1, 0, 0},
			point:       glm.Vec3{10.1, 0, 0},
			restitution: 0.5,
			penetration: 0.1,
			contact:     true,
		},
		{
			// ok here we're gonna use a cheat when doing table testing :)
			// anon function calling
			slider: func() *SliderToWorldConstraint {
				const (
					restitution = 0.5
				)
				sphereBody := NewRigidBody()
				sphereBody.SetPosition3f(4.9, 0, 0)
				sphereBody.SetVelocity3f(0, 0, 0)
				sphereBody.SetRestitution(restitution)
				sphereBody.calculateDerivedData()

				return NewSliderToWorldConstraint(
					glm.Vec3{0, 0, 0},
					glm.Vec3{0, 0, 0},
					sphereBody,
					5, 10, restitution)
			}(),
			normal:      glm.Vec3{1, 0, 0},
			point:       glm.Vec3{10.1, 0, 0},
			restitution: 0.5,
			penetration: 0.1,
			contact:     true,
		},
		{
			// ok here we're gonna use a cheat when doing table testing :)
			// anon function calling
			slider: func() *SliderToWorldConstraint {
				const (
					restitution = 0.5
				)
				sphereBody := NewRigidBody()
				sphereBody.SetPosition3f(7.5, 0, 0)
				sphereBody.SetVelocity3f(0, 0, 0)
				sphereBody.SetRestitution(restitution)
				sphereBody.calculateDerivedData()

				return NewSliderToWorldConstraint(
					glm.Vec3{0, 0, 0},
					glm.Vec3{0, 0, 0},
					sphereBody,
					5, 10, restitution)
			}(),
			normal:      glm.Vec3{1, 0, 0},
			point:       glm.Vec3{10.1, 0, 0},
			restitution: 0.5,
			penetration: 0.1,
			contact:     false,
		},
	}
	for i, test := range tests {

		contacts := make([]Contact, 1)
		n := test.slider.GenerateContacts(contacts)
		if !test.contact {
			if n != 0 {
				t.Errorf("[%d] GenerateContacts = %d, want 0", i, n)
				continue
			}
			continue
		}
		if n != 1 {
			t.Errorf("[%d] GenerateContacts = %d, want 1", i, n)
			continue
		}

		c := contacts[0]

		// TODO(hydroflame): revisit this when we clear the floating robustness
		// issue. (this definitely SHOULD pass)
		if !c.normal.Equal(&test.normal) {
			t.Errorf("[%d] slider normal = %s, want %s", i, c.normal.String(), test.normal.String())
		}

		// TODO(hydroflame): same issue as line 38
		if !flops.Eq(c.Penetration(), test.penetration) {
			t.Errorf("[%d] slider penetration = %v, want %v", i, c.Penetration(), test.penetration)
		}

		// TODO(hydroflame): same issue as line 38
		if !c.point.Equal(&test.point) {
			t.Errorf("[%d] slider point = %s, want %s", i, c.point.String(), test.point.String())
		}

		// thats just gonna work ;)
		if !flops.Eq(c.Restitution(), test.restitution) {
			t.Errorf("[%d] slider restitution = %v, want %v", i, c.Restitution(), test.restitution)
		}

		if !flops.Eq(c.Friction(), friction) {
			t.Errorf("[%d] slider friction = %v, want %v", i, c.Friction(), friction)
		}
	}
}
