package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

func boxAndBoxEarly(b1, b2 *CollisionBox) bool {
	// % = cross product
	// * = dot product
	p1 := glm.Vec3{X: b1.body.transformMatrix[9], Y: b1.body.transformMatrix[10], Z: b1.body.transformMatrix[11]}
	p2 := glm.Vec3{X: b2.body.transformMatrix[9], Y: b2.body.transformMatrix[10], Z: b2.body.transformMatrix[11]}
	toCenter := p2.Sub(&p1)

	// just store a copy of all these vectors
	b10 := glm.Vec3{X: b1.body.transformMatrix[0], Y: b1.body.transformMatrix[1], Z: b1.body.transformMatrix[2]}
	b11 := glm.Vec3{X: b1.body.transformMatrix[3], Y: b1.body.transformMatrix[4], Z: b1.body.transformMatrix[5]}
	b12 := glm.Vec3{X: b1.body.transformMatrix[6], Y: b1.body.transformMatrix[7], Z: b1.body.transformMatrix[8]}

	b20 := glm.Vec3{X: b2.body.transformMatrix[0], Y: b2.body.transformMatrix[1], Z: b2.body.transformMatrix[2]}
	b21 := glm.Vec3{X: b2.body.transformMatrix[3], Y: b2.body.transformMatrix[4], Z: b2.body.transformMatrix[5]}
	b22 := glm.Vec3{X: b2.body.transformMatrix[6], Y: b2.body.transformMatrix[7], Z: b2.body.transformMatrix[8]}

	b10b20 := b10.Cross(&b20)
	b10b21 := b10.Cross(&b21)
	b10b22 := b10.Cross(&b22)
	b11b20 := b11.Cross(&b20)
	b11b21 := b11.Cross(&b21)
	b11b22 := b11.Cross(&b22)
	b12b20 := b12.Cross(&b20)
	b12b21 := b12.Cross(&b21)
	b12b22 := b12.Cross(&b22)

	return penetrationOnAxis(b1, b2, &b10, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b11, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b12, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b20, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b21, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b22, &toCenter) >= 0 &&

		penetrationOnAxis(b1, b2, &b10b20, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b10b21, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b10b22, &toCenter) >= 0 &&

		penetrationOnAxis(b1, b2, &b11b20, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b11b21, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b11b22, &toCenter) >= 0 &&

		penetrationOnAxis(b1, b2, &b12b20, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b12b21, &toCenter) >= 0 &&
		penetrationOnAxis(b1, b2, &b12b22, &toCenter) >= 0
}

func penetrationOnAxis(b1, b2 *CollisionBox, axis, toCenter *glm.Vec3) float32 {
	b1Project, b2Project := transformToAxis(b1, axis), transformToAxis(b2, axis)
	distance := math.Abs(toCenter.Dot(axis))
	return b1Project + b2Project - distance
}

func transformToAxis(b *CollisionBox, axis *glm.Vec3) float32 {
	b1 := glm.Vec3{X: b.body.transformMatrix[0], Y: b.body.transformMatrix[1], Z: b.body.transformMatrix[2]}
	b2 := glm.Vec3{X: b.body.transformMatrix[3], Y: b.body.transformMatrix[4], Z: b.body.transformMatrix[5]}
	b3 := glm.Vec3{X: b.body.transformMatrix[6], Y: b.body.transformMatrix[7], Z: b.body.transformMatrix[8]}

	return b.halfSize.X*math.Abs(axis.Dot(&b1)) +
		b.halfSize.Y*math.Abs(axis.Dot(&b2)) +
		b.halfSize.Z*math.Abs(axis.Dot(&b3))
}

func tryAxis(one, two *CollisionBox, axis glm.Vec3, toCentre *glm.Vec3, index uint32, smallestPenetration *float32, smallestCase *uint32) bool {
	// Make sure we have a normalized axis, and don't check almost parallel axes
	if axis.Len2() < 1e-7 {
		return true
	}
	axis.Normalize()

	penetration := penetrationOnAxis(one, two, &axis, toCentre)

	if penetration < 0 {
		return false
	}

	if penetration < *smallestPenetration {
		*smallestPenetration = penetration
		*smallestCase = index
	}
	return true
}

