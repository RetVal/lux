package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type getObj struct{}

// Get is a shortcut for all the Get* functions.
var Get = getObj{}

// Booleanv returns a boolean value from OpenGL.
func (getObj) Booleanv(pname uint32, params *bool) {
	gl.GetBooleanv(pname, params)
}

// Doublev returns a float64 value from OpenGL.
func (getObj) Doublev(pname uint32, params *float64) {
	gl.GetDoublev(pname, params)
}

// Floatv returns a float32 value from OpenGL.
func (getObj) Floatv(pname uint32, params *float32) {
	gl.GetFloatv(pname, params)
}

// Integerv returns a int32 value from OpenGL.
func (getObj) Integerv(pname uint32, params *int32) {
	gl.GetIntegerv(pname, params)
}

// Integer64v returns a int64 value from OpenGL.
func (getObj) Integer64v(pname uint32, params *int64) {
	gl.GetInteger64v(pname, params)
}

// Booleani_v returns a boolean value from OpenGL at the specified index.
func (getObj) Booleaniv(pname uint32, index uint32, data *bool) {
	gl.GetBooleani_v(pname, index, data)
}

// Doubleiv returns a float64 value from OpenGL at the specified index.
func (getObj) Doubleiv(pname uint32, params *float64) {
	gl.GetDoublev(pname, params)
}

// Floativ returns a float32 value from OpenGL at the specified index.
func (getObj) Floativ(pname uint32, params *float32) {
	gl.GetFloatv(pname, params)
}

// Integeri_v returns a int32 value from OpenGL at the specified index.
func (getObj) Integeriv(pname uint32, index uint32, data *int32) {
	gl.GetIntegeri_v(pname, index, data)
}

// Integer64i_v returns a int64 value from OpenGL at the specified index.
func (getObj) Integer64iv(pname uint32, index uint32, data *int64) {
	gl.GetInteger64i_v(pname, index, data)
}

//params returns a single value indicating the active multitexture unit. The initial value is GL_TEXTURE0. See glActiveTexture.
func (getObj) ActiveTexture() int32 {
	var params int32
	gl.GetIntegerv(gl.ACTIVE_TEXTURE, &params)
	return params
}

//params returns a pair of values indicating the range of widths supported for aliased lines. See glLineWidth.
func (getObj) AliasedLineWidthRange() [2]float32 {
	var params [2]float32
	gl.GetFloatv(gl.ALIASED_LINE_WIDTH_RANGE, &params[0])
	return params
}

//params returns a pair of values indicating the range of widths supported for smooth (antialiased) lines. See glLineWidth.
func (getObj) SmoothLineWidthRange() [2]float32 {
	var params [2]float32
	gl.GetFloatv(gl.SMOOTH_LINE_WIDTH_RANGE, &params[0])
	return params
}

//params returns a single value indicating the level of quantization applied to smooth line width parameters.
func (getObj) SmoothLineWidthGranularity() float32 {
	var params float32
	gl.GetFloatv(gl.SMOOTH_LINE_WIDTH_GRANULARITY, &params)
	return params
}

//params returns a single value, the name of the buffer object currently bound to the target GL_ARRAY_BUFFER. If no buffer object is bound to this target, 0 is returned. The initial value is 0. See glBindBuffer.
func (getObj) ArrayBufferBinding() Buffer {
	var params int32
	gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &params)
	return Buffer(params)
}

//params returns a single boolean value indicating whether blending is enabled. The initial value is GL_FALSE. See glBlendFunc.
func (getObj) Blend() bool {
	var params bool
	gl.GetBooleanv(gl.BLEND, &params)
	return params
}

//params returns four values, the red, green, blue, and alpha values which are the components of the blend color. See glBlendColor.
func (getObj) BlendColor() [4]float32 {
	var params [4]float32
	gl.GetFloatv(gl.BLEND_COLOR, &params[0])
	return params
}

//params returns one value, the symbolic constant identifying the alpha destination blend function. The initial value is GL_ZERO. See glBlendFunc and glBlendFuncSeparate.
func (getObj) BlendDstAlpha() int32 {
	var params int32
	gl.GetIntegerv(gl.BLEND_DST_ALPHA, &params)
	return params
}

//params returns one value, the symbolic constant identifying the RGB destination blend function. The initial value is GL_ZERO. See glBlendFunc and glBlendFuncSeparate.
func (getObj) BlendDstRgb() int32 {
	var params int32
	gl.GetIntegerv(gl.BLEND_DST_RGB, &params)
	return params
}

//params returns one value, a symbolic constant indicating whether the RGB blend equation is GL_FUNC_ADD, GL_FUNC_SUBTRACT, GL_FUNC_REVERSE_SUBTRACT, GL_MIN or GL_MAX. See glBlendEquationSeparate.
func (getObj) BlendEquationRgb() int32 {
	var params int32
	gl.GetIntegerv(gl.BLEND_EQUATION_RGB, &params)
	return params
}

//params returns one value, a symbolic constant indicating whether the Alpha blend equation is GL_FUNC_ADD, GL_FUNC_SUBTRACT, GL_FUNC_REVERSE_SUBTRACT, GL_MIN or GL_MAX. See glBlendEquationSeparate.
func (getObj) BlendEquationAlpha() int32 {
	var params int32
	gl.GetIntegerv(gl.BLEND_EQUATION_ALPHA, &params)
	return params
}

//params returns one value, the symbolic constant identifying the alpha source blend function. The initial value is GL_ONE. See glBlendFunc and glBlendFuncSeparate.
func (getObj) BlendSrcAlpha() int32 {
	var params int32
	gl.GetIntegerv(gl.BLEND_SRC_ALPHA, &params)
	return params
}

//params returns one value, the symbolic constant identifying the RGB source blend function. The initial value is GL_ONE. See glBlendFunc and glBlendFuncSeparate.
func (getObj) BlendSrcRgb() int32 {
	var params int32
	gl.GetIntegerv(gl.BLEND_SRC_RGB, &params)
	return params
}

