package tornago

import (
	"github.com/luxengine/lux/glm"
	"math/rand"
	"sync"
	"testing"
)

const (
	benchmarkNumObjects = 1000
	benchmarkWorldSize  = 300.0
)

func TestBroadphases(t *testing.T) {
	rand.Seed(9999)
	const (
		numObjects = 10
		worldsize  = 1.0
	)
	type Object struct {
		body   *RigidBody
		volume *BoundingSphere
	}
	objects := make([]Object, 0, numObjects)

	// just make up a bunch of objects.
	for x := 0; x < cap(objects); x++ {
		var b RigidBody
		volume := BoundingSphere{
			center: glm.Vec3{rand.Float32() * worldsize, rand.Float32() * worldsize, rand.Float32() * worldsize},
			radius: rand.Float32(),
		}
		objects = append(objects, Object{
			body:   &b,
			volume: &volume,
		})
	}

	// find all potential contacts according to overlapping spheres. The
	// algorithm could generate more then these but it must at least generate
	// these because it can't know any better then this.
	expected := make(map[potentialContact]bool)
	for x := 0; x < len(objects); x++ {
		for y := x + 1; y < len(objects); y++ {
			if objects[x].volume.Overlaps(objects[y].volume) {
				expected[potentialContact{
					bodies: [2]*RigidBody{objects[x].body, objects[y].body},
				}] = false
			}
		}
	}
	t.Logf("%v\n", expected)

	//bvh := NewBVHNode(objects[0].body, objects[0].volume)
	broadphases := []struct {
		broadphase Broadphase
		name       string
	}{
		//{&bvh, "BVH"},
		{&NaiveBroadphase{}, "Naive"},
		{&SAP{}, "Non-Persistent sweep and prune"},
	}

	// we hope this is enough
	contacts := make([]potentialContact, len(objects)*100)

	for _, o := range broadphases {
		broadphase := o.broadphase
		for _, object := range objects {
			broadphase.Insert(object.body, object.volume)
		}

		var overdetect int
		gen := broadphase.GeneratePotentialContacts(contacts)

		for x := 0; x < gen; x++ {
			pc := contacts[x]
			pci := potentialContact{
				bodies: [2]*RigidBody{pc.bodies[1], pc.bodies[0]},
			}
			_, ok1 := expected[pc]
			_, ok2 := expected[pci]
			if !ok1 && !ok2 {
				overdetect++
			} else {
				if ok1 {
					expected[pc] = true
				} else {
					expected[pci] = true
				}
			}
		}

		var detected int
		for key, b := range expected {
			if !b {
				t.Errorf("we did not detect %v\n", key)
			} else {
				detected++
			}
		}

		for key := range expected {
			expected[key] = false
		}
		t.Logf("%s detected:  %d/%d\n", o.name, detected, len(expected))
		t.Logf("%s overdetected: %d\n", o.name, overdetect)
	}
}

/*
func BenchmarkBroadphaseBVH(b *testing.B) {
	rand.Seed(9999)
	const (
		numObjects = benchmarkNumObjects
		worldsize  = benchmarkWorldSize
	)
	type Object struct {
		body   *RigidBody
		volume *BoundingSphere
	}
	objects := make([]Object, 0, numObjects)

	// just make up a bunch of objects.
	for x := 0; x < cap(objects); x++ {
		var b RigidBody
		volume := BoundingSphere{
			center: glm.Vec3{rand.Float32() * worldsize, rand.Float32() * worldsize, rand.Float32() * worldsize},
			radius: rand.Float32(),
		}
		objects = append(objects, Object{
			body:   &b,
			volume: &volume,
		})
	}

	for x := 0; x < b.N; x++ {

		broadphase := NewBVHNode(objects[0].body, objects[0].volume)

		for i, object := range objects {
			// we don't need to add the first item its the starting element of
			// the tree
			if i == 0 {
				continue
			}
			broadphase.Insert(object.body, object.volume)
		}
		contactchan := make(chan PotentialContact)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			for range contactchan {
			}
			wg.Done()
		}()

		broadphase.GeneratePotentialContacts(contactchan)
		close(contactchan)
		wg.Wait()
	}
}*/

func BenchmarkBroadphaseNaive(b *testing.B) {
	rand.Seed(9999)
	const (
		numObjects = benchmarkNumObjects
		worldsize  = benchmarkWorldSize
	)
	type Object struct {
		body   *RigidBody
		volume *BoundingSphere
	}
	objects := make([]Object, 0, numObjects)

	// just make up a bunch of objects.
	for x := 0; x < cap(objects); x++ {
		var b RigidBody
		volume := BoundingSphere{
			center: glm.Vec3{rand.Float32() * worldsize, rand.Float32() * worldsize, rand.Float32() * worldsize},
			radius: rand.Float32(),
		}
		objects = append(objects, Object{
			body:   &b,
			volume: &volume,
		})
	}

	contacts := make([]potentialContact, len(objects)*100)

	for x := 0; x < b.N; x++ {
		broadphase := NaiveBroadphase{}

		for _, object := range objects {
			// we don't need to add the first item its the starting element of
			// the tree
			broadphase.Insert(object.body, object.volume)
		}

		broadphase.GeneratePotentialContacts(contacts)
	}
}

