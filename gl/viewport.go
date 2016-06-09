package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type viewport struct{}

// Viewport is the global variable used to access the viewport API.
var Viewport viewport

func (viewport) Set(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

func (viewport) Get() (int32, int32, int32, int32) {
	var values [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &values[0])
	return values[0], values[1], values[2], values[3]
}

func (viewport) GetMaxDims() (int32, int32) {
	var values [2]int32
	gl.GetIntegerv(gl.MAX_VIEWPORT_DIMS, &values[0])
	return values[0], values[1]
}