//params returns four values: the red, green, blue, and alpha values used to clear the color buffers. Integer values, if requested, are linearly mapped from the internal floating-point representation such that 1.0 returns the most positive representable integer value, and -1.0 returns the most negative representable integer value. The initial value is (0, 0, 0, 0). See glClearColor.
func (getObj) ColorClearValue() [4]float32 {
	var params [4]float32
	gl.GetFloatv(gl.COLOR_CLEAR_VALUE, &params[0])
	return params
}

//params returns a single boolean value indicating whether a fragment's RGBA color values are merged into the framebuffer using a logical operation. The initial value is GL_FALSE. See glLogicOp.
func (getObj) ColorLogicOp() bool {
	var params bool
	gl.GetBooleanv(gl.COLOR_LOGIC_OP, &params)
	return params
}

//params returns four boolean values: the red, green, blue, and alpha write enables for the color buffers. The initial value is (GL_TRUE, GL_TRUE, GL_TRUE, GL_TRUE). See glColorMask.
func (getObj) ColorWritemask() bool {
	var params bool
	gl.GetBooleanv(gl.COLOR_WRITEMASK, &params)
	return params
}

//params returns a list of symbolic constants of length GL_NUM_COMPRESSED_TEXTURE_FORMATS indicating which compressed texture formats are available. See glCompressedTexImage2D.
func (getObj) CompressedTextureFormats() []int32 {
	var params = make([]int32, Get.NumCompressedTextureFormats())
	gl.GetIntegerv(gl.COMPRESSED_TEXTURE_FORMATS, &params[0])
	return params
}

//params returns a single boolean value indicating whether polygon culling is enabled. The initial value is GL_FALSE. See glCullFace.
func (getObj) CullFace() bool {
	var params bool
	gl.GetBooleanv(gl.CULL_FACE, &params)
	return params
}

//params returns one value, the name of the program object that is currently active, or 0 if no program object is active. See glUseProgram.
func (getObj) CurrentProgram() Program {
	var params int32
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &params)
	return Program(params)
}

//params returns one value, the value that is used to clear the depth buffer. Integer values, if requested, are linearly mapped from the internal floating-point representation such that 1.0 returns the most positive representable integer value, and -1.0 returns the most negative representable integer value. The initial value is 1. See glClearDepth.
func (getObj) DepthClearValue() float32 {
	var params float32
	gl.GetFloatv(gl.DEPTH_CLEAR_VALUE, &params)
	return params
}

//params returns one value, the symbolic constant that indicates the depth comparison function. The initial value is GL_LESS. See glDepthFunc.
func (getObj) DepthFunc() int32 {
	var params int32
	gl.GetIntegerv(gl.DEPTH_FUNC, &params)
	return params
}

//params returns two values: the near and far mapping limits for the depth buffer. Integer values, if requested, are linearly mapped from the internal floating-point representation such that 1.0 returns the most positive representable integer value, and -1.0 returns the most negative representable integer value. The initial value is (0, 1). See glDepthRange.
func (getObj) DepthRange() [2]float32 {
	var params [2]float32
	gl.GetFloatv(gl.DEPTH_RANGE, &params[0])
	return params
}

//params returns a single boolean value indicating whether depth testing of fragments is enabled. The initial value is GL_FALSE. See glDepthFunc and glDepthRange.
func (getObj) DepthTest() bool {
	var params bool
	gl.GetBooleanv(gl.DEPTH_TEST, &params)
	return params
}

//params returns a single boolean value indicating if the depth buffer is enabled for writing. The initial value is GL_TRUE. See glDepthMask.
func (getObj) DepthWritemask() bool {
	var params bool
	gl.GetBooleanv(gl.DEPTH_WRITEMASK, &params)
	return params
}

//params returns a single boolean value indicating whether dithering of fragment colors and indices is enabled. The initial value is GL_TRUE.
func (getObj) Dither() bool {
	var params bool
	gl.GetBooleanv(gl.DITHER, &params)
	return params
}

//params returns a single boolean value indicating whether double buffering is supported.
func (getObj) Doublebuffer() bool {
	var params bool
	gl.GetBooleanv(gl.DOUBLEBUFFER, &params)
	return params
}

//params returns one value, a symbolic constant indicating which buffers are being drawn to. See glDrawBuffer. The initial value is GL_BACK if there are back buffers, otherwise it is GL_FRONT.
func (getObj) DrawBuffer() int32 {
	var params int32
	gl.GetIntegerv(gl.DRAW_BUFFER, &params)
	return params
}

//params returns one value, the name of the framebuffer object currently bound to the GL_DRAW_FRAMEBUFFER target. If the default framebuffer is bound, this value will be zero. The initial value is zero. See glBindFramebuffer.
func (getObj) DrawFramebufferBinding() Framebuffer {
	var params int32
	gl.GetIntegerv(gl.DRAW_FRAMEBUFFER_BINDING, &params)
	return Framebuffer(params)
}

//params returns one value, the name of the framebuffer object currently bound to the GL_READ_FRAMEBUFFER target. If the default framebuffer is bound, this value will be zero. The initial value is zero. See glBindFramebuffer.
func (getObj) ReadFramebufferBinding() Framebuffer {
	var params int32
	gl.GetIntegerv(gl.READ_FRAMEBUFFER_BINDING, &params)
	return Framebuffer(params)
}

//params returns a single value, the name of the buffer object currently bound to the target GL_ELEMENT_ARRAY_BUFFER. If no buffer object is bound to this target, 0 is returned. The initial value is 0. See glBindBuffer.
func (getObj) ElementArrayBufferBinding() Buffer {
	var params int32
	gl.GetIntegerv(gl.ELEMENT_ARRAY_BUFFER_BINDING, &params)
	return Buffer(params)
}

//params returns a single value, the name of the renderbuffer object currently bound to the target GL_RENDERBUFFER. If no renderbuffer object is bound to this target, 0 is returned. The initial value is 0. See glBindRenderbuffer.
func (getObj) RenderbufferBinding() RenderBuffer {
	var params int32
	gl.GetIntegerv(gl.RENDERBUFFER_BINDING, &params)
	return RenderBuffer(params)
}

