package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
	"math/rand"
	"testing"
)

func TestRigidBody_SetPosition3f(t *testing.T) {
	var b RigidBody
	b.SetPosition3f(1, 2, 3)
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	if b.position != want {
		t.Errorf("SetPosition3f = %v, want %v", b.position, want)
		return
	}
}

func TestRigidBody_SetPositionVec(t *testing.T) {
	var b RigidBody
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	b.SetPositionVec3(&want)
	if b.position != want {
		t.Errorf("SetPositionVec = %v, want %v", b.position, want)
		return
	}
}

func TestRigidBody_Position(t *testing.T) {
	var b RigidBody
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	b.SetPositionVec3(&want)
	if pos := b.Position(); pos != want {
		t.Errorf("Position = %v, want %v", pos, want)
		return
	}
}

func TestRigidBody_SetOrientation4f(t *testing.T) {
	var b RigidBody
	q := glm.Quat{W: 1, Vec3: glm.Vec3{X: 2, Y: 3, Z: 4}}
	b.SetOrientation4f(q.W, q.X, q.Y, q.Z)
	if res := b.orientation; res != q {
		t.Errorf("SetOrientation4f = %v, want %v", res, q)
		return
	}
}

func TestRigidBody_SetOrientationQuat(t *testing.T) {
	var b RigidBody
	q := glm.Quat{W: 1, Vec3: glm.Vec3{X: 2, Y: 3, Z: 4}}
	b.SetOrientationQuat(&q)
	if res := b.orientation; res != q {
		t.Errorf("SetOrientationQuat = %v, want %v", res, q)
		return
	}
}

func TestRigidBody_Orientation(t *testing.T) {
	var b RigidBody

	q := glm.Quat{W: 1, Vec3: glm.Vec3{X: 2, Y: 3, Z: 4}}
	b.SetOrientationQuat(&q)

	if ori := b.Orientation(); ori != q {
		t.Errorf("Orientation = %v, want %v", ori, q)
		return
	}
}

func TestRigidBody_SetVelocity3f(t *testing.T) {
	var b RigidBody
	b.SetVelocity3f(1, 2, 3)
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	if b.velocity != want {
		t.Errorf("SetVelocity3f = %v, want %v", b.velocity, want)
		return
	}
}

func TestRigidBody_SetVelocityVec(t *testing.T) {
	var b RigidBody
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	b.SetVelocityVec3(&want)
	if b.velocity != want {
		t.Errorf("SetVelocityVec = %v, want %v", b.velocity, want)
		return
	}
}

func TestRigidBody_Velocity(t *testing.T) {
	var b RigidBody
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	b.SetVelocityVec3(&want)
	if b.Velocity() != want {
		t.Errorf("Velocity = %v, want %v", b.Velocity(), want)
		return
	}
}

func TestRigidBody_SetAcceleration3f(t *testing.T) {
	var b RigidBody
	b.SetAcceleration3f(1, 2, 3)
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	if b.acceleration != want {
		t.Errorf("SetAcceleration3f = %v, want %v", b.acceleration, want)
		return
	}
}

func TestRigidBody_SetAccelerationVec(t *testing.T) {
	var b RigidBody
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	b.SetAccelerationVec3(&want)
	if b.acceleration != want {
		t.Errorf("SetAccelerationVec = %v, want %v", b.acceleration, want)
		return
	}
}

func TestRigidBody_Acceleration(t *testing.T) {
	var b RigidBody
	want := glm.Vec3{X: 1, Y: 2, Z: 3}
	b.SetAccelerationVec3(&want)
	if b.Acceleration() != want {
		t.Errorf("Acceleration = %v, want %v", b.Acceleration(), want)
		return
	}
}

