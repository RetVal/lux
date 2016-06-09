package geo

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestDOP8FromPoints(t *testing.T) {
	tests := []struct {
		points []glm.Vec3
		dop8   DOP8
	}{
		{[]glm.Vec3{{0, 0, 0}, {0, 1, 0}, {1, 0, 0}, {1, 1, 0}},
			DOP8{[4]float32{0, 0, -1, -1}, [4]float32{2, 2, 1, 1}}},
		{[]glm.Vec3{{5 + 0, 0, 0}, {5 + 0, 1, 0}, {5 + 1, 0, 0}, {5 + 1, 1, 0}},
			DOP8{[4]float32{5, 5, 4, -6}, [4]float32{7, 7, 6, -4}}},
		{[]glm.Vec3{{1, 1 + 1, 1 + 1}, {1, 1 + 1, 1}, {1, 1, 1 + 1}, {1, 1, 1},
			{1 + 1, 1 + 1, 1 + 1}, {1 + 1, 1 + 1, 1}, {1 + 1, 1, 1 + 1}, {1 + 1, 1, 1}},
			DOP8{[4]float32{3, 0, 0, 0}, [4]float32{6, 3, 3, 2}}},
	}
	for i, test := range tests {
		var d DOP8
		DOP8FromPoints(&d, test.points)
		if d != test.dop8 {
			t.Errorf("[%d] 8-dop = %+v, want %+v", i, d, test.dop8)
		}
	}
}

func TestTestDOP8DOP8(t *testing.T) {
	tests := []struct {
		a, b      DOP8
		intersect bool
	}{
		{DOP8{[4]float32{0, 0, -1, -1}, [4]float32{2, 2, 1, 1}},
			DOP8{[4]float32{0, 0, -1, -1}, [4]float32{2, 2, 1, 1}},
			true},
		{DOP8{[4]float32{0, 0, 0, -6}, [4]float32{7, 7, 6, 0}},
			DOP8{[4]float32{0, 0, 0, -6}, [4]float32{7, 7, 6, 0}},
			true},
		{
			DOP8{[4]float32{0, 0, -1, -1}, [4]float32{2, 2, 1, 1}},
			DOP8{[4]float32{3, 0, 0, 0}, [4]float32{6, 3, 3, 2}},
			false,
		},
	}
	for i, test := range tests {
		intersect := TestDOP8DOP8(&test.a, &test.b)
		if intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
		if !test.intersect {
			continue // if they don't overlap then t and q are junk data.
		}
	}
}
