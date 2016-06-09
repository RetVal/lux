package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

//Texture2D is the high level representation of OpenGL TEXTURE_2D object. It restrict availlable functions and automatically fills the GL_TEXTURE_2D target.
type Texture2D Texture

//GenTexture2D is an alias to glGenTextures(1,&tex).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenTextures.xml
func GenTexture2D() Texture2D {
	var tex uint32
	gl.GenTextures(1, &tex)
	return Texture2D(tex)
}

//GenTextures2D is an alias to glGenTextures(n,&tex).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenTextures.xml
func GenTextures2D(n int32) []Texture2D {
	tex := make([]Texture2D, n)
	gl.GenTextures(1, (*uint32)(&tex[0]))
	return tex
}

//Bind (Texture2D) is an alias to glBindTexture(gl.TEXTURE_2D, t).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindTexture.xml
func (t Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, uint32(t))
}

//Unbind is an alias to glBindTexture(0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindTexture.xml
func (Texture2D) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

//Delete is an alias to glDeleteTextures. This texture should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteTextures.xml
func (t Texture2D) Delete() {
	gl.DeleteTextures(1, (*uint32)(&t))
}

//TexImage2D is an alias to glTexImage2D.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage2D.xml
func (Texture2D) TexImage2D(level, internalformat, width, height, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(gl.TEXTURE_2D, level, internalformat, width, height, border, format, xtype, pixels)
}

//TexParameteriv is an alias to glTexParameteriv(target, pname, param)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) TexParameteriv(target, pname uint32, param *int32) {
	gl.TexParameteriv(target, pname, param)
}

//TexParameteri is an alias to glTexParameteri.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) TexParameteri(pname uint32, param int32) {
	gl.TexParameteri(gl.TEXTURE_2D, pname, param)
}

func (Texture2D) GetTexParameteriv(pname uint32, params *int32) {
	gl.GetTexParameteriv(gl.TEXTURE_2D, pname, params)
}

//TexParameterf is an alias to glTexParameterf.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) TexParameterf(pname uint32, param float32) {
	gl.TexParameterf(gl.TEXTURE_2D, pname, param)
}

//TexParameterfv is an alias to glTexParameterfv.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) TexParameterfv(pname uint32, param *float32) {
	gl.TexParameterfv(gl.TEXTURE_2D, pname, param)
}

//unc (Texture2D) GetTexLevelParameteriv(target uit32, level int32, pname uint32, params *int32) {GetTexLevelParameteriv is an alias to glGetTexLevelParameteriv(target, level, pname, params)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameter.xml
func (Texture2D) GetTexLevelParameteriv(target uint32, level int32, pname uint32, params *int32) {
	gl.GetTexLevelParameteriv(target, level, pname, params)
}

//unc (Texture2D) GetTexImage(level in32, format, xtype uint32, pixels unsafe.Pointer) {GetTexImage is an alias to glGetTexImage(gl.TEXTURE_2D, level, format, xtype, pixels)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexImage.xml
func (Texture2D) GetTexImage(level int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.GetTexImage(gl.TEXTURE_2D, level, format, xtype, pixels)
}

//ReadPixels is an alias to glReadPixels(x, y, width, height, format, xtype, pixels)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glReadPixels.xml
func (Texture2D) ReadPixels(x, y, width, height int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.ReadPixels(x, y, width, height, format, xtype, pixels)
}

//Width is an alias to glGetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_WIDTH, &w)
//
//Documentation reference:
func (Texture2D) Width(miplevel int32) int32 {
	var w int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_WIDTH, &w)
	return w
}

//Height is an alias to glGetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_HEIGHT, &h)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameter.xml
func (Texture2D) Height(miplevel int32) int32 {
	var h int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_HEIGHT, &h)
	return h
}

//InternalFormat is an alias to glGetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_INTERNAL_FORMAT, &x)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameter.xml
func (t Texture2D) InternalFormat(miplevel int32) uint32 {
	var x int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_INTERNAL_FORMAT, &x)
	return uint32(x)
}

//=========================================================================================================
//=================================================Utility=================================================
//=========================================================================================================

//BaseLevel is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_BASE_LEVEL, level).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) BaseLevel(level int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_BASE_LEVEL, level)
}

