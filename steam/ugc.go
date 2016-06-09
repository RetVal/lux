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

const (
	EUserUGCList_Published EUserUGCList = iota
	EUserUGCList_VotedOn
	EUserUGCList_VotedUp
	EUserUGCList_VotedDown
	EUserUGCList_WillVoteLater
	EUserUGCList_Favorited
	EUserUGCList_Subscribed
	EUserUGCList_UsedOrPlayed
	EUserUGCList_Followed
)

// Matching UGC types for queries
const (
	EUGCMatchingUGCType_Items              EUGCMatchingUGCType = 0 // both mtx items and ready-to-use items
	EUGCMatchingUGCType_Items_Mtx          EUGCMatchingUGCType = 1
	EUGCMatchingUGCType_Items_ReadyToUse   EUGCMatchingUGCType = 2
	EUGCMatchingUGCType_Collections        EUGCMatchingUGCType = 3
	EUGCMatchingUGCType_Artwork            EUGCMatchingUGCType = 4
	EUGCMatchingUGCType_Videos             EUGCMatchingUGCType = 5
	EUGCMatchingUGCType_Screenshots        EUGCMatchingUGCType = 6
	EUGCMatchingUGCType_AllGuides          EUGCMatchingUGCType = 7 // both web guides and integrated guides
	EUGCMatchingUGCType_WebGuides          EUGCMatchingUGCType = 8
	EUGCMatchingUGCType_IntegratedGuides   EUGCMatchingUGCType = 9
	EUGCMatchingUGCType_UsableInGame       EUGCMatchingUGCType = 10 // ready-to-use items and integrated guides
	EUGCMatchingUGCType_ControllerBindings EUGCMatchingUGCType = 11
	EUGCMatchingUGCType_GameManagedItems   EUGCMatchingUGCType = 12 // game managed items (not managed by users)
)

// Sort order for user published UGC lists (defaults to creation order descending)
const (
	EUserUGCListSortOrder_CreationOrderDesc EUserUGCListSortOrder = iota
	EUserUGCListSortOrder_CreationOrderAsc
	EUserUGCListSortOrder_TitleAsc
	EUserUGCListSortOrder_LastUpdatedDesc
	EUserUGCListSortOrder_SubscriptionDateDesc
	EUserUGCListSortOrder_VoteScoreDesc
	EUserUGCListSortOrder_ForModeration
)

// Combination of sorting and filtering for queries across all UGC
const (
	EUGCQuery_RankedByVote                                  EUGCQuery = 0
	EUGCQuery_RankedByPublicationDate                       EUGCQuery = 1
	EUGCQuery_AcceptedForGameRankedByAcceptanceDate         EUGCQuery = 2
	EUGCQuery_RankedByTrend                                 EUGCQuery = 3
	EUGCQuery_FavoritedByFriendsRankedByPublicationDate     EUGCQuery = 4
	EUGCQuery_CreatedByFriendsRankedByPublicationDate       EUGCQuery = 5
	EUGCQuery_RankedByNumTimesReported                      EUGCQuery = 6
	EUGCQuery_CreatedByFollowedUsersRankedByPublicationDate EUGCQuery = 7
	EUGCQuery_NotYetRated                                   EUGCQuery = 8
	EUGCQuery_RankedByTotalVotesAsc                         EUGCQuery = 9
	EUGCQuery_RankedByVotesUp                               EUGCQuery = 10
	EUGCQuery_RankedByTextSearch                            EUGCQuery = 11
	EUGCQuery_RankedByTotalUniqueSubscriptions              EUGCQuery = 12
)

const (
	EItemStatistic_NumSubscriptions       EItemStatistic = 0
	EItemStatistic_NumFavorites           EItemStatistic = 1
	EItemStatistic_NumFollowers           EItemStatistic = 2
	EItemStatistic_NumUniqueSubscriptions EItemStatistic = 3
	EItemStatistic_NumUniqueFavorites     EItemStatistic = 4
	EItemStatistic_NumUniqueFollowers     EItemStatistic = 5
	EItemStatistic_NumUniqueWebsiteViews  EItemStatistic = 6
	EItemStatistic_ReportScore            EItemStatistic = 7
)

