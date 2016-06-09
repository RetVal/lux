package qhull

import (
	"fmt"
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

var _ = fmt.Errorf

const (
	epsilonbase = 0.0001 // arbitrary value that will require testing
)

type (
	// Edge is a quickhull utility struct for edges
	Edge struct {
		Tail             int
		Prev, Next, Twin *Edge
		Face             *Face
	}
	// Face is a quickhull utility struct for faces
	Face struct {
		Edges     [3]*Edge
		Faces     [3]*Face
		Vertices  [3]int
		Conflicts []Conflict

		Plane
		// tmp variables
		Visited bool
	}

	// Conflict is a vertex that isn't inside the convex hull yet
	Conflict struct {
		Distance float32
		Index    int
	}
)

// FindHorizon finds the horizon of the conflict
func FindHorizon(face *Face, edge *Edge, point *glm.Vec3) []*Edge {
	face.Visited = true
	var horizon []*Edge
	for n := 0; n < 3; n++ {
		if DistToPlane(&edge.Twin.Face.Plane, point) < 0 {
			horizon = append(horizon, edge)
			edge = edge.Next
			continue
		} else {
		}
		// if we can see it
		if edge.Twin.Face.Visited {
			edge = edge.Next
			continue
		}

		horizon = append(horizon, FindHorizon(edge.Twin.Face, edge.Twin, point)...)
		edge = edge.Next
	}

	return horizon
}

// FindExtremums returns the 6 indices and 6 vec3 of the extremums for each axis
// fomatted [minx, miny, minz, maxx, maxy, maxz].
func FindExtremums(points []glm.Vec3) (extremumIndices [6]int, extremums [6]glm.Vec3) {
	extremums = [6]glm.Vec3{
		{math.MaxFloat32, 0, 0}, {0, math.MaxFloat32, 0}, {0, 0, math.MaxFloat32},
		{-math.MaxFloat32, 0, 0}, {0, -math.MaxFloat32, 0}, {0, 0, -math.MaxFloat32},
	}
	for i := range points {
		for n := 0; n < 3; n++ {
			if *extremums[n].I(n) > *points[i].I(n) {
				extremums[n] = points[i]
				extremumIndices[n] = i
			}
			if *extremums[n].I(n) > *points[i].I(n) {
				extremums[n] = points[i]
				extremumIndices[n] = i
			}
			if *extremums[n].I(n) > *points[i].I(n) {
				extremums[n] = points[i]
				extremumIndices[n] = i
			}

			if *extremums[3+n].I(n) < *points[i].I(n) {
				extremums[3+n] = points[i]
				extremumIndices[3+n] = i
			}
			if *extremums[3+n].I(n) < *points[i].I(n) {
				extremums[3+n] = points[i]
				extremumIndices[3+n] = i
			}
			if *extremums[3+n].I(n) < *points[i].I(n) {
				extremums[3+n] = points[i]
				extremumIndices[3+n] = i
			}
		}
	}
	return
}

// CalculateEpsilon calculates the epsilon the algorithm should use given the
// extremums of the point cloud [minx, miny, minz, maxx, maxy, maxz]
func CalculateEpsilon(extremums [6]glm.Vec3) float32 {
	var maxima float32

	maxima += math.Max(math.Abs(extremums[0].X), math.Abs(extremums[3+0].X))
	maxima += math.Max(math.Abs(extremums[1].Y), math.Abs(extremums[3+1].Y))
	maxima += math.Max(math.Abs(extremums[2].Z), math.Abs(extremums[3+2].Z))

	return epsilonbase * maxima * 3
}

// BuildInitialTetrahedron builds the initial tetrahedron from the given 4 indices
func BuildInitialTetrahedron(a, b, c, d int, points []glm.Vec3, center *glm.Vec3) []*Face {
	pts := [4][3]int{
		{a, b, c},
		{c, a, d},
		{b, d, a},
		{d, b, c},
	}
	var planes [4]Plane
	for n := 0; n < len(pts); n++ {
		planes[n] = ComputePlane(&points[pts[n][0]], &points[pts[n][1]], &points[pts[n][2]])
		if DistToPlane(&planes[n], center) > 0 {
			planes[n].Normal.Invert()
			planes[n].Offset = -planes[n].Offset
			pts[n] = [3]int{pts[n][0], pts[n][2], pts[n][1]}
		}
	}

	f0 := &Face{Vertices: pts[0], Plane: planes[0]}
	f1 := &Face{Vertices: pts[1], Plane: planes[1]}
	f2 := &Face{Vertices: pts[2], Plane: planes[2]}
	f3 := &Face{Vertices: pts[3], Plane: planes[3]}

	f0.Faces = [3]*Face{f2, f3, f1}
	f1.Faces = [3]*Face{f3, f2, f0}
	f2.Faces = [3]*Face{f0, f1, f3}
	f3.Faces = [3]*Face{f1, f0, f2}

	// edges of f0
	e00 := &Edge{Tail: pts[0][0], Face: f0}
	e01 := &Edge{Tail: pts[0][1], Face: f0}
	e02 := &Edge{Tail: pts[0][2], Face: f0}

	// edges of f1
	e10 := &Edge{Tail: pts[1][0], Face: f1}
	e11 := &Edge{Tail: pts[1][1], Face: f1}
	e12 := &Edge{Tail: pts[1][2], Face: f1}

	// edges of f2
	e20 := &Edge{Tail: pts[2][0], Face: f2}
	e21 := &Edge{Tail: pts[2][1], Face: f2}
	e22 := &Edge{Tail: pts[2][2], Face: f2}

	// edges of f3
	e30 := &Edge{Tail: pts[3][0], Face: f3}
	e31 := &Edge{Tail: pts[3][1], Face: f3}
	e32 := &Edge{Tail: pts[3][2], Face: f3}

	// Connect the faces to the edges
	f0.Edges = [3]*Edge{e00, e01, e02}
	f1.Edges = [3]*Edge{e10, e11, e12}
	f2.Edges = [3]*Edge{e20, e21, e22}
	f3.Edges = [3]*Edge{e30, e31, e32}

	//Setup twin edges
	e00.Twin, e20.Twin = e20, e00
	e01.Twin, e31.Twin = e31, e01
	e02.Twin, e12.Twin = e12, e02
	e10.Twin, e30.Twin = e30, e10
	e11.Twin, e21.Twin = e21, e11
	e22.Twin, e32.Twin = e32, e22

	// Circular connect the edges
	// e0*
	e00.Next, e00.Prev = e01, e02
	e01.Next, e01.Prev = e02, e00
	e02.Next, e02.Prev = e00, e01

	// e1*
	e10.Next, e10.Prev = e11, e12
	e11.Next, e11.Prev = e12, e10
	e12.Next, e12.Prev = e10, e11

	// e2*
	e20.Next, e20.Prev = e21, e22
	e21.Next, e21.Prev = e22, e20
	e22.Next, e22.Prev = e20, e21

	// e3*
	e30.Next, e30.Prev = e31, e32
	e31.Next, e31.Prev = e32, e30
	e32.Next, e32.Prev = e30, e31

	return []*Face{f0, f1, f2, f3}
}

// Plane is a hyperplane in 3D.
type Plane struct {
	Normal glm.Vec3
	Offset float32
}

// ComputePlane Given three noncollinear points (ordered ccw), compute plane
// equation.
func ComputePlane(a, b, c *glm.Vec3) Plane {
	ab, ac := b.Sub(a), c.Sub(a)
	abac := ab.Cross(&ac)
	abac.Normalize()

	return Plane{
		Normal: abac,
		Offset: abac.Dot(a),
	}
}

// DistToPlane returns the signed distance of point to plane.
func DistToPlane(plane *Plane, point *glm.Vec3) float32 {
	return plane.Normal.Dot(point) - plane.Offset
}
