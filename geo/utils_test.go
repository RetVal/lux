package geo

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/glm/glmtesting"
	"github.com/luxengine/lux/math"
	"testing"
)

func TestIsConvexQuad(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a, b, c, d glm.Vec3
		isconvex   bool
	}{
		{
			a:        glm.Vec3{X: 0, Y: 0, Z: 0},
			b:        glm.Vec3{X: 0, Y: 1, Z: 0},
			c:        glm.Vec3{X: 1, Y: 1, Z: 0},
			d:        glm.Vec3{X: 1, Y: 0, Z: 0},
			isconvex: true,
		},
		{
			a:        glm.Vec3{X: 0, Y: 0, Z: 0},
			b:        glm.Vec3{X: 1, Y: 1, Z: 0},
			c:        glm.Vec3{X: 0, Y: 1, Z: 0},
			d:        glm.Vec3{X: 1, Y: 0, Z: 0},
			isconvex: false,
		},
		{
			a:        glm.Vec3{X: 0, Y: 0, Z: 0},
			b:        glm.Vec3{X: 0, Y: 0, Z: 1},
			c:        glm.Vec3{X: 1, Y: 0, Z: 4},
			d:        glm.Vec3{X: 1, Y: 0, Z: 0},
			isconvex: true,
		},
		{
			a:        glm.Vec3{X: 0, Y: 0, Z: 0},
			b:        glm.Vec3{X: 0, Y: 1, Z: 0},
			c:        glm.Vec3{X: 0, Y: 4, Z: 1},
			d:        glm.Vec3{X: 0, Y: 0, Z: 1},
			isconvex: true,
		},
		{
			a:        glm.Vec3{X: 0, Y: 0, Z: 0},
			b:        glm.Vec3{X: 1, Y: 0, Z: 0},
			c:        glm.Vec3{X: 4, Y: 0, Z: 1},
			d:        glm.Vec3{X: 0, Y: 0, Z: 1},
			isconvex: true,
		},
		{
			a:        glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			b:        glm.Vec3{X: 1, Y: 0, Z: 0},
			c:        glm.Vec3{X: 4, Y: 0, Z: 1},
			d:        glm.Vec3{X: 0, Y: 0, Z: 1},
			isconvex: false,
		},
		{
			a:        glm.Vec3{X: 1, Y: 0, Z: 0},
			b:        glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			c:        glm.Vec3{X: 4, Y: 0, Z: 1},
			d:        glm.Vec3{X: 0, Y: 0, Z: 1},
			isconvex: false,
		},
		{
			a:        glm.Vec3{X: 4, Y: 0, Z: 1},
			b:        glm.Vec3{X: 1, Y: 0, Z: 0},
			c:        glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			d:        glm.Vec3{X: 0, Y: 0, Z: 1},
			isconvex: false,
		},
		{
			a:        glm.Vec3{X: 0, Y: 0, Z: 1},
			b:        glm.Vec3{X: 1, Y: 0, Z: 0},
			c:        glm.Vec3{X: 4, Y: 0, Z: 1},
			d:        glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			isconvex: false,
		},
	}

	for i, test := range tests {
		if isconvex := IsConvexQuad(&test.a, &test.b, &test.c, &test.d); isconvex != test.isconvex {
			t.Errorf("[%d] a(%v), b(%v), c(%v), d(%v) = %t, want %t", i,
				test.a, test.b, test.c, test.d, isconvex, test.isconvex)
		}
	}
}

