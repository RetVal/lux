package glm

import (
	"fmt"
	"github.com/luxengine/lux/flops"
	"github.com/luxengine/lux/math"
	"unsafe"
)

// Vec2 is the representation of a vector with 2 components.
type Vec2 struct {
	X, Y float32
}

// Vec3 is the representation of a vector with 3 components.
type Vec3 struct {
	X, Y, Z float32
}

// Vec4 is the representation of a vector with 4 components.
type Vec4 struct {
	X, Y, Z, W float32
}

// String returns a pretty string for this vector. eg.
// {-1.00000, 0.00000}
func (v1 *Vec2) String() string {
	ret := "{"
	//for n := 0; n < len(v1); n++ {
	elems := [...]float32{v1.X, v1.Y}
	for n, v := range elems {
		if v >= 0 {
			ret += " "
		}
		ret += fmt.Sprintf("%.6f", v)
		if n < len(elems)-1 {
			ret += ", "
		}
	}
	return ret + "}"
}

// String returns a pretty string for this vector. eg.
// {-1.00000, 0.00000, 0.00000}
func (v1 *Vec3) String() string {
	ret := "{"
	//for n := 0; n < len(v1); n++ {
	elems := [...]float32{v1.X, v1.Y, v1.Z}
	for n, v := range elems {
		if v >= 0 {
			ret += " "
		}
		ret += fmt.Sprintf("%.6f", v)
		if n < len(elems)-1 {
			ret += ", "
		}
	}
	return ret + "}"
}

// String returns a pretty string for this vector. eg.
// {-1.00000, 0.00000, 0.00000, 0.00000}
func (v1 *Vec4) String() string {
	ret := "{"
	//for n := 0; n < len(v1); n++ {
	elems := [...]float32{v1.X, v1.Y, v1.Z, v1.W}
	for n, v := range elems {
		if v >= 0 {
			ret += " "
		}
		ret += fmt.Sprintf("%.6f", v)
		if n < len(elems)-1 {
			ret += ", "
		}
	}
	return ret + "}"
}

// Vec3 return a Vec3 from this Vec2 with {z}. Similar to GLSL
//    vec3(v2, z);
func (v1 *Vec2) Vec3(z float32) Vec3 {
	return Vec3{v1.X, v1.Y, z}
}

// Vec4 return a Vec4 from this Vec2 with {z,w}. Similar to GLSL
//    vec4(v2, z, w);
func (v1 *Vec2) Vec4(z, w float32) Vec4 {
	return Vec4{v1.X, v1.Y, z, w}
}

// Vec2 return a Vec2 from the first 2 components of this Vec3. Similar to GLSL
//    vec2(v3);
func (v1 *Vec3) Vec2() Vec2 {
	return Vec2{v1.X, v1.Y}
}

// Vec4 return a Vec4 from this Vec3 with {w}. Similar to GLSL
//    vec4(v3, w);
func (v1 *Vec3) Vec4(w float32) Vec4 {
	return Vec4{v1.X, v1.Y, v1.Z, w}
}

// Vec2 return a Vec2 from the first 2 components of this Vec4. Similar to GLSL
//    vec2(v4);
func (v1 *Vec4) Vec2() Vec2 {
	return Vec2{v1.X, v1.Y}
}

// Vec3 return a Vec3 from the first 3 components of this Vec4. Similar to GLSL
//    vec3(v4);
func (v1 *Vec4) Vec3() Vec3 {
	return Vec3{v1.X, v1.Y, v1.Z}
}

// Elem extracts the elements of the vector for direct value assignment.
func (v1 Vec2) Elem() (x, y float32) {
	return v1.X, v1.Y
}

// Elem extracts the elements of the vector for direct value assignment.
func (v1 Vec3) Elem() (x, y, z float32) {
	return v1.X, v1.Y, v1.Z
}

// Elem extracts the elements of the vector for direct value assignment.
func (v1 Vec4) Elem() (x, y, z, w float32) {
	return v1.X, v1.Y, v1.Z, v1.W
}

// Perp returns the vector perpendicular to v1
func (v1 *Vec2) Perp() Vec2 {
	return Vec2{-v1.Y, v1.X}
}

