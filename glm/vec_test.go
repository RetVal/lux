package glm

import (
	"github.com/luxengine/lux/flops"
	"testing"
)

func Test2DVecAdd(t *testing.T) {
	t.Parallel()
	v1 := Vec2{1.0, 2.5}
	v2 := Vec2{0.0, 1.0}

	v3 := v1.Add(&v2)

	if !flops.Eq(v3.X, 1.0) || !flops.Eq(v3.Y, 3.5) {
		t.Errorf("Add not adding properly")
	}

	v4 := v2.Add(&v1)

	if !flops.Eq(v3.X, v4.X) || !flops.Eq(v3.Y, v4.Y) {
		t.Errorf("Addition is somehow not commutative")
	}

}

func Test3DVecAdd(t *testing.T) {
	t.Parallel()
	v1 := Vec3{1.0, 2.5, 1.1}
	v2 := Vec3{0.0, 1.0, 9.9}

	v3 := v1.Add(&v2)

	if !flops.Eq(v3.X, 1.0) || !flops.Eq(v3.Y, 3.5) || !flops.Eq(v3.Z, 11.0) {
		t.Errorf("Add not adding properly")
	}

	v4 := v2.Add(&v1)

	if !flops.Eq(v3.X, v4.X) || !flops.Eq(v3.Y, v4.Y) || !flops.Eq(v3.Z, v4.Z) {
		t.Errorf("Addition is somehow not commutative")
	}

}

func Test4DVecAdd(t *testing.T) {
	t.Parallel()
	v1 := Vec4{1.0, 2.5, 1.1, 2.0}
	v2 := Vec4{0.0, 1.0, 9.9, 100.0}

	v3 := v1.Add(&v2)

	if !flops.Eq(v3.X, 1.0) || !flops.Eq(v3.Y, 3.5) || !flops.Eq(v3.Z, 11.0) || !flops.Eq(v3.W, 102.0) {
		t.Errorf("Add not adding properly")
	}

	v4 := v2.Add(&v1)

	if !flops.Eq(v3.X, v4.X) || !flops.Eq(v3.Y, v4.Y) || !flops.Eq(v3.Z, v4.Z) || !flops.Eq(v3.W, v4.W) {
		t.Errorf("Addition is somehow not commutative")
	}

}

func Test2DVecSub(t *testing.T) {
	t.Parallel()
	v1 := Vec2{1.0, 2.5}
	v2 := Vec2{0.0, 1.0}

	v3 := v1.Sub(&v2)

	if !flops.Eq(v3.X, 1.0) || !flops.Eq(v3.Y, 1.5) {
		t.Errorf("Sub not subtracting properly [%f, %f]", v3.X, v3.Y)
	}

}

func Test3DVecSub(t *testing.T) {
	t.Parallel()
	v1 := Vec3{1.0, 2.5, 1.1}
	v2 := Vec3{0.0, 1.0, 9.9}

	v3 := v1.Sub(&v2)

	// 1.1-9.9 does stupid things to floats, so we need threshold
	if !flops.Eq(v3.X, 1.0) || !flops.Eq(v3.Y, 1.5) || !FloatEqualThreshold(v3.Z, -8.8, 1e-5) {
		t.Errorf("Sub not subtracting properly [%f, %f, %f]", v3.X, v3.Y, v3.Z)
	}

}

func Test4DVecSub(t *testing.T) {
	t.Parallel()
	v1 := Vec4{1.0, 2.5, 1.1, 2.0}
	v2 := Vec4{0.0, 1.0, 9.9, 100.0}

	v3 := v1.Sub(&v2)

	// 1.1-9.9 does stupid things to floats, so we need a more tolerant threshold
	if !flops.Eq(v3.X, 1.0) || !flops.Eq(v3.Y, 1.5) || !FloatEqualThreshold(v3.Z, -8.8, 1e-5) || !flops.Eq(v3.W, -98.0) {
		t.Errorf("Sub not subtracting properly [%f, %f, %f, %f]", v3.X, v3.Y, v3.Z, v3.W)
	}

}

