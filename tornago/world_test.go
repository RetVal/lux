package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestWorld_New(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)
	wd, ok := w.dispatcher.(*ContactResolver)
	if !ok {
		t.Error("world dispatcher is not a *ContactResolver")
	}

	if wd != &d {
		t.Error("world dispatcher is not the initial dispatcher")
	}

	wb, ok := w.broadphase.(*SAP)
	if !ok {
		t.Error("world broadphase is not a *SAP")
	}

	if wb != &b {
		t.Error("world broadphase is not the initial broadphase")
	}
}

func TestWorld_AddRigidBody(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)

	bodeh := NewRigidBody()
	bodeh.SetCollisionShape(NewCollisionSphere(1))

	bodeh2 := NewRigidBody()
	bodeh2.SetCollisionShape(NewCollisionSphere(2))

	t.Log(w.broadphase)

	w.AddRigidBody(bodeh)
	if len(w.bodies) != 1 {
		t.Errorf("World should have 1 rigid body in the world: %d", len(w.bodies))
	}

	if len(b.axisListX) != 2 { // remember theres 2 SAP nodes per body
		t.Errorf("SAP should have 1 rigid body in the world: %d", len(b.axisListX))
	}

	w.AddRigidBody(bodeh)

	if len(w.bodies) != 1 {
		t.Errorf("World should still have 1 rigid body in the world: %d", len(w.bodies))
	}

	if len(b.axisListX) != 2 { // remember theres 2 SAP nodes per body
		t.Errorf("SAP should have 1 rigid body in the world: %d", len(b.axisListX))
	}

	w.AddRigidBody(bodeh2)

	if len(w.bodies) != 2 {
		t.Errorf("World should still have 2 rigid body in the world: %d", len(w.bodies))
	}

	if len(b.axisListX) != 4 { // remember theres 2 SAP nodes per body
		t.Errorf("SAP should have 2 rigid body in the world: %d", len(b.axisListX))
	}
}

func TestWorld_RemoveRigidBody(t *testing.T) {
	b := NaiveBroadphase{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)

	bodeh := NewRigidBody()
	bodeh.SetCollisionShape(NewCollisionSphere(1))

	bodeh2 := NewRigidBody()
	bodeh2.SetCollisionShape(NewCollisionSphere(2))

	t.Log(w.broadphase)

	w.AddRigidBody(bodeh)
	w.AddRigidBody(bodeh2)

	w.RemoveRigidBody(bodeh)

	if len(w.bodies) != 1 {
		t.Errorf("World should still have 1 rigid body in the world: %d", len(w.bodies))
	}

	if len(b.objects) != 1 {
		t.Errorf("Broadphase should have 1 rigid body in the world: %d", len(b.objects))
	}

	w.RemoveRigidBody(bodeh)

	if len(w.bodies) != 1 {
		t.Errorf("World should still have 1 rigid body in the world: %d", len(w.bodies))
	}

	if len(b.objects) != 1 {
		t.Errorf("Broadphase should have 1 rigid body in the world: %d", len(b.objects))
	}
}

func TestWorld_AddForceGenerator(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)

	bodeh := NewRigidBody()
	bodeh.SetCollisionShape(NewCollisionSphere(1))

	bodeh2 := NewRigidBody()
	bodeh2.SetCollisionShape(NewCollisionSphere(2))

	earth := glm.Vec3{X: 0, Y: -9.8, Z: 0}
	gravity := NewGravityForceGenerator(&earth)

	w.AddForceGenerator(bodeh, gravity)

	if len(w.forceGeneratorEntries) != 1 {
		t.Errorf("World should contain 1 force generator entry: %d", len(w.forceGeneratorEntries))
	}

	w.AddForceGenerator(bodeh, gravity)

	if len(w.forceGeneratorEntries) != 1 {
		t.Errorf("World should contain 1 force generator entry: %d", len(w.forceGeneratorEntries))
	}

	w.AddForceGenerator(bodeh2, gravity)

	if len(w.forceGeneratorEntries) != 2 {
		t.Errorf("World should contain 2 force generator entry: %d", len(w.forceGeneratorEntries))
	}

}