// SetPerp sets this vector to its perpendicular
func (v1 *Vec2) SetPerp() {
	v1.X, v1.Y = -v1.Y, v1.X
}

// Cross computes the pseudo 2D cross product, Dot(Perp(u), v)
func (v1 *Vec2) Cross(v2 *Vec2) float32 {
	return v1.X*v2.Y - v1.Y*v2.X
}

// Cross is an operation only defined on 3D vectors, commonly referred to as
// "the cross product". It is equivalent to
// Vec3{v1.Y*v2.Z-v1.Z*v2.Y, v1.Z*v2.X-v1.X*v2.Z, v1.X*v2.Y - v1.Y*v2.X}.
// Another interpretation is that it's the vector whose magnitude is
// |v1||v2|sin(theta) where theta is the angle between v1 and v2.
//
// The cross product is most often used for finding surface normals. The cross
// product of vectors will generate a vector that is perpendicular to the plane
// they form.
//
// Technically, a generalized cross product exists as an "(N-1)ary" operation
// (that is, the 4D cross product requires 3 4D vectors). But the binary
// 3D (and 7D) cross product is the most important. It can be considered
// the area of a parallelogram with sides v1 and v2.
//
// Like the dot product, the cross product is roughly a measure of
// directionality. Two normalized perpendicular vectors will return a vector
// with a magnitude of 1.0 or -1.0 and two parallel vectors will return a vector
// with magnitude 0.0. The cross product is "anticommutative" meaning
// v1.Cross(v2) = -v2.Cross(v1), this property can be useful to know when
// finding normals, as taking the wrong cross product can lead to the opposite
// normal of the one you want.
//
// https://en.wikipedia.org/wiki/Cross_product
func (v1 *Vec3) Cross(v2 *Vec3) Vec3 {
	return Vec3{v1.Y*v2.Z - v1.Z*v2.Y, v1.Z*v2.X - v1.X*v2.Z, v1.X*v2.Y - v1.Y*v2.X}
}

// CrossOf is the same as Cross but with destination vector. v1 = v2 X v3.
func (v1 *Vec3) CrossOf(v2, v3 *Vec3) {
	v1.X = v2.Y*v3.Z - v2.Z*v3.Y
	v1.Y = v2.Z*v3.X - v2.X*v3.Z
	v1.Z = v2.X*v3.Y - v2.Y*v3.X
}

// CrossWith is the same as cross except it stores the result in v1.
func (v1 *Vec3) CrossWith(v2 *Vec3) {
	vx, vy, vz := v1.X, v1.Y, v1.Z
	v1.X = vy*v2.Z - vz*v2.Y
	v1.Y = vz*v2.X - vx*v2.Z
	v1.Z = vx*v2.Y - vy*v2.X
}

// ScalarTripleProduct returns Dot(v1, Cross(v2,v3)), its also called the box or
// mixed product.
//
// https://en.wikipedia.org/wiki/Triple_product
func ScalarTripleProduct(v0, v1, v2 *Vec3) float32 {
	return v0.X*(v1.Y*v2.Z-v1.Z*v2.Y) +
		v0.Y*(v1.Z*v2.X-v1.X*v2.Z) +
		v0.Z*(v1.X*v2.Y-v1.Y*v2.X)
}

// Add is equivalent to v3 := v1+v2
func (v1 *Vec2) Add(v2 *Vec2) Vec2 {
	return Vec2{v1.X + v2.X, v1.Y + v2.Y}
}

// AddOf is equivalent to v1 = v2+v3
func (v1 *Vec2) AddOf(v2, v3 *Vec2) {
	v1.X, v1.Y = v2.X+v3.X, v2.Y+v3.Y
}

// AddWith is equivalent to v1+=v2
func (v1 *Vec2) AddWith(v2 *Vec2) {
	v1.X += v2.X
	v1.Y += v2.Y
}

// AddScaledVec is a shortcut for v1 += c*v2
func (v1 *Vec2) AddScaledVec(c float32, v2 *Vec2) {
	v1.X += c * v2.X
	v1.Y += c * v2.Y
}

