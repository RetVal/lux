package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

//TODO: make subtype buffers with restrained functions
//TODO: make an enum for targets

//Buffer is the high level representation of OpenGL Buffer.
type Buffer uint32

//GenBuffer is an alias to glGenBuffers(1, &b).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenBuffers.xml
func GenBuffer() Buffer {
	var buff uint32
	gl.GenBuffers(1, &buff)
	return Buffer(buff)
}

//GenBuffers is an alias to glGenBuffers(n, &b[0]).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenBuffers.xml
func GenBuffers(n int32) []Buffer {
	buffs := make([]Buffer, n)
	gl.GenBuffers(n, (*uint32)(&buffs[0]))
	return buffs
}

//Bind is an alias to glBindBuffer(target, b).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindBuffer.xml
func (b Buffer) Bind(target uint32) {
	gl.BindBuffer(uint32(target), uint32(b))
}

//Unbind is an alias to glBindBuffer(target, 0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindBuffer.xml
func (Buffer) Unbind(target uint32) {
	gl.BindBuffer(target, 0)
}

//Data is an alias for glBufferData.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBufferData.xml
func (Buffer) Data(target uint32, size int, data unsafe.Pointer, usage uint32) {
	gl.BufferData(target, size, data, usage)
}

//Delete is an alias to glDeleteBuffers(&b). The buffer should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteBuffers.xml
func (b Buffer) Delete() {
	gl.DeleteBuffers(1, (*uint32)(&b))
}
