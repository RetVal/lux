package glfwd

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/luxengine/lux/ctrl"
)

// Init registers the glfw driver with ctrl.
func Init(w *glfw.Window) {
	ctrl.AddDriver(driver{w})
}

// driver is a driver implemented for glfw.
type driver struct {
	w *glfw.Window
}

// GetConnectedControllers returns the list of controllers this Driver can
// support
func (d driver) GetConnectedControllers() []ctrl.Controller {
	return nil
}

// Update gives a chance to poll the controllers.
func (d driver) Update() {
	glfw.PollEvents()
}

// LoadFormat explain the different actions sets that are available to the
// players
func (d driver) LoadFormat(format ctrl.ActionSetFormat) {
	for _, set := range format.Actionsets {
		acsets = append(acsets, ActionSet{
			name: set.Name,
		})
	}
}

// controller is a controller implemented for glfw controllers.
type controller struct {
}

// ActionSet is an action set implemented for glfw actions sets.
type ActionSet struct {
	name    string
	Actions []struct{}
}

var acsets []ActionSet

// Name returns the name of the controller, "Keyboard+Mouse", "ps3 1", etc
// different controllers from the same driver should return different names.
// Names should adjust when controllers disconnect/reconnect/are added.
func (c *controller) Name() string {
	return "Keyboard+Mouse"
}

// GetActionSet returns the action set with the given name. Should be called
// once at game startup but needs it needs to be possible to call this
// repetitively without any problem. Return nil if the set is not found.
func (c *controller) GetActionSet(name string) ctrl.ActionSet {
	as := acsets[name]
	return &as
}

// LoadControllerConfiguration maps the button to the action sets.
func (c *controller) LoadControllerConfiguration(something struct{}) error {
	return nil
}

// GetDigitalAction returns the DigitalAction associated with the given
// string.
func (a *ActionSet) GetDigitalAction(string) ctrl.DigitalAction {
	return nil
}

// GetAnalogAction returns the DigitalAction associated with the given
// string.
func (a *ActionSet) GetAnalogAction(string) ctrl.AnalogAction {
	return nil
}

// Activate sets this ActionSet as the active action set with the associated
// controller. It needs to be possible to call this repetitively without
// problems.
func (a *ActionSet) Activate() {}

// Name returns the localized name of the actionset
func (a *ActionSet) Name() string {
	return a.name
}
