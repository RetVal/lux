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

// DepotID represents the steam depot id.
type DepotID uint32

// Apps is a handler for the Steam Apps API.
type Apps struct {
	unsafe.Pointer
}

// IApps returns the steam apps interface handle.
func IApps() Apps {
	return Apps{C.SteamCAPI_SteamApps()}
}

// IsSubscribed has no valve documentation.
func (a Apps) IsSubscribed() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsSubscribed(a.Pointer))
}

// IsLowViolence has no valve documentation.
func (a Apps) IsLowViolence() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsLowViolence(a.Pointer))
}

// IsCybercafe has no valve documentation.
func (a Apps) IsCybercafe() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsCybercafe(a.Pointer))
}

// IsVACBanned has no valve documentation.
func (a Apps) IsVACBanned() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsVACBanned(a.Pointer))
}

// GetCurrentGameLanguage has no valve documentation.
func (a Apps) GetCurrentGameLanguage() string {
	ret := C.SteamCAPI_ISteamApps_GetCurrentGameLanguage(a.Pointer)
	// Are we suppose to free it ?
	out := C.GoString(ret)
	return out
}

// GetAvailableGameLanguages has no valve documentation.
func (a Apps) GetAvailableGameLanguages() string {
	ret := C.SteamCAPI_ISteamApps_GetAvailableGameLanguages(a.Pointer)
	// Are we suppose to free it ?
	out := C.GoString(ret)
	return out
}

// IsSubscribedApp check ownership of another game related to yours, a demo for
// example.
func (a Apps) IsSubscribedApp(ID AppID) bool {
	return bool(C.SteamCAPI_ISteamApps_BIsSubscribedApp(a.Pointer, C.AppId_t(ID)))
}

// IsDLCInstalled takes AppID of DLC and checks if the user owns the DLC and if
// the DLC is installed.
func (a Apps) IsDLCInstalled(ID AppID) bool {
	return bool(C.SteamCAPI_ISteamApps_BIsDlcInstalled(a.Pointer, C.AppId_t(ID)))
}

// GetEarliestPurchaseUnixTime returns the Unix time of the purchase of the app.
func (a Apps) GetEarliestPurchaseUnixTime(ID AppID) uint32 {
	return uint32(C.SteamCAPI_ISteamApps_GetEarliestPurchaseUnixTime(a.Pointer, C.AppId_t(ID)))
}

// IsSubscribedFromFreeWeekend checks if the user is subscribed to the current
// app through a free weekend. This function will return false for users who
// have a retail or other type of license. Before using, please ask your Valve
// technical contact how to package and secure your free weekend.
func (a Apps) IsSubscribedFromFreeWeekend() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsSubscribedFromFreeWeekend(a.Pointer))
}

// GetDLCCount returns the number of DLC pieces for the running app.
func (a Apps) GetDLCCount() int {
	return int(C.SteamCAPI_ISteamApps_GetDLCCount(a.Pointer))
}

// GetDLCDataByIndex returns metadata for DLC by index, of range [0, GetDLCCount()]
func (a Apps) GetDLCDataByIndex(iDLC int, ID *AppID, available *bool, metadata []byte) bool {
	return bool(C.SteamCAPI_ISteamApps_BGetDLCDataByIndex(a.Pointer, C.int(iDLC), (*C.AppId_t)(ID), (*C._Bool)(available), (*C.char)(unsafe.Pointer(&metadata[0])), C.int(len(metadata))))
}

// InstallDLC is the install control for optional DLC.
func (a Apps) InstallDLC(ID AppID) {
	C.SteamCAPI_ISteamApps_InstallDLC(a.Pointer, C.AppId_t(ID))
}

// UninstallDLC is the uninstall control for optional DLC.
func (a Apps) UninstallDLC(ID AppID) {
	C.SteamCAPI_ISteamApps_UninstallDLC(a.Pointer, C.AppId_t(ID))
}

// RequestAppProofOfPurchaseKey requests cd-key for yourself or owned DLC. If
// you are interested in this data then make sure you provide us with a list of
// valid keys to be distributed to users when they purchase the game, before the
// game ships. You'll receive an AppProofOfPurchaseKeyResponse_t callback when
// the key is available (which may be immediately).
func (a Apps) RequestAppProofOfPurchaseKey(ID AppID) {
	C.SteamCAPI_ISteamApps_RequestAppProofOfPurchaseKey(a.Pointer, C.AppId_t(ID))
}

// GetCurrentBetaName returns current beta branch name, 'public' is the default
// branch.
func (a Apps) GetCurrentBetaName(name []byte) bool {
	return bool(C.SteamCAPI_ISteamApps_GetCurrentBetaName(a.Pointer, (*C.char)(unsafe.Pointer(&name[0])), C.int(len(name))))
}

// MarkContentCorrupt signals Steam that game files seems corrupt or missing.
func (a Apps) MarkContentCorrupt(missingFilesOnly bool) bool {
	return bool(C.SteamCAPI_ISteamApps_MarkContentCorrupt(a.Pointer, C._Bool(missingFilesOnly)))
}

// GetInstalledDepots returns installed depots in mount order.
func (a Apps) GetInstalledDepots(ID AppID, depots []DepotID) uint32 {
	return uint32(C.SteamCAPI_ISteamApps_GetInstalledDepots(a.Pointer, C.AppId_t(ID), (*C.DepotId_t)(&depots[0]), C.uint(len(depots))))
}

// GetAppInstallDir returns current app install folder for AppID, returns folder
// name length.
func (a Apps) GetAppInstallDir(ID AppID, folder []byte) string {
	return string(folder[:C.SteamCAPI_ISteamApps_GetAppInstallDir(a.Pointer, C.AppId_t(ID), (*C.char)(unsafe.Pointer(&folder[0])), C.uint(len(folder)))])
}

// IsAppInstalled returns true if that app is installed (not necessarily owned).
func (a Apps) IsAppInstalled(ID AppID) bool {
	return bool(C.SteamCAPI_ISteamApps_BIsAppInstalled(a.Pointer, C.AppId_t(ID)))
}

// GetAppOwner returns the SteamID of the original owner. If different from
// current user, it's borrowed.
func (a Apps) GetAppOwner() SteamID {
	ret := C.SteamCAPI_ISteamApps_GetAppOwner(a.Pointer)
	return *(*SteamID)(unsafe.Pointer(&ret))
}

// GetLaunchQueryParam returns the associated launch param if the game is run
// via steam://run/<appid>//?param1=value1;param2=value2;param3=value3 etc.
// Parameter names starting with the character '@' are reserved for internal use
// and will always return and empty string. Parameter names starting with an
// underscore '_' are reserved for steam features -- they can be queried by the
// game, but it is advised that you not param names beginning with an underscore
// for your own features.
func (a Apps) GetLaunchQueryParam(key string) string {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	ret := C.SteamCAPI_ISteamApps_GetLaunchQueryParam(a.Pointer, ckey)
	out := C.GoString(ret)
	return out
}

// GetDLCDownloadProgress gets download progress for optional DLC.
func (a Apps) GetDLCDownloadProgress(ID AppID, bytesDownloaded *uint64, bytesTotal *uint64) bool {
	return bool(C.SteamCAPI_ISteamApps_GetDlcDownloadProgress(a.Pointer, C.AppId_t(ID), (*C.ulonglong)(bytesDownloaded), (*C.ulonglong)(bytesTotal)))
}

// GetAppBuildID return the buildid of this app, may change at any time based on
// backend updates to the game.
func (a Apps) GetAppBuildID() int32 {
	return int32(C.SteamCAPI_ISteamApps_GetAppBuildId(a.Pointer))
}