const (
	EItemUpdateStatusInvalid              EItemUpdateStatus = 0 // The item update handle was invalid, job might be finished, listen too SubmitItemUpdateResult_t
	EItemUpdateStatusPreparingConfig      EItemUpdateStatus = 1 // The item update is processing configuration data
	EItemUpdateStatusPreparingContent     EItemUpdateStatus = 2 // The item update is reading and processing content files
	EItemUpdateStatusUploadingContent     EItemUpdateStatus = 3 // The item update is uploading content changes to Steam
	EItemUpdateStatusUploadingPreviewFile EItemUpdateStatus = 4 // The item update is uploading new preview file image
	EItemUpdateStatusCommittingChanges    EItemUpdateStatus = 5 // The item update is committing all changes
)

type EResult int

type UGCQueryHandle uint64
type AccountID uint32
type EUserUGCList int32
type EUGCMatchingUGCType int32
type EUserUGCListSortOrder int32
type AppId uint32
type EUGCQuery int32
type PublishedFileId uint64
type EItemStatistic int32

type EWorkshopFileType int32
type ERemoteStoragePublishedFileVisibility int32
type EItemUpdateStatus int32

type SteamAPICall uint64
type UGCUpdateHandle uint64
type UGCHandle uint64

type ISteamUGC struct {
	unsafe.Pointer
}

type SteamUGCDetails struct {
	nPublishedFileId PublishedFileId
	// The result of the operation.
	eResult EResult
	// Type of the file
	eFileType EWorkshopFileType
	// ID of the app that created this file.
	nCreatorAppID AppId
	// ID of the app that will consume this file.
	nConsumerAppID AppId
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

func (d *SteamUGCDetails) refill() {
	d.nPublishedFileId = PublishedFileId(d.c.m_nPublishedFileId)
	d.eResult = EResult(d.c.m_eResult)
	d.eFileType = EWorkshopFileType(d.c.m_eFileType)
	d.nCreatorAppID = AppId(d.c.m_nCreatorAppID)
	d.nConsumerAppID = AppId(d.c.m_nConsumerAppID)
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

func UGC() ISteamUGC {
	return ISteamUGC{C.CSteamUGC()}
}

// CreateQueryUserUGCRequest creates a query UGC associated with a user. Creator app id or consumer app id must be valid and be set to the current running app. unPage should start at 1.
func (u ISteamUGC) CreateQueryUserUGCRequest(unAccountID AccountID, eListType EUserUGCList, eMatchingUGCType EUGCMatchingUGCType, eSortOrder EUserUGCListSortOrder, nCreatorAppID AppId, nConsumerAppID AppId, unPage uint32) UGCQueryHandle {
	return UGCQueryHandle(C.SteamUGC_CreateQueryUserUGCRequest(u.Pointer, C.AccountID_t(unAccountID), C.EUserUGCList(eListType), C.EUGCMatchingUGCType(eMatchingUGCType), C.EUserUGCListSortOrder(eSortOrder), C.AppId_t(nCreatorAppID), C.AppId_t(nConsumerAppID), C.uint(unPage)))
}

// CreateQueryAllUGCRequest creates a query for all matching UGC. Creator app id or consumer app id must be valid and be set to the current running app. unPage should start at 1.
func (u ISteamUGC) CreateQueryAllUGCRequest(eQueryType EUGCQuery, eMatchingeMatchingUGCTypeFileType EUGCMatchingUGCType, nCreatorAppID AppId, nConsumerAppID AppId, unPage uint32) UGCQueryHandle {
	return UGCQueryHandle(C.SteamUGC_CreateQueryAllUGCRequest(u.Pointer, C.EUGCQuery(eQueryType), C.EUGCMatchingUGCType(eMatchingeMatchingUGCTypeFileType), C.AppId_t(nCreatorAppID), C.AppId_t(nConsumerAppID), C.uint(unPage)))
}

// CreateQueryUGCDetailsRequest creates a query for the details of the given published file ids (the RequestUGCDetails call is deprecated and replaced with this)
func (u ISteamUGC) CreateQueryUGCDetailsRequest(pvecPublishedFileID *PublishedFileId, unNumPublishedFileIDs uint32) UGCQueryHandle {
	return UGCQueryHandle(C.SteamUGC_CreateQueryUGCDetailsRequest(u.Pointer, (*C.PublishedFileId_t)(pvecPublishedFileID), C.uint(unNumPublishedFileIDs)))
}

// SendQueryUGCRequest sends the query to Steam
func (u ISteamUGC) SendQueryUGCRequest(handle UGCQueryHandle) SteamAPICall {
	return SteamAPICall(C.SteamUGC_SendQueryUGCRequest(u.Pointer, C.UGCQueryHandle_t(handle)))
}

// Retrieves an individual result after receiving the callback for querying UGC

func (u ISteamUGC) GetQueryUGCResult(handle UGCQueryHandle, index uint32, pDetails *SteamUGCDetails) bool {
	ret := bool(C.SteamUGC_GetQueryUGCResult(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), &pDetails.c))
	pDetails.refill()
	return ret
}

func (u ISteamUGC) GetQueryUGCPreviewURL(handle UGCQueryHandle, index uint32, pchURL string, cchURLSize uint32) bool {
	cpchURL := C.CString(pchURL)
	defer C.free(unsafe.Pointer(cpchURL))
	return bool(C.SteamUGC_GetQueryUGCPreviewURL(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), cpchURL, C.uint(cchURLSize)))
}

