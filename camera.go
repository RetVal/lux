package lux

import (
	"github.com/luxengine/glm"
)

// Camera contains a view and projection matrix.
type Camera struct {
	View       glm.Mat4
	Projection glm.Mat4
	Pos        glm.Vec3
}

// SetPerspective sets the projection of this camera to a perspective
// projection.
func (c *Camera) SetPerspective(angle, ratio, zNear, zFar float32) {
	c.Projection = glm.Perspective(angle, ratio, zNear, zFar)
}

// SetOrtho sets the projection of this camera to an orthographic projection.
func (c *Camera) SetOrtho(left, right, bottom, top, near, far float32) {
	c.Projection = glm.Ortho(left, right, bottom, top, near, far)
}

// func to project from 2d to 3d
// func to project from 3d to 2d

// LookAtval sets the camera view direction by value.
func (c *Camera) LookAtval(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float32) {
	c.View = glm.LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ)
	c.Pos = glm.Vec3{eyeX, eyeY, eyeZ}
}

// LookAtVec sets the camera view direction by vectors.
func (c *Camera) LookAtVec(eye, center, up *glm.Vec3) {
	c.View = glm.LookAt(eye.X, eye.Y, eye.Z, center.X, center.Y, center.Z, up.X, up.Y, up.Z)
	c.Pos = *eye
}
