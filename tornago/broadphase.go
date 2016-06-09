package tornago

// Broadphase is an algorithm that does a first pass over all rigid bodies and
// detects possible collisions.
type Broadphase interface {
	// Insert take a rigid body and its bounding volume and adds it to the
	// structure. This function is called only once for every rigid body in the
	// world.
	Insert(b *RigidBody, volume *BoundingSphere)

	// Remove removes that rigid body from the world. Signaling that we do not
	// want it to participate in the collision.
	Remove(b *RigidBody)

	// GeneratePotentialContacts generates potential contacts. It can generate
	// false positives (potential collisions that turn out to not be actual
	// collisions) but must not generate false negative (not generating a
	// potential contact between 2 rigid bodies when there actually is a
	// collision between them). returns how many contacts we're actually generated.
	GeneratePotentialContacts(contacts []potentialContact) int
}
