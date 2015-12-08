package lux

import (
	"errors"
	"github.com/go-gl/gl/v3.3-core/gl"
	gl2 "github.com/luxengine/gl"
	"github.com/luxengine/glm"
)

// ShadowFBO is the structure to hold all the resources required to render
// shadow maps.
type ShadowFBO struct {
	framebuffer          gl2.Framebuffer
	texture              gl2.Texture2D
	projection, view, vp glm.Mat4
	program              gl2.Program
	mvpUni               gl2.UniformLocation
	width, height        int32
}

// NewShadowFBO will create a new FBO, a new depth texture and the program
// needed to generate shadow map currently not optimized, we should probably
// reuse the FBO and absolutely reuse the program.
func NewShadowFBO(width, height int32) (*ShadowFBO, error) {
	sfbo := ShadowFBO{}
	fbo := gl2.GenFramebuffer()
	sfbo.width, sfbo.height = width, height
	sfbo.framebuffer = fbo
	fbo.Bind(gl2.FRAMEBUFFER)
	defer fbo.Unbind(gl2.FRAMEBUFFER)

	shadowtex := gl2.GenTexture2D()
	sfbo.texture = shadowtex
	shadowtex.Bind()
	defer shadowtex.Unbind()
	shadowtex.TexParameteri(gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	shadowtex.TexParameteri(gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	shadowtex.TexParameteri(gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	shadowtex.TexParameteri(gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	shadowtex.TexParameteri(gl.TEXTURE_COMPARE_MODE, gl.COMPARE_REF_TO_TEXTURE)
	shadowtex.TexParameteri(gl.TEXTURE_COMPARE_FUNC, gl.LEQUAL)
	shadowtex.TexImage2D(0, gl.DEPTH_COMPONENT16, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)

	fbo.Texture(gl2.FRAMEBUFFER, gl2.DEPTH_ATTACHMENT, gl2.Texture(shadowtex), 0 /*level*/)

	fbo.DrawBuffer(gl2.NONE)
	fbo.ReadBuffer(gl2.NONE)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return &sfbo, errors.New("framebuffer incomplete")
	}

	vs, err := CompileShader(_shadowFboVertexShader, gl2.VERTEX_SHADER)
	if err != nil {
		return &sfbo, err
	}
	defer vs.Delete()
	fs, err := CompileShader(_shadowFboFragmentShader, gl2.FRAGMENT_SHADER)
	if err != nil {
		return &sfbo, err
	}
	defer fs.Delete()
	prog, err := NewProgram(vs, fs)
	if err != nil {
		return &sfbo, err
	}
	sfbo.program = prog
	prog.Use()
	sfbo.mvpUni = prog.GetUniformLocation("mvp")

	return &sfbo, nil
}

// SetOrtho sets the projection matrix to be used.
func (sfbo *ShadowFBO) SetOrtho(left, right, bottom, top, near, far float32) {
	sfbo.projection = glm.Ortho(left, right, bottom, top, near, far)
}

// LookAt sets the view matrix to look at (tx,ty,tz) from (ex,ey,ez), the up
// direction is always (0,1,0).
func (sfbo *ShadowFBO) LookAt(ex, ey, ez, tx, ty, tz float32) {
	sfbo.view = glm.LookAt(ex, ey, ez, tx, ty, tz, 0, 1, 0)
}

// BindForDrawing binds this fbo, change face culling for back face, start using
// the shadow program, calculate projection and clears the texture.
func (sfbo *ShadowFBO) BindForDrawing() {
	sfbo.framebuffer.Bind(gl2.FRAMEBUFFER)
	sfbo.program.Use()
	sfbo.vp = sfbo.projection.Mul4(&sfbo.view)
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	ViewportChange(sfbo.width, sfbo.height)
	gl.CullFace(gl.FRONT)
}

// Unbind return cull face to front and unbind this fbo.
func (sfbo *ShadowFBO) Unbind() {
	sfbo.framebuffer.Unbind(gl2.FRAMEBUFFER)
	gl.CullFace(gl.BACK)
}

// Render takes a mesh and a transform and render them, adding them to the depth
// texture data.
func (sfbo *ShadowFBO) Render(mesh Mesh, transform *Transform) {
	m := transform.Mat4()
	mvpmat := sfbo.vp.Mul4(&m)
	sfbo.mvpUni.UniformMatrix4fv(1, false, &mvpmat[0])
	mesh.Bind()
	mesh.DrawCall()
}

// ShadowMap return the depth texture.
func (sfbo *ShadowFBO) ShadowMap() gl2.Texture2D {
	return sfbo.texture
}

// ShadowMat return the 4x4 matric that represent world-to-screen transform used
// to check pixel occlusion.
func (sfbo *ShadowFBO) ShadowMat() glm.Mat4 {
	return depthscaling.Mul4(&sfbo.vp)
}

// Delete will clean up all the resources allocated to this FBO.
func (sfbo *ShadowFBO) Delete() {
	sfbo.texture.Delete()
	sfbo.framebuffer.Delete()
	sfbo.program.Delete()
}

var _shadowFboVertexShader = `
#version 330 core
uniform mat4 mvp;

layout(location = 0) in vec3 position;
 
void main(){
	gl_Position =  mvp * vec4(position,1);
}
` + "\x00"

var _shadowFboFragmentShader = `
#version 330 core

layout(location=0) out float x;

void main(){
	x=1;
}
` + "\x00"

var depthscaling = glm.Mat4{0.5, 0, 0, 0,
	0, 0.5, 0, 0,
	0, 0, 0.5, 0,
	0.5, 0.5, 0.5, 1}