func TestExtremePointsAlongDirection(t *testing.T) {
	t.Parallel()
	tests := []struct {
		direction  glm.Vec3
		points     []glm.Vec3
		imin, imax int
	}{
		{
			direction: glm.Vec3{X: 0, Y: 1, Z: 0},
			points: []glm.Vec3{
				{},
				{X: 4, Y: -9, Z: 0},
				{X: 2, Y: 1, Z: 0},
				{X: 5.4, Y: 7, Z: 0},
				{X: 1, Y: 2, Z: 0},
				{X: -4, Y: -5, Z: 0},
			},
			imin: 1,
			imax: 3,
		},
		{
			direction: glm.Vec3{X: 0, Y: -1, Z: 0},
			points: []glm.Vec3{
				{X: 0, Y: 0, Z: 0},
				{X: 4, Y: -9, Z: 0},
				{X: 2, Y: 1, Z: 0},
				{X: 5.4, Y: 7, Z: 0},
				{X: 1, Y: 2, Z: 0},
				{X: -4, Y: -5, Z: 0},
			},
			imin: 3,
			imax: 1,
		},
		{
			direction: glm.Vec3{X: 1, Y: 1, Z: 0},
			points: []glm.Vec3{
				{X: 0, Y: 0, Z: 0},
				{X: 10, Y: 10, Z: 0},
				{X: 0, Y: -10, Z: 0},
				{X: 1, Y: 0, Z: 0},
				{X: 0, Y: 1, Z: 0},
			},
			imin: 2,
			imax: 1,
		},
		{
			direction: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			points: []glm.Vec3{
				{X: 0, Y: 0, Z: 0},
				{X: 10, Y: 10, Z: 0},
				{X: 0, Y: -10, Z: 0},
				{X: 1, Y: 0, Z: 0},
				{X: 0, Y: 1, Z: 0},
			},
			imin: -1,
			imax: -1,
		},
		{
			direction: glm.Vec3{X: math.Inf(1), Y: 0, Z: 0},
			points: []glm.Vec3{
				{X: 0, Y: 0, Z: 0},
				{X: 10, Y: 10, Z: 0},
				{X: 0, Y: -10, Z: 0},
				{X: 1, Y: 0, Z: 0},
				{X: 0, Y: 1, Z: 0},
			},
			imin: -1,
			imax: 1,
		},
		{
			direction: glm.Vec3{X: 1, Y: 1, Z: 0},
			points: []glm.Vec3{
				{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
				{X: 10, Y: 10, Z: 0},
				{X: 0, Y: -10, Z: 0},
				{X: 1, Y: 0, Z: 0},
				{X: 0, Y: 1, Z: 0},
			},
			imin: 2,
			imax: 1,
		},
	}

	for i, test := range tests {
		imin, imax := ExtremePointsAlongDirection(&test.direction, test.points)
		if imin != test.imin || imax != test.imax {
			t.Errorf("[%d] direction(%v), points(%v) = %d, %d want %d, %d",
				i, test.direction, test.points, imin, imax, test.imin, test.imax)
		}
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		slice    []float32
		variance float32
	}{
		{
			[]float32{1, 1},
			0,
		},
		{
			[]float32{1, 2, 3},
			2.0 / 3.0,
		},
		{
			[]float32{1, 3, 5},
			8.0 / 3.0,
		},
		{
			[]float32{600, 470, 170, 430, 300},
			21704,
		},
		{
			[]float32{math.NaN(), 470, 170, 430, 300},
			math.NaN(),
		},
	}
	for i, test := range tests {
		if variance := Variance(test.slice); !glmtesting.FloatEqual(variance, test.variance) {
			t.Errorf("[%d] Variance(%v) = %f, want %f", i, test.slice, variance, test.variance)
		}
	}
}

func TestMinimumAreaRectangle(t *testing.T) {
	tests := []struct {
		points      []glm.Vec2
		minArea     float32
		center      glm.Vec2
		orientation [2]glm.Vec2
	}{
		{
			points:      []glm.Vec2{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {0.25, 0.25}, {0.5, 0.5}},
			minArea:     1,
			center:      glm.Vec2{X: 0.5, Y: 0.5},
			orientation: [2]glm.Vec2{{0, 1}, {-1, 0}},
		},
		{
			points:      []glm.Vec2{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {0.25, 0.25}, {0.5, 0.5}, {0.5, 1.5}},
			minArea:     1.5,
			center:      glm.Vec2{X: 0.5, Y: 0.75},
			orientation: [2]glm.Vec2{{0, 1}, {-1, 0}},
		},
		{
			points:      []glm.Vec2{{math.NaN(), 0}, {0, 1}, {1, 0}, {1, 1}, {0.25, 0.25}, {0.5, 0.5}, {0.5, 1.5}},
			minArea:     0,
			center:      glm.Vec2{X: math.NaN(), Y: math.NaN()},
			orientation: [2]glm.Vec2{{math.NaN(), math.NaN()}, {math.NaN(), math.NaN()}},
		},
	}
	for i, test := range tests {
		minArea, center, orientation := MinimumAreaRectangle(test.points)
		if !glmtesting.FloatEqual(minArea, test.minArea) {
			t.Errorf("[%d] minArea = %f, want %f", i, minArea, test.minArea)
		}

		if !glmtesting.Vec2Equal(center, test.center) {
			t.Errorf("[%d] center = %s, want %s", i, center.String(), test.center.String())
		}

		if !glmtesting.Vec2Equal(orientation[0], test.orientation[0]) {
			t.Errorf("[%d] orientation[0] = %s, want %s", i, orientation[0].String(), test.orientation[0].String())
		}
		if !glmtesting.Vec2Equal(orientation[1], test.orientation[1]) {
			t.Errorf("[%d] orientation[1] = %s, want %s", i, orientation[1].String(), test.orientation[1].String())
		}
	}
}

func TestClosestPointSegmentSegment(t *testing.T) {
	tests := []struct {
		p0, q0, p1, q1 glm.Vec3
		s, t, u        float32
		c0, c1         glm.Vec3
	}{
		{ // 0
			p0: glm.Vec3{X: 0, Y: 0, Z: 0},
			q0: glm.Vec3{X: 0, Y: 1, Z: 0},
			p1: glm.Vec3{X: 0, Y: 0, Z: 0},
			q1: glm.Vec3{X: 1, Y: 0, Z: 0},
			s:  0,
			t:  0,
			u:  0,
			c0: glm.Vec3{X: 0, Y: 0, Z: 0},
			c1: glm.Vec3{X: 0, Y: 0, Z: 0},
		},
		{ // 1
			p0: glm.Vec3{X: 0, Y: -0.5, Z: 0},
			q0: glm.Vec3{X: 0, Y: 0.5, Z: 0},
			p1: glm.Vec3{X: -0.5, Y: 0, Z: 0},
			q1: glm.Vec3{X: 0.5, Y: 0, Z: 0},
			s:  0.5,
			t:  0.5,
			u:  0,
			c0: glm.Vec3{},
			c1: glm.Vec3{},
		},
		{ // 2
			p0: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			q0: glm.Vec3{X: 0, Y: 0.5, Z: 0},
			p1: glm.Vec3{X: -0.5, Y: 0, Z: 0},
			q1: glm.Vec3{X: 0.5, Y: 0, Z: 0},
			s:  math.NaN(),
			t:  math.NaN(),
			u:  math.NaN(),
			c0: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			c1: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
		},
		{ // 3
			p0: glm.Vec3{X: 0, Y: -0.5, Z: 0},
			q0: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			p1: glm.Vec3{X: -0.5, Y: 0, Z: 0},
			q1: glm.Vec3{X: 0.5, Y: 0, Z: 0},
			s:  math.NaN(),
			t:  math.NaN(),
			u:  math.NaN(),
			c0: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			c1: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
		},
		{ // 4
			p0: glm.Vec3{X: 0, Y: -0.5, Z: 0},
			q0: glm.Vec3{X: 0, Y: 0.5, Z: 0},
			p1: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			q1: glm.Vec3{X: 0.5, Y: 0, Z: 0},
			s:  math.NaN(),
			t:  math.NaN(),
			u:  math.NaN(),
			c0: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			c1: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
		},
		{ // 5
			p0: glm.Vec3{X: 0, Y: -0.5, Z: 0},
			q0: glm.Vec3{X: 0, Y: 0.5, Z: 0},
			p1: glm.Vec3{X: -0.5, Y: 0, Z: 0},
			q1: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			s:  math.NaN(),
			t:  math.NaN(),
			u:  math.NaN(),
			c0: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			c1: glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
		},
		{ // 6
			p0: glm.Vec3{},
			q0: glm.Vec3{},
			p1: glm.Vec3{X: -0.5, Y: 0, Z: 0},
			q1: glm.Vec3{},
			s:  0,
			t:  1,
			u:  0,
			c0: glm.Vec3{},
			c1: glm.Vec3{},
		},
		{ // 7
			p0: glm.Vec3{},
			q0: glm.Vec3{X: -0.5, Y: 0, Z: 0},
			p1: glm.Vec3{},
			q1: glm.Vec3{},
			s:  -0,
			t:  0,
			u:  0,
			c0: glm.Vec3{},
			c1: glm.Vec3{},
		},
		{ // 8
			p0: glm.Vec3{},
			q0: glm.Vec3{},
			p1: glm.Vec3{},
			q1: glm.Vec3{},
			s:  0,
			t:  0,
			u:  0,
			c0: glm.Vec3{},
			c1: glm.Vec3{},
		},
		{ // 9
			p0: glm.Vec3{},
			q0: glm.Vec3{X: 0, Y: 1, Z: 0},
			p1: glm.Vec3{X: 0.1, Y: 0.5, Z: 0},
			q1: glm.Vec3{X: 1.1, Y: 0.5, Z: 0},
			s:  0.5,
			t:  0,
			u:  0.01000000070780516,
			c0: glm.Vec3{X: 0, Y: 0.5, Z: 0},
			c1: glm.Vec3{X: 0.1, Y: 0.5, Z: 0},
		},
		{ // 10
			p0: glm.Vec3{X: 0.1, Y: 0.5, Z: 0},
			q0: glm.Vec3{X: 1.1, Y: 0.5, Z: 0},
			p1: glm.Vec3{},
			q1: glm.Vec3{X: 0, Y: 1, Z: 0},
			s:  0,
			t:  0.5,
			u:  0.01000000070780516,
			c0: glm.Vec3{X: 0.1, Y: 0.5, Z: 0},
			c1: glm.Vec3{X: 0, Y: 0.5, Z: 0},
		},
		{ // 11
			p0: glm.Vec3{X: 0.1, Y: 0.5, Z: 0},
			q0: glm.Vec3{X: 1.1, Y: 0.5, Z: 0},
			p1: glm.Vec3{X: 2, Y: 0, Z: 0},
			q1: glm.Vec3{X: 2, Y: 1, Z: 0},
			s:  1,
			t:  0.5,
			u:  0.80999994277954102,
			c0: glm.Vec3{X: 1.1, Y: 0.5, Z: 0},
			c1: glm.Vec3{X: 2, Y: 0.5, Z: 0},
		},
	}
	for i, test := range tests {
		s, v, u, c0, c1 := ClosestPointSegmentSegment(&test.p0, &test.q0, &test.p1, &test.q1)
		if !glmtesting.FloatEqual(s, test.s) {
			t.Errorf("[%d] s = %f, want %f", i, s, test.s)
		}

		if !glmtesting.FloatEqual(v, test.t) {
			t.Errorf("[%d] t = %f, want %f", i, v, test.t)
		}

		if !glmtesting.FloatEqual(u, test.u) {
			t.Logf("%.17f", u)
			t.Errorf("[%d] u = %f, want %f", i, u, test.u)
		}

		if !glmtesting.Vec3Equal(c0, test.c0) {
			t.Errorf("[%d] c0 = %s, want %s", i, c0.String(), test.c0.String())
		}
		if !glmtesting.Vec3Equal(c1, test.c1) {
			t.Errorf("[%d] c1 = %s, want %s", i, c1.String(), test.c1.String())
		}

	}
}

func TestSqDistPointSegment(t *testing.T) {
	tests := []struct {
		a, b, c glm.Vec3
		sqdist  float32
	}{
		{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0}, math.NaN()},
		{glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: 1, Y: 0, Z: 0}, math.NaN()},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, math.NaN()},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0}, 1},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 2, Y: 0, Z: 0}, 4},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: -2, Y: 0, Z: 0}, 4},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: -2, Y: 1, Z: 0}, 4},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: -2, Y: 0.5, Z: 0}, 4},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: -2, Y: -1, Z: 0}, 5},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: -2, Y: 2, Z: 0}, 5},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 2, Y: -1, Z: 0}, 5},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 2, Y: 2, Z: 0}, 5},
	}
	for i, test := range tests {
		if sqdist := SqDistPointSegment(&test.a, &test.b, &test.c); !glmtesting.FloatEqual(sqdist, test.sqdist) {
			t.Errorf("[%d] sqdist = %f, want %f", i, sqdist, test.sqdist)
		}
	}
}