func (u ISteamUGC) GetQueryUGCMetadata(handle UGCQueryHandle, index uint32, pchMetadata string, cchMetadatasize uint32) bool {
	cpchMetadata := C.CString(pchMetadata)
	defer C.free(unsafe.Pointer(cpchMetadata))
	return bool(C.SteamUGC_GetQueryUGCMetadata(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), cpchMetadata, C.uint(cchMetadatasize)))
}

func (u ISteamUGC) GetQueryUGCChildren(handle UGCQueryHandle, index uint32, pvecPublishedFileID *PublishedFileId, cMaxEntries uint32) bool {
	return bool(C.SteamUGC_GetQueryUGCChildren(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), (*C.PublishedFileId_t)(pvecPublishedFileID), C.uint(cMaxEntries)))
}

func (u ISteamUGC) GetQueryUGCStatistic(handle UGCQueryHandle, index uint32, eStatType EItemStatistic, pStatValue *uint32) bool {
	return bool(C.SteamUGC_GetQueryUGCStatistic(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), C.EItemStatistic(eStatType), (*C.uint)(pStatValue)))
}

func (u ISteamUGC) GetQueryUGCNumAdditionalPreviews(handle UGCQueryHandle, index uint32) uint32 {
	return uint32(C.SteamUGC_GetQueryUGCNumAdditionalPreviews(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index)))
}

func (u ISteamUGC) GetQueryUGCAdditionalPreview(handle UGCQueryHandle, index uint32, previewIndex uint32, pchURLOrVideoID string, cchURLSize uint32, pbIsImage *bool) bool {
	cpchURLOrVideoID := C.CString(pchURLOrVideoID)
	defer C.free(unsafe.Pointer(cpchURLOrVideoID))
	return bool(C.SteamUGC_GetQueryUGCAdditionalPreview(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index), C.uint(previewIndex), cpchURLOrVideoID, C.uint(cchURLSize), (*C._Bool)(pbIsImage)))
}

