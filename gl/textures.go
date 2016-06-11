package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

//texture interface maybe ?
/*
type Texture interface{
	Bind()
	Unbind()
	Delete()
	//think about parameters
	//Parameteri()
}
*/

//Texture is a high-level representation of the OpenGL texture object, can be any type (TEXTURE_2D, TEXTURE_1D, etc).
type Texture uint32

// GenTexture is an alias for
//	var t uint32
//	glGenTexture(1, &t)
//	return t
func GenTexture() Texture {
	var tex uint32
	gl.GenTextures(1, &tex)
	return Texture(tex)
}

// GenTextures is an alias for
//	t := make([]Texture, n)
//	gl.GenTextures(1, (*uint32)(&t[0]))
//	return t
func GenTextures(n int32) []Texture {
	tex := make([]Texture, n)
	gl.GenTextures(1, (*uint32)(&tex[0]))
	return tex
}

// Delete is an alias for gl.DeleteTextures(1, (*uint32)(&t))
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteTextures.xml
func (t Texture) Delete() {
	gl.DeleteTextures(1, (*uint32)(&t))
}

// Bind is an alias for gl.BindTexture(target, uint32(t))
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindTexture.xml
func (t Texture) Bind(target uint32) {
	gl.BindTexture(target, uint32(t))
}

// Unbind is an alias for gl.BindTexture(target, 0)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindTexture.xml
func (t Texture) Unbind(target uint32) {
	gl.BindTexture(target, 0)
}

// CopyTexImage1D is an alias for gl.CopyTexImage1D(target, level, internalformat, x, y, width, border)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexImage1D.xml
func (Texture) CopyTexImage1D(target uint32, level int32, internalformat uint32, x, y, width, border int32) {
	gl.CopyTexImage1D(target, level, internalformat, x, y, width, border)
}

// CopyTexImage2D is an alias for gl.CopyTexImage2D(target, level, internalformat, x, y, width, height, border)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexImage2D.xml
func (Texture) CopyTexImage2D(target uint32, level int32, internalformat uint32, x, y, width, height, border int32) {
	gl.CopyTexImage2D(target, level, internalformat, x, y, width, height, border)
}

// TexImage1D is an alias for gl.TexImage1D(target, level, internalFormat, width, border, format, xtype, data)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage1D.xml
func (Texture) TexImage1D(target uint32, level, internalFormat, width, border int32, format, xtype uint32, data unsafe.Pointer) {
	gl.TexImage1D(target, level, internalFormat, width, border, format, xtype, data)
}

// TexImage2D is an alias for gl.TexImage2D(target, level, internalFormat, width, height, border, format, xtype, data)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage2D.xml
func (Texture) TexImage2D(target uint32, level, internalFormat, width, height, border int32, format, xtype uint32, data unsafe.Pointer) {
	gl.TexImage2D(target, level, internalFormat, width, height, border, format, xtype, data)
}

// TexImage3D is an alias for gl.TexImage3D(target, level, internalFormat, width, height, depth, border, format, xtype, data)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage3D.xml
func (Texture) TexImage3D(target uint32, level, internalFormat, width, height, depth, border int32, format, xtype uint32, data unsafe.Pointer) {
	gl.TexImage3D(target, level, internalFormat, width, height, depth, border, format, xtype, data)
}

// TexParameterfv is an alias for gl.TexParameterfv(target, pname, params)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterfv.xml
func (Texture) TexParameterfv(target, pname uint32, params *float32) {
	gl.TexParameterfv(target, pname, params)
}

// TexParameteriv is an alias for gl.TexParameteriv(target, pname, params)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameteriv.xml
func (Texture) TexParameteriv(target, pname uint32, params *int32) {
	gl.TexParameteriv(target, pname, params)
}

// TexParameterIiv is an alias for gl.TexParameterIiv(target, pname, params)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterIiv.xml
func (Texture) TexParameterIiv(target, pname uint32, params *int32) {
	gl.TexParameterIiv(target, pname, params)
}

// TexParameteri is an alias for gl.TexParameteri(target, pname, param)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameteri.xml
func (Texture) TexParameteri(target, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}

// TexParameterIuiv is an alias for gl.TexParameterIuiv(target, pname, params)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterIuiv.xml
func (Texture) TexParameterIuiv(target, pname uint32, params *uint32) {
	gl.TexParameterIuiv(target, pname, params)
}

// GetTexParameterfv is an alias for gl.GetTexParameterfv(target, pname, params)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterfv.xml
func (Texture) GetTexParameterfv(target, pname uint32, params *float32) {
	gl.GetTexParameterfv(target, pname, params)
}

// GetTexParameteriv is an alias for gl.GetTexParameteriv(target, pname, params)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameteriv.xml
func (Texture) GetTexParameteriv(target, pname uint32, params *int32) {
	gl.GetTexParameteriv(target, pname, params)
}

// GetTexParameterIiv is an alias for gl.GetTexParameterIiv(target, pname, params)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterIiv.xml
func (Texture) GetTexParameterIiv(target, pname uint32, params *int32) {
	gl.GetTexParameterIiv(target, pname, params)
}

// GetTexParameterIuiv is an alias for gl.GetTexParameterIuiv(target, pname, params)
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterIuiv.xml
func (Texture) GetTexParameterIuiv(target, pname uint32, params *uint32) {
	gl.GetTexParameterIuiv(target, pname, params)
}

// IsTexture is an alias for gl.IsTexture(uint32(t))
//
// Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glIsTexture.xml
func (t Texture) IsTexture() bool {
	return gl.IsTexture(uint32(t))
}
