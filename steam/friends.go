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

type FriendsGroupID int16

type PersonaState int

type EOverlayToStoreFlag int

const (
	PersonaStateOffline        PersonaState = iota // friend is not currently logged on
	PersonaStateOnline                             // friend is logged on
	PersonaStateBusy                               // user is on, but busy
	PersonaStateAway                               // auto-away feature
	PersonaStateSnooze                             // auto-away for a long time
	PersonaStateLookingToTrade                     // Online, trading
	PersonaStateLookingToPlay                      // Online, wanting to play
)

//-----------------------------------------------------------------------------
// Purpose: set of relationships to other users
//-----------------------------------------------------------------------------
/*enum EFriendRelationship
{
	k_EFriendRelationshipNone = 0,
	k_EFriendRelationshipBlocked = 1,			// this doesn't get stored; the user has just done an Ignore on an friendship invite
	k_EFriendRelationshipRequestRecipient = 2,
	k_EFriendRelationshipFriend = 3,
	k_EFriendRelationshipRequestInitiator = 4,
	k_EFriendRelationshipIgnored = 5,			// this is stored; the user has explicit blocked this other user from comments/chat/etc
	k_EFriendRelationshipIgnoredFriend = 6,
	k_EFriendRelationshipSuggested = 7,

	// keep this updated
	k_EFriendRelationshipMax = 8,
};


enum EOverlayToStoreFlag
{
	k_EOverlayToStoreFlag_None = 0,
	k_EOverlayToStoreFlag_AddToCart = 1,
	k_EOverlayToStoreFlag_AddToCartAndShow = 2,
};


//-----------------------------------------------------------------------------
// Purpose: Chat Entry Types (previously was only friend-to-friend message types)
//-----------------------------------------------------------------------------
enum EChatEntryType
{
	k_EChatEntryTypeInvalid = 0,
	k_EChatEntryTypeChatMsg = 1,		// Normal text message from another user
	k_EChatEntryTypeTyping = 2,			// Another user is typing (not used in multi-user chat)
	k_EChatEntryTypeInviteGame = 3,		// Invite from other user into that users current game
	k_EChatEntryTypeEmote = 4,			// text emote message (deprecated, should be treated as ChatMsg)
	//k_EChatEntryTypeLobbyGameStart = 5,	// lobby game is starting (dead - listen for LobbyGameCreated_t callback instead)
	k_EChatEntryTypeLeftConversation = 6, // user has left the conversation ( closed chat window )
	// Above are previous FriendMsgType entries, now merged into more generic chat entry types
	k_EChatEntryTypeEntered = 7,		// user has entered the conversation (used in multi-user chat and group chat)
	k_EChatEntryTypeWasKicked = 8,		// user was kicked (data: 64-bit steamid of actor performing the kick)
	k_EChatEntryTypeWasBanned = 9,		// user was banned (data: 64-bit steamid of actor performing the ban)
	k_EChatEntryTypeDisconnected = 10,	// user disconnected
	k_EChatEntryTypeHistoricalChat = 11,	// a chat message from user's chat history or offilne message
	k_EChatEntryTypeReserved1 = 12,
	k_EChatEntryTypeReserved2 = 13,
	k_EChatEntryTypeLinkBlocked = 14, // a link was removed by the chat filter.
};

*/

// ISteamFriends is a handler for the steam friends API.
type ISteamFriends struct {
	unsafe.Pointer
}

// Friends returns the current steam user as a player.
func Friends() ISteamFriends {
	return ISteamFriends{C.CSteamFriends()}
}

// GetPersonaName returns the local players name - guaranteed to not be empty.
// this is the same name as on the users community profile page this is stored
// in UTF-8 format like all the other interface functions that return a string,
// it's important that this pointer is not saved off; it will eventually be
// free'd or re-allocated.
func (f ISteamFriends) GetPersonaName() string {
	return C.GoString(C.SteamFriends_GetPersonaName(f.Pointer))
}

// SetPersonaName Sets the player name, stores it on the server and publishes
// the changes to all friends who are online. Changes take place locally
// immediately, and a PersonaStateChange_t is posted, presuming success.
// The final results are available through the return value SteamAPICall_t,
// using SetPersonaNameResponse_t. If the name change fails to happen on the
// server, then an additional global PersonaStateChange_t will be posted to
// change the name back, in addition to the SetPersonaNameResponse_t callback.
func (f ISteamFriends) SetPersonaName(name string) APICall {
	pchPersonaName := C.CString(name)
	defer C.free(unsafe.Pointer(pchPersonaName))
	return APICall(C.SteamFriends_SetPersonaName(f.Pointer, pchPersonaName))
}

