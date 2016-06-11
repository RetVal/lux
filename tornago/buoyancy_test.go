package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"github.com/luxengine/lux/math"
	"testing"
)

var _ ForceGenerator = &BuoyancyForceGenerator{}

func TestBuoyancyForceGenerator_UpdateForce(t *testing.T) {
	tests := []struct {
		body     *RigidBody
		buoyancy *BuoyancyForceGenerator
		expected glm.Vec3
	}{
		{
			body: func() *RigidBody {
				b := NewRigidBody()
				b.SetMass(1)
				b.SetPosition3f(0, -1, 0)
				b.calculateDerivedData()
				return b
			}(),
			buoyancy: &BuoyancyForceGenerator{
				Height:   0,
				MaxDepth: 1,
				Volume:   1,
				Density:  1000,
			},
			expected: glm.Vec3{X: 0, Y: 1000, Z: 0},
		},
		{
			body: func() *RigidBody {
				b := NewRigidBody()
				b.SetMass(1)
				b.SetPosition3f(0, -1, 0)
				b.calculateDerivedData()
				return b
			}(),
			buoyancy: &BuoyancyForceGenerator{
				Height:   math.NaN(),
				MaxDepth: 1,
				Volume:   1,
				Density:  1000,
			},
			expected: glm.Vec3{X: 0, Y: math.NaN(), Z: 0},
		},
		{
			body: func() *RigidBody {
				b := NewRigidBody()
				b.SetMass(1)
				b.SetPosition3f(0, -1, 0)
				b.calculateDerivedData()
				return b
			}(),
			buoyancy: &BuoyancyForceGenerator{
				Height:   0,
				MaxDepth: math.NaN(),
				Volume:   1,
				Density:  1000,
			},
			expected: glm.Vec3{X: 0, Y: math.NaN(), Z: 0},
		},
		{
			body: func() *RigidBody {
				b := NewRigidBody()
				b.SetMass(1)
				b.SetPosition3f(0, -1, 0)
				b.calculateDerivedData()
				return b
			}(),
			buoyancy: &BuoyancyForceGenerator{
				Height:   0,
				MaxDepth: 1,
				Volume:   math.NaN(),
				Density:  1000,
			},
			expected: glm.Vec3{X: 0, Y: math.NaN(), Z: 0},
		},
		{
			body: func() *RigidBody {
				b := NewRigidBody()
				b.SetMass(1)
				b.SetPosition3f(0, -1, 0)
				b.calculateDerivedData()
				return b
			}(),
			buoyancy: &BuoyancyForceGenerator{
				Height:   0,
				MaxDepth: 1,
				Volume:   1,
				Density:  math.NaN(),
			},
			expected: glm.Vec3{X: 0, Y: math.NaN(), Z: 0},
		},
	}
	for i, test := range tests {
		test.buoyancy.UpdateForce(test.body, 0)
		fa := test.body.forceAccumulator
		ex := test.expected
		if !glmtesting.Vec3Equal(fa, ex) {
			t.Errorf("[%d] force = %s, want %s", i, fa.String(), ex.String())
		}
	}
}
