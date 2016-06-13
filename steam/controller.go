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
	"math"
	"unsafe"
)

// consts related to steam controllers
const (
	MaxCount          = 16
	MaxAnalogActions  = 16
	MaxDigitalActions = 32
	MaxOrigins        = 8

	// When sending an option to a specific controller handle, you can send to all controllers via this command
	HandleAllControllers = math.MaxUint64
	MinAnalogActionData  = -1.0
	MaxAnalogActionData  = 1.0

	//MaxSteamControllers = 16

	//SteamRightTriggerMask          = 0x0000001
	//SteamLeftTriggerMask           = 0x0000002
	//SteamRightBumperMask           = 0x0000004
	//SteamLeftBumperMask            = 0x0000008
	//SteamButton0Mask               = 0x0000010
	//SteamButton1Mask               = 0x0000020
	//SteamButton2Mask               = 0x0000040
	//SteamButton3Mask               = 0x0000080
	//SteamTouch0Mask                = 0x0000100
	//SteamTouch1Mask                = 0x0000200
	//SteamTouch2Mask                = 0x0000400
	//SteamTouch3Mask                = 0x0000800
	//SteamButtonMenuMask            = 0x0001000
	//SteamButtonSteamMask           = 0x0002000
	//SteamButtonEscapeMask          = 0x0004000
	//SteamButtonBackLeftMask        = 0x0008000
	//SteamButtonBackRightMask       = 0x0010000
	//SteamButtonLeftpadClickedMask  = 0x0020000
	//SteamButtonRightpadClickedMask = 0x0040000
	//SteamLeftpadFingerdownMask     = 0x0080000
	//SteamRightpadFingerdownMask    = 0x0100000
	//SteamJoystickButtonMask        = 0x0400000
)

// ESteamControllerPad is an enum to differentiate the 2 different steam
// controller pads.
type ESteamControllerPad int32

// enum to differentiate the 2 different steam controller pads.
const (
	SteamControllerPadLeft ESteamControllerPad = iota
	SteamControllerPadRight
)

// EControllerSourceMode is an enum that defines the different controller source
// modes.
type EControllerSourceMode int32

// controller source mode enums
const (
	EControllerSourceModeNone EControllerSourceMode = iota
	EControllerSourceModeDpad
	EControllerSourceModeButtons
	EControllerSourceModeFourButtons
	EControllerSourceModeAbsoluteMouse
	EControllerSourceModeRelativeMouse
	EControllerSourceModeJoystickMove
	EControllerSourceModeJoystickCamera
	EControllerSourceModeScrollWheel
	EControllerSourceModeTrigger
	EControllerSourceModeTouchMenu
)

// EControllerActionOrigin is an enum to differentiate the different steam
// controller action origins.
type EControllerActionOrigin int32

// ControllerActionSetHandle is a handle to a steam controller action set.
type ControllerActionSetHandle uint64

// ControllerDigitalActionHandle is a handle to a steam controller digital
// action.
type ControllerDigitalActionHandle uint64

// ControllerAnalogActionHandle is a handle to a steam controller analog
// action.
type ControllerAnalogActionHandle uint64

// ControllerHandle is a handle to a steam controller.
type ControllerHandle uint64

// ControllerDigitalActionData is a struct returned by digital action queries.
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

// ControllerAnalogActionData is a struct returned by analog action queries.
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

// IController is a handler for the steam controller interface.
type IController struct{ unsafe.Pointer }

// Controller returns the controller interface, will return an invalid
// interface if Init returned false or has not been called yet.
func Controller() IController {
	return IController{C.CSteamController()}
}

// Init must be called when starting the use of the IController interface.
func (c IController) Init() bool {
	return bool(C.SteamCAPI_SteamController_Init(c.Pointer))
}

// Shutdown must be called when ending the use of the IController
// interface.
func (c IController) Shutdown() bool {
	return bool(C.SteamCAPI_SteamController_Shutdown(c.Pointer))
}

// RunFrame pumps callback/callresult events
// Note: SteamAPI_RunCallbacks will do this for you, so you should never need to
// call this directly.
func (c IController) RunFrame() {
	C.SteamCAPI_SteamController_RunFrame(c.Pointer)
}

// GetConnectedControllers enumerate currently connected controllers.
func (c IController) GetConnectedControllers(handles *[MaxCount]ControllerHandle) int {
	return int(C.SteamCAPI_SteamController_GetConnectedControllers(c.Pointer, (*C.ControllerHandle_t)(&handles[0])))
}

// ShowBindingPanel invokes the Steam overlay and brings up the binding screen
// Returns false if overlay is disabled / unavailable, or the user is not in Big
// Picture mode
func (c IController) ShowBindingPanel(controllerHandle ControllerHandle) bool {
	return bool(C.SteamCAPI_SteamController_ShowBindingPanel(c.Pointer, C.ControllerHandle_t(controllerHandle)))
}

