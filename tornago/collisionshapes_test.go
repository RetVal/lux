package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
	"testing"
)

var _ CollisionShape = &CollisionBox{}
var _ CollisionShape = &CollisionSphere{}

func TestCollisionBox_GetBoundingVolume(t *testing.T) {
	b := NewCollisionBox(glm.Vec3{X: 1, Y: 2, Z: 3})
	b.body = &RigidBody{position: glm.Vec3{X: 5, Y: 5, Z: 5}}
	vol := b.GetBoundingVolume()
	if vol.Center() != (glm.Vec3{X: 5, Y: 5, Z: 5}) {
		t.Error("center not as expected")
	}
	if vol.radius != 3 {
		t.Error("radius not as expected")
	}
}

// RayResultNothing does nothing with the result, mostly used for benchmarking.
type RayResultNothing struct{}

// AddResult does nothing, just implements RayResult interface.
func (RayResultNothing) AddResult(*RigidBody, glm.Vec3) bool { return true }

func TestCollisionSphere_RayTest(t *testing.T) {
	var originBody, movedBody, rotatedBody, movedRotatedBody RigidBody
	movedBody.SetPosition3f(5, 5, 5)
	movedRotatedBody.SetPosition3f(5, 5, 5)
	qi := glm.QuatIdent()
	originBody.SetOrientationQuat(&qi)
	movedBody.SetOrientationQuat(&qi)
	q2 := glm.Quat{W: 0.75, Vec3: glm.Vec3{X: 1, Y: 2, Z: 3}}
	rotatedBody.SetOrientationQuat(&q2)
	movedRotatedBody.SetOrientationQuat(&q2)

	tests := []struct {
		sphere CollisionSphere
		ray    Ray
		hit    bool
		point  glm.Vec3
	}{
		{ //miss
			sphere: CollisionSphere{
				body:   &originBody,
				radius: 1,
			},
			ray: NewRayFromTo(glm.Vec3{X: -2, Y: -2, Z: 0}, glm.Vec3{X: 2, Y: -2, Z: 0}),
			hit: false,
		},
		{ //hit axis aligned
			sphere: CollisionSphere{
				body:   &originBody,
				radius: 1,
			},
			ray:   NewRayFromTo(glm.Vec3{X: -2, Y: 0, Z: 0}, glm.Vec3{X: 2, Y: 0, Z: 0}),
			hit:   true,
			point: glm.Vec3{X: -1, Y: 0, Z: 0},
		},
		{ // hit diagonally
			sphere: CollisionSphere{
				body:   &originBody,
				radius: 1,
			},
			ray:   NewRayFromTo(glm.Vec3{X: -2, Y: -2, Z: -2}, glm.Vec3{X: 2, Y: 2, Z: 2}),
			hit:   true,
			point: glm.Vec3{X: -0.5773504, Y: -0.5773504, Z: -0.5773504},
		},
		{ // miss hit something at {5, 5, 5}
			sphere: CollisionSphere{
				body:   &movedBody,
				radius: 1,
			},
			ray:   NewRayFromTo(glm.Vec3{X: -2, Y: 0, Z: 0}, glm.Vec3{X: 2, Y: 0, Z: 0}),
			hit:   false,
			point: glm.Vec3{X: -1, Y: 0, Z: 0},
		},
		{ // hit something at {5, 5, 5}
			sphere: CollisionSphere{
				body:   &movedBody,
				radius: 1,
			},
			ray:   NewRayFromTo(glm.Vec3{X: -2 + 5, Y: 0 + 5, Z: 0 + 5}, glm.Vec3{X: 2 + 5, Y: 0 + 5, Z: 0 + 5}),
			hit:   true,
			point: glm.Vec3{X: -1 + 5, Y: 0 + 5, Z: 0 + 5},
		},
		{ //hit something diagonally at {5, 5, 5}
			sphere: CollisionSphere{
				body:   &movedBody,
				radius: 1,
			},
			ray:   NewRayFromTo(glm.Vec3{X: -2 + 5, Y: -2 + 5, Z: -2 + 5}, glm.Vec3{X: 2 + 5, Y: 2 + 5, Z: 2 + 5}),
			hit:   true,
			point: glm.Vec3{X: -0.5773504 + 5, Y: -0.5773504 + 5, Z: -0.5773504 + 5},
		},
	}

	for i, test := range tests {
		var res RayResultAny
		test.sphere.RayTest(test.ray, &res)
		if !test.hit {
			if res.Body != nil {
				t.Errorf("[%d] unexpected hit", i)
			}
			continue
		}
		if test.hit && res.Body == nil {
			t.Errorf("[%d] expected hit got nothing.", i)
			continue
		}

		if test.point != res.Hit {
			t.Errorf("[%d] hit = %v, want %v", i, res.Hit, test.point)
		}
	}
}

