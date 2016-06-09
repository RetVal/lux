package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Shader is a high-level representation of the OpenGL shader object, can be any
// type (VERTEX_SHADER, FRAGMENT_SHADER, etc).
type Shader uint32

// CreateShader is an alias to glCreateShader().
//
// Doc: https://www.opengl.org/sdk/docs/man3/xhtml/glCreateShader.xml
func CreateShader(shaderType uint32) Shader {
	return Shader(gl.CreateShader(shaderType))
}

// Delete is an alias to glDeleteShader().
//
// Doc: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteShader.xml
func (s Shader) Delete() {
	gl.DeleteShader(uint32(s))
}

// Source is an alias to glShaderSource()
//
// Doc: https://www.opengl.org/sdk/docs/man3/xhtml/glShaderSource.xml
func (s Shader) Source(count int32, xstring **uint8, length *int32) {
	gl.ShaderSource(uint32(s), count, xstring, length)
}

// Compile is an alias to glCompileShader(s)
//
// Doc: https://www.opengl.org/sdk/docs/man3/xhtml/glCompileShader.xml
func (s Shader) Compile() {
	gl.CompileShader(uint32(s))
}

// GetInfoLog is an alias to GetShaderInfoLog
//
// Doc: https://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderInfoLog.xml
func (s Shader) GetInfoLog() string {
	l := int32(s.GetInfoLogLength())
	buf := make([]byte, l+1, l+1)
	var length int32
	gl.GetShaderInfoLog(uint32(s), l, &length, &buf[0])
	return string(buf)
}

// GetSource is an alias to GetShaderInfoLog
//
// Doc: https://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderSource.xml
func (s Shader) GetSource() string {
	l := int32(s.GetInfoLogLength())
	buf := make([]byte, l+1, l+1)
	var length int32
	gl.GetShaderSource(uint32(s), l, &length, &buf[0])
	return string(buf)
}

// GetShaderType returns this shaders shader type.
func (s Shader) GetShaderType() int32 {
	var t int32
	gl.GetShaderiv(uint32(s), gl.SHADER_TYPE, &t)
	return t
}

// GetDeleteStatus returns true if shader is currently flagged for deletion,
// and false otherwise.
func (s Shader) GetDeleteStatus() bool {
	var t int32
	gl.GetShaderiv(uint32(s), gl.DELETE_STATUS, &t)
	return t == gl.TRUE
}

// GetCompileStatus returns true if the last compile operation on shader was
// successful, and false otherwise.
func (s Shader) GetCompileStatus() bool {
	var t int32
	gl.GetShaderiv(uint32(s), gl.COMPILE_STATUS, &t)
	return t == gl.TRUE
}

// GetInfoLogLength returns the number of characters in the information log for
// shader including the null termination character (i.e., the size of the
// character buffer required to store the information log). If shader has no
// information log, a value of 0 is returned.
func (s Shader) GetInfoLogLength() int {
	var t int32
	gl.GetShaderiv(uint32(s), gl.INFO_LOG_LENGTH, &t)
	return int(t)
}

// GetShaderSourceLength returns the length of the concatenation of the source
// strings that make up the shader source for the shader, including the null
// termination character. (i.e., the size of the character buffer required to
// store the shader source). If no source code exists, 0 is returned.
func (s Shader) GetShaderSourceLength() int {
	var t int32
	gl.GetShaderiv(uint32(s), gl.SHADER_SOURCE_LENGTH, &t)
	return int(t)
}