//params returns one value, a symbolic constant indicating the mode of the derivative accuracy hint for fragment shaders. The initial value is GL_DONT_CARE. See glHint.
func (getObj) FragmentShaderDerivativeHint() int32 {
	var params int32
	gl.GetIntegerv(gl.FRAGMENT_SHADER_DERIVATIVE_HINT, &params)
	return params
}

//params returns a single boolean value indicating whether antialiasing of lines is enabled. The initial value is GL_FALSE. See glLineWidth.
func (getObj) LineSmooth() bool {
	var params bool
	gl.GetBooleanv(gl.LINE_SMOOTH, &params)
	return params
}

//params returns one value, a symbolic constant indicating the mode of the line antialiasing hint. The initial value is GL_DONT_CARE. See glHint.
func (getObj) LineSmoothHint() int32 {
	var params int32
	gl.GetIntegerv(gl.LINE_SMOOTH_HINT, &params)
	return params
}

//params returns one value, the line width as specified with glLineWidth. The initial value is 1.
func (getObj) LineWidth() float32 {
	var params float32
	gl.GetFloatv(gl.LINE_WIDTH, &params)
	return params
}

//params returns one value, a symbolic constant indicating the selected logic operation mode. The initial value is GL_COPY. See glLogicOp.
func (getObj) LogicOpMode() int32 {
	var params int32
	gl.GetIntegerv(gl.LOGIC_OP_MODE, &params)
	return params
}

//params returns one value, a rough estimate of the largest 3D texture that the GL can handle. The value must be at least 64. Use GL_PROXY_TEXTURE_3D to determine if a texture is too large. See glTexImage3D.
func (getObj) Max3dTextureSize() int64 {
	var params int64
	gl.GetInteger64v(gl.MAX_3D_TEXTURE_SIZE, &params)
	return params
}

//params returns one value, the maximum number of application-defined clipping distances. The value must be at least 8.
func (getObj) MaxClipDistances() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_CLIP_DISTANCES, &params)
	return params
}

//params returns one value, the number of words for fragment shader uniform variables in all uniform blocks (including default). The value must be at least 1. See glUniform.
func (getObj) MaxCombinedFragmentUniformComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, &params)
	return params
}

//params returns one value, the maximum supported texture image units that can be used to access texture maps from the vertex shader and the fragment processor combined. If both the vertex shader and the fragment processing stage access the same texture image unit, then that counts as using two texture image units against this limit. The value must be at least 48. See glActiveTexture.
func (getObj) MaxCombinedTextureImageUnits() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS, &params)
	return params
}

//params returns one value, the number of words for vertex shader uniform variables in all uniform blocks (including default). The value must be at least 1. See glUniform.
func (getObj) MaxCombinedVertexUniformComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, &params)
	return params
}

//params returns one value, the number of words for geometry shader uniform variables in all uniform blocks (including default). The value must be at least 1. See glUniform.
func (getObj) MaxCombinedGeometryUniformComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS, &params)
	return params
}

//params returns one value, the number components for varying variables, which must be at least 60.
func (getObj) MaxVaryingComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_VARYING_COMPONENTS, &params)
	return params
}

//params returns one value, the maximum number of uniform blocks per program. The value must be at least 36. See glUniformBlockBinding.
func (getObj) MaxCombinedUniformBlocks() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_COMBINED_UNIFORM_BLOCKS, &params)
	return params
}

//params returns one value. The value gives a rough estimate of the largest cube-map texture that the GL can handle. The value must be at least 1024. Use GL_PROXY_TEXTURE_CUBE_MAP to determine if a texture is too large. See glTexImage2D.
func (getObj) MaxCubeMapTextureSize() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_CUBE_MAP_TEXTURE_SIZE, &params)
	return params
}

//params returns one value, the maximum number of simultaneous outputs that may be written in a fragment shader. The value must be at least 8. See glDrawBuffers.
func (getObj) MaxDrawBuffers() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_DRAW_BUFFERS, &params)
	return params
}

//params returns one value, the maximum number of active draw buffers when using dual-source blending. The value must be at least 1. See glBlendFunc and glBlendFuncSeparate.
func (getObj) MaxDualSourceDrawBuffers() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_DUAL_SOURCE_DRAW_BUFFERS, &params)
	return params
}

//params returns one value, the recommended maximum number of vertex array indices. See glDrawRangeElements.
func (getObj) MaxElementsIndices() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_ELEMENTS_INDICES, &params)
	return params
}

//params returns one value, the recommended maximum number of vertex array vertices. See glDrawRangeElements.
func (getObj) MaxElementsVertices() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_ELEMENTS_VERTICES, &params)
	return params
}

//params returns one value, the maximum number of individual floating-point, integer, or boolean values that can be held in uniform variable storage for a fragment shader. The value must be at least 1024. See glUniform.
func (getObj) MaxFragmentUniformComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_FRAGMENT_UNIFORM_COMPONENTS, &params)
	return params
}

//params returns one value, the maximum number of uniform blocks per fragment shader. The value must be at least 12. See glUniformBlockBinding.
func (getObj) MaxFragmentUniformBlocks() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_FRAGMENT_UNIFORM_BLOCKS, &params)
	return params
}

//params returns one value, the maximum number of components of the inputs read by the fragment shader, which must be at least 128.
func (getObj) MaxFragmentInputComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_FRAGMENT_INPUT_COMPONENTS, &params)
	return params
}

//params returns one value, the minimum texel offset allowed in a texture lookup, which must be at most -8.
func (getObj) MinProgramTexelOffset() int32 {
	var params int32
	gl.GetIntegerv(gl.MIN_PROGRAM_TEXEL_OFFSET, &params)
	return params
}

//params returns one value, the maximum texel offset allowed in a texture lookup, which must be at least 7.
func (getObj) MaxProgramTexelOffset() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_PROGRAM_TEXEL_OFFSET, &params)
	return params
}

//params returns one value. The value gives a rough estimate of the largest rectangular texture that the GL can handle. The value must be at least 1024. Use GL_PROXY_TEXTURE_RECTANGLE to determine if a texture is too large. See glTexImage2D.
func (getObj) MaxRectangleTextureSize() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_RECTANGLE_TEXTURE_SIZE, &params)
	return params
}