func TestClosestPointSegmentPoint(t *testing.T) {
	tests := []struct {
		a, b, c glm.Vec3
		v       float32
		point   glm.Vec3
	}{
		{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0},
			math.NaN(), glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: 1, Y: 0, Z: 0},
			math.NaN(), glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			math.NaN(), glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0},
			0, glm.Vec3{}},
	}
	for i, test := range tests {
		if v, point := ClosestPointSegmentPoint(&test.a, &test.b, &test.c); !glmtesting.Vec3Equal(point, test.point) || !glmtesting.FloatEqual(v, test.v) {
			t.Errorf("[%d] t, point = %f, %s, want %f, %s", i, v, point.String(), test.v, test.point.String())
		}
	}
}

func TestClosestPointPointRect2(t *testing.T) {
	tests := []struct {
		p, a, b, c glm.Vec3
		out        glm.Vec3
	}{
		{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: 2, Y: 0, Z: -1}, glm.Vec3{X: -1, Y: 0, Z: 2}, glm.Vec3{},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: -1, Y: 0, Z: 2}, glm.Vec3{},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 2, Y: 0, Z: -1}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 2, Y: 0, Z: -1}, glm.Vec3{X: -1, Y: 0, Z: 2}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 2, Y: 0, Z: -1}, glm.Vec3{X: -1, Y: 0, Z: 2},
			glm.Vec3{}},
		{glm.Vec3{X: 100, Y: 1, Z: 100}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 2, Y: 0, Z: -1}, glm.Vec3{X: -1, Y: 0, Z: 2},
			glm.Vec3{X: 2, Y: 0, Z: 2}},
	}
	for i, test := range tests {
		if out := ClosestPointPointRect2(&test.p, &test.a, &test.b, &test.c); !glmtesting.Vec3Equal(out, test.out) {
			t.Errorf("[%d] out = %s, want %s", i, out.String(), test.out.String())
		}
	}
}

func TestClosestPointPointTriangle(t *testing.T) {
	tests := []struct {
		p, a, b, c glm.Vec3
		out        glm.Vec3
	}{
		{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 2, Y: 0, Z: -1}, glm.Vec3{X: -1, Y: 0, Z: 2},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: 2, Y: 0, Z: -1}, glm.Vec3{X: -1, Y: 0, Z: 2},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: -1, Y: 0, Z: 2},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 2, Y: 0, Z: -1}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{}, glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0},
			glm.Vec3{}},
		{glm.Vec3{X: -1, Y: -1, Z: -1}, glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 0.33333328366279602, Y: 0.33333334326744080, Z: 0.33333334326744080}},
		{glm.Vec3{X: 2, Y: 0, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 1, Y: 0, Z: 0}},
		{glm.Vec3{X: 0, Y: 5, Z: 0}, glm.Vec3{X: 0, Y: 5, Z: 0}, glm.Vec3{X: 1, Y: 5, Z: 0}, glm.Vec3{X: 0, Y: 6, Z: 0},
			glm.Vec3{X: 0, Y: 5, Z: 0}},
		{glm.Vec3{X: -2, Y: 0, Z: -2}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 1, Y: 0, Z: -1},
			glm.Vec3{X: -1, Y: 0, Z: -1}},
		{glm.Vec3{X: 0, Y: 0, Z: 2}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 1, Y: 0, Z: -1},
			glm.Vec3{X: 0, Y: 0, Z: 1}},
		{glm.Vec3{X: 2, Y: 0, Z: -2}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 1, Y: 0, Z: -1},
			glm.Vec3{X: 1, Y: 0, Z: -1}},
		{glm.Vec3{X: -2, Y: 0, Z: 0}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 1, Y: 0, Z: -1},
			glm.Vec3{X: -0.8, Y: 0, Z: -0.6}},
		{glm.Vec3{X: 2, Y: 0, Z: 0}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 1, Y: 0, Z: -1},
			glm.Vec3{X: 0.8, Y: 0, Z: -0.6}},
		{glm.Vec3{X: 0, Y: 0, Z: -2}, glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 1, Y: 0, Z: -1},
			glm.Vec3{X: 0, Y: 0, Z: -1}},
	}
	for i, test := range tests {
		if out := ClosestPointPointTriangle(&test.p, &test.a, &test.b, &test.c); !glmtesting.Vec3Equal(out, test.out) {
			t.Errorf("[%d] out = %s, want %s", i, out.String(), test.out.String())
		}
	}
}
func TestPointOutsidePlane(t *testing.T) {
	tests := []struct {
		p, a, b, c glm.Vec3
		outside    bool
	}{
		{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			false},
		{glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{}, glm.Vec3{},
			false},
		{glm.Vec3{}, glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{},
			false},
		{glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			false},
		{glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{}, glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 0, Y: 0, Z: 1},
			false},
		{glm.Vec3{X: 0, Y: -1, Z: 0}, glm.Vec3{}, glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 0, Y: 0, Z: 1},
			true},
	}
	for i, test := range tests {
		if outside := PointOutsidePlane(&test.p, &test.a, &test.b, &test.c); !outside == test.outside {
			t.Errorf("[%d] outside = %t, want %t", i, outside, test.outside)
		}
	}
}