func TestRigidBody_SetMass(t *testing.T) {
	var b RigidBody
	//nan +inf, -inf, normal
	normal := float32(0.5)
	b.SetMass(normal)
	if b.inverseMass != 1.0/normal {
		t.Errorf("SetMass = %v, want %v", b.inverseMass, 1.0/normal)
		return
	}

	b.SetMass(math.Inf(1))
	if b.inverseMass != 0 {
		t.Errorf("SetMass = %v, want %v", b.inverseMass, 1.0/normal)
		return
	}

	b.SetMass(0)
	if b.inverseMass != 0 {
		t.Errorf("setting mass not acting as expected, got: %+v, wanted %+v", b.inverseMass, 1.0/normal)
		return
	}

	readyForNaN := true
	defer func() {
		recover()
		if !readyForNaN {
			t.Errorf("NaN not acting as expected")
		}
	}()
	b.SetMass(math.NaN())
}

func TestRigidBody_Mass(t *testing.T) {
	var b RigidBody
	//nan +inf, -inf, normal
	normal := float32(0.5)
	b.SetMass(normal)
	if b.Mass() != normal {
		t.Errorf("get mass not acting as expected, got: %+v, wanted %+v", b.Mass(), normal)
		return
	}

	b.SetMass(math.Inf(1))
	if b.Mass() != 0 {
		t.Errorf("get mass not acting as expected, got: %+v, wanted %+v", b.Mass(), 1.0/normal)
		return
	}

	b.SetMass(0)
	if b.Mass() != 0 {
		t.Errorf("get mass not acting as expected, got: %+v, wanted %+v", b.Mass(), 1.0/normal)
		return
	}
}

func TestRigidBody_HasFiniteMass(t *testing.T) {
	var b RigidBody
	b.SetMass(0.5)
	if !b.HasFiniteMass() {
		t.Errorf("0.5 is not a finite mass ?")
		return
	}

	b.SetMass(0)

	if b.HasFiniteMass() {
		t.Errorf("0 is a finite mass ?")
		return
	}

	b.SetMass(math.Inf(1))

	if b.HasFiniteMass() {
		t.Errorf("0 is a finite mass ?")
		return
	}
}

func TestRigidBody_InverseMass(t *testing.T) {
	var tests = []struct {
		mass, expected float32
	}{
		{1, 1},
		{math.Inf(1), 0},
		{0, 0},
	}

	for _, test := range tests {
		var b RigidBody
		b.SetMass(test.mass)
		if imass := b.InverseMass(); imass != test.expected {
			t.Errorf("Mass(%f), InverseMass() = %f, want %f", test.mass, imass, test.expected)
		}
	}
}

func TestRigidBody_SetLinearDamping(t *testing.T) {
	var b RigidBody
	damping := float32(1)
	b.SetLinearDamping(damping)
	if b.linearDamping != damping {
		t.Errorf("SetLinearDamping() = %v, want %v", b.linearDamping, damping)
		return
	}
}

func TestRigidBody_LinearDamping(t *testing.T) {
	var b RigidBody
	damping := float32(1)
	b.SetLinearDamping(damping)
	if d := b.LinearDamping(); d != damping {
		t.Errorf("LinearDamping() = %v, want %v", d, damping)
		return
	}
}

func TestRigidBody_SetAngularDamping(t *testing.T) {
	var b RigidBody
	damping := float32(1)
	b.SetAngularDamping(damping)
	if b.angularDamping != damping {
		t.Errorf("SetAngularDamping() = %v, want %v", b.angularDamping, damping)
		return
	}
}

func TestRigidBody_AngularDamping(t *testing.T) {
	var b RigidBody
	damping := float32(1)
	b.SetAngularDamping(damping)
	if d := b.AngularDamping(); d != damping {
		t.Errorf("AngularDamping() = %v, want %v", d, damping)
		return
	}
}

func TestRigidBody_InterniaTensor(t *testing.T) {
	var b RigidBody
	sit := sphereInertiaTensor(1, 1)
	b.SetInertiaTensor(&sit)
	if it := b.InertiaTensor(); it != sit {
		t.Errorf("InertiaTensor() = %v, want %v", it, sit)
	}
}

func TestRigidBody_InterniaTensorIn(t *testing.T) {
	var b RigidBody
	sit := sphereInertiaTensor(1, 1)
	b.SetInertiaTensor(&sit)
	var it glm.Mat3
	b.InertiaTensorIn(&it)
	if it != sit {
		t.Errorf("InertiaTensor() = %v, want %v", it, sit)
	}
}

