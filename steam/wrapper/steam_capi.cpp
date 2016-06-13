#include "sdk/public/steam/steam_api.h"
#include "sdk/public/steam/steam_gameserver.h"
#include "sdk/public/steam/isteamgamecoordinator.h"
#include "sdk/public/steam/steamtypes.h"
#include "sdk/public/steam/isteamappticket.h"

extern "C" {

typedef int EHTMLMouseButton;
typedef int EHTMLKeyModifiers;

//==============================================================================
//==================================Steam API==================================
//==============================================================================

void* CSteamUser()        { return SteamUser(); }
void* CSteamHTMLSurface() { return SteamHTMLSurface(); }

bool SteamCAPI_Init()			{ return SteamAPI_Init(); }
void SteamCAPI_Shutdown()		{ 		 SteamAPI_Shutdown(); }
bool SteamCAPI_IsSteamRunning()	{ return SteamAPI_IsSteamRunning(); }


//==============================================================================
//==============================Steam App List API==============================
//==============================================================================
void* CSteamAppList() { return SteamAppList(); }

unsigned int 	SteamAppList_GetNumInstalledApps(void* appList) 													{ return static_cast<ISteamAppList*>(appList)->GetNumInstalledApps(); }
unsigned int 	SteamAppList_GetInstalledApps(void* appList, AppId_t* pvecAppID, unsigned int unMaxAppIDs) 			{ return static_cast<ISteamAppList*>(appList)->GetInstalledApps(pvecAppID, unMaxAppIDs); }
int  			SteamAppList_GetAppName(void* appList, AppId_t nAppID, char *pchName, int cchNameMax) 				{ return static_cast<ISteamAppList*>(appList)->GetAppName(nAppID, pchName, cchNameMax); }
int  			SteamAppList_GetAppInstallDir(void* appList, AppId_t nAppID, char *pchDirectory, int cchNameMax) 	{ return static_cast<ISteamAppList*>(appList)->GetAppInstallDir(nAppID, pchDirectory, cchNameMax); }
int 			SteamAppList_GetAppBuildId(void* appList, AppId_t nAppID) 											{ return static_cast<ISteamAppList*>(appList)->GetAppBuildId(nAppID); }


//======================================================================
//============================Steam Apps API============================
//======================================================================
void* SteamCAPI_SteamApps() { return SteamApps(); }

bool 			SteamCAPI_ISteamApps_BIsSubscribed(void* apps) 																										{ return static_cast<ISteamApps*>(apps)->BIsSubscribed(); }
bool 			SteamCAPI_ISteamApps_BIsLowViolence(void* apps) 																									{ return static_cast<ISteamApps*>(apps)->BIsLowViolence(); }
bool 			SteamCAPI_ISteamApps_BIsCybercafe(void* apps) 																										{ return static_cast<ISteamApps*>(apps)->BIsCybercafe(); }
bool 			SteamCAPI_ISteamApps_BIsVACBanned(void* apps) 																										{ return static_cast<ISteamApps*>(apps)->BIsVACBanned(); }
const char* 	SteamCAPI_ISteamApps_GetCurrentGameLanguage(void* apps) 																							{ return static_cast<ISteamApps*>(apps)->GetCurrentGameLanguage(); }
const char* 	SteamCAPI_ISteamApps_GetAvailableGameLanguages(void* apps) 																							{ return static_cast<ISteamApps*>(apps)->GetAvailableGameLanguages(); }
bool 			SteamCAPI_ISteamApps_BIsSubscribedApp(void* apps,  AppId_t appID) 																					{ return static_cast<ISteamApps*>(apps)->BIsSubscribedApp(appID); }
bool 			SteamCAPI_ISteamApps_BIsDlcInstalled(void* apps,  AppId_t appID) 																					{ return static_cast<ISteamApps*>(apps)->BIsDlcInstalled(appID); }
unsigned int 	SteamCAPI_ISteamApps_GetEarliestPurchaseUnixTime(void* apps,  AppId_t nAppID) 																		{ return static_cast<ISteamApps*>(apps)->GetEarliestPurchaseUnixTime(nAppID); }
bool 			SteamCAPI_ISteamApps_BIsSubscribedFromFreeWeekend(void* apps) 																						{ return static_cast<ISteamApps*>(apps)->BIsSubscribedFromFreeWeekend(); }
int 			SteamCAPI_ISteamApps_GetDLCCount(void* apps) 																										{ return static_cast<ISteamApps*>(apps)->GetDLCCount(); }
bool 			SteamCAPI_ISteamApps_BGetDLCDataByIndex(void* apps,  int iDLC, AppId_t *pAppID, bool *pbAvailable, char* pchName, int cchNameBufferSize) 			{ return static_cast<ISteamApps*>(apps)->BGetDLCDataByIndex(iDLC, pAppID, pbAvailable, pchName, cchNameBufferSize); }
void 			SteamCAPI_ISteamApps_InstallDLC(void* apps,  AppId_t nAppID) 																						{ return static_cast<ISteamApps*>(apps)->InstallDLC(nAppID); }
void 			SteamCAPI_ISteamApps_UninstallDLC(void* apps,  AppId_t nAppID) 																						{ return static_cast<ISteamApps*>(apps)->UninstallDLC(nAppID); }
void 			SteamCAPI_ISteamApps_RequestAppProofOfPurchaseKey(void* apps,  AppId_t nAppID) 																		{ return static_cast<ISteamApps*>(apps)->RequestAppProofOfPurchaseKey(nAppID); }
bool 			SteamCAPI_ISteamApps_GetCurrentBetaName(void* apps,  char* pchName, int cchNameBufferSize) 															{ return static_cast<ISteamApps*>(apps)->GetCurrentBetaName(pchName, cchNameBufferSize); }
bool 			SteamCAPI_ISteamApps_MarkContentCorrupt(void* apps,  bool bMissingFilesOnly) 																		{ return static_cast<ISteamApps*>(apps)->MarkContentCorrupt(bMissingFilesOnly); }
unsigned int 	SteamCAPI_ISteamApps_GetInstalledDepots(void* apps,  AppId_t appID, DepotId_t *pvecDepots, unsigned int cMaxDepots) 								{ return static_cast<ISteamApps*>(apps)->GetInstalledDepots(appID, pvecDepots, cMaxDepots); }
unsigned int 	SteamCAPI_ISteamApps_GetAppInstallDir(void* apps,  AppId_t appID, char* pchFolder, unsigned int cchFolderBufferSize) 								{ return static_cast<ISteamApps*>(apps)->GetAppInstallDir(appID, pchFolder, cchFolderBufferSize); }
bool 			SteamCAPI_ISteamApps_BIsAppInstalled(void* apps,  AppId_t appID) 																					{ return static_cast<ISteamApps*>(apps)->BIsAppInstalled(appID); }
CSteamID 		SteamCAPI_ISteamApps_GetAppOwner(void* apps) 																										{ return static_cast<ISteamApps*>(apps)->GetAppOwner(); }
const char* 	SteamCAPI_ISteamApps_GetLaunchQueryParam(void* apps,  const char* pchKey) 																			{ return static_cast<ISteamApps*>(apps)->GetLaunchQueryParam(pchKey); }
bool 			SteamCAPI_ISteamApps_GetDlcDownloadProgress(void* apps,  AppId_t nAppID, unsigned long long* punBytesDownloaded, unsigned long long* punBytesTotal)	{ return static_cast<ISteamApps*>(apps)->GetDlcDownloadProgress(nAppID, punBytesDownloaded, punBytesTotal); }
int 			SteamCAPI_ISteamApps_GetAppBuildId(void* apps) 																										{ return static_cast<ISteamApps*>(apps)->GetAppBuildId(); }


//==============================================================================
//=============================Steam App Ticket API=============================
//==============================================================================

uint32 SteamCAPI_ISteamAppTicket_GetAppOwnershipTicketData(void* ticket, uint32 nAppID, void* pvBuffer, uint32 cbBufferLength, uint32 *piAppId, uint32 *piSteamId, uint32 *piSignature, uint32 *pcbSignature) { return static_cast<ISteamAppTicket*>(ticket)->GetAppOwnershipTicketData(nAppID, pvBuffer, cbBufferLength, piAppId, piSteamId, piSignature, pcbSignature); }

//==============================================================================
//===============================Steam client API===============================
//==============================================================================
void* SteamCAPI_SteamClient() { return SteamClient(); }

HSteamPipe 		SteamCAPI_SteamClient_CreateSteamPipe(void* client) 																					{ return static_cast<ISteamClient*>(client)->CreateSteamPipe(); }
bool 			SteamCAPI_SteamClient_BReleaseSteamPipe(void* client, HSteamPipe hSteamPipe) 															{ return static_cast<ISteamClient*>(client)->BReleaseSteamPipe(hSteamPipe); }
HSteamUser 		SteamCAPI_SteamClient_ConnectToGlobalUser(void* client, HSteamPipe hSteamPipe) 															{ return static_cast<ISteamClient*>(client)->ConnectToGlobalUser(hSteamPipe); }
HSteamUser 		SteamCAPI_SteamClient_CreateLocalUser(void* client, HSteamPipe *phSteamPipe, EAccountType eAccountType) 								{ return static_cast<ISteamClient*>(client)->CreateLocalUser(phSteamPipe, eAccountType); }
void 			SteamCAPI_SteamClient_ReleaseUser(void* client, HSteamPipe hSteamPipe, HSteamUser hUser) 												{ 		 static_cast<ISteamClient*>(client)->ReleaseUser(hSteamPipe, hUser); }
void* 			SteamCAPI_SteamClient_GetISteamUser(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 				{ return static_cast<ISteamClient*>(client)->GetISteamUser(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamGameServer(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamGameServer(hSteamUser, hSteamPipe, pchVersion); }
void 			SteamCAPI_SteamClient_SetLocalIPBinding(void* client, unsigned int unIP, unsigned short int usPort) 									{ 		 static_cast<ISteamClient*>(client)->SetLocalIPBinding(unIP, usPort); }
void* 			SteamCAPI_SteamClient_GetISteamFriends(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 				{ return static_cast<ISteamClient*>(client)->GetISteamFriends(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamUtils(void* client, HSteamPipe hSteamPipe, const char *pchVersion) 										{ return static_cast<ISteamClient*>(client)->GetISteamUtils(hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamMatchmaking(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamMatchmaking(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamMatchmakingServers(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 	{ return static_cast<ISteamClient*>(client)->GetISteamMatchmakingServers(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamGenericInterface(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 	{ return static_cast<ISteamClient*>(client)->GetISteamGenericInterface(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamUserStats(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamUserStats(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamGameServerStats(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 		{ return static_cast<ISteamClient*>(client)->GetISteamGameServerStats(hSteamuser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamApps(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 				{ return static_cast<ISteamClient*>(client)->GetISteamApps(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamNetworking(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamNetworking(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamRemoteStorage(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 		{ return static_cast<ISteamClient*>(client)->GetISteamRemoteStorage(hSteamuser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamScreenshots(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamScreenshots(hSteamuser, hSteamPipe, pchVersion); }
void 			SteamCAPI_SteamClient_RunFrame(void* client) 																							{ 		 static_cast<ISteamClient*>(client)->RunFrame(); }
unsigned int 	SteamCAPI_SteamClient_GetIPCCallCount(void* client) 																					{ return static_cast<ISteamClient*>(client)->GetIPCCallCount(); }
void 			SteamCAPI_SteamClient_SetWarningMessageHook(void* client, SteamAPIWarningMessageHook_t pFunction) 										{ 		 static_cast<ISteamClient*>(client)->SetWarningMessageHook(pFunction); }
bool 			SteamCAPI_SteamClient_BShutdownIfAllPipesClosed(void* client) 																			{ return static_cast<ISteamClient*>(client)->BShutdownIfAllPipesClosed(); }
void* 			SteamCAPI_SteamClient_GetISteamHTTP(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 				{ return static_cast<ISteamClient*>(client)->GetISteamHTTP(hSteamuser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamUnifiedMessages(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 		{ return static_cast<ISteamClient*>(client)->GetISteamUnifiedMessages(hSteamuser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamController(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamController(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamUGC(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 					{ return static_cast<ISteamClient*>(client)->GetISteamUGC(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamAppList(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion) 				{ return static_cast<ISteamClient*>(client)->GetISteamAppList(hSteamUser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamMusic(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 				{ return static_cast<ISteamClient*>(client)->GetISteamMusic(hSteamuser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamMusicRemote(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamMusicRemote(hSteamuser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamHTMLSurface(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamHTMLSurface(hSteamuser, hSteamPipe, pchVersion); }
void 			SteamCAPI_SteamClient_Set_SteamAPI_CPostAPIResultInProcess(void* client, SteamAPI_PostAPIResultInProcess_t func) 						{ 		 static_cast<ISteamClient*>(client)->Set_SteamAPI_CPostAPIResultInProcess(func); }
void 			SteamCAPI_SteamClient_Remove_SteamAPI_CPostAPIResultInProcess(void* client, SteamAPI_PostAPIResultInProcess_t func) 					{ 		 static_cast<ISteamClient*>(client)->Remove_SteamAPI_CPostAPIResultInProcess(func); }
void 			SteamCAPI_SteamClient_Set_SteamAPI_CCheckCallbackRegisteredInProcess(void* client, SteamAPI_CheckCallbackRegistered_t func) 			{ 		 static_cast<ISteamClient*>(client)->Set_SteamAPI_CCheckCallbackRegisteredInProcess(func); }
void* 			SteamCAPI_SteamClient_GetISteamInventory(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 			{ return static_cast<ISteamClient*>(client)->GetISteamInventory(hSteamuser, hSteamPipe, pchVersion); }
void* 			SteamCAPI_SteamClient_GetISteamVideo(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion) 				{ return static_cast<ISteamClient*>(client)->GetISteamVideo(hSteamuser, hSteamPipe, pchVersion); }



//========================================================
//==================Steam Controller API==================
//========================================================

void* CSteamController() { return SteamController(); }

bool 								SteamCAPI_SteamController_Init(void* controller) 																																																{ return static_cast<ISteamController*>(controller)->Init(); }
bool 								SteamCAPI_SteamController_Shutdown(void* controller) 																																															{ return static_cast<ISteamController*>(controller)->Shutdown(); }
void 								SteamCAPI_SteamController_RunFrame(void* controller) 																																															{ return static_cast<ISteamController*>(controller)->RunFrame(); }
int 								SteamCAPI_SteamController_GetConnectedControllers(void* controller, ControllerHandle_t *handlesOut) 																																			{ return static_cast<ISteamController*>(controller)->GetConnectedControllers(handlesOut); }
bool 								SteamCAPI_SteamController_ShowBindingPanel(void* controller, ControllerHandle_t controllerHandle) 																																				{ return static_cast<ISteamController*>(controller)->ShowBindingPanel(controllerHandle); }
ControllerActionSetHandle_t 		SteamCAPI_SteamController_GetActionSetHandle(void* controller, const char *pszActionSetName) 																																					{ return static_cast<ISteamController*>(controller)->GetActionSetHandle(pszActionSetName); }
void 								SteamCAPI_SteamController_ActivateActionSet(void* controller, ControllerHandle_t controllerHandle, ControllerActionSetHandle_t actionSetHandle) 																								{ return static_cast<ISteamController*>(controller)->ActivateActionSet(controllerHandle, actionSetHandle); }
ControllerActionSetHandle_t 		SteamCAPI_SteamController_GetCurrentActionSet(void* controller, ControllerHandle_t controllerHandle) 																																			{ return static_cast<ISteamController*>(controller)->GetCurrentActionSet(controllerHandle); }
ControllerDigitalActionHandle_t 	SteamCAPI_SteamController_GetDigitalActionHandle(void* controller, const char *pszActionName) 																																					{ return static_cast<ISteamController*>(controller)->GetDigitalActionHandle(pszActionName); }
ControllerDigitalActionData_t 		SteamCAPI_SteamController_GetDigitalActionData(void* controller, ControllerHandle_t controllerHandle, ControllerDigitalActionHandle_t digitalActionHandle) 																						{ return static_cast<ISteamController*>(controller)->GetDigitalActionData(controllerHandle, digitalActionHandle); }
int 								SteamCAPI_SteamController_GetDigitalActionOrigins(void* controller, ControllerHandle_t controllerHandle, ControllerActionSetHandle_t actionSetHandle, ControllerDigitalActionHandle_t digitalActionHandle, EControllerActionOrigin *originsOut) { return static_cast<ISteamController*>(controller)->GetDigitalActionOrigins(controllerHandle, actionSetHandle, digitalActionHandle, originsOut); }
ControllerAnalogActionHandle_t 		SteamCAPI_SteamController_GetAnalogActionHandle(void* controller, const char *pszActionName) 																																					{ return static_cast<ISteamController*>(controller)->GetAnalogActionHandle(pszActionName); }
ControllerAnalogActionData_t 		SteamCAPI_SteamController_GetAnalogActionData(void* controller, ControllerHandle_t controllerHandle, ControllerAnalogActionHandle_t analogActionHandle) 																						{ return static_cast<ISteamController*>(controller)->GetAnalogActionData(controllerHandle, analogActionHandle); }
int 								SteamCAPI_SteamController_GetAnalogActionOrigins(void* controller, ControllerHandle_t controllerHandle, ControllerActionSetHandle_t actionSetHandle, ControllerAnalogActionHandle_t analogActionHandle, EControllerActionOrigin *originsOut) 	{ return static_cast<ISteamController*>(controller)->GetAnalogActionOrigins(controllerHandle, actionSetHandle, analogActionHandle, originsOut); }
void 								SteamCAPI_SteamController_StopAnalogActionMomentum(void* controller, ControllerHandle_t controllerHandle, ControllerAnalogActionHandle_t eAction) 																								{ return static_cast<ISteamController*>(controller)->StopAnalogActionMomentum(controllerHandle, eAction); }
void 								SteamCAPI_SteamController_TriggerHapticPulse(void* controller, ControllerHandle_t controllerHandle, ESteamControllerPad eTargetPad, unsigned short usDurationMicroSec) 																			{ return static_cast<ISteamController*>(controller)->TriggerHapticPulse(controllerHandle, eTargetPad, usDurationMicroSec); }


//========================================================
//===================Steam friends API====================
//========================================================


void* CSteamFriends() { return SteamFriends(); }

const char* 		SteamFriends_GetPersonaName(void* steamFriends)																																												{ return static_cast<ISteamFriends*>(steamFriends)->GetPersonaName(); }
SteamAPICall_t 		SteamFriends_SetPersonaName(void* steamFriends, char* pchPersonaName)																																						{ return static_cast<ISteamFriends*>(steamFriends)->SetPersonaName(pchPersonaName); }
EPersonaState 		SteamFriends_GetPersonaState(void* steamFriends)																																											{ return static_cast<ISteamFriends*>(steamFriends)->GetPersonaState(); }
int 				SteamFriends_GetFriendCount(void* steamFriends, int iFriendFlags)																																							{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendCount(iFriendFlags); }
CSteamID 			SteamFriends_GetFriendByIndex(void* steamFriends, int iFriend, int iFriendFlags)																																			{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendByIndex(iFriend, iFriendFlags); }
EFriendRelationship SteamFriends_GetFriendRelationship(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendRelationship(steamIDFriend); }
EPersonaState 		SteamFriends_GetFriendPersonaState(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendPersonaState(steamIDFriend); }
const char* 		SteamFriends_GetFriendPersonaName(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendPersonaName(steamIDFriend); }
bool 				SteamFriends_GetFriendGamePlayed(void* steamFriends, CSteamID steamIDFriend, OUT_STRUCT() FriendGameInfo_t *pFriendGameInfo)																								{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendGamePlayed(steamIDFriend, pFriendGameInfo); }
const char* 		SteamFriends_GetFriendPersonaNameHistory(void* steamFriends, CSteamID steamIDFriend, int iPersonaName)																														{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendPersonaNameHistory(steamIDFriend, iPersonaName); }
int 				SteamFriends_GetFriendSteamLevel(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendSteamLevel(steamIDFriend); }
const char* 		SteamFriends_GetPlayerNickname(void* steamFriends, CSteamID steamIDPlayer)																																					{ return static_cast<ISteamFriends*>(steamFriends)->GetPlayerNickname(steamIDPlayer); }
int 				SteamFriends_GetFriendsGroupCount(void* steamFriends)																																										{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendsGroupCount(); }
FriendsGroupID_t 	SteamFriends_GetFriendsGroupIDByIndex(void* steamFriends, int iFG)																																							{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendsGroupIDByIndex(iFG); }
const char* 		SteamFriends_GetFriendsGroupName(void* steamFriends, FriendsGroupID_t friendsGroupID)																																		{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendsGroupName(friendsGroupID); }
int 				SteamFriends_GetFriendsGroupMembersCount(void* steamFriends, FriendsGroupID_t friendsGroupID)																																{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendsGroupMembersCount(friendsGroupID); }
void 				SteamFriends_GetFriendsGroupMembersList(void* steamFriends, FriendsGroupID_t friendsGroupID, OUT_ARRAY_CALL(nMembersCount, GetFriendsGroupMembersCount, friendsGroupID) CSteamID *pOutSteamIDMembers, int nMembersCount)	{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendsGroupMembersList(friendsGroupID, pOutSteamIDMembers, nMembersCount); }
bool 				SteamFriends_HasFriend(void* steamFriends, CSteamID steamIDFriend, int iFriendFlags)																																		{ return static_cast<ISteamFriends*>(steamFriends)->HasFriend(steamIDFriend, iFriendFlags); }
int 				SteamFriends_GetClanCount(void* steamFriends)																																												{ return static_cast<ISteamFriends*>(steamFriends)->GetClanCount(); }
CSteamID 			SteamFriends_GetClanByIndex(void* steamFriends, int iClan)																																									{ return static_cast<ISteamFriends*>(steamFriends)->GetClanByIndex(iClan); }
const char *		SteamFriends_GetClanName(void* steamFriends, CSteamID steamIDClan)																																							{ return static_cast<ISteamFriends*>(steamFriends)->GetClanName(steamIDClan); }
const char *		SteamFriends_GetClanTag(void* steamFriends, CSteamID steamIDClan)																																							{ return static_cast<ISteamFriends*>(steamFriends)->GetClanTag(steamIDClan); }
bool 				SteamFriends_GetClanActivityCounts(void* steamFriends, CSteamID steamIDClan, int* pnOnline, int* pnInGame, int* pnChatting)																									{ return static_cast<ISteamFriends*>(steamFriends)->GetClanActivityCounts(steamIDClan, pnOnline, pnInGame, pnChatting); }
SteamAPICall_t 		SteamFriends_DownloadClanActivityCounts(void* steamFriends, ARRAY_COUNT(cClansToRequest) CSteamID *psteamIDClans, int cClansToRequest)																						{ return static_cast<ISteamFriends*>(steamFriends)->DownloadClanActivityCounts(psteamIDClans, cClansToRequest); }
int 				SteamFriends_GetFriendCountFromSource(void* steamFriends, CSteamID steamIDSource)																																			{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendCountFromSource(steamIDSource); }
CSteamID 			SteamFriends_GetFriendFromSourceByIndex(void* steamFriends, CSteamID steamIDSource, int iFriend)																															{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendFromSourceByIndex(steamIDSource, iFriend); }
bool 				SteamFriends_IsUserInSource(void* steamFriends, CSteamID steamIDUser, CSteamID steamIDSource)																																{ return static_cast<ISteamFriends*>(steamFriends)->IsUserInSource(steamIDUser, steamIDSource); }
void 				SteamFriends_SetInGameVoiceSpeaking(void* steamFriends, CSteamID steamIDUser, bool bSpeaking)																																{ return static_cast<ISteamFriends*>(steamFriends)->SetInGameVoiceSpeaking(steamIDUser, bSpeaking); }
void 				SteamFriends_ActivateGameOverlay(void* steamFriends, const char* pchDialog)																																					{ return static_cast<ISteamFriends*>(steamFriends)->ActivateGameOverlay(pchDialog); }
void 				SteamFriends_ActivateGameOverlayToUser(void* steamFriends, const char* pchDialog, CSteamID steamID)																															{ return static_cast<ISteamFriends*>(steamFriends)->ActivateGameOverlayToUser(pchDialog, steamID); }
void 				SteamFriends_ActivateGameOverlayToWebPage(void* steamFriends, const char* pchURL)																																			{ return static_cast<ISteamFriends*>(steamFriends)->ActivateGameOverlayToWebPage(pchURL); }
void 				SteamFriends_ActivateGameOverlayToStore(void* steamFriends, AppId_t nAppID, EOverlayToStoreFlag eFlag)																														{ return static_cast<ISteamFriends*>(steamFriends)->ActivateGameOverlayToStore(nAppID, eFlag); }
void 				SteamFriends_SetPlayedWith(void* steamFriends, CSteamID steamIDUserPlayedWith)																																				{ return static_cast<ISteamFriends*>(steamFriends)->SetPlayedWith(steamIDUserPlayedWith); }
void 				SteamFriends_ActivateGameOverlayInviteDialog(void* steamFriends, CSteamID steamIDLobby)																																		{ return static_cast<ISteamFriends*>(steamFriends)->ActivateGameOverlayInviteDialog(steamIDLobby); }
int 				SteamFriends_GetSmallFriendAvatar(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetSmallFriendAvatar(steamIDFriend); }
int 				SteamFriends_GetMediumFriendAvatar(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetMediumFriendAvatar(steamIDFriend); }
int 				SteamFriends_GetLargeFriendAvatar(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetLargeFriendAvatar(steamIDFriend); }
bool 				SteamFriends_RequestUserInformation(void* steamFriends, CSteamID steamIDUser, bool bRequireNameOnly)																														{ return static_cast<ISteamFriends*>(steamFriends)->RequestUserInformation(steamIDUser, bRequireNameOnly); }
SteamAPICall_t 		SteamFriends_RequestClanOfficerList(void* steamFriends, CSteamID steamIDClan)																																				{ return static_cast<ISteamFriends*>(steamFriends)->RequestClanOfficerList(steamIDClan); }
CSteamID 			SteamFriends_GetClanOwner(void* steamFriends, CSteamID steamIDClan)																																							{ return static_cast<ISteamFriends*>(steamFriends)->GetClanOwner(steamIDClan); }
int 				SteamFriends_GetClanOfficerCount(void* steamFriends, CSteamID steamIDClan)																																					{ return static_cast<ISteamFriends*>(steamFriends)->GetClanOfficerCount(steamIDClan); }
CSteamID 			SteamFriends_GetClanOfficerByIndex(void* steamFriends, CSteamID steamIDClan, int iOfficer)																																	{ return static_cast<ISteamFriends*>(steamFriends)->GetClanOfficerByIndex(steamIDClan, iOfficer); }
unsigned int 		SteamFriends_GetUserRestrictions(void* steamFriends)																																										{ return static_cast<ISteamFriends*>(steamFriends)->GetUserRestrictions(); }
bool 				SteamFriends_SetRichPresence(void* steamFriends, const char* pchKey, const char* pchValue)																																	{ return static_cast<ISteamFriends*>(steamFriends)->SetRichPresence(pchKey, pchValue); }
void 				SteamFriends_ClearRichPresence(void* steamFriends)																																											{ 		 static_cast<ISteamFriends*>(steamFriends)->ClearRichPresence(); }
const char* 		SteamFriends_GetFriendRichPresence(void* steamFriends, CSteamID steamIDFriend, const char* pchKey)																															{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendRichPresence(steamIDFriend, pchKey); }
int 				SteamFriends_GetFriendRichPresenceKeyCount(void* steamFriends, CSteamID steamIDFriend)																																		{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendRichPresenceKeyCount(steamIDFriend); }
const char *		SteamFriends_GetFriendRichPresenceKeyByIndex(void* steamFriends, CSteamID steamIDFriend, int iKey)																															{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendRichPresenceKeyByIndex(steamIDFriend, iKey); }
void 				SteamFriends_RequestFriendRichPresence(void* steamFriends, CSteamID steamIDFriend)																																			{ return static_cast<ISteamFriends*>(steamFriends)->RequestFriendRichPresence(steamIDFriend); }
bool 				SteamFriends_InviteUserToGame(void* steamFriends, CSteamID steamIDFriend, const char* pchConnectString)																														{ return static_cast<ISteamFriends*>(steamFriends)->InviteUserToGame(steamIDFriend, pchConnectString); }
int 				SteamFriends_GetCoplayFriendCount(void* steamFriends)																																										{ return static_cast<ISteamFriends*>(steamFriends)->GetCoplayFriendCount(); }
CSteamID 			SteamFriends_GetCoplayFriend(void* steamFriends, int iCoplayFriend)																																							{ return static_cast<ISteamFriends*>(steamFriends)->GetCoplayFriend(iCoplayFriend); }
int 				SteamFriends_GetFriendCoplayTime(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendCoplayTime(steamIDFriend); }
AppId_t 			SteamFriends_GetFriendCoplayGame(void* steamFriends, CSteamID steamIDFriend)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendCoplayGame(steamIDFriend); }
SteamAPICall_t 		SteamFriends_JoinClanChatRoom(void* steamFriends, CSteamID steamIDClan)																																						{ return static_cast<ISteamFriends*>(steamFriends)->JoinClanChatRoom(steamIDClan); }
bool 				SteamFriends_LeaveClanChatRoom(void* steamFriends, CSteamID steamIDClan)																																					{ return static_cast<ISteamFriends*>(steamFriends)->LeaveClanChatRoom(steamIDClan); }
int 				SteamFriends_GetClanChatMemberCount(void* steamFriends, CSteamID steamIDClan)																																				{ return static_cast<ISteamFriends*>(steamFriends)->GetClanChatMemberCount(steamIDClan); }
CSteamID 			SteamFriends_GetChatMemberByIndex(void* steamFriends, CSteamID steamIDClan, int iUser)																																		{ return static_cast<ISteamFriends*>(steamFriends)->GetChatMemberByIndex(steamIDClan, iUser); }
bool 				SteamFriends_SendClanChatMessage(void* steamFriends, CSteamID steamIDClanChat, const char* pchText)																															{ return static_cast<ISteamFriends*>(steamFriends)->SendClanChatMessage(steamIDClanChat, pchText); }
int 				SteamFriends_GetClanChatMessage(void* steamFriends, CSteamID steamIDClanChat, int iMessage, void *prgchText, int cchTextMax, EChatEntryType *peChatEntryType, OUT_STRUCT() CSteamID *psteamidChatter)						{ return static_cast<ISteamFriends*>(steamFriends)->GetClanChatMessage(steamIDClanChat, iMessage, prgchText, cchTextMax, peChatEntryType, psteamidChatter); }
bool 				SteamFriends_IsClanChatAdmin(void* steamFriends, CSteamID steamIDClanChat, CSteamID steamIDUser)																															{ return static_cast<ISteamFriends*>(steamFriends)->IsClanChatAdmin(steamIDClanChat, steamIDUser); }
bool 				SteamFriends_IsClanChatWindowOpenInSteam(void* steamFriends, CSteamID steamIDClanChat)																																		{ return static_cast<ISteamFriends*>(steamFriends)->IsClanChatWindowOpenInSteam(steamIDClanChat); }
bool 				SteamFriends_OpenClanChatWindowInSteam(void* steamFriends, CSteamID steamIDClanChat)																																		{ return static_cast<ISteamFriends*>(steamFriends)->OpenClanChatWindowInSteam(steamIDClanChat); }
bool 				SteamFriends_CloseClanChatWindowInSteam(void* steamFriends, CSteamID steamIDClanChat)																																		{ return static_cast<ISteamFriends*>(steamFriends)->CloseClanChatWindowInSteam(steamIDClanChat); }
bool 				SteamFriends_SetListenForFriendsMessages(void* steamFriends, bool bInterceptEnabled)																																		{ return static_cast<ISteamFriends*>(steamFriends)->SetListenForFriendsMessages(bInterceptEnabled); }
bool 				SteamFriends_ReplyToFriendMessage(void* steamFriends, CSteamID steamIDFriend, const char* pchMsgToSend)																														{ return static_cast<ISteamFriends*>(steamFriends)->ReplyToFriendMessage(steamIDFriend, pchMsgToSend); }
int 				SteamFriends_GetFriendMessage(void* steamFriends, CSteamID steamIDFriend, int iMessageID, void *pvData, int cubData, EChatEntryType *peChatEntryType)																		{ return static_cast<ISteamFriends*>(steamFriends)->GetFriendMessage(steamIDFriend, iMessageID, pvData, cubData, peChatEntryType); }
SteamAPICall_t 		SteamFriends_GetFollowerCount(void* steamFriends, CSteamID steamID)																																							{ return static_cast<ISteamFriends*>(steamFriends)->GetFollowerCount(steamID); }
SteamAPICall_t 		SteamFriends_IsFollowing(void* steamFriends, CSteamID steamID)																																								{ return static_cast<ISteamFriends*>(steamFriends)->IsFollowing(steamID); }
SteamAPICall_t 		SteamFriends_EnumerateFollowingList(void* steamFriends, uint32 unStartIndex)																																				{ return static_cast<ISteamFriends*>(steamFriends)->EnumerateFollowingList(unStartIndex); }



//==============================================================================
//==========================Steam game coordinator API==========================
//==============================================================================

EGCResults	SteamCAPI_ISteamGameCoordinator_SendMessage(void* gc, unsigned int unMsgType, const void *pubData, unsigned int cubData)							{ return static_cast<ISteamGameCoordinator*>(gc)->SendMessage(unMsgType, pubData, cubData); }
bool		SteamCAPI_ISteamGameCoordinator_IsMessageAvailable(void* gc, unsigned int* pcubMsgSize)																{ return static_cast<ISteamGameCoordinator*>(gc)->IsMessageAvailable(pcubMsgSize); }
EGCResults	SteamCAPI_ISteamGameCoordinator_RetrieveMessage(void* gc, unsigned int* punMsgType, void *pubDest, unsigned int cubDest, unsigned int *pcubMsgSize)	{ return static_cast<ISteamGameCoordinator*>(gc)->RetrieveMessage(punMsgType, pubDest, cubDest, pcubMsgSize); }


//======================================================================
//========================Steam game server API=========================
//======================================================================

bool 						SteamCAPI_ISteamGameServer_InitGameServer(void* server, unsigned int unIP, unsigned short int usGamePort, unsigned short int usQueryPort, unsigned int unFlags, AppId_t nGameAppId, const char *pchVersionString)	{ return static_cast<ISteamGameServer*>(server)->InitGameServer(unIP, usGamePort, usQueryPort, unFlags, nGameAppId, pchVersionString); }
void 						SteamCAPI_ISteamGameServer_SetProduct(void* server, const char *pszProduct)																																			{ 		 static_cast<ISteamGameServer*>(server)->SetProduct(pszProduct); }
void 						SteamCAPI_ISteamGameServer_SetGameDescription(void* server, const char *pszGameDescription)																															{ 		 static_cast<ISteamGameServer*>(server)->SetGameDescription(pszGameDescription); }
void 						SteamCAPI_ISteamGameServer_SetModDir(void* server, const char *pszModDir)																																			{ 		 static_cast<ISteamGameServer*>(server)->SetModDir(pszModDir); }
void 						SteamCAPI_ISteamGameServer_SetDedicatedServer(void* server, bool bDedicated)																																		{ 		 static_cast<ISteamGameServer*>(server)->SetDedicatedServer(bDedicated); }
void 						SteamCAPI_ISteamGameServer_LogOn(void* server, const char *pszToken)																																				{ 		 static_cast<ISteamGameServer*>(server)->LogOn(pszToken); }
void 						SteamCAPI_ISteamGameServer_LogOnAnonymous(void* server)																																								{ 		 static_cast<ISteamGameServer*>(server)->LogOnAnonymous(); }
void 						SteamCAPI_ISteamGameServer_LogOff(void* server)																																										{ 		 static_cast<ISteamGameServer*>(server)->LogOff(); }
bool 						SteamCAPI_ISteamGameServer_BLoggedOn(void* server)																																									{ return static_cast<ISteamGameServer*>(server)->BLoggedOn(); }
bool 						SteamCAPI_ISteamGameServer_BSecure(void* server) 																																									{ return static_cast<ISteamGameServer*>(server)->BSecure(); }
CSteamID 					SteamCAPI_ISteamGameServer_GetSteamID(void* server)																																									{ return static_cast<ISteamGameServer*>(server)->GetSteamID(); }
bool 						SteamCAPI_ISteamGameServer_WasRestartRequested(void* server)																																						{ return static_cast<ISteamGameServer*>(server)->WasRestartRequested(); }
void 						SteamCAPI_ISteamGameServer_SetMaxPlayerCount(void* server, int cPlayersMax)																																			{ 		 static_cast<ISteamGameServer*>(server)->SetMaxPlayerCount(cPlayersMax); }
void 						SteamCAPI_ISteamGameServer_SetBotPlayerCount(void* server, int cBotplayers)																																			{ 		 static_cast<ISteamGameServer*>(server)->SetBotPlayerCount(cBotplayers); }
void 						SteamCAPI_ISteamGameServer_SetServerName(void* server, const char *pszServerName)																																	{ 		 static_cast<ISteamGameServer*>(server)->SetServerName(pszServerName); }
void 						SteamCAPI_ISteamGameServer_SetMapName(void* server, const char *pszMapName)																																			{ 		 static_cast<ISteamGameServer*>(server)->SetMapName(pszMapName); }
void 						SteamCAPI_ISteamGameServer_SetPasswordProtected(void* server, bool bPasswordProtected)																																{ 		 static_cast<ISteamGameServer*>(server)->SetPasswordProtected(bPasswordProtected); }
void 						SteamCAPI_ISteamGameServer_SetSpectatorPort(void* server, unsigned short int unSpectatorPort)																														{ 		 static_cast<ISteamGameServer*>(server)->SetSpectatorPort(unSpectatorPort); }
void 						SteamCAPI_ISteamGameServer_SetSpectatorServerName(void* server, const char *pszSpectatorServerName)																													{ 		 static_cast<ISteamGameServer*>(server)->SetSpectatorServerName(pszSpectatorServerName); }
void 						SteamCAPI_ISteamGameServer_ClearAllKeyValues(void* server)																																							{ 		 static_cast<ISteamGameServer*>(server)->ClearAllKeyValues(); }
void 						SteamCAPI_ISteamGameServer_SetKeyValue(void* server, const char *pKey, const char *pValue)																															{ 		 static_cast<ISteamGameServer*>(server)->SetKeyValue(pKey, pValue); }
void 						SteamCAPI_ISteamGameServer_SetGameTags(void* server, const char *pchGameTags)																																		{ 		 static_cast<ISteamGameServer*>(server)->SetGameTags(pchGameTags); }
void 						SteamCAPI_ISteamGameServer_SetGameData(void* server, const char *pchGameData)																																		{ 		 static_cast<ISteamGameServer*>(server)->SetGameData(pchGameData); }
void 						SteamCAPI_ISteamGameServer_SetRegion(void* server, const char *pszRegion)																																			{ 		 static_cast<ISteamGameServer*>(server)->SetRegion(pszRegion); }
bool 						SteamCAPI_ISteamGameServer_SendUserConnectAndAuthenticate(void* server, unsigned int unIPClient, const void *pvAuthBlob, unsigned int cubAuthBlobSize, CSteamID *pSteamIDUser)										{ return static_cast<ISteamGameServer*>(server)->SendUserConnectAndAuthenticate(unIPClient, pvAuthBlob, cubAuthBlobSize, pSteamIDUser); }
CSteamID 					SteamCAPI_ISteamGameServer_CreateUnauthenticatedUserConnection(void* server)																																		{ return static_cast<ISteamGameServer*>(server)->CreateUnauthenticatedUserConnection(); }
void 						SteamCAPI_ISteamGameServer_SendUserDisconnect(void* server, CSteamID steamIDUser)																																	{ 		 static_cast<ISteamGameServer*>(server)->SendUserDisconnect(steamIDUser); }
bool 						SteamCAPI_ISteamGameServer_BUpdateUserData(void* server, CSteamID steamIDUser, const char *pchPlayerName, unsigned int uScore)																						{ return static_cast<ISteamGameServer*>(server)->BUpdateUserData(steamIDUser, pchPlayerName, uScore); }
HAuthTicket 				SteamCAPI_ISteamGameServer_GetAuthSessionTicket(void* server, void *pTicket, int cbMaxTicket, unsigned int *pcbTicket)																								{ return static_cast<ISteamGameServer*>(server)->GetAuthSessionTicket(pTicket, cbMaxTicket, pcbTicket); }
EBeginAuthSessionResult 	SteamCAPI_ISteamGameServer_BeginAuthSession(void* server, const void *pAuthTicket, int cbAuthTicket, CSteamID steamID)																								{ return static_cast<ISteamGameServer*>(server)->BeginAuthSession(pAuthTicket, cbAuthTicket, steamID); }
void 						SteamCAPI_ISteamGameServer_EndAuthSession(void* server, CSteamID steamID)																																			{ 		 static_cast<ISteamGameServer*>(server)->EndAuthSession(steamID); }
void 						SteamCAPI_ISteamGameServer_CancelAuthTicket(void* server, HAuthTicket hAuthTicket)																																	{ 		 static_cast<ISteamGameServer*>(server)->CancelAuthTicket(hAuthTicket); }
EUserHasLicenseForAppResult SteamCAPI_ISteamGameServer_UserHasLicenseForApp(void* server, CSteamID steamID, AppId_t appID)																														{ return static_cast<ISteamGameServer*>(server)->UserHasLicenseForApp(steamID, appID); }
bool 						SteamCAPI_ISteamGameServer_RequestUserGroupStatus(void* server, CSteamID steamIDUser, CSteamID steamIDGroup)																										{ return static_cast<ISteamGameServer*>(server)->RequestUserGroupStatus(steamIDUser, steamIDGroup); }
void 						SteamCAPI_ISteamGameServer_GetGameplayStats(void* server)																																							{ 		 static_cast<ISteamGameServer*>(server)->GetGameplayStats(); }
SteamAPICall_t 				SteamCAPI_ISteamGameServer_GetServerReputation(void* server)																																						{ return static_cast<ISteamGameServer*>(server)->GetServerReputation(); }
unsigned int 				SteamCAPI_ISteamGameServer_GetPublicIP(void* server)																																								{ return static_cast<ISteamGameServer*>(server)->GetPublicIP(); }
bool 						SteamCAPI_ISteamGameServer_HandleIncomingPacket(void* server, const void *pData, int cbData, unsigned int srcIP, unsigned short int srcPort)																		{ return static_cast<ISteamGameServer*>(server)->HandleIncomingPacket(pData, cbData, srcIP, srcPort); }
int 						SteamCAPI_ISteamGameServer_GetNextOutgoingPacket(void* server, void *pOut, int cbMaxOut, unsigned int *pNetAdr, unsigned short int *pPort)																			{ return static_cast<ISteamGameServer*>(server)->GetNextOutgoingPacket(pOut, cbMaxOut, pNetAdr, pPort); }
void 						SteamCAPI_ISteamGameServer_EnableHeartbeats(void* server, bool bActive)																																				{ 		 static_cast<ISteamGameServer*>(server)->EnableHeartbeats(bActive); }
void 						SteamCAPI_ISteamGameServer_SetHeartbeatInterval(void* server, int iHeartbeatInterval)																																{ 		 static_cast<ISteamGameServer*>(server)->SetHeartbeatInterval(iHeartbeatInterval); }
void 						SteamCAPI_ISteamGameServer_ForceHeartbeat(void* server)																																								{ 		 static_cast<ISteamGameServer*>(server)->ForceHeartbeat(); }
SteamAPICall_t 				SteamCAPI_ISteamGameServer_AssociateWithClan(void* server, CSteamID steamIDClan)																																	{ return static_cast<ISteamGameServer*>(server)->AssociateWithClan(steamIDClan); }
SteamAPICall_t 				SteamCAPI_ISteamGameServer_ComputeNewPlayerCompatibility(void* server, CSteamID steamIDNewPlayer)																													{ return static_cast<ISteamGameServer*>(server)->ComputeNewPlayerCompatibility(steamIDNewPlayer); }

//==============================================================================
//=========================Steam game server stats API==========================
//==============================================================================

void* CSteamGameServerStats() { return SteamGameServerStats(); }

SteamAPICall_t 	ISteamGameServerStats_RequestUserStats(void* stats, CSteamID steamIDUser) 																				{ return static_cast<ISteamGameServerStats*>(stats)->RequestUserStats(steamIDUser); }
bool 			ISteamGameServerStats_GetUserStati(void* stats, CSteamID steamIDUser, const char *pchName, int32 *pData) 												{ return static_cast<ISteamGameServerStats*>(stats)->GetUserStat(steamIDUser, pchName, pData); }
bool 			ISteamGameServerStats_GetUserStatf(void* stats, CSteamID steamIDUser, const char *pchName, float *pData) 												{ return static_cast<ISteamGameServerStats*>(stats)->GetUserStat(steamIDUser, pchName, pData); }
bool 			ISteamGameServerStats_GetUserAchievement(void* stats, CSteamID steamIDUser, const char *pchName, bool *pbAchieved) 										{ return static_cast<ISteamGameServerStats*>(stats)->GetUserAchievement(steamIDUser, pchName, pbAchieved); }
bool 			ISteamGameServerStats_SetUserStati(void* stats, CSteamID steamIDUser, const char *pchName, int32 nData) 												{ return static_cast<ISteamGameServerStats*>(stats)->SetUserStat(steamIDUser, pchName, nData); }
bool 			ISteamGameServerStats_SetUserStatf(void* stats, CSteamID steamIDUser, const char *pchName, float fData) 												{ return static_cast<ISteamGameServerStats*>(stats)->SetUserStat(steamIDUser, pchName, fData); }
bool 			ISteamGameServerStats_UpdateUserAvgRateStat(void* stats, CSteamID steamIDUser, const char *pchName, float flCountThisSession, double dSessionLength)	{ return static_cast<ISteamGameServerStats*>(stats)->UpdateUserAvgRateStat(steamIDUser, pchName, flCountThisSession, dSessionLength); }
bool 			ISteamGameServerStats_SetUserAchievement(void* stats, CSteamID steamIDUser, const char *pchName) 														{ return static_cast<ISteamGameServerStats*>(stats)->SetUserAchievement(steamIDUser, pchName); }
bool 			ISteamGameServerStats_ClearUserAchievement(void* stats, CSteamID steamIDUser, const char *pchName) 														{ return static_cast<ISteamGameServerStats*>(stats)->ClearUserAchievement(steamIDUser, pchName); }
SteamAPICall_t 	ISteamGameServerStats_StoreUserStats(void* stats, CSteamID steamIDUser) 																				{ return static_cast<ISteamGameServerStats*>(stats)->StoreUserStats(steamIDUser); }

//==============================================================================
//============================Steam html surface API============================
//==============================================================================

bool			SteamCAPI_ISteamHTMLSurface_Init(void* surface) {
	return static_cast<ISteamHTMLSurface*>(surface)->Init();
}

bool			SteamCAPI_ISteamHTMLSurface_Shutdown(void* surface)																																				{ return static_cast<ISteamHTMLSurface*>(surface)->Shutdown(); }
SteamAPICall_t	SteamCAPI_ISteamHTMLSurface_CreateBrowser(void* surface, const char *pchUserAgent, const char *pchUserCSS)																						{ return static_cast<ISteamHTMLSurface*>(surface)->CreateBrowser(pchUserAgent, pchUserCSS); }
void			SteamCAPI_ISteamHTMLSurface_RemoveBrowser(void* surface, HHTMLBrowser unBrowserHandle)																											{ 		 static_cast<ISteamHTMLSurface*>(surface)->RemoveBrowser(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_LoadURL(void* surface, HHTMLBrowser unBrowserHandle, const char *pchURL, const char *pchPostData)																	{ 		 static_cast<ISteamHTMLSurface*>(surface)->LoadURL(unBrowserHandle, pchURL, pchPostData); }
void			SteamCAPI_ISteamHTMLSurface_SetSize(void* surface, HHTMLBrowser unBrowserHandle, unsigned int unWidth, unsigned int unHeight)																	{ 		 static_cast<ISteamHTMLSurface*>(surface)->SetSize(unBrowserHandle, unWidth, unHeight); }
void			SteamCAPI_ISteamHTMLSurface_StopLoad(void* surface, HHTMLBrowser unBrowserHandle)																												{ 		 static_cast<ISteamHTMLSurface*>(surface)->StopLoad(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_Reload(void* surface, HHTMLBrowser unBrowserHandle)																													{ 		 static_cast<ISteamHTMLSurface*>(surface)->Reload(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_GoBack(void* surface, HHTMLBrowser unBrowserHandle)																													{ 		 static_cast<ISteamHTMLSurface*>(surface)->GoBack(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_GoForward(void* surface, HHTMLBrowser unBrowserHandle)																												{ 		 static_cast<ISteamHTMLSurface*>(surface)->GoForward(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_AddHeader(void* surface, HHTMLBrowser unBrowserHandle, const char *pchKey, const char *pchValue)																	{ 		 static_cast<ISteamHTMLSurface*>(surface)->AddHeader(unBrowserHandle, pchKey, pchValue); }
void			SteamCAPI_ISteamHTMLSurface_ExecuteJavascript(void* surface, HHTMLBrowser unBrowserHandle, const char *pchScript)																				{ 		 static_cast<ISteamHTMLSurface*>(surface)->ExecuteJavascript(unBrowserHandle, pchScript); }
void			SteamCAPI_ISteamHTMLSurface_MouseUp(void* surface, HHTMLBrowser unBrowserHandle, EHTMLMouseButton eMouseButton)																					{ 		 static_cast<ISteamHTMLSurface*>(surface)->MouseUp(unBrowserHandle, static_cast<ISteamHTMLSurface::EHTMLMouseButton>(eMouseButton)); }
void			SteamCAPI_ISteamHTMLSurface_MouseDown(void* surface, HHTMLBrowser unBrowserHandle, EHTMLMouseButton eMouseButton)																				{ 		 static_cast<ISteamHTMLSurface*>(surface)->MouseDown(unBrowserHandle, static_cast<ISteamHTMLSurface::EHTMLMouseButton>(eMouseButton)); }
void			SteamCAPI_ISteamHTMLSurface_MouseDoubleClick(void* surface, HHTMLBrowser unBrowserHandle, EHTMLMouseButton eMouseButton)																		{ 		 static_cast<ISteamHTMLSurface*>(surface)->MouseDoubleClick(unBrowserHandle, static_cast<ISteamHTMLSurface::EHTMLMouseButton>(eMouseButton)); }
void			SteamCAPI_ISteamHTMLSurface_MouseMove(void* surface, HHTMLBrowser unBrowserHandle, int x, int y)																								{ 		 static_cast<ISteamHTMLSurface*>(surface)->MouseMove(unBrowserHandle, x, y); }
void			SteamCAPI_ISteamHTMLSurface_MouseWheel(void* surface, HHTMLBrowser unBrowserHandle, int32 nDelta)																								{ 		 static_cast<ISteamHTMLSurface*>(surface)->MouseWheel(unBrowserHandle, nDelta); }
void			SteamCAPI_ISteamHTMLSurface_KeyDown(void* surface, HHTMLBrowser unBrowserHandle, unsigned int nNativeKeyCode, EHTMLKeyModifiers eHTMLKeyModifiers)												{ 		 static_cast<ISteamHTMLSurface*>(surface)->KeyDown(unBrowserHandle, nNativeKeyCode, static_cast<ISteamHTMLSurface::EHTMLKeyModifiers>(eHTMLKeyModifiers)); }
void			SteamCAPI_ISteamHTMLSurface_KeyUp(void* surface, HHTMLBrowser unBrowserHandle, unsigned int nNativeKeyCode, EHTMLKeyModifiers eHTMLKeyModifiers)												{ 		 static_cast<ISteamHTMLSurface*>(surface)->KeyUp(unBrowserHandle, nNativeKeyCode, static_cast<ISteamHTMLSurface::EHTMLKeyModifiers>(eHTMLKeyModifiers)); }
void			SteamCAPI_ISteamHTMLSurface_KeyChar(void* surface, HHTMLBrowser unBrowserHandle, unsigned int cUnicodeChar, EHTMLKeyModifiers eHTMLKeyModifiers)												{ 		 static_cast<ISteamHTMLSurface*>(surface)->KeyChar(unBrowserHandle, cUnicodeChar, static_cast<ISteamHTMLSurface::EHTMLKeyModifiers>(eHTMLKeyModifiers)); }
void			SteamCAPI_ISteamHTMLSurface_SetHorizontalScroll(void* surface, HHTMLBrowser unBrowserHandle, unsigned int nAbsolutePixelScroll)																	{ 		 static_cast<ISteamHTMLSurface*>(surface)->SetHorizontalScroll(unBrowserHandle, nAbsolutePixelScroll); }
void			SteamCAPI_ISteamHTMLSurface_SetVerticalScroll(void* surface, HHTMLBrowser unBrowserHandle, unsigned int nAbsolutePixelScroll)																	{ 		 static_cast<ISteamHTMLSurface*>(surface)->SetVerticalScroll(unBrowserHandle, nAbsolutePixelScroll); }
void			SteamCAPI_ISteamHTMLSurface_SetKeyFocus(void* surface, HHTMLBrowser unBrowserHandle, bool bHasKeyFocus)																							{ 		 static_cast<ISteamHTMLSurface*>(surface)->SetKeyFocus(unBrowserHandle, bHasKeyFocus); }
void			SteamCAPI_ISteamHTMLSurface_ViewSource(void* surface, HHTMLBrowser unBrowserHandle)																												{ 		 static_cast<ISteamHTMLSurface*>(surface)->ViewSource(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_CopyToClipboard(void* surface, HHTMLBrowser unBrowserHandle)																										{ 		 static_cast<ISteamHTMLSurface*>(surface)->CopyToClipboard(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_PasteFromClipboard(void* surface, HHTMLBrowser unBrowserHandle)																										{ 		 static_cast<ISteamHTMLSurface*>(surface)->PasteFromClipboard(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_Find(void* surface, HHTMLBrowser unBrowserHandle, const char *pchSearchStr, bool bCurrentlyInFind, bool bReverse)													{ 		 static_cast<ISteamHTMLSurface*>(surface)->Find(unBrowserHandle, pchSearchStr, bCurrentlyInFind, bReverse); }
void			SteamCAPI_ISteamHTMLSurface_StopFind(void* surface, HHTMLBrowser unBrowserHandle)																												{ 		 static_cast<ISteamHTMLSurface*>(surface)->StopFind(unBrowserHandle); }
void			SteamCAPI_ISteamHTMLSurface_GetLinkAtPosition(void* surface,  HHTMLBrowser unBrowserHandle, int x, int y)																						{ 		 static_cast<ISteamHTMLSurface*>(surface)->GetLinkAtPosition(unBrowserHandle, x, y); }
void			SteamCAPI_ISteamHTMLSurface_SetCookie(void* surface, const char *pchHostname, const char *pchKey, const char *pchValue, const char *pchPath, RTime32 nExpires, bool bSecure, bool bHTTPOnly)	{ 		 static_cast<ISteamHTMLSurface*>(surface)->SetCookie(pchHostname, pchKey, pchValue, pchPath, nExpires, bSecure, bHTTPOnly); }
void			SteamCAPI_ISteamHTMLSurface_SetPageScaleFactor(void* surface, HHTMLBrowser unBrowserHandle, float flZoom, int nPointX, int nPointY)																{ 		 static_cast<ISteamHTMLSurface*>(surface)->SetPageScaleFactor(unBrowserHandle, flZoom, nPointX, nPointY); }
void			SteamCAPI_ISteamHTMLSurface_SetBackgroundMode(void* surface, HHTMLBrowser unBrowserHandle, bool bBackgroundMode)																				{ 		 static_cast<ISteamHTMLSurface*>(surface)->SetBackgroundMode(unBrowserHandle, bBackgroundMode); }
void			SteamCAPI_ISteamHTMLSurface_AllowStartRequest(void* surface, HHTMLBrowser unBrowserHandle, bool bAllowed)																						{ 		 static_cast<ISteamHTMLSurface*>(surface)->AllowStartRequest(unBrowserHandle, bAllowed); }
void			SteamCAPI_ISteamHTMLSurface_JSDialogResponse(void* surface, HHTMLBrowser unBrowserHandle, bool bResult)																							{ 		 static_cast<ISteamHTMLSurface*>(surface)->JSDialogResponse(unBrowserHandle, bResult); }
void			SteamCAPI_ISteamHTMLSurface_FileLoadDialogResponse(void* surface, HHTMLBrowser unBrowserHandle, const char **pchSelectedFiles)																	{ 		 static_cast<ISteamHTMLSurface*>(surface)->FileLoadDialogResponse(unBrowserHandle, pchSelectedFiles); }

//==============================================================================
//================================Steam HTTP API================================
//==============================================================================

HTTPRequestHandle 			SteamCAPI_ISteamHTTP_CreateHTTPRequest(void* http, EHTTPMethod eHTTPRequestMethod, const char *pchAbsoluteURL)																		{ return static_cast<ISteamHTTP*>(http)->CreateHTTPRequest(eHTTPRequestMethod, pchAbsoluteURL); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestContextValue(void* http, HTTPRequestHandle hRequest, unsigned long long ulContextValue)															{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestContextValue(hRequest, ulContextValue); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestNetworkActivityTimeout(void* http, HTTPRequestHandle hRequest, unsigned int unTimeoutSeconds)													{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestNetworkActivityTimeout(hRequest, unTimeoutSeconds); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestHeaderValue(void* http, HTTPRequestHandle hRequest, const char *pchHeaderName, const char *pchHeaderValue)										{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestHeaderValue(hRequest, pchHeaderName, pchHeaderValue); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestGetOrPostParameter(void* http, HTTPRequestHandle hRequest, const char *pchParamName, const char *pchParamValue)									{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestGetOrPostParameter(hRequest, pchParamName, pchParamValue); }
bool 						SteamCAPI_ISteamHTTP_SendHTTPRequest(void* http, HTTPRequestHandle hRequest, SteamAPICall_t *pCallHandle)																			{ return static_cast<ISteamHTTP*>(http)->SendHTTPRequest(hRequest, pCallHandle); }
bool 						SteamCAPI_ISteamHTTP_SendHTTPRequestAndStreamResponse(void* http, HTTPRequestHandle hRequest, SteamAPICall_t *pCallHandle)															{ return static_cast<ISteamHTTP*>(http)->SendHTTPRequestAndStreamResponse(hRequest, pCallHandle); }
bool 						SteamCAPI_ISteamHTTP_DeferHTTPRequest(void* http, HTTPRequestHandle hRequest)																										{ return static_cast<ISteamHTTP*>(http)->DeferHTTPRequest(hRequest); }
bool 						SteamCAPI_ISteamHTTP_PrioritizeHTTPRequest(void* http, HTTPRequestHandle hRequest)																									{ return static_cast<ISteamHTTP*>(http)->PrioritizeHTTPRequest(hRequest); }
bool 						SteamCAPI_ISteamHTTP_GetHTTPResponseHeaderSize(void* http, HTTPRequestHandle hRequest, const char *pchHeaderName, unsigned int *unResponseHeaderSize)								{ return static_cast<ISteamHTTP*>(http)->GetHTTPResponseHeaderSize(hRequest, pchHeaderName, unResponseHeaderSize); }
bool 						SteamCAPI_ISteamHTTP_GetHTTPResponseHeaderValue(void* http, HTTPRequestHandle hRequest, const char *pchHeaderName, unsigned char *pHeaderValueBuffer, unsigned int unBufferSize)	{ return static_cast<ISteamHTTP*>(http)->GetHTTPResponseHeaderValue(hRequest, pchHeaderName, pHeaderValueBuffer, unBufferSize); }
bool 						SteamCAPI_ISteamHTTP_GetHTTPResponseBodySize(void* http, HTTPRequestHandle hRequest, unsigned int *unBodySize)																		{ return static_cast<ISteamHTTP*>(http)->GetHTTPResponseBodySize(hRequest, unBodySize); }
bool 						SteamCAPI_ISteamHTTP_GetHTTPResponseBodyData(void* http, HTTPRequestHandle hRequest, unsigned char *pBodyDataBuffer, unsigned int unBufferSize)										{ return static_cast<ISteamHTTP*>(http)->GetHTTPResponseBodyData(hRequest, pBodyDataBuffer, unBufferSize); }
bool 						SteamCAPI_ISteamHTTP_GetHTTPStreamingResponseBodyData(void* http, HTTPRequestHandle hRequest, unsigned int cOffset, unsigned char *pBodyDataBuffer, unsigned int unBufferSize)		{ return static_cast<ISteamHTTP*>(http)->GetHTTPStreamingResponseBodyData(hRequest, cOffset, pBodyDataBuffer, unBufferSize); }
bool 						SteamCAPI_ISteamHTTP_ReleaseHTTPRequest(void* http, HTTPRequestHandle hRequest)																										{ return static_cast<ISteamHTTP*>(http)->ReleaseHTTPRequest(hRequest); }
bool 						SteamCAPI_ISteamHTTP_GetHTTPDownloadProgressPct(void* http, HTTPRequestHandle hRequest, float *pflPercentOut)																		{ return static_cast<ISteamHTTP*>(http)->GetHTTPDownloadProgressPct(hRequest, pflPercentOut); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestRawPostBody(void* http, HTTPRequestHandle hRequest, const char *pchContentType, unsigned char *pubBody, unsigned int unBodyLen)					{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestRawPostBody(hRequest, pchContentType, pubBody, unBodyLen); }
HTTPCookieContainerHandle 	SteamCAPI_ISteamHTTP_CreateCookieContainer(void* http, bool bAllowResponsesToModify)																								{ return static_cast<ISteamHTTP*>(http)->CreateCookieContainer(bAllowResponsesToModify); }
bool 						SteamCAPI_ISteamHTTP_ReleaseCookieContainer(void* http, HTTPCookieContainerHandle hCookieContainer)																					{ return static_cast<ISteamHTTP*>(http)->ReleaseCookieContainer(hCookieContainer); }
bool 						SteamCAPI_ISteamHTTP_SetCookie(void* http, HTTPCookieContainerHandle hCookieContainer, const char *pchHost, const char *pchUrl, const char *pchCookie)								{ return static_cast<ISteamHTTP*>(http)->SetCookie(hCookieContainer, pchHost, pchUrl, pchCookie); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestCookieContainer(void* http, HTTPRequestHandle hRequest, HTTPCookieContainerHandle hCookieContainer)												{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestCookieContainer(hRequest, hCookieContainer); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestUserAgentInfo(void* http, HTTPRequestHandle hRequest, const char *pchUserAgentInfo)																{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestUserAgentInfo(hRequest, pchUserAgentInfo); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestRequiresVerifiedCertificate(void* http, HTTPRequestHandle hRequest, bool bRequireVerifiedCertificate)											{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestRequiresVerifiedCertificate(hRequest, bRequireVerifiedCertificate); }
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestAbsoluteTimeoutMS(void* http, HTTPRequestHandle hRequest, unsigned int unMilliseconds)															{ return static_cast<ISteamHTTP*>(http)->SetHTTPRequestAbsoluteTimeoutMS(hRequest, unMilliseconds); }
bool 						SteamCAPI_ISteamHTTP_GetHTTPRequestWasTimedOut(void* http, HTTPRequestHandle hRequest, bool *pbWasTimedOut)																			{ return static_cast<ISteamHTTP*>(http)->GetHTTPRequestWasTimedOut(hRequest, pbWasTimedOut); }

//==============================================================================
//=============================Steam inventory API==============================
//==============================================================================

EResult	SteamCAPI_ISteamInventory_GetResultStatus(void* inventory,  SteamInventoryResult_t resultHandle)																																																															{ return static_cast<ISteamInventory*>(inventory)->GetResultStatus(resultHandle); }
bool 	SteamCAPI_ISteamInventory_GetResultItems(void* inventory,  SteamInventoryResult_t resultHandle,  SteamItemDetails_t *pOutItemsArray, unsigned int *punOutItemsArraySize)																																													{ return static_cast<ISteamInventory*>(inventory)->GetResultItems(resultHandle,  pOutItemsArray,  punOutItemsArraySize); }
unsigned int 	SteamCAPI_ISteamInventory_GetResultTimestamp(void* inventory,  SteamInventoryResult_t resultHandle)																																																													{ return static_cast<ISteamInventory*>(inventory)->GetResultTimestamp(resultHandle); }
bool 	SteamCAPI_ISteamInventory_CheckResultSteamID(void* inventory,  SteamInventoryResult_t resultHandle, CSteamID steamIDExpected)																																																								{ return static_cast<ISteamInventory*>(inventory)->CheckResultSteamID(resultHandle, steamIDExpected); }
void 	SteamCAPI_ISteamInventory_DestroyResult(void* inventory,  SteamInventoryResult_t resultHandle)																																																																{ return static_cast<ISteamInventory*>(inventory)->DestroyResult(resultHandle); }
bool 	SteamCAPI_ISteamInventory_GetAllItems(void* inventory,  SteamInventoryResult_t *pResultHandle)																																																																{ return static_cast<ISteamInventory*>(inventory)->GetAllItems(pResultHandle); }
bool 	SteamCAPI_ISteamInventory_GetItemsByID(void* inventory,  SteamInventoryResult_t *pResultHandle, const SteamItemInstanceID_t *pInstanceIDs, unsigned int unCountInstanceIDs)																																													{ return static_cast<ISteamInventory*>(inventory)->GetItemsByID(pResultHandle,  pInstanceIDs, unCountInstanceIDs); }
bool 	SteamCAPI_ISteamInventory_SerializeResult(void* inventory,  SteamInventoryResult_t resultHandle,  void *pOutBuffer, unsigned int *punOutBufferSize)																																																			{ return static_cast<ISteamInventory*>(inventory)->SerializeResult(resultHandle,  pOutBuffer,  punOutBufferSize); }
bool 	SteamCAPI_ISteamInventory_DeserializeResult(void* inventory,  SteamInventoryResult_t *pOutResultHandle, const void *pBuffer, unsigned int unBufferSize, bool bRESERVED_MUST_BE_FALSE)																																										{ return static_cast<ISteamInventory*>(inventory)->DeserializeResult(pOutResultHandle,  pBuffer, unBufferSize, bRESERVED_MUST_BE_FALSE); }
bool 	SteamCAPI_ISteamInventory_GenerateItems(void* inventory,  SteamInventoryResult_t *pResultHandle, const SteamItemDef_t *pArrayItemDefs, const unsigned int *punArrayQuantity, unsigned int unArrayLength)																																					{ return static_cast<ISteamInventory*>(inventory)->GenerateItems(pResultHandle,  pArrayItemDefs,  punArrayQuantity, unArrayLength); }
bool 	SteamCAPI_ISteamInventory_GrantPromoItems(void* inventory,  SteamInventoryResult_t *pResultHandle)																																																															{ return static_cast<ISteamInventory*>(inventory)->GrantPromoItems(pResultHandle); }
bool 	SteamCAPI_ISteamInventory_AddPromoItem(void* inventory,  SteamInventoryResult_t *pResultHandle, SteamItemDef_t itemDef)																																																										{ return static_cast<ISteamInventory*>(inventory)->AddPromoItem(pResultHandle, itemDef); }
bool 	SteamCAPI_ISteamInventory_AddPromoItems(void* inventory,  SteamInventoryResult_t *pResultHandle, const SteamItemDef_t *pArrayItemDefs, unsigned int unArrayLength)																																															{ return static_cast<ISteamInventory*>(inventory)->AddPromoItems(pResultHandle,  pArrayItemDefs, unArrayLength); }
bool 	SteamCAPI_ISteamInventory_ConsumeItem(void* inventory,  SteamInventoryResult_t *pResultHandle, SteamItemInstanceID_t itemConsume, unsigned int unQuantity)																																																	{ return static_cast<ISteamInventory*>(inventory)->ConsumeItem(pResultHandle, itemConsume, unQuantity); }
bool 	SteamCAPI_ISteamInventory_ExchangeItems(void* inventory,  SteamInventoryResult_t *pResultHandle, const SteamItemDef_t *pArrayGenerate, const unsigned int *punArrayGenerateQuantity, unsigned int unArrayGenerateLength, const SteamItemInstanceID_t *pArrayDestroy, const unsigned int *punArrayDestroyQuantity, unsigned int unArrayDestroyLength)		{ return static_cast<ISteamInventory*>(inventory)->ExchangeItems(pResultHandle,  pArrayGenerate,  punArrayGenerateQuantity, unArrayGenerateLength,  pArrayDestroy,  punArrayDestroyQuantity, unArrayDestroyLength); }
bool 	SteamCAPI_ISteamInventory_TransferItemQuantity(void* inventory,  SteamInventoryResult_t *pResultHandle, SteamItemInstanceID_t itemIdSource, unsigned int unQuantity, SteamItemInstanceID_t itemIdDest)																																						{ return static_cast<ISteamInventory*>(inventory)->TransferItemQuantity(pResultHandle, itemIdSource, unQuantity, itemIdDest); }
void 	SteamCAPI_ISteamInventory_SendItemDropHeartbeat(void* inventory)																																																																							{ return static_cast<ISteamInventory*>(inventory)->SendItemDropHeartbeat(); }
bool 	SteamCAPI_ISteamInventory_TriggerItemDrop(void* inventory,  SteamInventoryResult_t *pResultHandle, SteamItemDef_t dropListDefinition)																																																						{ return static_cast<ISteamInventory*>(inventory)->TriggerItemDrop(pResultHandle, dropListDefinition); }
bool 	SteamCAPI_ISteamInventory_TradeItems(void* inventory,  SteamInventoryResult_t *pResultHandle, CSteamID steamIDTradePartner, const SteamItemInstanceID_t *pArrayGive, const unsigned int *pArrayGiveQuantity, unsigned int nArrayGiveLength, const SteamItemInstanceID_t *pArrayGet, const unsigned int *pArrayGetQuantity, unsigned int nArrayGetLength)	{ return static_cast<ISteamInventory*>(inventory)->TradeItems(pResultHandle, steamIDTradePartner,  pArrayGive,  pArrayGiveQuantity, nArrayGiveLength,  pArrayGet,  pArrayGetQuantity, nArrayGetLength); }
bool 	SteamCAPI_ISteamInventory_LoadItemDefinitions(void* inventory)																																																																								{ return static_cast<ISteamInventory*>(inventory)->LoadItemDefinitions(); }
bool 	SteamCAPI_ISteamInventory_GetItemDefinitionIDs(void* inventory, SteamItemDef_t *pItemDefIDs, unsigned int *punItemDefIDsArraySize)																																																							{ return static_cast<ISteamInventory*>(inventory)->GetItemDefinitionIDs(pItemDefIDs,  punItemDefIDsArraySize); }
bool 	SteamCAPI_ISteamInventory_GetItemDefinitionProperty(void* inventory,  SteamItemDef_t iDefinition, const char *pchPropertyName, char *pchValueBuffer, unsigned int *punValueBufferSize)																																										{ return static_cast<ISteamInventory*>(inventory)->GetItemDefinitionProperty(iDefinition,  pchPropertyName,  pchValueBuffer,  punValueBufferSize); }


//==============================================================================
//==============================Steam Matchmaking===============================
//==============================================================================

int 			SteamCAPI_ISteamMatchmaking_GetFavoriteGameCount(void* mm)																																													{ return static_cast<ISteamMatchmaking*>(mm)->GetFavoriteGameCount(); }
bool 			SteamCAPI_ISteamMatchmaking_GetFavoriteGame(void* mm, int iGame, AppId_t *pnAppID, unsigned int *pnIP, unsigned short int *pnConnPort, unsigned short int *pnQueryPort, unsigned int *punFlags, unsigned int *pRTime32LastPlayedOnServer)	{ return static_cast<ISteamMatchmaking*>(mm)->GetFavoriteGame(iGame, pnAppID, pnIP, pnConnPort, pnQueryPort, punFlags, pRTime32LastPlayedOnServer); }
int 			SteamCAPI_ISteamMatchmaking_AddFavoriteGame(void* mm, AppId_t nAppID, unsigned int nIP, unsigned short int nConnPort, unsigned short int nQueryPort, unsigned int unFlags, unsigned int rTime32LastPlayedOnServer)							{ return static_cast<ISteamMatchmaking*>(mm)->AddFavoriteGame(nAppID, nIP, nConnPort, nQueryPort, unFlags, rTime32LastPlayedOnServer); }
bool 			SteamCAPI_ISteamMatchmaking_RemoveFavoriteGame(void* mm, AppId_t nAppID, unsigned int nIP, unsigned short int nConnPort, unsigned short int nQueryPort, unsigned int unFlags)																{ return static_cast<ISteamMatchmaking*>(mm)->RemoveFavoriteGame(nAppID, nIP, nConnPort, nQueryPort, unFlags); }
SteamAPICall_t 	SteamCAPI_ISteamMatchmaking_RequestLobbyList(void* mm)																																														{ return static_cast<ISteamMatchmaking*>(mm)->RequestLobbyList(); }
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListStringFilter(void* mm, const char *pchKeyToMatch, const char *pchValueToMatch, ELobbyComparison eComparisonType)																				{ return static_cast<ISteamMatchmaking*>(mm)->AddRequestLobbyListStringFilter(pchKeyToMatch, pchValueToMatch, eComparisonType); }
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListNumericalFilter(void* mm, const char *pchKeyToMatch, int nValueToMatch, ELobbyComparison eComparisonType)																					{ return static_cast<ISteamMatchmaking*>(mm)->AddRequestLobbyListNumericalFilter(pchKeyToMatch, nValueToMatch, eComparisonType); }
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListNearValueFilter(void* mm, const char *pchKeyToMatch, int nValueToBeCloseTo)																													{ return static_cast<ISteamMatchmaking*>(mm)->AddRequestLobbyListNearValueFilter(pchKeyToMatch, nValueToBeCloseTo); }
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable(void* mm, int nSlotsAvailable)																																			{ return static_cast<ISteamMatchmaking*>(mm)->AddRequestLobbyListFilterSlotsAvailable(nSlotsAvailable); }
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListDistanceFilter(void* mm, ELobbyDistanceFilter eLobbyDistanceFilter)																															{ return static_cast<ISteamMatchmaking*>(mm)->AddRequestLobbyListDistanceFilter(eLobbyDistanceFilter); }
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListResultCountFilter(void* mm, int cMaxResults)																																					{ return static_cast<ISteamMatchmaking*>(mm)->AddRequestLobbyListResultCountFilter(cMaxResults); }
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter(void* mm, CSteamID steamIDLobby)																																		{ return static_cast<ISteamMatchmaking*>(mm)->AddRequestLobbyListCompatibleMembersFilter(steamIDLobby); }
CSteamID 		SteamCAPI_ISteamMatchmaking_GetLobbyByIndex(void* mm, int iLobby)																																											{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyByIndex(iLobby); }
SteamAPICall_t 	SteamCAPI_ISteamMatchmaking_CreateLobby(void* mm, ELobbyType eLobbyType, int cMaxMembers)																																					{ return static_cast<ISteamMatchmaking*>(mm)->CreateLobby(eLobbyType, cMaxMembers); }
SteamAPICall_t 	SteamCAPI_ISteamMatchmaking_JoinLobby(void* mm, CSteamID steamIDLobby)																																										{ return static_cast<ISteamMatchmaking*>(mm)->JoinLobby(steamIDLobby); }
void 			SteamCAPI_ISteamMatchmaking_LeaveLobby(void* mm, CSteamID steamIDLobby)																																										{ return static_cast<ISteamMatchmaking*>(mm)->LeaveLobby(steamIDLobby); }
bool 			SteamCAPI_ISteamMatchmaking_InviteUserToLobby(void* mm, CSteamID steamIDLobby, CSteamID steamIDInvitee)																																		{ return static_cast<ISteamMatchmaking*>(mm)->InviteUserToLobby(steamIDLobby, steamIDInvitee); }
int 			SteamCAPI_ISteamMatchmaking_GetNumLobbyMembers(void* mm, CSteamID steamIDLobby)																																								{ return static_cast<ISteamMatchmaking*>(mm)->GetNumLobbyMembers(steamIDLobby); }
CSteamID 		SteamCAPI_ISteamMatchmaking_GetLobbyMemberByIndex(void* mm, CSteamID steamIDLobby, int iMember)																																				{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyMemberByIndex(steamIDLobby, iMember); }
const char*		SteamCAPI_ISteamMatchmaking_GetLobbyData(void* mm, CSteamID steamIDLobby, const char *pchKey)																																				{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyData(steamIDLobby, pchKey); }
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyData(void* mm, CSteamID steamIDLobby, const char *pchKey, const char *pchValue)																															{ return static_cast<ISteamMatchmaking*>(mm)->SetLobbyData(steamIDLobby, pchKey, pchValue); }
int 			SteamCAPI_ISteamMatchmaking_GetLobbyDataCount(void* mm, CSteamID steamIDLobby)																																								{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyDataCount(steamIDLobby); }
bool 			SteamCAPI_ISteamMatchmaking_GetLobbyDataByIndex(void* mm, CSteamID steamIDLobby, int iLobbyData, char *pchKey, int cchKeyBufferSize, char *pchValue, int cchValueBufferSize)																{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyDataByIndex(steamIDLobby, iLobbyData, pchKey, cchKeyBufferSize, pchValue, cchValueBufferSize); }
bool 			SteamCAPI_ISteamMatchmaking_DeleteLobbyData(void* mm, CSteamID steamIDLobby, const char *pchKey)																																			{ return static_cast<ISteamMatchmaking*>(mm)->DeleteLobbyData(steamIDLobby, pchKey); }
const char*		SteamCAPI_ISteamMatchmaking_GetLobbyMemberData(void* mm, CSteamID steamIDLobby, CSteamID steamIDUser, const char *pchKey)																													{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyMemberData(steamIDLobby, steamIDUser, pchKey); }
void 			SteamCAPI_ISteamMatchmaking_SetLobbyMemberData(void* mm, CSteamID steamIDLobby, const char *pchKey, const char *pchValue)																													{ return static_cast<ISteamMatchmaking*>(mm)->SetLobbyMemberData(steamIDLobby, pchKey, pchValue); }
bool 			SteamCAPI_ISteamMatchmaking_SendLobbyChatMsg(void* mm, CSteamID steamIDLobby, const void *pvMsgBody, int cubMsgBody)																														{ return static_cast<ISteamMatchmaking*>(mm)->SendLobbyChatMsg(steamIDLobby, pvMsgBody, cubMsgBody); }
int 			SteamCAPI_ISteamMatchmaking_GetLobbyChatEntry(void* mm, CSteamID steamIDLobby, int iChatID, CSteamID *pSteamIDUser, void *pvData, int cubData, EChatEntryType *peChatEntryType)																{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyChatEntry(steamIDLobby, iChatID, pSteamIDUser, pvData, cubData, peChatEntryType); }
bool 			SteamCAPI_ISteamMatchmaking_RequestLobbyData(void* mm, CSteamID steamIDLobby)																																								{ return static_cast<ISteamMatchmaking*>(mm)->RequestLobbyData(steamIDLobby); }
void 			SteamCAPI_ISteamMatchmaking_SetLobbyGameServer(void* mm, CSteamID steamIDLobby, unsigned int unGameServerIP, unsigned short int unGameServerPort, CSteamID steamIDGameServer)																{ return static_cast<ISteamMatchmaking*>(mm)->SetLobbyGameServer(steamIDLobby, unGameServerIP, unGameServerPort, steamIDGameServer); }
bool 			SteamCAPI_ISteamMatchmaking_GetLobbyGameServer(void* mm, CSteamID steamIDLobby, unsigned int *punGameServerIP, unsigned short int *punGameServerPort, CSteamID *psteamIDGameServer)															{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyGameServer(steamIDLobby, punGameServerIP, punGameServerPort, psteamIDGameServer); }
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyMemberLimit(void* mm, CSteamID steamIDLobby, int cMaxMembers)																																			{ return static_cast<ISteamMatchmaking*>(mm)->SetLobbyMemberLimit(steamIDLobby, cMaxMembers); }
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyType(void* mm, CSteamID steamIDLobby, ELobbyType eLobbyType)																																			{ return static_cast<ISteamMatchmaking*>(mm)->SetLobbyType(steamIDLobby, eLobbyType); }
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyJoinable(void* mm, CSteamID steamIDLobby, bool bLobbyJoinable)																																			{ return static_cast<ISteamMatchmaking*>(mm)->SetLobbyJoinable(steamIDLobby, bLobbyJoinable); }
CSteamID 		SteamCAPI_ISteamMatchmaking_GetLobbyOwner(void* mm, CSteamID steamIDLobby)																																									{ return static_cast<ISteamMatchmaking*>(mm)->GetLobbyOwner(steamIDLobby); }
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyOwner(void* mm, CSteamID steamIDLobby, CSteamID steamIDNewOwner)																																		{ return static_cast<ISteamMatchmaking*>(mm)->SetLobbyOwner(steamIDLobby, steamIDNewOwner); }
bool 			SteamCAPI_ISteamMatchmaking_SetLinkedLobby(void* mm, CSteamID steamIDLobby, CSteamID steamIDLobbyDependent)																																	{ return static_cast<ISteamMatchmaking*>(mm)->SetLinkedLobby(steamIDLobby, steamIDLobbyDependent); }








//==============================================================================
//================================Steam UGC API================================
//==============================================================================


void* CSteamUGC(){
	return SteamUGC();
}

UGCQueryHandle_t 	SteamUGC_CreateQueryUserUGCRequest(void* ugc, AccountID_t unAccountID, EUserUGCList eListType, EUGCMatchingUGCType eMatchingUGCType, EUserUGCListSortOrder eSortOrder, AppId_t nCreatorAppID, AppId_t nConsumerAppID, unsigned int unPage)	{ return static_cast<ISteamUGC*>(ugc)->CreateQueryUserUGCRequest(unAccountID, eListType, eMatchingUGCType, eSortOrder, nCreatorAppID, nConsumerAppID, unPage); }
UGCQueryHandle_t 	SteamUGC_CreateQueryAllUGCRequest(void* ugc, EUGCQuery eQueryType, EUGCMatchingUGCType eMatchingeMatchingUGCTypeFileType, AppId_t nCreatorAppID, AppId_t nConsumerAppID, unsigned int unPage)												{ return static_cast<ISteamUGC*>(ugc)->CreateQueryAllUGCRequest(eQueryType, eMatchingeMatchingUGCTypeFileType, nCreatorAppID, nConsumerAppID, unPage); }
UGCQueryHandle_t 	SteamUGC_CreateQueryUGCDetailsRequest(void* ugc, PublishedFileId_t* pvecPublishedFileID, unsigned int unNumPublishedFileIDs)																												{ return static_cast<ISteamUGC*>(ugc)->CreateQueryUGCDetailsRequest(pvecPublishedFileID, unNumPublishedFileIDs); }
SteamAPICall_t 		SteamUGC_SendQueryUGCRequest(void* ugc, UGCQueryHandle_t handle)																																											{ return static_cast<ISteamUGC*>(ugc)->SendQueryUGCRequest(handle); }
bool 				SteamUGC_GetQueryUGCResult(void* ugc, UGCQueryHandle_t handle, unsigned int index, SteamUGCDetails_t *pDetails)																															{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCResult(handle, index, pDetails); }
bool 				SteamUGC_GetQueryUGCPreviewURL(void* ugc, UGCQueryHandle_t handle, unsigned int index, char* pchURL, unsigned int cchURLSize)																												{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCPreviewURL(handle, index, pchURL, cchURLSize); }
bool 				SteamUGC_GetQueryUGCMetadata(void* ugc, UGCQueryHandle_t handle, unsigned int index, char* pchMetadata, unsigned int cchMetadatasize)																										{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCMetadata(handle, index, pchMetadata, cchMetadatasize); }
bool 				SteamUGC_GetQueryUGCChildren(void* ugc, UGCQueryHandle_t handle, unsigned int index, PublishedFileId_t* pvecPublishedFileID, unsigned int cMaxEntries)																						{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCChildren(handle, index, pvecPublishedFileID, cMaxEntries); }
bool 				SteamUGC_GetQueryUGCStatistic(void* ugc, UGCQueryHandle_t handle, unsigned int index, EItemStatistic eStatType, unsigned int *pStatValue)																									{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCStatistic(handle, index, eStatType, pStatValue); }
unsigned int 		SteamUGC_GetQueryUGCNumAdditionalPreviews(void* ugc, UGCQueryHandle_t handle, unsigned int index)																																			{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCNumAdditionalPreviews(handle, index); }
bool 				SteamUGC_GetQueryUGCAdditionalPreview(void* ugc, UGCQueryHandle_t handle, unsigned int index, unsigned int previewIndex, char* pchURLOrVideoID, unsigned int cchURLSize, bool *pbIsImage)													{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCAdditionalPreview(handle, index, previewIndex, pchURLOrVideoID, cchURLSize, pbIsImage); }
unsigned int 		SteamUGC_GetQueryUGCNumKeyValueTags(void* ugc, UGCQueryHandle_t handle, unsigned int index)																																				{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCNumKeyValueTags(handle, index); }
bool 				SteamUGC_GetQueryUGCKeyValueTag(void* ugc, UGCQueryHandle_t handle, unsigned int index, unsigned int keyValueTagIndex, char* pchKey, unsigned int cchKeySize, char* pchValue, unsigned int cchValueSize)									{ return static_cast<ISteamUGC*>(ugc)->GetQueryUGCKeyValueTag(handle, index, keyValueTagIndex, pchKey, cchKeySize, pchValue, cchValueSize); }
bool 				SteamUGC_ReleaseQueryUGCRequest(void* ugc, UGCQueryHandle_t handle)																																										{ return static_cast<ISteamUGC*>(ugc)->ReleaseQueryUGCRequest(handle); }
bool 				SteamUGC_AddRequiredTag(void* ugc, UGCQueryHandle_t handle, const char* pTagName)																																							{ return static_cast<ISteamUGC*>(ugc)->AddRequiredTag(handle, pTagName); }
bool 				SteamUGC_AddExcludedTag(void* ugc, UGCQueryHandle_t handle, const char* pTagName)																																							{ return static_cast<ISteamUGC*>(ugc)->AddExcludedTag(handle, pTagName); }
bool 				SteamUGC_SetReturnKeyValueTags(void* ugc, UGCQueryHandle_t handle, bool bReturnKeyValueTags)																																				{ return static_cast<ISteamUGC*>(ugc)->SetReturnKeyValueTags(handle, bReturnKeyValueTags); }
bool 				SteamUGC_SetReturnLongDescription(void* ugc, UGCQueryHandle_t handle, bool bReturnLongDescription)																																			{ return static_cast<ISteamUGC*>(ugc)->SetReturnLongDescription(handle, bReturnLongDescription); }
bool 				SteamUGC_SetReturnMetadata(void* ugc, UGCQueryHandle_t handle, bool bReturnMetadata)																																						{ return static_cast<ISteamUGC*>(ugc)->SetReturnMetadata(handle, bReturnMetadata); }
bool 				SteamUGC_SetReturnChildren(void* ugc, UGCQueryHandle_t handle, bool bReturnChildren)																																						{ return static_cast<ISteamUGC*>(ugc)->SetReturnChildren(handle, bReturnChildren); }
bool 				SteamUGC_SetReturnAdditionalPreviews(void* ugc, UGCQueryHandle_t handle, bool bReturnAdditionalPreviews)																																	{ return static_cast<ISteamUGC*>(ugc)->SetReturnAdditionalPreviews(handle, bReturnAdditionalPreviews); }
bool 				SteamUGC_SetReturnTotalOnly(void* ugc, UGCQueryHandle_t handle, bool bReturnTotalOnly)																																						{ return static_cast<ISteamUGC*>(ugc)->SetReturnTotalOnly(handle, bReturnTotalOnly); }
bool 				SteamUGC_SetLanguage(void* ugc, UGCQueryHandle_t handle, const char* pchLanguage)																																							{ return static_cast<ISteamUGC*>(ugc)->SetLanguage(handle, pchLanguage); }
bool 				SteamUGC_SetAllowCachedResponse(void* ugc, UGCQueryHandle_t handle, unsigned int unMaxAgeSeconds)																																			{ return static_cast<ISteamUGC*>(ugc)->SetAllowCachedResponse(handle, unMaxAgeSeconds); }
bool 				SteamUGC_SetCloudFileNameFilter(void* ugc, UGCQueryHandle_t handle, const char* pMatchCloudFileName)																																		{ return static_cast<ISteamUGC*>(ugc)->SetCloudFileNameFilter(handle, pMatchCloudFileName); }
bool 				SteamUGC_SetMatchAnyTag(void* ugc, UGCQueryHandle_t handle, bool bMatchAnyTag)																																								{ return static_cast<ISteamUGC*>(ugc)->SetMatchAnyTag(handle, bMatchAnyTag); }
bool 				SteamUGC_SetSearchText(void* ugc, UGCQueryHandle_t handle, const char* pSearchText)																																						{ return static_cast<ISteamUGC*>(ugc)->SetSearchText(handle, pSearchText); }
bool 				SteamUGC_SetRankedByTrendDays(void* ugc, UGCQueryHandle_t handle, unsigned int unDays)																																						{ return static_cast<ISteamUGC*>(ugc)->SetRankedByTrendDays(handle, unDays); }
bool 				SteamUGC_AddRequiredKeyValueTag(void* ugc, UGCQueryHandle_t handle, const char* pKey, const char* pValue)																																	{ return static_cast<ISteamUGC*>(ugc)->AddRequiredKeyValueTag(handle, pKey, pValue); }
SteamAPICall_t 		SteamUGC_RequestUGCDetails(void* ugc, PublishedFileId_t nPublishedFileID, unsigned int unMaxAgeSeconds)																																	{ return static_cast<ISteamUGC*>(ugc)->RequestUGCDetails(nPublishedFileID, unMaxAgeSeconds); }
SteamAPICall_t 		SteamUGC_CreateItem(void* ugc, AppId_t nConsumerAppId, EWorkshopFileType eFileType)																																						{ return static_cast<ISteamUGC*>(ugc)->CreateItem(nConsumerAppId, eFileType); }
UGCUpdateHandle_t 	SteamUGC_StartItemUpdate(void* ugc, AppId_t nConsumerAppId, PublishedFileId_t nPublishedFileID)																																			{ return static_cast<ISteamUGC*>(ugc)->StartItemUpdate(nConsumerAppId, nPublishedFileID); }
bool 				SteamUGC_SetItemTitle(void* ugc, UGCUpdateHandle_t handle, const char* pchTitle)																																							{ return static_cast<ISteamUGC*>(ugc)->SetItemTitle(handle, pchTitle); }
bool 				SteamUGC_SetItemDescription(void* ugc, UGCUpdateHandle_t handle, const char* pchDescription)																																				{ return static_cast<ISteamUGC*>(ugc)->SetItemDescription(handle, pchDescription); }
bool 				SteamUGC_SetItemUpdateLanguage(void* ugc, UGCUpdateHandle_t handle, const char* pchLanguage)																																				{ return static_cast<ISteamUGC*>(ugc)->SetItemUpdateLanguage(handle, pchLanguage); }
bool 				SteamUGC_SetItemMetadata(void* ugc, UGCUpdateHandle_t handle, const char* pchMetaData)																																						{ return static_cast<ISteamUGC*>(ugc)->SetItemMetadata(handle, pchMetaData); }
bool 				SteamUGC_SetItemVisibility(void* ugc, UGCUpdateHandle_t handle, ERemoteStoragePublishedFileVisibility eVisibility)																															{ return static_cast<ISteamUGC*>(ugc)->SetItemVisibility(handle, eVisibility); }
bool 				SteamUGC_SetItemTags(void* ugc, UGCUpdateHandle_t updateHandle, const SteamParamStringArray_t *pTags)																																		{ return static_cast<ISteamUGC*>(ugc)->SetItemTags(updateHandle, pTags); }
bool 				SteamUGC_SetItemContent(void* ugc, UGCUpdateHandle_t handle, const char* pszContentFolder)																																					{ return static_cast<ISteamUGC*>(ugc)->SetItemContent(handle, pszContentFolder); }
bool 				SteamUGC_SetItemPreview(void* ugc, UGCUpdateHandle_t handle, const char* pszPreviewFile)																																					{ return static_cast<ISteamUGC*>(ugc)->SetItemPreview(handle, pszPreviewFile); }
bool 				SteamUGC_RemoveItemKeyValueTags(void* ugc, UGCUpdateHandle_t handle, const char* pchKey)																																					{ return static_cast<ISteamUGC*>(ugc)->RemoveItemKeyValueTags(handle, pchKey); }
bool 				SteamUGC_AddItemKeyValueTag(void* ugc, UGCUpdateHandle_t handle, const char* pchKey, const char* pchValue)																																	{ return static_cast<ISteamUGC*>(ugc)->AddItemKeyValueTag(handle, pchKey, pchValue); }
SteamAPICall_t 		SteamUGC_SubmitItemUpdate(void* ugc, UGCUpdateHandle_t handle, const char* pchChangeNote)																																					{ return static_cast<ISteamUGC*>(ugc)->SubmitItemUpdate(handle, pchChangeNote); }
EItemUpdateStatus 	SteamUGC_GetItemUpdateProgress(void* ugc, UGCUpdateHandle_t handle, uint64 *punBytesProcessed, uint64* punBytesTotal)																														{ return static_cast<ISteamUGC*>(ugc)->GetItemUpdateProgress(handle, punBytesProcessed, punBytesTotal); }
SteamAPICall_t 		SteamUGC_SetUserItemVote(void* ugc, PublishedFileId_t nPublishedFileID, bool bVoteUp)																																						{ return static_cast<ISteamUGC*>(ugc)->SetUserItemVote(nPublishedFileID, bVoteUp); }
SteamAPICall_t 		SteamUGC_GetUserItemVote(void* ugc, PublishedFileId_t nPublishedFileID)																																									{ return static_cast<ISteamUGC*>(ugc)->GetUserItemVote(nPublishedFileID); }
SteamAPICall_t 		SteamUGC_AddItemToFavorites(void* ugc, AppId_t nAppId, PublishedFileId_t nPublishedFileID)																																					{ return static_cast<ISteamUGC*>(ugc)->AddItemToFavorites(nAppId, nPublishedFileID); }
SteamAPICall_t 		SteamUGC_RemoveItemFromFavorites(void* ugc, AppId_t nAppId, PublishedFileId_t nPublishedFileID)																																			{ return static_cast<ISteamUGC*>(ugc)->RemoveItemFromFavorites(nAppId, nPublishedFileID); }
SteamAPICall_t 		SteamUGC_SubscribeItem(void* ugc, PublishedFileId_t nPublishedFileID)																																										{ return static_cast<ISteamUGC*>(ugc)->SubscribeItem(nPublishedFileID); }
SteamAPICall_t 		SteamUGC_UnsubscribeItem(void* ugc, PublishedFileId_t nPublishedFileID)																																									{ return static_cast<ISteamUGC*>(ugc)->UnsubscribeItem(nPublishedFileID); }
unsigned int 		SteamUGC_GetNumSubscribedItems(void* ugc)																																																	{ return static_cast<ISteamUGC*>(ugc)->GetNumSubscribedItems(); }
unsigned int 		SteamUGC_GetSubscribedItems(void* ugc, PublishedFileId_t* pvecPublishedFileID, unsigned int cMaxEntries)																																	{ return static_cast<ISteamUGC*>(ugc)->GetSubscribedItems(pvecPublishedFileID, cMaxEntries); }
unsigned int 		SteamUGC_GetItemState(void* ugc, PublishedFileId_t nPublishedFileID)																																										{ return static_cast<ISteamUGC*>(ugc)->GetItemState(nPublishedFileID); }
bool 				SteamUGC_GetItemInstallInfo(void* ugc, PublishedFileId_t nPublishedFileID, uint64 *punSizeOnDisk, char* pchFolder, unsigned int cchFolderSize, unsigned int *punTimeStamp)																	{ return static_cast<ISteamUGC*>(ugc)->GetItemInstallInfo(nPublishedFileID, punSizeOnDisk, pchFolder, cchFolderSize, punTimeStamp); }
bool 				SteamUGC_GetItemDownloadInfo(void* ugc, PublishedFileId_t nPublishedFileID, uint64 *punBytesDownloaded, uint64 *punBytesTotal)																												{ return static_cast<ISteamUGC*>(ugc)->GetItemDownloadInfo(nPublishedFileID, punBytesDownloaded, punBytesTotal); }
bool 				SteamUGC_DownloadItem(void* ugc, PublishedFileId_t nPublishedFileID, bool bHighPriority)																																					{ return static_cast<ISteamUGC*>(ugc)->DownloadItem(nPublishedFileID, bHighPriority); }





//==============================================================================
//============================Steam game server API=============================
//==============================================================================
void* CSteamGameServer() { return SteamGameServer(); }

bool 						ISteamGameServer_InitGameServer(void* server, unsigned int unIP, unsigned short int usGamePort, unsigned short int usQueryPort, unsigned int unFlags, AppId_t nGameAppId, const char *pchVersionString)	{ return static_cast<ISteamGameServer*>(server)->InitGameServer(unIP,usGamePort,usQueryPort,unFlags,nGameAppId,pchVersionString); }
void 						ISteamGameServer_SetProduct(void* server, const char *pszProduct) 																																		{ return static_cast<ISteamGameServer*>(server)->SetProduct(pszProduct); }
void 						ISteamGameServer_SetGameDescription(void* server, const char *pszGameDescription) 																														{ return static_cast<ISteamGameServer*>(server)->SetGameDescription(pszGameDescription); }
void 						ISteamGameServer_SetModDir(void* server, const char *pszModDir) 																																		{ return static_cast<ISteamGameServer*>(server)->SetModDir(pszModDir); }
void 						ISteamGameServer_SetDedicatedServer(void* server, bool bDedicated) 																																		{ return static_cast<ISteamGameServer*>(server)->SetDedicatedServer(bDedicated); }
void 						ISteamGameServer_LogOn(void* server, const char *pszToken) 																																				{ return static_cast<ISteamGameServer*>(server)->LogOn(pszToken); }
void 						ISteamGameServer_LogOnAnonymous(void* server) 																																							{ return static_cast<ISteamGameServer*>(server)->LogOnAnonymous(); }
void 						ISteamGameServer_LogOff(void* server) 																																									{ return static_cast<ISteamGameServer*>(server)->LogOff(); }
bool 						ISteamGameServer_BLoggedOn(void* server) 																																								{ return static_cast<ISteamGameServer*>(server)->BLoggedOn(); }
bool 						ISteamGameServer_BSecure(void* server) 																																									{ return static_cast<ISteamGameServer*>(server)->BSecure(); }
CSteamID 					ISteamGameServer_GetSteamID(void* server) 																																								{ return static_cast<ISteamGameServer*>(server)->GetSteamID(); }
bool 						ISteamGameServer_WasRestartRequested(void* server) 																																						{ return static_cast<ISteamGameServer*>(server)->WasRestartRequested(); }
void 						ISteamGameServer_SetMaxPlayerCount(void* server, int cPlayersMax) 																																		{ return static_cast<ISteamGameServer*>(server)->SetMaxPlayerCount(cPlayersMax); }
void 						ISteamGameServer_SetBotPlayerCount(void* server, int cBotplayers) 																																		{ return static_cast<ISteamGameServer*>(server)->SetBotPlayerCount(cBotplayers); }
void 						ISteamGameServer_SetServerName(void* server, const char *pszServerName) 																																{ return static_cast<ISteamGameServer*>(server)->SetServerName(pszServerName); }
void 						ISteamGameServer_SetMapName(void* server, const char *pszMapName) 																																		{ return static_cast<ISteamGameServer*>(server)->SetMapName(pszMapName); }
void 						ISteamGameServer_SetPasswordProtected(void* server, bool bPasswordProtected) 																															{ return static_cast<ISteamGameServer*>(server)->SetPasswordProtected(bPasswordProtected); }
void 						ISteamGameServer_SetSpectatorPort(void* server, unsigned short int unSpectatorPort) 																													{ return static_cast<ISteamGameServer*>(server)->SetSpectatorPort(unSpectatorPort); }
void 						ISteamGameServer_SetSpectatorServerName(void* server, const char *pszSpectatorServerName) 																												{ return static_cast<ISteamGameServer*>(server)->SetSpectatorServerName(pszSpectatorServerName); }
void 						ISteamGameServer_ClearAllKeyValues(void* server) 																																						{ return static_cast<ISteamGameServer*>(server)->ClearAllKeyValues(); }
void 						ISteamGameServer_SetKeyValue(void* server, const char *pKey, const char *pValue) 																														{ return static_cast<ISteamGameServer*>(server)->SetKeyValue(pKey, pValue); }
void 						ISteamGameServer_SetGameTags(void* server, const char *pchGameTags) 																																	{ return static_cast<ISteamGameServer*>(server)->SetGameTags(pchGameTags); }
void 						ISteamGameServer_SetGameData(void* server, const char *pchGameData) 																																	{ return static_cast<ISteamGameServer*>(server)->SetGameData(pchGameData); }
void 						ISteamGameServer_SetRegion(void* server, const char *pszRegion) 																																		{ return static_cast<ISteamGameServer*>(server)->SetRegion(pszRegion); }
bool 						ISteamGameServer_SendUserConnectAndAuthenticate(void* server, unsigned int unIPClient, const void *pvAuthBlob, unsigned int cubAuthBlobSize, CSteamID *pSteamIDUser) 									{ return static_cast<ISteamGameServer*>(server)->SendUserConnectAndAuthenticate(unIPClient, pvAuthBlob, cubAuthBlobSize, pSteamIDUser); }
CSteamID 					ISteamGameServer_CreateUnauthenticatedUserConnection(void* server) 																																		{ return static_cast<ISteamGameServer*>(server)->CreateUnauthenticatedUserConnection(); }
void 						ISteamGameServer_SendUserDisconnect(void* server, CSteamID steamIDUser) 																																{ return static_cast<ISteamGameServer*>(server)->SendUserDisconnect(steamIDUser); }
bool 						ISteamGameServer_BUpdateUserData(void* server, CSteamID steamIDUser, const char *pchPlayerName, unsigned int uScore) 																					{ return static_cast<ISteamGameServer*>(server)->BUpdateUserData(steamIDUser, pchPlayerName, uScore); }
HAuthTicket 				ISteamGameServer_GetAuthSessionTicket(void* server, void *pTicket, int cbMaxTicket, unsigned int *pcbTicket) 																							{ return static_cast<ISteamGameServer*>(server)->GetAuthSessionTicket(pTicket, cbMaxTicket, pcbTicket); }
EBeginAuthSessionResult 	ISteamGameServer_BeginAuthSession(void* server, const void *pAuthTicket, int cbAuthTicket, CSteamID steamID) 																							{ return static_cast<ISteamGameServer*>(server)->BeginAuthSession(pAuthTicket, cbAuthTicket, steamID); }
void 						ISteamGameServer_EndAuthSession(void* server, CSteamID steamID) 																																		{ return static_cast<ISteamGameServer*>(server)->EndAuthSession(steamID); }
void 						ISteamGameServer_CancelAuthTicket(void* server, HAuthTicket hAuthTicket) 																																{ return static_cast<ISteamGameServer*>(server)->CancelAuthTicket(hAuthTicket); }
EUserHasLicenseForAppResult ISteamGameServer_UserHasLicenseForApp(void* server, CSteamID steamID, AppId_t appID) 																													{ return static_cast<ISteamGameServer*>(server)->UserHasLicenseForApp(steamID, appID); }
bool 						ISteamGameServer_RequestUserGroupStatus(void* server, CSteamID steamIDUser, CSteamID steamIDGroup) 																										{ return static_cast<ISteamGameServer*>(server)->RequestUserGroupStatus(steamIDUser, steamIDGroup); }
void 						ISteamGameServer_GetGameplayStats(void* server) 																																						{ return static_cast<ISteamGameServer*>(server)->GetGameplayStats(); }
SteamAPICall_t 				ISteamGameServer_GetServerReputation(void* server) 																																						{ return static_cast<ISteamGameServer*>(server)->GetServerReputation(); }
unsigned int 				ISteamGameServer_GetPublicIP(void* server) 																																								{ return static_cast<ISteamGameServer*>(server)->GetPublicIP(); }
bool 						ISteamGameServer_HandleIncomingPacket(void* server, const void *pData, int cbData, unsigned int srcIP, unsigned short int srcPort) 																		{ return static_cast<ISteamGameServer*>(server)->HandleIncomingPacket(pData, cbData, srcIP, srcPort); }
int 						ISteamGameServer_GetNextOutgoingPacket(void* server, void *pOut, int cbMaxOut, unsigned int *pNetAdr, unsigned short int *pPort) 																		{ return static_cast<ISteamGameServer*>(server)->GetNextOutgoingPacket(pOut, cbMaxOut, pNetAdr, pPort); }
void 						ISteamGameServer_EnableHeartbeats(void* server, bool bActive) 																																			{ return static_cast<ISteamGameServer*>(server)->EnableHeartbeats(bActive); }
void 						ISteamGameServer_SetHeartbeatInterval(void* server, int iHeartbeatInterval) 																															{ return static_cast<ISteamGameServer*>(server)->SetHeartbeatInterval(iHeartbeatInterval); }
void 						ISteamGameServer_ForceHeartbeat(void* server) 																																							{ return static_cast<ISteamGameServer*>(server)->ForceHeartbeat(); }
SteamAPICall_t 				ISteamGameServer_AssociateWithClan(void* server, CSteamID steamIDClan) 																																	{ return static_cast<ISteamGameServer*>(server)->AssociateWithClan(steamIDClan); }
SteamAPICall_t 				ISteamGameServer_ComputeNewPlayerCompatibility(void* server, CSteamID steamIDNewPlayer) 																												{ return static_cast<ISteamGameServer*>(server)->ComputeNewPlayerCompatibility(steamIDNewPlayer); }




}