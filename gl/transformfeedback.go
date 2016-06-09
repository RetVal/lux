package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//TODO:
//Add the glGet
//Add primitive mode enum
//Decide if we add IsTransformFeedback

//TransformFeedback is a high-level representation of OpenGL transform feedback object.
type TransformFeedback uint32

//GenTransformFeedback is an alias to glGenTransformFeedbacks(1,&tf).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glGenTransformFeedbacks.xhtml
func GenTransformFeedback() TransformFeedback {
	var tf uint32
	gl.GenTransformFeedbacks(1, &tf)
	return TransformFeedback(tf)
}

//GenTransformFeedbacks is an alias to glGenTransformFeedbacks(n,&va[0]).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glGenTransformFeedbacks.xhtml
func GenTransformFeedbacks(n int32) []TransformFeedback {
	tfs := make([]TransformFeedback, n)
	gl.GenTransformFeedbacks(n, (*uint32)(&tfs[0]))
	return tfs
}

//Bind is as alias to glBindTransformFeedback(tf).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glBindTransformFeedback.xhtml
func (tf TransformFeedback) Bind() {
	gl.BindTransformFeedback(gl.TRANSFORM_FEEDBACK, uint32(tf))
}

//Unbind is an alias to glBindTransformFeedback(0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glBindTransformFeedback.xhtml
func (TransformFeedback) Unbind() {
	gl.BindTransformFeedback(gl.TRANSFORM_FEEDBACK, 0)
}

//Delete is an alias for glDeleteTransformFeedbacks(tf). The TFO should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glDeleteTransformFeedbacks.xhtml
func (tf TransformFeedback) Delete() {
	gl.DeleteTransformFeedbacks(1, (*uint32)(&tf))
}

//Begin is an alias to glBeginTransformFeedback(primitiveMode).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glBeginTransformFeedback.xhtml
func (TransformFeedback) Begin(primitiveMode uint32) {
	gl.BeginTransformFeedback(primitiveMode)
}

//End is an alias to glEndTransformFeedback.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glEndTransformFeedback.xhtml (broken link?)
func (TransformFeedback) End() {
	gl.EndTransformFeedback()
}

//Pause is an alias to glPauseTransformFeedback.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glPauseTransformFeedback.xhtml
func (TransformFeedback) Pause() {
	gl.PauseTransformFeedback()
}

//Resume is an alias to glResumeTransformFeedback.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man/html/glResumeTransformFeedback.xhtml
func (TransformFeedback) Resume() {
	gl.ResumeTransformFeedback()
}

//BindBufferBase is an alias to glBindBufferBase
//
//Note: I'm not sure if this function is a transform feedback function or a buffer function.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindBufferBase.xml
func (TransformFeedback) BindBufferBase(target uint32, index uint32, buffer Buffer) {
	gl.BindBufferBase(target, index, uint32(buffer))
}