func TestCollisionBox_RayTest(t *testing.T) {
	var originBody, movedBody, rotatedBody, movedRotatedBody RigidBody
	movedBody.SetPosition3f(5, 5, 5)
	movedRotatedBody.SetPosition3f(5, 5, 5)
	qi := glm.QuatIdent()
	originBody.SetOrientationQuat(&qi)
	movedBody.SetOrientationQuat(&qi)
	q3 := glm.QuatRotate(math.Pi/4, &glm.Vec3{X: 0, Y: 1, Z: 0})
	rotatedBody.SetOrientationQuat(&q3)
	movedRotatedBody.SetOrientationQuat(&q3)

	originBody.calculateDerivedData()
	movedBody.calculateDerivedData()
	rotatedBody.calculateDerivedData()
	movedRotatedBody.calculateDerivedData()

	tests := []struct {
		box   CollisionBox
		ray   Ray
		hit   bool
		point glm.Vec3
	}{
		{ // full on hit
			box: CollisionBox{
				body:     &originBody,
				halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			ray:   NewRayFromTo(glm.Vec3{X: -3, Y: 0, Z: 0}, glm.Vec3{X: 3, Y: 0, Z: 0}),
			hit:   true,
			point: glm.Vec3{X: -1, Y: 0, Z: 0},
		},
		{ // completelly miss
			box: CollisionBox{
				body:     &originBody,
				halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			ray: NewRayFromTo(glm.Vec3{X: -3, Y: 5, Z: 0}, glm.Vec3{X: 3, Y: 5, Z: 0}),
			hit: false,
		},
		{ // full on hit, moved box
			box: CollisionBox{
				body:     &movedBody,
				halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			ray:   NewRayFromTo(glm.Vec3{X: 2, Y: 5, Z: 5}, glm.Vec3{X: 8, Y: 5, Z: 5}),
			hit:   true,
			point: glm.Vec3{X: 4, Y: 5, Z: 5},
		},
		{ // box rotates by 45 deg
			box: CollisionBox{
				body:     &rotatedBody,
				halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			ray:   NewRayFromTo(glm.Vec3{X: -3, Y: 0, Z: 0}, glm.Vec3{X: 3, Y: 0, Z: 0}),
			hit:   true,
			point: glm.Vec3{X: -1.4142135, Y: 0, Z: 0},
		},
		{ // box moved and rotated by pi/4
			box: CollisionBox{
				body:     &movedRotatedBody,
				halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			ray:   NewRayFromTo(glm.Vec3{X: 2, Y: 5, Z: 5}, glm.Vec3{X: 8, Y: 5, Z: 5}),
			hit:   true,
			point: glm.Vec3{X: -0.4142135 + 4, Y: 5, Z: 5},
		},
		{ // hit only a face
			box: CollisionBox{
				body:     &originBody,
				halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			ray:   NewRayFromTo(glm.Vec3{X: -3, Y: 1, Z: 0}, glm.Vec3{X: 3, Y: 1, Z: 0}),
			hit:   true,
			point: glm.Vec3{X: -1, Y: 1, Z: 0},
		},
		{ // full on hit reverse
			box: CollisionBox{
				body:     &originBody,
				halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			ray:   NewRayFromTo(glm.Vec3{X: 3, Y: 0, Z: 0}, glm.Vec3{X: -3, Y: 0, Z: 0}),
			hit:   true,
			point: glm.Vec3{X: 1, Y: 0, Z: 0},
		},
		{ // box moved and rotated by pi/4, reverse direction
			box: CollisionBox{
				body:     &movedRotatedBody,
				halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
			},
			ray:   NewRayFromTo(glm.Vec3{X: 3 + 5, Y: 0 + 5, Z: 0 + 5}, glm.Vec3{X: -3 + 5, Y: 0 + 5, Z: 0 + 5}),
			hit:   true,
			point: glm.Vec3{X: 6.4142135, Y: 5, Z: 5},
		},
	}

	for i, test := range tests {
		var res RayResultAny
		test.box.RayTest(test.ray, &res)
		if !test.hit {
			if res.Body != nil {
				t.Errorf("[%d] unexpected hit", i)
			}
			continue
		}
		if test.hit && res.Body == nil {
			t.Errorf("[%d] expected hit", i)
			continue
		}
		if !test.point.EqualThreshold(&res.Hit, 1e-2) {
			t.Errorf("[%d] hitpoint = %v, want %v", i, res.Hit, test.point)
		}
	}
}

func BenchmarkCollisionSphere_RayTest(b *testing.B) {
	var movedBody RigidBody
	movedBody.SetPosition3f(5, 5, 5)
	qi := glm.QuatIdent()
	movedBody.SetOrientationQuat(&qi)

	sphere := CollisionSphere{
		body:   &movedBody,
		radius: 1,
	}
	ray := NewRayFromTo(glm.Vec3{X: -2 + 5, Y: -2 + 5, Z: -2 + 5}, glm.Vec3{X: 2 + 5, Y: 2 + 5, Z: 2 + 5})
	var rr RayResult
	rr = RayResultNothing{}
	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		sphere.RayTest(ray, rr)
	}
}

func BenchmarkCollisionBox_RayTest(b *testing.B) {
	var movedRotatedBody RigidBody
	movedRotatedBody.SetPosition3f(5, 5, 5)
	q3 := glm.QuatRotate(math.Pi/4, &glm.Vec3{X: 0, Y: 1, Z: 0})
	movedRotatedBody.SetOrientationQuat(&q3)

	movedRotatedBody.calculateDerivedData()

	box := CollisionBox{
		body:     &movedRotatedBody,
		halfSize: glm.Vec3{X: 1, Y: 1, Z: 1},
	}
	ray := NewRayFromTo(glm.Vec3{X: -3 + 5, Y: 0 + 5, Z: 0 + 5}, glm.Vec3{X: 3 + 5, Y: 0 + 5, Z: 0 + 5})

	var rr RayResult
	rr = RayResultNothing{}
	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		box.RayTest(ray, rr)
	}
}