//params returns one value, the maximum supported texture image units that can be used to access texture maps from the fragment shader. The value must be at least 16. See glActiveTexture.
func (getObj) MaxTextureImageUnits() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_TEXTURE_IMAGE_UNITS, &params)
	return params
}

//params returns one value, the maximum, absolute value of the texture level-of-detail bias. The value must be at least 2.0.
func (getObj) MaxTextureLodBias() float32 {
	var params float32
	gl.GetFloatv(gl.MAX_TEXTURE_LOD_BIAS, &params)
	return params
}

//params returns one value. The value gives a rough estimate of the largest texture that the GL can handle. The value must be at least 1024. Use a proxy texture target such as GL_PROXY_TEXTURE_1D or GL_PROXY_TEXTURE_2D to determine if a texture is too large. See glTexImage1D and glTexImage2D.
func (getObj) MaxTextureSize() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &params)
	return params
}

//params returns one value. The value indicates the maximum supported size for renderbuffers. See glFramebufferRenderbuffer.
func (getObj) MaxRenderbufferSize() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_RENDERBUFFER_SIZE, &params)
	return params
}

//params returns one value. The value indicates the maximum number of layers allowed in an array texture, and must be at least 256. See glTexImage2D.
func (getObj) MaxArrayTextureLayers() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_ARRAY_TEXTURE_LAYERS, &params)
	return params
}

//params returns one value. The value gives the maximum number of texels allowed in the texel array of a texture buffer object. Value must be at least 65536.
func (getObj) MaxTextureBufferSize() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_TEXTURE_BUFFER_SIZE, &params)
	return params
}

//params returns one value, the maximum size in basic machine units of a uniform block. The value must be at least 16384. See glUniformBlockBinding.
func (getObj) MaxUniformBlockSize() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_UNIFORM_BLOCK_SIZE, &params)
	return params
}

//params returns one value, the maximum number of interpolators available for processing varying variables used by vertex and fragment shaders. This value represents the number of individual floating-point values that can be interpolated; varying variables declared as vectors, matrices, and arrays will all consume multiple interpolators. The value must be at least 32.
func (getObj) MaxVaryingFloats() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_VARYING_FLOATS, &params)
	return params
}

//params returns one value, the maximum number of 4-component generic vertex attributes accessible to a vertex shader. The value must be at least 16. See glVertexAttrib.
func (getObj) MaxVertexAttribs() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &params)
	return params
}

//params returns one value, the maximum supported texture image units that can be used to access texture maps from the vertex shader. The value may be at least 16. See glActiveTexture.
func (getObj) MaxVertexTextureImageUnits() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_VERTEX_TEXTURE_IMAGE_UNITS, &params)
	return params
}

//params returns one value, the maximum supported texture image units that can be used to access texture maps from the geometry shader. The value must be at least 16. See glActiveTexture.
func (getObj) MaxGeometryTextureImageUnits() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_GEOMETRY_TEXTURE_IMAGE_UNITS, &params)
	return params
}

//params returns one value, the maximum number of individual floating-point, integer, or boolean values that can be held in uniform variable storage for a vertex shader. The value must be at least 1024. See glUniform.
func (getObj) MaxVertexUniformComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_VERTEX_UNIFORM_COMPONENTS, &params)
	return params
}

//params returns one value, the maximum number of components of output written by a vertex shader, which must be at least 64.
func (getObj) MaxVertexOutputComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_VERTEX_OUTPUT_COMPONENTS, &params)
	return params
}

//params returns one value, the maximum number of individual floating-point, integer, or boolean values that can be held in uniform variable storage for a geometry shader. The value must be at least 1024. See glUniform.
func (getObj) MaxGeometryUniformComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_GEOMETRY_UNIFORM_COMPONENTS, &params)
	return params
}

//params returns one value, the maximum number of sample mask words.
func (getObj) MaxSampleMaskWords() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_SAMPLE_MASK_WORDS, &params)
	return params
}

//params returns one value, the maximum number of samples in a color multisample texture.
func (getObj) MaxColorTextureSamples() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_COLOR_TEXTURE_SAMPLES, &params)
	return params
}

//params returns one value, the maximum number of samples in a multisample depth or depth-stencil texture.
func (getObj) MaxDepthTextureSamples() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_DEPTH_TEXTURE_SAMPLES, &params)
	return params
}

//params returns one value, the maximum number of samples supported in integer format multisample buffers.
func (getObj) MaxIntegerSamples() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_INTEGER_SAMPLES, &params)
	return params
}

//params returns one value, the maximum glWaitSync timeout interval.
func (getObj) MaxServerWaitTimeout() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_SERVER_WAIT_TIMEOUT, &params)
	return params
}

//params returns one value, the maximum number of uniform buffer binding points on the context, which must be at least 36.
func (getObj) MaxUniformBufferBindings() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_UNIFORM_BUFFER_BINDINGS, &params)
	return params
}

//params returns one value, the minimum required alignment for uniform buffer sizes and offsets.
func (getObj) UniformBufferOffsetAlignment() int32 {
	var params int32
	gl.GetIntegerv(gl.UNIFORM_BUFFER_OFFSET_ALIGNMENT, &params)
	return params
}

//params returns one value, the maximum number of uniform blocks per vertex shader. The value must be at least 12. See glUniformBlockBinding.
func (getObj) MaxVertexUniformBlocks() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_VERTEX_UNIFORM_BLOCKS, &params)
	return params
}

//params returns one value, the maximum number of uniform blocks per geometry shader. The value must be at least 12. See glUniformBlockBinding.
func (getObj) MaxGeometryUniformBlocks() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_GEOMETRY_UNIFORM_BLOCKS, &params)
	return params
}

//params returns one value, the maximum number of components of inputs read by a geometry shader, which must be at least 64.
func (getObj) MaxGeometryInputComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_GEOMETRY_INPUT_COMPONENTS, &params)
	return params
}

