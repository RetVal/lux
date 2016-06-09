package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Program is a high level representation of OpenGL program object.
type Program uint32

// CurrentProgram returns the currently active program.
func CurrentProgram() Program {
	var v int32
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &v)
	return Program(v)
}

//CreateProgram is an alias to glCreateProgram.
func CreateProgram() Program {
	p := Program(gl.CreateProgram())
	return p
}

//AttachShader is an alias to glAttachShader.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glAttachShader.xml
func (p Program) AttachShader(shader uint32) {
	gl.AttachShader(uint32(p), shader)
}

//Link is an alias to glLinkProgram.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glLinkProgram.xml
func (p Program) Link() {
	gl.LinkProgram(uint32(p))
}

//Delete is an alias to glDeleteProgram. The progrma should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteProgram.xml
func (p Program) Delete() {
	gl.DeleteProgram(uint32(p))
}

//Use is an alias to glUseProgram(p).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
func (p Program) Use() {
	gl.UseProgram(uint32(p))
}

//Unuse is an alias to glUseProgram(0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
func (Program) Unuse() {
	gl.UseProgram(0)
}

//GetUniformLocation is an alias to glGetUniformLocation(p, name).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformLocation.xml
func (p Program) GetUniformLocation(name string) UniformLocation {
	return UniformLocation(gl.GetUniformLocation(uint32(p), gl.Str(name+"\x00")))
}

//BindFragDataLocation is an alias to glBindFragDataLocation(p, color, name).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindFragDataLocation.xml
func (p Program) BindFragDataLocation(color uint32, name string) {
	gl.BindFragDataLocation(uint32(p), color, gl.Str(name+"\x00"))
}

// GetDeleteStatus returns true if program is currently flagged for deletion,
// and false otherwise.
func (p Program) GetDeleteStatus() bool {
	var v int32
	gl.GetProgramiv(uint32(p), gl.DELETE_STATUS, &v)
	return v == gl.TRUE
}

// GetLinkStatus returns true if the last link operation on program was
// successful, and false otherwise.
func (p Program) GetLinkStatus() bool {
	var v int32
	gl.GetProgramiv(uint32(p), gl.LINK_STATUS, &v)
	return v == gl.TRUE
}

// GetValidateStatus returns true or if the last validation operation on
// program was successful, and false otherwise.
func (p Program) GetValidateStatus() bool {
	var v int32
	gl.GetProgramiv(uint32(p), gl.VALIDATE_STATUS, &v)
	return v == gl.TRUE
}

// GetInfoLogLength returns the number of characters in the information log for
// program including the null termination character (i.e., the size of the
// character buffer required to store the information log). If program has no
// information log, a value of 0 is returned.
func (p Program) GetInfoLogLength() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.INFO_LOG_LENGTH, &v)
	return int(v)
}

// GetInfoLog returns the information log for a program object.
func (p Program) GetInfoLog() string {
	infolength := int32(p.GetInfoLogLength())
	var actualLength int32
	l := make([]byte, infolength+1, infolength+1)
	gl.GetProgramInfoLog(uint32(p), infolength, &actualLength, &l[0])
	return string(l)
}

// GetNumAttachedShaders returns the number of shader objects attached to program.
func (p Program) GetNumAttachedShaders() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.ATTACHED_SHADERS, &v)
	return int(v)
}

// GetActiveAttributes returns the number of active attribute variables for
// program.
func (p Program) GetActiveAttributes() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.ACTIVE_ATTRIBUTES, &v)
	return int(v)
}

// GetActiveAttributeMaxLength returns the length of the longest active
// attribute name for program, including the null termination character (i.e.,
// the size of the character buffer required to store the longest attribute
// name). If no active attributes exist, 0 is returned.
func (p Program) GetActiveAttributeMaxLength() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.ACTIVE_ATTRIBUTE_MAX_LENGTH, &v)
	return int(v)
}

// GetNumActiveUniforms returns the number of active uniform variables for
// program.
func (p Program) GetNumActiveUniforms() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.ACTIVE_UNIFORMS, &v)
	return int(v)
}

// GetActiveUniformMaxLength returns the length of the longest active uniform
// variable name for program, including the null termination character (i.e.,
// the size of the character buffer required to store the longest uniform
// variable name). If no active uniform variables exist, 0 is returned.
func (p Program) GetActiveUniformMaxLength() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.ACTIVE_UNIFORM_MAX_LENGTH, &v)
	return int(v)
}

// GetTransformFeedbackBufferMode returns a symbolic constant indicating the
// buffer mode used when transform feedback is active. This may be
// GL_SEPARATE_ATTRIBS or GL_INTERLEAVED_ATTRIBS.
func (p Program) GetTransformFeedbackBufferMode() int32 {
	var v int32
	gl.GetProgramiv(uint32(p), gl.TRANSFORM_FEEDBACK_BUFFER_MODE, &v)
	return v
}

// GetNumTransformFeedbackVaryings returns the number of varying variables to
// capture in transform feedback mode for the program.
func (p Program) GetNumTransformFeedbackVaryings() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.TRANSFORM_FEEDBACK_VARYINGS, &v)
	return int(v)
}

// GetTransformFeedbackVaryingMaxLength returns the length of the longest
// variable name to be used for transform feedback, including the
// null-terminator.
func (p Program) GetTransformFeedbackVaryingMaxLength() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH, &v)
	return int(v)
}

// GetNumGeometryVerticesOut returns the maximum number of vertices that the
// geometry shader in program will output.
func (p Program) GetNumGeometryVerticesOut() int {
	var v int32
	gl.GetProgramiv(uint32(p), gl.GEOMETRY_VERTICES_OUT, &v)
	return int(v)
}

// GetGeometryInputType returns a symbolic constant indicating the primitive
// type accepted as input to the geometry shader contained in program.
func (p Program) GetGeometryInputType() int32 {
	var v int32
	gl.GetProgramiv(uint32(p), gl.GEOMETRY_INPUT_TYPE, &v)
	return v
}

// GetGeometryOutputType returns a symbolic constant indicating the primitive
// type that will be output by the geometry shader contained in program.
func (p Program) GetGeometryOutputType() int32 {
	var v int32
	gl.GetProgramiv(uint32(p), gl.GEOMETRY_OUTPUT_TYPE, &v)
	return v
}

// GetAttachedShaders returns a slice of all the attached shaders.
func (p Program) GetAttachedShaders() []Shader {
	num := int32(p.GetNumAttachedShaders())
	shaders := make([]uint32, num, num)
	var t int32
	gl.GetAttachedShaders(uint32(p), num, &t, &shaders[0])
	out := make([]Shader, 0, len(shaders))
	for _, s := range shaders {
		out = append(out, Shader(s))
	}
	return out
}
