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

func GenTexture() Texture {
	var tex uint32
	gl.GenTextures(1, &tex)
	return Texture(tex)
}

func GenTextures(n int32) []Texture {
	tex := make([]Texture, n)
	gl.GenTextures(1, (*uint32)(&tex[0]))
	return tex
}

func (t Texture) Delete() {
	gl.DeleteTextures(1, (*uint32)(&t))
}

func (t Texture) Bind(target uint32) {
	gl.BindTexture(target, uint32(t))
}

func (t Texture) Unbind(target uint32) {
	gl.BindTexture(target, 0)
}

func (Texture) CopyTexImage1D(target uint32, level int32, internalformat uint32, x, y, width, border int32) {
	gl.CopyTexImage1D(target, level, internalformat, x, y, width, border)
}

func (Texture) CopyTexImage2D(target uint32, level int32, internalformat uint32, x, y, width, height, border int32) {
	gl.CopyTexImage2D(target, level, internalformat, x, y, width, height, border)
}

func (Texture) TexImage1D(target uint32, level, internalFormat, width, border int32, format, xtype uint32, data unsafe.Pointer) {
	gl.TexImage1D(target, level, internalFormat, width, border, format, xtype, data)
}

func (Texture) TexImage2D(target uint32, level, internalFormat, width, height, border int32, format, xtype uint32, data unsafe.Pointer) {
	gl.TexImage2D(target, level, internalFormat, width, height, border, format, xtype, data)
}

func (Texture) TexImage3D(target uint32, level, internalFormat, width, height, depth, border int32, format, xtype uint32, data unsafe.Pointer) {
	gl.TexImage3D(target, level, internalFormat, width, height, depth, border, format, xtype, data)
}

func (Texture) TexParameterfv(target, pname uint32, params *float32) {
	gl.TexParameterfv(target, pname, params)
}

func (Texture) TexParameteriv(target, pname uint32, params *int32) {
	gl.TexParameteriv(target, pname, params)
}

func (Texture) TexParameterIiv(target, pname uint32, params *int32) {
	gl.TexParameterIiv(target, pname, params)
}

func (Texture) TexParameteri(target, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}

func (Texture) TexParameterIuiv(target, pname uint32, params *uint32) {
	gl.TexParameterIuiv(target, pname, params)
}

func (Texture) GetTexParameterfv(target, pname uint32, params *float32) {
	gl.GetTexParameterfv(target, pname, params)
}

func (Texture) GetTexParameteriv(target, pname uint32, params *int32) {
	gl.GetTexParameteriv(target, pname, params)
}

func (Texture) GetTexParameterIiv(target, pname uint32, params *int32) {
	gl.GetTexParameterIiv(target, pname, params)
}

func (Texture) GetTexParameterIuiv(target, pname uint32, params *uint32) {
	gl.GetTexParameterIuiv(target, pname, params)
}

func (t Texture) IsTexture() bool {
	return gl.IsTexture(uint32(t))
}
