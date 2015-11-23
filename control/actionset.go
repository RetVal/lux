package control

// LoadControls loads the control in the given file.
func LoadControls() {}

// ActionSet is a handle to a set of action grouped togheter in a logical unit.
type ActionSet int

// Action is a handle to a particular action within an action set.
type Action int

// FindActionSet returns the handle to the given action set or an error if it
// wasn't found.
func FindActionSet(name string) (ActionSet, error) {}

// FindAction returns the handle for an Action in the called ActionSet.
func (as ActionSet) FindAction(name string) (Action, error) {}

// Activate makes this action set the current action set.
func (as ActionSet) Activate() {}

// WasPressed returns true if the button bound to this action went from up to
// down in the last frame.
func (a Action) WasPressed() bool {}

// IsDown returns true if the button bound to this action is currently pressed.
func (a Action) IsDown() bool {}

// IsUp returns true if the button bound to this action is currently not
// pressed.
func (a Action) IsUp() bool
