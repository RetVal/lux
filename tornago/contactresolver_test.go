package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestMakeOrthonormal(t *testing.T) {
	var zero glm.Vec3

	var tests = []glm.Vec3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},

		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},

		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},

		{1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},

		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, -1},

		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},

		{0.70710677, 0.70710677, 0},
		{0.70710677, -0.70710677, 0},
		{-0.70710677, 0.70710677, 0},
		{-0.70710677, -0.70710677, 0},

		{0.70710677, 0, 0.70710677},
		{-0.70710677, 0, 0.70710677},
		{0.70710677, 0, -0.70710677},
		{-0.70710677, 0, -0.70710677},

		{0, 0.70710677, 0.70710677},
		{0, -0.70710677, 0.70710677},
		{0, 0.70710677, -0.70710677},
		{0, -0.70710677, -0.70710677},
	}

	for i, test := range tests {
		var x, y, z glm.Vec3
		x = test

		makeOrthonormal(&x, &y, &z)
		xcy := x.Cross(&y)
		xcz := x.Cross(&z)
		ycz := y.Cross(&z)
		v1, v2, v3 := xcy.Cross(&z), xcz.Cross(&y), ycz.Cross(&x)

		if !v1.EqualThreshold(&zero, 1e-4) || !v2.EqualThreshold(&zero, 1e-4) || !v3.EqualThreshold(&zero, 1e-4) {

			t.Logf("[%d] v1 %v=%v", i, v1, zero)
			t.Logf("[%d] v2 %v=%v", i, v2, zero)
			t.Logf("[%d] v3 %v=%v", i, v3, zero)

			t.Errorf("[%d] v1=%v, v2=%v, v3=%v", i, v1, v2, v3)
		}
	}
}

func TestCallback(t *testing.T) {
	var callbacked bool
	body0 := NewRigidBody()
	sphereShape0 := NewCollisionSphere(1)
	body0.SetCollisionShape(sphereShape0)
	body0.SetCallback(func(*RigidBody) {
		callbacked = true
	})

	body1 := NewRigidBody()
	sphereShape1 := NewCollisionSphere(1)
	body1.SetCollisionShape(sphereShape1)

	body0.calculateDerivedData()
	body1.calculateDerivedData()

	w := NewWorld(&NaiveBroadphase{}, ContactResolver{})
	w.AddRigidBody(body0)
	w.AddRigidBody(body1)
	w.Step(1)

	if !callbacked {
		t.Error("callback did not happen")
	}
}
