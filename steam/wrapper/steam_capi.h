#ifndef STEAM_C_API_H
#define STEAM_C_API_H

#include <stdbool.h>
#include <stdint.h>


// We need ifdef c++ because go will complain about seeing the `extern "C" {`
#ifdef _cplusplus
extern "C" { 
#endif

	#define k_cchPublishedDocumentTitleMax 129
	#define k_cchPublishedDocumentDescriptionMax 8000
	#define k_cchTagListMax 1025
	#define k_cchFilenameMax 260
	#define k_cchPublishedFileURLMax 256


	typedef int EBeginAuthSessionResult;
	typedef int EUserHasLicenseForAppResult;

	typedef int EResult;
	typedef int EWorkshopFileType;
	typedef int EHTMLMouseButton;
	typedef int EHTMLKeyModifiers;
	typedef void* CSteamID;
	typedef void* CGameID;
	typedef uint64_t SteamAPICall_t;
	typedef int EFriendRelationship;
	typedef int EPersonaState;
	typedef int16_t FriendsGroupID_t;
	typedef uint32_t AppId_t;
	typedef int EOverlayToStoreFlag;
	typedef int EChatEntryType;
	typedef int ERemoteStoragePublishedFileVisibility;
	typedef uint32_t DepotId_t;
	typedef int HSteamPipe;
	typedef int HSteamUser;
	typedef int EAccountType;
	typedef int EGCResults;
	typedef unsigned int HHTMLBrowser;
	typedef int EMatchMakingServerResponse;
	typedef void* HServerListRequest;
	typedef int ELobbyType;
	typedef int ELobbyDistanceFilter;
	typedef int ELobbyComparison;
	typedef int SteamItemDef_t;
	typedef unsigned long long SteamItemInstanceID_t;
	typedef int SteamInventoryResult_t;

	typedef unsigned long long ControllerActionSetHandle_t;
	typedef unsigned long long ControllerDigitalActionHandle_t;
	typedef unsigned long long ControllerAnalogActionHandle_t;
	typedef int EControllerSourceMode;
	typedef int EControllerActionOrigin;
	typedef unsigned long long ControllerHandle_t;
	typedef struct {
		// The current state of this action; will be true if currently pressed
		bool bState;
		
		// Whether or not this action is currently available to be bound in the active action set
		bool bActive;
	} ControllerDigitalActionData_t;
	typedef struct {
		// Type of data coming from this action, this will match what got specified in the action set
		EControllerSourceMode eMode;
		
		// The current state of this action; will be delta updates for mouse actions
		float x, y;
		
		// Whether or not this action is currently available to be bound in the active action set
		bool bActive;
	} ControllerAnalogActionData_t;



	typedef struct {
		SteamItemInstanceID_t m_itemId;
		SteamItemDef_t m_iDefinition;
		unsigned short m_unQuantity;
		unsigned short m_unFlags; // see ESteamItemFlags
	} SteamItemDetails_t;

	typedef uint32_t HTTPRequestHandle;
	typedef uint32_t HTTPCookieContainerHandle;
	typedef int EHTTPMethod;
	typedef uint32_t RTime32;

	typedef struct {
		CGameID m_gameID;
		uint32_t m_unGameIP;
		unsigned short m_usGamePort;
		unsigned short m_usQueryPort;
		CSteamID m_steamIDLobby;
	} FriendGameInfo_t;

	typedef uint32_t HAuthTicket;

	typedef uint64_t UGCUpdateHandle_t;
	typedef uint64_t UGCQueryHandle_t;
	typedef uint32_t AccountID_t;
	typedef int EUserUGCList;
	typedef int EUGCMatchingUGCType;
	typedef int EUserUGCListSortOrder;
	typedef int EUGCQuery;
	typedef int EItemStatistic;
	typedef int EItemUpdateStatus;
	typedef uint64_t PublishedFileId_t;
	typedef uint64_t UGCHandle_t;
	typedef struct {
		PublishedFileId_t m_nPublishedFileId;
		EResult m_eResult;												// The result of the operation.	
		EWorkshopFileType m_eFileType;									// Type of the file
		AppId_t m_nCreatorAppID;										// ID of the app that created this file.
		AppId_t m_nConsumerAppID;										// ID of the app that will consume this file.
		char m_rgchTitle[k_cchPublishedDocumentTitleMax];				// title of document
		char m_rgchDescription[k_cchPublishedDocumentDescriptionMax];	// description of document
		uint64_t m_ulSteamIDOwner;										// Steam ID of the user who created this content.
		uint32_t m_rtimeCreated;											// time when the published file was created
		uint32_t m_rtimeUpdated;											// time when the published file was last updated
		uint32_t m_rtimeAddedToUserList;									// time when the user added the published file to their list (not always applicable)
		ERemoteStoragePublishedFileVisibility m_eVisibility;			// visibility
		bool m_bBanned;													// whether the file was banned
		bool m_bAcceptedForUse;											// developer has specifically flagged this item as accepted in the Workshop
		bool m_bTagsTruncated;											// whether the list of tags was too long to be returned in the provided buffer
		char m_rgchTags[k_cchTagListMax];								// comma separated list of all tags associated with this file	
		// file/url information
		UGCHandle_t m_hFile;											// The handle of the primary file
		UGCHandle_t m_hPreviewFile;										// The handle of the preview file
		char m_pchFileName[k_cchFilenameMax];							// The cloud filename of the primary file
		int32_t m_nFileSize;												// Size of the primary file
		int32_t m_nPreviewFileSize;										// Size of the preview file
		char m_rgchURL[k_cchPublishedFileURLMax];						// URL (for a video or a website)
		// voting information
		uint32_t m_unVotesUp;												// number of votes up
		uint32_t m_unVotesDown;											// number of votes down
		float m_flScore;												// calculated score
		// collection details
		uint32_t m_unNumChildren;							
	} SteamUGCDetails_t;
	typedef struct {
		const char* * m_ppStrings;
		int32_t m_nNumStrings;
	} SteamParamStringArray_t ;

	typedef void ( *SteamAPIWarningMessageHook_t)(int, const char *);
	typedef void( *SteamAPI_PostAPIResultInProcess_t )(SteamAPICall_t callHandle, void *, unsigned int unCallbackSize, int iCallbackNum);
	typedef unsigned int ( *SteamAPI_CheckCallbackRegistered_t )( int iCallbackNum);

	extern bool SteamCAPI_Init();
	extern bool SteamCAPI_Shutdown();
	extern bool SteamCAPI_IsSteamRunning();
	extern void* CSteamUser();
	extern void* CSteamHTMLSurface();



//======================================================================
//==========================Steam App List API==========================
//======================================================================
void* CSteamAppList();

unsigned int SteamAppList_GetNumInstalledApps(void* appList);
unsigned int SteamAppList_GetInstalledApps(void* appList,  AppId_t *pvecAppID, unsigned int unMaxAppIDs);
int  SteamAppList_GetAppName(void* appList,  AppId_t nAppID, char* pchName, int cchNameMax);
int  SteamAppList_GetAppInstallDir(void* appList,  AppId_t nAppID, char* pchDirectory, int cchNameMax);
int SteamAppList_GetAppBuildId(void* appList,  AppId_t nAppID);



//======================================================================
//============================Steam App API=============================
//======================================================================
extern void* SteamCAPI_SteamApps();

extern bool SteamCAPI_ISteamApps_BIsSubscribed(void* apps);
extern bool SteamCAPI_ISteamApps_BIsLowViolence(void* apps);
extern bool SteamCAPI_ISteamApps_BIsCybercafe(void* apps);
extern bool SteamCAPI_ISteamApps_BIsVACBanned(void* apps);
extern const char* SteamCAPI_ISteamApps_GetCurrentGameLanguage(void* apps);
extern const char* SteamCAPI_ISteamApps_GetAvailableGameLanguages(void* apps);
extern bool SteamCAPI_ISteamApps_BIsSubscribedApp(void* apps,  AppId_t appID);
extern bool SteamCAPI_ISteamApps_BIsDlcInstalled(void* apps,  AppId_t appID);
extern unsigned int SteamCAPI_ISteamApps_GetEarliestPurchaseUnixTime(void* apps,  AppId_t nAppID);
extern bool SteamCAPI_ISteamApps_BIsSubscribedFromFreeWeekend(void* apps);
extern int SteamCAPI_ISteamApps_GetDLCCount(void* apps);
extern bool SteamCAPI_ISteamApps_BGetDLCDataByIndex(void* apps,  int iDLC, AppId_t *pAppID, bool *pbAvailable, char* pchName, int cchNameBufferSize);
extern void SteamCAPI_ISteamApps_InstallDLC(void* apps,  AppId_t nAppID);
extern void SteamCAPI_ISteamApps_UninstallDLC(void* apps,  AppId_t nAppID);
extern void SteamCAPI_ISteamApps_RequestAppProofOfPurchaseKey(void* apps,  AppId_t nAppID);
extern bool SteamCAPI_ISteamApps_GetCurrentBetaName(void* apps,  char* pchName, int cchNameBufferSize);
extern bool SteamCAPI_ISteamApps_MarkContentCorrupt(void* apps,  bool bMissingFilesOnly);
extern unsigned int SteamCAPI_ISteamApps_GetInstalledDepots(void* apps,  AppId_t appID, DepotId_t *pvecDepots, unsigned int cMaxDepots);
extern unsigned int SteamCAPI_ISteamApps_GetAppInstallDir(void* apps,  AppId_t appID, char* pchFolder, unsigned int cchFolderBufferSize);
extern bool SteamCAPI_ISteamApps_BIsAppInstalled(void* apps,  AppId_t appID);
extern CSteamID SteamCAPI_ISteamApps_GetAppOwner(void* apps);
extern const char* SteamCAPI_ISteamApps_GetLaunchQueryParam(void* apps,  const char* pchKey);
extern bool SteamCAPI_ISteamApps_GetDlcDownloadProgress(void* apps,  AppId_t nAppID, unsigned long long* punBytesDownloaded, unsigned long long* punBytesTotal); 
extern int SteamCAPI_ISteamApps_GetAppBuildId(void* apps);


//==============================================================================
//===============================Steam client API===============================
//==============================================================================
void* SteamCAPI_SteamClient();

HSteamPipe SteamCAPI_SteamClient_CreateSteamPipe(void* client);
bool SteamCAPI_SteamClient_BReleaseSteamPipe(void* client, HSteamPipe hSteamPipe);
HSteamUser SteamCAPI_SteamClient_ConnectToGlobalUser(void* client, HSteamPipe hSteamPipe);
HSteamUser SteamCAPI_SteamClient_CreateLocalUser(void* client, HSteamPipe *phSteamPipe, EAccountType eAccountType);
void SteamCAPI_SteamClient_ReleaseUser(void* client, HSteamPipe hSteamPipe, HSteamUser hUser);
void* SteamCAPI_SteamClient_GetISteamUser(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamGameServer(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void SteamCAPI_SteamClient_SetLocalIPBinding(void* client, unsigned int unIP, unsigned short int usPort);
void* SteamCAPI_SteamClient_GetISteamFriends(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamUtils(void* client, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamMatchmaking(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamMatchmakingServers(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamGenericInterface(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamUserStats(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamGameServerStats(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamApps(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamNetworking(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamRemoteStorage(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamScreenshots(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void SteamCAPI_SteamClient_RunFrame(void* client);
unsigned int SteamCAPI_SteamClient_GetIPCCallCount(void* client);
void SteamCAPI_SteamClient_SetWarningMessageHook(void* client, SteamAPIWarningMessageHook_t pFunction);
bool SteamCAPI_SteamClient_BShutdownIfAllPipesClosed(void* client);
void* SteamCAPI_SteamClient_GetISteamHTTP(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamUnifiedMessages(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamController(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamUGC(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamAppList(void* client, HSteamUser hSteamUser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamMusic(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamMusicRemote(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamHTMLSurface(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void SteamCAPI_SteamClient_Set_SteamAPI_CPostAPIResultInProcess(void* client, SteamAPI_PostAPIResultInProcess_t func);
void SteamCAPI_SteamClient_Remove_SteamAPI_CPostAPIResultInProcess(void* client, SteamAPI_PostAPIResultInProcess_t func);
void SteamCAPI_SteamClient_Set_SteamAPI_CCheckCallbackRegisteredInProcess(void* client, SteamAPI_CheckCallbackRegistered_t func);
void* SteamCAPI_SteamClient_GetISteamInventory(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);
void* SteamCAPI_SteamClient_GetISteamVideo(void* client, HSteamUser hSteamuser, HSteamPipe hSteamPipe, const char *pchVersion);




//==============================================================================
//=============================SteamController API==============================
//==============================================================================

typedef int ESteamControllerPad;
typedef struct {
	// If packet num matches that on your prior call, then the controller state hasn't been changed since 
	// your last call and there is no need to process it
	unsigned int unPacketNum;
	
	// bit flags for each of the buttons
	unsigned long int ulButtons;
	
	// Left pad coordinates
	short sLeftPadX;
	short sLeftPadY;
	
	// Right pad coordinates
	short sRightPadX;
	short sRightPadY;
	
} SteamControllerState_t;
extern void* CSteamController();

bool 								SteamCAPI_SteamController_Init(void* controller);
bool 								SteamCAPI_SteamController_Shutdown(void* controller);
void 								SteamCAPI_SteamController_RunFrame(void* controller);
int 								SteamCAPI_SteamController_GetConnectedControllers(void* controller, ControllerHandle_t *handlesOut);
bool 								SteamCAPI_SteamController_ShowBindingPanel(void* controller, ControllerHandle_t controllerHandle);
ControllerActionSetHandle_t 		SteamCAPI_SteamController_GetActionSetHandle(void* controller, const char *pszActionSetName);
void 								SteamCAPI_SteamController_ActivateActionSet(void* controller, ControllerHandle_t controllerHandle, ControllerActionSetHandle_t actionSetHandle);
ControllerActionSetHandle_t 		SteamCAPI_SteamController_GetCurrentActionSet(void* controller, ControllerHandle_t controllerHandle);
ControllerDigitalActionHandle_t 	SteamCAPI_SteamController_GetDigitalActionHandle(void* controller, const char *pszActionName);
ControllerDigitalActionData_t 		SteamCAPI_SteamController_GetDigitalActionData(void* controller, ControllerHandle_t controllerHandle, ControllerDigitalActionHandle_t digitalActionHandle);
int 								SteamCAPI_SteamController_GetDigitalActionOrigins(void* controller, ControllerHandle_t controllerHandle, ControllerActionSetHandle_t actionSetHandle, ControllerDigitalActionHandle_t digitalActionHandle, EControllerActionOrigin *originsOut);
ControllerAnalogActionHandle_t 		SteamCAPI_SteamController_GetAnalogActionHandle(void* controller, const char *pszActionName);
ControllerAnalogActionData_t 		SteamCAPI_SteamController_GetAnalogActionData(void* controller, ControllerHandle_t controllerHandle, ControllerAnalogActionHandle_t analogActionHandle);
int 								SteamCAPI_SteamController_GetAnalogActionOrigins(void* controller, ControllerHandle_t controllerHandle, ControllerActionSetHandle_t actionSetHandle, ControllerAnalogActionHandle_t analogActionHandle, EControllerActionOrigin *originsOut);
void 								SteamCAPI_SteamController_StopAnalogActionMomentum(void* controller, ControllerHandle_t controllerHandle, ControllerAnalogActionHandle_t eAction);
void 								SteamCAPI_SteamController_TriggerHapticPulse(void* controller, ControllerHandle_t controllerHandle, ESteamControllerPad eTargetPad, unsigned short usDurationMicroSec);



//==============================================================================
//==============================Steam friends API===============================
//==============================================================================

extern void* CSteamFriends();
extern const char* SteamFriends_GetPersonaName(void* steamFriends);
extern SteamAPICall_t SteamFriends_SetPersonaName(void* steamFriends, char* pchPersonaName);

extern EPersonaState SteamFriends_GetPersonaState(void* steamFriends);
extern int SteamFriends_GetFriendCount(void* steamFriends, int iFriendFlags);
extern CSteamID SteamFriends_GetFriendByIndex(void* steamFriends, int iFriend, int iFriendFlags);
extern EFriendRelationship SteamFriends_GetFriendRelationship(void* steamFriends, CSteamID steamIDFriend);
extern EPersonaState SteamFriends_GetFriendPersonaState(void* steamFriends, CSteamID steamIDFriend);
extern const char* SteamFriends_GetFriendPersonaName(void* steamFriends, CSteamID steamIDFriend);
extern bool SteamFriends_GetFriendGamePlayed(void* steamFriends, CSteamID steamIDFriend, FriendGameInfo_t *pFriendGameInfo);
extern const char* SteamFriends_GetFriendPersonaNameHistory(void* steamFriends, CSteamID steamIDFriend, int iPersonaName);
extern int SteamFriends_GetFriendSteamLevel(void* steamFriends, CSteamID steamIDFriend);
extern const char* SteamFriends_GetPlayerNickname(void* steamFriends, CSteamID steamIDPlayer);
extern int SteamFriends_GetFriendsGroupCount(void* steamFriends);
extern FriendsGroupID_t SteamFriends_GetFriendsGroupIDByIndex(void* steamFriends, int iFG);
extern const char* SteamFriends_GetFriendsGroupName(void* steamFriends, FriendsGroupID_t friendsGroupID);
extern int SteamFriends_GetFriendsGroupMembersCount(void* steamFriends, FriendsGroupID_t friendsGroupID);
extern void SteamFriends_GetFriendsGroupMembersList(void* steamFriends, FriendsGroupID_t friendsGroupID, CSteamID *pOutSteamIDMembers, int nMembersCount);
extern bool SteamFriends_HasFriend(void* steamFriends, CSteamID steamIDFriend, int iFriendFlags);
extern int SteamFriends_GetClanCount(void* steamFriends);
extern CSteamID SteamFriends_GetClanByIndex(void* steamFriends, int iClan);
extern const char* SteamFriends_GetClanName(void* steamFriends, CSteamID steamIDClan);
extern const char* SteamFriends_GetClanTag(void* steamFriends, CSteamID steamIDClan);
extern bool SteamFriends_GetClanActivityCounts(void* steamFriends, CSteamID steamIDClan, int *pnOnline, int *pnInGame, int *pnChatting);
extern SteamAPICall_t SteamFriends_DownloadClanActivityCounts(void* steamFriends, CSteamID *psteamIDClans, int cClansToRequest);
extern int SteamFriends_GetFriendCountFromSource(void* steamFriends, CSteamID steamIDSource);
extern CSteamID SteamFriends_GetFriendFromSourceByIndex(void* steamFriends, CSteamID steamIDSource, int iFriend);
extern bool SteamFriends_IsUserInSource(void* steamFriends, CSteamID steamIDUser, CSteamID steamIDSource);
extern void SteamFriends_SetInGameVoiceSpeaking(void* steamFriends, CSteamID steamIDUser, bool bSpeaking);
extern void SteamFriends_ActivateGameOverlay(void* steamFriends, const char* pchDialog);
extern void SteamFriends_ActivateGameOverlayToUser(void* steamFriends, const char* pchDialog, CSteamID steamID);
extern void SteamFriends_ActivateGameOverlayToWebPage(void* steamFriends, const char* pchURL);
extern void SteamFriends_ActivateGameOverlayToStore(void* steamFriends, AppId_t nAppID, EOverlayToStoreFlag eFlag);
extern void SteamFriends_SetPlayedWith(void* steamFriends, CSteamID steamIDUserPlayedWith);
extern void SteamFriends_ActivateGameOverlayInviteDialog(void* steamFriends, CSteamID steamIDLobby);
extern int SteamFriends_GetSmallFriendAvatar(void* steamFriends, CSteamID steamIDFriend);
extern int SteamFriends_GetMediumFriendAvatar(void* steamFriends, CSteamID steamIDFriend);
extern int SteamFriends_GetLargeFriendAvatar(void* steamFriends, CSteamID steamIDFriend);
extern bool SteamFriends_RequestUserInformation(void* steamFriends, CSteamID steamIDUser, bool bRequireNameOnly);
extern SteamAPICall_t SteamFriends_RequestClanOfficerList(void* steamFriends, CSteamID steamIDClan);
extern CSteamID SteamFriends_GetClanOwner(void* steamFriends, CSteamID steamIDClan);
extern int SteamFriends_GetClanOfficerCount(void* steamFriends, CSteamID steamIDClan);
extern CSteamID SteamFriends_GetClanOfficerByIndex(void* steamFriends, CSteamID steamIDClan, int iOfficer);
extern unsigned int SteamFriends_GetUserRestrictions(void* steamFriends);
extern bool SteamFriends_SetRichPresence(void* steamFriends, const char* pchKey, const char* pchValue);
extern void SteamFriends_ClearRichPresence(void* steamFriends);
extern const char* SteamFriends_GetFriendRichPresence(void* steamFriends, CSteamID steamIDFriend, const char* pchKey);
extern int SteamFriends_GetFriendRichPresenceKeyCount(void* steamFriends, CSteamID steamIDFriend);
extern const char* SteamFriends_GetFriendRichPresenceKeyByIndex(void* steamFriends, CSteamID steamIDFriend, int iKey);
extern void SteamFriends_RequestFriendRichPresence(void* steamFriends, CSteamID steamIDFriend);
extern bool SteamFriends_InviteUserToGame(void* steamFriends, CSteamID steamIDFriend, const char* pchConnectString);
extern int SteamFriends_GetCoplayFriendCount(void* steamFriends);
extern CSteamID SteamFriends_GetCoplayFriend(void* steamFriends, int iCoplayFriend);
extern int SteamFriends_GetFriendCoplayTime(void* steamFriends, CSteamID steamIDFriend);
extern AppId_t SteamFriends_GetFriendCoplayGame(void* steamFriends, CSteamID steamIDFriend);
extern SteamAPICall_t SteamFriends_JoinClanChatRoom(void* steamFriends, CSteamID steamIDClan);
extern bool SteamFriends_LeaveClanChatRoom(void* steamFriends, CSteamID steamIDClan);
extern int SteamFriends_GetClanChatMemberCount(void* steamFriends, CSteamID steamIDClan);
extern CSteamID SteamFriends_GetChatMemberByIndex(void* steamFriends, CSteamID steamIDClan, int iUser);
extern bool SteamFriends_SendClanChatMessage(void* steamFriends, CSteamID steamIDClanChat, const char* pchText);
extern int SteamFriends_GetClanChatMessage(void* steamFriends, CSteamID steamIDClanChat, int iMessage, void *prgchText, int cchTextMax, EChatEntryType *peChatEntryType, CSteamID *psteamidChatter);
extern bool SteamFriends_IsClanChatAdmin(void* steamFriends, CSteamID steamIDClanChat, CSteamID steamIDUser);
extern bool SteamFriends_IsClanChatWindowOpenInSteam(void* steamFriends, CSteamID steamIDClanChat);
extern bool SteamFriends_OpenClanChatWindowInSteam(void* steamFriends, CSteamID steamIDClanChat);
extern bool SteamFriends_CloseClanChatWindowInSteam(void* steamFriends, CSteamID steamIDClanChat);
extern bool SteamFriends_SetListenForFriendsMessages(void* steamFriends, bool bInterceptEnabled);
extern bool SteamFriends_ReplyToFriendMessage(void* steamFriends, CSteamID steamIDFriend, const char* pchMsgToSend);
extern int SteamFriends_GetFriendMessage(void* steamFriends, CSteamID steamIDFriend, int iMessageID, void *pvData, int cubData, EChatEntryType *peChatEntryType);
extern SteamAPICall_t SteamFriends_GetFollowerCount(void* steamFriends, CSteamID steamID);
extern SteamAPICall_t SteamFriends_IsFollowing(void* steamFriends, CSteamID steamID);
extern SteamAPICall_t SteamFriends_EnumerateFollowingList(void* steamFriends, unsigned int unStartIndex);




//======================================================================
//======================Steam game coordinator API======================
//======================================================================




EGCResults SteamCAPI_ISteamGameCoordinator_SendMessage(void* gc, unsigned int unMsgType, const void *pubData, unsigned int cubData);
bool SteamCAPI_ISteamGameCoordinator_IsMessageAvailable(void* gc, unsigned int *pcubMsgSize);
EGCResults SteamCAPI_ISteamGameCoordinator_RetrieveMessage(void* gc, unsigned int *punMsgType, void *pubDest, unsigned int cubDest, unsigned int *pcubMsgSize);

//======================================================================
//========================Steam game server API=========================
//======================================================================

bool SteamCAPI_ISteamGameServer_InitGameServer(void* server, unsigned int unIP, unsigned short int usGamePort, unsigned short int usQueryPort, unsigned int unFlags, AppId_t nGameAppId, const char *pchVersionString);
void SteamCAPI_ISteamGameServer_SetProduct(void* server, const char *pszProduct);
void SteamCAPI_ISteamGameServer_SetGameDescription(void* server, const char *pszGameDescription);
void SteamCAPI_ISteamGameServer_SetModDir(void* server, const char *pszModDir);
void SteamCAPI_ISteamGameServer_SetDedicatedServer(void* server, bool bDedicated);
void SteamCAPI_ISteamGameServer_LogOn(void* server, const char *pszToken);
void SteamCAPI_ISteamGameServer_LogOnAnonymous(void* server);
void SteamCAPI_ISteamGameServer_LogOff(void* server);
bool SteamCAPI_ISteamGameServer_BLoggedOn(void* server);
bool SteamCAPI_ISteamGameServer_BSecure(void* server) ;
CSteamID SteamCAPI_ISteamGameServer_GetSteamID(void* server);
bool SteamCAPI_ISteamGameServer_WasRestartRequested(void* server);
void SteamCAPI_ISteamGameServer_SetMaxPlayerCount(void* server, int cPlayersMax);
void SteamCAPI_ISteamGameServer_SetBotPlayerCount(void* server, int cBotplayers);
void SteamCAPI_ISteamGameServer_SetServerName(void* server, const char *pszServerName);
void SteamCAPI_ISteamGameServer_SetMapName(void* server, const char *pszMapName);
void SteamCAPI_ISteamGameServer_SetPasswordProtected(void* server, bool bPasswordProtected);
void SteamCAPI_ISteamGameServer_SetSpectatorPort(void* server, unsigned short int unSpectatorPort);
void SteamCAPI_ISteamGameServer_SetSpectatorServerName(void* server, const char *pszSpectatorServerName);
void SteamCAPI_ISteamGameServer_ClearAllKeyValues(void* server);
void SteamCAPI_ISteamGameServer_SetKeyValue(void* server, const char *pKey, const char *pValue);
void SteamCAPI_ISteamGameServer_SetGameTags(void* server, const char *pchGameTags);
void SteamCAPI_ISteamGameServer_SetGameData(void* server, const char *pchGameData);
void SteamCAPI_ISteamGameServer_SetRegion(void* server, const char *pszRegion);
bool SteamCAPI_ISteamGameServer_SendUserConnectAndAuthenticate(void* server, unsigned int unIPClient, const void *pvAuthBlob, unsigned int cubAuthBlobSize, CSteamID *pSteamIDUser);
CSteamID SteamCAPI_ISteamGameServer_CreateUnauthenticatedUserConnection(void* server);
void SteamCAPI_ISteamGameServer_SendUserDisconnect(void* server, CSteamID steamIDUser);
bool SteamCAPI_ISteamGameServer_BUpdateUserData(void* server, CSteamID steamIDUser, const char *pchPlayerName, unsigned int uScore);
HAuthTicket SteamCAPI_ISteamGameServer_GetAuthSessionTicket(void* server, void *pTicket, int cbMaxTicket, unsigned int *pcbTicket);
EBeginAuthSessionResult SteamCAPI_ISteamGameServer_BeginAuthSession(void* server, const void *pAuthTicket, int cbAuthTicket, CSteamID steamID);
void SteamCAPI_ISteamGameServer_EndAuthSession(void* server, CSteamID steamID);
void SteamCAPI_ISteamGameServer_CancelAuthTicket(void* server, HAuthTicket hAuthTicket);
EUserHasLicenseForAppResult SteamCAPI_ISteamGameServer_UserHasLicenseForApp(void* server, CSteamID steamID, AppId_t appID);
bool SteamCAPI_ISteamGameServer_RequestUserGroupStatus(void* server, CSteamID steamIDUser, CSteamID steamIDGroup);
void SteamCAPI_ISteamGameServer_GetGameplayStats(void* server);
SteamAPICall_t SteamCAPI_ISteamGameServer_GetServerReputation(void* server);
unsigned int SteamCAPI_ISteamGameServer_GetPublicIP(void* server);
bool SteamCAPI_ISteamGameServer_HandleIncomingPacket(void* server, const void *pData, int cbData, unsigned int srcIP, unsigned short int srcPort);
int SteamCAPI_ISteamGameServer_GetNextOutgoingPacket(void* server, void *pOut, int cbMaxOut, unsigned int *pNetAdr, unsigned short int *pPort);
void SteamCAPI_ISteamGameServer_EnableHeartbeats(void* server, bool bActive);
void SteamCAPI_ISteamGameServer_SetHeartbeatInterval(void* server, int iHeartbeatInterval);
void SteamCAPI_ISteamGameServer_ForceHeartbeat(void* server);
SteamAPICall_t SteamCAPI_ISteamGameServer_AssociateWithClan(void* server, CSteamID steamIDClan);
SteamAPICall_t SteamCAPI_ISteamGameServer_ComputeNewPlayerCompatibility(void* server, CSteamID steamIDNewPlayer);




//======================================================================
//=====================Steam game server stats API======================
//======================================================================
void* CSteamGameServerStats();

SteamAPICall_t ISteamGameServerStats_RequestUserStats(void* stats, CSteamID steamIDUser);
bool ISteamGameServerStats_GetUserStati(void* stats, CSteamID steamIDUser, const char* pchName, int *pData);
bool ISteamGameServerStats_GetUserStatf(void* stats, CSteamID steamIDUser, const char* pchName, float *pData);
bool ISteamGameServerStats_GetUserAchievement(void* stats, CSteamID steamIDUser, const char* pchName, bool *pbAchieved);
bool ISteamGameServerStats_SetUserStati(void* stats, CSteamID steamIDUser, const char* pchName, int nData);
bool ISteamGameServerStats_SetUserStatf(void* stats, CSteamID steamIDUser, const char* pchName, float fData);
bool ISteamGameServerStats_UpdateUserAvgRateStat(void* stats, CSteamID steamIDUser, const char* pchName, float flCountThisSession, double dSessionLength);
bool ISteamGameServerStats_SetUserAchievement(void* stats, CSteamID steamIDUser, const char* pchName);
bool ISteamGameServerStats_ClearUserAchievement(void* stats, CSteamID steamIDUser, const char* pchName);
SteamAPICall_t ISteamGameServerStats_StoreUserStats(void* stats, CSteamID steamIDUser);

//======================================================================
//========================Steam html surface API========================
//======================================================================

bool			SteamCAPI_ISteamHTMLSurface_Init(void* surface);
bool			SteamCAPI_ISteamHTMLSurface_Shutdown(void* surface);
SteamAPICall_t	SteamCAPI_ISteamHTMLSurface_CreateBrowser(void* surface, const char *pchUserAgent, const char *pchUserCSS);
void			SteamCAPI_ISteamHTMLSurface_RemoveBrowser(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_LoadURL(void* surface, HHTMLBrowser unBrowserHandle, const char *pchURL, const char *pchPostData);
void			SteamCAPI_ISteamHTMLSurface_SetSize(void* surface, HHTMLBrowser unBrowserHandle, unsigned int unWidth, unsigned int unHeight);
void			SteamCAPI_ISteamHTMLSurface_StopLoad(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_Reload(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_GoBack(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_GoForward(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_AddHeader(void* surface, HHTMLBrowser unBrowserHandle, const char *pchKey, const char *pchValue);
void			SteamCAPI_ISteamHTMLSurface_ExecuteJavascript(void* surface, HHTMLBrowser unBrowserHandle, const char *pchScript);
void			SteamCAPI_ISteamHTMLSurface_MouseUp(void* surface, HHTMLBrowser unBrowserHandle, EHTMLMouseButton eMouseButton);
void			SteamCAPI_ISteamHTMLSurface_MouseDown(void* surface, HHTMLBrowser unBrowserHandle, EHTMLMouseButton eMouseButton);
void			SteamCAPI_ISteamHTMLSurface_MouseDoubleClick(void* surface, HHTMLBrowser unBrowserHandle, EHTMLMouseButton eMouseButton);
void			SteamCAPI_ISteamHTMLSurface_MouseMove(void* surface, HHTMLBrowser unBrowserHandle, int x, int y);
void			SteamCAPI_ISteamHTMLSurface_MouseWheel(void* surface, HHTMLBrowser unBrowserHandle, int nDelta);
void			SteamCAPI_ISteamHTMLSurface_KeyDown(void* surface, HHTMLBrowser unBrowserHandle, unsigned int nNativeKeyCode, EHTMLKeyModifiers eHTMLKeyModifiers);
void			SteamCAPI_ISteamHTMLSurface_KeyUp(void* surface, HHTMLBrowser unBrowserHandle, unsigned int nNativeKeyCode, EHTMLKeyModifiers eHTMLKeyModifiers);
void			SteamCAPI_ISteamHTMLSurface_KeyChar(void* surface, HHTMLBrowser unBrowserHandle, unsigned int cUnicodeChar, EHTMLKeyModifiers eHTMLKeyModifiers);
void			SteamCAPI_ISteamHTMLSurface_SetHorizontalScroll(void* surface, HHTMLBrowser unBrowserHandle, unsigned int nAbsolutePixelScroll);
void			SteamCAPI_ISteamHTMLSurface_SetVerticalScroll(void* surface, HHTMLBrowser unBrowserHandle, unsigned int nAbsolutePixelScroll);
void			SteamCAPI_ISteamHTMLSurface_SetKeyFocus(void* surface, HHTMLBrowser unBrowserHandle, bool bHasKeyFocus);
void			SteamCAPI_ISteamHTMLSurface_ViewSource(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_CopyToClipboard(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_PasteFromClipboard(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_Find(void* surface, HHTMLBrowser unBrowserHandle, const char *pchSearchStr, bool bCurrentlyInFind, bool bReverse);
void			SteamCAPI_ISteamHTMLSurface_StopFind(void* surface, HHTMLBrowser unBrowserHandle);
void			SteamCAPI_ISteamHTMLSurface_GetLinkAtPosition(void* surface,  HHTMLBrowser unBrowserHandle, int x, int y);
void			SteamCAPI_ISteamHTMLSurface_SetCookie(void* surface, const char *pchHostname, const char *pchKey, const char *pchValue, const char *pchPath, RTime32 nExpires, bool bSecure, bool bHTTPOnly);
void			SteamCAPI_ISteamHTMLSurface_SetPageScaleFactor(void* surface, HHTMLBrowser unBrowserHandle, float flZoom, int nPointX, int nPointY);
void			SteamCAPI_ISteamHTMLSurface_SetBackgroundMode(void* surface, HHTMLBrowser unBrowserHandle, bool bBackgroundMode);
void			SteamCAPI_ISteamHTMLSurface_AllowStartRequest(void* surface, HHTMLBrowser unBrowserHandle, bool bAllowed);
void			SteamCAPI_ISteamHTMLSurface_JSDialogResponse(void* surface, HHTMLBrowser unBrowserHandle, bool bResult);
void			SteamCAPI_ISteamHTMLSurface_FileLoadDialogResponse(void* surface, HHTMLBrowser unBrowserHandle, const char **pchSelectedFiles);

//==============================================================================
//================================Steam HTTP API================================
//==============================================================================

HTTPRequestHandle 			SteamCAPI_ISteamHTTP_CreateHTTPRequest(void* http, EHTTPMethod eHTTPRequestMethod, const char *pchAbsoluteURL);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestContextValue(void* http, HTTPRequestHandle hRequest, unsigned long long ulContextValue);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestNetworkActivityTimeout(void* http, HTTPRequestHandle hRequest, unsigned int unTimeoutSeconds);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestHeaderValue(void* http, HTTPRequestHandle hRequest, const char *pchHeaderName, const char *pchHeaderValue);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestGetOrPostParameter(void* http, HTTPRequestHandle hRequest, const char *pchParamName, const char *pchParamValue);
bool 						SteamCAPI_ISteamHTTP_SendHTTPRequest(void* http, HTTPRequestHandle hRequest, SteamAPICall_t *pCallHandle);
bool 						SteamCAPI_ISteamHTTP_SendHTTPRequestAndStreamResponse(void* http, HTTPRequestHandle hRequest, SteamAPICall_t *pCallHandle);
bool 						SteamCAPI_ISteamHTTP_DeferHTTPRequest(void* http, HTTPRequestHandle hRequest);
bool 						SteamCAPI_ISteamHTTP_PrioritizeHTTPRequest(void* http, HTTPRequestHandle hRequest);
bool 						SteamCAPI_ISteamHTTP_GetHTTPResponseHeaderSize(void* http, HTTPRequestHandle hRequest, const char *pchHeaderName, unsigned int *unResponseHeaderSize);
bool 						SteamCAPI_ISteamHTTP_GetHTTPResponseHeaderValue(void* http, HTTPRequestHandle hRequest, const char *pchHeaderName, unsigned char *pHeaderValueBuffer, unsigned int unBufferSize);
bool 						SteamCAPI_ISteamHTTP_GetHTTPResponseBodySize(void* http, HTTPRequestHandle hRequest, unsigned int *unBodySize);
bool 						SteamCAPI_ISteamHTTP_GetHTTPResponseBodyData(void* http, HTTPRequestHandle hRequest, unsigned char *pBodyDataBuffer, unsigned int unBufferSize);
bool 						SteamCAPI_ISteamHTTP_GetHTTPStreamingResponseBodyData(void* http, HTTPRequestHandle hRequest, unsigned int cOffset, unsigned char *pBodyDataBuffer, unsigned int unBufferSize);
bool 						SteamCAPI_ISteamHTTP_ReleaseHTTPRequest(void* http, HTTPRequestHandle hRequest);
bool 						SteamCAPI_ISteamHTTP_GetHTTPDownloadProgressPct(void* http, HTTPRequestHandle hRequest, float *pflPercentOut);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestRawPostBody(void* http, HTTPRequestHandle hRequest, const char *pchContentType, unsigned char *pubBody, unsigned int unBodyLen);
HTTPCookieContainerHandle 	SteamCAPI_ISteamHTTP_CreateCookieContainer(void* http, bool bAllowResponsesToModify);
bool 						SteamCAPI_ISteamHTTP_ReleaseCookieContainer(void* http, HTTPCookieContainerHandle hCookieContainer);
bool 						SteamCAPI_ISteamHTTP_SetCookie(void* http, HTTPCookieContainerHandle hCookieContainer, const char *pchHost, const char *pchUrl, const char *pchCookie);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestCookieContainer(void* http, HTTPRequestHandle hRequest, HTTPCookieContainerHandle hCookieContainer);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestUserAgentInfo(void* http, HTTPRequestHandle hRequest, const char *pchUserAgentInfo);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestRequiresVerifiedCertificate(void* http, HTTPRequestHandle hRequest, bool bRequireVerifiedCertificate);
bool 						SteamCAPI_ISteamHTTP_SetHTTPRequestAbsoluteTimeoutMS(void* http, HTTPRequestHandle hRequest, unsigned int unMilliseconds);
bool 						SteamCAPI_ISteamHTTP_GetHTTPRequestWasTimedOut(void* http, HTTPRequestHandle hRequest, bool *pbWasTimedOut);

//==============================================================================
//=============================Steam inventory API==============================
//==============================================================================

EResult	SteamCAPI_ISteamInventory_GetResultStatus(void* inventory,  SteamInventoryResult_t resultHandle);																																																									
bool 	SteamCAPI_ISteamInventory_GetResultItems(void* inventory,  SteamInventoryResult_t resultHandle,  SteamItemDetails_t *pOutItemsArray, unsigned int *punOutItemsArraySize);																																									
unsigned int 	SteamCAPI_ISteamInventory_GetResultTimestamp(void* inventory,  SteamInventoryResult_t resultHandle);																																																									
bool 	SteamCAPI_ISteamInventory_CheckResultSteamID(void* inventory,  SteamInventoryResult_t resultHandle, CSteamID steamIDExpected);																																																		
void 	SteamCAPI_ISteamInventory_DestroyResult(void* inventory,  SteamInventoryResult_t resultHandle);																																																										
bool 	SteamCAPI_ISteamInventory_GetAllItems(void* inventory,  SteamInventoryResult_t *pResultHandle);																																																										
bool 	SteamCAPI_ISteamInventory_GetItemsByID(void* inventory,  SteamInventoryResult_t *pResultHandle, const SteamItemInstanceID_t *pInstanceIDs, unsigned int unCountInstanceIDs);																																								
bool 	SteamCAPI_ISteamInventory_SerializeResult(void* inventory,  SteamInventoryResult_t resultHandle,  void *pOutBuffer, unsigned int *punOutBufferSize);																																														
bool 	SteamCAPI_ISteamInventory_DeserializeResult(void* inventory,  SteamInventoryResult_t *pOutResultHandle, const void *pBuffer, unsigned int unBufferSize, bool bRESERVED_MUST_BE_FALSE);																																						
bool 	SteamCAPI_ISteamInventory_GenerateItems(void* inventory,  SteamInventoryResult_t *pResultHandle, const SteamItemDef_t *pArrayItemDefs, const unsigned int *punArrayQuantity, unsigned int unArrayLength);																																		
bool 	SteamCAPI_ISteamInventory_GrantPromoItems(void* inventory,  SteamInventoryResult_t *pResultHandle);																																																									
bool 	SteamCAPI_ISteamInventory_AddPromoItem(void* inventory,  SteamInventoryResult_t *pResultHandle, SteamItemDef_t itemDef);																																																				
bool 	SteamCAPI_ISteamInventory_AddPromoItems(void* inventory,  SteamInventoryResult_t *pResultHandle, const SteamItemDef_t *pArrayItemDefs, unsigned int unArrayLength);																																										
bool 	SteamCAPI_ISteamInventory_ConsumeItem(void* inventory,  SteamInventoryResult_t *pResultHandle, SteamItemInstanceID_t itemConsume, unsigned int unQuantity);																																												
bool 	SteamCAPI_ISteamInventory_ExchangeItems(void* inventory,  SteamInventoryResult_t *pResultHandle, const SteamItemDef_t *pArrayGenerate, const unsigned int *punArrayGenerateQuantity, unsigned int unArrayGenerateLength, const SteamItemInstanceID_t *pArrayDestroy, const unsigned int *punArrayDestroyQuantity, unsigned int unArrayDestroyLength);		
bool 	SteamCAPI_ISteamInventory_TransferItemQuantity(void* inventory,  SteamInventoryResult_t *pResultHandle, SteamItemInstanceID_t itemIdSource, unsigned int unQuantity, SteamItemInstanceID_t itemIdDest);																																	
void 	SteamCAPI_ISteamInventory_SendItemDropHeartbeat(void* inventory);																																																																	
bool 	SteamCAPI_ISteamInventory_TriggerItemDrop(void* inventory,  SteamInventoryResult_t *pResultHandle, SteamItemDef_t dropListDefinition);																																																
bool 	SteamCAPI_ISteamInventory_TradeItems(void* inventory,  SteamInventoryResult_t *pResultHandle, CSteamID steamIDTradePartner, const SteamItemInstanceID_t *pArrayGive, const unsigned int *pArrayGiveQuantity, unsigned int nArrayGiveLength, const SteamItemInstanceID_t *pArrayGet, const unsigned int *pArrayGetQuantity, unsigned int nArrayGetLength);	
bool 	SteamCAPI_ISteamInventory_LoadItemDefinitions(void* inventory);																																																																		
bool 	SteamCAPI_ISteamInventory_GetItemDefinitionIDs(void* inventory, SteamItemDef_t *pItemDefIDs, unsigned int *punItemDefIDsArraySize);																																																		
bool 	SteamCAPI_ISteamInventory_GetItemDefinitionProperty(void* inventory,  SteamItemDef_t iDefinition, const char *pchPropertyName, char *pchValueBuffer, unsigned int *punValueBufferSize);																																					

//==============================================================================
//==============================Steam Matchmaking===============================
//==============================================================================

int 			SteamCAPI_ISteamMatchmaking_GetFavoriteGameCount(void* mm);
bool 			SteamCAPI_ISteamMatchmaking_GetFavoriteGame(void* mm, int iGame, AppId_t *pnAppID, unsigned int *pnIP, unsigned short int *pnConnPort, unsigned short int *pnQueryPort, unsigned int *punFlags, unsigned int *pRTime32LastPlayedOnServer );
int 			SteamCAPI_ISteamMatchmaking_AddFavoriteGame(void* mm, AppId_t nAppID, unsigned int nIP, unsigned short int nConnPort, unsigned short int nQueryPort, unsigned int unFlags, unsigned int rTime32LastPlayedOnServer );
bool 			SteamCAPI_ISteamMatchmaking_RemoveFavoriteGame(void* mm, AppId_t nAppID, unsigned int nIP, unsigned short int nConnPort, unsigned short int nQueryPort, unsigned int unFlags );
SteamAPICall_t 	SteamCAPI_ISteamMatchmaking_RequestLobbyList(void* mm);
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListStringFilter(void* mm, const char *pchKeyToMatch, const char *pchValueToMatch, ELobbyComparison eComparisonType );
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListNumericalFilter(void* mm, const char *pchKeyToMatch, int nValueToMatch, ELobbyComparison eComparisonType );
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListNearValueFilter(void* mm, const char *pchKeyToMatch, int nValueToBeCloseTo );
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListFilterSlotsAvailable(void* mm, int nSlotsAvailable );
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListDistanceFilter(void* mm, ELobbyDistanceFilter eLobbyDistanceFilter );
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListResultCountFilter(void* mm, int cMaxResults );
void 			SteamCAPI_ISteamMatchmaking_AddRequestLobbyListCompatibleMembersFilter(void* mm, CSteamID steamIDLobby );
CSteamID 		SteamCAPI_ISteamMatchmaking_GetLobbyByIndex(void* mm, int iLobby );
SteamAPICall_t 	SteamCAPI_ISteamMatchmaking_CreateLobby(void* mm, ELobbyType eLobbyType, int cMaxMembers );
SteamAPICall_t 	SteamCAPI_ISteamMatchmaking_JoinLobby(void* mm, CSteamID steamIDLobby );
void 			SteamCAPI_ISteamMatchmaking_LeaveLobby(void* mm, CSteamID steamIDLobby );
bool 			SteamCAPI_ISteamMatchmaking_InviteUserToLobby(void* mm, CSteamID steamIDLobby, CSteamID steamIDInvitee );
int 			SteamCAPI_ISteamMatchmaking_GetNumLobbyMembers(void* mm, CSteamID steamIDLobby );
CSteamID 		SteamCAPI_ISteamMatchmaking_GetLobbyMemberByIndex(void* mm, CSteamID steamIDLobby, int iMember );
const char*		SteamCAPI_ISteamMatchmaking_GetLobbyData(void* mm, CSteamID steamIDLobby, const char *pchKey );
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyData(void* mm, CSteamID steamIDLobby, const char *pchKey, const char *pchValue );
int 			SteamCAPI_ISteamMatchmaking_GetLobbyDataCount(void* mm, CSteamID steamIDLobby );
bool 			SteamCAPI_ISteamMatchmaking_GetLobbyDataByIndex(void* mm, CSteamID steamIDLobby, int iLobbyData, char *pchKey, int cchKeyBufferSize, char *pchValue, int cchValueBufferSize );
bool 			SteamCAPI_ISteamMatchmaking_DeleteLobbyData(void* mm, CSteamID steamIDLobby, const char *pchKey );
const char*		SteamCAPI_ISteamMatchmaking_GetLobbyMemberData(void* mm, CSteamID steamIDLobby, CSteamID steamIDUser, const char *pchKey );
void 			SteamCAPI_ISteamMatchmaking_SetLobbyMemberData(void* mm, CSteamID steamIDLobby, const char *pchKey, const char *pchValue );
bool 			SteamCAPI_ISteamMatchmaking_SendLobbyChatMsg(void* mm, CSteamID steamIDLobby, const void *pvMsgBody, int cubMsgBody );
int 			SteamCAPI_ISteamMatchmaking_GetLobbyChatEntry(void* mm, CSteamID steamIDLobby, int iChatID, CSteamID *pSteamIDUser, void *pvData, int cubData, EChatEntryType *peChatEntryType );
bool 			SteamCAPI_ISteamMatchmaking_RequestLobbyData(void* mm, CSteamID steamIDLobby );
void 			SteamCAPI_ISteamMatchmaking_SetLobbyGameServer(void* mm, CSteamID steamIDLobby, unsigned int unGameServerIP, unsigned short int unGameServerPort, CSteamID steamIDGameServer );
bool 			SteamCAPI_ISteamMatchmaking_GetLobbyGameServer(void* mm, CSteamID steamIDLobby, unsigned int *punGameServerIP, unsigned short int *punGameServerPort, CSteamID *psteamIDGameServer );
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyMemberLimit(void* mm, CSteamID steamIDLobby, int cMaxMembers );
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyType(void* mm, CSteamID steamIDLobby, ELobbyType eLobbyType );
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyJoinable(void* mm, CSteamID steamIDLobby, bool bLobbyJoinable );
CSteamID 		SteamCAPI_ISteamMatchmaking_GetLobbyOwner(void* mm, CSteamID steamIDLobby );
bool 			SteamCAPI_ISteamMatchmaking_SetLobbyOwner(void* mm, CSteamID steamIDLobby, CSteamID steamIDNewOwner );
bool 			SteamCAPI_ISteamMatchmaking_SetLinkedLobby(void* mm, CSteamID steamIDLobby, CSteamID steamIDLobbyDependent );

//==============================================================================
//==================Steam matchmaking server list response API==================
//==============================================================================

void SteamCAPI_ISteamMatchmakingServerListResponse_ServerResponded( HServerListRequest hRequest, int iServer );
void SteamCAPI_ISteamMatchmakingServerListResponse_ServerFailedToRespond( HServerListRequest hRequest, int iServer );
void SteamCAPI_ISteamMatchmakingServerListResponse_RefreshComplete( HServerListRequest hRequest, EMatchMakingServerResponse response );








//==============================================================================
//================================Steam UGC API=================================
//==============================================================================
extern void* CSteamUGC();

extern UGCQueryHandle_t 	SteamUGC_CreateQueryUserUGCRequest(void* ugc, AccountID_t unAccountID, EUserUGCList eListType, EUGCMatchingUGCType eMatchingUGCType, EUserUGCListSortOrder eSortOrder, AppId_t nCreatorAppID, AppId_t nConsumerAppID, unsigned int unPage);
extern UGCQueryHandle_t 	SteamUGC_CreateQueryAllUGCRequest(void* ugc, EUGCQuery eQueryType, EUGCMatchingUGCType eMatchingeMatchingUGCTypeFileType, AppId_t nCreatorAppID, AppId_t nConsumerAppID, unsigned int unPage);
extern UGCQueryHandle_t 	SteamUGC_CreateQueryUGCDetailsRequest(void* ugc, PublishedFileId_t *pvecPublishedFileID, unsigned int unNumPublishedFileIDs);
extern SteamAPICall_t 		SteamUGC_SendQueryUGCRequest(void* ugc, UGCQueryHandle_t handle);
extern bool 				SteamUGC_GetQueryUGCResult(void* ugc, UGCQueryHandle_t handle, unsigned int index, SteamUGCDetails_t *pDetails);
extern bool 				SteamUGC_GetQueryUGCPreviewURL(void* ugc, UGCQueryHandle_t handle, unsigned int index, char* pchURL, unsigned int cchURLSize);
extern bool 				SteamUGC_GetQueryUGCMetadata(void* ugc, UGCQueryHandle_t handle, unsigned int index, char* pchMetadata, unsigned int cchMetadatasize);
extern bool 				SteamUGC_GetQueryUGCChildren(void* ugc, UGCQueryHandle_t handle, unsigned int index, PublishedFileId_t* pvecPublishedFileID, unsigned int cMaxEntries);
extern bool 				SteamUGC_GetQueryUGCStatistic(void* ugc, UGCQueryHandle_t handle, unsigned int index, EItemStatistic eStatType, unsigned int *pStatValue);
extern unsigned int 		SteamUGC_GetQueryUGCNumAdditionalPreviews(void* ugc, UGCQueryHandle_t handle, unsigned int index);
extern bool 				SteamUGC_GetQueryUGCAdditionalPreview(void* ugc, UGCQueryHandle_t handle, unsigned int index, unsigned int previewIndex, char* pchURLOrVideoID, unsigned int cchURLSize, bool *pbIsImage);
extern unsigned int 		SteamUGC_GetQueryUGCNumKeyValueTags(void* ugc, UGCQueryHandle_t handle, unsigned int index);
extern bool 				SteamUGC_GetQueryUGCKeyValueTag(void* ugc, UGCQueryHandle_t handle, unsigned int index, unsigned int keyValueTagIndex, char* pchKey, unsigned int cchKeySize, char* pchValue, unsigned int cchValueSize);
extern bool 				SteamUGC_ReleaseQueryUGCRequest(void* ugc, UGCQueryHandle_t handle);
extern bool 				SteamUGC_AddRequiredTag(void* ugc, UGCQueryHandle_t handle, const char* pTagName);
extern bool 				SteamUGC_AddExcludedTag(void* ugc, UGCQueryHandle_t handle, const char* pTagName);
extern bool 				SteamUGC_SetReturnKeyValueTags(void* ugc, UGCQueryHandle_t handle, bool bReturnKeyValueTags);
extern bool 				SteamUGC_SetReturnLongDescription(void* ugc, UGCQueryHandle_t handle, bool bReturnLongDescription);
extern bool 				SteamUGC_SetReturnMetadata(void* ugc, UGCQueryHandle_t handle, bool bReturnMetadata);
extern bool 				SteamUGC_SetReturnChildren(void* ugc, UGCQueryHandle_t handle, bool bReturnChildren);
extern bool 				SteamUGC_SetReturnAdditionalPreviews(void* ugc, UGCQueryHandle_t handle, bool bReturnAdditionalPreviews);
extern bool 				SteamUGC_SetReturnTotalOnly(void* ugc, UGCQueryHandle_t handle, bool bReturnTotalOnly);
extern bool 				SteamUGC_SetLanguage(void* ugc, UGCQueryHandle_t handle, const char* pchLanguage);
extern bool 				SteamUGC_SetAllowCachedResponse(void* ugc, UGCQueryHandle_t handle, unsigned int unMaxAgeSeconds);
extern bool 				SteamUGC_SetCloudFileNameFilter(void* ugc, UGCQueryHandle_t handle, const char* pMatchCloudFileName);
extern bool 				SteamUGC_SetMatchAnyTag(void* ugc, UGCQueryHandle_t handle, bool bMatchAnyTag);
extern bool 				SteamUGC_SetSearchText(void* ugc, UGCQueryHandle_t handle, const char* pSearchText);
extern bool 				SteamUGC_SetRankedByTrendDays(void* ugc, UGCQueryHandle_t handle, unsigned int unDays);
extern bool 				SteamUGC_AddRequiredKeyValueTag(void* ugc, UGCQueryHandle_t handle, const char* pKey, const char* pValue);
extern SteamAPICall_t 		SteamUGC_RequestUGCDetails(void* ugc, PublishedFileId_t nPublishedFileID, unsigned int unMaxAgeSeconds);
extern SteamAPICall_t 		SteamUGC_CreateItem(void* ugc, AppId_t nConsumerAppId, EWorkshopFileType eFileType); // create new item for this app with no content attached yet
extern UGCUpdateHandle_t 	SteamUGC_StartItemUpdate(void* ugc, AppId_t nConsumerAppId, PublishedFileId_t nPublishedFileID); // start an UGC item update. Set changed properties before commiting update with CommitItemUpdate()
extern bool 				SteamUGC_SetItemTitle(void* ugc, UGCUpdateHandle_t handle, const char* pchTitle); // change the title of an UGC item
extern bool 				SteamUGC_SetItemDescription(void* ugc, UGCUpdateHandle_t handle, const char* pchDescription); // change the description of an UGC item
extern bool 				SteamUGC_SetItemUpdateLanguage(void* ugc, UGCUpdateHandle_t handle, const char* pchLanguage); // specify the language of the title or description that will be set
extern bool 				SteamUGC_SetItemMetadata(void* ugc, UGCUpdateHandle_t handle, const char* pchMetaData); // change the metadata of an UGC item (max = k_cchDeveloperMetadataMax)
extern bool 				SteamUGC_SetItemVisibility(void* ugc, UGCUpdateHandle_t handle, ERemoteStoragePublishedFileVisibility eVisibility); // change the visibility of an UGC item
extern bool 				SteamUGC_SetItemTags(void* ugc, UGCUpdateHandle_t updateHandle, const SteamParamStringArray_t *pTags); // change the tags of an UGC item
extern bool 				SteamUGC_SetItemContent(void* ugc, UGCUpdateHandle_t handle, const char* pszContentFolder); // update item content from this local folder
extern bool 				SteamUGC_SetItemPreview(void* ugc, UGCUpdateHandle_t handle, const char* pszPreviewFile); //  change preview image file for this item. pszPreviewFile points to local image file, which must be under 1MB in size
extern bool 				SteamUGC_RemoveItemKeyValueTags(void* ugc, UGCUpdateHandle_t handle, const char* pchKey); // remove any existing key-value tags with the specified key
extern bool 				SteamUGC_AddItemKeyValueTag(void* ugc, UGCUpdateHandle_t handle, const char* pchKey, const char* pchValue); // add new key-value tags for the item. Note that there can be multiple values for a tag.
extern SteamAPICall_t 		SteamUGC_SubmitItemUpdate(void* ugc, UGCUpdateHandle_t handle, const char* pchChangeNote); // commit update process started with StartItemUpdate()
extern EItemUpdateStatus 	SteamUGC_GetItemUpdateProgress(void* ugc, UGCUpdateHandle_t handle, uint64_t *punBytesProcessed, uint64_t* punBytesTotal);
extern SteamAPICall_t 		SteamUGC_SetUserItemVote(void* ugc, PublishedFileId_t nPublishedFileID, bool bVoteUp);
extern SteamAPICall_t 		SteamUGC_GetUserItemVote(void* ugc, PublishedFileId_t nPublishedFileID);
extern SteamAPICall_t 		SteamUGC_AddItemToFavorites(void* ugc, AppId_t nAppId, PublishedFileId_t nPublishedFileID);
extern SteamAPICall_t 		SteamUGC_RemoveItemFromFavorites(void* ugc, AppId_t nAppId, PublishedFileId_t nPublishedFileID);
extern SteamAPICall_t 		SteamUGC_SubscribeItem(void* ugc, PublishedFileId_t nPublishedFileID); // subscribe to this item, will be installed ASAP
extern SteamAPICall_t 		SteamUGC_UnsubscribeItem(void* ugc, PublishedFileId_t nPublishedFileID); // unsubscribe from this item, will be uninstalled after game quits
extern unsigned int 		SteamUGC_GetNumSubscribedItems(void* ugc); // number of subscribed items 
extern unsigned int 		SteamUGC_GetSubscribedItems(void* ugc, PublishedFileId_t* pvecPublishedFileID, unsigned int cMaxEntries); // all subscribed item PublishFileIDs
extern unsigned int 		SteamUGC_GetItemState(void* ugc, PublishedFileId_t nPublishedFileID);
extern bool 				SteamUGC_GetItemInstallInfo(void* ugc, PublishedFileId_t nPublishedFileID, uint64_t *punSizeOnDisk, char* pchFolder, unsigned int cchFolderSize, unsigned int *punTimeStamp);
extern bool 				SteamUGC_GetItemDownloadInfo(void* ugc, PublishedFileId_t nPublishedFileID, uint64_t *punBytesDownloaded, uint64_t *punBytesTotal);
extern bool 				SteamUGC_DownloadItem(void* ugc, PublishedFileId_t nPublishedFileID, bool bHighPriority);




//======================================================================
//========================Steam game server API=========================
//======================================================================
void* CSteamGameServer();

extern bool ISteamGameServer_InitGameServer(void* server, unsigned int unIP, unsigned short int usGamePort, unsigned short int usQueryPort, unsigned int unFlags, AppId_t nGameAppId, const char* pchVersionString);
extern void ISteamGameServer_SetProduct(void* server, const char* pszProduct);
extern void ISteamGameServer_SetGameDescription(void* server, const char* pszGameDescription);
extern void ISteamGameServer_SetModDir(void* server, const char* pszModDir);
extern void ISteamGameServer_SetDedicatedServer(void* server, bool bDedicated);
extern void ISteamGameServer_LogOn(void* server, const char* pszToken);
extern void ISteamGameServer_LogOnAnonymous(void* server);
extern void ISteamGameServer_LogOff(void* server);
extern bool ISteamGameServer_BLoggedOn(void* server);
extern bool ISteamGameServer_BSecure(void* server);
extern CSteamID ISteamGameServer_GetSteamID(void* server);
extern bool ISteamGameServer_WasRestartRequested(void* server);
extern void ISteamGameServer_SetMaxPlayerCount(void* server, int cPlayersMax);
extern void ISteamGameServer_SetBotPlayerCount(void* server, int cBotplayers);
extern void ISteamGameServer_SetServerName(void* server, const char* pszServerName);
extern void ISteamGameServer_SetMapName(void* server, const char* pszMapName);
extern void ISteamGameServer_SetPasswordProtected(void* server, bool bPasswordProtected);
extern void ISteamGameServer_SetSpectatorPort(void* server, unsigned short int unSpectatorPort);
extern void ISteamGameServer_SetSpectatorServerName(void* server, const char* pszSpectatorServerName);
extern void ISteamGameServer_ClearAllKeyValues(void* server);
extern void ISteamGameServer_SetKeyValue(void* server, const char* pKey, const char* pValue);
extern void ISteamGameServer_SetGameTags(void* server, const char* pchGameTags);
extern void ISteamGameServer_SetGameData(void* server, const char* pchGameData);
extern void ISteamGameServer_SetRegion(void* server, const char* pszRegion);
extern bool ISteamGameServer_SendUserConnectAndAuthenticate(void* server, unsigned int unIPClient, const void *pvAuthBlob, unsigned int cubAuthBlobSize, CSteamID *pSteamIDUser);
extern CSteamID ISteamGameServer_CreateUnauthenticatedUserConnection(void* server);
extern void ISteamGameServer_SendUserDisconnect(void* server, CSteamID steamIDUser);
extern bool ISteamGameServer_BUpdateUserData(void* server, CSteamID steamIDUser, const char* pchPlayerName, unsigned int uScore);
extern HAuthTicket ISteamGameServer_GetAuthSessionTicket(void* server, void *pTicket, int cbMaxTicket, unsigned int *pcbTicket);
extern EBeginAuthSessionResult ISteamGameServer_BeginAuthSession(void* server, const void *pAuthTicket, int cbAuthTicket, CSteamID steamID);
extern void ISteamGameServer_EndAuthSession(void* server, CSteamID steamID);
extern void ISteamGameServer_CancelAuthTicket(void* server, HAuthTicket hAuthTicket);
extern EUserHasLicenseForAppResult ISteamGameServer_UserHasLicenseForApp(void* server, CSteamID steamID, AppId_t appID);
extern bool ISteamGameServer_RequestUserGroupStatus(void* server, CSteamID steamIDUser, CSteamID steamIDGroup);
extern void ISteamGameServer_GetGameplayStats(void* server);
extern SteamAPICall_t ISteamGameServer_GetServerReputation(void* server);
extern unsigned int ISteamGameServer_GetPublicIP(void* server);
extern bool ISteamGameServer_HandleIncomingPacket(void* server, const void *pData, int cbData, unsigned int srcIP, unsigned short int srcPort);
extern int ISteamGameServer_GetNextOutgoingPacket(void* server, void *pOut, int cbMaxOut, unsigned int *pNetAdr, unsigned short int *pPort);
extern void ISteamGameServer_EnableHeartbeats(void* server, bool bActive);
extern void ISteamGameServer_SetHeartbeatInterval(void* server, int iHeartbeatInterval);
extern void ISteamGameServer_ForceHeartbeat(void* server);
extern SteamAPICall_t ISteamGameServer_AssociateWithClan(void* server, CSteamID steamIDClan);
extern SteamAPICall_t ISteamGameServer_ComputeNewPlayerCompatibility(void* server, CSteamID steamIDNewPlayer);




// Closing statement of `extern "C"{`
#ifdef _cplusplus
}
#endif

#endif //STEAM_C_API_H