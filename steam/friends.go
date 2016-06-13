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

// FriendsGroupID is an ID for friends group.
type FriendsGroupID int16

// PersonaState is a enum for a user steam status (offline/online/away)
type PersonaState int

// EOverlayToStoreFlag are passed as parameters to the store
type EOverlayToStoreFlag int

// EFriendRelationship is an enum for friend relationships
type EFriendRelationship int32

// The different friend relationships.
const (
	EFriendRelationshipNone             EFriendRelationship = 0
	EFriendRelationshipBlocked          EFriendRelationship = 1 // this doesn't get stored; the user has just done an Ignore on an friendship invite
	EFriendRelationshipRequestRecipient EFriendRelationship = 2
	EFriendRelationshipFriend           EFriendRelationship = 3
	EFriendRelationshipRequestInitiator EFriendRelationship = 4
	EFriendRelationshipIgnored          EFriendRelationship = 5 // this is stored; the user has explicit blocked this other user from comments/chat/etc
	EFriendRelationshipIgnoredFriend    EFriendRelationship = 6
	EFriendRelationshipSuggested        EFriendRelationship = 7

	// keep this updated
	EFriendRelationshipMax = 8
)

// the different enums for PersonaState
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

// IFriends is an interface to accessing information about individual users,
// that can be a friend, in a group, on a game server or in a lobby with the
// local user.
type IFriends struct {
	unsafe.Pointer
}

// Friends returns the current steam user as a player.
func Friends() IFriends {
	return IFriends{C.CSteamFriends()}
}

// GetPersonaName returns the local players name - guaranteed to not be empty.
// this is the same name as on the users community profile page this is stored
// in UTF-8 format like all the other interface functions that return a string,
// it's important that this pointer is not saved off; it will eventually be
// free'd or re-allocated.
func (f IFriends) GetPersonaName() string {
	return C.GoString(C.SteamFriends_GetPersonaName(f.Pointer))
}

// SetPersonaName Sets the player name, stores it on the server and publishes
// the changes to all friends who are online. Changes take place locally
// immediately, and a PersonaStateChange_t is posted, presuming success.
// The final results are available through the return value SteamAPICall_t,
// using SetPersonaNameResponse_t. If the name change fails to happen on the
// server, then an additional global PersonaStateChange_t will be posted to
// change the name back, in addition to the SetPersonaNameResponse_t callback.
func (f IFriends) SetPersonaName(name string) APICall {
	cname := C.CString(name)
	ret := APICall(C.SteamFriends_SetPersonaName(f.Pointer, cname))
	C.free(unsafe.Pointer(cname))
	return ret
}

// GetPersonaState gets the status of the user.
func (f IFriends) GetPersonaState() PersonaState {
	return PersonaState(C.SteamFriends_GetPersonaState(f.Pointer))
}

// GetFriendCount takes a set of k_EFriendFlags and returns the number of users
// the client knows about who meet that criteria then GetFriendByIndex() can
// then be used to return the id's of each of those users.
func (f IFriends) GetFriendCount(flags int32) int {
	return int(C.SteamFriends_GetFriendCount(f.Pointer, C.int(flags)))
}

// GetFriendByIndex returns the steamID of a user. index is a index of range
// [0, GetFriendCount()). flags must be the same value as used in
// GetFriendCount(). The returned SteamID can then be used by all the functions
// below to access details about the user.
func (f IFriends) GetFriendByIndex(index, flags int32) SteamID {
	ret := C.SteamFriends_GetFriendByIndex(f.Pointer, C.int(index), C.int(flags))
	return *(*SteamID)(unsafe.Pointer(&ret))
}

// GetFriendRelationship returns a relationship to a user.
func (f IFriends) GetFriendRelationship(ID SteamID) EFriendRelationship {
	return EFriendRelationship(C.SteamFriends_GetFriendRelationship(f.Pointer, *(*C.CSteamID)(unsafe.Pointer(&ID))))
}

// GetFriendPersonaState returns the current status of the specified user this
// will only be known by the local user if steamIDFriend is in their friends
// list; on the same game server; in a chat room or lobby; or in a small group
// with the local user.
func (f IFriends) GetFriendPersonaState(ID SteamID) PersonaState {
	return PersonaState(C.SteamFriends_GetFriendPersonaState(f.Pointer, *(*C.CSteamID)(unsafe.Pointer(&ID))))
}

// GetFriendPersonaName returns the name of another user - guaranteed to not be
// empty. Same rules as GetFriendPersonaState() apply as to whether or not the
// user knowns the name of the other user note that on first joining a lobby,
// chat room or game server the local user will not known the name of the other
// users automatically; that information will arrive asynchronously.
func (f IFriends) GetFriendPersonaName(ID SteamID) string {
	return C.GoString(C.SteamFriends_GetFriendPersonaName(f.Pointer, *(*C.CSteamID)(unsafe.Pointer(&ID))))
}

