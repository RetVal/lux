package tornago

// naiveBroadphaseEntry is the type that naive broadphases have a list of. It
// holds all the objects that will be checked during the broadphase.
type naiveBroadphaseEntry struct {
	body   *RigidBody
	volume *BoundingSphere
}

// NaiveBroadphase is literally just a list and the contact generation is
// testing every possible combinations.
type NaiveBroadphase struct {
	objects []naiveBroadphaseEntry
}

// Insert inserts this rigid body in this node of one of it's childs.
func (n *NaiveBroadphase) Insert(b *RigidBody, volume *BoundingSphere) {
	// find if we have it before.
	var found bool
	for _, o := range n.objects {
		if b == o.body {
			found = true
			break
		}
	}

	// remove that guy if we dont already have it.
	if !found {
		n.objects = append(n.objects, naiveBroadphaseEntry{
			body:   b,
			volume: volume,
		})
	}
}

// GeneratePotentialContacts sends all colliding bounding sphere to the narrow
// phase detector. Should run in O(n^2).
func (n *NaiveBroadphase) GeneratePotentialContacts(contacts []potentialContact) int {
	var cnt int
	// for every pair.
	for x := 0; x < len(n.objects); x++ {
		for y := x + 1; y < len(n.objects); y++ {
			// if we still have place.
			if cnt < len(contacts) {
				// if they overlap.
				if n.objects[x].volume.Overlaps(n.objects[y].volume) {
					// we add it.
					contacts[cnt] = potentialContact{
						bodies: [2]*RigidBody{n.objects[x].body, n.objects[y].body},
					}
					cnt++
				}
			}
		}
	}
	return cnt
}

// Remove removes that rigid body from the world. Signaling that we do not
// want it to participate in the collision.
func (n *NaiveBroadphase) Remove(b *RigidBody) {
	for i, o := range n.objects {
		if o.body == b {
			// Remove the object. However that doesn't free its memory.
			copy(n.objects[i:], n.objects[i+1:])
			n.objects = n.objects[:len(n.objects)-1]
		}
	}
}
