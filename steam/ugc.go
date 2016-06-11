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

// EUserUGCList is an enum for different lists of published UGC for a user. If
// the current logged in user is different than the specified user, then some
// options may not be allowed.
type EUserUGCList int32

// the EUserUGCList enum values
const (
	EUserUGCListPublished EUserUGCList = iota
	EUserUGCListVotedOn
	EUserUGCListVotedUp
	EUserUGCListVotedDown
	EUserUGCListWillVoteLater
	EUserUGCListFavorited
	EUserUGCListSubscribed
	EUserUGCListUsedOrPlayed
	EUserUGCListFollowed
)

// Matching UGC types for queries
const (
	EUGCMatchingUGCTypeItems              EUGCMatchingUGCType = 0 // both mtx items and ready-to-use items
	EUGCMatchingUGCTypeItemsMtx           EUGCMatchingUGCType = 1
	EUGCMatchingUGCTypeItemsReadyToUse    EUGCMatchingUGCType = 2
	EUGCMatchingUGCTypeCollections        EUGCMatchingUGCType = 3
	EUGCMatchingUGCTypeArtwork            EUGCMatchingUGCType = 4
	EUGCMatchingUGCTypeVideos             EUGCMatchingUGCType = 5
	EUGCMatchingUGCTypeScreenshots        EUGCMatchingUGCType = 6
	EUGCMatchingUGCTypeAllGuides          EUGCMatchingUGCType = 7 // both web guides and integrated guides
	EUGCMatchingUGCTypeWebGuides          EUGCMatchingUGCType = 8
	EUGCMatchingUGCTypeIntegratedGuides   EUGCMatchingUGCType = 9
	EUGCMatchingUGCTypeUsableInGame       EUGCMatchingUGCType = 10 // ready-to-use items and integrated guides
	EUGCMatchingUGCTypeControllerBindings EUGCMatchingUGCType = 11
	EUGCMatchingUGCTypeGameManagedItems   EUGCMatchingUGCType = 12 // game managed items (not managed by users)
)

// Sort order for user published UGC lists (defaults to creation order descending)
const (
	EUserUGCListSortOrderCreationOrderDesc EUserUGCListSortOrder = iota
	EUserUGCListSortOrderCreationOrderAsc
	EUserUGCListSortOrderTitleAsc
	EUserUGCListSortOrderLastUpdatedDesc
	EUserUGCListSortOrderSubscriptionDateDesc
	EUserUGCListSortOrderVoteScoreDesc
	EUserUGCListSortOrderForModeration
)

// Combination of sorting and filtering for queries across all UGC
const (
	EUGCQueryRankedByVote                                  EUGCQuery = 0
	EUGCQueryRankedByPublicationDate                       EUGCQuery = 1
	EUGCQueryAcceptedForGameRankedByAcceptanceDate         EUGCQuery = 2
	EUGCQueryRankedByTrend                                 EUGCQuery = 3
	EUGCQueryFavoritedByFriendsRankedByPublicationDate     EUGCQuery = 4
	EUGCQueryCreatedByFriendsRankedByPublicationDate       EUGCQuery = 5
	EUGCQueryRankedByNumTimesReported                      EUGCQuery = 6
	EUGCQueryCreatedByFollowedUsersRankedByPublicationDate EUGCQuery = 7
	EUGCQueryNotYetRated                                   EUGCQuery = 8
	EUGCQueryRankedByTotalVotesAsc                         EUGCQuery = 9
	EUGCQueryRankedByVotesUp                               EUGCQuery = 10
	EUGCQueryRankedByTextSearch                            EUGCQuery = 11
	EUGCQueryRankedByTotalUniqueSubscriptions              EUGCQuery = 12
)

// EItemStatistic has no documentation. contact valve :)
type EItemStatistic int32

// the enum values of EItemStatistic
const (
	EItemStatisticNumSubscriptions       EItemStatistic = 0
	EItemStatisticNumFavorites           EItemStatistic = 1
	EItemStatisticNumFollowers           EItemStatistic = 2
	EItemStatisticNumUniqueSubscriptions EItemStatistic = 3
	EItemStatisticNumUniqueFavorites     EItemStatistic = 4
	EItemStatisticNumUniqueFollowers     EItemStatistic = 5
	EItemStatisticNumUniqueWebsiteViews  EItemStatistic = 6
	EItemStatisticReportScore            EItemStatistic = 7
)

