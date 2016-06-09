package glmtesting

import (
	"github.com/luxengine/lux/flops"
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

// nan vectors
var (
	NaN2 = glm.Vec2{math.NaN(), math.NaN()}
	NaN3 = glm.Vec3{math.NaN(), math.NaN(), math.NaN()}
	NaN4 = glm.Vec4{math.NaN(), math.NaN(), math.NaN(), math.NaN()}
)

// FloatEqual returns true if v0 == v1 for every component. Will also return true
// when the components of both vectors are NaN.
func FloatEqual(v0, v1 float32) bool {
	if !flops.Eq(v0, v1) && !(math.IsNaN(v0) && math.IsNaN(v1)) {
		return false
	}
	return true
}

// Vec2Equal returns true if v0 == v1 for every component. Will also return true
// when the components of both vectors are NaN.
func Vec2Equal(v0, v1 glm.Vec2) bool {
	if !flops.Eq(v0.X, v1.X) && !(math.IsNaN(v0.X) && math.IsNaN(v1.X)) {
		return false
	}
	if !flops.Eq(v0.Y, v1.Y) && !(math.IsNaN(v0.Y) && math.IsNaN(v1.Y)) {
		return false
	}
	return true
}

// Vec3Equal returns true if v0 == v1 for every component. Will also return true
// when the components of both vectors are NaN.
func Vec3Equal(v0, v1 glm.Vec3) bool {
	if !flops.Eq(v0.X, v1.X) && !(math.IsNaN(v0.X) && math.IsNaN(v1.X)) {
		return false
	}
	if !flops.Eq(v0.Y, v1.Y) && !(math.IsNaN(v0.Y) && math.IsNaN(v1.Y)) {
		return false
	}
	if !flops.Eq(v0.Z, v1.Z) && !(math.IsNaN(v0.Z) && math.IsNaN(v1.Z)) {
		return false
	}
	return true
}

// Vec4Equal returns true if v0 == v1 for every component. Will also return true
// when the components of both vectors are NaN.
func Vec4Equal(v0, v1 glm.Vec4) bool {
	if !flops.Eq(v0.X, v1.X) && !(math.IsNaN(v0.X) && math.IsNaN(v1.X)) {
		return false
	}
	if !flops.Eq(v0.Y, v1.Y) && !(math.IsNaN(v0.Y) && math.IsNaN(v1.Y)) {
		return false
	}
	if !flops.Eq(v0.Z, v1.Z) && !(math.IsNaN(v0.Z) && math.IsNaN(v1.Z)) {
		return false
	}
	if !flops.Eq(v0.W, v1.W) && !(math.IsNaN(v0.W) && math.IsNaN(v1.W)) {
		return false
	}
	return true
}
