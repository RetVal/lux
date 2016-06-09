package geo

import (
	"reflect"
	"sort"
	"unsafe"

	"github.com/luxengine/lux/geo/internal/qhull"
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

func convexhullFromFaces(points []glm.Vec3, faces []*qhull.Face) *Convexhull {
	// reduce the vertices to the essential set
	vertexMap := make(map[glm.Vec3]int)
	for _, face := range faces {
		for n := 0; n < len(face.Vertices); n++ {
			if _, ok := vertexMap[points[face.Vertices[n]]]; !ok {
				vertexMap[points[face.Vertices[n]]] = len(vertexMap)
			}
		}
	}

	// calculate memory usage of the entire hull
	verticesSize := int(unsafe.Sizeof(glm.Vec3{})) * len(vertexMap)
	trianglesSize := int(unsafe.Sizeof(HullTriangle{})) * len(faces)
	convexhullSize := int(unsafe.Sizeof(Convexhull{}))

	// allocate the required memory
	mem := make([]byte, verticesSize+trianglesSize+convexhullSize)

	// make the 2 slice headers for the hull
	verticesSliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&mem[convexhullSize])),
		Len:  len(vertexMap),
		Cap:  len(vertexMap),
	}
	trianglesSliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&mem[convexhullSize+verticesSize])),
		Len:  len(faces),
		Cap:  len(faces),
	}

	hull := (*Convexhull)(unsafe.Pointer(&mem[0]))

	// build the continuous hull
	hull.Vertices = *(*[]glm.Vec3)(unsafe.Pointer(&verticesSliceHeader))
	hull.Triangles = *(*[]HullTriangle)(unsafe.Pointer(&trianglesSliceHeader))

	// copy the vertices over to the hull
	for vertex, index := range vertexMap {
		hull.Vertices[index] = vertex
	}

	// map all face->index in face slice
	faceMap := make(map[*qhull.Face]int)
	for i, face := range faces {
		faceMap[face] = i
	}

	// link all the triangles
	for i, face := range faces {
		for n := 0; n < len(face.Vertices); n++ {
			hull.Triangles[i].Vertices[n] = &hull.Vertices[vertexMap[points[face.Vertices[n]]]]
		}
		for n := 0; n < len(hull.Triangles[i].Adjacent); n++ {
			hull.Triangles[i].Adjacent[n] = &hull.Triangles[faceMap[face.Faces[n]]]
		}
	}

	return hull
}