// GetPersonaState gets the status of the user.
func (f ISteamFriends) GetPersonaState() PersonaState {
	return PersonaState(C.SteamFriends_GetPersonaState(f.Pointer))
}

// GetFriendCount takes a set of k_EFriendFlags and returns the number of users
// the client knows about who meet that criteria then GetFriendByIndex() can
// then be used to return the id's of each of those users.
func (f ISteamFriends) GetFriendCount(iFriendFlags int32) int32 {
	return int32(C.SteamFriends_GetFriendCount(f.Pointer, C.int(iFriendFlags)))
}

/*
// GetFriendByIndex returns the steamID of a user. iFriend is a index of range
// [0, GetFriendCount()). iFriendsFlags must be the same value as used in
// GetFriendCount(). The returned CSteamID can then be used by all the functions
// below to access details about the user
func (f ISteamFriends) GetFriendByIndex(iFriend, iFriendFlags int32) CSteamID {
	return C.SteamFriends_GetFriendByIndex(f.Pointer, iFriend, iFriendFlags)
}

// GetFriendRelationship returns a relationship to a user.
func (f ISteamFriends) GetFriendRelationship(steamIDFriend CSteamID) EFriendRelationship {
	return C.SteamFriends_GetFriendRelationship(f.Pointer, steamIDFriend)
}

// GetFriendPersonaState returns the current status of the specified user this
// will only be known by the local user if steamIDFriend is in their friends
// list; on the same game server; in a chat room or lobby; or in a small group
// with the local user.
func (f ISteamFriends) GetFriendPersonaState(steamIDFriend CSteamID) PersonaState {
	return C.SteamFriends_GetFriendPersonaState(f.Pointer, steamIDFriend)
}

// GetFriendPersonaName returns the name of another user - guaranteed to not be
// empty. Same rules as GetFriendPersonaState() apply as to whether or not the
// user knowns the name of the other user note that on first joining a lobby,
// chat room or game server the local user will not known the name of the other
// users automatically; that information will arrive asynchronously.
func (f ISteamFriends) GetFriendPersonaName(steamIDFriend CSteamID) string {
	return C.SteamFriends_GetFriendPersonaName(f.Pointer, steamIDFriend)
}

// GetFriendGamePlayed returns true if the friend is actually in a game, and
// fills in pFriendGameInfo with an extra details.
func (f ISteamFriends) GetFriendGamePlayed(steamIDFriend CSteamID, pFriendGameInfo *FriendGameInfo_t) bool {
	return bool(C.SteamFriends_GetFriendGamePlayed(f.Pointer, steamIDFriend, pFriendGameInfo))
}

// GetFriendPersonaNameHistory accesses old friends names - returns an empty
// string when their are no more items in the history.
func (f ISteamFriends) GetFriendPersonaNameHistory(steamIDFriend CSteamID, iPersonaName int) string {
	return C.SteamFriends_GetFriendPersonaNameHistory(f.Pointer)
}

// friends steam level
func (f ISteamFriends) GetFriendSteamLevel(CSteamID steamIDFriend) int {
	return C.SteamFriends_GetFriendSteamLevel(f.Pointer)
}

// Returns nickname the current user has set for the specified player. Returns NULL if the no nickname has been set for that player.
func (f ISteamFriends) GetPlayerNickname(CSteamID steamIDPlayer) string {
	return C.SteamFriends_GetPlayerNickname(f.Pointer)
}
*/

// GetFriendsGroupCount returns the number of friends groups.
func (f ISteamFriends) GetFriendsGroupCount() int {
	return int(C.SteamFriends_GetFriendsGroupCount(f.Pointer))
}

// GetFriendsGroupIDByIndex returns the friends group ID for the given index
// (invalid indices return k_FriendsGroupID_Invalid).
func (f ISteamFriends) GetFriendsGroupIDByIndex(iFG int) FriendsGroupID {
	return FriendsGroupID(C.SteamFriends_GetFriendsGroupIDByIndex(f.Pointer, C.int(iFG)))
}

