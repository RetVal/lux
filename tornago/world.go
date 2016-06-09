package tornago

type forceGeneratorEntry struct {
	body           *RigidBody
	forceGenerator ForceGenerator
}

// World is the parent structure of a tornago instance, you interract with
// everything via a world. It constains all the simulated rigid bodies,
// constraints and force generators.
type World struct {
	// The broadphase this world uses.
	broadphase Broadphase

	// The dispatcher this world uses.
	dispatcher Dispatcher

	// The rigid bodies that we want to simulate.
	bodies []*RigidBody

	// All the constraints in the world.
	constraints []Constraint

	// All the force generator entries in the world.
	forceGeneratorEntries []forceGeneratorEntry
}

// NewWorld generates a new world with the given Broadphase and Dispatcher.
func NewWorld(broadphase Broadphase, dispatcher Dispatcher) *World {
	var w World
	w.New(broadphase, dispatcher)
	return &w
}

// New initialises this world with the given arguments. This is used for memory
// management.
func (w *World) New(broadphase Broadphase, dispatcher Dispatcher) {
	w.broadphase = broadphase
	w.dispatcher = dispatcher
}

// AddRigidBody adds the given rigid body to the world.
func (w *World) AddRigidBody(body *RigidBody) {
	var found bool
	for _, b := range w.bodies {
		if b == body {
			found = true
		}
	}
	if !found {
		w.bodies = append(w.bodies, body)
		w.broadphase.Insert(body, body.shape.GetBoundingVolume())
	}
}

// RemoveRigidBody removes the rigid body from the world.
func (w *World) RemoveRigidBody(body *RigidBody) {
	for i, b := range w.bodies {
		if b == body {
			copy(w.bodies[i:], w.bodies[i+1:])
			w.bodies = w.bodies[:len(w.bodies)-1]
			w.broadphase.Remove(body)
			return
		}
	}
}

// AddForceGenerator adds the given force generator to the world.
func (w *World) AddForceGenerator(body *RigidBody, forceGenerator ForceGenerator) {
	var found bool
	for _, f := range w.forceGeneratorEntries {
		if f.body == body && f.forceGenerator == forceGenerator {
			found = true
		}
	}
	if !found {
		w.forceGeneratorEntries = append(w.forceGeneratorEntries, forceGeneratorEntry{
			body:           body,
			forceGenerator: forceGenerator,
		})
	}
}

// RemoveForceGenerator removes this {force generator, body} from the world.
func (w *World) RemoveForceGenerator(body *RigidBody, forceGenerator ForceGenerator) {
	for i, e := range w.forceGeneratorEntries {
		if e.body == body && e.forceGenerator == forceGenerator {
			copy(w.forceGeneratorEntries[i:], w.forceGeneratorEntries[i+1:])
			w.forceGeneratorEntries = w.forceGeneratorEntries[:len(w.forceGeneratorEntries)-1]
			return
		}
	}
}

// SetBroadphase sets the broadphase to use.
func (w *World) SetBroadphase(broadphase Broadphase) {
	w.broadphase = broadphase
	for _, body := range w.bodies {
		w.broadphase.Insert(body, body.shape.GetBoundingVolume())
	}
}

// Broadphase returns this world current broadphase.
func (w *World) Broadphase() Broadphase {
	return w.broadphase
}

// SetDispatcher sets the dispatcher to use.
func (w *World) SetDispatcher(dispatcher Dispatcher) {
	w.dispatcher = dispatcher
}

// Dispatcher returns this world current dispatcher.
func (w *World) Dispatcher() Dispatcher {
	return w.dispatcher
}

// RayTest casts a ray in the world a calls RayResult.AddResult for every object
// hit.
func (w *World) RayTest(ray Ray, result RayResult) {
	for _, body := range w.bodies {
		body.shape.RayTest(ray, result)
	}
}

// Step steps the world forward in time by the given time amount.
func (w *World) Step(duration float32) {

	// iterate over the force generators.
	for _, e := range w.forceGeneratorEntries {
		e.forceGenerator.UpdateForce(e.body, duration)
	}

	// integrate all the rigid bodies.
	for _, b := range w.bodies {
		b.Integrate(duration)
	}

	// TODO: the broadphase needs to be updated every frame but for now we'll
	// just rebuild it.
	w.broadphase = &NaiveBroadphase{}
	for _, b := range w.bodies {
		w.broadphase.Insert(b, b.shape.GetBoundingVolume())
	}

	pcontacts := make([]potentialContact, len(w.bodies))
	gen := w.broadphase.GeneratePotentialContacts(pcontacts)

	contacts := make([]Contact, gen+10) // 10 is just a buffer
	gen = resolvePotentialContacts(pcontacts[:gen], contacts)

	for _, constraint := range w.constraints {
		if gen >= len(contacts) {
			break
		}

		n := constraint.GenerateContacts(contacts[gen:])
		gen += n
	}

	w.dispatcher.ResolveContacts(contacts[:gen], duration)
}

// AddConstraint adds a constraint to the world.
func (w *World) AddConstraint(constraint Constraint) {
	var found bool
	for _, c := range w.constraints {
		if c == constraint {
			found = true
		}
	}
	if !found {
		w.constraints = append(w.constraints, constraint)
	}
}

// RemoveConstraint removes a constraint from the world.
func (w *World) RemoveConstraint(constraint Constraint) {
	for i, c := range w.constraints {
		if c == constraint {
			copy(w.constraints[i:], w.constraints[i+1:])
			w.constraints = w.constraints[:len(w.constraints)-1]
		}
	}
}