/*
// GetFriendGamePlayed returns true if the friend is actually in a game, and
// fills in pFriendGameInfo with an extra details.
func (f IFriends) GetFriendGamePlayed(friendID SteamID, pFriendGameInfo *FriendGameInfo_t) bool {
	return bool(C.SteamFriends_GetFriendGamePlayed(f.Pointer, *(*C.CSteamID)(unsafe.Pointer(&friendID)), pFriendGameInfo))
}

// GetFriendPersonaNameHistory accesses old friends names - returns an empty
// string when their are no more items in the history.
func (f IFriends) GetFriendPersonaNameHistory(steamIDFriend SteamID, iPersonaName int) string {
	return C.SteamFriends_GetFriendPersonaNameHistory(f.Pointer)
}

// friends steam level
func (f IFriends) GetFriendSteamLevel(SteamID steamIDFriend) int {
	return C.SteamFriends_GetFriendSteamLevel(f.Pointer)
}

// Returns nickname the current user has set for the specified player. Returns NULL if the no nickname has been set for that player.
func (f IFriends) GetPlayerNickname(SteamID steamIDPlayer) string {
	return C.SteamFriends_GetPlayerNickname(f.Pointer)
}*/

// GetFriendsGroupCount returns the number of friends groups.
func (f IFriends) GetFriendsGroupCount() int {
	return int(C.SteamFriends_GetFriendsGroupCount(f.Pointer))
}

// GetFriendsGroupIDByIndex returns the friends group ID for the given index
// (invalid indices return k_FriendsGroupID_Invalid).
func (f IFriends) GetFriendsGroupIDByIndex(iFG int) FriendsGroupID {
	return FriendsGroupID(C.SteamFriends_GetFriendsGroupIDByIndex(f.Pointer, C.int(iFG)))
}

// GetFriendsGroupName returns the name for the given friends group (NULL in
// the case of invalid friends group IDs).
func (f IFriends) GetFriendsGroupName(friendsGroupID FriendsGroupID) string {
	ret := C.SteamFriends_GetFriendsGroupName(f.Pointer, C.FriendsGroupID_t(friendsGroupID))
	out := C.GoString(ret)
	C.free(unsafe.Pointer(ret))
	return out
}

// GetFriendsGroupMembersCount returns the number of members in a given friends
// group.
func (f IFriends) GetFriendsGroupMembersCount(friendsGroupID FriendsGroupID) int {
	return int(C.SteamFriends_GetFriendsGroupMembersCount(f.Pointer, C.FriendsGroupID_t(friendsGroupID)))
}

/*
//func (f IFriends) void SteamFriends_GetFriendsGroupMembersList( FriendsGroupID_t friendsGroupID, OUT_ARRAY_CALL(nMembersCount, GetFriendsGroupMembersCount, friendsGroupID ) SteamID *pOutSteamIDMembers, int nMembersCount ) = 0;
func (f IFriends) HasFriend(SteamID steamIDFriend, int iFriendFlags) bool {
	return C.SteamFriends_HasFriend(f.Pointer)
}
*/

// GetClanCount returns clan (group) iteration and access functions.
func (f IFriends) GetClanCount() int {
	return int(C.SteamFriends_GetClanCount(f.Pointer))
}

/*
func (f IFriends) GetClanByIndex(int iClan) SteamID {
	return C.SteamFriends_GetClanByIndex(f.Pointer)
}

func (f IFriends) GetClanName(SteamID steamIDClan) string {
	return C.SteamFriends_GetClanName(f.Pointer)
}

func (f IFriends) GetClanTag(SteamID steamIDClan) string {
	return C.SteamFriends_GetClanTag(f.Pointer)
}

func (f IFriends) GetClanActivityCounts(SteamID steamIDClan, int *pnOnline, int *pnInGame, int *pnChatting) bool {
	return C.SteamFriends_GetClanActivityCounts(f.Pointer)
}

//func (f IFriends) SteamAPICall_t SteamFriends_DownloadClanActivityCounts( ARRAY_COUNT(cClansToRequest) SteamID *psteamIDClans, int cClansToRequest ) = 0;
func (f IFriends) GetFriendCountFromSource(SteamID steamIDSource) int {
	return C.SteamFriends_GetFriendCountFromSource(f.Pointer)
}

func (f IFriends) GetFriendFromSourceByIndex(SteamID steamIDSource, int iFriend) SteamID {
	return C.SteamFriends_GetFriendFromSourceByIndex(f.Pointer)
}

func (f IFriends) IsUserInSource(SteamID steamIDUser, SteamID steamIDSource) bool {
	return C.SteamFriends_IsUserInSource(f.Pointer)
}

func (f IFriends) SetInGameVoiceSpeaking(SteamID steamIDUser, bool bSpeaking) void {
	return C.SteamFriends_SetInGameVoiceSpeaking(f.Pointer)
}
*/

