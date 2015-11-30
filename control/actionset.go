package control

//TODO:
// - add a way to differentiate between player 1-16
// - add drivers of some sort so you can chose between glfw, steam or Xinput (or other input methods, like http or AI)

type actionSetF struct {
	Name    string    `json:"name"`
	actions []actionF `json:"actions"`
}

type actionF struct {
	Name  string `json:"name"`
	Input input  `json:"input"`
}

type input struct {
	// is it the mouse or a joystick
	Mouse bool

	// joystick, mouse, accelerometer
	Inputmode string `json:"input_mode"`

	// is it a button
	Button bool

	//button only, can it rapid fire
	Rapidfire bool `json:"rapid_fire"`
}

/*
control file = list of action set
action set = list of action, name
action = name, input
input = button input or joystick/mouse input
*/

// Controller represents the controller of a particular player.
type Controller struct {
	driver Driver
}

// NewController returns a controller based on the given driver.
func NewController(driver Driver) Controller {
	return Controller{
		driver: driver,
	}
}

// ActionSet is a handle to a set of action grouped togheter in a logical unit.
type ActionSet int32

// Action is a handle to a particular action within an action set.
type Action int32

// Poll organises the data for the next frame. This needs to be called or you
// may get inaccurate results.
func (c Controller) Poll() {}

// FindActionSet returns the handle to the given action set or an error if it
// wasn't found. You would call this once per action set.
func (c Controller) FindActionSet(name string) (ActionSet, error) { return 0, nil }

// FindAction returns the handle for an Action in the called ActionSet. You
// would usually call this once per action.
func (c Controller) FindAction(as ActionSet, name string) (Action, error) { return 0, nil }

// Activate makes this action set the current action set.
func (c Controller) Activate(as ActionSet) {}

// WasPressed returns true if the button bound to this action went from up to
// down in the last frame.
func (c Controller) WasPressed(a Action) bool { return false }

// IsDown returns true if the button bound to this action is currently pressed.
func (c Controller) IsDown(a Action) bool { return false }

// IsUp returns true if the button bound to this action is currently not
// pressed.
func (c Controller) IsUp(a Action) bool { return false }

// AnalogAbs returns the absolute position of the analog device.
func (c Controller) AnalogAbs(a Action) (float32, float32) { return 0, 0 }

// AnalogD returns the difference between this frame and last frame absolute
// position for this action.
func (c Controller) AnalogD(a Action) (float32, float32) { return 0, 0 }