// the enum values of EItemUpdateStatus
const (
	EItemUpdateStatusInvalid              EItemUpdateStatus = 0 // The item update handle was invalid, job might be finished, listen too SubmitItemUpdateResult_t
	EItemUpdateStatusPreparingConfig      EItemUpdateStatus = 1 // The item update is processing configuration data
	EItemUpdateStatusPreparingContent     EItemUpdateStatus = 2 // The item update is reading and processing content files
	EItemUpdateStatusUploadingContent     EItemUpdateStatus = 3 // The item update is uploading content changes to Steam
	EItemUpdateStatusUploadingPreviewFile EItemUpdateStatus = 4 // The item update is uploading new preview file image
	EItemUpdateStatusCommittingChanges    EItemUpdateStatus = 5 // The item update is committing all changes
)

// EResult is an enum for the general result codes.
type EResult int

// UGCQueryHandle is a handle for UGC queries.
type UGCQueryHandle uint64

// AccountID has no documentation. Contact Valve :)
type AccountID uint32

// EUGCMatchingUGCType is an enum for matching UGC types for queries.
type EUGCMatchingUGCType int32

// EUserUGCListSortOrder is an enum to sort order for user published UGC lists
// (defaults to creation order descending).
type EUserUGCListSortOrder int32

// AppID has no documentation. Contact Valve :)
type AppID uint32

// EUGCQuery is an enum for combination of sorting and filtering for queries
// across all UGC.
type EUGCQuery int32

// PublishedFileID has no documentation. Contact Valve :)
type PublishedFileID uint64

// EWorkshopFileType has no documentation. Contact Valve :)
type EWorkshopFileType int32

// ERemoteStoragePublishedFileVisibility has no documentation. Contact Valve :)
type ERemoteStoragePublishedFileVisibility int32

// EItemUpdateStatus has no documentation. Contact Valve :)
type EItemUpdateStatus int32

// UGCUpdateHandle has no documentation. Contact Valve :)
type UGCUpdateHandle uint64

// UGCHandle is a handle to a piece of user generated content.
type UGCHandle uint64

// ISteamUGC is a interface to Steam UGC API.
type ISteamUGC struct {
	unsafe.Pointer
}

// UGCDetails holds the details for a single published file/UGC.
type UGCDetails struct {
	nPublishedFileID PublishedFileID
	// The result of the operation.
	eResult EResult
	// Type of the file
	eFileType EWorkshopFileType
	// ID of the app that created this file.
	nCreatorAppID AppID
	// ID of the app that will consume this file.
	nConsumerAppID AppID
	// title of document
	rgchTitle string
	// description of document
	rgchDescription string
	// Steam ID of the user who created this content.
	ulSteamIDOwner uint64
	// time when the published file was created
	rtimeCreated uint32
	// time when the published file was last updated
	rtimeUpdated uint32
	// time when the user added the published file to their list (not always applicable)
	rtimeAddedToUserList uint32
	// visibility
	eVisibility ERemoteStoragePublishedFileVisibility
	// whether the file was banned
	bBanned bool
	// developer has specifically flagged this item as accepted in the Workshop
	bAcceptedForUse bool
	// whether the list of tags was too long to be returned in the provided buffer
	bTagsTruncated bool
	// comma separated list of all tags associated with this file
	rgchTags string
	// file/url information
	// The handle of the primary file
	hFile UGCHandle
	// The handle of the preview file
	hPreviewFile UGCHandle
	// The cloud filename of the primary file
	pchFileName string
	// Size of the primary file
	nFileSize int32
	// Size of the preview file
	nPreviewFileSize int32
	// URL (for a video or a website)
	rgchURL string
	// voting information
	// number of votes up
	unVotesUp uint32
	// number of votes down
	unVotesDown uint32
	// calculated score
	flScore float32
	// collection details
	unNumChildren uint32
	c             C.SteamUGCDetails_t
}