func TestRigidBody_OpenGLMatrix(t *testing.T) {
	var b RigidBody
	q := glm.Quat{W: 1, Vec3: glm.Vec3{X: 2, Y: 3, Z: 4}}
	q.Normalize()
	m := q.Mat4()
	pos := glm.Vec3{X: 5, Y: 6, Z: 7}
	trans := glm.Translate3D(pos.X, pos.Y, pos.Z)
	trans.Mul4With(&m)

	m = trans

	b.SetOrientationQuat(&q)
	b.SetPositionVec3(&pos)
	b.calculateDerivedData()

	var ret glm.Mat4
	b.OpenGLMatrix(&ret)
	if ret != m {
		t.Errorf("OpenGLMatrix() = %v, want %v", ret, m)
	}
}

func TestRigidBody_AddForce(t *testing.T) {
	f1, f2 := glm.Vec3{X: 0, Y: 1, Z: 2}, glm.Vec3{X: 4, Y: 5, Z: 6}
	var b RigidBody
	b.AddForce(&f1)
	b.AddForce(&f2)

	expected := f1.Add(&f2)

	if b.forceAccumulator != expected {
		t.Errorf("AddForce(%v, %v) = %v, want %v", f1, f2, b.forceAccumulator, expected)
	}
}

func TestRigidBody_AddTorque(t *testing.T) {
	f1, f2 := glm.Vec3{X: 0, Y: 1, Z: 2}, glm.Vec3{X: 4, Y: 5, Z: 6}
	var b RigidBody
	b.AddTorque(&f1)
	b.AddTorque(&f2)

	expected := f1.Add(&f2)

	if b.torqueAccumulator != expected {
		t.Errorf("AddTorque(%v, %v) = %v, want %v", f1, f2, b.forceAccumulator, expected)
	}
}

func TestRigidBody_AddForceAtPoint(t *testing.T) {
	force, point := glm.Vec3{X: 0, Y: 10, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0}
	var b RigidBody
	b.AddForceAtPoint(&force, &point)
	b.AddForceAtPoint(&force, &point)
	ef, et := glm.Vec3{X: 0, Y: 20, Z: 0}, glm.Vec3{X: 0, Y: 0, Z: 20}
	if f, to := b.forceAccumulator, b.torqueAccumulator; !to.EqualThreshold(&et, 1e-4) || !f.EqualThreshold(&ef, 1e-4) {
		t.Errorf("forceAcc, torqueAcc = %v, %v, want %v, %v", f, to, ef, et)
		return
	}
}

func TestRigidBody_AddForceAtBodyPoint(t *testing.T) {
	force, point := glm.Vec3{X: 0, Y: 10, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0}

	q := glm.Quat{W: 1, Vec3: glm.Vec3{X: 2, Y: 3, Z: 4}}
	pos := glm.Vec3{X: 5, Y: 6, Z: 7}
	q.Normalize()
	var b RigidBody
	b.SetOrientationQuat(&q)
	b.SetPositionVec3(&pos)
	b.AddForceAtBodyPoint(&force, &point)
	b.AddForceAtBodyPoint(&force, &point)

	ef, et := glm.Vec3{X: 0, Y: 20, Z: 0}, glm.Vec3{X: 140, Y: 0, Z: -100}
	if fo, to := b.forceAccumulator, b.torqueAccumulator; !ef.EqualThreshold(&fo, 1e-4) || !et.EqualThreshold(&to, 1e-4) {
		t.Errorf("forceAcc, torqueAcc = %v, %v, want %v, %v", fo, to, ef, et)
		return
	}
}

func TestRigidBody_PointInWorldCoordinates(t *testing.T) {
	q := glm.Quat{W: 1, Vec3: glm.Vec3{X: 2, Y: 3, Z: 4}}
	pos := glm.Vec3{X: 5, Y: 6, Z: 7}
	q.Normalize()
	var b RigidBody
	b.SetOrientationQuat(&q)
	b.SetPositionVec3(&pos)
	b.calculateDerivedData()

	point := glm.Vec3{X: 1, Y: 0, Z: 0}
	var dst glm.Vec3

	b.PointInWorldCoordinates(&point, &dst)

	expected := glm.Vec3{X: 4.3333335, Y: 6.6666665, Z: 7.3333335}

	if !dst.EqualThreshold(&expected, 1e-4) {
		t.Errorf("world = %v, want %v", dst, expected)
		return
	}
}

