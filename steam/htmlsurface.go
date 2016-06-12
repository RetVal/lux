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

// EHTMLMouseButton is an enum for mouse buttons.
type EHTMLMouseButton int32

// The different mouse buttons you can give to steamworks.
const (
	HTMLMouseButton_Left   EHTMLMouseButton = 0
	HTMLMouseButton_Right  EHTMLMouseButton = 1
	HTMLMouseButton_Middle EHTMLMouseButton = 2
)

// EHTMLKeyModifiers is an enum for key modifiers.
type EHTMLKeyModifiers int32

// The different key modifiers you can give to steamworks.
const (
	HTMLKeyModifier_None      EHTMLKeyModifiers = 0
	HTMLKeyModifier_AltDown   EHTMLKeyModifiers = 1 << 0
	HTMLKeyModifier_CtrlDown  EHTMLKeyModifiers = 1 << 1
	HTMLKeyModifier_ShiftDown EHTMLKeyModifiers = 1 << 2
)

// Init must be called when starting use of the interface.
func (surface IHTMLSurface) Init() bool {
	return bool(C.SteamCAPI_ISteamHTMLSurface_Init(surface.Pointer))
}

// Shutdown must be called when ending use of the interface.
func (surface IHTMLSurface) Shutdown() bool {
	return bool(C.SteamCAPI_ISteamHTMLSurface_Shutdown(surface.Pointer))
}

// CreateBrowser creates a browser object for display of a html page, when
// creation is complete the call handle will return a HTML_BrowserReady_t
// callback for the HHTMLBrowser of your new browser. The user agent string is a
// substring to be added to the general user agent string so you can identify
// your client on web servers. The userCSS string lets you apply a CSS style
// sheet to every displayed page, leave "" if you do not require this
// functionality.
func (surface IHTMLSurface) CreateBrowser(userAgent, userCSS string) APICall {
	cuserAgent := C.CString(userAgent)
	cuserCSS := C.CString(userCSS)

	ret := APICall(C.SteamCAPI_ISteamHTMLSurface_CreateBrowser(surface.Pointer, cuserAgent, cuserCSS))

	C.free(unsafe.Pointer(cuserAgent))
	C.free(unsafe.Pointer(cuserCSS))

	return ret
}

// RemoveBrowser when you are done with a html surface, this lets Steam free the
// resources being used by it.
func (surface IHTMLSurface) RemoveBrowser(browser APICall) {
	C.SteamCAPI_ISteamHTMLSurface_RemoveBrowser(surface.Pointer, C.HHTMLBrowser(browser))
}

// LoadURL navigates to this URL, results in a HTML_StartRequest_t as the
// request commences.
func (surface IHTMLSurface) LoadURL(browser APICall, URL, postData string) {
	cURL := C.CString(URL)
	cpostData := C.CString(postData)

	C.SteamCAPI_ISteamHTMLSurface_LoadURL(surface.Pointer, C.HHTMLBrowser(browser), cURL, cpostData)

	C.free(unsafe.Pointer(cURL))
	C.free(unsafe.Pointer(cpostData))
}

// SetSize tells the surface the size in pixels to display the surface.
func (surface IHTMLSurface) SetSize(browserHandle APICall, width, height uint32) {
	C.SteamCAPI_ISteamHTMLSurface_SetSize(surface.Pointer, C.HHTMLBrowser(browserHandle), C.uint(width), C.uint(height))
}

// StopLoad stops the load of the current html page.
func (surface IHTMLSurface) StopLoad(browserHandle APICall) {
	C.SteamCAPI_ISteamHTMLSurface_StopLoad(surface.Pointer, C.HHTMLBrowser(browserHandle))
}

// Reload (most likely from local cache) the current page.
func (surface IHTMLSurface) Reload(browserHandle APICall) {
	C.SteamCAPI_ISteamHTMLSurface_Reload(surface.Pointer, C.HHTMLBrowser(browserHandle))
}

// GoBack navigates back in the page history.
func (surface IHTMLSurface) GoBack(browserHandle APICall) {
	C.SteamCAPI_ISteamHTMLSurface_GoBack(surface.Pointer, C.HHTMLBrowser(browserHandle))
}

// GoForward navigates forward in the page history.
func (surface IHTMLSurface) GoForward(browserHandle APICall) {
	C.SteamCAPI_ISteamHTMLSurface_GoForward(surface.Pointer, C.HHTMLBrowser(browserHandle))
}

// AddHeader adds this header to any url requests from this browser.
func (surface IHTMLSurface) AddHeader(browserHandle APICall, key, value string) {
	ckey := C.CString(key)
	cvalue := C.CString(value)

	C.SteamCAPI_ISteamHTMLSurface_AddHeader(surface.Pointer, C.HHTMLBrowser(browserHandle), ckey, cvalue)

	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(cvalue))
}

