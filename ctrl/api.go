package ctrl

import (
	"encoding/json"
	"io"
)

var (
	drivers []Driver
)

// ActionSetFormat holds the action sets + actions available to the players.
type ActionSetFormat struct {
	Actionsets []struct {
		Name         string `json:"name"`
		Stickpadgyro []struct {
			Name      string `json:"name"`
			Inputmode string `json:"inputmode"`
		} `json:"stickpadgyro"`
		Analogtriggers []struct {
			Name string `json:"name"`
		} `json:"analogtriggers"`
		Buttons []struct {
			Name string `json:"name"`
		} `json:"buttons"`
	} `json:"actionsets"`
}

// Driver is the generic driver interface for any of the supported controllers.
type Driver interface {
	// GetConnectedControllers returns the list of controllers this Driver can
	// support
	GetConnectedControllers() []Controller

	// Update gives a chance to poll the controllers.
	Update()

	// LoadFormat explain the different actions sets that are available to the
	// players
	LoadFormat(ActionSetFormat)
}

// Controller is the generic interface for any of the supported controllers.
type Controller interface {
	// Name returns the name of the controller, "Keyboard+Mouse", "ps3 1", etc
	// different controllers from the same driver should return different names.
	// Names should adjust when controllers disconnect/reconnect/are added.
	Name() string

	// GetActionSet returns the action set with the given name. Should be called
	// once at game startup but needs it needs to be possible to call this
	// repetitively without any problem. Return nil if the set is not found.
	GetActionSet(string) ActionSet

	// LoadControllerConfiguration maps the button to the action sets.
	LoadControllerConfiguration(struct{}) error
}

// ActionSet is a handle to a list of actions that the user can activate, such
// as jumping, firing, interracting with the environnement, opening menus.
// ActionSet instances are tied to a specific controller.
type ActionSet interface {
	// GetDigitalAction returns the DigitalAction associated with the given
	// string.
	GetDigitalAction(string) DigitalAction

	// GetAnalogAction returns the DigitalAction associated with the given
	// string.
	GetAnalogAction(string) AnalogAction

	// Activate sets this ActionSet as the active action set with the associated
	// controller. It needs to be possible to call this repetitively without
	// problems.
	Activate()

	// Name returns the localized name of the actionset.
	Name() string
}

// DigitalAction is a action that the player can take, they represent buttons
// that are either pressed or not. Actions are tied to a specific instance of a
// ActionSet.
type DigitalAction interface {
	// Data reads the state of the in-game action.
	Data() DigitalActionData

	// Name returns the localized name of the action
	Name() string

	// ActionSet returns the parent action set.
	ActionSet() ActionSet

	// GetOrigin returns the origin of the action input, this is used to find
	// the appropriate glyph (image) to use with the associated action. Eg.
	// display the A xbox image when trying to talk to an NPC.
	GetOrigin() struct{}
}

// AnalogAction is a action that the player can take, they represent all actions
// that have non boolean status (mouse, joysticks, etc). Actions are tied to a
// specific instance of a ActionSet.
type AnalogAction interface {
	// Data reads the state of the in-game action.
	Data() AnalogActionData

	// Name returns the localized name of the action
	Name() string

	// ActionSet returns the parent action set.
	ActionSet() ActionSet

	// GetOrigin returns the origin of the action input, this is used to find
	// the appropriate glyph (image) to use with the associated action. Eg.
	// display the A xbox image when trying to talk to an NPC.
	GetOrigin() struct{}
}

// ControllerSourceMode is an enum for the types of mode analog signals can
// have.
type ControllerSourceMode int16

// The different controller source mode
const (
	None ControllerSourceMode = iota
	Dpad
	Buttons
	FourButtons
	AbsoluteMouse
	RelativeMouse
	JoystickMove
	JoystickCamera
	ScrollWheel
	Trigger
	TouchMenu
)

// AnalogActionData is the data returned from AnalogAction queries.
type AnalogActionData struct {
	// Type of data coming from this action, this will match what got specified
	// in the action set
	Mode ControllerSourceMode

	// Whether or not this action is currently available to be bound in the
	// active action set
	Active bool

	// The current state of this action; will be delta updates for mouse or
	// joysticks actions and for trigger like mechanism only the X value will be
	// filled.
	X, Y float32
}

// DigitalActionData is the data returned from DigitalAction queries.
type DigitalActionData struct {
	// Whether or not this action is currently available to be bound in the
	// active action set.
	Active bool

	// The current state of this action; will be true if currently pressed
	State bool
}

// AddDriver adds the given driver {kb+m, steam, ps(3|4), xbox(360|one)} to be
// used. Only driver packages should be calling this function.
func AddDriver(d Driver) {
	// TODO(hydroflame): store a list of drivers
	drivers = append(drivers, d)
}

// LoadCtrlFormat will load the formatted controller file and passes the
// actionset/action structure to all controllers. It's important to understand
// that this data only explain what the different actions sets and actions are
// and not their bindings.
func LoadCtrlFormat(r io.Reader) error {
	// TODO(hydroflame): implement config loading.
	var format ActionSetFormat

	err := json.NewDecoder(r).Decode(&format)
	if err != nil {
		return err
	}
	return nil
}

// GetConnectedControllers returns the slice of currently connected controller.
func GetConnectedControllers() []Controller {
	var ctrls []Controller
	for _, d := range drivers {
		ctrls = append(ctrls, d.GetConnectedControllers()...)
	}
	return ctrls
}

// Update gives a change for every driver to poll their inputs. This must be
// called once per frame.
func Update() {
	for _, d := range drivers {
		d.Update()
	}
}
