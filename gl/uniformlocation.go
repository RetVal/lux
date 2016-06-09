package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//UniformLocation is the high-level representation of openGL shader uniform location, generated via the program object. The goal is to specify exactly which type this variable is so that you can restrict the Go compiler on what function you can use on this variable.
//    mvp := program.GetUniformLocation("MVP").ToMatrix4()
//    //you can only call UniformMatrix4fv on mvp
type UniformLocation int32

//Uniform1f is an alias to glUniform1f.
func (ul UniformLocation) Uniform1f(v0 float32) {
	gl.Uniform1f(int32(ul), v0)
}

//Uniform2f is an alias to glUniform2f.
func (ul UniformLocation) Uniform2f(v0, v1 float32) {
	gl.Uniform2f(int32(ul), v0, v1)
}

//Uniform3f is an alias to glUniform3f.
func (ul UniformLocation) Uniform3f(v0, v1, v2 float32) {
	gl.Uniform3f(int32(ul), v0, v1, v2)
}

//Uniform4f is an alias to glUniform4f.
func (ul UniformLocation) Uniform4f(v0, v1, v2, v3 float32) {
	gl.Uniform4f(int32(ul), v0, v1, v2, v3)
}

//Uniform1i is an alias to glUniform1i.
func (ul UniformLocation) Uniform1i(v0 int32) {
	gl.Uniform1i(int32(ul), v0)
}

//Uniform2i is an alias to glUniform2i.
func (ul UniformLocation) Uniform2i(v0, v1 int32) {
	gl.Uniform2i(int32(ul), v0, v1)
}

//Uniform3i is an alias to glUniform3i.
func (ul UniformLocation) Uniform3i(v0, v1, v2 int32) {
	gl.Uniform3i(int32(ul), v0, v1, v2)
}

//Uniform4i is an alias to glUniform4i.
func (ul UniformLocation) Uniform4i(v0, v1, v2, v3 int32) {
	gl.Uniform4i(int32(ul), v0, v1, v2, v3)
}

//Uniform1ui is an alias to glUniform1ui.
func (ul UniformLocation) Uniform1ui(v0 uint32) {
	gl.Uniform1ui(int32(ul), v0)
}

//Uniform2ui is an alias to glUniform2ui.
func (ul UniformLocation) Uniform2ui(v0, v1 uint32) {
	gl.Uniform2ui(int32(ul), v0, v1)
}

//Uniform3ui is an alias to glUniform3ui.
func (ul UniformLocation) Uniform3ui(v0, v1, v2 uint32) {
	gl.Uniform3ui(int32(ul), v0, v1, v2)
}

//Uniform4ui is an alias to glUniform4ui.
func (ul UniformLocation) Uniform4ui(v0, v1, v2, v3 uint32) {
	gl.Uniform4ui(int32(ul), v0, v1, v2, v3)
}

//Uniform1fv is an alias to glUniform1fv.
func (ul UniformLocation) Uniform1fv(count int32, value *float32) {
	gl.Uniform1fv(int32(ul), count, value)
}

//Uniform2fv is an alias to glUniform2fv.
func (ul UniformLocation) Uniform2fv(count int32, value *float32) {
	gl.Uniform2fv(int32(ul), count, value)
}

//Uniform3fv is an alias to glUniform3fv.
func (ul UniformLocation) Uniform3fv(count int32, value *float32) {
	gl.Uniform3fv(int32(ul), count, value)
}

//Uniform4fv is an alias to glUniform4fv.
func (ul UniformLocation) Uniform4fv(count int32, value *float32) {
	gl.Uniform4fv(int32(ul), count, value)
}

//Uniform1iv is an alias to glUniform1iv.
func (ul UniformLocation) Uniform1iv(count int32, value *int32) {
	gl.Uniform1iv(int32(ul), count, value)
}

//Uniform2iv is an alias to glUniform2iv.
func (ul UniformLocation) Uniform2iv(count int32, value *int32) {
	gl.Uniform2iv(int32(ul), count, value)
}

//Uniform3iv is an alias to glUniform3iv.
func (ul UniformLocation) Uniform3iv(count int32, value *int32) {
	gl.Uniform3iv(int32(ul), count, value)
}

//Uniform4iv is an alias to glUniform4iv.
func (ul UniformLocation) Uniform4iv(count int32, value *int32) {
	gl.Uniform4iv(int32(ul), count, value)
}

//Uniform1uiv is an alias to glUniform1uiv.
func (ul UniformLocation) Uniform1uiv(count int32, value *uint32) {
	gl.Uniform1uiv(int32(ul), count, value)
}

//Uniform2uiv is an alias to glUniform2uiv.
func (ul UniformLocation) Uniform2uiv(count int32, value *uint32) {
	gl.Uniform2uiv(int32(ul), count, value)
}

//Uniform3uiv is an alias to glUniform3uiv.
func (ul UniformLocation) Uniform3uiv(count int32, value *uint32) {
	gl.Uniform3uiv(int32(ul), count, value)
}

//Uniform4uiv is an alias to glUniform4uiv.
func (ul UniformLocation) Uniform4uiv(count int32, value *uint32) {
	gl.Uniform4uiv(int32(ul), count, value)
}

//UniformMatrix2fv is an alias to glUniformMatrix2fv.
func (ul UniformLocation) UniformMatrix2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2fv(int32(ul), count, transpose, value)
}

//UniformMatrix3fv is an alias to glUniformMatrix3fv.
func (ul UniformLocation) UniformMatrix3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3fv(int32(ul), count, transpose, value)
}

//UniformMatrix4fv is an alias to glUniformMatrix4fv.
func (ul UniformLocation) UniformMatrix4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4fv(int32(ul), count, transpose, value)
}

//UniformMatrix2x3fv is an alias to glUniformMatrix2x3fv.
func (ul UniformLocation) UniformMatrix2x3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2x3fv(int32(ul), count, transpose, value)
}

//UniformMatrix3x2fv is an alias to glUniformMatrix3x2fv.
func (ul UniformLocation) UniformMatrix3x2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3x2fv(int32(ul), count, transpose, value)
}

//UniformMatrix2x4fv is an alias to glUniformMatrix2x4fv.
func (ul UniformLocation) UniformMatrix2x4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2x4fv(int32(ul), count, transpose, value)
}

//UniformMatrix4x2fv is an alias to glUniformMatrix4x2fv.
func (ul UniformLocation) UniformMatrix4x2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4x2fv(int32(ul), count, transpose, value)
}

//UniformMatrix3x4fv is an alias to glUniformMatrix3x4fv.
func (ul UniformLocation) UniformMatrix3x4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3x4fv(int32(ul), count, transpose, value)
}

//UniformMatrix4x3fv is an alias to glUniformMatrix4x3fv.
func (ul UniformLocation) UniformMatrix4x3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4x3fv(int32(ul), count, transpose, value)
}