// Sub is equivalent to v3 := v1-v2
func (v1 *Vec2) Sub(v2 *Vec2) Vec2 {
	return Vec2{v1.X - v2.X, v1.Y - v2.Y}
}

// SubOf is equivalent to v1 = v2-v3
func (v1 *Vec2) SubOf(v2, v3 *Vec2) {
	v1.X, v1.Y = v2.X-v3.X, v2.Y-v3.Y
}

// SubWith is equivalent to v1-=v2
func (v1 *Vec2) SubWith(v2 *Vec2) {
	v1.X -= v2.X
	v1.Y -= v2.Y
}

// Mul is equivalent to v3 := c*v1
func (v1 *Vec2) Mul(c float32) Vec2 {
	return Vec2{v1.X * c, v1.Y * c}
}

// MulOf is equivalent to v1 = c*v2
func (v1 *Vec2) MulOf(c float32, v2 *Vec2) {
	v1.X, v1.Y = c*v2.X, c*v2.Y
}

// MulWith is equivalent to v1*=c
func (v1 *Vec2) MulWith(c float32) {
	v1.X *= c
	v1.Y *= c
}

// ComponentProduct returns {v1.X*v2.X,v1.Y*v2.Y, ... v1[n]*v2[n]}. It's
// equivalent to v3 := v1 * v2
func (v1 *Vec2) ComponentProduct(v2 *Vec2) Vec2 {
	return Vec2{v1.X * v2.X, v1.Y * v2.Y}
}

// ComponentProductOf is equivalent to v1 = v2*v3
func (v1 *Vec2) ComponentProductOf(v2, v3 *Vec2) {
	v1.X = v2.X * v3.X
	v1.Y = v2.Y * v3.Y
}

// ComponentProductWith is equivalent to v1 = v1*v2
func (v1 *Vec2) ComponentProductWith(v2 *Vec2) {
	v1.X = v1.X * v2.X
	v1.Y = v1.Y * v2.Y
}

// Dot returns the dot product of this vector with another. There are multiple
// ways to describe this value. One is the multiplication of their lengths and
// cos(theta) where theta is the angle between the vectors:
// v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise
// multiplication of all elements. So for instance, two Vec3s would yield
// v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z.
//
// This means that the dot product of a vector and itself is the square of its
// Len (within the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to
// pointing in the same direction. If both vectors are normalized, the value
// will be -1 for opposite pointing, one for same pointing, and 0 for
// perpendicular vectors.
func (v1 *Vec2) Dot(v2 *Vec2) float32 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the
// square root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v.X, v.Y).
func (v1 *Vec2) Len() float32 {
	return math.Hypot(v1.X, v1.Y)
}

// Len2 returns the square of the length, this function is used when optimising
// out the sqrt operation.
func (v1 *Vec2) Len2() float32 {
	return v1.X*v1.X + v1.Y*v1.Y
}

// Invert changes the sign of every component of this vector.
func (v1 *Vec2) Invert() {
	v1.X = -v1.X
	v1.Y = -v1.Y
}

// Inverse return a new vector with invert sign for every component
func (v1 *Vec2) Inverse() Vec2 {
	return Vec2{-v1.X, -v1.Y}
}

// Zero sets this vector to all zero components.
func (v1 *Vec2) Zero() {
	v1.X, v1.Y = 0, 0
}

// Normalized normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating
// point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 *Vec2) Normalized() Vec2 {
	l := 1.0 / v1.Len()
	return Vec2{v1.X * l, v1.Y * l}
}

// Normalize is the same as Normalize but doesn't return a new vector.
func (v1 *Vec2) Normalize() {
	l := 1.0 / v1.Len()
	v1.X *= l
	v1.Y *= l
}

// NormalizeVec2 normalizes given vector. shortcut for when you don't want to
// use pointers.
func NormalizeVec2(v Vec2) Vec2 {
	l := 1.0 / v.Len()
	v.X *= l
	v.Y *= l
	return v
}

// Equal takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 *Vec2) Equal(v2 *Vec2) bool {
	return flops.Eq(v1.X, v2.X) && flops.Eq(v1.Y, v2.Y)
}

