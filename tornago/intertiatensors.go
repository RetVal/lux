package tornago

import (
	"github.com/luxengine/lux/glm"
)

// sphereInertiaTensor returns the inertia tensor of a sphere.
func sphereInertiaTensor(mass, radius float32) glm.Mat3 {
	v := 0.4 * mass * radius * radius
	return glm.Mat3{v, 0, 0, 0, v, 0, 0, 0, v}
}

// cuboidInertiaTensor returns the inertiatensor of a cuboid shape.
func cuboidInertiaTensor(mass, dx, dy, dz float32) glm.Mat3 {
	const (
		f = 0.3
	)
	dx2, dy2, dz2 := dx*dx, dy*dy, dz*dz
	return glm.Mat3{
		f * mass * (dy2 + dz2), 0, 0,
		0, f * mass * (dx2 + dz2), 0,
		0, 0, f * mass * (dx2 + dy2),
	}
}
