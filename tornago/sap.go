package tornago

type sapNode struct {
	// Is this the beginning or the end of the rigid body
	start bool

	// Where in the axis is that thing.
	value float32

	// Which body is that representing.
	body *RigidBody
}

// SAP is a sweep and prune broadphase. It only works on 1 axis but this
// particular broadphase should probably not be used in a real application.
type SAP struct {
	axisListX []sapNode
}

// Remove removes that rigid body from the world. Signaling that we do not
// want it to participate in the collision.
func (s *SAP) Remove(b *RigidBody) {
	for i, o := range s.axisListX {
		// Doesn't matter if it's a start or end. Just remove it.
		if o.body == b {
			// Remove the object. However that doesn't free its memory.
			copy(s.axisListX[i:], s.axisListX[i+1:])
			s.axisListX = s.axisListX[:len(s.axisListX)-1]
		}
	}
}

// Insert inserts that node in the SAP.
func (s *SAP) Insert(body *RigidBody, volume *BoundingSphere) {
	s.insertNode(sapNode{
		start: true,
		value: volume.MinX(),
		body:  body,
	})
	s.insertNode(sapNode{
		start: false,
		value: volume.MaxX(),
		body:  body,
	})
}

func (s *SAP) insertNodeAt(n sapNode, i int) {
	s.axisListX = append(s.axisListX, sapNode{})
	copy(s.axisListX[i+1:], s.axisListX[i:])
	s.axisListX[i] = n
}

func (s *SAP) insertNode(n sapNode) {
	if len(s.axisListX) == 0 {
		s.axisListX = append(s.axisListX, n)
		return
	}

	for i, x := range s.axisListX {
		if x.value > n.value {
			s.insertNodeAt(n, i)
			return
		}
	}
	s.axisListX = append(s.axisListX, n)
}

// GeneratePotentialContacts generates all potential contacts with everybody
func (s *SAP) GeneratePotentialContacts(contacts []potentialContact) int {
	var cnt int
	active := make([]*RigidBody, 0, len(s.axisListX)/2)
	for _, n := range s.axisListX {
		// if its the start of an object, check if there are any active objects
		// spawn collisions for all of them and add it to the active list
		if n.start {
			for _, a := range active {
				// if we still have place.
				if cnt < len(contacts) {
					contacts[cnt] = potentialContact{
						bodies: [2]*RigidBody{a, n.body},
					}
					cnt++
				}
			}
			active = append(active, n.body)
		} else { // if its the end of one delete it from the active list
			for i, a := range active {
				if a == n.body {
					//remove it we found it
					copy(active[i:], active[i+1:])
					active = active[:len(active)-1]
					//allright we're out
					break
				}
			}
		}
	}
	return cnt
}

/*
LIST INSERT
list = append(list, nil)
copy(list[i+1:], list[i:])
list[i] = n

LIST REMOVE
copy(list[i:], list[i+1:])
list = list[:len(list)-1]
*/