//params returns one value, the maximum number of components of outputs written by a geometry shader, which must be at least 128.
func (getObj) MaxGeometryOutputComponents() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_GEOMETRY_OUTPUT_COMPONENTS, &params)
	return params
}

//params returns two values: the maximum supported width and height of the viewport. These must be at least as large as the visible dimensions of the display being rendered to. See glViewport.
func (getObj) MaxViewportDims() int32 {
	var params int32
	gl.GetIntegerv(gl.MAX_VIEWPORT_DIMS, &params)
	return params
}

//params returns a single integer value indicating the number of available compressed texture formats. The minimum value is 4. See glCompressedTexImage2D.
func (getObj) NumCompressedTextureFormats() int32 {
	var params int32
	gl.GetIntegerv(gl.NUM_COMPRESSED_TEXTURE_FORMATS, &params)
	return params
}

//params returns one value, the byte alignment used for writing pixel data to memory. The initial value is 4. See glPixelStore.
func (getObj) PackAlignment() int32 {
	var params int32
	gl.GetIntegerv(gl.PACK_ALIGNMENT, &params)
	return params
}

//params returns one value, the image height used for writing pixel data to memory. The initial value is 0. See glPixelStore.
func (getObj) PackImageHeight() int32 {
	var params int32
	gl.GetIntegerv(gl.PACK_IMAGE_HEIGHT, &params)
	return params
}

//params returns a single boolean value indicating whether single-bit pixels being written to memory are written first to the least significant bit of each unsigned byte. The initial value is GL_FALSE. See glPixelStore.
func (getObj) PackLsbFirst() bool {
	var params bool
	gl.GetBooleanv(gl.PACK_LSB_FIRST, &params)
	return params
}

//params returns one value, the row length used for writing pixel data to memory. The initial value is 0. See glPixelStore.
func (getObj) PackRowLength() int32 {
	var params int32
	gl.GetIntegerv(gl.PACK_ROW_LENGTH, &params)
	return params
}

//params returns one value, the number of pixel images skipped before the first pixel is written into memory. The initial value is 0. See glPixelStore.
func (getObj) PackSkipImages() int32 {
	var params int32
	gl.GetIntegerv(gl.PACK_SKIP_IMAGES, &params)
	return params
}

//params returns one value, the number of pixel locations skipped before the first pixel is written into memory. The initial value is 0. See glPixelStore.
func (getObj) PackSkipPixels() int32 {
	var params int32
	gl.GetIntegerv(gl.PACK_SKIP_PIXELS, &params)
	return params
}

//params returns one value, the number of rows of pixel locations skipped before the first pixel is written into memory. The initial value is 0. See glPixelStore.
func (getObj) PackSkipRows() int32 {
	var params int32
	gl.GetIntegerv(gl.PACK_SKIP_ROWS, &params)
	return params
}

//params returns a single boolean value indicating whether the bytes of two-byte and four-byte pixel indices and components are swapped before being written to memory. The initial value is GL_FALSE. See glPixelStore.
func (getObj) PackSwapBytes() bool {
	var params bool
	gl.GetBooleanv(gl.PACK_SWAP_BYTES, &params)
	return params
}

//params returns a single value, the name of the buffer object currently bound to the target GL_PIXEL_PACK_BUFFER. If no buffer object is bound to this target, 0 is returned. The initial value is 0. See glBindBuffer.
func (getObj) PixelPackBufferBinding() Buffer {
	var params int32
	gl.GetIntegerv(gl.PIXEL_PACK_BUFFER_BINDING, &params)
	return Buffer(params)
}

//params returns a single value, the name of the buffer object currently bound to the target GL_PIXEL_UNPACK_BUFFER. If no buffer object is bound to this target, 0 is returned. The initial value is 0. See glBindBuffer.
func (getObj) PixelUnpackBufferBinding() Buffer {
	var params int32
	gl.GetIntegerv(gl.PIXEL_UNPACK_BUFFER_BINDING, &params)
	return Buffer(params)
}

//params returns one value, the point size threshold for determining the point size. See glPointParameter.
func (getObj) PointFadeThresholdSize() float32 {
	var params float32
	gl.GetFloatv(gl.POINT_FADE_THRESHOLD_SIZE, &params)
	return params
}

//params returns one value, the current primitive restart index. The initial value is 0. See glPrimitiveRestartIndex.
func (getObj) PrimitiveRestartIndex() int32 {
	var params int32
	gl.GetIntegerv(gl.PRIMITIVE_RESTART_INDEX, &params)
	return params
}

//params returns a single boolean value indicating whether vertex program point size mode is enabled. If enabled, then the point size is taken from the shader built-in gl_PointSize. If disabled, then the point size is taken from the point state as specified by glPointSize. The initial value is GL_FALSE.
func (getObj) ProgramPointSize() bool {
	var params bool
	gl.GetBooleanv(gl.PROGRAM_POINT_SIZE, &params)
	return params
}

//params returns one value, the currently selected provoking vertex convention. The initial value is GL_LAST_VERTEX_CONVENTION. See glProvokingVertex.
func (getObj) ProvokingVertex() int32 {
	var params int32
	gl.GetIntegerv(gl.PROVOKING_VERTEX, &params)
	return params
}

//params returns one value, the point size as specified by glPointSize. The initial value is 1.
func (getObj) PointSize() float32 {
	var params float32
	gl.GetFloatv(gl.POINT_SIZE, &params)
	return params
}

//params returns one value, the size difference between adjacent supported sizes for antialiased points. See glPointSize.
func (getObj) PointSizeGranularity() float32 {
	var params float32
	gl.GetFloatv(gl.POINT_SIZE_GRANULARITY, &params)
	return params
}

//params returns two values: the smallest and largest supported sizes for antialiased points. The smallest size must be at most 1, and the largest size must be at least 1. See glPointSize.
func (getObj) PointSizeRange() [2]float32 {
	var params [2]float32
	gl.GetFloatv(gl.POINT_SIZE_RANGE, &params[0])
	return params
}