// ActivateGameOverlay activates the game overlay, with an optional dialog to
// open valid options are "Friends", "Community", "Players", "Settings",
// "OfficialGameGroup", "Stats", "Achievements"
func (f IFriends) ActivateGameOverlay(pchDialog string) {
	cpchDialog := C.CString(pchDialog)
	defer C.free(unsafe.Pointer(cpchDialog))
	C.SteamFriends_ActivateGameOverlay(f.Pointer, cpchDialog)
}

/*
func (f IFriends) ActivateGameOverlayToUser(string pchDialog, SteamID steamID) void {
	return C.SteamFriends_ActivateGameOverlayToUser(f.Pointer)
}
*/

// ActivateGameOverlayToWebPage activates game overlay web browser directly to
// the specified URL full address with protocol type is required,
// e.g. http://www.steamgames.com/
func (f IFriends) ActivateGameOverlayToWebPage(pchURL string) {
	cpchURL := C.CString(pchURL)
	defer C.free(unsafe.Pointer(cpchURL))
	C.SteamFriends_ActivateGameOverlayToWebPage(f.Pointer, cpchURL)
}

// ActivateGameOverlayToStore activates game overlay to store page for app.
func (f IFriends) ActivateGameOverlayToStore(nAppID AppID, eFlag EOverlayToStoreFlag) {
	C.SteamFriends_ActivateGameOverlayToStore(f.Pointer, C.AppId_t(nAppID), C.EOverlayToStoreFlag(eFlag))
}

/*
func (f IFriends) SetPlayedWith(SteamID steamIDUserPlayedWith) void {
	return C.SteamFriends_SetPlayedWith(f.Pointer)
}


func (f IFriends) ActivateGameOverlayInviteDialog(SteamID steamIDLobby) void {
	return C.SteamFriends_ActivateGameOverlayInviteDialog(f.Pointer)
}

func (f IFriends) GetSmallFriendAvatar(SteamID steamIDFriend) int {
	return C.SteamFriends_GetSmallFriendAvatar(f.Pointer)
}

func (f IFriends) GetMediumFriendAvatar(SteamID steamIDFriend) int {
	return C.SteamFriends_GetMediumFriendAvatar(f.Pointer)
}

func (f IFriends) GetLargeFriendAvatar(SteamID steamIDFriend) int {
	return C.SteamFriends_GetLargeFriendAvatar(f.Pointer)
}

func (f IFriends) RequestUserInformation(SteamID steamIDUser, bool bRequireNameOnly) bool {
	return C.SteamFriends_RequestUserInformation(f.Pointer)
}

func (f IFriends) RequestClanOfficerList(SteamID steamIDClan) SteamAPICall_t {
	return C.SteamFriends_RequestClanOfficerList(f.Pointer)
}

func (f IFriends) GetClanOwner(SteamID steamIDClan) SteamID {
	return C.SteamFriends_GetClanOwner(f.Pointer)
}

func (f IFriends) GetClanOfficerCount(SteamID steamIDClan) int {
	return C.SteamFriends_GetClanOfficerCount(f.Pointer)
}

func (f IFriends) GetClanOfficerByIndex(SteamID steamIDClan, int iOfficer) SteamID {
	return C.SteamFriends_GetClanOfficerByIndex(f.Pointer)
}
*/

