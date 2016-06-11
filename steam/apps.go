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

// ISteamApps is a handler for the SteamApp API.
type ISteamApps struct {
	unsafe.Pointer
}

// Apps return the default steam ISteamApp handle.
func Apps() ISteamApps {
	return ISteamApps{C.SteamCAPI_SteamApps()}
}

// IsSubscribed has no valve documentation.
func (a ISteamApps) IsSubscribed() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsSubscribed(a.Pointer))
}

// IsLowViolence has no valve documentation.
func (a ISteamApps) IsLowViolence() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsLowViolence(a.Pointer))
}

// IsCybercafe has no valve documentation.
func (a ISteamApps) IsCybercafe() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsCybercafe(a.Pointer))
}

// IsVACBanned has no valve documentation.
func (a ISteamApps) IsVACBanned() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsVACBanned(a.Pointer))
}

// GetCurrentGameLanguage has no valve documentation.
func (a ISteamApps) GetCurrentGameLanguage() string {
	ret := C.SteamCAPI_ISteamApps_GetCurrentGameLanguage(a.Pointer)
	out := C.GoString(ret)
	return out
}

// GetAvailableGameLanguages has no valve documentation.
func (a ISteamApps) GetAvailableGameLanguages() string {
	ret := C.SteamCAPI_ISteamApps_GetAvailableGameLanguages(a.Pointer)
	out := C.GoString(ret)
	return out
}

// IsSubscribedApp check ownership of another game related to yours, a demo for
// example.
func (a ISteamApps) IsSubscribedApp(appID AppID) bool {
	return bool(C.SteamCAPI_ISteamApps_BIsSubscribedApp(a.Pointer, C.AppId_t(appID)))
}

// IsDlcInstalled takes AppID of DLC and checks if the user owns the DLC and if
// the DLC is installed.
func (a ISteamApps) IsDlcInstalled(appID AppID) bool {
	return bool(C.SteamCAPI_ISteamApps_BIsDlcInstalled(a.Pointer, C.AppId_t(appID)))
}

// GetEarliestPurchaseUnixTime returns the Unix time of the purchase of the app.
func (a ISteamApps) GetEarliestPurchaseUnixTime(nAppID AppID) uint32 {
	return uint32(C.SteamCAPI_ISteamApps_GetEarliestPurchaseUnixTime(a.Pointer, C.AppId_t(nAppID)))
}

// IsSubscribedFromFreeWeekend checks if the user is subscribed to the current
// app through a free weekend. This function will return false for users who
// have a retail or other type of license. Before using, please ask your Valve
// technical contact how to package and secure your free weekend.
func (a ISteamApps) IsSubscribedFromFreeWeekend() bool {
	return bool(C.SteamCAPI_ISteamApps_BIsSubscribedFromFreeWeekend(a.Pointer))
}

// GetDLCCount returns the number of DLC pieces for the running app.
func (a ISteamApps) GetDLCCount() int32 {
	return int32(C.SteamCAPI_ISteamApps_GetDLCCount(a.Pointer))
}

// GetDLCDataByIndex returns metadata for DLC by index, of range [0, GetDLCCount()]
func (a ISteamApps) GetDLCDataByIndex(iDLC int32, pAppID *AppID, pbAvailable *bool, pchName string, cchNameBufferSize int32) bool {
	cpchName := C.CString(pchName)
	defer C.free(unsafe.Pointer(cpchName))
	return bool(C.SteamCAPI_ISteamApps_BGetDLCDataByIndex(a.Pointer, C.int(iDLC), (*C.AppId_t)(pAppID), (*C._Bool)(pbAvailable), cpchName, C.int(cchNameBufferSize)))
}

// InstallDLC is the install control for optional DLC.
func (a ISteamApps) InstallDLC(nAppID AppID) {
	C.SteamCAPI_ISteamApps_InstallDLC(a.Pointer, C.AppId_t(nAppID))
}

// UninstallDLC is the uninstall control for optional DLC.
func (a ISteamApps) UninstallDLC(nAppID AppID) {
	C.SteamCAPI_ISteamApps_UninstallDLC(a.Pointer, C.AppId_t(nAppID))
}