// EqualThreshold takes in a threshold for comparing two floats, and uses
// it to do an element-wise comparison of the vector to another.
func (v1 *Vec2) EqualThreshold(v2 *Vec2, threshold float32) bool {
	return FloatEqualThreshold(v1.X, v2.X, threshold) && FloatEqualThreshold(v1.Y, v2.Y, threshold)
}

// I gets the ith element of this vector.
func (v1 *Vec2) I(i int) *float32 {
	return (*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(v1)) + uintptr(i)*4))
}

/*
// X is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec2) X() float32 {
	return v1.X
}

// Y is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec2) Y() float32 {
	return v1.Y
}*/

// OuterProd2 does the vector outer product
// of two vectors. The outer product produces an
// 2x2 matrix. E.G. a Vec2 * Vec2 = Mat2.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 *Vec2) OuterProd2(v2 *Vec2) Mat2 {
	return Mat2{v1.X * v2.X, v1.Y * v2.X, v1.X * v2.Y, v1.Y * v2.Y}
}

// Add is equivalent to v3 := v1+v2
func (v1 *Vec3) Add(v2 *Vec3) Vec3 {
	return Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// AddOf is equivalent to v1 = v2+v3
func (v1 *Vec3) AddOf(v2, v3 *Vec3) {
	v1.X = v2.X + v3.X
	v1.Y = v2.Y + v3.Y
	v1.Z = v2.Z + v3.Z

}

// AddWith is equivalent to v1+=v2
func (v1 *Vec3) AddWith(v2 *Vec3) {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
}

// AddScaledVec is a shortcut for v1 += c*v2
func (v1 *Vec3) AddScaledVec(c float32, v2 *Vec3) {
	v1.X += c * v2.X
	v1.Y += c * v2.Y
	v1.Z += c * v2.Z
}

// Sub is equivalent to v3 := v1-v2
func (v1 *Vec3) Sub(v2 *Vec3) Vec3 {
	return Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

// SubOf is equivalent to v1 = v2-v3
func (v1 *Vec3) SubOf(v2, v3 *Vec3) {
	v1.X, v1.Y, v1.Z = v2.X-v3.X, v2.Y-v3.Y, v2.Z-v3.Z
}

// SubWith is equivalent to v1-=v2
func (v1 *Vec3) SubWith(v2 *Vec3) {
	v1.X -= v2.X
	v1.Y -= v2.Y
	v1.Z -= v2.Z
}

// Mul is equivalent to v3 := c*v1
func (v1 *Vec3) Mul(c float32) Vec3 {
	return Vec3{v1.X * c, v1.Y * c, v1.Z * c}
}

// MulOf is equivalent to v1 = c*v2
func (v1 *Vec3) MulOf(c float32, v2 *Vec3) {
	v1.X = c * v2.X
	v1.Y = c * v2.Y
	v1.Z = c * v2.Z
}

// MulWith is equivalent to v1*=c
func (v1 *Vec3) MulWith(c float32) {
	v1.X *= c
	v1.Y *= c
	v1.Z *= c
}

// ComponentProduct returns {v1.X*v2.X,v1.Y*v2.Y, ... v1[n]*v2[n]}. It's
// equivalent to v3 := v1 * v2
func (v1 *Vec3) ComponentProduct(v2 *Vec3) Vec3 {
	return Vec3{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

// ComponentProductOf is equivalent to v1 = v2*v3
func (v1 *Vec3) ComponentProductOf(v2, v3 *Vec3) {
	v1.X = v2.X * v3.X
	v1.Y = v2.Y * v3.Y
	v1.Z = v2.Z * v3.Z
}

// ComponentProductWith is equivalent to v1 = v1*v2
func (v1 *Vec3) ComponentProductWith(v2 *Vec3) {
	v1.X = v1.X * v2.X
	v1.Y = v1.Y * v2.Y
	v1.Z = v1.Z * v2.Z
}

// Dot returns the dot product of this vector with another. There are multiple
// ways to describe this value. One is the multiplication of their lengths and
// cos(theta) where theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise
// multiplication of all elements. So for instance, two Vec3s would yield
// v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z.
//
// This means that the dot product of a vector and itself is the square of its
// Len (within the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to
// pointing in the same direction. If both vectors are normalized, the value
// will be -1 for opposite pointing, one for same pointing, and 0 for
// perpendicular vectors.
func (v1 *Vec3) Dot(v2 *Vec3) float32 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the
// square root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v.X, v.Y).
func (v1 *Vec3) Len() float32 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
}

// Len2 returns the square of the length, this function is used when optimising
// out the sqrt operation.
func (v1 *Vec3) Len2() float32 {
	return v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z
}

// Invert changes the sign of every component of this vector.
func (v1 *Vec3) Invert() {
	v1.X = -v1.X
	v1.Y = -v1.Y
	v1.Z = -v1.Z
}

// Inverse return a new vector with invert sign for every component
func (v1 *Vec3) Inverse() Vec3 {
	return Vec3{-v1.X, -v1.Y, -v1.Z}
}

// Zero sets this vector to all zero components.
func (v1 *Vec3) Zero() {
	v1.X, v1.Y, v1.Z = 0, 0, 0
}

// Normalized normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating
// point error), while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 *Vec3) Normalized() Vec3 {
	l := 1.0 / v1.Len()
	return Vec3{v1.X * l, v1.Y * l, v1.Z * l}
}

// Normalize is the same as Normalize but doesn't return a new vector.
func (v1 *Vec3) Normalize() {
	l := 1.0 / v1.Len()
	v1.X *= l
	v1.Y *= l
	v1.Z *= l
}

// NormalizeVec3 normalizes given vector. shortcut for when you don't want to
// use pointers.
func NormalizeVec3(v Vec3) Vec3 {
	l := 1.0 / v.Len()
	v.X *= l
	v.Y *= l
	v.Z *= l
	return v
}

// Equal takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 *Vec3) Equal(v2 *Vec3) bool {
	return flops.Eq(v1.X, v2.X) && flops.Eq(v1.Y, v2.Y) && flops.Eq(v1.Z, v2.Z)
}

// EqualThreshold takes in a threshold for comparing two floats, and uses
// it to do an element-wise comparison of the vector to another.
func (v1 *Vec3) EqualThreshold(v2 *Vec3, threshold float32) bool {
	return FloatEqualThreshold(v1.X, v2.X, threshold) && FloatEqualThreshold(v1.Y, v2.Y, threshold) && FloatEqualThreshold(v1.Z, v2.Z, threshold)
}

// I gets the ith element of this vector.
func (v1 *Vec3) I(i int) *float32 {
	return (*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(v1)) + uintptr(i)*4))
}