func (d *UGCDetails) refill() {
	d.nPublishedFileID = PublishedFileID(d.c.m_nPublishedFileId)
	d.eResult = EResult(d.c.m_eResult)
	d.eFileType = EWorkshopFileType(d.c.m_eFileType)
	d.nCreatorAppID = AppID(d.c.m_nCreatorAppID)
	d.nConsumerAppID = AppID(d.c.m_nConsumerAppID)
	d.rgchTitle = C.GoString(&d.c.m_rgchTitle[0])
	d.rgchDescription = C.GoString(&d.c.m_rgchDescription[0])
	d.ulSteamIDOwner = uint64(d.c.m_ulSteamIDOwner)
	d.rtimeCreated = uint32(d.c.m_rtimeCreated)
	d.rtimeUpdated = uint32(d.c.m_rtimeUpdated)
	d.rtimeAddedToUserList = uint32(d.c.m_rtimeAddedToUserList)
	d.eVisibility = ERemoteStoragePublishedFileVisibility(d.c.m_eVisibility)
	d.bBanned = bool(d.c.m_bBanned)
	d.bAcceptedForUse = bool(d.c.m_bAcceptedForUse)
	d.bTagsTruncated = bool(d.c.m_bTagsTruncated)
	d.rgchTags = C.GoString(&d.c.m_rgchTags[0])
	d.hFile = UGCHandle(d.c.m_hFile)
	d.hPreviewFile = UGCHandle(d.c.m_hPreviewFile)
	d.pchFileName = C.GoString(&d.c.m_pchFileName[0])
	d.nFileSize = int32(d.c.m_nFileSize)
	d.nPreviewFileSize = int32(d.c.m_nPreviewFileSize)
	d.rgchURL = C.GoString(&d.c.m_rgchURL[0])
	d.unVotesUp = uint32(d.c.m_unVotesUp)
	d.unVotesDown = uint32(d.c.m_unVotesDown)
	d.flScore = float32(d.c.m_flScore)
	d.unNumChildren = uint32(d.c.m_unNumChildren)
}

// UGC returns the ISteamUGC.
func UGC() ISteamUGC {
	return ISteamUGC{C.CSteamUGC()}
}

// CreateQueryUserUGCRequest creates a query UGC associated with a user. Creator app id or consumer app id must be valid and be set to the current running app. unPage should start at 1.
func (u ISteamUGC) CreateQueryUserUGCRequest(unAccountID AccountID, eListType EUserUGCList, eMatchingUGCType EUGCMatchingUGCType, eSortOrder EUserUGCListSortOrder, nCreatorAppID AppID, nConsumerAppID AppID, unPage uint32) UGCQueryHandle {
	return UGCQueryHandle(C.SteamUGC_CreateQueryUserUGCRequest(u.Pointer, C.AccountID_t(unAccountID), C.EUserUGCList(eListType), C.EUGCMatchingUGCType(eMatchingUGCType), C.EUserUGCListSortOrder(eSortOrder), C.AppId_t(nCreatorAppID), C.AppId_t(nConsumerAppID), C.uint(unPage)))
}

// CreateQueryAllUGCRequest creates a query for all matching UGC. Creator app id or consumer app id must be valid and be set to the current running app. unPage should start at 1.
func (u ISteamUGC) CreateQueryAllUGCRequest(eQueryType EUGCQuery, eMatchingeMatchingUGCTypeFileType EUGCMatchingUGCType, nCreatorAppID AppID, nConsumerAppID AppID, unPage uint32) UGCQueryHandle {
	return UGCQueryHandle(C.SteamUGC_CreateQueryAllUGCRequest(u.Pointer, C.EUGCQuery(eQueryType), C.EUGCMatchingUGCType(eMatchingeMatchingUGCTypeFileType), C.AppId_t(nCreatorAppID), C.AppId_t(nConsumerAppID), C.uint(unPage)))
}

// CreateQueryUGCDetailsRequest creates a query for the details of the given published file ids (the RequestUGCDetails call is deprecated and replaced with this)
func (u ISteamUGC) CreateQueryUGCDetailsRequest(pvecPublishedFileID *PublishedFileID, unNumPublishedFileIDs uint32) UGCQueryHandle {
	return UGCQueryHandle(C.SteamUGC_CreateQueryUGCDetailsRequest(u.Pointer, (*C.PublishedFileId_t)(pvecPublishedFileID), C.uint(unNumPublishedFileIDs)))
}

// SendQueryUGCRequest sends the query to Steam
func (u ISteamUGC) SendQueryUGCRequest(handle UGCQueryHandle) APICall {
	return APICall(C.SteamUGC_SendQueryUGCRequest(u.Pointer, C.UGCQueryHandle_t(handle)))
}