// GetFriendsGroupName returns the name for the given friends group (NULL in
// the case of invalid friends group IDs).
func (f ISteamFriends) GetFriendsGroupName(friendsGroupID FriendsGroupID) string {
	ret := C.SteamFriends_GetFriendsGroupName(f.Pointer, C.FriendsGroupID_t(friendsGroupID))
	out := C.GoString(ret)
	C.free(unsafe.Pointer(ret))
	return out
}

// GetFriendsGroupMembersCount returns the number of members in a given friends
// group.
func (f ISteamFriends) GetFriendsGroupMembersCount(friendsGroupID FriendsGroupID) int {
	return int(C.SteamFriends_GetFriendsGroupMembersCount(f.Pointer, C.FriendsGroupID_t(friendsGroupID)))
}

/*
//func (f ISteamFriends) void SteamFriends_GetFriendsGroupMembersList( FriendsGroupID_t friendsGroupID, OUT_ARRAY_CALL(nMembersCount, GetFriendsGroupMembersCount, friendsGroupID ) CSteamID *pOutSteamIDMembers, int nMembersCount ) = 0;
func (f ISteamFriends) HasFriend(CSteamID steamIDFriend, int iFriendFlags) bool {
	return C.SteamFriends_HasFriend(f.Pointer)
}
*/

// GetClanCount returns clan (group) iteration and access functions.
func (f ISteamFriends) GetClanCount() int {
	return int(C.SteamFriends_GetClanCount(f.Pointer))
}

/*
func (f ISteamFriends) GetClanByIndex(int iClan) CSteamID {
	return C.SteamFriends_GetClanByIndex(f.Pointer)
}

func (f ISteamFriends) GetClanName(CSteamID steamIDClan) string {
	return C.SteamFriends_GetClanName(f.Pointer)
}

func (f ISteamFriends) GetClanTag(CSteamID steamIDClan) string {
	return C.SteamFriends_GetClanTag(f.Pointer)
}

func (f ISteamFriends) GetClanActivityCounts(CSteamID steamIDClan, int *pnOnline, int *pnInGame, int *pnChatting) bool {
	return C.SteamFriends_GetClanActivityCounts(f.Pointer)
}

//func (f ISteamFriends) SteamAPICall_t SteamFriends_DownloadClanActivityCounts( ARRAY_COUNT(cClansToRequest) CSteamID *psteamIDClans, int cClansToRequest ) = 0;
func (f ISteamFriends) GetFriendCountFromSource(CSteamID steamIDSource) int {
	return C.SteamFriends_GetFriendCountFromSource(f.Pointer)
}

func (f ISteamFriends) GetFriendFromSourceByIndex(CSteamID steamIDSource, int iFriend) CSteamID {
	return C.SteamFriends_GetFriendFromSourceByIndex(f.Pointer)
}

func (f ISteamFriends) IsUserInSource(CSteamID steamIDUser, CSteamID steamIDSource) bool {
	return C.SteamFriends_IsUserInSource(f.Pointer)
}

func (f ISteamFriends) SetInGameVoiceSpeaking(CSteamID steamIDUser, bool bSpeaking) void {
	return C.SteamFriends_SetInGameVoiceSpeaking(f.Pointer)
}
*/

// ActivateGameOverlay activates the game overlay, with an optional dialog to
// open valid options are "Friends", "Community", "Players", "Settings",
// "OfficialGameGroup", "Stats", "Achievements"
func (f ISteamFriends) ActivateGameOverlay(pchDialog string) {
	cpchDialog := C.CString(pchDialog)
	defer C.free(unsafe.Pointer(cpchDialog))
	C.SteamFriends_ActivateGameOverlay(f.Pointer, cpchDialog)
}

/*
func (f ISteamFriends) ActivateGameOverlayToUser(string pchDialog, CSteamID steamID) void {
	return C.SteamFriends_ActivateGameOverlayToUser(f.Pointer)
}
*/

// ActivateGameOverlayToWebPage activates game overlay web browser directly to
// the specified URL full address with protocol type is required,
// e.g. http://www.steamgames.com/
func (f ISteamFriends) ActivateGameOverlayToWebPage(pchURL string) {
	cpchURL := C.CString(pchURL)
	defer C.free(unsafe.Pointer(cpchURL))
	C.SteamFriends_ActivateGameOverlayToWebPage(f.Pointer, cpchURL)
}

// ActivateGameOverlayToStore activates game overlay to store page for app.
func (f ISteamFriends) ActivateGameOverlayToStore(nAppID AppId, eFlag EOverlayToStoreFlag) {
	C.SteamFriends_ActivateGameOverlayToStore(f.Pointer, C.AppId_t(nAppID), C.EOverlayToStoreFlag(eFlag))
}

