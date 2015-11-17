package lux

import (
	"errors"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	gl2 "github.com/luxengine/gl"
)

//the post process fullscreen quad mesh

var fstri struct {
	vao     gl2.VertexArray
	pos, uv gl2.Buffer
}

//InitPostProcessSystem will allocate all the resources required to make the Image Post Processing system work.
func InitPostProcessSystem() {
	//init model
	quadvao := gl2.GenVertexArray()
	quadvao.Bind()
	defer quadvao.Unbind()

	vertpos := gl2.GenBuffer()
	vertpos.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(_quadvertpos)*2*4, ptr(_quadvertpos), gl.STATIC_DRAW)

	vertuv := gl2.GenBuffer()
	vertuv.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(_quadvertuv)*2*4, ptr(_quadvertuv), gl.STATIC_DRAW)

	fstri.vao = quadvao
	fstri.pos = vertpos
	fstri.uv = vertuv
}

//PostProcessFramebufferer is an interface to represent a single PostProcess effect.
type PostProcessFramebufferer interface {
	PreRender()
	Render(gl2.Texture2D)
	SetNext(PostProcessFramebufferer)
}

//PostProcessFramebuffer is the generic post process framebuffer
type PostProcessFramebuffer struct {
	Fb           gl2.Framebuffer
	Tex          gl2.Texture2D
	Prog         gl2.Program
	next         PostProcessFramebufferer
	time, source gl2.UniformLocation
}

//NewPostProcessFramebuffer creates a new PostProcessFramebuffer and allocated all the ressources.
//You do not control the vertex shader but you can give a fragment shader. The fragment shader must have the following uniforms:
//	-resolution: float vec2, representing the size of the texture
//	-time: float, glfw time since the begining of the program
//	-tex: sampler2D, the input texture to this post process pass
func NewPostProcessFramebuffer(width, height int32, fragmentSource string) (*PostProcessFramebuffer, error) {
	ppf := PostProcessFramebuffer{}
	vs, err := CompileShader(_fullscreenVertexShader, gl2.VERTEX_SHADER)
	if err != nil {
		return &ppf, err
	}
	defer vs.Delete()
	fs, err := CompileShader(fragmentSource, gl2.FRAGMENT_SHADER)
	if err != nil {
		return &ppf, err
	}
	defer fs.Delete()
	prog, err := NewProgram(vs, fs)
	if err != nil {
		return &ppf, err
	}
	ppf.Prog = prog
	prog.Use()
	res := prog.GetUniformLocation("resolution")
	ppf.source = prog.GetUniformLocation("tex")

	res.Uniform2f(float32(width), float32(height))

	ppf.time = prog.GetUniformLocation("time")

	ppf.Fb = gl2.GenFramebuffer()
	ppf.Tex = GenRGBTexture2D(width, height)

	ppf.Fb.Bind(gl2.FRAMEBUFFER)
	defer ppf.Fb.Unbind(gl2.FRAMEBUFFER)

	ppf.Fb.Texture(gl2.FRAMEBUFFER, gl2.COLOR_ATTACHMENT0, gl2.Texture(ppf.Tex), 0)
	ppf.Fb.DrawBuffers(gl2.COLOR_ATTACHMENT0)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return &ppf, errors.New("framebuffer incomplete")
	}
	return &ppf, nil
}

//PreRender binds either the next post process fbo if there is one or unbinds any fbo to render to screen. Also disable depth test.
func (ppfb *PostProcessFramebuffer) PreRender() {
	if ppfb.next != nil {
		ppfb.Fb.Bind(gl2.FRAMEBUFFER)
	} else {
		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	}
	gl.Disable(gl.DEPTH_TEST)
}

//PostRender renable depth test
func (ppfb *PostProcessFramebuffer) PostRender() {
	gl.Enable(gl.DEPTH_TEST)
}

//Render takes a texture and feed it to the fragment shader as a fullscreen texture. It will call the next post process pass if there is one.
func (ppfb *PostProcessFramebuffer) Render(t gl2.Texture2D) {
	ppfb.Prog.Use()
	ppfb.time.Uniform1f(float32(glfw.GetTime()))

	gl.ActiveTexture(gl2.TEXTURE0)
	t.Bind()
	ppfb.source.Uniform1i(0)
	Fstri()
	if ppfb.next != nil {
		ppfb.next.PreRender()
		ppfb.next.Render(ppfb.Tex)
	}
}

//SetNext sets the post process effect to pass automatically after this post process.
func (ppfb *PostProcessFramebuffer) SetNext(n PostProcessFramebufferer) {
	ppfb.next = n
}

//Delete releases all the resources allocated to this post process fbo.
func (ppfb *PostProcessFramebuffer) Delete() {
	ppfb.Prog.Delete()
	ppfb.Tex.Delete()
	ppfb.Fb.Delete()
}

//Fstri draws a fullscreen triangle such that it covers the entire screen with uv coordinates from [0,0]-[1,1]
func Fstri() {
	fstri.vao.Bind()
	defer fstri.vao.Unbind()
	gl.EnableVertexAttribArray(0)
	fstri.pos.Bind(gl.ARRAY_BUFFER)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	fstri.uv.Bind(gl.ARRAY_BUFFER)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

var _fullscreenVertexShader = `
#version 330
layout (location=0) in vec2 vert;
layout (location=1) in vec2 vertTexCoord;
out vec2 uv;
void main() {
	uv = vertTexCoord;
	gl_Position = vec4(vert,0,1);
}
` + "\x00"

//technicly a fullscreen triangle now
var _quadvertpos = []float32{3, 1, -1, 1, -1, -3}
var _quadvertuv = []float32{2, 1, 0, 1, 0, -1}