func (u ISteamUGC) GetQueryUGCNumKeyValueTags(handle UGCQueryHandle, index uint32) uint32 {
	return uint32(C.SteamUGC_GetQueryUGCNumKeyValueTags(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(index)))
}

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

// Options to set for querying UGC

func (u ISteamUGC) AddRequiredTag(handle UGCQueryHandle, pTagName string) bool {
	cpTagName := C.CString(pTagName)
	defer C.free(unsafe.Pointer(cpTagName))
	return bool(C.SteamUGC_AddRequiredTag(u.Pointer, C.UGCQueryHandle_t(handle), cpTagName))
}

func (u ISteamUGC) AddExcludedTag(handle UGCQueryHandle, pTagName string) bool {
	cpTagName := C.CString(pTagName)
	defer C.free(unsafe.Pointer(cpTagName))
	return bool(C.SteamUGC_AddExcludedTag(u.Pointer, C.UGCQueryHandle_t(handle), cpTagName))
}

func (u ISteamUGC) SetReturnKeyValueTags(handle UGCQueryHandle, bReturnKeyValueTags bool) bool {
	return bool(C.SteamUGC_SetReturnKeyValueTags(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnKeyValueTags)))
}

func (u ISteamUGC) SetReturnLongDescription(handle UGCQueryHandle, bReturnLongDescription bool) bool {
	return bool(C.SteamUGC_SetReturnLongDescription(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnLongDescription)))
}

func (u ISteamUGC) SetReturnMetadata(handle UGCQueryHandle, bReturnMetadata bool) bool {
	return bool(C.SteamUGC_SetReturnMetadata(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnMetadata)))
}

func (u ISteamUGC) SetReturnChildren(handle UGCQueryHandle, bReturnChildren bool) bool {
	return bool(C.SteamUGC_SetReturnChildren(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnChildren)))
}

func (u ISteamUGC) SetReturnAdditionalPreviews(handle UGCQueryHandle, bReturnAdditionalPreviews bool) bool {
	return bool(C.SteamUGC_SetReturnAdditionalPreviews(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnAdditionalPreviews)))
}

func (u ISteamUGC) SetReturnTotalOnly(handle UGCQueryHandle, bReturnTotalOnly bool) bool {
	return bool(C.SteamUGC_SetReturnTotalOnly(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bReturnTotalOnly)))
}

func (u ISteamUGC) SetLanguage(handle UGCQueryHandle, pchLanguage string) bool {
	cpchLanguage := C.CString(pchLanguage)
	defer C.free(unsafe.Pointer(cpchLanguage))
	return bool(C.SteamUGC_SetLanguage(u.Pointer, C.UGCQueryHandle_t(handle), cpchLanguage))
}

func (u ISteamUGC) SetAllowCachedResponse(handle UGCQueryHandle, unMaxAgeSeconds uint32) bool {
	return bool(C.SteamUGC_SetAllowCachedResponse(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(unMaxAgeSeconds)))
}

// SetCloudFileNameFilter sets the options only for querying user UGC
func (u ISteamUGC) SetCloudFileNameFilter(handle UGCQueryHandle, pMatchCloudFileName string) bool {
	cpMatchCloudFileName := C.CString(pMatchCloudFileName)
	defer C.free(unsafe.Pointer(cpMatchCloudFileName))
	return bool(C.SteamUGC_SetCloudFileNameFilter(u.Pointer, C.UGCQueryHandle_t(handle), cpMatchCloudFileName))
}

// Options only for querying all UGC

func (u ISteamUGC) SetMatchAnyTag(handle UGCQueryHandle, bMatchAnyTag bool) bool {
	return bool(C.SteamUGC_SetMatchAnyTag(u.Pointer, C.UGCQueryHandle_t(handle), C._Bool(bMatchAnyTag)))
}

