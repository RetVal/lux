package glmtesting

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
	"testing"
)

func TestFloatEqual(t *testing.T) {
	tests := []struct {
		a, b   float32
		result bool
	}{
		{0, 0, true},
		{math.NaN(), 0, false},
		{0, math.NaN(), false},
		{math.NaN(), math.NaN(), true},
	}
	for i, test := range tests {
		if res := FloatEqual(test.a, test.b); res != test.result {
			t.Errorf("[%d] result = %t, want %t", i, res, test.result)
		}
	}
}

func TestVec2Equal(t *testing.T) {
	tests := []struct {
		a, b   glm.Vec2
		result bool
	}{
		{glm.Vec2{}, glm.Vec2{}, true},
		{glm.Vec2{X: math.NaN(), Y: 0}, glm.Vec2{}, false},
		{glm.Vec2{X: 0, Y: math.NaN()}, glm.Vec2{}, false},
		{glm.Vec2{X: math.NaN(), Y: math.NaN()}, glm.Vec2{X: math.NaN(), Y: math.NaN()}, true},
	}
	for i, test := range tests {
		if res := Vec2Equal(test.a, test.b); res != test.result {
			t.Errorf("[%d] result = %t, want %t", i, res, test.result)
		}
	}
}

func TestVec3Equal(t *testing.T) {
	tests := []struct {
		a, b   glm.Vec3
		result bool
	}{
		{glm.Vec3{}, glm.Vec3{}, true},
		{glm.Vec3{X: math.NaN(), Y: 0, Z: 0}, glm.Vec3{}, false},
		{glm.Vec3{X: 0, Y: math.NaN(), Z: 0}, glm.Vec3{}, false},
		{glm.Vec3{X: 0, Y: 0, Z: math.NaN()}, glm.Vec3{}, false},
		{glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, glm.Vec3{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}, true},
	}
	for i, test := range tests {
		if res := Vec3Equal(test.a, test.b); res != test.result {
			t.Errorf("[%d] result = %t, want %t", i, res, test.result)
		}
	}
}

func TestVec4Equal(t *testing.T) {
	tests := []struct {
		a, b   glm.Vec4
		result bool
	}{
		{glm.Vec4{}, glm.Vec4{}, true},
		{glm.Vec4{X: math.NaN(), Y: 0, Z: 0, W: 0}, glm.Vec4{}, false},
		{glm.Vec4{X: 0, Y: math.NaN(), Z: 0, W: 0}, glm.Vec4{}, false},
		{glm.Vec4{X: 0, Y: 0, Z: math.NaN(), W: 0}, glm.Vec4{}, false},
		{glm.Vec4{X: 0, Y: 0, Z: 0, W: math.NaN()}, glm.Vec4{}, false},
		{glm.Vec4{X: math.NaN(), Y: math.NaN(), Z: math.NaN(), W: math.NaN()}, glm.Vec4{X: math.NaN(), Y: math.NaN(), Z: math.NaN(), W: math.NaN()}, true},
	}
	for i, test := range tests {
		if res := Vec4Equal(test.a, test.b); res != test.result {
			t.Errorf("[%d] result = %t, want %t", i, res, test.result)
		}
	}
}