func TestPointsOnOppositeSideOfPlane(t *testing.T) {
	tests := []struct {
		p0, p1, a, b, c glm.Vec3
		opposite        bool
	}{
		{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			false},
		{glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			false},
		{glm.Vec3{}, glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{}, glm.Vec3{},
			false},
		{glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{},
			false},
		{glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
			false},
		{glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 0, Y: 0, Z: -1}, glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0},
			true},
	}
	for i, test := range tests {
		if opposite := PointsOnOppositeSideOfPlane(&test.p0, &test.p1, &test.a, &test.b, &test.c); opposite != test.opposite {
			t.Errorf("[%d] opposite = %t, want %t", i, opposite, test.opposite)
		}
	}
}

func TestClosestPointPointTetrahedron(t *testing.T) {
	tests := []struct {
		p, a, b, c, d glm.Vec3
		point         glm.Vec3
	}{
		{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
		{glm.Vec3{X: -1, Y: -1, Z: -1}, glm.Vec3{}, glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{X: 0, Y: 0, Z: 1},
			glm.Vec3{}},

		{glm.Vec3{X: 0, Y: -2, Z: 0}, glm.Vec3{X: -1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: -1, Z: 1}, glm.Vec3{X: 1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 0, Y: -1, Z: 0}},
		{glm.Vec3{X: -1, Y: 1, Z: -1}, glm.Vec3{X: -1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: -1, Z: 1}, glm.Vec3{X: 1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: -0.33333331346511841, Y: 0.33333337306976318, Z: -0.33333331346511841}},
		{glm.Vec3{X: 1, Y: 1, Z: -1}, glm.Vec3{X: -1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: -1, Z: 1}, glm.Vec3{X: 1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 0.33333331346511841, Y: 0.33333337306976318, Z: -0.33333331346511841}},
		{glm.Vec3{X: -1, Y: 1, Z: 1}, glm.Vec3{X: -1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: -1, Z: 1}, glm.Vec3{X: 1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 0, Y: 0.60000002384185791, Z: 0.20000000298023224}},
		{glm.Vec3{X: 1, Y: 1, Z: 1}, glm.Vec3{X: -1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: -1, Z: 1}, glm.Vec3{X: 1, Y: -1, Z: -1}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 0, Y: 0.60000002384185791, Z: 0.19999998807907104}},
	}
	for i, test := range tests {
		if point := ClosestPointPointTetrahedron(&test.p, &test.a, &test.b, &test.c, &test.d); !glmtesting.Vec3Equal(point, test.point) {
			t.Logf("%.17f,%.17f,%.17f", point.X, point.Y, point.Z)
			t.Errorf("[%d] point = %s, want %s", i, point.String(), test.point.String())
		}
	}
}

func TestTriangleAreaFromLengths(t *testing.T) {
	tests := []struct {
		a, b, c float32
		area    float32
	}{
		// http://www.mathopenref.com/heronsformula.html
		{math.NaN(), 1, 1, math.NaN()},
		{1, math.NaN(), 1, math.NaN()},
		{1, 1, math.NaN(), math.NaN()},
		{1, 1, 1, 0.43301269412040710},
		{1, 2, 2, 0.96824586391448975},
	}
	for i, test := range tests {
		if area := TriangleAreaFromLengths(test.a, test.b, test.c); !glmtesting.FloatEqual(area, test.area) {
			t.Errorf("[%d] area(%f, %f, %f) = %f, want %f", i, test.a, test.b, test.b, area, test.area)
		}
	}
}

func TestDistPointPlane(t *testing.T) {
	tests := []struct {
		p, a, b, c glm.Vec3
		dist       float32
	}{
		{glm.Vec3{X: 500, Y: 1, Z: 0}, glm.Vec3{}, glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 0, Y: 0, Z: 1},
			1},
	}
	for i, test := range tests {
		if dist := DistPointPlane(&test.p, &test.a, &test.b, &test.c); !glmtesting.FloatEqual(dist, test.dist) {
			t.Errorf("[%d] dist = %f, want %f", i, dist, test.dist)
		}
	}
}

