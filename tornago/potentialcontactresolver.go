package tornago

const (
	unsupportedcollisionshape = "Unsupported collision shape. Please notify the lux team."
)

// potentialContact holds 2 rigid bodies that might be in contact.
type potentialContact struct {
	bodies [2]*RigidBody
}

// resolvePotentialContacts checks every potential contact and calls the
// appropriate function to verify and generate it.
func resolvePotentialContacts(pcontacts []potentialContact, contacts []Contact) int {
	var size int
	for _, pc := range pcontacts {
		if size == len(contacts) {
			// uh oh, we ran out of space.
			return size
		}

		// check collision filtering
		if pc.bodies[0].Group()&pc.bodies[1].Mask() == 0 || pc.bodies[1].Group()&pc.bodies[0].Mask() == 0 {
			continue
		}

		switch shape1 := pc.bodies[0].shape.(type) {
		case *CollisionSphere:
			switch shape2 := pc.bodies[1].shape.(type) {
			case *CollisionSphere:
				c := sphereAndSphere(shape1, shape2, contacts[size:])
				size += c
			case *CollisionBox:
				c := sphereAndBox(shape1, shape2, contacts[size:])
				size += c
			default:
				panic(unsupportedcollisionshape)
			}
		case *CollisionBox:
			switch shape2 := pc.bodies[1].shape.(type) {
			case *CollisionSphere:
				c := sphereAndBox(shape2, shape1, contacts[size:])
				size += c
			case *CollisionBox:
				c := boxAndBox(shape1, shape2, contacts[size:])
				size += c
			default:
				panic(unsupportedcollisionshape)
			}
		default:
			panic(unsupportedcollisionshape)
		}
	}
	return size
}
