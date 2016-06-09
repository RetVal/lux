package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestSpringForceGenerator_New(t *testing.T) {
	const (
		k = 0.5
		l = 3.0
	)
	p1, p2 := glm.Vec3{1, 2, 3}, glm.Vec3{4, 5, 6}
	var b RigidBody
	g := NewSpringForceGenerator(&p1, &b, &p2, k, l)
	if g.localPoint != p1 {
		t.Errorf("local point = %v, want %v", g.localPoint, p1)
	}
	if g.otherPoint != p2 {
		t.Errorf("local point = %v, want %v", g.otherPoint, p2)
	}
	if g.other != &b {
		t.Errorf("other = %p, want %p", g.other, &b)
	}
	if g.k != k {
		t.Errorf("k = %f, want %f", g.k, k)
	}
	if g.l != l {
		t.Errorf("l = %f, want %f", g.l, l)
	}
}

func TestSpringForceGenerator_SettersAndGetters(t *testing.T) {
	const (
		k = 0.5
		l = 3.0
	)
	p1, p2 := glm.Vec3{0, 0, 0}, glm.Vec3{0, 0, 0}
	g := NewSpringForceGenerator(&p1, nil, &p2, 0, 0)
	g.SetSpringConstant(k)
	g.SetRestLength(l)
	if g.k != k {
		t.Errorf("SetSpringConstant() = %f, want %f", g.k, k)
	}
	if g.l != l {
		t.Errorf("SetRestLength() = %f, want %f", g.l, l)
	}
	if gk := g.SpringConstant(); gk != k {
		t.Errorf("SpringConstant() = %f, want %f", gk, k)
	}
	if gl := g.RestLength(); gl != l {
		t.Errorf("RestLength() = %f, want %f", gl, l)
	}

	g.SetLocalPoint3f(1, 2, 3)
	if g.localPoint != (glm.Vec3{1, 2, 3}) {
		t.Errorf("SetLocalPoint3f = %v, want %v", g.localPoint, (glm.Vec3{1, 2, 3}))
	}
	lp := glm.Vec3{4, 5, 6}
	g.SetLocalPointVec(&lp)
	if g.localPoint != lp {
		t.Errorf("SetLocalPointVec = %v, want %v", g.localPoint, lp)
	}

	if glp := g.LocalPoint(); glp != lp {
		t.Errorf("LocalPoint = %v, want %v", glp, lp)
	}

	g.SetOtherPoint3f(1, 2, 3)
	if g.otherPoint != (glm.Vec3{1, 2, 3}) {
		t.Errorf("SetOtherPoint3f = %v, want %v", g.otherPoint, (glm.Vec3{1, 2, 3}))
	}
	op := glm.Vec3{4, 5, 6}
	g.SetOtherPointVec(&op)
	if g.otherPoint != op {
		t.Errorf("SetOtherPointVec = %v, want %v", g.otherPoint, op)
	}

	if gop := g.OtherPoint(); gop != op {
		t.Errorf("OtherPoint = %v, want %v", gop, op)
	}

	var b RigidBody
	g.SetTargetRigidBody(&b)
	if g.other != &b {
		t.Errorf("SetTargetRigidBody() = %p, want %p", g.other, &b)
	}
	if gb := g.TargetRigidBody(); gb != &b {
		t.Errorf("TargetRigidBody() = %p, want %p", gb, &b)
	}
}

func TestSpringForceGenerator_UpdateForce_InfiniteMass(t *testing.T) {
	var b RigidBody
	var target RigidBody

	pos := glm.Vec3{99, 99, 99}
	b.SetLinearDamping(1)
	b.SetAngularDamping(1)
	b.SetMass(0)
	b.SetPositionVec3(&pos)

	var p1, p2 glm.Vec3

	g := NewSpringForceGenerator(&p1, &target, &p2, 999, 0)
	for x := 0; x < 4; x++ {
		g.UpdateForce(&b, 1)
	}
	if bpos := b.Position(); bpos != pos {
		t.Errorf("UpdateForce = %v, want %v", bpos, pos)
	}
}

func TestSpringForceGenerator_UpdateForce(t *testing.T) {
	var b RigidBody
	var target RigidBody

	pos := glm.Vec3{1, 0, 0}
	b.SetLinearDamping(0.95)
	b.SetAngularDamping(0.95)
	b.SetMass(1)
	b.SetPositionVec3(&pos)

	target.SetLinearDamping(0.95)
	target.SetAngularDamping(0.95)
	target.SetMass(1)

	b.calculateDerivedData()
	target.calculateDerivedData()

	var p1, p2 glm.Vec3

	g := NewSpringForceGenerator(&p1, &target, &p2, 0.5, 1.1)
	for x := 0; x < 4000; x++ {
		g.UpdateForce(&b, 1)
		b.Integrate(1)
	}

	expected := glm.Vec3{1.1, 0, 0}
	if bpos := b.Position(); !bpos.EqualThreshold(&expected, 1e-4) {
		t.Errorf("UpdateForce position = %v, want %v", bpos, expected)
	}
}

func BenchmarkSpringForceGenerator_UpdateForce(b *testing.B) {
	var body RigidBody
	var target RigidBody

	pos := glm.Vec3{1, 0, 0}
	body.SetLinearDamping(0.95)
	body.SetAngularDamping(0.95)
	body.SetMass(1)
	body.SetPositionVec3(&pos)

	target.SetLinearDamping(0.95)
	target.SetAngularDamping(0.95)
	target.SetMass(1)

	body.calculateDerivedData()
	target.calculateDerivedData()

	var p1, p2 glm.Vec3

	g := NewSpringForceGenerator(&p1, &target, &p2, 0.5, 1.1)
	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		g.UpdateForce(&body, 1)
	}

}
