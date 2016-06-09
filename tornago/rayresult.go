package tornago

import (
	"github.com/luxengine/lux/glm"
)

// verify, at compile time, that these type implement RayResult.
var _ RayResult = &RayResultAny{}
var _ RayResult = &RayResultClosest{}
var _ RayResult = &RayResultAll{}

// RayResult can be passed to World.RayTest to capture the results of the test.
type RayResult interface {
	// This function is notified every time the algorithm finds a contact point.
	// it returns true if it wished the algorithm to continue.
	AddResult(*RigidBody, glm.Vec3) bool
}

// RayResultAny keeps only the first ray found by the algorithm (not
// necessarelly the closest).
type RayResultAny struct {
	Body *RigidBody
	Hit  glm.Vec3
}

// AddResult takes the result and replaces the currently saved result.
func (r *RayResultAny) AddResult(body *RigidBody, hit glm.Vec3) bool {
	r.Body = body
	r.Hit = hit
	return false
}

// RayResultClosest keeps only the rigid body with the closest contact point
// with to the origin.
type RayResultClosest struct {
	Origin glm.Vec3
	Body   *RigidBody
	Hit    glm.Vec3
	Len2   float32
}

// AddResult is notified when the engine finds a intersection. It always return
// true (indicating it will continue until the engine cannot find further
// intersection).
func (r *RayResultClosest) AddResult(b *RigidBody, hit glm.Vec3) bool {
	dist := hit.Sub(&r.Origin)
	l := dist.Len2()

	if r.Body == nil {
		r.Body = b
		r.Hit = hit
		r.Len2 = l
		return true
	}

	// if it's closer
	if l < r.Len2 {
		r.Body = b
		r.Hit = hit
		r.Len2 = l
	}

	return true
}

// RayResultAll keeps track of all ray intersection and sorts them.
type RayResultAll struct {
	Origin glm.Vec3
	Bodies []*RigidBody
	Points []glm.Vec3
	Len2s  []float32
}

// AddResult is notified when the engine finds a intersection, it adds the
// result to the sorted list. It always return true (indicating it will continue
// until the engine cannot find further intersection).
func (r *RayResultAll) AddResult(b *RigidBody, hit glm.Vec3) bool {
	// Since the API does not support any kind of post process on ray results
	// we need to sort as we get receive results.
	d := hit.Sub(&r.Origin)
	l2 := d.Len2()
	if len(r.Bodies) == 0 {
		r.Bodies = append(r.Bodies, b)
		r.Points = append(r.Points, hit)
		r.Len2s = append(r.Len2s, l2)
		return true
	}
	i := len(r.Len2s)
	for n := range r.Len2s {
		if l2 < r.Len2s[n] {
			i = n
			break
		}
	}

	r.Bodies = append(r.Bodies, nil)
	copy(r.Bodies[i+1:], r.Bodies[i:])
	r.Bodies[i] = b

	r.Points = append(r.Points, glm.Vec3{})
	copy(r.Points[i+1:], r.Points[i:])
	r.Points[i] = hit

	r.Len2s = append(r.Len2s, 0)
	copy(r.Len2s[i+1:], r.Len2s[i:])
	r.Len2s[i] = l2
	return true
}