func TestVecMul(t *testing.T) {
	t.Parallel()
	v := Vec2{1.0, 0.0}
	v = v.Mul(15.0)

	if !flops.Eq(v.X, 15.0) || !flops.Eq(v.Y, 0.0) {
		t.Errorf("Vec mul does something weird [%f, %f]", v.X, v.Y)
	}

	v2 := Vec3{1.0, 0.0, 100.1}
	v2 = v2.Mul(15.0)

	if !flops.Eq(v2.X, 15.0) || !flops.Eq(v2.Y, 0.0) || !flops.Eq(v2.Z, 1501.5) {
		t.Errorf("Vec mul does something weird [%f, %f, %f]", v2.X, v2.Y, v2.Z)
	}

	v3 := Vec4{1.0, 0.0, 100.1, -1.0}
	v3 = v3.Mul(15.0)

	if !flops.Eq(v3.X, 15.0) || !flops.Eq(v3.Y, 0.0) || !flops.Eq(v3.Z, 1501.5) || !flops.Eq(v3.W, -15.0) {
		t.Errorf("Vec mul does something weird [%f, %f, %f, %f]", v3.X, v3.Y, v3.Z, v3.W)
	}
}

func TestVecCrossProduct(t *testing.T) {
	t.Parallel()
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{10, 11, 12}
	expected := Vec3{-9, 18, -9}
	result := v1.Cross(&v2)

	if !expected.Equal(&result) {
		t.Errorf("Vec3 cross product %v x %v Got: %v. Expected: %v.",
			v1, v2, result, expected)
	}
}

func TestVecCrossOfProduct(t *testing.T) {
	t.Parallel()
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{10, 11, 12}
	expected := v1.Cross(&v2)
	var result Vec3
	result.CrossOf(&v1, &v2)

	if !expected.Equal(&result) {
		t.Errorf("Vec3 cross product %v x %v Got: %v. Expected: %v.",
			v1, v2, result, expected)
	}
}

func TestVecCrossWithProduct(t *testing.T) {
	t.Parallel()
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{10, 11, 12}
	expected := v1.Cross(&v2)
	result := v1
	result.CrossWith(&v2)

	if !expected.Equal(&result) {
		t.Errorf("Vec3 cross product %v x %v Got: %v. Expected: %v.",
			v1, v2, result, expected)
	}
}

func TestVecDotProduct(t *testing.T) {
	t.Parallel()
	mustEqual := func(result float32, expected float32, name string) {
		if !flops.Eq(result, expected) {
			t.Errorf("%v.Dot(%v) failed. Got: %v. Expected %v.",
				name, name, result, expected)
		}
	}

	mustEqual((&Vec2{1, 2}).Dot(&Vec2{3, 4}), 11, "Vec2")
	mustEqual((&Vec3{-1, -5, -7}).Dot(&Vec3{10, 20, 30}), -320, "Vec3")
	mustEqual((&Vec4{1, 3, 5, 7}).Dot(&Vec4{10, 20, 30, 40}), 500, "Vec3")
}

func TestVecLen(t *testing.T) {
	t.Parallel()
	mustEqual := func(result float32, expected float32, name string) {
		if !flops.Eq(result, expected) {
			t.Errorf("%v failed. Got: %v. Expected %v.",
				name, result, expected)
		}
	}

	v2, v3, v4 := &Vec2{3, 4}, &Vec3{2, -5, 4}, &Vec4{2, 3, 5, 7}
	if v2 == nil {
		t.Errorf("v2 is nil")
	}
	mustEqual(v2.Len(), 5, "Vec2.Len()")
	mustEqual(v3.Len(), 6.708203932499, "Vec3.Len()")
	mustEqual(v4.Len(), 9.3273790530888, "Vec4.Len()")
}

func Test2DVecNormalize(t *testing.T) {
	t.Parallel()
	v := Vec2{3, 4}
	norm := v.Normalized()
	expected := &Vec2{0.6, 0.8}
	if !norm.Equal(expected) {
		t.Errorf("%v.Normalize() failed. Got: %v. Expected: %v",
			v, norm, expected)
	}
}