//params returns one value, the scaling factor used to determine the variable offset that is added to the depth value of each fragment generated when a polygon is rasterized. The initial value is 0. See glPolygonOffset.
func (getObj) PolygonOffsetFactor() float32 {
	var params float32
	gl.GetFloatv(gl.POLYGON_OFFSET_FACTOR, &params)
	return params
}

//params returns one value. This value is multiplied by an implementation-specific value and then added to the depth value of each fragment generated when a polygon is rasterized. The initial value is 0. See glPolygonOffset.
func (getObj) PolygonOffsetUnits() float32 {
	var params float32
	gl.GetFloatv(gl.POLYGON_OFFSET_UNITS, &params)
	return params
}

//params returns a single boolean value indicating whether polygon offset is enabled for polygons in fill mode. The initial value is GL_FALSE. See glPolygonOffset.
func (getObj) PolygonOffsetFill() bool {
	var params bool
	gl.GetBooleanv(gl.POLYGON_OFFSET_FILL, &params)
	return params
}

//params returns a single boolean value indicating whether polygon offset is enabled for polygons in line mode. The initial value is GL_FALSE. See glPolygonOffset.
func (getObj) PolygonOffsetLine() bool {
	var params bool
	gl.GetBooleanv(gl.POLYGON_OFFSET_LINE, &params)
	return params
}

//params returns a single boolean value indicating whether polygon offset is enabled for polygons in point mode. The initial value is GL_FALSE. See glPolygonOffset.
func (getObj) PolygonOffsetPoint() bool {
	var params bool
	gl.GetBooleanv(gl.POLYGON_OFFSET_POINT, &params)
	return params
}

//params returns a single boolean value indicating whether antialiasing of polygons is enabled. The initial value is GL_FALSE. See glPolygonMode.
func (getObj) PolygonSmooth() bool {
	var params bool
	gl.GetBooleanv(gl.POLYGON_SMOOTH, &params)
	return params
}

//params returns one value, a symbolic constant indicating the mode of the polygon antialiasing hint. The initial value is GL_DONT_CARE. See glHint.
func (getObj) PolygonSmoothHint() int32 {
	var params int32
	gl.GetIntegerv(gl.POLYGON_SMOOTH_HINT, &params)
	return params
}

//params returns one value, a symbolic constant indicating which color buffer is selected for reading. The initial value is GL_BACK if there is a back buffer, otherwise it is GL_FRONT. See glReadPixels.
func (getObj) ReadBuffer() int32 {
	var params int32
	gl.GetIntegerv(gl.READ_BUFFER, &params)
	return params
}

//params returns a single integer value indicating the number of sample buffers associated with the framebuffer. See glSampleCoverage.
func (getObj) SampleBuffers() int32 {
	var params int32
	gl.GetIntegerv(gl.SAMPLE_BUFFERS, &params)
	return params
}

//params returns a single positive floating-point value indicating the current sample coverage value. See glSampleCoverage.
func (getObj) SampleCoverageValue() float32 {
	var params float32
	gl.GetFloatv(gl.SAMPLE_COVERAGE_VALUE, &params)
	return params
}

//params returns a single boolean value indicating if the temporary coverage value should be inverted. See glSampleCoverage.
func (getObj) SampleCoverageInvert() bool {
	var params bool
	gl.GetBooleanv(gl.SAMPLE_COVERAGE_INVERT, &params)
	return params
}

//params returns a single value, the name of the sampler object currently bound to the active texture unit. The initial value is 0. See glBindSampler.
func (getObj) SamplerBinding() int32 {
	var params int32
	gl.GetIntegerv(gl.SAMPLER_BINDING, &params)
	return params
}

//params returns a single integer value indicating the coverage mask size. See glSampleCoverage.
func (getObj) Samples() int32 {
	var params int32
	gl.GetIntegerv(gl.SAMPLES, &params)
	return params
}

//params returns four values: the x and y window coordinates of the scissor box, followed by its width and height. Initially the x and y window coordinates are both 0 and the width and height are set to the size of the window. See glScissor.
func (getObj) ScissorBox() [4]int32 {
	var params [4]int32
	gl.GetIntegerv(gl.SCISSOR_BOX, &params[0])
	return params
}

//params returns a single boolean value indicating whether scissoring is enabled. The initial value is GL_FALSE. See glScissor.
func (getObj) ScissorTest() bool {
	var params bool
	gl.GetBooleanv(gl.SCISSOR_TEST, &params)
	return params
}

//params returns one value, a symbolic constant indicating what action is taken for back-facing polygons when the stencil test fails. The initial value is GL_KEEP. See glStencilOpSeparate.
func (getObj) StencilBackFail() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_BACK_FAIL, &params)
	return params
}

//params returns one value, a symbolic constant indicating what function is used for back-facing polygons to compare the stencil reference value with the stencil buffer value. The initial value is GL_ALWAYS. See glStencilFuncSeparate.
func (getObj) StencilBackFunc() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_BACK_FUNC, &params)
	return params
}

//params returns one value, a symbolic constant indicating what action is taken for back-facing polygons when the stencil test passes, but the depth test fails. The initial value is GL_KEEP. See glStencilOpSeparate.
func (getObj) StencilBackPassDepthFail() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_BACK_PASS_DEPTH_FAIL, &params)
	return params
}

//params returns one value, a symbolic constant indicating what action is taken for back-facing polygons when the stencil test passes and the depth test passes. The initial value is GL_KEEP. See glStencilOpSeparate.
func (getObj) StencilBackPassDepthPass() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_BACK_PASS_DEPTH_PASS, &params)
	return params
}

//params returns one value, the reference value that is compared with the contents of the stencil buffer for back-facing polygons. The initial value is 0. See glStencilFuncSeparate.
func (getObj) StencilBackRef() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_BACK_REF, &params)
	return params
}

//params returns one value, the mask that is used for back-facing polygons to mask both the stencil reference value and the stencil buffer value before they are compared. The initial value is all 1's. See glStencilFuncSeparate.
func (getObj) StencilBackValueMask() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_BACK_VALUE_MASK, &params)
	return params
}

