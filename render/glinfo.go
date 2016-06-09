package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Extensions contains all available OpenGL extensions.
var Extensions = make(map[string]struct{})

// GetOpenglVersion will return the current OpenGL version.
func GetOpenglVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

// QueryExtentions will grab every extension currently loaded and populate
// Extensions.
func QueryExtentions() {
	var numExtensions int32
	gl.GetIntegerv(gl.NUM_EXTENSIONS, &numExtensions)
	for i := int32(0); i < numExtensions; i++ {
		extension := gl.GoStr(gl.GetStringi(gl.EXTENSIONS, uint32(i)))
		Extensions[extension] = struct{}{}
	}
}