func TestVecElemAccessors(t *testing.T) {
	t.Parallel()
	mustEqual := func(desc string, expected float32, results ...float32) {
		for _, r := range results {
			if !flops.Eq(expected, r) {
				t.Errorf("%v failed. Got: %v. Expected %v...",
					desc, results, expected)
			}
		}
	}

	const x, y, z, w = 1, 2, 3, 4
	v2 := Vec2{x, y}
	v3 := Vec3{x, y, z}
	v4 := Vec4{x, y, z, w}

	mustEqual("Vec.X", x, v2.X, v3.X, v4.X)
	mustEqual("Vec.Y", y, v2.Y, v3.Y, v4.Y)
	mustEqual("Vec.Z", z, v3.Z, v4.Z)
	mustEqual("Vec.W", w, v4.W)

	x2, y2 := v2.Elem()
	x3, y3, z3 := v3.Elem()
	x4, y4, z4, w4 := v4.Elem()
	mustEqual("Vec.Elem() -> x", x, x2, x3, x4)
	mustEqual("Vec.Elem() -> y", y, y2, y3, y4)
	mustEqual("Vec.Elem() -> z", z, z3, z4)
	mustEqual("Vec.Elem() -> w", w, w4)
}

func TestVecEqual(t *testing.T) {
	t.Parallel()
	assert := func(res bool, desc string) {
		if !res {
			t.Errorf("%v failed.", desc)
		}
	}

	v2, errV2 := Vec2{1, 2}, Vec2{1, 0}
	v3, errV3 := Vec3{1, 2, 3}, Vec3{1, 2, 0}
	v4, errV4 := Vec4{1, 2, 3, 4}, Vec4{1, 2, 0, 4}

	assert(v2.Equal(&v2), "Vec2.Equal")
	assert(!v2.Equal(&errV2), "Vec2.Equal")

	assert(v2.EqualThreshold(&v2, 0.1), "Vec2.EqualThreshold")
	assert(!v2.EqualThreshold(&errV2, 0.1), "Vec2.EqualThreshold")

	assert(v3.Equal(&v3), "Vec3.Equal")
	assert(!v3.Equal(&errV3), "Vec3.Equal")

	assert(v3.EqualThreshold(&v3, 0.1), "Vec3.EqualThreshold")
	assert(!v3.EqualThreshold(&errV3, 0.1), "Vec3.EqualThreshold")

	assert(v4.Equal(&v4), "Vec4.Equal")
	assert(!v4.Equal(&errV4), "Vec4.Equal")

	assert(v4.EqualThreshold(&v4, 0.1), "Vec4.EqualThreshold")
	assert(!v4.EqualThreshold(&errV4, 0.1), "Vec4.EqualThreshold")

}

func TestOuterProd2(t *testing.T) {
	t.Parallel()
	v1 := Vec2{1, -1}
	v2 := Vec2{2, 3}
	m := v1.OuterProd2(&v2)
	expected := Mat2{2, -2, 3, -3}
	if m != expected {
		t.Errorf("unexpeted result from %+v\n%+v", m, expected)
	}
}

func TestOuterProd3(t *testing.T) {
	t.Parallel()
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{3, 2, 1}
	m := v1.OuterProd3(&v2)
	expected := Mat3{3, 6, 9, 2, 4, 6, 1, 2, 3}
	if m != expected {
		t.Errorf("unexpected result from Outerprod3 %+v, %+v", m, expected)
	}
}

func TestVec4Normalize(t *testing.T) {
	t.Parallel()
	v := Vec4{1, 2, 3, 4}
	v.Normalize()
	if !flops.Eq(v.Len(), 1) {
		t.Errorf("length after normalize not 1")
	}
}

var v2 = Vec4{1, 2, 3, 4}
var v3 = Vec4{1, 2, 3, 4}
var tmpvec4 Vec4

func BenchmarkVec4Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpvec4 = v2.Add(&v3)
	}
}

func BenchmarkVec4AddOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpvec4.AddOf(&v2, &v3)
	}
}

func BenchmarkVec4AddWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpvec4.AddWith(&v2)
	}
}

func BenchmarkVec4Sub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpvec4 = v2.Sub(&v3)
	}
}

func BenchmarkVec4Scale(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpvec4 = v2.Mul(0.5)
	}
}

var tmpf float32

func BenchmarkVec4Dot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpf = v2.Dot(&v3)
	}
}

func BenchmarkVec4Len(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpf = v2.Len()
	}
}

func BenchmarkVec4Normalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpvec4 = v2.Normalized()
	}
}

var v3_1 = Vec3{1, 2, 3}
var v3_2 = Vec3{3, 2, 1}
var v3_3 Vec3

func BenchmarkVecCross(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v3_3 = v3_1.Cross(&v3_2)
	}
}
