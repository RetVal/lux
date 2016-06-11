package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"github.com/luxengine/lux/math"
	"testing"
)

// check at compile time that they satisfy the interface
var _ ForceGenerator = &AttractionSphere{}
var _ ForceGenerator = &AttractionCylinder{}

func TestAttractionSphere_UpdateForce(t *testing.T) {
	tests := []struct {
		body          *RigidBody
		sphere        *AttractionSphere
		expectedForce glm.Vec3
	}{
		{
			body: func() *RigidBody {
				b := NewRigidBody()
				b.SetMass(1)
				b.calculateDerivedData()
				return b
			}(),
			sphere: &AttractionSphere{
				Force:  1,
				Center: glm.Vec3{X: 0, Y: -1, Z: 0},
			},
			expectedForce: glm.Vec3{X: 0, Y: -1, Z: 0},
		},
		{
			body: func() *RigidBody {
				b := NewRigidBody()
				b.SetMass(1)
				b.calculateDerivedData()
				return b
			}(),
			sphere: &AttractionSphere{
				Force:  -1,
				Center: glm.Vec3{X: 0, Y: -1, Z: 0},
			},
			expectedForce: glm.Vec3{X: 0, Y: 1, Z: 0},
		},
		{
			body: NewRigidBody(),
			sphere: &AttractionSphere{
				Force:  1,
				Center: glm.Vec3{},
			},
			expectedForce: glm.Vec3{},
		},
		{
			body: NewRigidBody(),
			sphere: &AttractionSphere{
				Force:  math.NaN(),
				Center: glm.Vec3{X: 0, Y: 1, Z: 0},
			},
			expectedForce: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
		},
		{
			body: NewRigidBody(),
			sphere: &AttractionSphere{
				Force:  1,
				Center: glm.Vec3{X: math.NaN(), Y: 1, Z: 0},
			},
			expectedForce: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
		},
		{
			body: NewRigidBody(),
			sphere: &AttractionSphere{
				Force:  1,
				Center: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			},
			expectedForce: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
		},
	}
	for i, test := range tests {
		if i != 2 {
			continue
		}
		test.sphere.UpdateForce(test.body, 0)
		fa := test.body.forceAccumulator
		ex := test.expectedForce
		if !glmtesting.Vec3Equal(fa, ex) {
			t.Errorf("[%d] force = %s, want %s", i, fa.String(), ex.String())
		}
	}
}
