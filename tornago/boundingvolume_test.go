package tornago

import (
	"github.com/luxengine/lux/glm"
	"testing"
)

func TestBoundingSphere_Overlaps(t *testing.T) {
	v1 := NewBoundingSphere(&glm.Vec3{2, 0, 0}, 2)
	v2 := NewBoundingSphere(&glm.Vec3{3, 0, 0}, 1)

	if !v1.Overlaps(&v2) {
		t.Errorf("spheres should overlap %v, %v", v1, v2)
	}

	v2.center = glm.Vec3{10, 0, 0}

	if v1.Overlaps(&v2) {
		t.Errorf("spheres should not overlap %v, %v", v1, v2)
	}

	v2.center = glm.Vec3{5, 0, 0}

	if v1.Overlaps(&v2) {
		t.Errorf("spheres should not overlap %v, %v", v1, v2)
	}
}

func TestBoundingSphere_New(t *testing.T) {
	pos := glm.Vec3{1, 2, 3}
	radius := float32(5)
	bs := NewBoundingSphere(&pos, radius)
	if bs.Radius() != radius {
		t.Errorf("Error settings radius, %f, want %f", bs.Radius(), radius)
	}
	if c := bs.Center(); c != pos {
		t.Errorf("Error setting center, %v, want %v", c, pos)
	}
}

func TestBoundingSphere_SphereFromSphere(t *testing.T) {
	{ //not one inside the other
		p1, p2 := glm.Vec3{0, 0, 0}, glm.Vec3{5, 0, 0}
		r1, r2 := float32(1), float32(1)
		s1, s2 := NewBoundingSphere(&p1, r1), NewBoundingSphere(&p2, r2)

		p3, r3 := glm.Vec3{2.5, 0, 0}, float32(3.5)
		s3 := NewBoundingSphereFromSpheres(&s1, &s2)

		if r := s3.Radius(); r != r3 {
			t.Errorf("Error settings radius, %f, want %f", r, r3)
		}
		if c := s3.Center(); c != p3 {
			t.Errorf("Error setting center, %v, want %v", c, p3)
		}
	}
	{ //s1 inside s2
		p1, p2 := glm.Vec3{0, 0, 0}, glm.Vec3{5, 0, 0}
		r1, r2 := float32(1), float32(10)
		s1, s2 := NewBoundingSphere(&p1, r1), NewBoundingSphere(&p2, r2)

		p3, r3 := glm.Vec3{5, 0, 0}, float32(10)
		s3 := NewBoundingSphereFromSpheres(&s1, &s2)

		if r := s3.Radius(); r != r3 {
			t.Errorf("Error settings radius, %f, want %f", r, r3)
		}
		if c := s3.Center(); c != p3 {
			t.Errorf("Error setting center, %v, want %v", c, p3)
		}
	}
	{ //s2 inside s1
		p1, p2 := glm.Vec3{0, 0, 0}, glm.Vec3{5, 0, 0}
		r1, r2 := float32(10), float32(1)
		s1, s2 := NewBoundingSphere(&p1, r1), NewBoundingSphere(&p2, r2)

		p3, r3 := glm.Vec3{0, 0, 0}, float32(10)
		s3 := NewBoundingSphereFromSpheres(&s1, &s2)

		if r := s3.Radius(); r != r3 {
			t.Errorf("Error settings radius, %f, want %f", r, r3)
		}
		if c := s3.Center(); c != p3 {
			t.Errorf("Error setting center, %v, want %v", c, p3)
		}
	}
}

func TestBoundingSphere_GetGrowth(t *testing.T) {
	p1, p2 := glm.Vec3{0, 0, 0}, glm.Vec3{5, 0, 0}
	r1, r2 := float32(1), float32(1)
	s1, s2 := NewBoundingSphere(&p1, r1), NewBoundingSphere(&p2, r2)

	s3 := NewBoundingSphereFromSpheres(&s1, &s2)

	if size, g := s3.Size()-s1.Size(), s1.Growth(&s2); size != g {
		t.Errorf("GetGrowth = %f, want %f", g, size)
	}
}
