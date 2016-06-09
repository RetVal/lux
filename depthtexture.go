package lux

import (
	"github.com/luxengine/lux/gl"
)

//GenDepthTexture is a utility function to generate a depth Texture2D
func GenDepthTexture(width, height int32) gl.Texture2D {
	tex := gl.GenTexture2D()
	tex.Bind()
	tex.MinFilter(gl.NEAREST)
	tex.MagFilter(gl.NEAREST)
	tex.WrapS(gl.CLAMP_TO_EDGE)
	tex.WrapT(gl.CLAMP_TO_EDGE)
	tex.TexImage2D(0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	tex.Unbind()
	return tex
}