func (u ISteamUGC) SetSearchText(handle UGCQueryHandle, pSearchText string) bool {
	cpSearchText := C.CString(pSearchText)
	defer C.free(unsafe.Pointer(cpSearchText))
	return bool(C.SteamUGC_SetSearchText(u.Pointer, C.UGCQueryHandle_t(handle), cpSearchText))
}

func (u ISteamUGC) SetRankedByTrendDays(handle UGCQueryHandle, unDays uint32) bool {
	return bool(C.SteamUGC_SetRankedByTrendDays(u.Pointer, C.UGCQueryHandle_t(handle), C.uint(unDays)))
}

func (u ISteamUGC) AddRequiredKeyValueTag(handle UGCQueryHandle, pKey string, pValue string) bool {
	cpKey := C.CString(pKey)
	defer C.free(unsafe.Pointer(cpKey))
	cpValue := C.CString(pValue)
	defer C.free(unsafe.Pointer(cpValue))
	return bool(C.SteamUGC_AddRequiredKeyValueTag(u.Pointer, C.UGCQueryHandle_t(handle), cpKey, cpValue))
}

// RequestUGCDetails is deprecated - Use CreateQueryUGCDetailsRequest call above instead!
func (u ISteamUGC) RequestUGCDetails(nPublishedFileID PublishedFileId, unMaxAgeSeconds uint32) SteamAPICall {
	return SteamAPICall(C.SteamUGC_RequestUGCDetails(u.Pointer, C.PublishedFileId_t(nPublishedFileID), C.uint(unMaxAgeSeconds)))
}

// Steam Workshop Creator API

// CreateItem creates a new item for this app with no content attached yet
func (u ISteamUGC) CreateItem(nConsumerAppId AppId, eFileType EWorkshopFileType) SteamAPICall {
	return SteamAPICall(C.SteamUGC_CreateItem(u.Pointer, C.AppId_t(nConsumerAppId), C.EWorkshopFileType(eFileType)))
}

// StartItemUpdate starts an UGC item update. Set changed properties before committing update with CommitItemUpdate()
func (u ISteamUGC) StartItemUpdate(nConsumerAppId AppId, nPublishedFileID PublishedFileId) UGCUpdateHandle {
	return UGCUpdateHandle(C.SteamUGC_StartItemUpdate(u.Pointer, C.AppId_t(nConsumerAppId), C.PublishedFileId_t(nPublishedFileID)))
}

// change the title of an UGC item
func (u ISteamUGC) SetItemTitle(handle UGCUpdateHandle, pchTitle string) bool {
	cpchTitle := C.CString(pchTitle)
	defer C.free(unsafe.Pointer(cpchTitle))
	return bool(C.SteamUGC_SetItemTitle(u.Pointer, C.UGCUpdateHandle_t(handle), cpchTitle))
}

// change the description of an UGC item
func (u ISteamUGC) SetItemDescription(handle UGCUpdateHandle, pchDescription string) bool {
	cpchDescription := C.CString(pchDescription)
	defer C.free(unsafe.Pointer(cpchDescription))
	return bool(C.SteamUGC_SetItemDescription(u.Pointer, C.UGCUpdateHandle_t(handle), cpchDescription))
}

// specify the language of the title or description that will be set
func (u ISteamUGC) SetItemUpdateLanguage(handle UGCUpdateHandle, pchLanguage string) bool {
	cpchLanguage := C.CString(pchLanguage)
	defer C.free(unsafe.Pointer(cpchLanguage))
	return bool(C.SteamUGC_SetItemUpdateLanguage(u.Pointer, C.UGCUpdateHandle_t(handle), cpchLanguage))
}

// change the metadata of an UGC item (max = k_cchDeveloperMetadataMax)
func (u ISteamUGC) SetItemMetadata(handle UGCUpdateHandle, pchMetaData string) bool {
	cpchMetaData := C.CString(pchMetaData)
	defer C.free(unsafe.Pointer(cpchMetaData))
	return bool(C.SteamUGC_SetItemMetadata(u.Pointer, C.UGCUpdateHandle_t(handle), cpchMetaData))
}