func BenchmarkBroadphaseNonPersistentSAP(b *testing.B) {
	rand.Seed(9999)
	const (
		numObjects = benchmarkNumObjects
		worldsize  = benchmarkWorldSize
	)
	type Object struct {
		body   *RigidBody
		volume *BoundingSphere
	}
	objects := make([]Object, 0, numObjects)

	// just make up a bunch of objects.
	for x := 0; x < cap(objects); x++ {
		var b RigidBody
		volume := BoundingSphere{
			center: glm.Vec3{rand.Float32() * worldsize, rand.Float32() * worldsize, rand.Float32() * worldsize},
			radius: rand.Float32(),
		}
		objects = append(objects, Object{
			body:   &b,
			volume: &volume,
		})
	}

	contacts := make([]potentialContact, len(objects)*100)

	for x := 0; x < b.N; x++ {

		broadphase := SAP{}

		for _, object := range objects {
			// we don't need to add the first item its the starting element of
			// the tree
			broadphase.Insert(object.body, object.volume)
		}

		broadphase.GeneratePotentialContacts(contacts)
	}
}

func BenchmarkBroadphase_Fake_NonPersistentSAP(b *testing.B) {
	rand.Seed(9999)
	const (
		numObjects = benchmarkNumObjects
		worldsize  = benchmarkWorldSize
	)
	type Object struct {
		body   *RigidBody
		volume *BoundingSphere
	}
	objects := make([]Object, 0, numObjects)

	// just make up a bunch of objects.
	for x := 0; x < cap(objects); x++ {
		var b RigidBody
		volume := BoundingSphere{
			center: glm.Vec3{rand.Float32() * worldsize, rand.Float32() * worldsize, rand.Float32() * worldsize},
			radius: rand.Float32(),
		}
		objects = append(objects, Object{
			body:   &b,
			volume: &volume,
		})
	}

	broadphase := SAP{}

	for _, object := range objects {
		// we don't need to add the first item its the starting element of
		// the tree
		broadphase.Insert(object.body, object.volume)
	}

	contacts := make([]potentialContact, len(objects)*100)

	for x := 0; x < b.N; x++ {
		broadphase.GeneratePotentialContacts(contacts)
	}
}

//use this as template for testing new broadphases.
func TestSAP(t *testing.T) {
	rand.Seed(9999)
	const (
		numObjects = 10
		worldsize  = 1.0
	)
	type Object struct {
		body   *RigidBody
		volume *BoundingSphere
	}
	objects := make([]Object, 0, numObjects)

	// just make up a bunch of objects.
	for x := 0; x < cap(objects); x++ {
		var b RigidBody
		volume := BoundingSphere{
			center: glm.Vec3{rand.Float32() * worldsize, rand.Float32() * worldsize, rand.Float32() * worldsize},
			radius: rand.Float32(),
		}
		objects = append(objects, Object{
			body:   &b,
			volume: &volume,
		})
	}

	// find all potential contacts according to overlapping spheres. The
	// algorithm could generate more then these but it must at least generate
	// these because it can't know any better then this.
	expected := make(map[potentialContact]bool)
	for x := 0; x < len(objects); x++ {
		for y := x + 1; y < len(objects); y++ {
			if objects[x].volume.Overlaps(objects[y].volume) {
				expected[potentialContact{
					bodies: [2]*RigidBody{objects[x].body, objects[y].body},
				}] = false
			}
		}
	}
	t.Logf("%v\n", expected)

	broadphase := SAP3{}

	for _, object := range objects {
		// we don't need to add the first item its the starting element of
		// the tree
		broadphase.Insert(object.body, object.volume)
	}

	contactchan := make(chan potentialContact)

	var wg sync.WaitGroup
	wg.Add(1)

	var overdetect int
	var detected int

	go func() {
		for contact := range contactchan {
			//t.Logf("detected: %p, %p\n", contact.bodies[0], contact.bodies[1])
			pc := potentialContact{
				bodies: contact.bodies,
			}
			pci := potentialContact{
				bodies: [2]*RigidBody{contact.bodies[1], contact.bodies[0]},
			}
			_, ok1 := expected[pc]
			_, ok2 := expected[pci]
			if !ok1 && !ok2 {
				overdetect++
			} else {
				if ok1 {
					expected[pc] = true
				} else {
					expected[pci] = true
				}
				detected++
			}
		}
		wg.Done()
	}()

	broadphase.GeneratePotentialContacts(contactchan)
	close(contactchan)
	wg.Wait()

	t.Logf("%s detected:  %d/%d\n", "SAP", detected, len(expected))
	t.Logf("%s overdetected: %d\n", "SAP", overdetect)
}

func TestBroadphaseBug(t *testing.T) {
	var r1, r2 RigidBody
	r1.SetRestitution(0.1)
	r2.SetRestitution(0.1)

	r1.SetLinearDamping(1)
	r1.SetAngularDamping(1)

	r2.SetLinearDamping(1)
	r2.SetAngularDamping(1)

	r1.SetMass(1)
	r2.SetMass(0)

	qi := glm.QuatIdent()
	r1.SetOrientationQuat(&qi)
	r2.SetOrientationQuat(&qi)

	r1.SetCollisionShape(NewCollisionSphere(0.5))
	r2.SetCollisionShape(NewCollisionBox(glm.Vec3{5, 5, 5}))

	r1.SetPosition3f(0, 0.1, 0)
	r2.SetPosition3f(0, -5, 0)

	r1.calculateDerivedData()
	r2.calculateDerivedData()

	b := NaiveBroadphase{}

	b.Insert(&r1, r1.shape.GetBoundingVolume())
	b.Insert(&r2, r2.shape.GetBoundingVolume())

	pcontacts := make([]potentialContact, 1)
	gen := b.GeneratePotentialContacts(pcontacts)
	t.Log(pcontacts)
	t.Log(gen)

	d := ContactResolver{}

	contacts := make([]Contact, gen)
	gen = resolvePotentialContacts(pcontacts[:gen], contacts)
	t.Log(gen)
	t.Logf("contacts %+v", contacts[:gen])

	d.ResolveContacts(contacts[:gen], 0.16)

	t.Log("r1 pos", r1.Position())
}
