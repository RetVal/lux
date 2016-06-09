package tornago

// ForceGenerator is an interface to model every force generators.
type ForceGenerator interface {
	//calculate and update the force applied to the given rigid body.
	UpdateForce(p *RigidBody, duration float32)
}
