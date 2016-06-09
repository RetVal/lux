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

// ISteamAppList is a handler for the SteamAppList API.
type ISteamAppList struct {
	unsafe.Pointer
}

// AppList return the default steam ISteamAppList handle.
func AppList() ISteamAppList {
	return ISteamAppList{C.CSteamAppList}
}

// GetNumInstalledApps returns the number of apps currently installed by this
// steam user.
func (a ISteamAppList) GetNumInstalledApps() uint32 {
	return uint32(C.SteamAppList_GetNumInstalledApps(a.Pointer))
}

// GetInstalledApps fills pvecAppID with max unMaxAppIDs with the AppId of the
// apps installed.
func (a ISteamAppList) GetInstalledApps(pvecAppID *AppId, unMaxAppIDs uint32) uint32 {
	return uint32(C.SteamAppList_GetInstalledApps(a.Pointer, (*C.AppId_t)(pvecAppID), C.uint(unMaxAppIDs)))
}

// GetAppName returns -1 if no name was found
// BUG(The strings are return value... so this method is actually wrong)
func (a ISteamAppList) GetAppName(nAppID AppId, pchName string, cchNameMax int32) int32 {
	cpchName := C.CString(pchName)
	defer C.free(unsafe.Pointer(cpchName))
	return int32(C.SteamAppList_GetAppName(a.Pointer, C.AppId_t(nAppID), cpchName, C.int(cchNameMax)))
}

// GetAppInstallDir returns -1 if no dir was found
// BUG(The strings are return value... so this method is actually wrong)
func (a ISteamAppList) GetAppInstallDir(nAppID AppId, pchDirectory string, cchNameMax int32) int32 {
	cpchDirectory := C.CString(pchDirectory)
	defer C.free(unsafe.Pointer(cpchDirectory))
	return int32(C.SteamAppList_GetAppInstallDir(a.Pointer, C.AppId_t(nAppID), cpchDirectory, C.int(cchNameMax)))
}

// GetAppBuildID return the buildid of this app, may change at any time based on
// backend updates to the game
func (a ISteamAppList) GetAppBuildID(nAppID AppId) int32 {
	return int32(C.SteamAppList_GetAppBuildId(a.Pointer, C.AppId_t(nAppID)))
}