// GetQueryUGCResult has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCResult(handle UGCQueryHandle, index uint32, pDetails *UGCDetails) bool {
	ret := bool(C.SteamUGC_GetQueryUGCResult(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), &pDetails.c))
	pDetails.refill()
	return ret
}

// GetQueryUGCPreviewURL has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCPreviewURL(handle UGCQueryHandle, index uint32, pchURL string, cchURLSize uint32) bool {
	cpchURL := C.CString(pchURL)
	defer C.free(unsafe.Pointer(cpchURL))
	return bool(C.SteamUGC_GetQueryUGCPreviewURL(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), cpchURL, C.uint(cchURLSize)))
}

// GetQueryUGCMetadata has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCMetadata(handle UGCQueryHandle, index uint32, pchMetadata string, cchMetadatasize uint32) bool {
	cpchMetadata := C.CString(pchMetadata)
	defer C.free(unsafe.Pointer(cpchMetadata))
	return bool(C.SteamUGC_GetQueryUGCMetadata(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), cpchMetadata, C.uint(cchMetadatasize)))
}

// GetQueryUGCChildren has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCChildren(handle UGCQueryHandle, index uint32, pvecPublishedFileID *PublishedFileID, cMaxEntries uint32) bool {
	return bool(C.SteamUGC_GetQueryUGCChildren(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), (*C.PublishedFileId_t)(pvecPublishedFileID), C.uint(cMaxEntries)))
}

// GetQueryUGCStatistic has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCStatistic(handle UGCQueryHandle, index uint32, eStatType EItemStatistic, pStatValue *uint32) bool {
	return bool(C.SteamUGC_GetQueryUGCStatistic(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), C.EItemStatistic(eStatType), (*C.uint)(pStatValue)))
}

// GetQueryUGCNumAdditionalPreviews has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCNumAdditionalPreviews(handle UGCQueryHandle, index uint32) uint32 {
	return uint32(C.SteamUGC_GetQueryUGCNumAdditionalPreviews(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index)))
}

// GetQueryUGCAdditionalPreview has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCAdditionalPreview(handle UGCQueryHandle, index uint32, previewIndex uint32, pchURLOrVideoID string, cchURLSize uint32, pbIsImage *bool) bool {
	cpchURLOrVideoID := C.CString(pchURLOrVideoID)
	defer C.free(unsafe.Pointer(cpchURLOrVideoID))
	return bool(C.SteamUGC_GetQueryUGCAdditionalPreview(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), C.uint(previewIndex), cpchURLOrVideoID, C.uint(cchURLSize), (*C._Bool)(pbIsImage)))
}

// GetQueryUGCNumKeyValueTags has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCNumKeyValueTags(handle UGCQueryHandle, index uint32) uint32 {
	return uint32(C.SteamUGC_GetQueryUGCNumKeyValueTags(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index)))
}

// GetQueryUGCKeyValueTag has no documentation but it does have this comment above a function block  "Retrieves an individual result after receiving the callback for querying UGC.""
func (u ISteamUGC) GetQueryUGCKeyValueTag(handle UGCQueryHandle, index uint32, keyValueTagIndex uint32, pchKey string, cchKeySize uint32, pchValue string, cchValueSize uint32) bool {
	cpchKey := C.CString(pchKey)
	defer C.free(unsafe.Pointer(cpchKey))
	cpchValue := C.CString(pchValue)
	defer C.free(unsafe.Pointer(cpchValue))
	return bool(C.SteamUGC_GetQueryUGCKeyValueTag(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), C.uint(keyValueTagIndex), cpchKey, C.uint(cchKeySize), cpchValue, C.uint(cchValueSize)))
}

// ReleaseQueryUGCRequest releases the request to free up memory, after retrieving results
func (u ISteamUGC) ReleaseQueryUGCRequest(handle UGCQueryHandle) bool {
	return bool(C.SteamUGC_ReleaseQueryUGCRequest(u.Pointer, C.UGCQueryHandle_t(handle)))
}

// AddRequiredTag has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) AddRequiredTag(handle UGCQueryHandle, pTagName string) bool {
	cpTagName := C.CString(pTagName)
	defer C.free(unsafe.Pointer(cpTagName))
	return bool(C.SteamUGC_AddRequiredTag(u.Pointer, C.UGCQueryHandle_t(handle), cpTagName))
}

