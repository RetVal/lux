package geo

import (
	"fmt"
)

// Shape just typedefs interface{} to differentiate between seemingly any data
// from collision shapes. Valid Shapes are {*Sphere, *AABB, *OBB, *Capsule}.
type Shape interface{}

// TestShapeShape takes 2 shape, their respective transforms and returns wether
// or not they collide.
func TestShapeShape(s0, s1 Shape) bool {
	switch c0 := s0.(type) {
	case *Sphere:
		switch c1 := s1.(type) {
		case *Sphere:
			return TestSphereSphere(c0, c1)
		case *AABB:
			return TestAABBSphere(c1, c0)
		case *OBB:
			return TestOBBSphere(c1, c0)
		case *Capsule:
			return TestCapsuleSphere(c1, c0)
		case *Convexhull:
			return TestConvexhullSphere(c1, c0)
		}
	case *AABB:
		switch c1 := s1.(type) {
		case *Sphere:
			return TestAABBSphere(c0, c1)
		case *AABB:
			return TestAABBAABB(c0, c1)
		case *OBB:
			return TestAABBOBB(c0, c1)
		case *Capsule:
			return TestAABBCapsule(c0, c1)
		case *Convexhull:
			return TestAABBConvexhull(c0, c1)
		}
	case *OBB:
		switch c1 := s1.(type) {
		case *Sphere:
			return TestOBBSphere(c0, c1)
		case *AABB:
			return TestAABBOBB(c1, c0)
		case *OBB:
			return TestOBBOBB(c0, c1)
		case *Capsule:
			return TestCapsuleOBB(c1, c0)
		case *Convexhull:
			break
			// return TestConvexhullOBB(c1,c0)
		}
	case *Capsule:
		switch c1 := s1.(type) {
		case *Sphere:
			return TestCapsuleSphere(c0, c1)
		case *AABB:
			return TestAABBCapsule(c1, c0)
		case *OBB:
			return TestCapsuleOBB(c0, c1)
		case *Capsule:
			return TestCapsuleCapsule(c0, c1)
		case *Convexhull:
			break
			// return TestCapsuleConvexhull(c0,c1)
		}

	case *Convexhull:
		switch c1 := s1.(type) {
		case *Sphere:
			return TestConvexhullSphere(c0, c1)
		case *AABB:
			return TestAABBConvexhull(c1, c0)
		case *OBB:
			break
			// return TestConvexhullOBB(c0,c1)
		case *Capsule:
			break
			// return TestCapsuleConvexhull(c1, c0)
		case *Convexhull:
			return TestConvexhullConvexhull(c0, c1)
		}
	}
	panic(fmt.Sprintf("Unsupported collision: %T, %T", s0, s1))
}
