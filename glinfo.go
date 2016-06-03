package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//On load will contain all available OpenGL extension.
var Extensions map[string]bool

func init() {
	Extensions = make(map[string]bool)
}

//GetOpenglVersion will return the current OpenGL version.
func GetOpenglVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

//QueryExtentions will grab every extension currently loaded and populate lux.Extensions.
func QueryExtentions() {
	var numExtensions int32
	gl.GetIntegerv(gl.NUM_EXTENSIONS, &numExtensions)
	for i := int32(0); i < numExtensions; i++ {
		extension := gl.GoStr(gl.GetStringi(gl.EXTENSIONS, uint32(i)))
		Extensions[extension] = true
	}
}