//params returns one value, the mask that controls writing of the stencil bitplanes for back-facing polygons. The initial value is all 1's. See glStencilMaskSeparate.
func (getObj) StencilBackWritemask() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_BACK_WRITEMASK, &params)
	return params
}

//params returns one value, the index to which the stencil bitplanes are cleared. The initial value is 0. See glClearStencil.
func (getObj) StencilClearValue() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_CLEAR_VALUE, &params)
	return params
}

//params returns one value, a symbolic constant indicating what action is taken when the stencil test fails. The initial value is GL_KEEP. See glStencilOp. This stencil state only affects non-polygons and front-facing polygons. Back-facing polygons use separate stencil state. See glStencilOpSeparate.
func (getObj) StencilFail() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_FAIL, &params)
	return params
}

//params returns one value, a symbolic constant indicating what function is used to compare the stencil reference value with the stencil buffer value. The initial value is GL_ALWAYS. See glStencilFunc. This stencil state only affects non-polygons and front-facing polygons. Back-facing polygons use separate stencil state. See glStencilFuncSeparate.
func (getObj) StencilFunc() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_FUNC, &params)
	return params
}

//params returns one value, a symbolic constant indicating what action is taken when the stencil test passes, but the depth test fails. The initial value is GL_KEEP. See glStencilOp. This stencil state only affects non-polygons and front-facing polygons. Back-facing polygons use separate stencil state. See glStencilOpSeparate.
func (getObj) StencilPassDepthFail() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_PASS_DEPTH_FAIL, &params)
	return params
}

//params returns one value, a symbolic constant indicating what action is taken when the stencil test passes and the depth test passes. The initial value is GL_KEEP. See glStencilOp. This stencil state only affects non-polygons and front-facing polygons. Back-facing polygons use separate stencil state. See glStencilOpSeparate.
func (getObj) StencilPassDepthPass() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_PASS_DEPTH_PASS, &params)
	return params
}

//params returns one value, the reference value that is compared with the contents of the stencil buffer. The initial value is 0. See glStencilFunc. This stencil state only affects non-polygons and front-facing polygons. Back-facing polygons use separate stencil state. See glStencilFuncSeparate.
func (getObj) StencilRef() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_REF, &params)
	return params
}

//params returns a single boolean value indicating whether stencil testing of fragments is enabled. The initial value is GL_FALSE. See glStencilFunc and glStencilOp.
func (getObj) StencilTest() bool {
	var params bool
	gl.GetBooleanv(gl.STENCIL_TEST, &params)
	return params
}

//params returns one value, the mask that is used to mask both the stencil reference value and the stencil buffer value before they are compared. The initial value is all 1's. See glStencilFunc. This stencil state only affects non-polygons and front-facing polygons. Back-facing polygons use separate stencil state. See glStencilFuncSeparate.
func (getObj) StencilValueMask() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_VALUE_MASK, &params)
	return params
}

//params returns one value, the mask that controls writing of the stencil bitplanes. The initial value is all 1's. See glStencilMask. This stencil state only affects non-polygons and front-facing polygons. Back-facing polygons use separate stencil state. See glStencilMaskSeparate.
func (getObj) StencilWritemask() int32 {
	var params int32
	gl.GetIntegerv(gl.STENCIL_WRITEMASK, &params)
	return params
}

//params returns a single boolean value indicating whether stereo buffers (left and right) are supported.
func (getObj) Stereo() bool {
	var params bool
	gl.GetBooleanv(gl.STEREO, &params)
	return params
}

//params returns one value, an estimate of the number of bits of subpixel resolution that are used to position rasterized geometry in window coordinates. The value must be at least 4.
func (getObj) SubpixelBits() int32 {
	var params int32
	gl.GetIntegerv(gl.SUBPIXEL_BITS, &params)
	return params
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_1D. The initial value is 0. See glBindTexture.
func (getObj) TextureBinding1D() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_1D, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_1D_ARRAY. The initial value is 0. See glBindTexture.
func (getObj) TextureBinding1DArray() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_1D_ARRAY, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_2D. The initial value is 0. See glBindTexture.
func (getObj) TextureBinding2D() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_2D_ARRAY. The initial value is 0. See glBindTexture.
func (getObj) TextureBinding2DArray() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D_ARRAY, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_2D_MULTISAMPLE. The initial value is 0. See glBindTexture.
func (getObj) TextureBinding2DMultisample() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D_MULTISAMPLE, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_2D_MULTISAMPLE_ARRAY. The initial value is 0. See glBindTexture.
func (getObj) TextureBinding2DMultisampleArray() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_3D. The initial value is 0. See glBindTexture.
func (getObj) TextureBinding3D() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_3D, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_BUFFER. The initial value is 0. See glBindTexture.
func (getObj) TextureBindingBuffer() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_BUFFER, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_CUBE_MAP. The initial value is 0. See glBindTexture.
func (getObj) TextureBindingCubeMap() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_CUBE_MAP, &params)
	return Texture(params)
}

//params returns a single value, the name of the texture currently bound to the target GL_TEXTURE_RECTANGLE. The initial value is 0. See glBindTexture.
func (getObj) TextureBindingRectangle() Texture {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_RECTANGLE, &params)
	return Texture(params)
}

//params returns a single value indicating the mode of the texture compression hint. The initial value is GL_DONT_CARE.
func (getObj) TextureCompressionHint() int32 {
	var params int32
	gl.GetIntegerv(gl.TEXTURE_COMPRESSION_HINT, &params)
	return params
}

//params returns a single value, the 64-bit value of the current GL time. See glQueryCounter.
func (getObj) Timestamp() int64 {
	var params int64
	gl.GetInteger64v(gl.TIMESTAMP, &params)
	return params
}

//When used with non-indexed variants of glGet (such as glGetIntegerv), params returns a single value, the name of the buffer object currently bound to the target GL_TRANSFORM_FEEDBACK_BUFFER. If no buffer object is bound to this target, 0 is returned. When used with indexed variants of glGet (such as glGetIntegeri_v), params returns a single value, the name of the buffer object bound to the indexed transform feedback attribute stream. The initial value is 0 for all targets. See glBindBuffer, glBindBufferBase, and glBindBufferRange.
func (getObj) TransformFeedbackBufferBinding() TransformFeedback {
	var params int32
	gl.GetIntegerv(gl.TRANSFORM_FEEDBACK_BUFFER_BINDING, &params)
	return TransformFeedback(params)
}