func TestRigidBody_Integrate_LinearOnly(t *testing.T) {
	const (
		mass      = 1
		radius    = 1
		expectedY = -5
	)
	var b RigidBody
	b.linearDamping = 1
	b.angularDamping = 0
	b.SetMass(mass)
	tensor := sphereInertiaTensor(mass, radius)
	b.SetInertiaTensor(&tensor)
	b.acceleration = glm.Vec3{X: 0, Y: -10, Z: 0}
	for x := 0; x < 10; x++ {
		b.Integrate(0.1)
	}
	if b.position.Y > expectedY {
		t.Errorf("linear integration did not work as planned, got %v, want %v", b.position, glm.Vec3{X: 0, Y: expectedY, Z: 0})
		return
	}
}

func TestRigidBody_Integrate_AngularOnly(t *testing.T) {
	const (
		mass      = 1
		radius    = 1
		expectedY = -5
	)
	var b RigidBody
	b.linearDamping = 0
	b.angularDamping = 1
	b.SetMass(mass)
	b.orientation.Iden()
	tensor := sphereInertiaTensor(mass, radius)
	b.SetInertiaTensor(&tensor)
	force := glm.Vec3{X: 0, Y: 1, Z: 0}
	point := glm.Vec3{X: 1, Y: 0, Z: 0}
	for x := 0; x < 10; x++ {
		b.AddForceAtPoint(&force, &point)
		b.Integrate(0.1)
		//d(b.orientation)
	}
}

func TestRigidBody_Integrate_Both(t *testing.T) {
	const (
		mass   = 1
		radius = 1
	)
	var b RigidBody
	b.linearDamping = 1
	b.angularDamping = 1
	b.SetMass(mass)
	b.orientation.Iden()
	tensor := sphereInertiaTensor(mass, radius)
	b.SetInertiaTensor(&tensor)
	force := glm.Vec3{X: 0, Y: 1, Z: 0}
	point := glm.Vec3{X: 1, Y: 0, Z: 0}
	for x := 0; x < 10; x++ {
		b.AddForceAtPoint(&force, &point)
		b.Integrate(0.1)
	}
	var m glm.Mat4
	b.OpenGLMatrix(&m)
	expected := glm.Mat4{0.43354082, 0.9011339, 0, 0, -0.9011339, 0.43354082, 0, 0, 0, 0, 1, 0, 0, 0.5500001, 0, 1}
	if !m.EqualThreshold(&expected, 1e-4) {
		t.Errorf("Integrate = %v, want %v", m, expected)
	}
}

func BenchmarkRigidBody_Integrate(b *testing.B) {
	var body RigidBody
	body.SetMass(2)
	tensor := sphereInertiaTensor(2, 1)
	body.SetInertiaTensor(&tensor)
	body.SetLinearDamping(0.95)
	body.SetAngularDamping(0.95)
	body.calculateDerivedData()
	b.ResetTimer()
	b.StopTimer()
	for x := 0; x < b.N; x++ {
		force, point := glm.Vec3{X: rand.Float32(), Y: rand.Float32(), Z: rand.Float32()}, glm.Vec3{X: rand.Float32(), Y: rand.Float32(), Z: rand.Float32()}
		body.AddForceAtBodyPoint(&force, &point)
		b.StartTimer()
		body.Integrate(0.1)
		b.StopTimer()
	}
}

func BenchmarkRigidBody_calculateDerivedData(b *testing.B) {
	var body RigidBody
	position := glm.Vec3{X: 1, Y: 2, Z: 3}
	orientation := glm.Quat{W: 1, Vec3: glm.Vec3{X: 2, Y: 3, Z: 4}}
	body.SetPositionVec3(&position)
	body.SetOrientationQuat(&orientation)
	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		body.calculateDerivedData()
	}
}