// AddExcludedTag has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) AddExcludedTag(handle UGCQueryHandle, pTagName string) bool {
	cpTagName := C.CString(pTagName)
	defer C.free(unsafe.Pointer(cpTagName))
	return bool(C.SteamUGC_AddExcludedTag(u.Pointer, C.UGCQueryHandle_t(handle), cpTagName))
}

// SetReturnKeyValueTags has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) SetReturnKeyValueTags(handle UGCQueryHandle, bReturnKeyValueTags bool) bool {
	return bool(C.SteamUGC_SetReturnKeyValueTags(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnKeyValueTags)))
}

// SetReturnLongDescription has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) SetReturnLongDescription(handle UGCQueryHandle, bReturnLongDescription bool) bool {
	return bool(C.SteamUGC_SetReturnLongDescription(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnLongDescription)))
}

// SetReturnMetadata has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) SetReturnMetadata(handle UGCQueryHandle, bReturnMetadata bool) bool {
	return bool(C.SteamUGC_SetReturnMetadata(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnMetadata)))
}

// SetReturnChildren has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) SetReturnChildren(handle UGCQueryHandle, bReturnChildren bool) bool {
	return bool(C.SteamUGC_SetReturnChildren(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnChildren)))
}

// SetReturnAdditionalPreviews has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) SetReturnAdditionalPreviews(handle UGCQueryHandle, bReturnAdditionalPreviews bool) bool {
	return bool(C.SteamUGC_SetReturnAdditionalPreviews(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnAdditionalPreviews)))
}

// SetReturnTotalOnly has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) SetReturnTotalOnly(handle UGCQueryHandle, bReturnTotalOnly bool) bool {
	return bool(C.SteamUGC_SetReturnTotalOnly(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnTotalOnly)))
}

// SetLanguage has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) SetLanguage(handle UGCQueryHandle, pchLanguage string) bool {
	cpchLanguage := C.CString(pchLanguage)
	defer C.free(unsafe.Pointer(cpchLanguage))
	return bool(C.SteamUGC_SetLanguage(u.Pointer, C.UGCQueryHandle_t(handle), cpchLanguage))
}

// SetAllowCachedResponse has no documentation but it has this comment above a function block:  "Options to set for querying UGC"
func (u ISteamUGC) SetAllowCachedResponse(handle UGCQueryHandle, unMaxAgeSeconds uint32) bool {
	return bool(C.SteamUGC_SetAllowCachedResponse(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(unMaxAgeSeconds)))
}

// SetCloudFileNameFilter sets the options only for querying user UGC
func (u ISteamUGC) SetCloudFileNameFilter(handle UGCQueryHandle, pMatchCloudFileName string) bool {
	cpMatchCloudFileName := C.CString(pMatchCloudFileName)
	defer C.free(unsafe.Pointer(cpMatchCloudFileName))
	return bool(C.SteamUGC_SetCloudFileNameFilter(u.Pointer, C.UGCQueryHandle_t(handle), cpMatchCloudFileName))
}

// SetMatchAnyTag has no documentation but has this comment above a function block:  "Options only for querying all UGC"
func (u ISteamUGC) SetMatchAnyTag(handle UGCQueryHandle, bMatchAnyTag bool) bool {
	return bool(C.SteamUGC_SetMatchAnyTag(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bMatchAnyTag)))
}

// SetSearchText has no documentation but has this comment above a function block:  "Options only for querying all UGC"
func (u ISteamUGC) SetSearchText(handle UGCQueryHandle, pSearchText string) bool {
	cpSearchText := C.CString(pSearchText)
	defer C.free(unsafe.Pointer(cpSearchText))
	return bool(C.SteamUGC_SetSearchText(u.Pointer, C.UGCQueryHandle_t(handle), cpSearchText))
}

// SetRankedByTrendDays has no documentation but has this comment above a function block:  "Options only for querying all UGC"
func (u ISteamUGC) SetRankedByTrendDays(handle UGCQueryHandle, unDays uint32) bool {
	return bool(C.SteamUGC_SetRankedByTrendDays(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(unDays)))
}

