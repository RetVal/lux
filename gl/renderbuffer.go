package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//RenderBuffer is the high-level representation of OpenGL render buffer.
type RenderBuffer uint32

//GenRenderBuffer is an alias to glGenRenderbuffers(1, &b).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenRenderbuffers.xml
func GenRenderBuffer() RenderBuffer {
	var buf uint32
	gl.GenRenderbuffers(1, &buf)
	return RenderBuffer(buf)
}

//GenRenderBuffers is an alias to glGenRenderbuffers(n, &b).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenRenderbuffers.xml
func GenRenderBuffers(n int32) []RenderBuffer {
	buf := make([]RenderBuffer, n)
	gl.GenRenderbuffers(n, (*uint32)(&buf[0]))
	return buf
}

//Bind is an alias to glBindRenderbuffer(gl.RENDERBUFFER, uint32(rb)).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindRenderbuffer.xml
func (rb RenderBuffer) Bind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, uint32(rb))
}

//Unbind is an alias to glBindRenderbuffer(gl.RENDERBUFFER, 0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindRenderbuffer.xml
func (RenderBuffer) Unbind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
}

//Delete is an alias to glDeleteRenderbuffers. This render buffer should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteRenderbuffers.xml
func (rb RenderBuffer) Delete() {
	gl.DeleteRenderbuffers(1, (*uint32)(&rb))
}

//Storage is an alias to glRenderbufferStorage.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glRenderbufferStorage.xml
func (RenderBuffer) Storage(internalformat uint32, width, height int32) {
	//RENDERBUFFER is the only possible value
	gl.RenderbufferStorage(gl.RENDERBUFFER, internalformat, width, height)
}