// This method is called when we know that a vertex from
// box two is in contact with box one.
func fillPointFaceBoxBox(one, two *CollisionBox, toCentre *glm.Vec3, contacts []Contact, best uint32, pen float32) {
	// We know which axis the collision is on (i.e. best),
	// but we need to work out which of the two faces on
	// this axis.
	normal := glm.Vec3{X: one.body.transformMatrix[best*3], Y: one.body.transformMatrix[(best*3)+1], Z: one.body.transformMatrix[(best*3)+2]}
	if normal.Dot(toCentre) > 0 {
		normal.Invert()
	}

	// Work out which vertex of box two we're colliding with.
	vertex := two.halfSize

	t0 := glm.Vec3{X: two.body.transformMatrix[0], Y: two.body.transformMatrix[1], Z: two.body.transformMatrix[2]}
	if t0.Dot(&normal) < 0 {
		vertex.X = -vertex.X
	}

	t1 := glm.Vec3{X: two.body.transformMatrix[3], Y: two.body.transformMatrix[4], Z: two.body.transformMatrix[5]}
	if t1.Dot(&normal) < 0 {
		vertex.Y = -vertex.Y
	}

	t2 := glm.Vec3{X: two.body.transformMatrix[6], Y: two.body.transformMatrix[7], Z: two.body.transformMatrix[8]}
	if t2.Dot(&normal) < 0 {
		vertex.Z = -vertex.Z
	}

	contacts[0] = Contact{
		bodies:      [2]*RigidBody{one.body, two.body},
		point:       two.body.transformMatrix.Mul3x1(&vertex),
		normal:      normal,
		penetration: pen,
		friction:    (one.body.friction + two.body.friction) / 2,
		restitution: (one.body.restitution + two.body.restitution) / 2,
	}
}

// boxAndBox generates collision between 2 boxes.
func boxAndBox(b1, b2 *CollisionBox, contacts []Contact) int {
	if !boxAndBoxEarly(b1, b2) {
		return 0
	}

	p1, p2 := b1.body.Position(), b2.body.Position()
	toCentre := p2.Sub(&p1)

	// We start by assuming theres is no contact.
	pen := float32(math.MaxFloat32)
	best := uint32(0xFFFFFF)

	// Now we check each axes, returning if it gives us
	// a separating axis, and keeping track of the axis with
	// the smallest penetration otherwise.

	b10 := glm.Vec3{X: b1.body.transformMatrix[0], Y: b1.body.transformMatrix[1], Z: b1.body.transformMatrix[2]}
	if !tryAxis(b1, b2, b10, &toCentre, 0, &pen, &best) {
		return 0
	}

	b11 := glm.Vec3{X: b1.body.transformMatrix[3], Y: b1.body.transformMatrix[4], Z: b1.body.transformMatrix[5]}
	if !tryAxis(b1, b2, b11, &toCentre, 1, &pen, &best) {
		return 0
	}

	b12 := glm.Vec3{X: b1.body.transformMatrix[6], Y: b1.body.transformMatrix[7], Z: b1.body.transformMatrix[8]}
	if !tryAxis(b1, b2, b12, &toCentre, 2, &pen, &best) {
		return 0
	}

	b20 := glm.Vec3{X: b2.body.transformMatrix[0], Y: b2.body.transformMatrix[1], Z: b2.body.transformMatrix[2]}
	if !tryAxis(b1, b2, b20, &toCentre, 3, &pen, &best) {
		return 0
	}

	b21 := glm.Vec3{X: b2.body.transformMatrix[3], Y: b2.body.transformMatrix[4], Z: b2.body.transformMatrix[5]}
	if !tryAxis(b1, b2, b21, &toCentre, 4, &pen, &best) {
		return 0
	}

	b22 := glm.Vec3{X: b2.body.transformMatrix[6], Y: b2.body.transformMatrix[7], Z: b2.body.transformMatrix[8]}
	if !tryAxis(b1, b2, b22, &toCentre, 5, &pen, &best) {
		return 0
	}

	// Store the best axis-major, in case we run into almost
	// parallel edge collisions later
	bestSingleAxis := best

	if !tryAxis(b1, b2, b10.Cross(&b20), &toCentre, 6, &pen, &best) {
		return 0
	}

	if !tryAxis(b1, b2, b10.Cross(&b21), &toCentre, 7, &pen, &best) {
		return 0
	}

	if !tryAxis(b1, b2, b10.Cross(&b22), &toCentre, 8, &pen, &best) {
		return 0
	}

	if !tryAxis(b1, b2, b11.Cross(&b20), &toCentre, 9, &pen, &best) {
		return 0
	}

	if !tryAxis(b1, b2, b11.Cross(&b21), &toCentre, 10, &pen, &best) {
		return 0
	}

	if !tryAxis(b1, b2, b11.Cross(&b22), &toCentre, 11, &pen, &best) {
		return 0
	}

	if !tryAxis(b1, b2, b12.Cross(&b20), &toCentre, 12, &pen, &best) {
		return 0
	}

	if !tryAxis(b1, b2, b12.Cross(&b21), &toCentre, 13, &pen, &best) {
		return 0
	}

	if !tryAxis(b1, b2, b12.Cross(&b22), &toCentre, 14, &pen, &best) {
		return 0
	}

	// We now know there's a collision, and we know which
	// of the axes gave the smallest penetration. We now
	// can deal with it in different ways depending on
	// the case.
	if best < 3 {

		// We've got a vertex of box two on a face of box one.
		fillPointFaceBoxBox(b1, b2, &toCentre, contacts, best, pen)
		return 1
	} else if best < 6 {
		// We've got a vertex of box one on a face of box two.
		// We use the same algorithm as above, but swap around
		// one and two (and therefore also the vector between their
		// centres).
		toCentre.Invert()
		fillPointFaceBoxBox(b2, b1, &toCentre, contacts, best-3, pen)
		return 1
	} else {

		// We've got an edge-edge contact. Find out which axes
		best -= 6
		oneAxisIndex := best / 3
		twoAxisIndex := best % 3

		oneAxis := glm.Vec3{X: b1.body.transformMatrix[(oneAxisIndex * 3)], Y: b1.body.transformMatrix[(oneAxisIndex*3)+1], Z: b1.body.transformMatrix[(oneAxisIndex*3)+2]}
		twoAxis := glm.Vec3{X: b2.body.transformMatrix[(twoAxisIndex * 3)], Y: b2.body.transformMatrix[(twoAxisIndex*3)+1], Z: b2.body.transformMatrix[(twoAxisIndex*3)+2]}
		axis := oneAxis.Cross(&twoAxis)
		axis.Normalize()

		// The axis should point from box one to box two.
		if axis.Dot(&toCentre) > 0 {
			axis.Invert()
		}

		// We have the axes, but not the edges: each axis has 4 edges parallel
		// to it, we need to find which of the 4 for each object. We do
		// that by finding the point in the centre of the edge. We know
		// its component in the direction of the box's collision axis is zero
		// (its a mid-point) and we determine which of the extremes in each
		// of the other axes is closest.
		ptOnOneEdge := b1.halfSize
		ptOnTwoEdge := b2.halfSize
		for i := uint32(0); i < 3; i++ {
			if i == oneAxisIndex {
				*ptOnOneEdge.I(int(i)) = 0
			} else if a := (glm.Vec3{X: b1.body.transformMatrix[i*3], Y: b1.body.transformMatrix[i*3+1], Z: b1.body.transformMatrix[i*3+2]}); a.Dot(&axis) > 0 {
				*ptOnOneEdge.I(int(i)) = -*ptOnOneEdge.I(int(i))
			}

			if i == twoAxisIndex {
				*ptOnTwoEdge.I(int(i)) = 0
			} else if a := (glm.Vec3{X: b2.body.transformMatrix[i*3], Y: b2.body.transformMatrix[i*3+1], Z: b2.body.transformMatrix[i*3+2]}); a.Dot(&axis) < 0 {
				*ptOnTwoEdge.I(int(i)) = -*ptOnTwoEdge.I(int(i))
			}
		}

		// Move them into world coordinates (they are already oriented
		// correctly, since they have been derived from the axes).
		ptOnOneEdge = b1.body.transformMatrix.Mul3x1(&ptOnOneEdge)
		ptOnTwoEdge = b2.body.transformMatrix.Mul3x1(&ptOnTwoEdge)

		// So we have a point and a direction for the colliding edges.
		// We need to find out point of closest approach of the two
		// line-segments.
		vertex := contactPoint(
			&ptOnOneEdge, &oneAxis, *b1.halfSize.I(int(oneAxisIndex)),
			&ptOnTwoEdge, &twoAxis, *b2.halfSize.I(int(twoAxisIndex)),
			bestSingleAxis > 2)

		contacts[0] = Contact{
			bodies:      [2]*RigidBody{b1.body, b2.body},
			point:       vertex,
			normal:      axis,
			penetration: pen,
			friction:    (b1.body.friction + b2.body.friction) / 2,
			restitution: (b1.body.restitution + b2.body.restitution) / 2,
		}
		return 1
	}
}

