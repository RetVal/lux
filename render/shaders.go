package lux

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Shader represent an OpenGL shader object along with it's type.
type Shader struct {
	Stype uint32
	Loc   uint32
}

// Delete releases all the resources held by this shader object.
func (s Shader) Delete() {
	gl.DeleteShader(s.Loc)
}

// CompileShader will take the shader source and generate a shader object based
// on which type you give it. Returns an error if it fails.
func CompileShader(source string, shaderType uint32) (Shader, error) {
	shader := gl.CreateShader(uint32(shaderType))

	csource := gl.Str(source)
	gl.ShaderSource(shader, 1, &csource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return Shader{}, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return Shader{shaderType, shader}, nil
}
