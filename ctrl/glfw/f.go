package glfwd

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/luxengine/ctrl"
)

type actionSet struct {
	Name string
}

var acsets []actionSet

// GetConnectedControllers returns the list of controllers this Driver can
// support
func (d *Driver) GetConnectedControllers() []Controller {
	return nil
}

// Update gives a chance to poll the controllers.
func (d *Driver) Update() {
	glfw.PollEvents()
}

// LoadFormat explain the different actions sets that are available to the
// players
func (d *Driver) LoadFormat(format ctrl.ActionSetFormat) {
	for _, set := range format.Actionsets {
		acsets = append(acsets, actionSet{
			Name: set.Name,
		})
	}
}

// Name returns the name of the controller, "Keyboard+Mouse", "ps3 1", etc
// different controllers from the same driver should return different names.
// Names should adjust when controllers disconnect/reconnect/are added.
func (c *Controller) Name() string {
	return "Glfw controller"
}

// GetActionSet returns the action set with the given name. Should be called
// once at game startup but needs it needs to be possible to call this
// repetitively without any problem. Return nil if the set is not found.
func (c *Controller) GetActionSet(name string) ctrl.ActionSet {
	return c.sets[name]
}

// LoadControllerConfiguration maps the button to the action sets.
func (c *Controller) LoadControllerConfiguration(something struct{}) error {
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
	return ""
}