// ExecuteJavascript runs this javascript script in the currently loaded page.
func (surface IHTMLSurface) ExecuteJavascript(browserHandle APICall, script string) {
	cscript := C.CString(script)

	C.SteamCAPI_ISteamHTMLSurface_ExecuteJavascript(surface.Pointer, C.HHTMLBrowser(browserHandle), cscript)

	C.free(unsafe.Pointer(cscript))
}

// MouseUp should be called when a mouse up event happens.
func (surface IHTMLSurface) MouseUp(browserHandle APICall, mouseButton EHTMLMouseButton) {
	C.SteamCAPI_ISteamHTMLSurface_MouseUp(surface.Pointer, C.HHTMLBrowser(browserHandle), C.EHTMLMouseButton(mouseButton))
}

// MouseDown should be called when a mouse down event happens.
func (surface IHTMLSurface) MouseDown(browserHandle APICall, mouseButton EHTMLMouseButton) {
	C.SteamCAPI_ISteamHTMLSurface_MouseDown(surface.Pointer, C.HHTMLBrowser(browserHandle), C.EHTMLMouseButton(mouseButton))
}

// MouseDoubleClick should be called when a mouse double click event happens.
func (surface IHTMLSurface) MouseDoubleClick(browserHandle APICall, mouseButton EHTMLMouseButton) {
	C.SteamCAPI_ISteamHTMLSurface_MouseDoubleClick(surface.Pointer, C.HHTMLBrowser(browserHandle), C.EHTMLMouseButton(mouseButton))
}

// MouseMove should be called when a mouse move event happens.
func (surface IHTMLSurface) MouseMove(browserHandle APICall, x, y int32) {
	C.SteamCAPI_ISteamHTMLSurface_MouseMove(surface.Pointer, C.HHTMLBrowser(browserHandle), C.int(x), C.int(y))
}

// MouseWheel should be called when a mouse wheel event happens.
func (surface IHTMLSurface) MouseWheel(browserHandle APICall, delta int32) {
	C.SteamCAPI_ISteamHTMLSurface_MouseWheel(surface.Pointer, C.HHTMLBrowser(browserHandle), C.int(delta))
}

// KeyDown should be called when a key down event happens.
func (surface IHTMLSurface) KeyDown(browserHandle APICall, nativeKeyCode uint32, keyModifiers EHTMLKeyModifiers) {
	C.SteamCAPI_ISteamHTMLSurface_KeyDown(surface.Pointer, C.HHTMLBrowser(browserHandle), C.uint(nativeKeyCode), C.EHTMLKeyModifiers(keyModifiers))
}

// KeyUp should be called when a key up event happens.
func (surface IHTMLSurface) KeyUp(browserHandle APICall, nativeKeyCode uint32, keyModifiers EHTMLKeyModifiers) {
	C.SteamCAPI_ISteamHTMLSurface_KeyUp(surface.Pointer, C.HHTMLBrowser(browserHandle), C.uint(nativeKeyCode), C.EHTMLKeyModifiers(keyModifiers))
}

// KeyChar is for unicode character (and therefore multiple character per press)
func (surface IHTMLSurface) KeyChar(browserHandle APICall, unicodeChar uint32, keyModifiers EHTMLKeyModifiers) {
	C.SteamCAPI_ISteamHTMLSurface_KeyChar(surface.Pointer, C.HHTMLBrowser(browserHandle), C.uint(unicodeChar), C.EHTMLKeyModifiers(keyModifiers))
}

// SetHorizontalScroll is used to programmatically scroll.
func (surface IHTMLSurface) SetHorizontalScroll(browserHandle APICall, absolutePixelScroll uint32) {
	C.SteamCAPI_ISteamHTMLSurface_SetHorizontalScroll(surface.Pointer, C.HHTMLBrowser(browserHandle), C.uint(absolutePixelScroll))
}

// SetVerticalScroll is used to programmatically scroll.
func (surface IHTMLSurface) SetVerticalScroll(browserHandle APICall, absolutePixelScroll uint32) {
	C.SteamCAPI_ISteamHTMLSurface_SetVerticalScroll(surface.Pointer, C.HHTMLBrowser(browserHandle), C.uint(absolutePixelScroll))
}

// SetKeyFocus tells the html control if it has key focus currently, controls
// showing the I-beam cursor in text controls amongst other things.
func (surface IHTMLSurface) SetKeyFocus(browserHandle APICall, hasKeyFocus bool) {
	C.SteamCAPI_ISteamHTMLSurface_SetKeyFocus(surface.Pointer, C.HHTMLBrowser(browserHandle), C.bool(hasKeyFocus))
}

// ViewSource opens the current pages html code in the local editor of choice,
// used for debugging.
func (surface IHTMLSurface) ViewSource(browserHandle APICall) {
	C.SteamCAPI_ISteamHTMLSurface_ViewSource(surface.Pointer, C.HHTMLBrowser(browserHandle))
}