//BorderColor is an alias to glTexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_BORDER_COLOR, color ).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) BorderColor(color *float32) {
	gl.TexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_BORDER_COLOR, color)
}

//CompareFunc is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_FUNC, cfunc).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) CompareFunc(cfunc int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_FUNC, cfunc)
}

//CompareMode is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_MODE, mode).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) CompareMode(mode int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_MODE, mode)
}

//LODBias is an alias to glTexParameterf(gl.TEXTURE_2D, gl.TEXTURE_LOD_BIAS, bias).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) LODBias(bias float32) {
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_LOD_BIAS, bias)
}

//MinFilter is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, filter).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) MinFilter(filter int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, filter)
}

//MagFilter is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, filter).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) MagFilter(filter int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, filter)
}

//MinLod is an alias to glTexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_LOD, param).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) MinLod(param float32) {
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_LOD, param)
}

//MaxLod is an alias to glTexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAX_LOD, param).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) MaxLod(param float32) {
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAX_LOD, param)
}

//MaxLevel is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAX_LEVEL, param).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) MaxLevel(param int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAX_LEVEL, param)
}

//SwizzleR is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_R, swizzle).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) SwizzleR(swizzle int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_R, swizzle)
}

//SwizzleG is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_G, swizzle).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) SwizzleG(swizzle int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_G, swizzle)
}

//SwizzleB is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_B, swizzle).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) SwizzleB(swizzle int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_B, swizzle)
}

//SwizzleA is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_A, swizzle).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) SwizzleA(swizzle int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_A, swizzle)
}

//WrapS is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, wrap).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) WrapS(wrap int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, wrap)
}

//WrapT is an alias to glTexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, wrap).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (Texture2D) WrapT(wrap int32) {
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, wrap)
}

//GetBaseLevel is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_BASE_LEVEL, &params)
func (Texture2D) GetBaseLevel() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_BASE_LEVEL, &params)
	return params
}

//GetBorderColor is an alias to glGetTexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_BORDER_COLOR, &params[0])
func (Texture2D) GetBorderColor() [4]float32 {
	var params [4]float32
	gl.GetTexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_BORDER_COLOR, &params[0])
	return params
}

//GetCompareMode is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_MODE, &params)
func (Texture2D) GetCompareMode() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_MODE, &params)
	return params
}

//GetCompareFunc is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_FUNC, &params)
func (Texture2D) GetCompareFunc() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_FUNC, &params)
	return params
}

//GetLODBias is an alias to glGetTexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_LOD_BIAS, &params)
func (Texture2D) GetLODBias() float32 {
	var params float32
	gl.GetTexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_LOD_BIAS, &params)
	return params
}

//GetMagFilter is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, &params)
func (Texture2D) GetMagFilter() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, &params)
	return params
}

//GetMaxLevel is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MAX_LEVEL, &params)
func (Texture2D) GetMaxLevel() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MAX_LEVEL, &params)
	return params
}

//GetMaxLOD is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MAX_LOD, &params)
func (Texture2D) GetMaxLOD() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MAX_LOD, &params)
	return params
}

//GetMinFilter is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, &params)
func (Texture2D) GetMinFilter() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, &params)
	return params
}

//GetMinLOD is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MIN_LOD, &params)
func (Texture2D) GetMinLOD() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MIN_LOD, &params)
	return params
}

//GetSwizzleR is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_R, &params)
func (Texture2D) GetSwizzleR() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_R, &params)
	return params
}

//GetSwizzleG is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_G, &params)
func (Texture2D) GetSwizzleG() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_G, &params)
	return params
}

//GetSwizzleB is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_B, &params)
func (Texture2D) GetSwizzleB() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_B, &params)
	return params
}

//GetSwizzleA is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_A, &params)
func (Texture2D) GetSwizzleA() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_A, &params)
	return params
}

//GetSwizzleRGBA is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_RGBA, &params)
func (Texture2D) GetSwizzleRGBA() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_SWIZZLE_RGBA, &params)
	return params
}

//GetWrapS is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, &params)
func (Texture2D) GetWrapS() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, &params)
	return params
}

//GetWrapT is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, &params)
func (Texture2D) GetWrapT() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, &params)
	return params
}

//GetWrapR is an alias to glGetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, &params)
func (Texture2D) GetWrapR() int32 {
	var params int32
	gl.GetTexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, &params)
	return params
}