// GetUserRestrictions returns the user restrictions. If current user is chat
// restricted, he can't send or receive any text and voice chat messages. The
// user can't see custom avatars. But the user can be online and send or receive
// game invites. A chat restricted user can't add friends or join any groups.
func (f IFriends) GetUserRestrictions() uint32 {
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
func (f IFriends) SetRichPresence(pchKey, pchValue string) bool {
	cpchKey := C.CString(pchKey)
	defer C.free(unsafe.Pointer(cpchKey))
	cpchValue := C.CString(pchValue)
	defer C.free(unsafe.Pointer(cpchValue))
	return bool(C.SteamFriends_SetRichPresence(f.Pointer, cpchKey, cpchValue))
}

// ClearRichPresence clears all Rich Presence data.
func (f IFriends) ClearRichPresence() {
	C.SteamFriends_ClearRichPresence(f.Pointer)
}

/*
func (f IFriends) GetFriendRichPresence(SteamID steamIDFriend, string pchKey) string {
	return C.SteamFriends_GetFriendRichPresence(f.Pointer)
}

func (f IFriends) GetFriendRichPresenceKeyCount(SteamID steamIDFriend) int {
	return C.SteamFriends_GetFriendRichPresenceKeyCount(f.Pointer)
}

func (f IFriends) GetFriendRichPresenceKeyByIndex(SteamID steamIDFriend, int iKey) string {
	return C.SteamFriends_GetFriendRichPresenceKeyByIndex(f.Pointer)
}

func (f IFriends) RequestFriendRichPresence(SteamID steamIDFriend) void {
	return C.SteamFriends_RequestFriendRichPresence(f.Pointer)
}

func (f IFriends) InviteUserToGame(SteamID steamIDFriend, string pchConnectString) bool {
	return C.SteamFriends_InviteUserToGame(f.Pointer)
}
*/

// GetCoplayFriendCount returns the number of users recently played with.
func (f IFriends) GetCoplayFriendCount() int {
	return int(C.SteamFriends_GetCoplayFriendCount(f.Pointer))
}

/*
func (f IFriends) GetCoplayFriend(int iCoplayFriend) SteamID {
	return C.SteamFriends_GetCoplayFriend(f.Pointer)
}


func (f IFriends) GetFriendCoplayTime(SteamID steamIDFriend) int {
	return C.SteamFriends_GetFriendCoplayTime(f.Pointer)
}

func (f IFriends) GetFriendCoplayGame(SteamID steamIDFriend) AppId_t {
	return C.SteamFriends_GetFriendCoplayGame(f.Pointer)
}

func (f IFriends) JoinClanChatRoom(SteamID steamIDClan) SteamAPICall_t {
	return C.SteamFriends_JoinClanChatRoom(f.Pointer)
}

func (f IFriends) LeaveClanChatRoom(SteamID steamIDClan) bool {
	return C.SteamFriends_LeaveClanChatRoom(f.Pointer)
}

func (f IFriends) GetClanChatMemberCount(SteamID steamIDClan) int {
	return C.SteamFriends_GetClanChatMemberCount(f.Pointer)
}

func (f IFriends) GetChatMemberByIndex(SteamID steamIDClan, int iUser) SteamID {
	return C.SteamFriends_GetChatMemberByIndex(f.Pointer)
}

func (f IFriends) SendClanChatMessage(SteamID steamIDClanChat, string pchText) bool {
	return C.SteamFriends_SendClanChatMessage(f.Pointer)
}

//func (f IFriends) int SteamFriends_GetClanChatMessage( SteamID steamIDClanChat, int iMessage, void *prgchText, int cchTextMax, EChatEntryType *peChatEntryType, OUT_STRUCT() SteamID *psteamidChatter );
func (f IFriends) IsClanChatAdmin(SteamID steamIDClanChat, SteamID steamIDUser) bool {
	return C.SteamFriends_IsClanChatAdmin(f.Pointer)
}

func (f IFriends) IsClanChatWindowOpenInSteam(SteamID steamIDClanChat) bool {
	return C.SteamFriends_IsClanChatWindowOpenInSteam(f.Pointer)
}

func (f IFriends) OpenClanChatWindowInSteam(SteamID steamIDClanChat) bool {
	return C.SteamFriends_OpenClanChatWindowInSteam(f.Pointer)
}

func (f IFriends) CloseClanChatWindowInSteam(SteamID steamIDClanChat) bool {
	return C.SteamFriends_CloseClanChatWindowInSteam(f.Pointer)
}
*/

// SetListenForFriendsMessages has no documentation. But does has "this is so
// you can show P2P chats inline in the game" near the declaration of this
// function.
func (f IFriends) SetListenForFriendsMessages(interceptEnabled bool) bool {
	return bool(C.SteamFriends_SetListenForFriendsMessages(f.Pointer, C._Bool(interceptEnabled)))
}

/*
func (f IFriends) ReplyToFriendMessage(SteamID steamIDFriend, string pchMsgToSend) bool {
	return C.SteamFriends_ReplyToFriendMessage(f.Pointer)
}

func (f IFriends) GetFriendMessage(SteamID steamIDFriend, int iMessageID, void *pvData, int cubData, EChatEntryType *peChatEntryType) int {
	return C.SteamFriends_GetFriendMessage(f.Pointer)
}

func (f IFriends) GetFollowerCount(SteamID steamID) SteamAPICall_t {
	return C.SteamFriends_GetFollowerCount(f.Pointer)
}

func (f IFriends) IsFollowing(SteamID steamID) SteamAPICall_t {
	return C.SteamFriends_IsFollowing(f.Pointer)
}
*/

// EnumerateFollowingList has no documentation beside "following apis"
func (f IFriends) EnumerateFollowingList(unStartIndex uint32) APICall {
	return APICall(C.SteamFriends_EnumerateFollowingList(f.Pointer, C.uint(unStartIndex)))
}