// CopyToClipboard copies the currently selected text on the html page to the
// local clipboard.
func (surface IHTMLSurface) CopyToClipboard(browserHandle APICall) {
	C.SteamCAPI_ISteamHTMLSurface_CopyToClipboard(surface.Pointer, C.HHTMLBrowser(browserHandle))
}

// PasteFromClipboard pastes from the local clipboard to the current html page.
func (surface IHTMLSurface) PasteFromClipboard(browserHandle APICall) {
	C.SteamCAPI_ISteamHTMLSurface_PasteFromClipboard(surface.Pointer, C.HHTMLBrowser(browserHandle))
}

// Find finds this string in the browser, if bCurrentlyInFind is true then
// instead cycle to the next matching element.
func (surface IHTMLSurface) Find(browserHandle APICall, search string, currentlyInFind, reverse bool) {
	csearch := C.CString(search)

	C.SteamCAPI_ISteamHTMLSurface_Find(surface.Pointer, C.HHTMLBrowser(browserHandle), csearch, C.bool(currentlyInFind), C.bool(reverse))

	C.free(unsafe.Pointer(csearch))
}

// StopFind cancels a currently running find.
func (surface IHTMLSurface) StopFind(browserHandle APICall) {
	C.SteamCAPI_ISteamHTMLSurface_StopFind(surface.Pointer, C.HHTMLBrowser(browserHandle))
}

// GetLinkAtPosition returns details about the link at position x, y on the
// current page.
func (surface IHTMLSurface) GetLinkAtPosition(browserHandle APICall, x, y int32) {
	C.SteamCAPI_ISteamHTMLSurface_GetLinkAtPosition(surface.Pointer, C.HHTMLBrowser(browserHandle), C.int(x), C.int(y))
}

// SetCookie sets a webcookie for the hostname in question.
func (surface IHTMLSurface) SetCookie(hostname, key, value, path string, expires RTime32, secure, HTTPOnly bool) {
	chostname := C.CString(hostname)
	ckey := C.CString(key)
	cvalue := C.CString(value)
	cpath := C.CString(path)

	C.SteamCAPI_ISteamHTMLSurface_SetCookie(surface.Pointer, chostname, ckey, cvalue, cpath, C.RTime32(expires), C.bool(secure), C.bool(HTTPOnly))

	C.free(unsafe.Pointer(chostname))
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(cvalue))
	C.free(unsafe.Pointer(cpath))
}

// SetPageScaleFactor zooms the current page by zoom ( from 0.0 to 2.0, so to
// zoom to 120% use 1.2 ), zooming around point x, y in the page (use 0,0 if you
// don't care).
func (surface IHTMLSurface) SetPageScaleFactor(browserHandle APICall, zoom float32, x, y int32) {
	C.SteamCAPI_ISteamHTMLSurface_SetPageScaleFactor(surface.Pointer, C.HHTMLBrowser(browserHandle), C.float(zoom), C.int(x), C.int(y))
}

// SetBackgroundMode enables/disables low-resource background mode, where
// javascript and repaint timers are throttled, resources are more aggressively
// purged from memory, and audio/video elements are paused. When background mode
// is enabled, all HTML5 video and audio objects will execute ".pause()" and
// gain the property "._steam_background_paused = 1". When background mode is
// disabled, any video or audio objects with that property will resume with
// ".play()".
func (surface IHTMLSurface) SetBackgroundMode(browserHandle APICall, backgroundMode bool) {
	C.SteamCAPI_ISteamHTMLSurface_SetBackgroundMode(surface.Pointer, C.HHTMLBrowser(browserHandle), C.bool(backgroundMode))
}

// AllowStartRequest *must* be called in response to a HTML_StartRequest_t
// callback. Set allowed to true to allow this navigation, false to cancel it
// and stay on the current page. You can use this feature to limit the valid
// pages allowed in your HTML surface.
func (surface IHTMLSurface) AllowStartRequest(browserHandle APICall, allowed bool) {
	C.SteamCAPI_ISteamHTMLSurface_AllowStartRequest(surface.Pointer, C.HHTMLBrowser(browserHandle), C.bool(allowed))
}

// JSDialogResponse *must* be called in response to a HTML_JSAlert_t or
// HTML_JSConfirm_t callback. Set result to true for the OK option of a confirm,
// use false otherwise.
func (surface IHTMLSurface) JSDialogResponse(browserHandle APICall, result bool) {
	C.SteamCAPI_ISteamHTMLSurface_JSDialogResponse(surface.Pointer, C.HHTMLBrowser(browserHandle), C.bool(result))
}

/*
// FileLoadDialogResponse *must* be called in response to a
// HTML_FileOpenDialog_t callback.
func (surface IHTMLSurface) FileLoadDialogResponse(browserHandle APICall) {
	chostname := C.CString(hostname)

	C.SteamCAPI_ISteamHTMLSurface_FileLoadDialogResponse(surface.Pointer, C.HHTMLBrowser(browserHandle), char**pchSelectedFiles)

	C.free(unsafe.Pointer(chostname))
}
*/
