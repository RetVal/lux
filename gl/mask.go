package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Depth is a global variable to encapsulate depth mask related functions.
var Depth DepthMask

//DepthMask is a type to encapsulate depth mask related functions.
type DepthMask struct{}

//Total is an alias to glDepthMask(false)
func (DepthMask) Total() {
	gl.DepthMask(false)
}

//None is an alias to glDepthMask(false)
func (DepthMask) None() {
	gl.DepthMask(true)
}

//Mask is an alias to glDepthMask(mask)
func (DepthMask) Mask(mask bool) {
	gl.DepthMask(mask)
}

//Color is a global varialbe to encapsulate color mask related functions.
var Color ColorMask

//ColorMask is a type to encapsulate color mask related functions.
type ColorMask struct{}

//Total is an alias to glColorMask(false, false, false, false)
func (ColorMask) Total() {
	gl.ColorMask(false, false, false, false)
}

//None is an alias to glColorMask(true, true, true, true)
func (ColorMask) None() {
	gl.ColorMask(true, true, true, true)
}

//Mask is an alias to glColorMask(r, g, b, a)
func (ColorMask) Mask(r, g, b, a bool) {
	gl.ColorMask(r, g, b, a)
}
