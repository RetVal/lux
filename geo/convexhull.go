package geo

import (
	"bytes"
	"fmt"
	"io"

	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// HullTriangle is a convex hull triangle.
type HullTriangle struct {
	// the 3 vertices that form this triangle
	Vertices [3]*glm.Vec3
	// triangles adjacent to this triangle.
	Adjacent [3]*HullTriangle
}

// Convexhull is a generic convex hull type. The primary way to generate these
// is via Quickhull.
type Convexhull struct {
	// the optimised vertex slice for this hull.
	Vertices []glm.Vec3

	// the triangles of the hull.
	Triangles []HullTriangle
	// the volume of the convex hull in unit/m^3
	volume float32

	// the center of mass of the convex hull.
	Center glm.Vec3

	// inertia tensor of the convex hull.
	inertia glm.Mat3
}

// ShapeType returns the shape type for aabbs.
func (*Convexhull) ShapeType() int {
	return convexhullShapeType
}

// Volume returns the volume of the hull.
func (c *Convexhull) Volume() float32 {
	return c.volume
}

// Mass returns the mass of the hull, in mass unit, given the density in
// (mass unit/distance unit^3).
func (c *Convexhull) Mass(density float32) float32 {
	return density * c.volume
}

// CalculateInternals calculates the volume, center of mass and the inertia
// tensor.
func (c *Convexhull) CalculateInternals() {
	const (
		oneDiv6   = 1.0 / 6.0
		oneDiv24  = 1.0 / 24.0
		oneDiv60  = 1.0 / 60.0
		oneDiv120 = 1.0 / 120.0
	)

	// order:  1, x, y, z, x^2, y^2, z^2, xy, yz, zx
	var integral [10]float32

	for _, tri := range c.Triangles {
		// Get vertices of triangle i.
		v0, v1, v2 := *tri.Vertices[0], *tri.Vertices[1], *tri.Vertices[2]

		// Get cross product of edges and normal vector.
		V1mV0 := v1.Sub(&v0)     // Vector3<Real>
		V2mV0 := v2.Sub(&v0)     // Vector3<Real>
		N := V1mV0.Cross(&V2mV0) // Vector3<Real>

		// Compute integral terms.
		var tmp0, tmp1, tmp2 float32
		var f1x, f2x, f3x, g0x, g1x, g2x float32
		tmp0 = v0.X + v1.X
		f1x = tmp0 + v2.X
		tmp1 = v0.X * v0.X
		tmp2 = tmp1 + v1.X*tmp0
		f2x = tmp2 + v2.X*f1x
		f3x = v0.X*tmp1 + v1.X*tmp2 + v2.X*f2x
		g0x = f2x + v0.X*(f1x+v0.X)
		g1x = f2x + v1.X*(f1x+v1.X)
		g2x = f2x + v2.X*(f1x+v2.X)

		var f1y, f2y, f3y, g0y, g1y, g2y float32
		tmp0 = v0.Y + v1.Y
		f1y = tmp0 + v2.Y
		tmp1 = v0.Y * v0.Y
		tmp2 = tmp1 + v1.Y*tmp0
		f2y = tmp2 + v2.Y*f1y
		f3y = v0.Y*tmp1 + v1.Y*tmp2 + v2.Y*f2y
		g0y = f2y + v0.Y*(f1y+v0.Y)
		g1y = f2y + v1.Y*(f1y+v1.Y)
		g2y = f2y + v2.Y*(f1y+v2.Y)

		var f1z, f2z, f3z, g0z, g1z, g2z float32
		tmp0 = v0.Z + v1.Z
		f1z = tmp0 + v2.Z
		tmp1 = v0.Z * v0.Z
		tmp2 = tmp1 + v1.Z*tmp0
		f2z = tmp2 + v2.Z*f1z
		f3z = v0.Z*tmp1 + v1.Z*tmp2 + v2.Z*f2z
		g0z = f2z + v0.Z*(f1z+v0.Z)
		g1z = f2z + v1.Z*(f1z+v1.Z)
		g2z = f2z + v2.Z*(f1z+v2.Z)

		// Update integrals.
		integral[0] += N.X * f1x
		integral[1] += N.X * f2x
		integral[2] += N.Y * f2y
		integral[3] += N.Z * f2z
		integral[4] += N.X * f3x
		integral[5] += N.Y * f3y
		integral[6] += N.Z * f3z
		integral[7] += N.X * (v0.Y*g0x + v1.Y*g1x + v2.Y*g2x)
		integral[8] += N.Y * (v0.Z*g0y + v1.Z*g1y + v2.Z*g2y)
		integral[9] += N.Z * (v0.X*g0z + v1.X*g1z + v2.X*g2z)
	}

	integral[0] *= oneDiv6
	integral[1] *= oneDiv24
	integral[2] *= oneDiv24
	integral[3] *= oneDiv24
	integral[4] *= oneDiv60
	integral[5] *= oneDiv60
	integral[6] *= oneDiv60
	integral[7] *= oneDiv120
	integral[8] *= oneDiv120
	integral[9] *= oneDiv120

	// mass
	c.volume = integral[0]

	// center of mass
	oomass := 1.0 / c.volume
	c.Center = glm.Vec3{
		X: integral[1] * oomass,
		Y: integral[2] * oomass,
		Z: integral[3] * oomass,
	}

	// inertia relative to world origin
	c.inertia[0] = integral[5] + integral[6]
	c.inertia[1] = -integral[7]
	c.inertia[2] = -integral[9]
	c.inertia[3] = c.inertia[1]
	c.inertia[4] = integral[4] + integral[6]
	c.inertia[5] = -integral[8]
	c.inertia[6] = c.inertia[2]
	c.inertia[7] = c.inertia[5]
	c.inertia[8] = integral[4] + integral[5]
	return
}

// MoveToOrigin moves the rigid body data so that it's center of mass is it's
// local origin. It calls CalculateInternals.
func (c *Convexhull) MoveToOrigin() {
	for n := range c.Vertices {
		c.Vertices[n].SubWith(&c.Center)
	}
	c.CalculateInternals()
}

// supportSlow returns the vertex that is the most in the direction of the given
// axis. This is the slow, reference implementation.
func (c *Convexhull) supportSlow(axis *glm.Vec3) glm.Vec3 {
	// very shitty simple implementation. No caching, no optimisation
	var max float32 = -math.MaxFloat32
	var vertex glm.Vec3
	for _, t := range c.Triangles {
		for n := 0; n < len(t.Vertices); n++ {
			if dist := axis.Dot(t.Vertices[n]); dist > max {
				max = dist
				vertex = *t.Vertices[n]
			}
		}
	}
	return vertex
}

// Support returns the vertex that is the most in the direction of the given
// axis.
func (c *Convexhull) Support(axis *glm.Vec3, cache *HullTriangle) (glm.Vec3, *HullTriangle) {
	if cache == nil {
		cache = &c.Triangles[0]
	}
	var max float32 = -math.MaxFloat32
	var vertex glm.Vec3
	// start with the cache
	for n := 0; n < 3; n++ {
		if dist := axis.Dot(cache.Vertices[n]); dist > max {
			max = dist
			vertex = *cache.Vertices[n]
		}
	}
	for {
	loop:
		for n := 0; n < 3; n++ {
			for m := 0; m < 3; m++ {
				if dist := axis.Dot(cache.Adjacent[n].Vertices[m]); dist > max {
					max = dist
					vertex = *cache.Adjacent[n].Vertices[m]
					cache = cache.Adjacent[n]
					continue loop
				}
			}
		}
		return vertex, cache
	}
}

// writeWavefront writes the faces to a writer as the .obj format.
// Import in blender with -Z forward and Y up.
func writeWavefront(writer io.Writer, hull *Convexhull) {
	pointsbuf, facebuf := new(bytes.Buffer), new(bytes.Buffer)

	vmap := make(map[*glm.Vec3]int)
	for i, v := range hull.Vertices {
		// +1 because wavefront starts at 1
		vmap[&hull.Vertices[i]] = i + 1
		fmt.Fprintf(pointsbuf, "v %f %f %f\n", v.X, v.Y, v.Z)
	}

	for _, triangle := range hull.Triangles {
		fmt.Fprintf(facebuf, "f %d %d %d\n", vmap[triangle.Vertices[0]], vmap[triangle.Vertices[1]], vmap[triangle.Vertices[2]])
	}

	fmt.Fprint(writer, pointsbuf.String())
	fmt.Fprint(writer, facebuf.String())
}

//	function GJK_intersection(shape p, shape q, vector initial_axis):
//       vector  A = Support(p, initial_axis) - Support(q, -initial_axis)
//       simplex s = {A}
//       vector  D = -A
//       loop:
//           A = Support(p, D) - Support(q, -D)
//           if dot(A, D) < 0:
//              reject
//           s = s ∪ A
//           s, D, contains_origin = NearestSimplex(s)
//           if contains_origin:
//              accept

// TestConvexhullConvexhull tests wether 2 convex hull intersects using GJK.
func TestConvexhullConvexhull(c0, c1 *Convexhull) bool {
	initial := c1.Center.Sub(&c0.Center)
	iinitial := initial.Inverse()
	s0, cache0 := c0.Support(&initial, nil)
	s1, cache1 := c1.Support(&iinitial, nil)
	A := s0.Sub(&s1)
	var s Simplex
	s.Merge(&A)
	D := A.Inverse()
	for {
		iD := D.Inverse()
		s0, cache0 = c0.Support(&D, cache0)
		s1, cache1 = c1.Support(&iD, cache1)
		A = s0.Sub(&s1)
		if A.Dot(&D) < 0 {
			return false
		}
		s.Merge(&A)
		var contain bool
		D, contain = s.NearestToOrigin()
		if contain {
			return true
		}
	}
}

// testConvexhullConvexhullSlow tests wether 2 convex hull intersects using GJK.
func testConvexhullConvexhullSlow(c0, c1 *Convexhull) bool {
	initial := c1.Center.Sub(&c0.Center)
	iinitial := initial.Inverse()
	s0 := c0.supportSlow(&initial)
	s1 := c1.supportSlow(&iinitial)
	A := s0.Sub(&s1)
	var s Simplex
	s.Merge(&A)
	D := A.Inverse()
	for {
		iD := D.Inverse()
		s0 = c0.supportSlow(&D)
		s1 = c1.supportSlow(&iD)
		A = s0.Sub(&s1)
		if A.Dot(&D) < 0 {
			return false
		}
		s.Merge(&A)
		var contain bool
		D, contain = s.NearestToOrigin()
		if contain {
			return true
		}
	}
}
