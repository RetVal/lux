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

// Pipe is a handle to a communication pipe to the Steam client.
type Pipe int32

// HSteamUser is a handle to single instance of a steam user.
type HSteamUser int32

// AccountType is an enum for all the possible account type.
type AccountType int32

// Client is a handler for the SteamClient API.
type Client struct {
	unsafe.Pointer
}

// IClient returns steam default Client handle.
func IClient() Client {
	return Client{C.SteamCAPI_SteamClient()}
}

// CreatePipe creates a communication pipe to the Steam client.
func (c Client) CreatePipe() Pipe {
	return Pipe(C.SteamCAPI_SteamClient_CreateSteamPipe(c.Pointer))
}

// ReleasePipe releases a previously created communications pipe.
func (c Client) ReleasePipe(pipe Pipe) bool {
	return bool(C.SteamCAPI_SteamClient_BReleaseSteamPipe(c.Pointer, C.HSteamPipe(pipe)))
}

// ConnectToGlobalUser connects to an existing global user, failing if none
// exists used by the game to coordinate with the steamUI.
func (c Client) ConnectToGlobalUser(pipe Pipe) HSteamUser {
	return HSteamUser(C.SteamCAPI_SteamClient_ConnectToGlobalUser(c.Pointer, C.HSteamPipe(pipe)))
}

// CreateLocalUser is used by game servers, it creates a steam user that won't
// be shared with anyone else.
func (c Client) CreateLocalUser(pipe *Pipe, accountType AccountType) HSteamUser {
	return HSteamUser(C.SteamCAPI_SteamClient_CreateLocalUser(c.Pointer, (*C.HSteamPipe)(pipe), C.EAccountType(accountType)))
}

// ReleaseUser removes an allocated user.
func (c Client) ReleaseUser(pipe Pipe, user HSteamUser) {
	C.SteamCAPI_SteamClient_ReleaseUser(c.Pointer, C.HSteamPipe(pipe), C.HSteamUser(user))
}