// AddRequiredKeyValueTag has no documentation but has this comment above a function block:  "Options only for querying all UGC"
func (u ISteamUGC) AddRequiredKeyValueTag(handle UGCQueryHandle, pKey string, pValue string) bool {
	cpKey := C.CString(pKey)
	defer C.free(unsafe.Pointer(cpKey))
	cpValue := C.CString(pValue)
	defer C.free(unsafe.Pointer(cpValue))
	return bool(C.SteamUGC_AddRequiredKeyValueTag(u.Pointer, C.UGCQueryHandle_t(handle), cpKey, cpValue))
}

// If it's deprecated we really shouldn't expose it.
// RequestUGCDetails is deprecated - Use CreateQueryUGCDetailsRequest call above instead!
//func (u ISteamUGC) RequestUGCDetails(nPublishedFileID PublishedFileId, unMaxAgeSeconds uint32) SteamAPICall {
//	return SteamAPICall(C.SteamUGC_RequestUGCDetails(u.Pointer, C.PublishedFileId_t(nPublishedFileID), C.uint(unMaxAgeSeconds)))
//}

// Steam Workshop Creator API

// CreateItem creates a new item for this app with no content attached yet
func (u ISteamUGC) CreateItem(nConsumerAppID AppID, eFileType EWorkshopFileType) APICall {
	return APICall(C.SteamUGC_CreateItem(u.Pointer, C.AppId_t(nConsumerAppID), C.EWorkshopFileType(eFileType)))
}

// StartItemUpdate starts an UGC item update. Set changed properties before committing update with CommitItemUpdate()
func (u ISteamUGC) StartItemUpdate(nConsumerAppID AppID, nPublishedFileID PublishedFileID) UGCUpdateHandle {
	return UGCUpdateHandle(C.SteamUGC_StartItemUpdate(u.Pointer, C.AppId_t(nConsumerAppID), C.PublishedFileId_t(nPublishedFileID)))
}

// SetItemTitle change the title of an UGC item
func (u ISteamUGC) SetItemTitle(handle UGCUpdateHandle, pchTitle string) bool {
	cpchTitle := C.CString(pchTitle)
	defer C.free(unsafe.Pointer(cpchTitle))
	return bool(C.SteamUGC_SetItemTitle(u.Pointer, C.UGCUpdateHandle_t(handle), cpchTitle))
}

// SetItemDescription change the description of an UGC item
func (u ISteamUGC) SetItemDescription(handle UGCUpdateHandle, pchDescription string) bool {
	cpchDescription := C.CString(pchDescription)
	defer C.free(unsafe.Pointer(cpchDescription))
	return bool(C.SteamUGC_SetItemDescription(u.Pointer, C.UGCUpdateHandle_t(handle), cpchDescription))
}

// SetItemUpdateLanguage specify the language of the title or description that will be set
func (u ISteamUGC) SetItemUpdateLanguage(handle UGCUpdateHandle, pchLanguage string) bool {
	cpchLanguage := C.CString(pchLanguage)
	defer C.free(unsafe.Pointer(cpchLanguage))
	return bool(C.SteamUGC_SetItemUpdateLanguage(u.Pointer, C.UGCUpdateHandle_t(handle), cpchLanguage))
}

// SetItemMetadata change the metadata of an UGC item (max = k_cchDeveloperMetadataMax)
func (u ISteamUGC) SetItemMetadata(handle UGCUpdateHandle, pchMetaData string) bool {
	cpchMetaData := C.CString(pchMetaData)
	defer C.free(unsafe.Pointer(cpchMetaData))
	return bool(C.SteamUGC_SetItemMetadata(u.Pointer, C.UGCUpdateHandle_t(handle), cpchMetaData))
}

// SetItemVisibility change the visibility of an UGC item
func (u ISteamUGC) SetItemVisibility(handle UGCUpdateHandle, eVisibility ERemoteStoragePublishedFileVisibility) bool {
	return bool(C.SteamUGC_SetItemVisibility(u.Pointer, C.UGCUpdateHandle_t(handle), C.ERemoteStoragePublishedFileVisibility(eVisibility)))
}

