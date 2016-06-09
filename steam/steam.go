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

// Init returns true to indicates all required interfaces have been acquired and are accessible.
// A return of false indicates one of three conditions:
//     The Steam client isn't running. A running Steam client is required to provide implementations of the various Steamworks interfaces.
//     The Steam client couldn't determine the AppID of game. Make sure you have Steam_appid.txt in your game directory.
//     Your application is not running under the same user context as the Steam client, including admin privileges.
func Init() bool {
	return bool(C.SteamCAPI_Init())
}

func Shutdown() {
	C.SteamCAPI_Shutdown()
}

// IsSteamRunning checks if a local Steam client is running.
func IsSteamRunning() bool {
	return bool(C.SteamCAPI_IsSteamRunning())
}

// ISteamUser is the interface to the currently connected steam user.
type IUser struct {
	unsafe.Pointer
}

type IGameServer struct {
	unsafe.Pointer
}

type IFriends struct{ unsafe.Pointer }
type IUtils struct{ unsafe.Pointer }
type IMatchmaking struct{ unsafe.Pointer }
type IMatchmakingServers struct{ unsafe.Pointer }
type IGenericInterface struct{ unsafe.Pointer }
type IUserStats struct{ unsafe.Pointer }
type IGameServerStats struct{ unsafe.Pointer }
type IApps struct{ unsafe.Pointer }
type INetworking struct{ unsafe.Pointer }
type IRemoteStorage struct{ unsafe.Pointer }
type IScreenshots struct{ unsafe.Pointer }

type IHTTP struct{ unsafe.Pointer }
type IUnifiedMessages struct{ unsafe.Pointer }
type IController struct{ unsafe.Pointer }
type IUGC struct{ unsafe.Pointer }
type IAppList struct{ unsafe.Pointer }
type IMusic struct{ unsafe.Pointer }
type IMusicRemote struct{ unsafe.Pointer }
type IHTMLSurface struct{ unsafe.Pointer }

// SteamUser returns the ISteamUser interface for the currently connected user.
func User() IUser {
	return IUser{C.CSteamUser()}
}
