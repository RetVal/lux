package lux

import (
	"fmt" //for error

	"github.com/luxengine/lux/gl"
)

// NewProgram will create an OpenGL program from the given shaders, any
// combinations can be used.
func NewProgram(shaders ...Shader) (gl.Program, error) {
	program := gl.CreateProgram()

	for _, shader := range shaders {
		program.AttachShader(shader.Loc)
	}
	program.Link()

	if !program.GetLinkStatus() {
		return 0, fmt.Errorf("failed to link program: %v", program.GetInfoLog())
	}

	return program, nil
}

// MustNotGLError will check opengl for error and panic if one was generated.
func MustNotGLError() {
	if err := gl.GetError(); err != nil {
		panic(err)
	}
}

// GenRGBTexture2D is a utility function to generate an empty 2D textures of
// size (width,height), internal format RGB and data type UNSIGNED_BYTE.
func GenRGBTexture2D(width, height int32) gl.Texture2D {
	tex := gl.GenTexture2D()
	tex.Bind()
	tex.MinFilter(gl.LINEAR)
	tex.MagFilter(gl.LINEAR)
	tex.WrapS(gl.CLAMP_TO_EDGE)
	tex.WrapT(gl.CLAMP_TO_EDGE)
	tex.TexImage2D(0, gl.RGB, width, height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	return tex
}