// SetItemTags changes the tags of an UGC item.
// WARNING: untested and might cause segfault if steam expected
// SteamParamStringArray_t to not be collected and Go decides that we don't need
// it anymore
func (u ISteamUGC) SetItemTags(handle UGCUpdateHandle, pTags []string) bool {
	var cstrings []*C.char
	for _, s := range pTags {
		c := C.CString(s)
		cstrings = append(cstrings, c)
		defer C.free(unsafe.Pointer(c))
	}
	sa := C.SteamParamStringArray_t{
		m_ppStrings:   &cstrings[0],
		m_nNumStrings: C.int32_t(len(pTags)),
	}
	return bool(C.SteamUGC_SetItemTags(u.Pointer, C.UGCUpdateHandle_t(handle), &sa))
}

// SetItemContent update item content from this local folder
func (u ISteamUGC) SetItemContent(handle UGCUpdateHandle, pszContentFolder string) bool {
	cpszContentFolder := C.CString(pszContentFolder)
	defer C.free(unsafe.Pointer(cpszContentFolder))
	return bool(C.SteamUGC_SetItemContent(u.Pointer, C.UGCUpdateHandle_t(handle), cpszContentFolder))
}

// SetItemPreview  change preview image file for this item. pszPreviewFile points to local image file, which must be under 1MB in size
func (u ISteamUGC) SetItemPreview(handle UGCUpdateHandle, pszPreviewFile string) bool {
	cpszPreviewFile := C.CString(pszPreviewFile)
	defer C.free(unsafe.Pointer(cpszPreviewFile))
	return bool(C.SteamUGC_SetItemPreview(u.Pointer, C.UGCUpdateHandle_t(handle), cpszPreviewFile))
}

// RemoveItemKeyValueTags remove any existing key-value tags with the specified key
func (u ISteamUGC) RemoveItemKeyValueTags(handle UGCUpdateHandle, pchKey string) bool {
	cpchKey := C.CString(pchKey)
	defer C.free(unsafe.Pointer(cpchKey))
	return bool(C.SteamUGC_RemoveItemKeyValueTags(u.Pointer, C.UGCUpdateHandle_t(handle), cpchKey))
}

// AddItemKeyValueTag add new key-value tags for the item. Note that there can be multiple values for a tag.
func (u ISteamUGC) AddItemKeyValueTag(handle UGCUpdateHandle, pchKey string, pchValue string) bool {
	cpchKey := C.CString(pchKey)
	defer C.free(unsafe.Pointer(cpchKey))
	cpchValue := C.CString(pchValue)
	defer C.free(unsafe.Pointer(cpchValue))
	return bool(C.SteamUGC_AddItemKeyValueTag(u.Pointer, C.UGCUpdateHandle_t(handle), cpchKey, cpchValue))
}

// SubmitItemUpdate commit update process started with StartItemUpdate()
func (u ISteamUGC) SubmitItemUpdate(handle UGCUpdateHandle, pchChangeNote string) APICall {
	cpchChangeNote := C.CString(pchChangeNote)
	defer C.free(unsafe.Pointer(cpchChangeNote))
	return APICall(C.SteamUGC_SubmitItemUpdate(u.Pointer, C.UGCUpdateHandle_t(handle), cpchChangeNote))
}

// GetItemUpdateProgress has no documentation. Contact Valve :)
func (u ISteamUGC) GetItemUpdateProgress(handle UGCUpdateHandle, punBytesProcessed *uint64, punBytesTotal *uint64) EItemUpdateStatus {
	return EItemUpdateStatus(C.SteamUGC_GetItemUpdateProgress(u.Pointer, C.UGCUpdateHandle_t(handle), (*C.uint64_t)(punBytesProcessed), (*C.uint64_t)(punBytesTotal)))
}

// SetUserItemVote has no documentation, but it does have this comment above the function block "Steam Workshop Consumer API"
func (u ISteamUGC) SetUserItemVote(nPublishedFileID PublishedFileID, bVoteUp bool) APICall {
	return APICall(C.SteamUGC_SetUserItemVote(u.Pointer, C.PublishedFileId_t(nPublishedFileID), C._Bool(bVoteUp)))
}

// GetUserItemVote has no documentation, but it does have this comment above the function block "Steam Workshop Consumer API"
func (u ISteamUGC) GetUserItemVote(nPublishedFileID PublishedFileID) APICall {
	return APICall(C.SteamUGC_GetUserItemVote(u.Pointer, C.PublishedFileId_t(nPublishedFileID)))
}