// change the visibility of an UGC item
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

// update item content from this local folder
func (u ISteamUGC) SetItemContent(handle UGCUpdateHandle, pszContentFolder string) bool {
	cpszContentFolder := C.CString(pszContentFolder)
	defer C.free(unsafe.Pointer(cpszContentFolder))
	return bool(C.SteamUGC_SetItemContent(u.Pointer, C.UGCUpdateHandle_t(handle), cpszContentFolder))
}

//  change preview image file for this item. pszPreviewFile points to local image file, which must be under 1MB in size
func (u ISteamUGC) SetItemPreview(handle UGCUpdateHandle, pszPreviewFile string) bool {
	cpszPreviewFile := C.CString(pszPreviewFile)
	defer C.free(unsafe.Pointer(cpszPreviewFile))
	return bool(C.SteamUGC_SetItemPreview(u.Pointer, C.UGCUpdateHandle_t(handle), cpszPreviewFile))
}

// remove any existing key-value tags with the specified key
func (u ISteamUGC) RemoveItemKeyValueTags(handle UGCUpdateHandle, pchKey string) bool {
	cpchKey := C.CString(pchKey)
	defer C.free(unsafe.Pointer(cpchKey))
	return bool(C.SteamUGC_RemoveItemKeyValueTags(u.Pointer, C.UGCUpdateHandle_t(handle), cpchKey))
}

// add new key-value tags for the item. Note that there can be multiple values for a tag.
func (u ISteamUGC) AddItemKeyValueTag(handle UGCUpdateHandle, pchKey string, pchValue string) bool {
	cpchKey := C.CString(pchKey)
	defer C.free(unsafe.Pointer(cpchKey))
	cpchValue := C.CString(pchValue)
	defer C.free(unsafe.Pointer(cpchValue))
	return bool(C.SteamUGC_AddItemKeyValueTag(u.Pointer, C.UGCUpdateHandle_t(handle), cpchKey, cpchValue))
}

// commit update process started with StartItemUpdate()
func (u ISteamUGC) SubmitItemUpdate(handle UGCUpdateHandle, pchChangeNote string) SteamAPICall {
	cpchChangeNote := C.CString(pchChangeNote)
	defer C.free(unsafe.Pointer(cpchChangeNote))
	return SteamAPICall(C.SteamUGC_SubmitItemUpdate(u.Pointer, C.UGCUpdateHandle_t(handle), cpchChangeNote))
}

func (u ISteamUGC) GetItemUpdateProgress(handle UGCUpdateHandle, punBytesProcessed *uint64, punBytesTotal *uint64) EItemUpdateStatus {
	return EItemUpdateStatus(C.SteamUGC_GetItemUpdateProgress(u.Pointer, C.UGCUpdateHandle_t(handle), (*C.uint64_t)(punBytesProcessed), (*C.uint64_t)(punBytesTotal)))
}

// Steam Workshop Consumer API

func (u ISteamUGC) SetUserItemVote(nPublishedFileID PublishedFileId, bVoteUp bool) SteamAPICall {
	return SteamAPICall(C.SteamUGC_SetUserItemVote(u.Pointer, C.PublishedFileId_t(nPublishedFileID), C._Bool(bVoteUp)))
}

func (u ISteamUGC) GetUserItemVote(nPublishedFileID PublishedFileId) SteamAPICall {
	return SteamAPICall(C.SteamUGC_GetUserItemVote(u.Pointer, C.PublishedFileId_t(nPublishedFileID)))
}

func (u ISteamUGC) AddItemToFavorites(nAppId AppId, nPublishedFileID PublishedFileId) SteamAPICall {
	return SteamAPICall(C.SteamUGC_AddItemToFavorites(u.Pointer, C.AppId_t(nAppId), C.PublishedFileId_t(nPublishedFileID)))
}

