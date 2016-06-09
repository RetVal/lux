package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"testing"
)

func TestSimplex_Merge(t *testing.T) {
	t.Parallel()
	tests := []struct {
		simplex Simplex
		point   glm.Vec3
		out     Simplex
	}{
		{Simplex{[4]glm.Vec3{{}, {}, {}, {}}, 0},
			glm.Vec3{0, 1, 0},
			Simplex{[4]glm.Vec3{{0, 1, 0}, {}, {}, {}}, 1}},
		{Simplex{[4]glm.Vec3{{0, 1, 0}, {}, {}, {}}, 1},
			glm.Vec3{0, 2, 0},
			Simplex{[4]glm.Vec3{{0, 1, 0}, {0, 2, 0}, {}, {}}, 2}},
		{Simplex{[4]glm.Vec3{{0, 1, 0}, {0, 2, 0}, {}, {}}, 2},
			glm.Vec3{0, 3, 0},
			Simplex{[4]glm.Vec3{{0, 1, 0}, {0, 2, 0}, {0, 3, 0}, {}}, 3}},
		{Simplex{[4]glm.Vec3{{0, 1, 0}, {0, 2, 0}, {0, 3, 0}, {}}, 3},
			glm.Vec3{0, 4, 0},
			Simplex{[4]glm.Vec3{{0, 1, 0}, {0, 2, 0}, {0, 3, 0}, {0, 4, 0}}, 4}},
	}
	for i, test := range tests {
		simplex := test.simplex
		simplex.Merge(&test.point)

		if simplex != test.out {
			t.Errorf("[%d] simplex = %+v, want %+v", i, simplex, test.out)
		}

	}
}

