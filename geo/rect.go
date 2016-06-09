package geo

import (
	"github.com/luxengine/lux/glm"
)

// Rect is a rectangle in 3D space, they are a simpler version of OBBs.
type Rect struct {
	// the center of the rectangle
	Center glm.Vec3
	// the orientation of the rectangle in space.
	Orientation [2]glm.Vec3
	// the half extends of the rectangle.
	HalfExtend glm.Vec2
}

// Area returns the area of the rectangle
func (r *Rect) Area() float32 {
	return 4 * r.HalfExtend.X * r.HalfExtend.Y
}

// ClosestPointPointRect returns the point on the rectangle closest to p
func ClosestPointPointRect(p *glm.Vec3, r *Rect) glm.Vec3 {
	d := p.Sub(&r.Center)
	closestPoint := r.Center

	// Start result at center of box; make steps from there

	dist := d.Dot(&r.Orientation[0])

	// If distance farther than the box extents, clamp to the box
	if dist > r.HalfExtend.X {
		dist = r.HalfExtend.X
	} else if dist < -r.HalfExtend.X {
		dist = -r.HalfExtend.X
	}

	closestPoint.AddScaledVec(dist, &r.Orientation[0])

	dist = d.Dot(&r.Orientation[1])

	// If distance farther than the box extents, clamp to the box
	if dist > r.HalfExtend.Y {
		dist = r.HalfExtend.Y
	} else if dist < -r.HalfExtend.Y {
		dist = -r.HalfExtend.Y
	}

	closestPoint.AddScaledVec(dist, &r.Orientation[1])

	return closestPoint
}

// SqDistPointRect returns the square distance of p to the rectangle.
func SqDistPointRect(p *glm.Vec3, r *Rect) float32 {
	v := p.Sub(&r.Center)

	var sqDist float32

	for i := 0; i < len(r.Orientation); i++ {
		// Project vector from box center to p on each axis, getting the
		// distance of p along that axis, and count any excess distance
		// outside box extents.
		var excess float32
		d := v.Dot(&r.Orientation[i])

		if d < -*r.HalfExtend.I(i) {
			excess = d + *r.HalfExtend.I(i)
		} else if d > *r.HalfExtend.I(i) {
			excess = d - *r.HalfExtend.I(i)
		}
		sqDist += excess * excess
	}
	return sqDist
}
