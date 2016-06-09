package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type stencilObj struct{}

//Stencil is a proxy for all the stencil functions. It should make your code cleaner.
var Stencil = stencilObj{}

//Enable is an alias to glEnable(gl.STENCIL_TEST)
func (stencilObj) Enable() {
	gl.Enable(gl.STENCIL_TEST)
}

//Disable is an alias to glDisable(gl.STENCIL_TEST)
func (stencilObj) Disable() {
	gl.Disable(gl.STENCIL_TEST)
}

//Func is an alias to glStencilFunc(f, ref, mask)
func (stencilObj) Func(f uint32, ref int32, mask uint32) {
	gl.StencilFunc(f, ref, mask)
}

//Op is an alias to glStencilOp(sfail, zfail, zpass)
func (stencilObj) Op(sfail, zfail, zpass uint32) {
	gl.StencilOp(sfail, zfail, zpass)
}

//Mask is an alias to glStencilMask(mask)
func (stencilObj) Mask(mask uint32) {
	gl.StencilMask(mask)
}

//Stencil op possible values
const (
	Keep     = gl.KEEP
	Zero     = gl.ZERO
	Replace  = gl.REPLACE
	Incr     = gl.INCR
	IncrWrap = gl.INCR_WRAP
	Decr     = gl.DECR
	DecrWrap = gl.DECR_WRAP
	Invert   = gl.INVERT
)
