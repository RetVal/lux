package geo

import (
	"github.com/luxengine/lux/glm"
)

// Plane is a hyperplane in 3D.
type Plane struct {
	// The normal to the plane.
	Normal glm.Vec3

	// Offset from origin in the direction of the normal. NOT the up direction.
	Offset float32
}

// PlaneFromPoints computes the plane given by [ab,ac].
func PlaneFromPoints(a, b, c *glm.Vec3) Plane {
	ab, ac := b.Sub(a), c.Sub(a)
	abac := ab.Cross(&ac)
	abac.Normalize()

	return Plane{
		Normal: abac,
		Offset: abac.Dot(a),
	}
}

// DistanceToPlane returns the distance of v to plane p.
func DistanceToPlane(plane *Plane, point *glm.Vec3) float32 {
	return plane.Normal.Dot(point) - plane.Offset
}