func TestSimplex_NearestToOrigin(t *testing.T) {
	tests := []struct {
		simplex        Simplex
		direction      glm.Vec3
		containsOrigin bool
		reduced        Simplex
	}{
		// 0 element tests
		{ // 0
			simplex:        Simplex{[4]glm.Vec3{{}, {}, {}, {}}, 0},
			direction:      glm.Vec3{},
			containsOrigin: true,
			reduced:        Simplex{[4]glm.Vec3{{}, {}, {}, {}}, 0},
		},
		// 1 element tests
		{ // 1
			simplex:        Simplex{[4]glm.Vec3{{0, 1, 0}, {}, {}, {}}, 1},
			direction:      glm.Vec3{0, -1, 0},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{0, 1, 0}, {}, {}, {}}, 1},
		},
		{ // 2
			simplex:        Simplex{[4]glm.Vec3{{0, 0, 0}, {}, {}, {}}, 1},
			direction:      glm.Vec3{0, 0, 0},
			containsOrigin: true,
			reduced:        Simplex{[4]glm.Vec3{{0, 0, 0}, {}, {}, {}}, 0},
		},
		// 2 element tests
		{ // 3
			simplex:        Simplex{[4]glm.Vec3{{0, 1, 0}, {1, 2, 0}, {}, {}}, 2},
			direction:      glm.Vec3{0, -1, 0},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{0, 1, 0}, {}, {}, {}}, 1},
		},
		{ // 4
			simplex:        Simplex{[4]glm.Vec3{{-2, 1, 0}, {-1, 1, 0}, {}, {}}, 2},
			direction:      glm.Vec3{0.70710676908493042, -0.70710676908493042, 0},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{-1, 1, 0}, {}, {}, {}}, 1},
		},
		{ // 5
			simplex:        Simplex{[4]glm.Vec3{{-2, 1, 0}, {-1, 2, 0}, {}, {}}, 2},
			direction:      glm.Vec3{0.70710682868957520, -0.70710682868957520, 0},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{-2, 1, 0}, {-1, 2, 0}, {}, {}}, 2},
		},
		{ // 6
			simplex:        Simplex{[4]glm.Vec3{{-1, 0, 0}, {1, 0, 0}, {}, {}}, 2},
			direction:      glm.Vec3{},
			containsOrigin: true,
			reduced:        Simplex{[4]glm.Vec3{{-2, 1, 0}, {-1, 2, 0}, {}, {}}, 2},
		},
		// 3 element tests
		{ // 7
			simplex:        Simplex{[4]glm.Vec3{{-1, 0, -1}, {0, 0, 1}, {1, 0, -1}, {}}, 3},
			direction:      glm.Vec3{},
			containsOrigin: true,
			reduced:        Simplex{[4]glm.Vec3{{}, {}, {}, {}}, 2},
		},
		{ // 8
			simplex:        Simplex{[4]glm.Vec3{{-1, 0, -1 - 5}, {0, 0, 1 - 5}, {1, 0, -1 - 5}, {}}, 3},
			direction:      glm.Vec3{0, 0, 1},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{0, 0, -4}, {}, {}, {}}, 1},
		},
		{ // 9
			simplex:        Simplex{[4]glm.Vec3{{-1 - 2, 0, -1}, {0 - 2, 0, 1}, {1 - 2, 0, -1}, {}}, 3},
			direction:      glm.Vec3{0.89442718029022217, 0, 0.44721359014511108},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{1 - 2, 0, -1}, {0 - 2, 0, 1}, {}, {}}, 2},
		},
		{ // 10
			simplex:        Simplex{[4]glm.Vec3{{-1 - 5, 0, -1 + 1}, {0 - 5, 0, 1 + 1}, {1 - 5, 0, -1 + 1}, {}}, 3},
			direction:      glm.Vec3{1, 0, 0},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{1 - 5, 0, -1 + 1}, {}, {}, {}}, 1},
		},
		{ // 11
			simplex:        Simplex{[4]glm.Vec3{{-1, 0, -1 + 5}, {0, 0, 1 + 5}, {1, 0, -1 + 5}, {}}, 3},
			direction:      glm.Vec3{0, 0, -1},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{-1, 0, -1 + 5}, {1, 0, -1 + 5}, {}, {}}, 2},
		},
		{ // 12
			simplex:        Simplex{[4]glm.Vec3{{-1 + 5, 0, -1}, {0 + 5, 0, 1}, {1 + 5, 0, -1}, {}}, 3},
			direction:      glm.Vec3{-0.97014254331588745, 0, 0.24253563582897186},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{-1 + 5, 0, -1}, {}, {}, {}}, 1},
		},
		{ // 13
			simplex:        Simplex{[4]glm.Vec3{{-1 + 2, 0, -1}, {0 + 2, 0, 1}, {1 + 2, 0, -1}, {}}, 3},
			direction:      glm.Vec3{-0.89442718029022217, 0, 0.44721359014511108},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{-1 + 2, 0, -1}, {0 + 2, 0, 1}, {}, {}}, 2},
		},
		{ // 14
			simplex:        Simplex{[4]glm.Vec3{{-1, -1, -1}, {0, -1, 1}, {1, -1, -1}, {}}, 3},
			direction:      glm.Vec3{0, 1, 0},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{-1, -1, -1}, {0, -1, 1}, {1, -1, -1}, {}}, 3},
		},
		// 4 element tests
		{ // 15
			simplex:        Simplex{[4]glm.Vec3{{-1, -1, -1}, {0, -1, 1}, {1, -1, -1}, {0, 1, 0}}, 4},
			direction:      glm.Vec3{},
			containsOrigin: true,
			reduced:        Simplex{[4]glm.Vec3{{}, {}, {}, {}}, 3},
		},
		{ // 16 botom forward
			simplex:        Simplex{[4]glm.Vec3{{-1, -1, -1 - 10}, {0, -1, 1 - 10}, {1, -1, -1 - 10}, {0, 1, 0 - 10}}, 4},
			direction:      glm.Vec3{0, 0.11043152213096619, 0.99388372898101807},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{0, -1, 1 - 10}, {}, {}, {}}, 1},
		},
		{ // 17 left vertex
			simplex:        Simplex{[4]glm.Vec3{{-1 + 10, -1, -1}, {0 + 10, -1, 1}, {1 + 10, -1, -1}, {0 + 10, 1, 0}}, 4},
			direction:      glm.Vec3{-0.98787838220596313, 0.10976426303386688, 0.10976426303386688},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{-1 + 10, -1, -1}, {}, {}, {}}, 1},
		},
		{ // 18 right vertex
			simplex:        Simplex{[4]glm.Vec3{{-1 - 10, -1, -1}, {0 - 10, -1, 1}, {1 - 10, -1, -1}, {0 - 10, 1, 0}}, 4},
			direction:      glm.Vec3{0.98787838220596313, 0.10976426303386688, 0.10976426303386688},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{1 - 10, -1, -1}, {}, {}, {}}, 1},
		},
		{ // 19 top vertex
			simplex:        Simplex{[4]glm.Vec3{{-1, -1 - 10, -1}, {0, -1 - 10, 1}, {1, -1 - 10, -1}, {0, 1 - 10, 0}}, 4},
			direction:      glm.Vec3{0, 1, 0},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{0, 1 - 10, 0}, {}, {}, {}}, 1},
		},
		{ // 20 botom face
			simplex:        Simplex{[4]glm.Vec3{{-1, -1 + 10, -1}, {0, -1 + 10, 1}, {1, -1 + 10, -1}, {0, 1 + 10, 0}}, 4},
			direction:      glm.Vec3{0, -1, 0},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{-1, -1 + 10, -1}, {0, -1 + 10, 1}, {1, -1 + 10, -1}, {}}, 3},
		},
		/*{ // 18
			simplex:        Simplex{[4]glm.Vec3{{-1, -1, -1 - 2}, {0, -1, 1 - 2}, {1, -1, -1 - 2}, {0, 1, 0 - 2}}, 4},
			direction:      glm.Vec3{0, 0.44721359014511108, 0.89442718029022217},
			containsOrigin: false,
			reduced:        Simplex{[4]glm.Vec3{{0, 1, 0 - 2}, {0, -1, 1 - 2}, {}, {}}, 2},
		},*/
	}
	for i, test := range tests {
		//if i != 16 {
		//	continue
		//}
		simplex := test.simplex
		direction, contain := simplex.NearestToOrigin()
		if contain != test.containsOrigin {
			t.Errorf("[%d] contain = %t, want %t", i, contain, test.containsOrigin)
		}
		if test.containsOrigin {
			continue
		}
		if direction.X == -0 {
			direction.X = 0
		}
		if direction.Y == -0 {
			direction.Y = 0
		}
		if direction.Z == -0 {
			direction.Z = 0
		}

		if !glmtesting.Vec3Equal(direction, test.direction) {
			t.Logf("%.17f,%.17f,%.17f", direction.X, direction.Y, direction.Z)
			t.Errorf("[%d] direction = %s, want %s", i, direction.String(), test.direction.String())
		}
		if simplex.Size != test.reduced.Size {
			t.Errorf("[%d] reduce size = %d, want %d", i, simplex.Size, test.reduced.Size)
			continue // don't check the points if the size is wrong
		}
		for n := 0; n < simplex.Size; n++ {
			if !glmtesting.Vec3Equal(simplex.Points[n], test.reduced.Points[n]) {
				t.Errorf("[%d] points[%d] = %s, want %s", i, n, simplex.Points[n].String(), test.reduced.Points[n].String())
			}
		}
	}
}

func TestSimplex_NearestToOrigin_Extra(t *testing.T) {
	t.Parallel()
	defer func() { recover() }()
	var s Simplex
	// force a panic
	s.Size = 5
	s.NearestToOrigin()
}
