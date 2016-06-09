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

// HSteamPipe is a handle to a communication pipe to the Steam client.
type HSteamPipe int32

// HSteamUser is a handle to single instance of a steam user.
type HSteamUser int32
type EAccountType int32

// ISteamClient is a handler for the SteamClient API.
type ISteamClient struct {
	unsafe.Pointer
}

// Client returns steam default ISteamClient handle.
func Client() ISteamClient {
	return ISteamClient{C.SteamCAPI_SteamClient()}
}

// CreateSteamPipe creates a communication pipe to the Steam client.
func (c ISteamClient) CreateSteamPipe() HSteamPipe {
	return HSteamPipe(C.SteamCAPI_SteamClient_CreateSteamPipe(c.Pointer))
}

// ReleaseSteamPipe releases a previously created communications pipe.
func (c ISteamClient) ReleaseSteamPipe(hSteamPipe HSteamPipe) bool {
	return bool(C.SteamCAPI_SteamClient_BReleaseSteamPipe(c.Pointer, C.HSteamPipe(hSteamPipe)))
}

// ConnectToGlobalUser connects to an existing global user, failing if none
// exists used by the game to coordinate with the steamUI.
func (c ISteamClient) ConnectToGlobalUser(hSteamPipe HSteamPipe) HSteamUser {
	return HSteamUser(C.SteamCAPI_SteamClient_ConnectToGlobalUser(c.Pointer, C.HSteamPipe(hSteamPipe)))
}

// CreateLocalUser is used by game servers, it creates a steam user that won't
// be shared with anyone else.
func (c ISteamClient) CreateLocalUser(phSteamPipe *HSteamPipe, eAccountType EAccountType) HSteamUser {
	return HSteamUser(C.SteamCAPI_SteamClient_CreateLocalUser(c.Pointer, (*C.HSteamPipe)(phSteamPipe), C.EAccountType(eAccountType)))
}

// ReleaseUser removes an allocated user.
func (c ISteamClient) ReleaseUser(hSteamPipe HSteamPipe, hUser HSteamUser) {
	C.SteamCAPI_SteamClient_ReleaseUser(c.Pointer, C.HSteamPipe(hSteamPipe), C.HSteamUser(hUser))
}

