package lux

import (
	"github.com/luxengine/gl"
)

//ViewportChange is an alias to glViewport(0, 0, width, height)
func ViewportChange(width, height int32) {
	gl.Viewport.Set(0, 0, width, height)
}