func TestWorld_RemoveForceGenerator(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)

	bodeh := NewRigidBody()
	bodeh.SetCollisionShape(NewCollisionSphere(1))

	bodeh2 := NewRigidBody()
	bodeh2.SetCollisionShape(NewCollisionSphere(2))

	earth := glm.Vec3{X: 0, Y: -9.8, Z: 0}
	gravity := NewGravityForceGenerator(&earth)

	w.AddForceGenerator(bodeh, gravity)
	w.AddForceGenerator(bodeh2, gravity)

	w.RemoveForceGenerator(bodeh, gravity)

	if len(w.forceGeneratorEntries) != 1 {
		t.Errorf("World should contain 1 force generator entry: %d", len(w.forceGeneratorEntries))
	}

	w.RemoveForceGenerator(bodeh2, gravity)

	if len(w.forceGeneratorEntries) != 0 {
		t.Errorf("World should contain 0 force generator entry: %d", len(w.forceGeneratorEntries))
	}
}

func TestWorld_SetBroadphase(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)

	nb := NaiveBroadphase{}

	bodeh := NewRigidBody()
	bodeh.SetCollisionShape(NewCollisionSphere(1))

	w.AddRigidBody(bodeh)

	w.SetBroadphase(&nb)

	if br := w.Broadphase(); br != &nb {

		t.Errorf("w.Broadphase() = %p, want %p", &br, &nb)
	}

	if len(nb.objects) != 1 {
		t.Errorf("NaiveBroadphase should contain 1 rigid body: %d", len(nb.objects))
	}

}

func TestWorld_Dispatcher(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, nil)

	w.SetDispatcher(&d)

	if di := w.Dispatcher(); di != &d {
		t.Errorf("w.Dispatcher() = %p, want %p", &di, &d)
	}
}

func TestWorld_RayTest(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)

	bodeh := NewRigidBody()
	bodeh.SetCollisionShape(NewCollisionSphere(1))
	bodeh.SetPosition3f(0, 5, 0)

	w.AddRigidBody(bodeh)

	r := NewRay(glm.Vec3{X: -5, Y: 5, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0}, 100)
	rr := RayResultAny{}

	w.RayTest(r, &rr)

	if rr.Body != bodeh {
		t.Errorf("rr.Body = %p, want %p", rr.Body, bodeh)
	}
}

func TestWorld_AddConstraints(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)

	bodeh := NewRigidBody()
	bodeh.SetCollisionShape(NewCollisionSphere(1))

	str := NewStringToWorldConstraint(glm.Vec3{X: 0, Y: 10, Z: 0}, glm.Vec3{}, bodeh, 10, 0)

	str2 := NewStringToWorldConstraint(glm.Vec3{X: 0, Y: 11, Z: 0}, glm.Vec3{}, bodeh, 1, 0)

	if len(w.constraints) != 0 {
		t.Errorf("World should contain 0 constraints: %d", len(w.constraints))
	}

	w.AddConstraint(str)

	if len(w.constraints) != 1 {
		t.Errorf("World should contain 1 constraints: %d", len(w.constraints))
	}

	w.AddConstraint(str)

	if len(w.constraints) != 1 {
		t.Errorf("World should contain 1 constraints: %d", len(w.constraints))
	}

	w.AddConstraint(str2)

	if len(w.constraints) != 2 {
		t.Errorf("World should contain 2 constraints: %d", len(w.constraints))
	}
}

func TestWorld_RemoveConstraints(t *testing.T) {
	b := SAP{}
	d := ContactResolver{}
	w := NewWorld(&b, &d)

	bodeh := NewRigidBody()
	bodeh.SetCollisionShape(NewCollisionSphere(1))

	str := NewStringToWorldConstraint(glm.Vec3{X: 0, Y: 10, Z: 0}, glm.Vec3{}, bodeh, 10, 0)
	str2 := NewStringToWorldConstraint(glm.Vec3{X: 0, Y: 11, Z: 0}, glm.Vec3{}, bodeh, 1, 0)

	w.AddConstraint(str)
	w.AddConstraint(str2)

	w.RemoveConstraint(str)

	if len(w.constraints) != 1 {
		t.Errorf("World should contain 1 constraints: %d", len(w.constraints))
	}

	w.RemoveConstraint(str2)

	if len(w.constraints) != 0 {
		t.Errorf("World should contain 0 constraints: %d", len(w.constraints))
	}
}