// GetISteamUser retrieves the ISteamUser interface associated with the handle.
func (c ISteamClient) GetISteamUser(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IUser {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IUser{C.SteamCAPI_SteamClient_GetISteamUser(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamGameServer retrieves the ISteamGameServer interface associated with
// the handle.
func (c ISteamClient) GetISteamGameServer(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IGameServer {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IGameServer{C.SteamCAPI_SteamClient_GetISteamGameServer(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// SetLocalIPBinding set the local IP and Port to bind to this must be set
// before CreateLocalUser().
func (c ISteamClient) SetLocalIPBinding(unIP uint32, usPort uint16) {
	C.SteamCAPI_SteamClient_SetLocalIPBinding(c.Pointer, C.uint(unIP), C.ushort(usPort))
}

// GetISteamFriends returns the ISteamFriends interface.
func (c ISteamClient) GetISteamFriends(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IFriends {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IFriends{C.SteamCAPI_SteamClient_GetISteamFriends(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamUtils returns the ISteamUtils interface.
func (c ISteamClient) GetISteamUtils(hSteamPipe HSteamPipe, pchVersion string) IUtils {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IUtils{C.SteamCAPI_SteamClient_GetISteamUtils(c.Pointer, C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamMatchmaking returns the ISteamMatchmaking interface.
func (c ISteamClient) GetISteamMatchmaking(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IMatchmaking {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IMatchmaking{C.SteamCAPI_SteamClient_GetISteamMatchmaking(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamMatchmakingServers returns the ISteamMatchmakingServers interface.
func (c ISteamClient) GetISteamMatchmakingServers(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IMatchmakingServers {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IMatchmakingServers{C.SteamCAPI_SteamClient_GetISteamMatchmakingServers(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamGenericInterface returns the a generic interface.
func (c ISteamClient) GetISteamGenericInterface(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IGenericInterface {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IGenericInterface{C.SteamCAPI_SteamClient_GetISteamGenericInterface(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamUserStats returns the ISteamUserStats interface.
func (c ISteamClient) GetISteamUserStats(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IUserStats {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IUserStats{C.SteamCAPI_SteamClient_GetISteamUserStats(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamGameServerStats returns the ISteamGameServerStats interface.
func (c ISteamClient) GetISteamGameServerStats(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IGameServerStats {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IGameServerStats{C.SteamCAPI_SteamClient_GetISteamGameServerStats(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamApps returns apps interface.
func (c ISteamClient) GetISteamApps(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IApps {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IApps{C.SteamCAPI_SteamClient_GetISteamApps(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamNetworking returns the networking interface.
func (c ISteamClient) GetISteamNetworking(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) INetworking {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return INetworking{C.SteamCAPI_SteamClient_GetISteamNetworking(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamRemoteStorage returns the remote storage interface.
func (c ISteamClient) GetISteamRemoteStorage(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IRemoteStorage {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IRemoteStorage{C.SteamCAPI_SteamClient_GetISteamRemoteStorage(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamScreenshots returns the user screenshots interface.
func (c ISteamClient) GetISteamScreenshots(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IScreenshots {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IScreenshots{C.SteamCAPI_SteamClient_GetISteamScreenshots(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// RunFrame needs to be called every frame to process matchmaking results
// redundant if you're already calling SteamAPI_RunCallbacks().
func (c ISteamClient) RunFrame() {
	C.SteamCAPI_SteamClient_RunFrame(c.Pointer)
}

// GetIPCCallCount returns the number of IPC calls made since the last time this
// function was called Used for perf debugging so you can understand how many
// IPC calls your game makes per frame Every IPC call is at minimum a thread
// context switch if not a process one so you want to rate control how often
// you do them.
func (c ISteamClient) GetIPCCallCount() uint32 {
	return uint32(C.SteamCAPI_SteamClient_GetIPCCallCount(c.Pointer))
}

/*
// SetWarningMessageHook is for API warning handling 'int' is the severity; 0
// for msg, 1 for warning 'const char *' is the text of the message callbacks
// will occur directly after the API function is called that generated the
// warning or message.
func (c ISteamClient) SetWarningMessageHook(pFunction SteamAPIWarningMessageHook_t) {
	C.SteamCAPI_SteamClient_SetWarningMessageHook(c.Pointer, pFunction)
}
*/

// ShutdownIfAllPipesClosed triggers global shutdown for the DLL.
func (c ISteamClient) ShutdownIfAllPipesClosed() bool {
	return bool(C.SteamCAPI_SteamClient_BShutdownIfAllPipesClosed(c.Pointer))
}

// GetISteamHTTP returns the HTTP interface.
func (c ISteamClient) GetISteamHTTP(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IHTTP {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IHTTP{C.SteamCAPI_SteamClient_GetISteamHTTP(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamUnifiedMessages returns the UnifiedMessages interface.
func (c ISteamClient) GetISteamUnifiedMessages(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IUnifiedMessages {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IUnifiedMessages{C.SteamCAPI_SteamClient_GetISteamUnifiedMessages(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamController returns the controller interface.
func (c ISteamClient) GetISteamController(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IController {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IController{C.SteamCAPI_SteamClient_GetISteamController(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamUGC returns the UGC interface.
func (c ISteamClient) GetISteamUGC(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IUGC {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IUGC{C.SteamCAPI_SteamClient_GetISteamUGC(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamAppList returns app list interface, this is only available on
// specially registered apps.
func (c ISteamClient) GetISteamAppList(hSteamUser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IAppList {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IAppList{C.SteamCAPI_SteamClient_GetISteamAppList(c.Pointer, C.HSteamUser(hSteamUser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamMusic returns the Music Player interface.
func (c ISteamClient) GetISteamMusic(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IMusic {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IMusic{C.SteamCAPI_SteamClient_GetISteamMusic(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamMusicRemote returns the Music Player Remote interface.
func (c ISteamClient) GetISteamMusicRemote(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IMusicRemote {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IMusicRemote{C.SteamCAPI_SteamClient_GetISteamMusicRemote(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

// GetISteamHTMLSurface returns the html page display interface.
func (c ISteamClient) GetISteamHTMLSurface(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) IHTMLSurface {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return IHTMLSurface{C.SteamCAPI_SteamClient_GetISteamHTMLSurface(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion)}
}

/*
// Helper functions for internal Steam usage

func (c ISteamClient) Set_SteamAPI_CPostAPIResultInProcess(xfunc SteamAPI_PostAPIResultInProcess_t) {
	C.SteamCAPI_SteamClient_Set_SteamAPI_CPostAPIResultInProcess(c.Pointer, xfunc)
}

func (c ISteamClient) Remove_SteamAPI_CPostAPIResultInProcess(xfunc SteamAPI_PostAPIResultInProcess_t) {
	C.SteamCAPI_SteamClient_Remove_SteamAPI_CPostAPIResultInProcess(c.Pointer, xfunc)
}

func (c ISteamClient) Set_SteamAPI_CCheckCallbackRegisteredInProcess(xfunc SteamAPI_CheckCallbackRegistered_t) {
	C.SteamCAPI_SteamClient_Set_SteamAPI_CCheckCallbackRegisteredInProcess(c.Pointer, xfunc)
}
*/
// GetISteamInventory returns the inventory interface.
func (c ISteamClient) GetISteamInventory(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) unsafe.Pointer {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return unsafe.Pointer(C.SteamCAPI_SteamClient_GetISteamInventory(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion))
}

// GetISteamVideo returns the Video interface.
func (c ISteamClient) GetISteamVideo(hSteamuser HSteamUser, hSteamPipe HSteamPipe, pchVersion string) unsafe.Pointer {
	cpchVersion := C.CString(pchVersion)
	defer C.free(unsafe.Pointer(cpchVersion))
	return unsafe.Pointer(C.SteamCAPI_SteamClient_GetISteamVideo(c.Pointer, C.HSteamUser(hSteamuser), C.HSteamPipe(hSteamPipe), cpchVersion))
}