/*
// X is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec3) X() float32 {
	return v1.X
}

// Y is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec3) Y() float32 {
	return v1.Y
}

// Z is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec3) Z() float32 {
	return v1.Z
}*/

// OuterProd3 does the vector outer product
// of two vectors. The outer product produces an
// 3x3 matrix. E.G. a Vec3 * Vec3 = Mat3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 *Vec3) OuterProd3(v2 *Vec3) Mat3 {
	return Mat3{v1.X * v2.X, v1.Y * v2.X, v1.Z * v2.X, v1.X * v2.Y, v1.Y * v2.Y, v1.Z * v2.Y, v1.X * v2.Z, v1.Y * v2.Z, v1.Z * v2.Z}
}

// Add is equivalent to v3 := v1+v2
func (v1 *Vec4) Add(v2 *Vec4) Vec4 {
	return Vec4{
		v1.X + v2.X,
		v1.Y + v2.Y,
		v1.Z + v2.Z,
		v1.W + v2.W,
	}
}

// AddOf is equivalent to v1 = v2+v3
func (v1 *Vec4) AddOf(v2, v3 *Vec4) {
	v1.X = v2.X + v3.X
	v1.Y = v2.Y + v3.Y
	v1.Z = v2.Z + v3.Z
	v1.W = v2.W + v3.W
}

// AddWith is equivalent to v1+=v2
func (v1 *Vec4) AddWith(v2 *Vec4) {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
	v1.W += v2.W
}

// AddScaledVec is a shortcut for v1 += c*v2
func (v1 *Vec4) AddScaledVec(c float32, v2 *Vec4) {
	v1.X += c * v2.X
	v1.Y += c * v2.Y
	v1.Z += c * v2.Z
	v1.W += c * v2.W
}

