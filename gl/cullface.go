package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type cullfacer struct{}

// CullFace is the global variable to the CullFace API
var CullFace cullfacer

// Enable enables GL_CULL_FACE
func (cullfacer) Enable() {
	gl.Enable(gl.CULL_FACE)
}

// Disable disables GL_CULL_FACE
func (cullfacer) Disable() {
	gl.Disable(gl.CULL_FACE)
}

// Front calls CullFace front
func (cullfacer) Front() {
	gl.CullFace(gl.FRONT)
}

// Back calls CullFace back
func (cullfacer) Back() {
	gl.CullFace(gl.FRONT)
}

// FrontAndBack calls CullFace frontAndBack
func (cullfacer) FrontAndBack() {
	gl.CullFace(gl.FRONT_AND_BACK)
}
