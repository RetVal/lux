package geo

import (
	"strconv"

	"github.com/luxengine/lux/glm"
)

// Simplex represents a simplex in 3D. Either a point, a line, a triangle,
// or a tetrahedron.
type Simplex struct {
	// the Points contained in the simplex. Data past Points[Size] is assumed to
	// be garbage.
	Points [4]glm.Vec3 // use an array to keep the memory all in 1 spot
	// the current extend of the simplex.
	Size int
}

// Merge merges the given vector to the simplex. This will panic if you add a
// 5th vertex.
func (s *Simplex) Merge(u *glm.Vec3) {
	s.Points[s.Size] = *u
	s.Size++
}

func (s *Simplex) nearestToOrigin4() (glm.Vec3, bool) {
	const (
		a = iota
		b
		c
		d
	)
	var zero glm.Vec3
	if PointsOnOppositeSideOfPlane(&zero, &s.Points[d], &s.Points[a], &s.Points[b], &s.Points[c]) {
		s.Size = 3
		return s.NearestToOrigin()
	}

	if PointsOnOppositeSideOfPlane(&zero, &s.Points[b], &s.Points[a], &s.Points[c], &s.Points[d]) {
		// a c d
		s.Size = 3
		s.Points[b] = s.Points[d]
		return s.NearestToOrigin()
	}

	if PointsOnOppositeSideOfPlane(&zero, &s.Points[c], &s.Points[a], &s.Points[d], &s.Points[b]) {
		s.Size = 3
		s.Points[c] = s.Points[d]
		return s.NearestToOrigin()
	}

	if PointsOnOppositeSideOfPlane(&zero, &s.Points[a], &s.Points[b], &s.Points[d], &s.Points[c]) {
		s.Size = 3
		s.Points[a] = s.Points[d]
		return s.NearestToOrigin()
	}

	// if its not outside anything its inside.
	return glm.Vec3{}, true
}

func (s *Simplex) nearestToOrigin3() (glm.Vec3, bool) {
	const (
		a = iota
		b
		c
		d
	)
	var zero glm.Vec3
	ab := s.Points[b].Sub(&s.Points[a])
	ac := s.Points[c].Sub(&s.Points[a])
	ap := zero.Sub(&s.Points[a])
	var bp, cp, closest glm.Vec3
	var va, vb, vc, d3, d4, d5, d6, denom, v, w float32

	// Check if P in vertex region outside A
	d1, d2 := ab.Dot(&ap), ac.Dot(&ap)
	if d1 <= 0 && d2 <= 0 {
		closest = s.Points[a] // barycentric coordinates (1, 0, 0)
		s.Size = 1
		//println("here1")
		goto checkclosest
	}

	bp = zero.Sub(&s.Points[b])
	d3, d4 = ab.Dot(&bp), ac.Dot(&bp)
	if d3 >= 0 && d4 <= d3 {
		closest = s.Points[b] // barycentric coordinates (0, 1, 0)
		s.Points[a] = s.Points[b]
		s.Size = 1
		//println("here2")
		goto checkclosest
	}

	// Check if P in edge region of AB, if so return projection of P onto AB.
	vc = d1*d4 - d3*d2
	if vc <= 0 && d1 >= 0 && d3 <= 0 {
		ret := s.Points[a]
		ret.AddScaledVec(d1/(d1-d3), &ab)
		closest = ret
		// ab, already good
		s.Size = 2
		//println("here3")
		goto checkclosest
	}

	// Check if P in vertex region outside C
	cp = zero.Sub(&s.Points[c])
	d5, d6 = ab.Dot(&cp), ac.Dot(&cp)
	if d6 >= 0 && d5 <= d6 {
		closest = s.Points[c] // barycentric coordinates (0, 0, 1)
		s.Size = 1
		s.Points[a] = s.Points[c]
		//println("here4")
		goto checkclosest
	}

	// Check if P in edge region of BC, if so return projection of P onto AC
	vb = d5*d2 - d1*d6
	if vb <= 0 && d2 >= 0 && d6 <= 0 {
		closest = s.Points[a]
		closest.AddScaledVec(d2/(d2-d6), &ac)
		s.Size = 2
		s.Points[b] = s.Points[c]
		//println("here5")
		goto checkclosest
	}

	// Check if P in edge region of BC, if so return projection of P onto BC
	va = d3*d6 - d5*d4
	if va <= 0 && (d4-d3) >= 0 && (d5-d6) >= 0 {
		bc := s.Points[c].Sub(&s.Points[b])
		closest = s.Points[b]
		closest.AddScaledVec((d4-d3)/((d4-d3)+(d5-d6)), &bc)
		s.Size = 2
		s.Points[a] = s.Points[c] // don't even move b
		//println("here6")
		goto checkclosest
	}

	// P inside face region. Compute Q through it's barycentric coordinates
	denom = 1 / (va + vb + vc)
	v = vb * denom
	w = vc * denom
	closest = s.Points[a]
	closest.AddScaledVec(v, &ab)
	closest.AddScaledVec(w, &ac)

checkclosest:
	// if the point is zero ...
	if closest.Equal(&zero) {
		// ... we contain the origin.
		return zero, true
	}
	// ... else
	closest.Normalize()
	return closest.Inverse(), false
}

func (s *Simplex) nearestToOrigin2() (glm.Vec3, bool) {
	const (
		a = iota
		b
		c
		d
	)
	var zero, closest glm.Vec3
	var t float32
	t, closest = ClosestPointSegmentPoint(&s.Points[a], &s.Points[b], &zero)

	// if it's at the beginning of ab ...
	if t == 0 {
		// .. no need to move points but reduce the simplex.
		s.Size = 1
	} else if t == 1 {
		// ... or at the end, move b to a and reduce
		s.Size = 1
		s.Points[a] = s.Points[b]
	}

	// if the point is zero ...
	if closest.Equal(&zero) {
		// ... we contain the origin.
		return zero, true
	}
	// ... else
	closest.Normalize()
	return closest.Inverse(), false
}

func (s *Simplex) nearestToOrigin1() (glm.Vec3, bool) {
	const (
		a = iota
		b
	)
	var zero glm.Vec3
	var closest glm.Vec3
	closest = s.Points[a]
	// if the point is zero ...
	if closest.Equal(&zero) {
		// ... then we can reduce to a zero simplex and contain.
		s.Size = 0
		return zero, true
	}
	// ... else no reduce and no contain.
	closest.Normalize()
	return closest.Inverse(), false
}

// NearestToOrigin modifies the simplex to contain only the minimum amount of
// points required to describe the direction to origin, it also returns the next
// direction to search in GJK and true if the origin is contained in the simplex
func (s *Simplex) NearestToOrigin() (direction glm.Vec3, containsOrigin bool) {
	switch s.Size {
	case 4:
		return s.nearestToOrigin4()
	case 3:
		return s.nearestToOrigin3()
	case 2:
		return s.nearestToOrigin2()
	case 1:
		return s.nearestToOrigin1()
	case 0:
		return glm.Vec3{}, true
	}
	panic("Simplex.Size=" + strconv.Itoa(int(s.Size)) + ", need 1, 2, 3, or 4")
}