func contactPoint(pOne, dOne *glm.Vec3, oneSize float32, pTwo, dTwo *glm.Vec3, twoSize float32, useOne bool) glm.Vec3 {
	//Vector3 toSt, cOne, cTwo;
	//real dpStaOne, dpStaTwo, dpOneTwo, smOne, smTwo;
	//real denom, mua, mub;

	smOne := dOne.Len2()       // float32
	smTwo := dTwo.Len2()       // float32
	dpOneTwo := dTwo.Dot(dOne) // float32

	toSt := pOne.Sub(pTwo)      // vec3
	dpStaOne := dOne.Dot(&toSt) // float32
	dpStaTwo := dTwo.Dot(&toSt) // float32

	denom := smOne*smTwo - dpOneTwo*dpOneTwo //float32

	// Zero denominator indicates parallel lines
	if math.Abs(denom) < 0.0001 {
		if useOne {
			return *pOne
		}
		return *pTwo
	}

	mua := (dpOneTwo*dpStaTwo - smTwo*dpStaOne) / denom // float32
	mub := (smOne*dpStaTwo - dpOneTwo*dpStaOne) / denom // float32

	// If either of the edges has the nearest point out
	// of bounds, then the edges aren't crossed, we have
	// an edge-face contact. Our point is on the edge, which
	// we know from the useOne parameter.
	if mua > oneSize || mua < -oneSize || mub > twoSize || mub < -twoSize {
		if useOne {
			return *pOne
		}
		return *pTwo
	}

	tmp := dOne.Mul(mua)
	cOne := pOne.Add(&tmp)

	tmp = dTwo.Mul(mub)
	cTwo := pTwo.Add(&tmp)

	t1 := cOne.Mul(0.5)
	t1.AddScaledVec(0.5, &cTwo)
	return t1
}
