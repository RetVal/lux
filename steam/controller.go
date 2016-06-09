package steam

/*
#cgo CFLAGS: -Ih
#cgo linux LDFLAGS: -Llib -lsteam_capi -lsteam_api
#cgo darwin LDFLAGS: -Llib -lsteam_capi
#include <steam_capi.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

const (
	MAX_STEAM_CONTROLLERS = 16

	STEAM_RIGHT_TRIGGER_MASK           = 0x0000001
	STEAM_LEFT_TRIGGER_MASK            = 0x0000002
	STEAM_RIGHT_BUMPER_MASK            = 0x0000004
	STEAM_LEFT_BUMPER_MASK             = 0x0000008
	STEAM_BUTTON_0_MASK                = 0x0000010
	STEAM_BUTTON_1_MASK                = 0x0000020
	STEAM_BUTTON_2_MASK                = 0x0000040
	STEAM_BUTTON_3_MASK                = 0x0000080
	STEAM_TOUCH_0_MASK                 = 0x0000100
	STEAM_TOUCH_1_MASK                 = 0x0000200
	STEAM_TOUCH_2_MASK                 = 0x0000400
	STEAM_TOUCH_3_MASK                 = 0x0000800
	STEAM_BUTTON_MENU_MASK             = 0x0001000
	STEAM_BUTTON_STEAM_MASK            = 0x0002000
	STEAM_BUTTON_ESCAPE_MASK           = 0x0004000
	STEAM_BUTTON_BACK_LEFT_MASK        = 0x0008000
	STEAM_BUTTON_BACK_RIGHT_MASK       = 0x0010000
	STEAM_BUTTON_LEFTPAD_CLICKED_MASK  = 0x0020000
	STEAM_BUTTON_RIGHTPAD_CLICKED_MASK = 0x0040000
	STEAM_LEFTPAD_FINGERDOWN_MASK      = 0x0080000
	STEAM_RIGHTPAD_FINGERDOWN_MASK     = 0x0100000
	STEAM_JOYSTICK_BUTTON_MASK         = 0x0400000
)

type ESteamControllerPad int32

const (
	SteamControllerPad_Left ESteamControllerPad = iota
	SteamControllerPad_Right
)

type EControllerSourceMode int32

const (
	EControllerSourceMode_None EControllerSourceMode = iota
	EControllerSourceMode_Dpad
	EControllerSourceMode_Buttons
	EControllerSourceMode_FourButtons
	EControllerSourceMode_AbsoluteMouse
	EControllerSourceMode_RelativeMouse
	EControllerSourceMode_JoystickMove
	EControllerSourceMode_JoystickCamera
	EControllerSourceMode_ScrollWheel
	EControllerSourceMode_Trigger
	EControllerSourceMode_TouchMenu
)

type (
	ControllerActionSetHandle     uint64
	ControllerDigitalActionHandle uint64
	ControllerAnalogActionHandle  uint64
	ControllerHandle              uint64
	EControllerActionOrigin       int32
)

type ControllerDigitalActionData struct {
	// The current state of this action; will be true if currently pressed
	State bool

	// Whether or not this action is currently available to be bound in the active action set
	Active bool

	// The C version for some function.
	c C.ControllerDigitalActionData_t
}

func controllerDigitalActionDataFromC(c C.ControllerDigitalActionData_t) ControllerDigitalActionData {
	return ControllerDigitalActionData{
		State:  bool(c.bState),
		Active: bool(c.bActive),
	}
}

type ControllerAnalogActionData struct {
	// Type of data coming from this action, this will match what got specified in the action set
	Mode EControllerSourceMode

	// The current state of this action; will be delta updates for mouse actions
	X, Y float32

	// Whether or not this action is currently available to be bound in the active action set
	Active bool
}

func controllerAnalogActionDatafromC(c C.ControllerAnalogActionData_t) ControllerAnalogActionData {
	return ControllerAnalogActionData{
		Mode:   EControllerSourceMode(c.eMode),
		X:      float32(c.x),
		Y:      float32(c.y),
		Active: bool(c.bActive),
	}
}

// ISteamController is a handler for the steam controller API.
type ISteamController struct{ unsafe.Pointer }

// SteamController returns the controller interface, will return an invalid
// interface if Init returned false or has not been called yet.
func SteamController() ISteamController {
	return ISteamController{C.CSteamController()}
}

// Init must be called when starting the use of the ISteamController interface.
func (c ISteamController) Init() bool {
	return bool(C.SteamCAPI_SteamController_Init(c.Pointer))
}

// Shutdown must be called when ending the use of the ISteamController
// interface.
func (c ISteamController) Shutdown() bool {
	return bool(C.SteamCAPI_SteamController_Shutdown(c.Pointer))
}

// RunFrame pumps callback/callresult events
// Note: SteamAPI_RunCallbacks will do this for you, so you should never need to
// call this directly.
func (c ISteamController) RunFrame() {
	C.SteamCAPI_SteamController_RunFrame(c.Pointer)
}

// GetConnectedControllers enumerate currently connected controllers
// handlesOut should point to a STEAM_CONTROLLER_MAX_COUNT sized array of ControllerHandle_t handles
// Returns the number of handles written to handlesOut
func (c ISteamController) GetConnectedControllers(handlesOut *ControllerHandle) int {
	return int(C.SteamCAPI_SteamController_GetConnectedControllers(c.Pointer, (*C.ControllerHandle_t)(handlesOut)))
}

// ShowBindingPanel invokes the Steam overlay and brings up the binding screen
// Returns false if overlay is disabled / unavailable, or the user is not in Big
// Picture mode
func (c ISteamController) ShowBindingPanel(controllerHandle ControllerHandle) bool {
	return bool(C.SteamCAPI_SteamController_ShowBindingPanel(c.Pointer, C.ControllerHandle_t(controllerHandle)))
}

// GetActionSetHandle lookups the handle for an Action Set. Best to do this once
// on startup, and store the handles for all future API calls.
func (c ISteamController) GetActionSetHandle(actionSetName string) ControllerActionSetHandle {
	cActionSetName := C.CString(actionSetName)
	r := ControllerActionSetHandle(C.SteamCAPI_SteamController_GetActionSetHandle(c.Pointer, cActionSetName))
	C.free(unsafe.Pointer(cActionSetName))
	return r
}

// ActivateActionSet reconfigures the controller to use the specified action set
// (ie 'Menu', 'Walk' or 'Drive') This is cheap, and can be safely called
// repeatedly. It's often easier to repeatedly call it in your state loops,
// instead of trying to place it in all of your state transitions.
func (c ISteamController) ActivateActionSet(controllerHandle ControllerHandle, actionSetHandle ControllerActionSetHandle) {
	C.SteamCAPI_SteamController_ActivateActionSet(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerActionSetHandle_t(actionSetHandle))
}

// GetCurrentActionSet returns the currently active Action Set.
func (c ISteamController) GetCurrentActionSet(controllerHandle ControllerHandle) ControllerActionSetHandle {
	return ControllerActionSetHandle(C.SteamCAPI_SteamController_GetCurrentActionSet(c.Pointer, C.ControllerHandle_t(controllerHandle)))
}

// GetDigitalActionHandle lookups the handle for a digital action. Best to do this
// once on startup, and store the handles for all future API calls.
func (c ISteamController) GetDigitalActionHandle(actionName string) ControllerDigitalActionHandle {
	cActionSetName := C.CString(actionName)
	r := ControllerDigitalActionHandle(C.SteamCAPI_SteamController_GetDigitalActionHandle(c.Pointer, cActionSetName))
	C.free(unsafe.Pointer(cActionSetName))
	return r
}

// GetDigitalActionData returns the current state of the supplied digital game
// action.
func (c ISteamController) GetDigitalActionData(controllerHandle ControllerHandle, digitalActionHandle ControllerDigitalActionHandle) ControllerDigitalActionData {
	return controllerDigitalActionDataFromC(C.SteamCAPI_SteamController_GetDigitalActionData(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerDigitalActionHandle_t(digitalActionHandle)))
}

// GetDigitalActionOrigins gets the origin(s) for a digital action within an
// action set. Returns the number of origins supplied in originsOut. Use this to
// display the appropriate on-screen prompt for the action. originsOut should
// point to a STEAM_CONTROLLER_MAX_ORIGINS sized array of
// EControllerActionOrigin handles
func (c ISteamController) GetDigitalActionOrigins(controllerHandle ControllerHandle, actionSetHandle ControllerActionSetHandle, digitalActionHandle ControllerDigitalActionHandle, originsOut *EControllerActionOrigin) int {
	return int(C.SteamCAPI_SteamController_GetDigitalActionOrigins(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerActionSetHandle_t(actionSetHandle), C.ControllerDigitalActionHandle_t(digitalActionHandle), (*C.EControllerActionOrigin)(originsOut)))
}

// GetAnalogActionHandle lookups the handle for an analog action. Best to do
// this once on startup, and store the handles for all future API calls.
func (c ISteamController) GetAnalogActionHandle(actionName string) ControllerAnalogActionHandle {
	cActionSetName := C.CString(actionName)
	r := ControllerAnalogActionHandle(C.SteamCAPI_SteamController_GetAnalogActionHandle(c.Pointer, cActionSetName))
	C.free(unsafe.Pointer(cActionSetName))
	return r
}

// GetAnalogActionData returns the current state of these supplied analog game
// action.
func (c ISteamController) GetAnalogActionData(controllerHandle ControllerHandle, analogActionHandle ControllerAnalogActionHandle) ControllerAnalogActionData {
	return controllerAnalogActionDatafromC(C.SteamCAPI_SteamController_GetAnalogActionData(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerAnalogActionHandle_t(analogActionHandle)))
}

// GetAnalogActionOrigins gets the origin(s) for an analog action within an
// action set. Returns the number of origins supplied in originsOut. Use this to
// display the appropriate on-screen prompt for the action. originsOut should
// point to a STEAM_CONTROLLER_MAX_ORIGINS sized array of
// EControllerActionOrigin handles
func (c ISteamController) GetAnalogActionOrigins(controllerHandle ControllerHandle, actionSetHandle ControllerActionSetHandle, analogActionHandle ControllerAnalogActionHandle, originsOut *EControllerActionOrigin) int {
	return int(C.SteamCAPI_SteamController_GetAnalogActionOrigins(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerActionSetHandle_t(actionSetHandle), C.ControllerAnalogActionHandle_t(analogActionHandle), (*C.EControllerActionOrigin)(originsOut)))
}

func (c ISteamController) StopAnalogActionMomentum(controllerHandle ControllerHandle, eAction ControllerAnalogActionHandle) {
	C.SteamCAPI_SteamController_StopAnalogActionMomentum(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerAnalogActionHandle_t(eAction))
}

// TriggerHapticPulse triggers a haptic pulse on a controller.
func (c ISteamController) TriggerHapticPulse(controllerHandle ControllerHandle, eTargetPad ESteamControllerPad, usDurationMicroSec uint16) {
	C.SteamCAPI_SteamController_TriggerHapticPulse(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ESteamControllerPad(eTargetPad), C.ushort(usDurationMicroSec))
}
