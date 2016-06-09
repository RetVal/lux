package lux

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/lux/gl"
)

// ShadowCubeFBO is a packed framebuffer and cubemap with depth only for
// shadows.
type ShadowCubeFBO struct {
	framebuffer          gl.Framebuffer
	texture              gl.Texture
	projection, view, vp glm.Mat4
	program              gl.Program
	mvpUni               gl.UniformLocation
	width, height        int32
}

// NewShadowCubeFBO makes a new ShadowCubeFBO.
func NewShadowCubeFBO(width, height int32) *ShadowCubeFBO {
	fbo := gl.GenFramebuffer()
	fbo.Bind(gl.FRAMEBUFFER)

	tex := gl.GenTexture()
	tex.Bind(gl.TEXTURE_CUBE_MAP)

	tex.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	tex.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	tex.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	tex.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	tex.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	tex.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_COMPARE_MODE, gl.COMPARE_REF_TO_TEXTURE)
	tex.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_COMPARE_FUNC, gl.LEQUAL)

	tex.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_X, 0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)
	tex.TexImage2D(gl.TEXTURE_CUBE_MAP_NEGATIVE_X, 0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)

	tex.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_Y, 0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)
	tex.TexImage2D(gl.TEXTURE_CUBE_MAP_NEGATIVE_Y, 0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)

	tex.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_Z, 0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)
	tex.TexImage2D(gl.TEXTURE_CUBE_MAP_NEGATIVE_Z, 0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)

	fbo.Texture(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, tex, 0)

	fbo.DrawBuffer(gl.NONE)
	fbo.ReadBuffer(gl.NONE)

	fbo.Unbind(gl.FRAMEBUFFER)
	tex.Unbind(gl.TEXTURE_CUBE_MAP)

	return &ShadowCubeFBO{
		framebuffer: fbo,
		texture:     tex,
		//program:     a,
		width:  width,
		height: height,
	}
}
