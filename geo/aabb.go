package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// AABB is an axis-aligned bounding box
type AABB struct {
	// Center represents the center of the bounding box.
	Center glm.Vec3

	// HalfExtend represents the 3 half extends of the bounding box.
	HalfExtend glm.Vec3
}

// ShapeType returns the shape type for aabbs.
func (*AABB) ShapeType() int {
	return aabbShapeType
}

// Volume returns the volume of this aabb.
func (aabb *AABB) Volume() float32 {
	return 8 * aabb.HalfExtend.X * aabb.HalfExtend.Y * aabb.HalfExtend.Z
}

// Mass returns the mass of the aabb, in mass unit, given the density in
// (mass unit/distance unit^3).
func (aabb *AABB) Mass(density float32) float32 {
	return density * aabb.Volume()
}

// InertiaTensor returns the inertia tensor of the aabb. Note that all AABB have
// a zero inertia tensor as they cannot rotate.
func (aabb *AABB) InertiaTensor() glm.Mat3 {
	return glm.Mat3{}
}

// TestAABBAABB returns true if these AABB overlap.
func TestAABBAABB(a, b *AABB) bool {
	if math.Abs(a.Center.X-b.Center.X) > a.HalfExtend.X+b.HalfExtend.X ||
		math.Abs(a.Center.Y-b.Center.Y) > a.HalfExtend.Y+b.HalfExtend.Y ||
		math.Abs(a.Center.Z-b.Center.Z) > a.HalfExtend.Z+b.HalfExtend.Z {
		return false
	}
	return true
}

// UpdateAABB3x4 computes an enclosing AABB base transformed by t and puts the
// result in fill. base and fill must not be the same.
func UpdateAABB3x4(base, fill *AABB, t *glm.Mat3x4) {
	for i := 0; i < 3; i++ {
		*fill.Center.I(i) = t[i+9]
		*fill.HalfExtend.I(i) = 0
		for j := 0; j < 3; j++ {
			*fill.Center.I(i) += t[j*3+i] * *base.Center.I(j)
			*fill.HalfExtend.I(i) += math.Abs(t[j*3+i]) * *base.HalfExtend.I(j)
		}
	}
}

// UpdateAABB4 computes an enclosing AABB base transformed by t and puts the
// result in fill. base and fill must not be the same.
func UpdateAABB4(base, fill *AABB, t *glm.Mat4) {
	for i := 0; i < 3; i++ {
		*fill.Center.I(i) = t[i+12]
		*fill.HalfExtend.I(i) = 0
		for j := 0; j < 3; j++ {
			*fill.Center.I(i) += t[j*4+i] * *base.Center.I(j)
			*fill.HalfExtend.I(i) += math.Abs(t[j*4+i]) * *base.HalfExtend.I(j)
		}
	}
}

// ClosestPointPointAABB returns the point in or on the AABB closest to p.
func ClosestPointPointAABB(p *glm.Vec3, a *AABB) glm.Vec3 {
	return glm.Vec3{
		X: math.Clamp(p.X, a.Center.X-a.HalfExtend.X, a.Center.X+a.HalfExtend.X),
		Y: math.Clamp(p.Y, a.Center.Y-a.HalfExtend.Y, a.Center.Y+a.HalfExtend.Y),
		Z: math.Clamp(p.Z, a.Center.Z-a.HalfExtend.Z, a.Center.Z+a.HalfExtend.Z),
	}
}

// SqDistPointAABB returns the square distance of p to the AABB.
func SqDistPointAABB(p *glm.Vec3, a *AABB) float32 {
	var sqDist float32

	// For each axis count any excess distance outside box extents
	v := p.X
	min := a.Center.X - a.HalfExtend.X
	max := a.Center.X + a.HalfExtend.X
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	v = p.Y
	min = a.Center.Y - a.HalfExtend.Y
	max = a.Center.Y + a.HalfExtend.Y
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	v = p.Z
	min = a.Center.Z - a.HalfExtend.Z
	max = a.Center.Z + a.HalfExtend.Z
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	return sqDist
}
