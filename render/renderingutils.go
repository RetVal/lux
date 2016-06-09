package lux

import (
	"io/ioutil"

	"github.com/luxengine/lux/gl"
)

// RenderProgram is the lux representation of a OpenGL program vertex-fragment
// along with all the common uniforms.
type RenderProgram struct {
	Prog                            gl.Program
	M, V, P, Diffuse, Light, N, Eye gl.UniformLocation
}

// LoadProgram loads a vertex-fragment program and gathers:
// "M": model matrix uniform
// "V": view matrix uniform
// "P": projection matrix uniform
// "N": normal matrix uniform
// "diffuse":diffuse texture sampler2d
// "pointlight":array of vec3 for light position
func LoadProgram(vertexfile, fragfile string) (out RenderProgram, err error) {
	vssource, err := ioutil.ReadFile(vertexfile)
	if err != nil {
		return
	}
	fssource, err := ioutil.ReadFile(fragfile)
	if err != nil {
		return
	}
	vs, err := CompileShader(string(vssource)+"\x00", gl.VERTEX_SHADER)
	defer vs.Delete()
	if err != nil {
		return
	}
	fs, err := CompileShader(string(fssource)+"\x00", gl.FRAGMENT_SHADER)
	defer fs.Delete()
	if err != nil {
		return
	}
	p, err := NewProgram(vs, fs)
	if err != nil {
		return
	}
	out.Prog = p
	out.M = p.GetUniformLocation("M")
	out.V = p.GetUniformLocation("V")
	out.P = p.GetUniformLocation("P")
	out.N = p.GetUniformLocation("N")
	out.Diffuse = p.GetUniformLocation("diffuse")
	out.Light = p.GetUniformLocation("pointlight")
	out.Eye = p.GetUniformLocation("eye")
	return
}

// Delete releases all resources held by this program object
func (rp *RenderProgram) Delete() {
	rp.Prog.Delete()
}