//When used with indexed variants of glGet (such as glGetInteger64i_v), params returns a single value, the start offset of the binding range for each transform feedback attribute stream. The initial value is 0 for all streams. See glBindBufferRange.
func (getObj) TransformFeedbackBufferStart(index uint32) int64 {
	var params int64
	gl.GetInteger64i_v(gl.TRANSFORM_FEEDBACK_BUFFER_START, index, &params)
	return params
}

//When used with indexed variants of glGet (such as glGetInteger64i_v), params returns a single value, the size of the binding range for each transform feedback attribute stream. The initial value is 0 for all streams. See glBindBufferRange.
func (getObj) TransformFeedbackBufferSize(index uint32) int64 {
	var params int64
	gl.GetInteger64i_v(gl.TRANSFORM_FEEDBACK_BUFFER_SIZE, index, &params)
	return params
}

//When used with non-indexed variants of glGet (such as glGetIntegerv), params returns a single value, the name of the buffer object currently bound to the target GL_UNIFORM_BUFFER. If no buffer object is bound to this target, 0 is returned. When used with indexed variants of glGet (such as glGetIntegeri_v), params returns a single value, the name of the buffer object bound to the indexed uniform buffer binding point. The initial value is 0 for all targets. See glBindBuffer, glBindBufferBase, and glBindBufferRange.
func (getObj) UniformBufferBinding() Buffer {
	var params int32
	gl.GetIntegerv(gl.UNIFORM_BUFFER_BINDING, &params)
	return Buffer(params)
}

//When used with indexed variants of glGet (such as glGetInteger64i_v), params returns a single value, the start offset of the binding range for each indexed uniform buffer binding. The initial value is 0 for all bindings. See glBindBufferRange.
func (getObj) UniformBufferStart(index uint32) int64 {
	var params int64
	gl.GetInteger64i_v(gl.UNIFORM_BUFFER_START, index, &params)
	return params
}

//When used with indexed variants of glGet (such as glGetInteger64i_v), params returns a single value, the size of the binding range for each indexed uniform buffer binding. The initial value is 0 for all bindings. See glBindBufferRange.
func (getObj) UniformBufferSize(index uint32) int64 {
	var params int64
	gl.GetInteger64i_v(gl.UNIFORM_BUFFER_SIZE, index, &params)
	return params
}

//params returns one value, the byte alignment used for reading pixel data from memory. The initial value is 4. See glPixelStore.
func (getObj) UnpackAlignment() int32 {
	var params int32
	gl.GetIntegerv(gl.UNPACK_ALIGNMENT, &params)
	return params
}

//params returns one value, the image height used for reading pixel data from memory. The initial is 0. See glPixelStore.
func (getObj) UnpackImageHeight() int32 {
	var params int32
	gl.GetIntegerv(gl.UNPACK_IMAGE_HEIGHT, &params)
	return params
}

//params returns a single boolean value indicating whether single-bit pixels being read from memory are read first from the least significant bit of each unsigned byte. The initial value is GL_FALSE. See glPixelStore.
func (getObj) UnpackLsbFirst() bool {
	var params bool
	gl.GetBooleanv(gl.UNPACK_LSB_FIRST, &params)
	return params
}

//params returns one value, the row length used for reading pixel data from memory. The initial value is 0. See glPixelStore.
func (getObj) UnpackRowLength() int32 {
	var params int32
	gl.GetIntegerv(gl.UNPACK_ROW_LENGTH, &params)
	return params
}

//params returns one value, the number of pixel images skipped before the first pixel is read from memory. The initial value is 0. See glPixelStore.
func (getObj) UnpackSkipImages() int32 {
	var params int32
	gl.GetIntegerv(gl.UNPACK_SKIP_IMAGES, &params)
	return params
}

//params returns one value, the number of pixel locations skipped before the first pixel is read from memory. The initial value is 0. See glPixelStore.
func (getObj) UnpackSkipPixels() int32 {
	var params int32
	gl.GetIntegerv(gl.UNPACK_SKIP_PIXELS, &params)
	return params
}

//params returns one value, the number of rows of pixel locations skipped before the first pixel is read from memory. The initial value is 0. See glPixelStore.
func (getObj) UnpackSkipRows() int32 {
	var params int32
	gl.GetIntegerv(gl.UNPACK_SKIP_ROWS, &params)
	return params
}

//params returns a single boolean value indicating whether the bytes of two-byte and four-byte pixel indices and components are swapped after being read from memory. The initial value is GL_FALSE. See glPixelStore.
func (getObj) UnpackSwapBytes() bool {
	var params bool
	gl.GetBooleanv(gl.UNPACK_SWAP_BYTES, &params)
	return params
}

//params returns one value, the number of extensions supported by the GL implementation for the current context. See glGetString.
func (getObj) NumExtensions() int32 {
	var params int32
	gl.GetIntegerv(gl.NUM_EXTENSIONS, &params)
	return params
}

//params returns one value, the major version number of the OpenGL API supported by the current context.
func (getObj) MajorVersion() int32 {
	var params int32
	gl.GetIntegerv(gl.MAJOR_VERSION, &params)
	return params
}

//params returns one value, the minor version number of the OpenGL API supported by the current context.
func (getObj) MinorVersion() int32 {
	var params int32
	gl.GetIntegerv(gl.MINOR_VERSION, &params)
	return params
}

//params returns one value, the flags with which the context was created (such as debugging functionality).
func (getObj) ContextFlags() int32 {
	var params int32
	gl.GetIntegerv(gl.CONTEXT_FLAGS, &params)
	return params
}

//params returns four values: the x and y window coordinates of the viewport, followed by its width and height. Initially the x and y window coordinates are both set to 0, and the width and height are set to the width and height of the window into which the GL will do its rendering. See glViewport.
func (getObj) Viewport() [4]int32 {
	var params [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &params[0])
	return params
}
