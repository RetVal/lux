package tornago

type sap3Node struct {
	// Is this the beginning or the end of the rigid body
	start bool

	// Where in the axis is that thing.
	value float32

	// Which body is that representing.
	body *RigidBody
}

// SAP3 is a sweep and prune broadphase
type SAP3 struct {
	axisList [3][]sapNode
}

// Insert inserts that node in the SAP.
func (s *SAP3) Insert(body *RigidBody, volume *BoundingSphere) {
	s.insertNode(sapNode{start: true, value: volume.MinX(), body: body}, 0)
	s.insertNode(sapNode{start: false, value: volume.MaxX(), body: body}, 0)

	s.insertNode(sapNode{start: true, value: volume.MinY(), body: body}, 1)
	s.insertNode(sapNode{start: false, value: volume.MaxY(), body: body}, 1)

	s.insertNode(sapNode{start: true, value: volume.MinZ(), body: body}, 2)
	s.insertNode(sapNode{start: false, value: volume.MaxZ(), body: body}, 2)
}

// Remove removes this rigid body from the broadphase. It will no longer be used
// in the simulation.
func (s *SAP3) Remove(body *RigidBody, volume *BoundingSphere) {
	s.remove(body, 0)
	s.remove(body, 1)
	s.remove(body, 2)
}

// remove removes the given body from the specified axis list.
func (s *SAP3) remove(body *RigidBody, axis int) {
	var found int
	for i := 0; i < len(s.axisList[axis]); i++ {
		if s.axisList[axis][i].body == body {
			copy(s.axisList[axis][i:], s.axisList[axis][i+1:])
			s.axisList[axis] = s.axisList[axis][:len(s.axisList[axis])-1]
			found++
			i--
			if found == 2 {
				return
			}
		}
	}
}

func (s *SAP3) insertNodeAt(n sapNode, i, axis int) {
	s.axisList[axis] = append(s.axisList[axis], sapNode{})
	copy(s.axisList[axis][i+1:], s.axisList[axis][i:])
	s.axisList[axis][i] = n
}

func (s *SAP3) insertNode(n sapNode, axis int) {
	if len(s.axisList[axis]) == axis {
		s.axisList[axis] = append(s.axisList[axis], n)
		return
	}

	for i, x := range s.axisList[axis] {
		if x.value > n.value {
			s.insertNodeAt(n, i, axis)
			return
		}
	}
	s.axisList[axis] = append(s.axisList[axis], n)
}

// GeneratePotentialContacts generates all potential contacts with everybody
func (s *SAP3) GeneratePotentialContacts(narrowPhaseDetector chan<- potentialContact) {
	active := make([]*RigidBody, 0, len(s.axisList[0])/2)
	for _, n := range s.axisList[0] {

		// if its the start of an object, check if there are any active objects
		// spawn collisions for all of them and add it to the active list
		if n.start {
			for _, a := range active {
				narrowPhaseDetector <- potentialContact{
					bodies: [2]*RigidBody{a, n.body},
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
