// Package ctrl gives the users the ability to use a valve style controller API
// to deal with all major controller {kb+m, steam, ps3, ps4, xbox360, xboxone}.
// To use the API start by importing ctrl and any driver package you want,
// verify if your platform supports it.
//	import (
//		"github.com/luxengine/ctrl"
//		_ "github.com/luxengine/ctrl/glfw"
//		_ "github.com/luxengine/ctrl/steam"
//	)
// Folling Valve's model, the first you have to do is declare a file that
// explains what are the actions you expect to be exposed. Super Mario Bros
// (NES) would define 4-7 actions (depending on how you count), "move", "jump",
// "run" and "pause". Note that this initial file does NOT bind any key to the
// actions, it only defines what are the expected actions.
//
// There are 2 types of possible actions, digital actions and analog actions.
// digital actiosn represent everything is either on or off (like buttons),
// while digital actions represent everything that have a range of values,
// including mouse, joysticks and touch points.
//
// Since the same button can be used both inside eg. in-menus and in-game, we
// wrap actions in action sets. Only one action set can be active at a time and
// actions that aren't in the active set are disabled until their owner set is
// activated.
//
// Once you have your actions defined you must make a default binding in another
// file (or any io.Reader), these bind actions to specific devices
// (buttons/mouse/joystick). This allows players to change their configuration
// much more easilly.
//
// The API usage should look like this:
//
//	//import all the stuff
//	//somehow read the necessary action file and default binding
//	p1 := GetConnectedControllers()[0] // or give one to each player
//	ingame := p1.GetActionSet("callofdutyingame")
//	ingame.Activate() // this makes the in-game action set the acive one
//	fire := ingame.GetDigitalAction("fire")
//	...
//	if fire.Data().Pressed() {
//		// fire some bullets
//	}
//
// Having this setup, as valve explained, allow for configuration to be swapped
// much more easily while the app is running and, according to a lot of people
// who use steam controllers, is really fun to use. Also glfw/ps/xbox do not
// define any sort of API to use the device while steam does, so why force
// things.
package ctrl