func TestTestSpherePlane(t *testing.T) {
	tests := []struct {
		sphere    Sphere
		plane     Plane
		intersect bool
	}{
		{
			Sphere{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{
			Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, 0},
			false},
		{
			Sphere{glm.Vec3{}, math.NaN()},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{
			Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, math.NaN()},
			false},
		{
			Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			true},
		{
			Sphere{glm.Vec3{X: 0, Y: 1.5, Z: 0}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{
			Sphere{glm.Vec3{X: 0, Y: 1.5, Z: 0}, 1},
			Plane{glm.NormalizeVec3(glm.Vec3{X: 1, Y: 1, Z: 1}), 0},
			true},
		{
			Sphere{glm.Vec3{X: 0, Y: -1.5, Z: 0}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{
			Sphere{glm.Vec3{X: 0, Y: -1.5, Z: 0}, 1},
			Plane{glm.NormalizeVec3(glm.Vec3{X: 1, Y: 1, Z: 1}), 0},
			true},
	}
	for i, test := range tests {
		if intersect := TestPlaneSphere(&test.plane, &test.sphere); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestInsideSpherePlane(t *testing.T) {
	tests := []struct {
		sphere Sphere
		plane  Plane
		inside bool
	}{
		{ // 0
			Sphere{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{ // 1
			Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, 0},
			false},
		{ // 2
			Sphere{glm.Vec3{}, math.NaN()},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{ // 3
			Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, math.NaN()},
			false},
		{ // 4
			Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{ // 5
			Sphere{glm.Vec3{X: 0, Y: 1.5, Z: 0}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{ // 6
			Sphere{glm.Vec3{X: 0, Y: 1.5, Z: 0}, 1},
			Plane{glm.NormalizeVec3(glm.Vec3{X: 1, Y: 1, Z: 1}), 0},
			false},
		{ // 7
			Sphere{glm.Vec3{X: 0, Y: -1.5, Z: 0}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{ // 8
			Sphere{glm.Vec3{X: 0, Y: -1.5, Z: 0}, 1},
			Plane{glm.NormalizeVec3(glm.Vec3{X: 1, Y: 1, Z: 1}), 0},
			false},
		{ // 9
			Sphere{glm.Vec3{X: 0, Y: -2, Z: 0}, 1},
			Plane{glm.NormalizeVec3(glm.Vec3{X: 1, Y: 1, Z: 1}), 0},
			false},
	}
	for i, test := range tests {
		if inside := InsidePlaneSphere(&test.plane, &test.sphere); inside != test.inside {
			t.Errorf("[%d] inside = %t, want %t", i, inside, test.inside)
		}
	}
}

func TestTestSphereHalfspace(t *testing.T) {
	tests := []struct {
		sphere    Sphere
		plane     Plane
		intersect bool
	}{
		{Sphere{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false}, // 0
		{Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, 0},
			false}, // 1
		{Sphere{glm.Vec3{}, math.NaN()},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false}, // 2
		{Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, math.NaN()},
			false}, // 3
		{Sphere{glm.Vec3{}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			true}, // 4
		{Sphere{glm.Vec3{X: 0, Y: 1.5, Z: 0}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false}, // 5
		{Sphere{glm.Vec3{X: 0, Y: 1.5, Z: 0}, 1},
			Plane{glm.NormalizeVec3(glm.Vec3{X: 1, Y: 1, Z: 1}), 0},
			true}, // 6
		{Sphere{glm.Vec3{X: 0, Y: -1.5, Z: 0}, 1},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			true}, // 7
		{Sphere{glm.Vec3{X: 0, Y: -1.5, Z: 0}, 1},
			Plane{glm.NormalizeVec3(glm.Vec3{X: 1, Y: 1, Z: 1}), 0},
			true}, // 8
		{Sphere{glm.Vec3{X: 0, Y: -2, Z: 0}, 1},
			Plane{glm.NormalizeVec3(glm.Vec3{X: 1, Y: 1, Z: 1}), 0},
			true}, // 9
	}
	for i, test := range tests {
		if intersect := TestHalfspaceSphere(&test.plane, &test.sphere); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestTestOBBPlane(t *testing.T) {
	tests := []struct {
		obb       OBB
		plane     Plane
		intersect bool
	}{
		{
			OBB{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()},
				glm.Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false,
		},
		{
			OBB{glm.Vec3{},
				glm.Mat3{X: math.NaN(), Y: math.NaN(), Z: math.NaN(), 1, 0, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false,
		},
		{
			OBB{glm.Vec3{},
				glm.Mat3{0, 1, 0, 1, 0, 0, 0, 0, 1},
				glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false,
		},
		{
			OBB{glm.Vec3{},
				glm.Mat3{0, 1, 0, 1, 0, 0, 0, 0, 1},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			true,
		},
		{
			OBB{glm.Vec3{},
				glm.Mat3{0, 1, 0, 0, 0, 1, 1, 0, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			true,
		},
		{
			OBB{glm.Vec3{},
				glm.Mat3{0, 1, 0, 0, 0, 1, 1, 0, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 2},
			false,
		},
		{
			OBB{glm.Vec3{},
				glm.Mat3{0, 1, 0, 0, 0, 1, 1, 0, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 1, Y: 0, Z: 0}, 0},
			true,
		},
	}
	for i, test := range tests {

		if intersect := TestOBBPlane(&test.obb, &test.plane); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}

	}
}

func TestTestAABBPlane(t *testing.T) {
	tests := []struct {
		aabb      AABB
		plane     Plane
		intersect bool
	}{
		{AABB{glmtesting.NaN3, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{AABB{glm.Vec3{}, glmtesting.NaN3},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glmtesting.NaN3, 0},
			false},
		{AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, math.NaN()},
			false},
		{AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			true},
		{AABB{glm.Vec3{X: 0, Y: 5, Z: 0}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 5},
			true},
		{AABB{glm.Vec3{X: 0, Y: -5, Z: 0}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0},
			false},
		{AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 5},
			false},
		{AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			Plane{glm.Vec3{X: 1, Y: 0, Z: 0}, 0},
			true},
	}
	for i, test := range tests {
		if intersect := TestAABBPlane(&test.aabb, &test.plane); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestTestSphereAABB(t *testing.T) {
	tests := []struct {
		sphere    Sphere
		aabb      AABB
		intersect bool
	}{
		{Sphere{glm.Vec3{}, math.NaN()}, AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			false},
		{Sphere{glm.Vec3{}, 1}, AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true},
		{Sphere{glm.Vec3{X: 0, Y: 1, Z: 0}, 1}, AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true},
		{Sphere{glm.Vec3{X: 0, Y: 1.5, Z: 0}, 1}, AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true},
		{Sphere{glm.Vec3{X: 0, Y: 2, Z: 0}, 1}, AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			false},
		{Sphere{glm.Vec3{X: 0, Y: 2, Z: 0}, 1}, AABB{glm.Vec3{X: 0, Y: 2, Z: 0}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true},
	}
	for i, test := range tests {
		if intersect := TestAABBSphere(&test.aabb, &test.sphere); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestTestSphereOBB(t *testing.T) {
	//t.Skip("there seems to be a bug")
	tests := []struct {
		sphere    Sphere
		obb       OBB
		intersect bool
	}{
		{Sphere{glmtesting.NaN3, 1},
			OBB{},
			false}, // 0
		{Sphere{glm.Vec3{}, math.NaN()},
			OBB{},
			false}, // 1
		{Sphere{glm.Vec3{}, 1},
			OBB{glmtesting.NaN3,
				glm.Mat3{0, 1, 0, 0, 0, 1, 1, 0, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			false}, // 2
		{Sphere{glm.Vec3{}, 1},
			OBB{glm.Vec3{},
				glm.Mat3{math.NaN(), math.NaN(), math.NaN(), 0, 0, 1, 1, 0, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			false}, // 3
		{Sphere{glm.Vec3{}, 1},
			OBB{glm.Vec3{},
				glm.Mat3{0, 1, 0, 0, 0, 1, 1, 0, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true}, // 4
		{Sphere{glm.Vec3{X: 0, Y: 1, Z: 0}, 1},
			OBB{glm.Vec3{},
				glm.Mat3{0, 1, 0, 0, 0, 1, 1, 0, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true}, // 5
		{Sphere{glm.Vec3{X: 0, Y: 1, Z: 0}, 1},
			OBB{glm.Vec3{X: 0, Y: -2, Z: 0},
				glm.Mat3{0, 1, 0, 0, 0, 1, 1, 0, 0},
				glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			false}, // 6
	}
	for i, test := range tests {
		if intersect := TestOBBSphere(&test.obb, &test.sphere); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestTestSphereTriangle(t *testing.T) {
	tests := []struct {
		sphere    Sphere
		a, b, c   glm.Vec3
		intersect bool
	}{
		{Sphere{glmtesting.NaN3, 1},
			glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			false}, // 0
		{Sphere{glm.Vec3{}, math.NaN()},
			glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			false}, // 1
		{Sphere{glm.Vec3{}, 1},
			glmtesting.NaN3, glm.Vec3{}, glm.Vec3{},
			false}, // 2
		{Sphere{glm.Vec3{}, 1},
			glm.Vec3{}, glmtesting.NaN3, glm.Vec3{},
			false}, // 3
		{Sphere{glm.Vec3{}, 1},
			glm.Vec3{}, glm.Vec3{}, glmtesting.NaN3,
			false}, // 4

		{Sphere{glm.Vec3{}, 1},
			glm.Vec3{}, glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 0, Y: 1, Z: 0},
			true}, // 5
		{Sphere{glm.Vec3{X: 0, Y: 4, Z: 0}, 1},
			glm.Vec3{}, glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 0, Y: 1, Z: 0},
			false}, // 6
		{Sphere{glm.Vec3{X: 0, Y: 5, Z: 0}, 1},
			glm.Vec3{X: 0, Y: 5, Z: 0}, glm.Vec3{X: 1, Y: 5, Z: 0}, glm.Vec3{X: 0, Y: 6, Z: 0},
			true}, // 7
	}
	for i, test := range tests {
		if intersect := TestSphereTriangle(&test.sphere, &test.a, &test.b, &test.c); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestIntersectSegmentPlane(t *testing.T) {
	tests := []struct {
		a, b      glm.Vec3
		plane     Plane
		t         float32
		q         glm.Vec3
		intersect bool
	}{
		{glmtesting.NaN3, glm.Vec3{},
			Plane{glm.Vec3{}, 0},
			math.NaN(), glm.Vec3{}, false},
		{glm.Vec3{}, glmtesting.NaN3,
			Plane{glm.Vec3{}, 0},
			math.NaN(), glm.Vec3{}, false},
		{glm.Vec3{}, glm.Vec3{},
			Plane{glmtesting.NaN3, 0},
			math.NaN(), glm.Vec3{}, false},
		{glm.Vec3{}, glm.Vec3{},
			Plane{glm.Vec3{}, math.NaN()},
			math.NaN(), glm.Vec3{}, false},

		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0.5},
			0.5, glm.Vec3{X: 0, Y: 0.5, Z: 0}, true},

		{glm.Vec3{}, glm.Vec3{X: 0, Y: 0.4, Z: 0},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0.5},
			0, glm.Vec3{}, false},
		{glm.Vec3{X: 0, Y: 0.6, Z: 0}, glm.Vec3{X: 0, Y: 0.5, Z: 0},
			Plane{glm.Vec3{X: 0, Y: 1, Z: 0}, 0.5},
			1, glm.Vec3{X: 0, Y: 0.5, Z: 0}, true},
	}
	for i, test := range tests {
		v, q, intersect := IntersectSegmentPlane(&test.a, &test.b, &test.plane)
		if intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
		if !test.intersect {
			continue // if they don't overlap then t and q are junk data
		}

		if !glmtesting.FloatEqual(v, test.t) {
			t.Errorf("[%d] t = %f, want %f", i, v, test.t)
		}
		if !glmtesting.Vec3Equal(q, test.q) {
			t.Errorf("[%d] q = %s, want %s", i, q.String(), test.q.String())
		}
	}
}

func TestIntersectRaySphere(t *testing.T) {
	tests := []struct {
		p, d      glm.Vec3
		sphere    Sphere
		t         float32
		q         glm.Vec3
		intersect bool
	}{
		{glmtesting.NaN3, glm.Vec3{},
			Sphere{glm.Vec3{}, 0},
			0, glmtesting.NaN3, false},
		{glm.Vec3{}, glmtesting.NaN3,
			Sphere{glm.Vec3{}, 0},
			0, glmtesting.NaN3, false},
		{glm.Vec3{}, glm.Vec3{},
			Sphere{glmtesting.NaN3, 0},
			0, glmtesting.NaN3, false},
		{glm.Vec3{}, glm.Vec3{},
			Sphere{glm.Vec3{}, math.NaN()},
			0, glm.Vec3{}, false},

		{glm.Vec3{X: -10, Y: 1, Z: 0}, glm.Vec3{X: 10, Y: 1, Z: 0},
			Sphere{glm.Vec3{}, 2},
			0.49111938476562500, glm.Vec3{X: -5.08880615234375, Y: 1.491119384765625, Z: 0}, true}, // 4

		{glm.Vec3{X: -10, Y: 10, Z: 0}, glm.Vec3{X: 10, Y: 10, Z: 0},
			Sphere{glm.Vec3{}, 2},
			0, glm.Vec3{}, false}, // 5
		{glm.Vec3{X: 0, Y: 5, Z: 0}, glm.Vec3{X: 0, Y: 1, Z: 0},
			Sphere{glm.Vec3{}, 2},
			0, glm.Vec3{}, false}, // 6
		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0},
			Sphere{glm.Vec3{}, 2},
			0, glm.Vec3{}, true}, // 7
	}
	for i, test := range tests {
		v, q, intersect := IntersectRaySphere(&test.p, &test.d, &test.sphere)
		if intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
		if !test.intersect {
			continue // if they don't overlap then t and q are junk data
		}

		if !glmtesting.FloatEqual(v, test.t) {
			t.Errorf("[%d] t = %f, want %f", i, v, test.t)
		}
		if !glmtesting.Vec3Equal(q, test.q) {
			t.Errorf("[%d] q = %s, want %s", i, q.String(), test.q.String())
		}
	}
}

func TestTestRaySphere(t *testing.T) {
	tests := []struct {
		p, d      glm.Vec3
		sphere    Sphere
		intersect bool
	}{

		{glmtesting.NaN3, glm.Vec3{},
			Sphere{glm.Vec3{}, 0},
			false}, // 0
		/*{glm.Vec3{}, glmtesting.NaN3,
		Sphere{glm.Vec3{}, 1},
		false},*/ // this one actually doesn't get the opportunity to mess up
		{glm.Vec3{}, glm.Vec3{},
			Sphere{glmtesting.NaN3, 0},
			false}, // 1
		{glm.Vec3{}, glm.Vec3{},
			Sphere{glm.Vec3{}, math.NaN()},
			false}, // 2

		{glm.Vec3{X: -10, Y: 1, Z: 0}, glm.Vec3{X: 10, Y: 1, Z: 0},
			Sphere{glm.Vec3{}, 2},
			true}, // 3
		{glm.Vec3{X: -10, Y: 10, Z: 0}, glm.Vec3{X: 10, Y: 10, Z: 0},
			Sphere{glm.Vec3{}, 2},
			false}, // 4
		{glm.Vec3{X: 0, Y: 0, Z: -5}, glm.Vec3{X: 0, Y: 0.707, Z: 0.707},
			Sphere{glm.Vec3{X: 0, Y: 0, Z: 5}, 2},
			false}, // 5
		{glm.Vec3{X: 0, Y: 0, Z: 5}, glm.Vec3{X: 0, Y: 0.707, Z: 0.707},
			Sphere{glm.Vec3{X: 0, Y: 0, Z: 5}, 2},
			true}, // 6
	}
	for i, test := range tests {
		if intersect := TestRaySphere(&test.p, &test.d, &test.sphere); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestIntersectRayAABB(t *testing.T) {
	tests := []struct {
		p, d      glm.Vec3
		aabb      AABB
		t         float32
		q         glm.Vec3
		intersect bool
	}{
		{glmtesting.NaN3, glm.Vec3{},
			AABB{glm.Vec3{}, glm.Vec3{}},
			0, glm.Vec3{}, false},
		{glm.Vec3{}, glmtesting.NaN3,
			AABB{glm.Vec3{}, glm.Vec3{}},
			0, glm.Vec3{}, false},
		{glm.Vec3{}, glm.Vec3{},
			AABB{glmtesting.NaN3, glm.Vec3{}},
			0, glm.Vec3{}, false},
		{glm.Vec3{}, glm.Vec3{},
			AABB{glm.Vec3{}, glmtesting.NaN3},
			0, glm.Vec3{}, false},

		{glm.Vec3{X: -5, Y: 0, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			4.5, glm.Vec3{X: -0.5, Y: 0, Z: 0}, true},
		{glm.Vec3{X: 0, Y: 5, Z: 0}, glm.Vec3{X: 0, Y: -1, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			4.5, glm.Vec3{X: 0, Y: 0.5, Z: 0}, true},
		{glm.Vec3{X: 0, Y: 0, Z: -5}, glm.Vec3{X: 0, Y: 0, Z: 1},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			4.5, glm.Vec3{X: 0, Y: 0, Z: -0.5}, true},
		{glm.Vec3{X: 0, Y: 0, Z: -5}, glm.Vec3{X: 0, Y: 0, Z: 1},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: math.Inf(1)}},
			0, glm.Vec3{X: 0, Y: 0, Z: -5}, true},
	}
	for i, test := range tests {
		v, q, intersect := IntersectRayAABB(&test.p, &test.d, &test.aabb)
		if intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
		if !test.intersect {
			continue // if they don't overlap then t and q are junk data.
		}

		if !glmtesting.FloatEqual(v, test.t) {
			t.Errorf("[%d] t = %f, want %f", i, v, test.t)
		}
		if !glmtesting.Vec3Equal(q, test.q) {
			t.Errorf("[%d] q = %s, want %s", i, q.String(), test.q.String())
		}
	}
}

func TestTestSegmentAABB(t *testing.T) {
	tests := []struct {
		a, b      glm.Vec3
		aabb      AABB
		intersect bool
	}{
		{glmtesting.NaN3, glm.Vec3{},
			AABB{glm.Vec3{}, glm.Vec3{}},
			false},
		{glm.Vec3{}, glmtesting.NaN3,
			AABB{glm.Vec3{}, glm.Vec3{}},
			false},
		{glm.Vec3{}, glm.Vec3{},
			AABB{glmtesting.NaN3, glm.Vec3{}},
			false},
		{glm.Vec3{}, glm.Vec3{},
			AABB{glm.Vec3{}, glmtesting.NaN3},
			false},

		{glm.Vec3{X: -5, Y: 0, Z: 0}, glm.Vec3{X: 5, Y: 0, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			true},
		{glm.Vec3{X: -5, Y: 10, Z: 0}, glm.Vec3{X: 5, Y: 10, Z: 0},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			false},
		{glm.Vec3{X: 0, Y: 0, Z: 2}, glm.Vec3{X: 0, Y: 0, Z: 3},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}},
			false},
		{glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5}, glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5},
			AABB{glm.Vec3{}, glm.Vec3{X: 0.25, Y: 0.25, Z: 0.25}},
			false},
	}
	for i, test := range tests {
		if intersect := TestAABBSegment(&test.aabb, &test.a, &test.b); intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
	}
}

func TestIntersectSegmentTriangle(t *testing.T) {
	tests := []struct {
		p, q, a, b, c glm.Vec3
		u, v, w, t    float32
		intersect     bool
	}{
		{glmtesting.NaN3, glm.Vec3{},
			glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			math.NaN(), math.NaN(), math.NaN(), math.NaN(),
			true},
		{glm.Vec3{}, glmtesting.NaN3,
			glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			math.NaN(), math.NaN(), math.NaN(), math.NaN(),
			true},
		{glm.Vec3{}, glm.Vec3{},
			glmtesting.NaN3, glm.Vec3{}, glm.Vec3{},
			math.NaN(), math.NaN(), math.NaN(), math.NaN(),
			true},
		{glm.Vec3{}, glm.Vec3{},
			glm.Vec3{}, glmtesting.NaN3, glm.Vec3{},
			math.NaN(), math.NaN(), math.NaN(), math.NaN(),
			true},
		{glm.Vec3{}, glm.Vec3{},
			glm.Vec3{}, glm.Vec3{}, glmtesting.NaN3,
			math.NaN(), math.NaN(), math.NaN(), math.NaN(),
			true},

		{glm.Vec3{}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: 0, Y: 0.5, Z: 0}, glm.Vec3{X: 1, Y: 0.5, Z: 0}, glm.Vec3{X: 0, Y: 0.5, Z: 1},
			1, 0, 0, 0.5,
			true}, // 5
		{glm.Vec3{X: 0, Y: 0, Z: 1}, glm.Vec3{X: 0, Y: 1, Z: 1},
			glm.Vec3{X: 0, Y: 0.5, Z: 0}, glm.Vec3{X: 1, Y: 0.5, Z: 0}, glm.Vec3{X: 0, Y: 0.5, Z: 1},
			0, 0, 1, 0.5,
			true}, // 6
		/*{glm.Vec3{0, -1, 0}, glm.Vec3{0, -2, 0},
		glm.Vec3{-1, 0, -1}, glm.Vec3{0, 0, 1}, glm.Vec3{1, 0, -1},
		0, 1, 0, 0.5,
		true}, // 7*/
	}
	for i, test := range tests {
		{
			u, v, w, intersect := IntersectSegmentTriangle(&test.p, &test.q, &test.a, &test.b, &test.c)
			if intersect != test.intersect {
				t.Errorf("1. [%d] intersect = %t, want %t", i, intersect, test.intersect)
			}
			if !test.intersect {
				continue // if they don't overlap then t and q are junk data.
			}

			if !glmtesting.FloatEqual(u, test.u) {
				t.Errorf("1. [%d] u = %f, want %f", i, u, test.u)
			}
			if !glmtesting.FloatEqual(v, test.v) {
				t.Errorf("1. [%d] v = %f, want %f", i, v, test.v)
			}
			if !glmtesting.FloatEqual(w, test.w) {
				t.Errorf("1. [%d] w = %f, want %f", i, w, test.w)
			}
		}
		{
			u, v, w, q, intersect := IntersectSegmentTriangle2(&test.p, &test.q, &test.a, &test.b, &test.c)
			if intersect != test.intersect {
				t.Errorf("2. [%d] intersect = %t, want %t", i, intersect, test.intersect)
			}
			if !test.intersect {
				continue // if they don't overlap then t and q are junk data.
			}

			if !glmtesting.FloatEqual(u, test.u) {
				t.Errorf("2. [%d] u = %f, want %f", i, u, test.u)
			}
			if !glmtesting.FloatEqual(v, test.v) {
				t.Errorf("2. [%d] v = %f, want %f", i, v, test.v)
			}
			if !glmtesting.FloatEqual(w, test.w) {
				t.Errorf("2. [%d] w = %f, want %f", i, w, test.w)
			}
			if !glmtesting.FloatEqual(q, test.t) {
				t.Errorf("2. [%d] q = %f, want %f", i, q, test.t)
			}
		}
	}
}

func TestIntersectSegmentQuad(t *testing.T) {
	tests := []struct {
		p, q, a, b, c, d glm.Vec3
		point            glm.Vec3
		intersect        bool
	}{
		{glmtesting.NaN3, glm.Vec3{},
			glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			glmtesting.NaN3,
			true},
		{glm.Vec3{}, glmtesting.NaN3,
			glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			glmtesting.NaN3,
			true},
		{glm.Vec3{}, glm.Vec3{},
			glmtesting.NaN3, glm.Vec3{}, glm.Vec3{}, glm.Vec3{},
			glmtesting.NaN3,
			true},
		{glm.Vec3{}, glm.Vec3{},
			glm.Vec3{}, glmtesting.NaN3, glm.Vec3{}, glm.Vec3{},
			glmtesting.NaN3,
			true},
		{glm.Vec3{}, glm.Vec3{},
			glm.Vec3{}, glm.Vec3{}, glmtesting.NaN3, glm.Vec3{},
			glmtesting.NaN3,
			true},
		{glm.Vec3{}, glm.Vec3{},
			glm.Vec3{}, glm.Vec3{}, glm.Vec3{}, glmtesting.NaN3,
			glmtesting.NaN3,
			true},
		{glm.Vec3{X: 0, Y: -0.5, Z: 0}, glm.Vec3{X: 0, Y: 0.5, Z: 0},
			glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 1, Y: 0, Z: -1}, glm.Vec3{X: 1, Y: 0, Z: 1}, glm.Vec3{X: -1, Y: 0, Z: 1},
			glm.Vec3{},
			true}, // 6
		{glm.Vec3{X: 0.3, Y: -0.5, Z: 0}, glm.Vec3{X: 0.3, Y: 0.5, Z: 0},
			glm.Vec3{X: -1, Y: 0, Z: -1}, glm.Vec3{X: 1, Y: 0, Z: -1}, glm.Vec3{X: 1, Y: 0, Z: 1}, glm.Vec3{X: -1, Y: 0, Z: 1},
			glm.Vec3{X: 0.3, Y: 0, Z: 0},
			true}, // 7
	}
	for i, test := range tests {
		point, intersect := IntersectSegmentQuad(&test.p, &test.q, &test.a, &test.b, &test.c, &test.d)
		if intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
		if !test.intersect {
			continue // if they don't overlap then t and q are junk data.
		}

		if !glmtesting.Vec3Equal(point, test.point) {
			t.Errorf("[%d] point = %f, want %f", i, point, test.point)
		}
	}
}

func TestIntersectSegmentCylinder(t *testing.T) {
	tests := []struct {
		a, b, p, q glm.Vec3
		r          float32
		t          float32
		intersect  bool
	}{
		{glmtesting.NaN3, glm.Vec3{},
			glm.Vec3{}, glm.Vec3{}, 0,
			math.NaN(), false},
		{glm.Vec3{}, glmtesting.NaN3,
			glm.Vec3{}, glm.Vec3{}, 0,
			math.NaN(), false},
		{glm.Vec3{}, glm.Vec3{},
			glmtesting.NaN3, glm.Vec3{}, 0,
			math.NaN(), false},
		{glm.Vec3{}, glm.Vec3{},
			glm.Vec3{}, glmtesting.NaN3, 0,
			math.NaN(), false},
		{glm.Vec3{}, glm.Vec3{},
			glm.Vec3{}, glm.Vec3{}, math.NaN(),
			math.NaN(), false},

		{glm.Vec3{X: -10, Y: 0, Z: 0}, glm.Vec3{X: 10, Y: 0, Z: 0},
			glm.Vec3{X: 0, Y: -10, Z: 0}, glm.Vec3{X: 0, Y: 10, Z: 0}, 1,
			0.45, true},
		{glm.Vec3{X: -10, Y: 1, Z: 0}, glm.Vec3{X: 10, Y: 1, Z: 0},
			glm.Vec3{X: 0, Y: -10, Z: 0}, glm.Vec3{X: 0, Y: 10, Z: 0}, 1,
			0.45, true},
		{glm.Vec3{X: -10, Y: 0, Z: 1}, glm.Vec3{X: 10, Y: 0, Z: 1},
			glm.Vec3{X: 0, Y: -10, Z: 0}, glm.Vec3{X: 0, Y: 10, Z: 0}, 1,
			0.5, true},

		{glm.Vec3{X: 0, Y: 1, Z: 0}, glm.Vec3{},
			glm.Vec3{X: 0, Y: -10, Z: 0}, glm.Vec3{X: 0, Y: 10, Z: 0}, 1,
			0, true},
		{glm.Vec3{X: -5, Y: 5, Z: 0}, glm.Vec3{X: 5, Y: 5, Z: 0},
			glm.Vec3{X: 0, Y: -1, Z: 0}, glm.Vec3{X: 0, Y: 1, Z: 0}, 1,
			0, false},
		{glm.Vec3{X: 0, Y: -1, Z: 0}, glm.Vec3{X: 0, Y: 1, Z: 0},
			glm.Vec3{X: -5, Y: 5, Z: 0}, glm.Vec3{X: 5, Y: 5, Z: 0}, 1,
			0, false},
		{glm.Vec3{}, glm.Vec3{X: 2, Y: 0, Z: 0},
			glm.Vec3{X: 1, Y: 0, Z: 0}, glm.Vec3{X: 3, Y: 0, Z: 0}, 1,
			0.5, true}, // 11
		{glm.Vec3{X: 3, Y: 0, Z: 0}, glm.Vec3{X: 1, Y: 0, Z: 0},
			glm.Vec3{}, glm.Vec3{X: 2, Y: 0, Z: 0}, 1,
			0.5, true}, // 12
		/*{glm.Vec3{0.7, 0.8, 0}, glm.Vec3{0.5, 0.7, 0},
		glm.Vec3{0, -1, 0}, glm.Vec3{0, 1, 0}, 0.5,
		0, true}, // 13*/
	}
	for i, test := range tests {
		v, intersect := IntersectSegmentCylinder(&test.a, &test.b, &test.p, &test.q, test.r)
		if intersect != test.intersect {
			t.Errorf("[%d] intersect = %t, want %t", i, intersect, test.intersect)
		}
		if !test.intersect {
			continue // if they don't overlap then t and q are junk data.
		}

		if !glmtesting.FloatEqual(v, test.t) {
			t.Errorf("[%d] v = %f, want %f", i, v, test.t)
		}
	}
}

func TestIsPointInTriangle(t *testing.T) {
	tests := []struct {
		p, a, b, c glm.Vec3
		inside     bool
	}{
		{
			p:      glm.Vec3{X: 0, Y: 1, Z: 0},
			a:      glm.Vec3{X: -1, Y: 0, Z: -1},
			b:      glm.Vec3{X: 0, Y: 0, Z: 1},
			c:      glm.Vec3{X: 1, Y: 0, Z: -1},
			inside: true},
		{
			p:      glm.Vec3{X: -1, Y: 1, Z: -1},
			a:      glm.Vec3{X: -1, Y: 0, Z: -1},
			b:      glm.Vec3{X: 0, Y: 0, Z: 1},
			c:      glm.Vec3{X: 1, Y: 0, Z: -1},
			inside: true},
	}
	for i, test := range tests {
		inside := IsPointInTriangle(&test.p, &test.a, &test.b, &test.c)
		if inside != test.inside {
			t.Errorf("[%d] inside = %t, want %t", i, inside, test.inside)
		}
	}
}
