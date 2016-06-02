package lux

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/luxengine/gl"
)

//obvious ryze quote
func letsglowLETSGLOW() {
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
}

//SetContext will set a OpenGL core 3.3 context with foward compatibility and debug context
func SetContext() {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    // Necessary for OS X
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // Necessary for OS X
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)
}

func headlessContext() {
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    // Necessary for OS X
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // Necessary for OS X
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.False)
}

func glbs() {
	//gl.Enable(gl.DEPTH_TEST)
	//gl.DepthFunc(gl.LESS)
	gl.ClearColor.Set(0.3, 0.3, 0.3, 1.0)
	gl.CullFace.Enable()
}

//InitGLFW will call glfw.Init and panic if it fails
func InitGLFW() {
	if err := glfw.Init(); err != nil {
		panic("failed to initialize glfw")
	}
}

//TerminateGLFW is an alias for glfw.Terminate
func TerminateGLFW() {
	glfw.Terminate()
}

//CreateWindow creates a new glfw window. If fullscreen will place the screen on primary monitor.
func CreateWindow(width, height int, title string, fullscreen bool) (window *glfw.Window) {
	SetContext()
	var x *glfw.Monitor
	if fullscreen {
		x = glfw.GetPrimaryMonitor()
	}
	window, err := glfw.CreateWindow(width, height, title, x, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	letsglowLETSGLOW()
	glbs()
	QueryExtentions()
	return
}

//StartHeadless will initialize everything but wont actually create a window, so you can test your application.
func StartHeadless() (window *glfw.Window) {
	var err error
	headlessContext()
	window, err = glfw.CreateWindow(1, 1, "", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	letsglowLETSGLOW()
	glbs()
	QueryExtentions()
	return
}