/*
func (f ISteamFriends) SetPlayedWith(CSteamID steamIDUserPlayedWith) void {
	return C.SteamFriends_SetPlayedWith(f.Pointer)
}


func (f ISteamFriends) ActivateGameOverlayInviteDialog(CSteamID steamIDLobby) void {
	return C.SteamFriends_ActivateGameOverlayInviteDialog(f.Pointer)
}

func (f ISteamFriends) GetSmallFriendAvatar(CSteamID steamIDFriend) int {
	return C.SteamFriends_GetSmallFriendAvatar(f.Pointer)
}

func (f ISteamFriends) GetMediumFriendAvatar(CSteamID steamIDFriend) int {
	return C.SteamFriends_GetMediumFriendAvatar(f.Pointer)
}

func (f ISteamFriends) GetLargeFriendAvatar(CSteamID steamIDFriend) int {
	return C.SteamFriends_GetLargeFriendAvatar(f.Pointer)
}

func (f ISteamFriends) RequestUserInformation(CSteamID steamIDUser, bool bRequireNameOnly) bool {
	return C.SteamFriends_RequestUserInformation(f.Pointer)
}

func (f ISteamFriends) RequestClanOfficerList(CSteamID steamIDClan) SteamAPICall_t {
	return C.SteamFriends_RequestClanOfficerList(f.Pointer)
}

func (f ISteamFriends) GetClanOwner(CSteamID steamIDClan) CSteamID {
	return C.SteamFriends_GetClanOwner(f.Pointer)
}

func (f ISteamFriends) GetClanOfficerCount(CSteamID steamIDClan) int {
	return C.SteamFriends_GetClanOfficerCount(f.Pointer)
}

func (f ISteamFriends) GetClanOfficerByIndex(CSteamID steamIDClan, int iOfficer) CSteamID {
	return C.SteamFriends_GetClanOfficerByIndex(f.Pointer)
}
*/

// GetUserRestrictions returns the user restrictions. If current user is chat
// restricted, he can't send or receive any text and voice chat messages. The
// user can't see custom avatars. But the user can be online and send or receive
// game invites. A chat restricted user can't add friends or join any groups.
func (f ISteamFriends) GetUserRestrictions() uint32 {
	return uint32(C.SteamFriends_GetUserRestrictions(f.Pointer))
}

// SetRichPresence sets a Rich Presence data.
//
// Rich Presence data is automatically shared between friends who are in the same game
// Each user has a set of Key/Value pairs
// Up to 20 different keys can be set
// There are two magic keys:
//		"status"  - a UTF-8 string that will show up in the 'view game info' dialog in the Steam friends list
//		"connect" - a UTF-8 string that contains the command-line for how a friend can connect to a game
// GetFriendRichPresence() returns an empty string "" if no value is set
// SetRichPresence() to a NULL or an empty string deletes the key
// You can iterate the current set of keys for a friend with GetFriendRichPresenceKeyCount()
// and GetFriendRichPresenceKeyByIndex() (typically only used for debugging)
func (f ISteamFriends) SetRichPresence(pchKey, pchValue string) bool {
	cpchKey := C.CString(pchKey)
	defer C.free(unsafe.Pointer(cpchKey))
	cpchValue := C.CString(pchValue)
	defer C.free(unsafe.Pointer(cpchValue))
	return bool(C.SteamFriends_SetRichPresence(f.Pointer, cpchKey, cpchValue))
}

// ClearRichPresence clears all Rich Presence data.
func (f ISteamFriends) ClearRichPresence() {
	C.SteamFriends_ClearRichPresence(f.Pointer)
}

/*
func (f ISteamFriends) GetFriendRichPresence(CSteamID steamIDFriend, string pchKey) string {
	return C.SteamFriends_GetFriendRichPresence(f.Pointer)
}

func (f ISteamFriends) GetFriendRichPresenceKeyCount(CSteamID steamIDFriend) int {
	return C.SteamFriends_GetFriendRichPresenceKeyCount(f.Pointer)
}

func (f ISteamFriends) GetFriendRichPresenceKeyByIndex(CSteamID steamIDFriend, int iKey) string {
	return C.SteamFriends_GetFriendRichPresenceKeyByIndex(f.Pointer)
}

func (f ISteamFriends) RequestFriendRichPresence(CSteamID steamIDFriend) void {
	return C.SteamFriends_RequestFriendRichPresence(f.Pointer)
}

func (f ISteamFriends) InviteUserToGame(CSteamID steamIDFriend, string pchConnectString) bool {
	return C.SteamFriends_InviteUserToGame(f.Pointer)
}
*/