// GetActionSetHandle lookups the handle for an Action Set. Best to do this once
// on startup, and store the handles for all future API calls.
func (c IController) GetActionSetHandle(name string) ControllerActionSetHandle {
	cname := C.CString(name)
	r := ControllerActionSetHandle(C.SteamCAPI_SteamController_GetActionSetHandle(c.Pointer, cname))
	C.free(unsafe.Pointer(cname))
	return r
}

// ActivateActionSet reconfigures the controller to use the specified action set
// (ie 'Menu', 'Walk' or 'Drive') This is cheap, and can be safely called
// repeatedly. It's often easier to repeatedly call it in your state loops,
// instead of trying to place it in all of your state transitions.
func (c IController) ActivateActionSet(controllerHandle ControllerHandle, actionSetHandle ControllerActionSetHandle) {
	C.SteamCAPI_SteamController_ActivateActionSet(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerActionSetHandle_t(actionSetHandle))
}

// GetCurrentActionSet returns the currently active Action Set.
func (c IController) GetCurrentActionSet(controllerHandle ControllerHandle) ControllerActionSetHandle {
	return ControllerActionSetHandle(C.SteamCAPI_SteamController_GetCurrentActionSet(c.Pointer, C.ControllerHandle_t(controllerHandle)))
}

// GetDigitalActionHandle lookups the handle for a digital action. Best to do this
// once on startup, and store the handles for all future API calls.
func (c IController) GetDigitalActionHandle(name string) ControllerDigitalActionHandle {
	cname := C.CString(name)
	r := ControllerDigitalActionHandle(C.SteamCAPI_SteamController_GetDigitalActionHandle(c.Pointer, cname))
	C.free(unsafe.Pointer(cname))
	return r
}

// GetDigitalActionData returns the current state of the supplied digital game
// action.
func (c IController) GetDigitalActionData(controllerHandle ControllerHandle, digitalActionHandle ControllerDigitalActionHandle) ControllerDigitalActionData {
	return controllerDigitalActionDataFromC(C.SteamCAPI_SteamController_GetDigitalActionData(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerDigitalActionHandle_t(digitalActionHandle)))
}

// GetDigitalActionOrigins gets the origin(s) for a digital action within an
// action set. Returns the number of origins supplied in originsOut. Use this to
// display the appropriate on-screen prompt for the action. originsOut should
// point to a STEAM_CONTROLLER_MAX_ORIGINS sized array of
// EControllerActionOrigin handles
func (c IController) GetDigitalActionOrigins(controllerHandle ControllerHandle, actionSetHandle ControllerActionSetHandle, digitalActionHandle ControllerDigitalActionHandle, originsOut *EControllerActionOrigin) int {
	return int(C.SteamCAPI_SteamController_GetDigitalActionOrigins(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerActionSetHandle_t(actionSetHandle), C.ControllerDigitalActionHandle_t(digitalActionHandle), (*C.EControllerActionOrigin)(originsOut)))
}

// GetAnalogActionHandle lookups the handle for an analog action. Best to do
// this once on startup, and store the handles for all future API calls.
func (c IController) GetAnalogActionHandle(name string) ControllerAnalogActionHandle {
	cname := C.CString(name)
	r := ControllerAnalogActionHandle(C.SteamCAPI_SteamController_GetAnalogActionHandle(c.Pointer, cname))
	C.free(unsafe.Pointer(cname))
	return r
}

// GetAnalogActionData returns the current state of these supplied analog game
// action.
func (c IController) GetAnalogActionData(controllerHandle ControllerHandle, analogActionHandle ControllerAnalogActionHandle) ControllerAnalogActionData {
	return controllerAnalogActionDatafromC(C.SteamCAPI_SteamController_GetAnalogActionData(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerAnalogActionHandle_t(analogActionHandle)))
}

// GetAnalogActionOrigins gets the origin(s) for an analog action within an
// action set. Returns the number of origins supplied in originsOut. Use this to
// display the appropriate on-screen prompt for the action.
func (c IController) GetAnalogActionOrigins(controllerHandle ControllerHandle, actionSetHandle ControllerActionSetHandle, analogActionHandle ControllerAnalogActionHandle, origins *[MaxOrigins]EControllerActionOrigin) int {
	return int(C.SteamCAPI_SteamController_GetAnalogActionOrigins(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerActionSetHandle_t(actionSetHandle), C.ControllerAnalogActionHandle_t(analogActionHandle), (*C.EControllerActionOrigin)(&origins[0])))
}

// StopAnalogActionMomentum has no documentation.
func (c IController) StopAnalogActionMomentum(controllerHandle ControllerHandle, action ControllerAnalogActionHandle) {
	C.SteamCAPI_SteamController_StopAnalogActionMomentum(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ControllerAnalogActionHandle_t(action))
}

// TriggerHapticPulse triggers a haptic pulse on a controller.
func (c IController) TriggerHapticPulse(controllerHandle ControllerHandle, targetPad ESteamControllerPad, durationms uint16) {
	C.SteamCAPI_SteamController_TriggerHapticPulse(c.Pointer, C.ControllerHandle_t(controllerHandle), C.ESteamControllerPad(targetPad), C.ushort(durationms))
}