// RequestAppProofOfPurchaseKey requests cd-key for yourself or owned DLC. If
// you are interested in this data then make sure you provide us with a list of
// valid keys to be distributed to users when they purchase the game, before the
// game ships. You'll receive an AppProofOfPurchaseKeyResponse_t callback when
// the key is available (which may be immediately).
func (a ISteamApps) RequestAppProofOfPurchaseKey(nAppID AppID) {
	C.SteamCAPI_ISteamApps_RequestAppProofOfPurchaseKey(a.Pointer, C.AppId_t(nAppID))
}

// GetCurrentBetaName returns current beta branch name, 'public' is the default
// branch.
func (a ISteamApps) GetCurrentBetaName(pchName string, cchNameBufferSize int32) bool {
	cpchName := C.CString(pchName)
	defer C.free(unsafe.Pointer(cpchName))
	return bool(C.SteamCAPI_ISteamApps_GetCurrentBetaName(a.Pointer, cpchName, C.int(cchNameBufferSize)))
}

// MarkContentCorrupt signals Steam that game files seems corrupt or missing.
func (a ISteamApps) MarkContentCorrupt(bMissingFilesOnly bool) bool {
	return bool(C.SteamCAPI_ISteamApps_MarkContentCorrupt(a.Pointer, C._Bool(bMissingFilesOnly)))
}

// GetInstalledDepots returns installed depots in mount order
func (a ISteamApps) GetInstalledDepots(appID AppID, pvecDepots *DepotID, cMaxDepots uint32) uint32 {
	return uint32(C.SteamCAPI_ISteamApps_GetInstalledDepots(a.Pointer, C.AppId_t(appID), (*C.DepotId_t)(pvecDepots), C.uint(cMaxDepots)))
}

// GetAppInstallDir returns current app install folder for AppID, returns folder
// name length.
func (a ISteamApps) GetAppInstallDir(appID AppID, pchFolder string, cchFolderBufferSize uint32) uint32 {
	cpchFolder := C.CString(pchFolder)
	defer C.free(unsafe.Pointer(cpchFolder))
	return uint32(C.SteamCAPI_ISteamApps_GetAppInstallDir(a.Pointer, C.AppId_t(appID), cpchFolder, C.uint(cchFolderBufferSize)))
}

// IsAppInstalled returns true if that app is installed (not necessarily owned).
func (a ISteamApps) IsAppInstalled(appID AppID) bool {
	return bool(C.SteamCAPI_ISteamApps_BIsAppInstalled(a.Pointer, C.AppId_t(appID)))
}

/*
// GetAppOwner returns the SteamID of the original owner. If different from
// current user, it's borrowed.
func (a ISteamApps) GetAppOwner() CSteamID {
	return C.SteamCAPI_ISteamApps_GetAppOwner(a.Pointer)
}*/

// GetLaunchQueryParam returns the associated launch param if the game is run
// via steam://run/<appid>//?param1=value1;param2=value2;param3=value3 etc.
// Parameter names starting with the character '@' are reserved for internal use
// and will always return and empty string. Parameter names starting with an
// underscore '_' are reserved for steam features -- they can be queried by the
// game, but it is advised that you not param names beginning with an underscore
// for your own features.
func (a ISteamApps) GetLaunchQueryParam(pchKey string) string {
	cpchKey := C.CString(pchKey)
	defer C.free(unsafe.Pointer(cpchKey))
	ret := C.SteamCAPI_ISteamApps_GetLaunchQueryParam(a.Pointer, cpchKey)
	out := C.GoString(ret)
	return out
}

// GetDlcDownloadProgress gets download progress for optional DLC.
func (a ISteamApps) GetDlcDownloadProgress(nAppID AppID, punBytesDownloaded *uint64, punBytesTotal *uint64) bool {
	return bool(C.SteamCAPI_ISteamApps_GetDlcDownloadProgress(a.Pointer, C.AppId_t(nAppID), (*C.ulonglong)(punBytesDownloaded), (*C.ulonglong)(punBytesTotal)))
}

// GetAppBuildID return the buildid of this app, may change at any time based on
// backend updates to the game.
func (a ISteamApps) GetAppBuildID() int32 {
	return int32(C.SteamCAPI_ISteamApps_GetAppBuildId(a.Pointer))
}
