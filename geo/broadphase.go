package geo

import (
//	"github.com/luxengine/lux/glm"
)

/*
// Supporter return the vertex that is the most in the given direction.
type Supporter interface {
	Support(*glm.Vec3) glm.Vec3
}

// Broadphase is an spacial sorting algorithm used to improve performance of
// the narrow phase.
type Broadphase interface {
	// Insert the given object in the search space.
	Insert(Supporter)
	// Remove the object from the search space.
	Remove()

	// Update the search space (called between frames)
	Update()
}

// SAP (a.k.a. Sweep And Prune) is an algorithm that sorts elements on a set of
// orthogonal axis
// [Baraff92], [Cohen95]
type SAP struct {
	axis [3][]sapNode
}

type sapNode struct {
	// elem is a pointer to the object.
	elem Supporter
	// start is true if this element is the start of the object or false if it's
	// the end.
	pos   float32
	start bool
}

var sapAxis = [3][2]glm.Vec3{
	[2]glm.Vec3{{-1, 0, 0}, {1, 0, 0}},
	[2]glm.Vec3{{0, -1, 0}, {0, 1, 0}},
	[2]glm.Vec3{{0, 0, -1}, {0, 0, 1}},
}

// Insert the given object in the search space.
func (s *SAP) Insert(sup Supporter) {
	var supports [3][2]glm.Vec3
	for n := range supports {
		for m := range supports[n] {
			supports[n][m] = sup.Support(&sapAxis[n][m])
		}
	}
}

// Remove the object from the search space.
func (s *SAP) Remove() {}

// Update the search space (called between frames)
func (s *SAP) Update() {
	for n := 0; n < 3; n++ {
		for m := 0; m < len(s.axis[n])-1; m++ {
			var i int
			if !s.axis[n][m].start {
				i = 1
			}
			s.axis[n][m].pos = s.axis[n][m].elem.Support(&sapAxis[n][i]).I(n)
		}
	}

	// This is *technically O(n2) sorting, but temporal coherence should be this
	// almost O(n).
	for n := 0; n < 3; n++ {
		var m int
		for m < len(s.axis[n])-1 {
			if s.axis[n][m].pos < s.axis[n][m+1].pos {
				m++
				continue
			}
			s.axis[n][m], s.axis[n][m+1] = s.axis[n][m+1], s.axis[n][m]
			m-- // in case this needs to move more then 1 position.
		}
	}
}*/
