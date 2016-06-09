package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// Ray represents a ray to be cast in the world.
type Ray struct {
	// origin is where the ray starts in world coordinates.
	origin glm.Vec3

	// direction is the direction of the ray starting at origin.
	direction glm.Vec3

	// len is how far the ray extends starting from origin in the rays
	// direction.
	len float32
}

// NewRay returns a ray from the given origin, direction and length.
func NewRay(origin, direction glm.Vec3, len float32) Ray {
	return Ray{
		origin:    origin,
		direction: direction.Normalized(),
		len:       len,
	}
}

// NewRayFromTo takes the start and end in world coordinates and makes a ray
// with them. if from is equal to to  we return a Ray that points up and a Len
// of 0.
func NewRayFromTo(from, to glm.Vec3) Ray {
	dir := to.Sub(&from)
	normal := dir.Normalized()
	if math.IsNaN(normal.X) || math.IsNaN(normal.Y) || math.IsNaN(normal.Z) {
		normal = glm.Vec3{X: 0, Y: 1, Z: 0}
	}
	return Ray{
		origin:    from,
		direction: normal,
		len:       dir.Len(),
	}
}

// Direction returns the direction of the ray in world coordinates. This Vec3 is
// guaranteed to be normal.
func (r Ray) Direction() glm.Vec3 {
	return r.direction
}

// Origin returns the origin of this ray in world coordinates.
func (r Ray) Origin() glm.Vec3 {
	return r.origin
}

// Len returns the length of this ray.
func (r Ray) Len() float32 {
	return r.len
}

// Destination returns the other end of the ray.
func (r Ray) Destination() glm.Vec3 {
	tmp := r.direction.Mul(r.len)
	return tmp.Add(&r.origin)
}

// At takes a float and returns a point f unit away from the origin
// in the ray direction.
func (r Ray) At(f float32) glm.Vec3 {
	tmp := r.origin
	tmp.AddScaledVec(f, &r.direction)
	return tmp
}
