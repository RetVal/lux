package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// CollisionShape is just an interface that all collision shape should
// implement (implicitelly everything implements CollisionShape). It doesn't
// have any special function we just need a specific interface for them.
type CollisionShape interface {
	GetBoundingVolume() *BoundingSphere
	GetInertiaTensor(*RigidBody) glm.Mat3
	RayTest(Ray, RayResult)
}

// CollisionSphere represents a sphere.
type CollisionSphere struct {
	body   *RigidBody
	radius float32
}

// NewCollisionSphere makes a new collision sphere with the radius given.
func NewCollisionSphere(radius float32) *CollisionSphere {
	return &CollisionSphere{
		radius: radius,
	}
}

// Position returns the position of the sphere
func (s *CollisionSphere) Position() glm.Vec3 {
	return s.body.Position()
}

// Radius returns the radius of the sphere.
func (s *CollisionSphere) Radius() float32 {
	return s.radius
}

// GetBoundingVolume returns a bounding volume for this collision shape. Kinda
// redundant on a sphere...
func (s *CollisionSphere) GetBoundingVolume() *BoundingSphere {
	return &BoundingSphere{
		center: s.Position(),
		radius: s.Radius(),
	}
}

// GetInertiaTensor returns the inertia tensor for this collision shape.
func (s *CollisionSphere) GetInertiaTensor(b *RigidBody) glm.Mat3 {
	s.body = b
	return sphereInertiaTensor(b.Mass(), s.radius)
}

// RayTest tests this ray against the sphere annd the result if there is
// intersection.
func (s *CollisionSphere) RayTest(ray Ray, res RayResult) {
	o := ray.Origin()
	c := s.Position()
	l := c.Sub(&o)
	d := ray.Direction()
	sf := l.Dot(&d)
	l2 := l.Dot(&l)
	r2 := s.Radius() * s.Radius()
	// if the projection of the vector to the centre of the sphere on the
	// direction is smaller then 0 and the origin isn't inside the sphere then
	// we have a no-hit.
	if sf < 0 && l2 > r2 {
		return
	}

	m2 := l2 - sf*sf
	if m2 > r2 {
		return
	}

	q := math.Sqrt(r2 + m2)

	t := sf
	if l2 > r2 {
		t -= q
	} else {
		t += q
	}

	o.AddScaledVec(t, &d)
	res.AddResult(s.body, o)
}

// CollisionPlane represents an infinite plane.
type CollisionPlane struct {
	body   *RigidBody
	offset float32
	normal glm.Vec3
}

// Direction returns the plane normal.
func (p *CollisionPlane) Direction() glm.Vec3 {
	return p.normal
}

// Offset returns the offset of the ground.
func (p *CollisionPlane) Offset() float32 {
	return p.offset
}

// CollisionBox represents a box of any size.
type CollisionBox struct {
	body     *RigidBody
	halfSize glm.Vec3
}

// NewCollisionBox returns a collision box from the given arguments.
func NewCollisionBox(halfSize glm.Vec3) *CollisionBox {
	return &CollisionBox{
		halfSize: halfSize,
	}
}

// GetBoundingVolume returns a bounding volume for this collision shape.
func (b *CollisionBox) GetBoundingVolume() *BoundingSphere {
	r := b.halfSize.X
	if b.halfSize.Y > r {
		r = b.halfSize.Y
	}
	if b.halfSize.Z > r {
		r = b.halfSize.Z
	}

	return &BoundingSphere{
		center: b.body.Position(),
		radius: r,
	}
}

// GetInertiaTensor returns the inertia tensor for this collision shape.
func (b *CollisionBox) GetInertiaTensor(rb *RigidBody) glm.Mat3 {
	b.body = rb
	return cuboidInertiaTensor(rb.Mass(), b.halfSize.X, b.halfSize.Y, b.halfSize.Z)
}

// Position returns this collision shape position.
func (b *CollisionBox) Position() glm.Vec3 {
	return b.body.Position()
}

// RayTest tests this ray against the box annd the result if there is
// intersection.
func (b *CollisionBox) RayTest(ray Ray, res RayResult) {
	var tmin, tmax, tymin, tymax, tzmin, tzmax float32
	// We transform the ray in local space and use a AABB algorithm instead.
	ro := ray.Origin()
	ro = b.body.transformMatrix.TransformInverse(&ro)
	dir := ray.Direction()
	dir = b.body.transformMatrix.TransformInverseDirection(&dir)

	idir0 := 1.0 / dir.X
	if idir0 >= 0 {
		tmin = (-b.halfSize.X - ro.X) * idir0
		tmax = (b.halfSize.X - ro.X) * idir0
	} else {
		tmin = (b.halfSize.X - ro.X) * idir0
		tmax = (-b.halfSize.X - ro.X) * idir0
	}

	idir1 := 1.0 / dir.Y
	if idir1 >= 0 {
		tymin = (-b.halfSize.Y - ro.Y) * idir1
		tymax = (b.halfSize.Y - ro.Y) * idir1
	} else {
		tymin = (b.halfSize.Y - ro.Y) * idir1
		tymax = (-b.halfSize.Y - ro.Y) * idir1
	}

	if (tmin > tymax) || (tymin > tmax) {
		return
	}
	if tymin > tmin {
		tmin = tymin
	}
	if tymax < tmax {
		tmax = tymax
	}

	idir2 := 1.0 / dir.Z
	if idir2 >= 0 {
		tzmin = (-b.halfSize.Z - ro.Z) * idir2
		tzmax = (b.halfSize.Z - ro.Z) * idir2
	} else {
		tzmin = (b.halfSize.Z - ro.Z) * idir2
		tzmax = (-b.halfSize.Z - ro.Z) * idir2
	}

	if (tmin > tzmax) || (tzmin > tmax) {
		return
	}
	if tzmin > tmin {
		tmin = tzmin
	}
	if tzmax < tmax {
		tmax = tzmax
	}

	if tmin < ray.Len() && tmax > 0 {
		res.AddResult(b.body, ray.At(tmin))
	}
}