// GetCoplayFriendCount returns the number of users recently played with.
func (f ISteamFriends) GetCoplayFriendCount() int {
	return int(C.SteamFriends_GetCoplayFriendCount(f.Pointer))
}

/*
func (f ISteamFriends) GetCoplayFriend(int iCoplayFriend) CSteamID {
	return C.SteamFriends_GetCoplayFriend(f.Pointer)
}


func (f ISteamFriends) GetFriendCoplayTime(CSteamID steamIDFriend) int {
	return C.SteamFriends_GetFriendCoplayTime(f.Pointer)
}

func (f ISteamFriends) GetFriendCoplayGame(CSteamID steamIDFriend) AppId_t {
	return C.SteamFriends_GetFriendCoplayGame(f.Pointer)
}

func (f ISteamFriends) JoinClanChatRoom(CSteamID steamIDClan) SteamAPICall_t {
	return C.SteamFriends_JoinClanChatRoom(f.Pointer)
}

func (f ISteamFriends) LeaveClanChatRoom(CSteamID steamIDClan) bool {
	return C.SteamFriends_LeaveClanChatRoom(f.Pointer)
}

func (f ISteamFriends) GetClanChatMemberCount(CSteamID steamIDClan) int {
	return C.SteamFriends_GetClanChatMemberCount(f.Pointer)
}

func (f ISteamFriends) GetChatMemberByIndex(CSteamID steamIDClan, int iUser) CSteamID {
	return C.SteamFriends_GetChatMemberByIndex(f.Pointer)
}

func (f ISteamFriends) SendClanChatMessage(CSteamID steamIDClanChat, string pchText) bool {
	return C.SteamFriends_SendClanChatMessage(f.Pointer)
}

//func (f ISteamFriends) int SteamFriends_GetClanChatMessage( CSteamID steamIDClanChat, int iMessage, void *prgchText, int cchTextMax, EChatEntryType *peChatEntryType, OUT_STRUCT() CSteamID *psteamidChatter );
func (f ISteamFriends) IsClanChatAdmin(CSteamID steamIDClanChat, CSteamID steamIDUser) bool {
	return C.SteamFriends_IsClanChatAdmin(f.Pointer)
}

func (f ISteamFriends) IsClanChatWindowOpenInSteam(CSteamID steamIDClanChat) bool {
	return C.SteamFriends_IsClanChatWindowOpenInSteam(f.Pointer)
}

func (f ISteamFriends) OpenClanChatWindowInSteam(CSteamID steamIDClanChat) bool {
	return C.SteamFriends_OpenClanChatWindowInSteam(f.Pointer)
}

func (f ISteamFriends) CloseClanChatWindowInSteam(CSteamID steamIDClanChat) bool {
	return C.SteamFriends_CloseClanChatWindowInSteam(f.Pointer)
}
*/

func (f ISteamFriends) SetListenForFriendsMessages(bInterceptEnabled bool) bool {
	return bool(C.SteamFriends_SetListenForFriendsMessages(f.Pointer, C._Bool(bInterceptEnabled)))
}

/*
func (f ISteamFriends) ReplyToFriendMessage(CSteamID steamIDFriend, string pchMsgToSend) bool {
	return C.SteamFriends_ReplyToFriendMessage(f.Pointer)
}

func (f ISteamFriends) GetFriendMessage(CSteamID steamIDFriend, int iMessageID, void *pvData, int cubData, EChatEntryType *peChatEntryType) int {
	return C.SteamFriends_GetFriendMessage(f.Pointer)
}

func (f ISteamFriends) GetFollowerCount(CSteamID steamID) SteamAPICall_t {
	return C.SteamFriends_GetFollowerCount(f.Pointer)
}

func (f ISteamFriends) IsFollowing(CSteamID steamID) SteamAPICall_t {
	return C.SteamFriends_IsFollowing(f.Pointer)
}
*/

func (f ISteamFriends) EnumerateFollowingList(unStartIndex uint32) SteamAPICall {
	return SteamAPICall(C.SteamFriends_EnumerateFollowingList(f.Pointer, C.uint(unStartIndex)))
}
