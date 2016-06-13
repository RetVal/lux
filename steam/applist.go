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

// IAppList is a handler for the SteamAppList API.
type IAppList struct {
	unsafe.Pointer
}

// AppList return the default steam AppList interface handle.
func AppList() IAppList {
	return IAppList{C.CSteamAppList()}
}

// GetNumInstalledApps returns the number of apps currently installed by this
// steam user.
func (a IAppList) GetNumInstalledApps() int {
	return int(C.SteamAppList_GetNumInstalledApps(a.Pointer))
}

// GetInstalledApps fills apps with the AppId of the apps installed. and return
// the subslice filled.
func (a IAppList) GetInstalledApps(apps []AppID) []AppID {
	return apps[:C.SteamAppList_GetInstalledApps(a.Pointer, (*C.AppId_t)(&apps[0]), C.uint(len(apps)))]
}

// GetAppName returns -1 if no name was found. steamworks will fill the given
// slice as much as possible and return the string (either all of or part of the
// byte slice) representing the name.
func (a IAppList) GetAppName(ID AppID, name []byte) string {
	return string(name[:C.SteamAppList_GetAppName(a.Pointer, C.AppId_t(ID), (*C.char)(unsafe.Pointer(&name[0])), C.int(len(name)))])
}

// GetAppInstallDir returns -1 if no dir was found. steamworks will fill the
// given slice as much as possible and return the string (either all of or part of the
// byte slice) representing the directory.
func (a IAppList) GetAppInstallDir(ID AppID, directory []byte) string {
	return string(directory[:C.SteamAppList_GetAppInstallDir(a.Pointer, C.AppId_t(ID), (*C.char)(unsafe.Pointer(&directory[0])), C.int(len(directory)))])
}

// GetAppBuildID return the buildid of this app, may change at any time based on
// backend updates to the game.
func (a IAppList) GetAppBuildID(ID AppID) int32 {
	return int32(C.SteamAppList_GetAppBuildId(a.Pointer, C.AppId_t(ID)))
}
