package glfwd

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

// Driver is a driver implemented for glfw.
type Driver struct {
	w *glfw.Window
}

// Controller is a controller implemented for glfw controllers.
type Controller struct {
	sets map[string]ActionSet
}

// ActionSet is an action set implemented for glfw actions sets.
type ActionSet struct {
	Name    string
	Actions []struct{}
}
