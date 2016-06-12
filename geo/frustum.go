package geo

import (
	"fmt"
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// Frustum is a pyramid with the top cut off although with orthographic
// projections it's more of a cube then a pyramid. It is mostly used with
// rendering algorithms.
type Frustum struct {
	Planes [6]Plane
}

// FrustumFromPerspective fills the given plane by the 6 plane defining this
// projection.
func FrustumFromPerspective(fovy, aspect, near, far float32, frustum *Frustum) {
	frustum.Planes[0] = Plane{
		Offset: -near,
		Normal: glm.Vec3{0, 0, 1},
	}
	frustum.Planes[1] = Plane{
		Offset: far,
		Normal: glm.Vec3{0, 0, -1},
	}

	ay := (math.Pi - fovy) / 2
	say, cay := math.Sincos(ay)
	frustum.Planes[2] = Plane{
		Normal: glm.Vec3{0, cay, say},
	}
	frustum.Planes[3] = Plane{
		Normal: glm.Vec3{0, -cay, say},
	}

	ax := (math.Pi - fovy*aspect) / 2
	sax, cax := math.Sincos(ax)
	frustum.Planes[4] = Plane{
		Normal: glm.Vec3{cax, 0, sax},
	}
	frustum.Planes[5] = Plane{
		Normal: glm.Vec3{-cax, 0, sax},
	}
}

// FrustumFromOrthographic fills the given plane by the 6 plane defining this
// projection.
func FrustumFromOrthographic(left, right, bottom, top, near, far float32, frustum *Frustum) {
	frustum.Planes[0] = Plane{
		Offset: -near,
		Normal: glm.Vec3{0, 0, 1},
	}
	frustum.Planes[1] = Plane{
		Offset: far,
		Normal: glm.Vec3{0, 0, -1},
	}

	frustum.Planes[2] = Plane{
		Offset: math.Abs(right),
		Normal: glm.Vec3{1, 0, 0},
	}
	frustum.Planes[3] = Plane{
		Offset: math.Abs(left),
		Normal: glm.Vec3{-1, 0, 0},
	}

	frustum.Planes[4] = Plane{
		Offset: math.Abs(top),
		Normal: glm.Vec3{0, 1, 0},
	}
	frustum.Planes[5] = Plane{
		Offset: math.Abs(bottom),
		Normal: glm.Vec3{0, -1, 0},
	}
}

// TestAABBFrustum returns true if this aabb and plane intersect. It can return
// false positives but no false negatives.
func TestAABBFrustum(aabb *AABB, frustum *Frustum, view *glm.Mat4) bool {
	var taabb AABB
	var _ = fmt.Print
	UpdateAABB4(aabb, &taabb, view)
	for n := 0; n < 6; n++ {
		if !TestAABBHalfspace(aabb, &frustum.Planes[n]) {
			return false
		}
	}
	return true
}
