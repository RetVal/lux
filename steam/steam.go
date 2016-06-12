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

// Shutdown terminates the Steam API.
func Shutdown() {
	C.SteamCAPI_Shutdown()
}

// IsSteamRunning checks if a local Steam client is running.
func IsSteamRunning() bool {
	return bool(C.SteamCAPI_IsSteamRunning())
}

// IUser is the interface to the currently connected steam user.
type IUser struct {
	unsafe.Pointer
}

// IGameServer is used to authenticate users via Steam to play on a game server.
type IGameServer struct{ unsafe.Pointer }

// IFriends is an interface to accessing information about individual users,
// that can be a friend, in a group, on a game server or in a lobby with the
// local user.
type IFriends struct{ unsafe.Pointer }

// IUtils is an interface to user independent utility functions.
type IUtils struct{ unsafe.Pointer }

// IMatchmaking holds methods for match making services for clients to get to
// favorites and to operate on game lobbies.
type IMatchmaking struct{ unsafe.Pointer }

// IMatchmakingServers holds functions for match making services for clients to
// get to game lists and details.
type IMatchmakingServers struct{ unsafe.Pointer }

// IGenericInterface I'm not even sure what that is, it's defined in the flat
// API but doesn't seem to be defined in the actual code.
type IGenericInterface struct{ unsafe.Pointer }

// IUserStats holds functions for accessing stats, achievements, and leaderboard
// information.
type IUserStats struct{ unsafe.Pointer }

// IGameServerStats holds functions for authenticating users via Steam to play
// on a game server.
type IGameServerStats struct{ unsafe.Pointer }

// IApps is an interface to app data.
type IApps struct{ unsafe.Pointer }

// INetworking holds functions for making connections and sending data between
// clients, traversing NAT's where possible
type INetworking struct{ unsafe.Pointer }

// IRemoteStorage holds functions for accessing, reading and writing files
// stored remotely and cached locally.
type IRemoteStorage struct{ unsafe.Pointer }

// IScreenshots holds functions for adding screenshots to the user's screenshot
// library.
type IScreenshots struct{ unsafe.Pointer }

// IHTTP is an interface to http client.
type IHTTP struct{ unsafe.Pointer }

// IUnifiedMessages is an interface to unified messages client.
type IUnifiedMessages struct{ unsafe.Pointer }

// IController holds the function for the steam controller API.
type IController struct{ unsafe.Pointer }

// IUGC is an interface to steam ugc.
type IUGC struct{ unsafe.Pointer }

// IAppList is an interface to app data in Steam.
type IAppList struct{ unsafe.Pointer }

// IMusic holds functions to control music playback in the steam client.
type IMusic struct{ unsafe.Pointer }

// IMusicRemote has no documentation. Contact Valve :)
type IMusicRemote struct{ unsafe.Pointer }

// IHTMLSurface holds functions for displaying HTML pages and interacting with
// them.
type IHTMLSurface struct{ unsafe.Pointer }

// User returns the ISteamUser interface for the currently connected user.
func User() IUser {
	return IUser{C.CSteamUser()}
}

// HTMLSurface returns the default HTML surface.
func HTMLSurface() IHTMLSurface {
	return IHTMLSurface{C.CSteamHTMLSurface()}
}
