package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// BoundingSphere represents a bounding sphere. It is used by the broadphase
// detector.
type BoundingSphere struct {
	center glm.Vec3
	radius float32
}

// NewBoundingSphere returns a bounding sphere with the given properties.
func NewBoundingSphere(center *glm.Vec3, radius float32) BoundingSphere {
	return BoundingSphere{
		center: *center,
		radius: radius,
	}
}

// NewBoundingSphereFromSpheres returns a bounding sphere that bounds the 2 given
// spheres.
func NewBoundingSphereFromSpheres(s1, s2 *BoundingSphere) BoundingSphere {
	centerOffset := s2.center
	centerOffset.SubWith(&s1.center)
	distance := centerOffset.Len2()
	radiusDiff := s2.radius - s1.radius

	if radiusDiff*radiusDiff >= distance {
		if s1.radius > s2.radius {
			return BoundingSphere{
				center: s1.center,
				radius: s1.radius,
			}
		}
		return BoundingSphere{
			center: s2.center,
			radius: s2.radius,
		}
	}
	distance = math.Sqrt(distance)
	radius := (distance + s1.radius + s2.radius) * 0.5

	// The new centre is based on s1's centre, moved towards
	// s2's centre by an amount proportional to the spheres'
	// radii.
	center := s1.center
	if distance > 0 {
		centerOffset.MulWith((radius - s1.radius) / distance)
		center.AddWith(&centerOffset)
	}
	return BoundingSphere{
		center: center,
		radius: radius,
	}
}

// Center returns the center of this bounding sphere.
func (b *BoundingSphere) Center() glm.Vec3 {
	return b.center
}

// Radius returns the radius of this sphere.
func (b *BoundingSphere) Radius() float32 {
	return b.radius
}

// Overlaps returns true if the 2 bounding spheres overlap.
func (b *BoundingSphere) Overlaps(o *BoundingSphere) bool {
	dist := b.center
	dist.SubWith(&o.center)

	r := b.radius + o.radius
	return dist.Len2() < r*r
}

// Size returns the size of this bounding volume, it is used by the bounding
// volume hierarchy. It doesn't really matter what the size if just that the
// comparison of 2 size is correct.
func (b *BoundingSphere) Size() float32 {
	return b.radius
}

// Growth returns the amount by which this bounding sphere would grow if we
// were to add this volume
func (b *BoundingSphere) Growth(o *BoundingSphere) float32 {
	new := NewBoundingSphereFromSpheres(b, o)
	return new.Size() - b.Size()
}

// MinX returns the minimum X coordinate that the sphere intersects, this is
// used by the sweep and prune broadphase.
func (b *BoundingSphere) MinX() float32 {
	return b.center.X - b.radius
}

// MaxX returns the maximum X coordinate that the sphere intersects, this is
// used by the sweep and prune broadphase.
func (b *BoundingSphere) MaxX() float32 {
	return b.center.X + b.radius
}

// MinY returns the minimum Y coordinate that the sphere intersects, this is
// used by the sweep and prune broadphase.
func (b *BoundingSphere) MinY() float32 {
	return b.center.Y - b.radius
}

// MaxY returns the maximum Y coordinate that the sphere intersects, this is
// used by the sweep and prune broadphase.
func (b *BoundingSphere) MaxY() float32 {
	return b.center.Y + b.radius
}

// MinZ returns the minimum Z coordinate that the sphere intersects, this is
// used by the sweep and prune broadphase.
func (b *BoundingSphere) MinZ() float32 {
	return b.center.Z - b.radius
}

// MaxZ returns the maximum Z coordinate that the sphere intersects, this is
// used by the sweep and prune broadphase.
func (b *BoundingSphere) MaxZ() float32 {
	return b.center.Z + b.radius
}

/*
// BoundingBox represents a bounding box. It is used by the broadphase detector.
type BoundingBox struct {
	center   glm.Vec3
	halfSize glm.Vec3
}*/
