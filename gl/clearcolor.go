package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type clearcolor struct{}

// ClearColor is the global variable to the CleaColor API
var ClearColor clearcolor

// Set sets clear color.
func (clearcolor) Set(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

// Get returns the clear color.
func (clearcolor) Get() (float32, float32, float32, float32) {
	var color [4]float32
	gl.GetFloatv(gl.COLOR_CLEAR_VALUE, &color[0])
	return color[0], color[1], color[2], color[3]
}
