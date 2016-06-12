package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/tensors"
	"github.com/luxengine/lux/math"
)

// Sphere is a bounding volume for spheres.
type Sphere struct {
	Center glm.Vec3
	Radius float32
}

// ShapeType returns the shape type for spheres.
func (*Sphere) ShapeType() int {
	return sphereShapeType
}

// Volume returns the volume of this sphere
func (sphere *Sphere) Volume() float32 {
	// 4/3 pi r^3
	return 4.0 / 3.0 * math.Pi * sphere.Radius * sphere.Radius * sphere.Radius
}

// Mass returns the mass of the sphere, in mass unit, given the density in
// (mass unit/distance unit^3).
func (sphere *Sphere) Mass(density float32) float32 {
	return density * sphere.Volume()
}

// InertiaTensor returns the inertia tensor of the sphere.
func (sphere *Sphere) InertiaTensor(density float32) glm.Mat3 {
	return tensors.Sphere(sphere.Mass(density), sphere.Radius)
}

// TestSphereSphere return true if the spheres overlap.
func TestSphereSphere(a, b *Sphere) bool {
	d := b.Center.Sub(&a.Center)
	l2 := d.Len2()
	r := a.Radius + b.Radius
	return l2 <= r*r
}

// AABBFromSphere returns the AABB bounding this sphere.
//
// NOTE: If you need to use this function you better start questioning the
// algorithm you're implementing as the sphere is both faster and bounds the
// underlying object better.
func AABBFromSphere(s *Sphere) AABB {
	return AABB{
		Center:     s.Center,
		HalfExtend: glm.Vec3{X: s.Radius, Y: s.Radius, Z: s.Radius},
	}
}

// MergePoint updates the bounding sphere to encompass v if needed.
func (sphere *Sphere) MergePoint(v *glm.Vec3) {
	// Compute squared distance between point and sphere center
	d := v.Sub(&sphere.Center)
	dist2 := d.Len2()
	// Only update s if point p is outside it
	if dist2 > sphere.Radius*sphere.Radius {
		dist := math.Sqrt(dist2)
		newRadius := (sphere.Radius + dist) * 0.5
		k := (newRadius - sphere.Radius) / dist
		sphere.Radius = newRadius
		sphere.Center.AddScaledVec(k, &d)
	}
}

// eigenSphere sets this sphere to the bounding sphere of the given point set
// using eigen values algorithm, this doesn't necessarily wrap all the points so
// use RitterEigenSphere.
func eigenSphere(points []glm.Vec3) Sphere {
	var m glm.Mat3

	// Compute the covariance matrix m
	CovarianceMatrix(&m, points)

	var v glm.Mat3
	// Decompose it into eigenvectors (in v) and eigenvalues (in m)
	Jacobi(&m, &v)

	// Find the component with largest magnitude eigenvalue (largest spread)
	maxe := math.Abs(m[0])

	var maxc int
	if maxf := math.Abs(m[3*1+1]); maxf > maxe {
		maxc = 1
		maxe = maxf
	}
	if maxf := math.Abs(m[3*2+2]); maxf > maxe {
		maxc = 2
		// not used but in the algorithm ?
		//maxe = maxf
	}

	var e glm.Vec3

	e.X = v[3*maxc+0]
	e.Y = v[3*maxc+1]
	e.Z = v[3*maxc+2]

	// Find the most extreme points along direction e
	imin, imax := ExtremePointsAlongDirection(&e, points)
	minpt := points[imin]
	maxpt := points[imax]
	u := maxpt.Sub(&minpt)
	dist := u.Len()

	var s Sphere
	s.Radius = dist * 0.5

	t := minpt.Add(&maxpt)
	s.Center.MulOf(0.5, &t)
	return s
}

// RitterEigenSphere sets this sphere to wrap all the given points using eigen
// values as base.
func RitterEigenSphere(points []glm.Vec3) Sphere {
	// Start with sphere from maximum spread
	s := eigenSphere(points)
	// Grow sphere to include all points
	for i := range points {
		s.MergePoint(&points[i])
	}
	return s
}