// AddItemToFavorites has no documentation, but it does have this comment above the function block "Steam Workshop Consumer API"
func (u ISteamUGC) AddItemToFavorites(nAppID AppID, nPublishedFileID PublishedFileID) APICall {
	return APICall(C.SteamUGC_AddItemToFavorites(u.Pointer, C.AppId_t(nAppID), C.PublishedFileId_t(nPublishedFileID)))
}

// RemoveItemFromFavorites has no documentation, but it does have this comment above the function block "Steam Workshop Consumer API"
func (u ISteamUGC) RemoveItemFromFavorites(nAppID AppID, nPublishedFileID PublishedFileID) APICall {
	return APICall(C.SteamUGC_RemoveItemFromFavorites(u.Pointer, C.AppId_t(nAppID), C.PublishedFileId_t(nPublishedFileID)))
}

// SubscribeItem subscribes to this item, will be installed ASAP
func (u ISteamUGC) SubscribeItem(nPublishedFileID PublishedFileID) APICall {
	return APICall(C.SteamUGC_SubscribeItem(u.Pointer, C.PublishedFileId_t(nPublishedFileID)))
}

// UnsubscribeItem unsubscribes from this item, will be uninstalled after game quits
func (u ISteamUGC) UnsubscribeItem(nPublishedFileID PublishedFileID) APICall {
	return APICall(C.SteamUGC_UnsubscribeItem(u.Pointer, C.PublishedFileId_t(nPublishedFileID)))
}

// GetNumSubscribedItems returns the number of subscribed items
func (u ISteamUGC) GetNumSubscribedItems() uint32 {
	return uint32(C.SteamUGC_GetNumSubscribedItems(u.Pointer))
}

// GetSubscribedItems returns all subscribed item PublishFileIDs
func (u ISteamUGC) GetSubscribedItems(pvecPublishedFileID *PublishedFileID, cMaxEntries uint32) uint32 {
	return uint32(C.SteamUGC_GetSubscribedItems(u.Pointer, (*C.PublishedFileId_t)(pvecPublishedFileID), C.uint(cMaxEntries)))
}

// GetItemState returns the EItemState flags about item on this client
func (u ISteamUGC) GetItemState(nPublishedFileID PublishedFileID) uint32 {
	return uint32(C.SteamUGC_GetItemState(u.Pointer, C.PublishedFileId_t(nPublishedFileID)))
}

// GetItemInstallInfo returns the info about currently installed content on disc for items that have k_EItemStateInstalled set
// if k_EItemStateLegacyItem is set, pchFolder contains the path to the legacy file itself (not a folder)
func (u ISteamUGC) GetItemInstallInfo(nPublishedFileID PublishedFileID, punSizeOnDisk *uint64, pchFolder string, cchFolderSize uint32, punTimeStamp *uint32) bool {
	cpchFolder := C.CString(pchFolder)
	defer C.free(unsafe.Pointer(cpchFolder))
	return bool(C.SteamUGC_GetItemInstallInfo(u.Pointer, C.PublishedFileId_t(nPublishedFileID), (*C.uint64_t)(punSizeOnDisk), cpchFolder, C.uint(cchFolderSize), (*C.uint)(punTimeStamp)))
}

// GetItemDownloadInfo returns the info about pending update for items that have k_EItemStateNeedsUpdate set. punBytesTotal will be valid after download started once
func (u ISteamUGC) GetItemDownloadInfo(nPublishedFileID PublishedFileID, punBytesDownloaded *uint64, punBytesTotal *uint64) bool {
	return bool(C.SteamUGC_GetItemDownloadInfo(u.Pointer, C.PublishedFileId_t(nPublishedFileID), (*C.uint64_t)(punBytesDownloaded), (*C.uint64_t)(punBytesTotal)))
}

// DownloadItem downloads new or update already installed item. If function returns true, wait for DownloadItemResult_t. If the item is already installed,
// then files on disk should not be used until callback received. If item is not subscribed to, it will be cached for some time.
// If bHighPriority is set, any other item download will be suspended and this item downloaded ASAP.
func (u ISteamUGC) DownloadItem(nPublishedFileID PublishedFileID, bHighPriority bool) bool {
	return bool(C.SteamUGC_DownloadItem(u.Pointer, C.PublishedFileId_t(nPublishedFileID), C._Bool(bHighPriority)))
}
