package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

//TODO:
//Add all the glGet

//VertexArray is the high level representation of OpenGL vertex array object.
type VertexArray uint32

//GenVertexArray is an alias for glGenVertexArrays(1, &vao).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenVertexArrays.xml
func GenVertexArray() VertexArray {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	return VertexArray(vao)
}

//GenVertexArrays is an alias for glGenVertexArrays(n, &vao[0]).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenVertexArrays.xml
func GenVertexArrays(n int32) []VertexArray {
	vao := make([]VertexArray, n)
	gl.GenVertexArrays(n, (*uint32)(&vao[0]))
	return vao
}

//Bind is an alias to glBindVertexArray(vao).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindVertexArray.xml
func (vao VertexArray) Bind() {
	gl.BindVertexArray(uint32(vao))
}

//Unbind is an alias to glBindVertexArray(0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindVertexArray.xml
func (VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

//Delete is an alias to glDeleteVertexArray(vao). The vao should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteVertexArrays.xml
func (vao VertexArray) Delete() {
	gl.DeleteVertexArrays(1, (*uint32)(&vao))
}

//EnableVertexAttribArray is an alias to glEnableVertexAttribArray(index).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glEnableVertexAttribArray.xml
func (VertexArray) EnableVertexAttribArray(index uint32) {
	gl.EnableVertexAttribArray(index)
}

//DisableVertexAttribArray is an alias to gl.DisableVertexAttribArray(index).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDisableVertexAttribArray.xml
func (VertexArray) DisableVertexAttribArray(index uint32) {
	gl.DisableVertexAttribArray(index)
}

//VertexAttribPointer is an alias for glVertexAttribPointer.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribPointer.xml
func (VertexArray) VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	gl.VertexAttribPointer(index, size, xtype, normalized, stride, pointer)
}

//VertexAttribIPointer is an alias for glVertexAttribIPointer.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribPointer.xml
func (VertexArray) VertexAttribIPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	gl.VertexAttribIPointer(index, size, xtype, stride, pointer)
}

//VertexAttribLPointer is an alias for glVertexAttribLPointer.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribPointer.xml
func (VertexArray) VertexAttribLPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	gl.VertexAttribLPointer(index, size, xtype, stride, pointer)
}