// Sub is equivalent to v3 := v1-v2
func (v1 *Vec4) Sub(v2 *Vec4) Vec4 {
	return Vec4{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z, v1.W - v2.W}
}

// SubOf is equivalent to v1 = v2-v3
func (v1 *Vec4) SubOf(v2, v3 *Vec4) {
	v1.X, v1.Y, v1.Z, v1.W = v2.X-v3.X, v2.Y-v3.Y, v2.Z-v3.Z, v2.W-v3.W
}

// SubWith is equivalent to v1-=v2
func (v1 *Vec4) SubWith(v2 *Vec4) {
	v1.X -= v2.X
	v1.Y -= v2.Y
	v1.Z -= v2.Z
	v1.W -= v2.W
}

// Mul is equivalent to v3 := c*v1
func (v1 *Vec4) Mul(c float32) Vec4 {
	return Vec4{v1.X * c, v1.Y * c, v1.Z * c, v1.W * c}
}

// MulOf is equivalent to v1 = c*v2
func (v1 *Vec4) MulOf(c float32, v2 *Vec4) {
	v1.X, v1.Y, v1.Z, v1.W = c*v2.X, c*v2.Y, c*v2.Z, c*v2.W
}

// MulWith is equivalent to v1*=c
func (v1 *Vec4) MulWith(c float32) {
	v1.X *= c
	v1.Y *= c
	v1.Z *= c
	v1.W *= c
}

// ComponentProduct returns {v1.X*v2.X,v1.Y*v2.Y, ... v1[n]*v2[n]}. It's
// equivalent to v3 := v1 * v2
func (v1 *Vec4) ComponentProduct(v2 *Vec4) Vec4 {
	return Vec4{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z, v1.W * v2.W}
}

// ComponentProductOf is equivalent to v1 = v2*v3
func (v1 *Vec4) ComponentProductOf(v2, v3 *Vec4) {
	v1.X = v2.X * v3.X
	v1.Y = v2.Y * v3.Y
	v1.Z = v2.Z * v3.Z
	v1.W = v2.W * v3.W
}

// ComponentProductWith is equivalent to v1 = v1*v2
func (v1 *Vec4) ComponentProductWith(v2 *Vec4) {
	v1.X = v1.X * v2.X
	v1.Y = v1.Y * v2.Y
	v1.Z = v1.Z * v2.Z
	v1.W = v1.W * v2.W
}

// Dot returns the dot product of this vector with another. There are multiple
// ways to describe this value. One is the multiplication of their lengths and
// cos(theta) where theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise
// multiplication of all elements. So for instance, two Vec3s would yield
// v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z.
//
// This means that the dot product of a vector and itself is the square of its
// Len (within the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to
// pointing in the same direction. If both vectors are normalized, the value
// will be -1 for opposite pointing, one for same pointing, and 0 for
// perpendicular vectors.
func (v1 *Vec4) Dot(v2 *Vec4) float32 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z + v1.W*v2.W
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the
// square root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v.X, v.Y).
func (v1 *Vec4) Len() float32 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z + v1.W*v1.W)
}

// Len2 returns the square of the length, this function is used when optimising
// out the sqrt operation.
func (v1 *Vec4) Len2() float32 {
	return v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z + v1.W*v1.W
}

// Invert changes the sign of every component of this vector.
func (v1 *Vec4) Invert() {
	v1.X = -v1.X
	v1.Y = -v1.Y
	v1.Z = -v1.Z
	v1.W = -v1.W
}

// Inverse return a new vector with invert sign for every component
func (v1 *Vec4) Inverse() Vec4 {
	return Vec4{-v1.X, -v1.Y, -v1.Z, -v1.W}
}

// Zero sets this vector to all zero components.
func (v1 *Vec4) Zero() {
	v1.X, v1.Y, v1.Z, v1.W = 0, 0, 0, 0
}

// Normalized normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating
// point error), while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 *Vec4) Normalized() Vec4 {
	l := 1.0 / v1.Len()
	return Vec4{v1.X * l, v1.Y * l, v1.Z * l, v1.W * l}
}

