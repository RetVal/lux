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

// AppTicket hands out a reasonable "future proof" view of an app ownership
// ticket the raw (signed) buffer, and indices into that buffer where the appid
// and steamid are located.  the sizes of the appid and steamid are implicit in
// (each version of) the interface - currently uin32 appid and uint64 steamid
type AppTicket struct {
	unsafe.Pointer
}

// GetAppOwnershipTicketData has no documentation. Contact Valve :)
func (t AppTicket) GetAppOwnershipTicketData(nAppID uint32, pvBuffer unsafe.Pointer, cbBufferLength uint32, piAppID, piSteamID, piSignature, pcbSignature *uint32) uint32 {
	return uint32(C.SteamCAPI_ISteamAppTicket_GetAppOwnershipTicketData(t.Pointer, C.uint32_t(nAppID), pvBuffer, C.uint32_t(cbBufferLength), (*C.uint32_t)(piAppID), (*C.uint32_t)(piSteamID), (*C.uint32_t)(piSignature), (*C.uint32_t)(pcbSignature)))
}
