package glm

import (
	"github.com/luxengine/lux/math"
)

// Ortho returns a Mat4 that represents a orthographic projection from the given
// arguments.
func Ortho(left, right, bottom, top, near, far float32) Mat4 {
	rml, tmb, fmn := 1/(right-left), 1/(top-bottom), 1/(far-near)

	return Mat4{
		2 * rml, 0, 0, 0,
		0, 2 * tmb, 0, 0,
		0, 0, -2 * fmn, 0,
		-(right + left) * rml, -(top + bottom) * tmb, -(far + near) * fmn, 1,
	}
}

// OrthoIn is a memory friendly version of Ortho.
func OrthoIn(left, right, bottom, top, near, far float32, p *Mat4) {
	rml, tmb, fmn := 1/(right-left), 1/(top-bottom), 1/(far-near)

	p[0] = 2 * rml
	p[1] = 0
	p[2] = 0
	p[3] = 0

	p[4] = 0
	p[5] = 2 * tmb
	p[6] = 0
	p[7] = 0

	p[8] = 0
	p[9] = 0
	p[10] = -2 * fmn
	p[11] = 0

	p[12] = -(right + left) * rml
	p[13] = -(top + bottom) * tmb
	p[14] = -(far + near) * fmn
	p[15] = 1
}

// Ortho2D is equivalent to Ortho with the near and far planes being -1 and 1,
// respectively.
func Ortho2D(left, right, bottom, top float32) Mat4 {
	return Ortho(left, right, bottom, top, -1, 1)
}

// Perspective returns a Mat4 representing a perspective projection given fovy
// in radian, aspect as width/height, near and far as the distance from origin.
func Perspective(fovy, aspect, near, far float32) Mat4 {
	nmf, f := 1./(near-far), 1./math.Tan(fovy/2.0)
	return Mat4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (near + far) * nmf, -1,
		0, 0, (2. * far * near) * nmf, 0,
	}
}

// PerspectiveIn is a memory friendly version of Perspective.
func PerspectiveIn(fovy, aspect, near, far float32, p *Mat4) {
	nmf, f := 1/(near-far), 1./math.Tan(fovy/2.0)

	p[0] = f / aspect
	p[1] = 0
	p[2] = 0
	p[3] = 0

	p[4] = 0
	p[5] = f
	p[6] = 0
	p[7] = 0

	p[8] = 0
	p[9] = 0
	p[10] = (near + far) * nmf
	p[11] = -1

	p[12] = 0
	p[13] = 0
	p[14] = (2. * far * near) * nmf
	p[15] = 0
}

// Frustum returns a Mat4 representing a frustrum transform (squared pyramid with the top cut off)
func Frustum(left, right, bottom, top, near, far float32) Mat4 {
	rml, tmb, fmn := 1/(right-left), 1/(top-bottom), 1/(far-near)
	A, B, C, D := (right+left)*rml, (top+bottom)*tmb, -(far+near)*fmn, -(2*far*near)*fmn

	return Mat4{
		(2 * near) * rml, 0, 0, 0,
		0, (2 * near) * tmb, 0, 0,
		A, B, C, -1,
		0, 0, D, 0,
	}
}

// LookAt returns a Mat4 that represents a camera transform from the given
// arguments.
func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float32) Mat4 {
	return LookAtV(
		&Vec3{eyeX, eyeY, eyeZ},
		&Vec3{centerX, centerY, centerZ},
		&Vec3{upX, upY, upZ},
	)
}

// LookAtV generates a transform matrix from world space into the specific eye
// space. Up must be normalized
func LookAtV(eye, center, up *Vec3) Mat4 {
	var f Vec3
	f.SubOf(center, eye)
	f.Normalize()
	var s Vec3
	s.CrossOf(&f, up)
	//s.Normalize()
	var u Vec3
	u.CrossOf(&s, &f)

	M := Mat4{
		s.X, u.X, -f.X, 0,
		s.Y, u.Y, -f.Y, 0,
		s.Z, u.Z, -f.Z, 0,
		0, 0, 0, 1,
	}

	t := Translate3D(-eye.X, -eye.Y, -eye.Z)
	return M.Mul4(&t)
}

// LookAtVIn is a memory friendly version o LookAtV.
func LookAtVIn(eye, center, up *Vec3, v *Mat4) {
	f := Vec3{
		center.X - eye.X,
		center.Y - eye.Y,
		center.Z - eye.Z,
	}
	flen := 1.0 / (f.X*f.X + f.Y*f.Y + f.Z*f.Z)
	f.X *= flen
	f.Y *= flen
	f.Z *= flen

	s := Vec3{f.Y*up.Z - f.Z*up.Y, f.Z*up.X - f.X*up.Z, f.X*up.Y - f.Y*up.X}

	v[0] = s.X
	v[1] = s.Y*f.Z - s.Z*f.Y
	v[2] = -f.X
	v[3] = 0

	v[4] = s.Y
	v[5] = s.Z*f.X - s.X*f.Z
	v[6] = -f.Y
	v[7] = 0

	v[8] = s.Z
	v[9] = s.X*f.Y - s.Y*f.X
	v[10] = -f.Z
	v[11] = 0
	v[12] = -(v[0]*eye.X + v[4]*eye.Y + v[8]*eye.Z)
	v[13] = -(v[1]*eye.X + v[5]*eye.Y + v[9]*eye.Z)
	v[14] = -(v[2]*eye.X + v[6]*eye.Y + v[10]*eye.Z)
	v[15] = 1
}

// Project transforms a set of coordinates from object space (in obj) to window
// coordinates (with depth)
//
// Window coordinates are continuous, not discrete, so you won't get exact pixel
// locations without rounding.
func Project(obj *Vec3, modelview, projection *Mat4, initialX, initialY, width, height int) Vec3 {
	obj4 := obj.Vec4(1)

	pm := projection.Mul4(modelview)
	vpp := pm.Mul4x1(&obj4)
	return Vec3{
		float32(initialX) + (float32(width)*(vpp.X+1))*0.5,
		float32(initialY) + (float32(height)*(vpp.Y+1))*0.5,
		(vpp.Z + 1) * 0.5,
	}
}

// UnProject transforms a set of window coordinates to object space. If your MVP
// matrix is not invertible this will return garbage.
//
// Note that the projection may not be perfect if you use strict pixel locations
// rather than the exact values given by Project.
func UnProject(win *Vec3, modelview, projection *Mat4, initialX, initialY, width, height int) Vec3 {
	pm := projection.Mul4(modelview)
	inv := pm.Inverse()

	obj4 := inv.Mul4x1(&Vec4{
		(2 * (win.X - float32(initialX)) / float32(width)) - 1,
		(2 * (win.Y - float32(initialY)) / float32(height)) - 1,
		2*win.Z - 1,
		1.0,
	})
	obj := obj4.Vec3()

	//if obj4[3] > MinValue {}
	over := 1 / obj4.W
	obj.X *= over
	obj.Y *= over
	obj.Z *= over

	return obj
}
