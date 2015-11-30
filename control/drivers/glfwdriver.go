package drivers

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/luxengine/lux/control"
)

var _ control.Driver = &GLFWDriver{}

var glfwDriverInitialized bool

var glfwDriverInstance GLFWDriver

// GLFWDriver is a driver to lux/control that is based on glfw's input.
type GLFWDriver struct {
	w *glfw.Window
}

// Poll polls glfw events.
func (d *GLFWDriver) Poll() { glfw.PollEvents() }

func (d *GLFWDriver) keycb(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
}

// NewGLFWDriver initializes the glfw controller driver.
func NewGLFWDriver(w *glfw.Window) GLFWDriver {
	if !glfwDriverInitialized {
		initializeGLFWDriver(w)
	}
	return glfwDriverInstance
}

func initializeGLFWDriver(w *glfw.Window) {
	glfwDriverInstance.w = w
	w.SetKeyCallback(glfwDriverInstance.keycb)
}
