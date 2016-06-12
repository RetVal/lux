package main

import (
	"github.com/luxengine/lux/geo"
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
	"github.com/luxengine/lux/render"
	"testing"
)

func TestCameraCulling(t *testing.T) {
	var cam lux.Camera
	cameraAngle := glm.DegToRad(150)
	aspect := float32(3) / float32(3)
	var znear, zfar float32 = 0.1, 100.0
	cam.SetPerspective(cameraAngle, aspect, znear, zfar)
	var fr geo.Frustum
	geo.FrustumFromPerspective(cameraAngle, aspect, znear, zfar, &fr)

	for n := 0; n < 6; n++ {
		t.Log(fr.Planes[n])
	}
	t.Log("\n")

	aabb := geo.AABB{
		Center:     glm.Vec3{X: 0, Y: 0, Z: -15},
		HalfExtend: glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5},
	}

	const slice = 16
	for n := 0; n < slice; n++ {
		cam.LookAtval(
			0, 0, 0,
			math.Sin((float32(math.Pi)*2/float32(slice))*float32(n)), 0, math.Cos((float32(math.Pi)*2/float32(slice))*float32(n)),
			0, 1, 0)
		t.Errorf("%t", geo.TestAABBFrustum(&aabb, &fr, &cam.View))
		if geo.TestAABBFrustum(&aabb, &fr, &cam.View) {
			t.Errorf("%f %f %f", math.Sin((float32(math.Pi)*2/float32(slice))*float32(n)), 0.0, math.Cos((float32(math.Pi)*2/float32(slice))*float32(n)))
		}
	}

}
