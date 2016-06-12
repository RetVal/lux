package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/tensors"
	"github.com/luxengine/lux/math"
)

// OBB is a Oriented Bounding Box.
type OBB struct {
	// The center of the OBB
	Center glm.Vec3
	// The orientation of the OBB, these need to be orthonormal.
	Orientation glm.Mat3
	// The half extends of the OBB.
	HalfExtend glm.Vec3
}

// ShapeType returns the shape type for obbs.
func (*OBB) ShapeType() int {
	return obbShapeType
}

// Volume returns the volume of the obb.
func (obb *OBB) Volume() float32 {
	return 8 * obb.HalfExtend.X * obb.HalfExtend.Y * obb.HalfExtend.Z
}

// Mass returns the mass of the obb, in mass unit, given the density in
// (mass unit/distance unit^3).
func (obb *OBB) Mass(density float32) float32 {
	return density * obb.Volume()
}

// InertiaTensor returns the inertia tensor of the obb.
func (obb *OBB) InertiaTensor(density float32) glm.Mat3 {
	return tensors.Cuboid(obb.Mass(density), obb.HalfExtend.X*2, obb.HalfExtend.Y*2, obb.HalfExtend.Z*2)
}

// ClosestPointOBBPoint returns the point in or on the OBB closest to p
func ClosestPointOBBPoint(o *OBB, p *glm.Vec3) glm.Vec3 {
	closestPoint := o.Center

	d := p.Sub(&o.Center)

	// Start result at center of box; make steps from there

	// For each OBB axis...
	for i := 0; i < 3; i++ {
		// ...project d onto that axis and get the distance along the axis of d
		// from the box center
		r := o.Orientation.Row(i)
		dist := d.Dot(&r)

		// If distance farther than the box extents, clamp to the box
		if dist > *o.HalfExtend.I(i) {
			dist = *o.HalfExtend.I(i)
		} else if dist < -*o.HalfExtend.I(i) {
			dist = -*o.HalfExtend.I(i)
		}

		closestPoint.AddScaledVec(dist, &r)
	}
	return closestPoint
}

// SqDistOBBPoint returns the square distance of p to the OBB.
func SqDistOBBPoint(o *OBB, p *glm.Vec3) float32 {
	v := p.Sub(&o.Center)

	var sqDist float32

	for i := 0; i < 3; i++ {
		// Project vector from box center to 'p' on each axis, getting the
		// distance of 'p' along that axis, and count any excess distance
		// outside box extents
		var excess float32
		r := o.Orientation.Row(i)
		d := v.Dot(&r)

		if d < -*o.HalfExtend.I(i) {
			excess = d + *o.HalfExtend.I(i)
		} else if d > *o.HalfExtend.I(i) {
			excess = d - *o.HalfExtend.I(i)
		}
		sqDist += excess * excess
	}
	return sqDist
}

