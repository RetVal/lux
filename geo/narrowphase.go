package geo

const (
	aabbShapeType = iota
	sphereShapeType
	obbShapeType
	capsuleShapeType
	convexhullShapeType
	shapeTypeLen
)

var testTable = [shapeTypeLen][shapeTypeLen]func(Shape, Shape) bool{
	{ // AABB test table
		// AABB AABB.
		func(s0, s1 Shape) bool {
			aabb0, aabb1 := s0.(*AABB), s1.(*AABB)
			return TestAABBAABB(aabb0, aabb1)
		},
		// AABB Sphere
		func(s0, s1 Shape) bool {
			aabb, sphere := s0.(*AABB), s1.(*Sphere)
			return TestAABBSphere(aabb, sphere)
		},
		// AABB OBB
		func(s0, s1 Shape) bool {
			aabb, obb := s0.(*AABB), s1.(*OBB)
			return TestAABBOBB(aabb, obb)
		},
		// AABB Capsule
		func(s0, s1 Shape) bool {
			aabb, capsule := s0.(*AABB), s1.(*Capsule)
			return TestAABBCapsule(aabb, capsule)
		},
		// AABB Convexhull
		func(s0, s1 Shape) bool {
			aabb, hull := s0.(*AABB), s1.(*Convexhull)
			return TestAABBConvexhull(aabb, hull)
		},
	},
	{ // Sphere test table
		// Sphere AABB.
		func(s0, s1 Shape) bool {
			sphere, aabb := s0.(*Sphere), s1.(*AABB)
			return TestAABBSphere(aabb, sphere)
		},
		// Sphere Sphere
		func(s0, s1 Shape) bool {
			sphere0, sphere1 := s0.(*Sphere), s1.(*Sphere)
			return TestSphereSphere(sphere0, sphere1)
		},
		// Sphere OBB
		func(s0, s1 Shape) bool {
			sphere, obb := s0.(*Sphere), s1.(*OBB)
			return TestOBBSphere(obb, sphere)
		},
		// Sphere Capsule
		func(s0, s1 Shape) bool {
			sphere, capsule := s0.(*Sphere), s1.(*Capsule)
			return TestCapsuleSphere(capsule, sphere)
		},
		// Sphere Convexhull
		func(s0, s1 Shape) bool {
			sphere, hull := s0.(*Sphere), s1.(*Convexhull)
			return TestConvexhullSphere(hull, sphere)
		},
	},
	{ // OBB test table
		// OBB AABB.
		func(s0, s1 Shape) bool {
			obb, aabb := s0.(*OBB), s1.(*AABB)
			return TestAABBOBB(aabb, obb)
		},
		// OBB Sphere
		func(s0, s1 Shape) bool {
			obb, sphere := s0.(*OBB), s1.(*Sphere)
			return TestOBBSphere(obb, sphere)
		},
		// OBB OBB
		func(s0, s1 Shape) bool {
			obb0, obb1 := s0.(*OBB), s1.(*OBB)
			return TestOBBOBB(obb0, obb1)
		},
		// OBB Capsule
		func(s0, s1 Shape) bool {
			obb, capsule := s0.(*OBB), s1.(*Capsule)
			return TestCapsuleOBB(capsule, obb)
		},
		// OBB Convexhull
		func(s0, s1 Shape) bool {
			panic("obb - convexhull test not implemented")
			//obb, hull := s0.(*OBB), s1.(*Convexhull)
			//return TestConvexhullOBB(hull, obb)
		},
	},
	{ // Capsule test table
		// Capsule AABB.
		func(s0, s1 Shape) bool {
			capsule, aabb := s0.(*Capsule), s1.(*AABB)
			return TestAABBCapsule(aabb, capsule)
		},
		// Capsule Sphere
		func(s0, s1 Shape) bool {
			capsule, sphere := s0.(*Capsule), s1.(*Sphere)
			return TestCapsuleSphere(capsule, sphere)
		},
		// Capsule OBB
		func(s0, s1 Shape) bool {
			capsule, obb := s0.(*Capsule), s1.(*OBB)
			return TestCapsuleOBB(capsule, obb)
		},
		// Capsule Capsule
		func(s0, s1 Shape) bool {
			capsule0, capsule1 := s0.(*Capsule), s1.(*Capsule)
			return TestCapsuleCapsule(capsule0, capsule1)
		},
		// Capsule Convexhull
		func(s0, s1 Shape) bool {
			panic("capsule - convexhull test not implemented")
			//capsule, hull := s0.(*Capsule), s1.(*Convexhull)
			//return TestCapsuleConvexhull(capsule, hull)
		},
	},
	{ // Convexhull test table
		// Convexhull AABB.
		func(s0, s1 Shape) bool {
			hull, aabb := s0.(*Convexhull), s1.(*AABB)
			return TestAABBConvexhull(aabb, hull)
		},
		// Convexhull Sphere
		func(s0, s1 Shape) bool {
			hull, sphere := s0.(*Convexhull), s1.(*Sphere)
			return TestConvexhullSphere(hull, sphere)
		},
		// Convexhull OBB
		func(s0, s1 Shape) bool {
			panic("convexhull - obb not implemented")
			//hull, obb := s0.(*Convexhull), s1.(*OBB)
			//return TestConvexhullOBB(hull, obb)
		},
		// Convexhull Capsule
		func(s0, s1 Shape) bool {
			panic("convexhull - capsule not implemented")
			//hull, capsule := s0.(*Convexhull), s1.(*Capsule)
			//return TestCapsuleConvexhull(capsule, hull)
		},
		// Convexhull Convexhull
		func(s0, s1 Shape) bool {
			hull0, hull1 := s0.(*Convexhull), s1.(*Convexhull)
			return TestConvexhullConvexhull(hull0, hull1)
		},
	},
}

// Shape just typedefs interface{} to differentiate between seemingly any data
// from collision shapes. Valid Shapes are {*Sphere, *AABB, *OBB, *Capsule}.
type Shape interface {
	ShapeType() int
}

// TestShapeShape takes 2 shape, their respective transforms and returns wether
// or not they collide.
func TestShapeShape(s0, s1 Shape) bool {
	return testTable[s0.ShapeType()][s1.ShapeType()](s0, s1)
}
