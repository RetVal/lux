package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Framebuffer is a high level representation of OpenGL framebuffer object.
type Framebuffer uint32

// GenFramebuffer is an alias to glGenFramebuffers(1, &fbo).
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenFramebuffers.xml
func GenFramebuffer() Framebuffer {
	var fbo uint32
	gl.GenFramebuffers(1, &fbo)
	return Framebuffer(fbo)
}

// GenFramebuffers is an alias to glGenFramebuffers(n, &fbo).
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenFramebuffers.xml
func GenFramebuffers(n int32) []Framebuffer {
	fbos := make([]Framebuffer, n)
	gl.GenFramebuffers(n, (*uint32)(&fbos[0]))
	return fbos
}

// Bind is an alias to glBindFramebuffer(gl.FRAMEBUFFER, fbo).
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindFramebuffer.xml
func (fbo Framebuffer) Bind(target uint32) {
	gl.BindFramebuffer(target, uint32(fbo))
}

// Unbind is an alias to glBindFramebuffer(gl.FRAMEBUFFER, 0)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindFramebuffer.xml
func (Framebuffer) Unbind(target uint32) {
	gl.BindFramebuffer(target, 0)
}

// RenderBuffer is an alias to glFramebufferRenderbuffer(target, attachement, gl.RENDERBUFFER, renderbuffer).
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferRenderbuffer.xml
func (Framebuffer) RenderBuffer(target, attachement uint32, renderbuffer RenderBuffer) {
	gl.FramebufferRenderbuffer(target, attachement, gl.RENDERBUFFER, uint32(renderbuffer))
}

// Texture is an alias to glFramebufferTexture(target, attachement, texture, level).
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture.xml
func (Framebuffer) Texture(target, attachement uint32, texture Texture, level int32) {
	gl.FramebufferTexture(target, attachement, uint32(texture), level)
}

// DrawBuffers is an alias to glDrawBuffers(len(attachements), &attachements[0]).
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDrawBuffers.xml
func (Framebuffer) DrawBuffers(attachements ...uint32) {
	gl.DrawBuffers(int32(len(attachements)), (*uint32)(&attachements[0]))
}

// DrawBuffer is an alias to glDrawBuffer(attachement).
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDrawBuffer.xml
func (Framebuffer) DrawBuffer(attachement uint32) {
	gl.DrawBuffer(attachement)
}

// ReadBuffer is as alias to glReadBuffer(attachement).
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glReadBuffer.xml
func (Framebuffer) ReadBuffer(attachement uint32) {
	gl.ReadBuffer(attachement)
}

// Delete is an alias to glDeleteFramebuffers(1, &fbo). The FBO should not be used after calling this.
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteFramebuffers.xml
func (fbo Framebuffer) Delete() {
	gl.DeleteFramebuffers(1, (*uint32)(&fbo))
}

// Status is an alis for glCheckFramebufferStatus.
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glCheckFramebufferStatus.xml
func (fbo Framebuffer) Status(target uint32) uint32 {
	return gl.CheckFramebufferStatus(target)
}
