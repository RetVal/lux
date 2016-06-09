package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// sphereAndSphere generates contacts between 2 collision sphere.
func sphereAndSphere(s1, s2 *CollisionSphere, contacts []Contact) int {
	var midline glm.Vec3
	p1, p2 := s1.Position(), s2.Position()
	midline.SubOf(&p1, &p2)
	size := midline.Len()

	//if they're not touching we can leave.
	if size > s1.Radius()+s2.Radius() {
		return 0
	}

	var normal glm.Vec3
	normal.MulOf(1.0/size, &midline)

	point := p2
	point.AddScaledVec(0.5, &midline)
	contacts[0] = Contact{
		bodies:      [2]*RigidBody{s1.body, s2.body},
		point:       point,
		normal:      normal,
		penetration: s1.Radius() + s2.Radius() - size,
		friction:    (s1.body.friction + s2.body.friction) / 2,
		restitution: (s1.body.restitution + s2.body.restitution) / 2,
	}
	return 1
}

/*
// sphereAndHalfSpace collides a sphere and a half space.
func sphereAndHalfSpace(s *CollisionSphere, p *CollisionPlane, contacts []Contact) int {
	spos := s.Position()

	dir := p.Direction()
	ballDistance := dir.Dot(&spos) - s.Radius() - p.Offset()

	if ballDistance >= 0 {
		return 0
	}

	point := dir
	point.MulWith(ballDistance + s.Radius())
	point.SubOf(&spos, &point)

	contacts[0] = Contact{
		bodies:      [2]*RigidBody{s.body, nil},
		point:       point,
		normal:      dir,
		penetration: -ballDistance,
		friction:    0,
	}
	return 1
}

// boxAndHalfSpace makes the collision between a box and a half space. We can't
// produce a face-face collision with boxes and half spaces (or plane for that
// matter), neither can they produce a edge-face collision. So really we just
// need to check every vertex-face collisions.
func boxAndHalfSpace(b *CollisionBox, p *CollisionPlane, contacts []Contact) int {
	var numcontact int
	// make all vertices
	vertices := [...]glm.Vec3{
		{-b.halfSize[0], -b.halfSize[1], -b.halfSize[2]},
		{-b.halfSize[0], -b.halfSize[1], b.halfSize[2]},
		{b.halfSize[0], -b.halfSize[1], -b.halfSize[2]},
		{b.halfSize[0], -b.halfSize[1], b.halfSize[2]},
		{-b.halfSize[0], b.halfSize[1], -b.halfSize[2]},
		{-b.halfSize[0], b.halfSize[1], b.halfSize[2]},
		{b.halfSize[0], b.halfSize[1], -b.halfSize[2]},
		{b.halfSize[0], b.halfSize[1], b.halfSize[2]},
	}

	dir := p.Direction()

	for x := 0; x < len(vertices); x++ {
		// transform it.
		vertexPos := b.body.transformMatrix.Mul3x1(&vertices[x])

		// basically we transform the vertex to the same axis as the offset as
		// the plane offset represent (Y axis?) and then just check if we're
		// lower
		vertexDistance := vertexPos.Dot(&dir)

		// Is it lower ?
		if vertexDistance <= p.Offset() {
			//we have a contact.
			var point glm.Vec3
			point.MulOf(vertexDistance-p.Offset(), &dir)
			contacts[numcontact] = Contact{
				bodies:      [2]*RigidBody{b.body, nil},
				point:       point,
				normal:      dir,
				penetration: p.Offset() - vertexDistance,
				friction:    0,
			}
			numcontact++
		}
		if len(contacts) <= numcontact {
			return numcontact
		}
	}
	return numcontact
}
*/

// sphereAndBox check for collision between a sphere and a box.
func sphereAndBox(s *CollisionSphere, b *CollisionBox, contacts []Contact) int {
	scenter := s.Position()
	transform := b.body.transformMatrix
	relsCenter := transform.TransformInverse(&scenter)

	closestPoint := glm.Vec3{
		X: math.Clamp(relsCenter.X, -b.halfSize.X, b.halfSize.X),
		Y: math.Clamp(relsCenter.Y, -b.halfSize.Y, b.halfSize.Y),
		Z: math.Clamp(relsCenter.Z, -b.halfSize.Z, b.halfSize.Z),
	}

	// At this point we have the closest point from the box to the sphere. If
	// that one is not colliding then clearly there's no collision.

	var diff glm.Vec3
	diff.SubOf(&closestPoint, &relsCenter)
	dist := diff.Len2()

	// no collision
	if dist > s.Radius()*s.Radius() {
		return 0
	}

	closestPointWorld := transform.Transform(&closestPoint)

	var normal glm.Vec3
	normal.SubOf(&scenter, &closestPointWorld)

	pen := s.Radius()

	// If the sphere is inside the box its normal will be zero.
	var zero glm.Vec3
	if normal.EqualThreshold(&zero, 1e-2) {
		index := 1
		d := b.halfSize.X - closestPoint.X
		if f := math.Abs(-b.halfSize.X - closestPoint.X); f < d {
			d = f
			index = -1
		}
		if f := b.halfSize.Y - closestPoint.Y; f < d {
			d = f
			index = 2
		}
		if f := math.Abs(-b.halfSize.Y - closestPoint.Y); f < d {
			d = f
			index = -2
		}

		if f := b.halfSize.Z - closestPoint.Z; f < d {
			d = f
			index = 3
		}
		if f := math.Abs(-b.halfSize.Z - closestPoint.Z); f < d {
			d = f
			index = -3
		}
		if index > 0 {
			*normal.I(index - 1) = 1
		} else {
			*normal.I((-index) - 1) = -1
		}

		pen += d

	} else { //the centre of the sphere is not inside the box.
		normal.Normalize()
		pen -= math.Sqrt(dist)
	}

	contacts[0] = Contact{
		bodies:      [2]*RigidBody{s.body, b.body},
		point:       closestPointWorld,
		normal:      normal,
		penetration: pen,
		friction:    (s.body.friction + b.body.friction) / 2,
		restitution: (s.body.restitution + b.body.restitution) / 2,
	}
	return 1
}