func (u ISteamUGC) RemoveItemFromFavorites(nAppId AppId, nPublishedFileID PublishedFileId) SteamAPICall {
	return SteamAPICall(C.SteamUGC_RemoveItemFromFavorites(u.Pointer, C.AppId_t(nAppId), C.PublishedFileId_t(nPublishedFileID)))
}

// SubscribeItem subscribes to this item, will be installed ASAP
func (u ISteamUGC) SubscribeItem(nPublishedFileID PublishedFileId) SteamAPICall {
	return SteamAPICall(C.SteamUGC_SubscribeItem(u.Pointer, C.PublishedFileId_t(nPublishedFileID)))
}

// UnsubscribeItem unsubscribes from this item, will be uninstalled after game quits
func (u ISteamUGC) UnsubscribeItem(nPublishedFileID PublishedFileId) SteamAPICall {
	return SteamAPICall(C.SteamUGC_UnsubscribeItem(u.Pointer, C.PublishedFileId_t(nPublishedFileID)))
}

// GetNumSubscribedItems returns the number of subscribed items
func (u ISteamUGC) GetNumSubscribedItems() uint32 {
	return uint32(C.SteamUGC_GetNumSubscribedItems(u.Pointer))
}

// all subscribed item PublishFileIDs
func (u ISteamUGC) GetSubscribedItems(pvecPublishedFileID *PublishedFileId, cMaxEntries uint32) uint32 {
	return uint32(C.SteamUGC_GetSubscribedItems(u.Pointer, (*C.PublishedFileId_t)(pvecPublishedFileID), C.uint(cMaxEntries)))
}

// GetItemState returns the EItemState flags about item on this client
func (u ISteamUGC) GetItemState(nPublishedFileID PublishedFileId) uint32 {
	return uint32(C.SteamUGC_GetItemState(u.Pointer, C.PublishedFileId_t(nPublishedFileID)))
}

// GetItemInstallInfo returns the info about currently installed content on disc for items that have k_EItemStateInstalled set
// if k_EItemStateLegacyItem is set, pchFolder contains the path to the legacy file itself (not a folder)
func (u ISteamUGC) GetItemInstallInfo(nPublishedFileID PublishedFileId, punSizeOnDisk *uint64, pchFolder string, cchFolderSize uint32, punTimeStamp *uint32) bool {
	cpchFolder := C.CString(pchFolder)
	defer C.free(unsafe.Pointer(cpchFolder))
	return bool(C.SteamUGC_GetItemInstallInfo(u.Pointer, C.PublishedFileId_t(nPublishedFileID), (*C.uint64_t)(punSizeOnDisk), cpchFolder, C.uint(cchFolderSize), (*C.uint)(punTimeStamp)))
}

// GetItemDownloadInfo returns the info about pending update for items that have k_EItemStateNeedsUpdate set. punBytesTotal will be valid after download started once
func (u ISteamUGC) GetItemDownloadInfo(nPublishedFileID PublishedFileId, punBytesDownloaded *uint64, punBytesTotal *uint64) bool {
	return bool(C.SteamUGC_GetItemDownloadInfo(u.Pointer, C.PublishedFileId_t(nPublishedFileID), (*C.uint64_t)(punBytesDownloaded), (*C.uint64_t)(punBytesTotal)))
}

// DownloadItem downloads new or update already installed item. If function returns true, wait for DownloadItemResult_t. If the item is already installed,
// then files on disk should not be used until callback received. If item is not subscribed to, it will be cached for some time.
// If bHighPriority is set, any other item download will be suspended and this item downloaded ASAP.
func (u ISteamUGC) DownloadItem(nPublishedFileID PublishedFileId, bHighPriority bool) bool {
	return bool(C.SteamUGC_DownloadItem(u.Pointer, C.PublishedFileId_t(nPublishedFileID), C._Bool(bHighPriority)))
}