// Normalize is the same as Normalize but doesn't return a new vector.
func (v1 *Vec4) Normalize() {
	l := 1.0 / v1.Len()
	v1.X *= l
	v1.Y *= l
	v1.Z *= l
	v1.W *= l
}

// NormalizeVec4 normalizes given vector. shortcut for when you don't want to
// use pointers.
func NormalizeVec4(v Vec4) Vec4 {
	l := 1.0 / v.Len()
	v.X *= l
	v.Y *= l
	v.Z *= l
	v.W *= l
	return v
}

// Equal takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 *Vec4) Equal(v2 *Vec4) bool {
	return flops.Eq(v1.X, v2.X) && flops.Eq(v1.Y, v2.Y) && flops.Eq(v1.Z, v2.Z) && flops.Eq(v1.W, v2.W)
}

// EqualThreshold takes in a threshold for comparing two floats, and uses
// it to do an element-wise comparison of the vector to another.
func (v1 *Vec4) EqualThreshold(v2 *Vec4, threshold float32) bool {
	return FloatEqualThreshold(v1.X, v2.X, threshold) && FloatEqualThreshold(v1.Y, v2.Y, threshold) && FloatEqualThreshold(v1.Z, v2.Z, threshold) && FloatEqualThreshold(v1.W, v2.W, threshold)
}

// I gets the ith element of this vector.
func (v1 *Vec4) I(i int) *float32 {
	return (*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(v1)) + uintptr(i)*4))
}

/*
// X is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec4) X() float32 {
	return v1.X
}

// Y is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec4) Y() float32 {
	return v1.Y
}

// Z is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec4) Z() float32 {
	return v1.Z
}

// W is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v.X or v.X() depending on personal preference.
func (v1 Vec4) W() float32 {
	return v1.W
}*/

// SetNormalizeOf sets this vector as v2 normalized. v1 = normalize(v2).
func (v1 *Vec2) SetNormalizeOf(v2 *Vec2) {
	l := 1.0 / v2.Len()
	v1.X = l * v2.X
	v1.Y = l * v2.Y
}

// SetNormalizeOf sets this vector as v2 normalized. v1 = normalize(v2).
func (v1 *Vec3) SetNormalizeOf(v2 *Vec3) {
	l := 1.0 / v2.Len()
	v1.X = l * v2.X
	v1.Y = l * v2.Y
	v1.Z = l * v2.Z
}

// SetNormalizeOf sets this vector as v2 normalized. v1 = normalize(v2).
func (v1 *Vec4) SetNormalizeOf(v2 *Vec4) {
	l := 1.0 / v2.Len()
	v1.X = l * v2.X
	v1.Y = l * v2.Y
	v1.Z = l * v2.Z
	v1.W = l * v2.W
}

// Dotf is the same as Dot but takes 2 float32 as input instead (API convinience
// function)
func (v1 *Vec2) Dotf(x, y float32) float32 {
	return v1.X*x + v1.Y*y
}

// Dotf is the same as Dot but takes 3 float32 as input instead (API convinience
// function)
func (v1 *Vec3) Dotf(x, y, z float32) float32 {
	return v1.X*x + v1.Y*y + v1.Z*z
}

// Dotf is the same as Dot but takes 4 float32 as input instead (API convinience
// function)
func (v1 *Vec4) Dotf(x, y, z, w float32) float32 {
	return v1.X*x + v1.Y*y + v1.Z*z + v1.W*w
}

// AngleBetween returns the angle between v1 and v2 in radian.
func (v1 *Vec2) AngleBetween(v2 *Vec2) float32 {
	return math.Acos((v1.Dot(v2)) / (v1.Len() * v2.Len()))
}

// AngleBetween returns the angle between v1 and v2 in radian.
func (v1 *Vec3) AngleBetween(v2 *Vec3) float32 {
	return math.Acos((v1.Dot(v2)) / (v1.Len() * v2.Len()))
}

// AngleBetween returns the angle between v1 and v2 in radian.
func (v1 *Vec4) AngleBetween(v2 *Vec4) float32 {
	return math.Acos((v1.Dot(v2)) / (v1.Len() * v2.Len()))
}