// GetISteamUser retrieves the ISteamUser interface associated with the handle.
func (c Client) GetISteamUser(user HSteamUser, pipe Pipe, version string) IUser {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IUser{C.SteamCAPI_SteamClient_GetISteamUser(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamGameServer retrieves the ISteamGameServer interface associated with
// the handle.
func (c Client) GetISteamGameServer(user HSteamUser, pipe Pipe, version string) IGameServer {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IGameServer{C.SteamCAPI_SteamClient_GetISteamGameServer(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// SetLocalIPBinding set the local IP and Port to bind to this must be set
// before CreateLocalUser().
func (c Client) SetLocalIPBinding(IP uint32, port uint16) {
	C.SteamCAPI_SteamClient_SetLocalIPBinding(c.Pointer, C.uint(IP), C.ushort(port))
}

// GetISteamFriends returns the ISteamFriends interface.
func (c Client) GetISteamFriends(user HSteamUser, pipe Pipe, version string) IFriends {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IFriends{C.SteamCAPI_SteamClient_GetISteamFriends(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamUtils returns the ISteamUtils interface.
func (c Client) GetISteamUtils(pipe Pipe, version string) IUtils {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IUtils{C.SteamCAPI_SteamClient_GetISteamUtils(c.Pointer, C.HSteamPipe(pipe), cversion)}
}

// GetISteamMatchmaking returns the ISteamMatchmaking interface.
func (c Client) GetISteamMatchmaking(user HSteamUser, pipe Pipe, version string) IMatchmaking {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IMatchmaking{C.SteamCAPI_SteamClient_GetISteamMatchmaking(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamMatchmakingServers returns the ISteamMatchmakingServers interface.
func (c Client) GetISteamMatchmakingServers(user HSteamUser, pipe Pipe, version string) IMatchmakingServers {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IMatchmakingServers{C.SteamCAPI_SteamClient_GetISteamMatchmakingServers(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamGenericInterface returns the a generic interface.
func (c Client) GetISteamGenericInterface(user HSteamUser, pipe Pipe, version string) IGenericInterface {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IGenericInterface{C.SteamCAPI_SteamClient_GetISteamGenericInterface(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamUserStats returns the ISteamUserStats interface.
func (c Client) GetISteamUserStats(user HSteamUser, pipe Pipe, version string) IUserStats {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IUserStats{C.SteamCAPI_SteamClient_GetISteamUserStats(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamGameServerStats returns the ISteamGameServerStats interface.
func (c Client) GetISteamGameServerStats(user HSteamUser, pipe Pipe, version string) IGameServerStats {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IGameServerStats{C.SteamCAPI_SteamClient_GetISteamGameServerStats(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamApps returns apps interface.
func (c Client) GetISteamApps(user HSteamUser, pipe Pipe, version string) Apps {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return Apps{C.SteamCAPI_SteamClient_GetISteamApps(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamNetworking returns the networking interface.
func (c Client) GetISteamNetworking(user HSteamUser, pipe Pipe, version string) INetworking {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return INetworking{C.SteamCAPI_SteamClient_GetISteamNetworking(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamRemoteStorage returns the remote storage interface.
func (c Client) GetISteamRemoteStorage(user HSteamUser, pipe Pipe, version string) IRemoteStorage {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IRemoteStorage{C.SteamCAPI_SteamClient_GetISteamRemoteStorage(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamScreenshots returns the user screenshots interface.
func (c Client) GetISteamScreenshots(user HSteamUser, pipe Pipe, version string) IScreenshots {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IScreenshots{C.SteamCAPI_SteamClient_GetISteamScreenshots(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// RunFrame needs to be called every frame to process matchmaking results
// redundant if you're already calling SteamAPI_RunCallbacks().
func (c Client) RunFrame() {
	C.SteamCAPI_SteamClient_RunFrame(c.Pointer)
}

// GetIPCCallCount returns the number of IPC calls made since the last time this
// function was called Used for perf debugging so you can understand how many
// IPC calls your game makes per frame Every IPC call is at minimum a thread
// context switch if not a process one so you want to rate control how often
// you do them.
func (c Client) GetIPCCallCount() uint32 {
	return uint32(C.SteamCAPI_SteamClient_GetIPCCallCount(c.Pointer))
}

/*
// SetWarningMessageHook is for API warning handling 'int' is the severity; 0
// for msg, 1 for warning 'const char *' is the text of the message callbacks
// will occur directly after the API function is called that generated the
// warning or message.
func (c Client) SetWarningMessageHook(pFunction SteamAPIWarningMessageHook_t) {
	C.SteamCAPI_SteamClient_SetWarningMessageHook(c.Pointer, pFunction)
}
*/

// ShutdownIfAllPipesClosed triggers global shutdown for the DLL.
func (c Client) ShutdownIfAllPipesClosed() bool {
	return bool(C.SteamCAPI_SteamClient_BShutdownIfAllPipesClosed(c.Pointer))
}

// GetISteamHTTP returns the HTTP interface.
func (c Client) GetISteamHTTP(user HSteamUser, pipe Pipe, version string) IHTTP {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IHTTP{C.SteamCAPI_SteamClient_GetISteamHTTP(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamUnifiedMessages returns the UnifiedMessages interface.
func (c Client) GetISteamUnifiedMessages(user HSteamUser, pipe Pipe, version string) IUnifiedMessages {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IUnifiedMessages{C.SteamCAPI_SteamClient_GetISteamUnifiedMessages(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamController returns the controller interface.
func (c Client) GetISteamController(user HSteamUser, pipe Pipe, version string) IController {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IController{C.SteamCAPI_SteamClient_GetISteamController(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamUGC returns the UGC interface.
func (c Client) GetISteamUGC(user HSteamUser, pipe Pipe, version string) IUGC {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IUGC{C.SteamCAPI_SteamClient_GetISteamUGC(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamAppList returns app list interface, this is only available on
// specially registered apps.
func (c Client) GetISteamAppList(user HSteamUser, pipe Pipe, version string) IAppList {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IAppList{C.SteamCAPI_SteamClient_GetISteamAppList(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamMusic returns the Music Player interface.
func (c Client) GetISteamMusic(user HSteamUser, pipe Pipe, version string) IMusic {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IMusic{C.SteamCAPI_SteamClient_GetISteamMusic(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamMusicRemote returns the Music Player Remote interface.
func (c Client) GetISteamMusicRemote(user HSteamUser, pipe Pipe, version string) IMusicRemote {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return IMusicRemote{C.SteamCAPI_SteamClient_GetISteamMusicRemote(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion)}
}

// GetISteamHTMLSurface returns the html page display interface.
func (c Client) GetISteamHTMLSurface(user HSteamUser, steamPipe Pipe, version string) IHTMLSurface {
	cversion := C.CString(version)
	ret := IHTMLSurface{C.SteamCAPI_SteamClient_GetISteamHTMLSurface(c.Pointer, C.HSteamUser(user), C.HSteamPipe(steamPipe), cversion)}
	C.free(unsafe.Pointer(cversion))
	return ret
}

/*
// Helper functions for internal Steam usage

func (c Client) Set_SteamAPI_CPostAPIResultInProcess(xfunc SteamAPI_PostAPIResultInProcess_t) {
	C.SteamCAPI_SteamClient_Set_SteamAPI_CPostAPIResultInProcess(c.Pointer, xfunc)
}

func (c Client) Remove_SteamAPI_CPostAPIResultInProcess(xfunc SteamAPI_PostAPIResultInProcess_t) {
	C.SteamCAPI_SteamClient_Remove_SteamAPI_CPostAPIResultInProcess(c.Pointer, xfunc)
}

func (c Client) Set_SteamAPI_CCheckCallbackRegisteredInProcess(xfunc SteamAPI_CheckCallbackRegistered_t) {
	C.SteamCAPI_SteamClient_Set_SteamAPI_CCheckCallbackRegisteredInProcess(c.Pointer, xfunc)
}
*/

// GetISteamInventory returns the inventory interface.
func (c Client) GetISteamInventory(user HSteamUser, pipe Pipe, version string) unsafe.Pointer {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return unsafe.Pointer(C.SteamCAPI_SteamClient_GetISteamInventory(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion))
}

// GetISteamVideo returns the Video interface.
func (c Client) GetISteamVideo(user HSteamUser, pipe Pipe, version string) unsafe.Pointer {
	cversion := C.CString(version)
	defer C.free(unsafe.Pointer(cversion))
	return unsafe.Pointer(C.SteamCAPI_SteamClient_GetISteamVideo(c.Pointer, C.HSteamUser(user), C.HSteamPipe(pipe), cversion))
}