// TestOBBOBB returns true if these OBB overlap.
func TestOBBOBB(a, b *OBB) bool {
	// TODO(hydroflame): find a good value for that said epsilon
	const (
		epsilon = 0.0001
	)

	var R glm.Mat3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r0, r1 := a.Orientation.Row(i), b.Orientation.Row(i)
			R[j*3+i] = r0.Dot(&r1)
		}
	}

	// Compute translation vector t
	t := b.Center.Sub(&a.Center)

	r0 := a.Orientation.Row(0)
	r1 := a.Orientation.Row(1)
	r2 := a.Orientation.Row(2)
	// Bring translation into a's coordinate frame
	t = glm.Vec3{
		X: t.Dot(&r0),
		Y: t.Dot(&r1),
		Z: t.Dot(&r2),
	}

	var AbsR glm.Mat3
	// Compute common subexpressions. Add in an epsilon term to counteract
	// arithmetic errors when two edges are parallel and their cross product is
	// (near) zero.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			AbsR[j*3+i] = math.Abs(R[j*3+i]) + epsilon
		}
	}

	var ra, rb float32
	// Test axes L = A0, L = A1, L = A2
	for i := 0; i < 3; i++ {
		ra = *a.HalfExtend.I(i)
		rb = b.HalfExtend.X*AbsR[0*3+i] + b.HalfExtend.Y*AbsR[1*3+i] + b.HalfExtend.Z*AbsR[2*3+i]
		if math.Abs(*t.I(i)) > ra+rb {
			return false
		}
	}

	// Test axes L = B0, L = B1, L = B2
	for i := 0; i < 3; i++ {
		ra = a.HalfExtend.X*AbsR[i*3+0] + a.HalfExtend.Y*AbsR[i*3+1] + a.HalfExtend.Z*AbsR[i*3+2]
		rb = *b.HalfExtend.I(i)
		if math.Abs(t.X*R[i*3+0]+t.Y*R[i*3+1]+t.Z*R[i*3+2]) > ra+rb {
			return false
		}
	}

	// Test axis L = A0 x B0
	ra = a.HalfExtend.Y*AbsR[2] + a.HalfExtend.Z*AbsR[1]
	rb = b.HalfExtend.Y*AbsR[6] + b.HalfExtend.Z*AbsR[3]
	if math.Abs(t.Z*R[1]-t.Y*R[2]) > ra+rb {
		return false
	}

	// Test axis L = A0 x B1
	ra = a.HalfExtend.Y*AbsR[5] + a.HalfExtend.Z*AbsR[4]
	rb = b.HalfExtend.X*AbsR[6] + b.HalfExtend.Z*AbsR[0]
	if math.Abs(t.Z*R[4]-t.Y*R[5]) > ra+rb {
		return false
	}

	// Test axis L = A0 x B2
	ra = a.HalfExtend.Y*AbsR[8] + a.HalfExtend.Z*AbsR[7]
	rb = b.HalfExtend.X*AbsR[3] + b.HalfExtend.Y*AbsR[0]
	if math.Abs(t.Z*R[7]-t.Y*R[8]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B0
	ra = a.HalfExtend.X*AbsR[2] + a.HalfExtend.Z*AbsR[0]
	rb = b.HalfExtend.Y*AbsR[7] + b.HalfExtend.Z*AbsR[4]
	if math.Abs(t.X*R[2]-t.Z*R[0]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B1
	ra = a.HalfExtend.X*AbsR[5] + a.HalfExtend.Z*AbsR[3]
	rb = b.HalfExtend.X*AbsR[7] + b.HalfExtend.Z*AbsR[1]
	if math.Abs(t.X*R[5]-t.Z*R[3]) > ra+rb {
		return false
	}

	// Test axis L = A1 x B2
	ra = a.HalfExtend.X*AbsR[8] + a.HalfExtend.Z*AbsR[6]
	rb = b.HalfExtend.X*AbsR[4] + b.HalfExtend.Y*AbsR[1]
	if math.Abs(t.X*R[8]-t.Z*R[6]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B0
	ra = a.HalfExtend.X*AbsR[1] + a.HalfExtend.Y*AbsR[0]
	rb = b.HalfExtend.Y*AbsR[8] + b.HalfExtend.Z*AbsR[5]
	if math.Abs(t.Y*R[0]-t.X*R[1]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B1
	ra = a.HalfExtend.X*AbsR[4] + a.HalfExtend.Y*AbsR[3]
	rb = b.HalfExtend.X*AbsR[8] + b.HalfExtend.Z*AbsR[2]
	if math.Abs(t.Y*R[3]-t.X*R[4]) > ra+rb {
		return false
	}

	// Test axis L = A2 x B2
	ra = a.HalfExtend.X*AbsR[7] + a.HalfExtend.Y*AbsR[6]
	rb = b.HalfExtend.X*AbsR[5] + b.HalfExtend.Y*AbsR[2]
	if math.Abs(t.Y*R[6]-t.X*R[7]) > ra+rb {
		return false
	}

	return true
}