// Quickhull returns the convex hull of the given points. The point slice will
// be modified, give a copy if you don't want your original data to be touched.
func Quickhull(points []glm.Vec3) *Convexhull {
	// 1 Find Initial tetrahedron
	// 1.1 Find Initial Triangle
	extremumIndices, extremums := qhull.FindExtremums(points)

	// 0.2 calculate the epsilon
	epsilon := qhull.CalculateEpsilon(extremums)

	// 1.2 Find the 3 most extreme points
	var triangleIndices [3]int
	var maxArea float32
	for i := 0; i < len(extremums); i++ {
		for j := i + 1; j < len(extremums); j++ {
			for k := j + 1; k < len(extremums); k++ {
				l1, l2, l3 := extremums[i].Sub(&extremums[j]), extremums[i].Sub(&extremums[k]), extremums[j].Sub(&extremums[k])
				area := TriangleAreaFromLengths(l1.Len(), l2.Len(), l3.Len())
				if area > maxArea {
					maxArea = area
					triangleIndices = [3]int{i, j, k}
				}
			}
		}
	}

	// 1.3 Finish the tetrahedron

	// find the last point
	l1, l2 := extremums[triangleIndices[1]].Sub(&extremums[triangleIndices[0]]), extremums[triangleIndices[2]].Sub(&extremums[triangleIndices[0]])
	dir := l1.Cross(&l2)

	imin, imax := ExtremePointsAlongDirection(&dir, points)
	vmin, vmax := points[imin].Sub(&extremums[triangleIndices[0]]), points[imax].Sub(&extremums[triangleIndices[0]])
	p1, p2 := math.Abs(vmin.Dot(&dir)), math.Abs(vmax.Dot(&dir))

	var tetraIndex int
	if p1 > p2 {
		tetraIndex = imin
	} else {
		tetraIndex = imax
	}

	// make the slice of indices that we have to deal with, and remove the
	// indices that are already in the tetrahedron.
	uneaten := make([]int, len(points))
	for i := range uneaten {
		uneaten[i] = i
	}

	tetraindices := []int{tetraIndex, extremumIndices[triangleIndices[0]], extremumIndices[triangleIndices[1]], extremumIndices[triangleIndices[2]]}
	sort.Ints(tetraindices)
	for n := len(tetraindices) - 1; n >= 0; n-- {
		uneaten = uneaten[:tetraindices[n]+copy(uneaten[tetraindices[n]:], uneaten[tetraindices[n]+1:])]
	}

	// find a point that will be inside everything, always
	var center glm.Vec3
	for _, i := range tetraindices {
		center.AddWith(&points[i])
	}
	center.MulWith(0.25) // 1.0 / len(tetraindices)

	faces := qhull.BuildInitialTetrahedron(tetraIndex, extremumIndices[triangleIndices[0]], extremumIndices[triangleIndices[1]], extremumIndices[triangleIndices[2]], points, &center)

	// initial partitioning
	var maxDist float32 = -math.MaxFloat32
	var maxDistFace, maxDistConflict int = -1, -1
	for n := 0; n < len(uneaten); n++ {
		var assigned bool
		for fn, face := range faces {
			if dist := qhull.DistToPlane(&face.Plane, &points[uneaten[n]]); dist > epsilon {
				face.Conflicts = append(face.Conflicts, qhull.Conflict{Distance: dist, Index: uneaten[n]})
				assigned = true
				if dist > maxDist {
					maxDist = dist
					maxDistFace = fn
					maxDistConflict = len(face.Conflicts) - 1
				}
			}
		}
		if !assigned {
			uneaten = uneaten[:n+copy(uneaten[n:], uneaten[n+1:])]
			n--
		}
	}

eat:
	if maxDistFace == -1 {
		return convexhullFromFaces(points, faces)
	}

	edges := qhull.FindHorizon(faces[maxDistFace], faces[maxDistFace].Edges[0], &points[faces[maxDistFace].Conflicts[maxDistConflict].Index])

	newfaces := make([]*qhull.Face, len(edges))
	{ // create the new triangles and link edges/faces togheter.
		for n := range edges {
			e0, e1, e2 := new(qhull.Edge), new(qhull.Edge), new(qhull.Edge)
			e0.Tail = edges[n].Tail
			e0.Next = e1
			e0.Prev = e2
			e0.Twin = edges[n]

			e1.Tail = edges[n].Next.Tail
			e1.Next = e2
			e1.Prev = e0

			e2.Tail = faces[maxDistFace].Conflicts[maxDistConflict].Index
			e2.Next = e0
			e2.Prev = e1

			f := &qhull.Face{
				Vertices: [3]int{edges[n].Tail, edges[n].Next.Tail, faces[maxDistFace].Conflicts[maxDistConflict].Index},
				Faces:    [3]*qhull.Face{edges[n].Twin.Face, nil, nil},
				Edges:    [3]*qhull.Edge{e0, e1, e2},
			}
			f.Plane = qhull.ComputePlane(&points[f.Vertices[0]], &points[f.Vertices[1]], &points[f.Vertices[2]])
			if qhull.DistToPlane(&f.Plane, &center) > 0 {
				f.Plane.Normal.Invert()
				f.Plane.Offset = -f.Plane.Offset
				f.Vertices[0], f.Vertices[1] = f.Vertices[1], f.Vertices[0]
			}
			e0.Face = f
			e1.Face = f
			e2.Face = f
			twin := edges[n].Twin
			edges[n].Twin = nil
			twin.Twin = f.Edges[0]
			f.Edges[0].Twin = twin
			edges[n] = twin

			faces = append(faces, f)
			newfaces[n] = f
		}
		for n := range edges {
			edges[n].Twin.Face.Edges[1].Twin = edges[(n+1)%len(edges)].Twin.Face.Edges[2]
			edges[n].Twin.Face.Edges[2].Twin = edges[(n-1+len(edges))%len(edges)].Twin.Face.Edges[1]
			edges[n].Twin.Face.Faces[1] = edges[(n+1)%len(edges)].Twin.Face
			edges[n].Twin.Face.Faces[2] = edges[(n-1+len(edges))%len(edges)].Twin.Face
		}
	}
	{ // reassign conflicts
		for _, face := range faces {
			if face.Visited {
				for _, conflict := range face.Conflicts {
					var assigned bool
					for _, newface := range newfaces {
						if qhull.DistToPlane(&newface.Plane, &points[conflict.Index]) > epsilon {
							newface.Conflicts = append(newface.Conflicts, conflict)
							assigned = true
						}
					}
					if !assigned { // if the conflict was resolved somehow during this iteration we can remove it.
						for i, uneat := range uneaten {
							if uneat == conflict.Index {
								uneaten = uneaten[:i+copy(uneaten[i:], uneaten[i+1:])]
							}
						}
					}
				}
			}
		}
	}
	{ // clean old faces.
		for n := 0; n < len(faces); n++ {
			if faces[n].Visited {
				faces = faces[:n+copy(faces[n:], faces[n+1:])]
				n--
			}
		}
	}
	{ // find the new king conflict.
		maxDist = -math.MaxFloat32
		maxDistFace, maxDistConflict = -1, -1
		for f, face := range faces {
			for c, conflict := range face.Conflicts {
				if conflict.Distance > maxDist {
					maxDist = conflict.Distance
					maxDistFace = f
					maxDistConflict = c
				}
			}
		}
	}

	//uneaten = make([]int, len(points))
	//for i := range uneaten {
	//	uneaten[i] = i
	//}

	// sanity check
	for _, face := range faces {
		if qhull.DistToPlane(&face.Plane, &center) > 0 {
			panic("bad construction")
		}
	}

	goto eat
}
