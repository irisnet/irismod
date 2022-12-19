package keeper

import (
	"reflect"
	"unsafe"
)

// As iris/nft uses x/nft as its base module, new added key must start from 0x06.
// Key of each plugin must add an PrefixPluginXxx to distinguish.
var (
	PluginRental = []byte{0x06}

	RentalInfoKey = []byte{0x01}

	Delimiter = []byte{0x00}
)

// rentalInfoKey returns the byte representation of the rental info key.
// This key comprises of <0x06><0x01><class-id><delimiter><nft-id>
func rentalInfoKey(classId, nftId string) []byte {
	classIdBz := unsafeStrToBytes(classId)
	nftIdBz := unsafeStrToBytes(nftId)

	key := make([]byte, len(PluginRental)+len(RentalInfoKey)+len(classIdBz)+len(Delimiter)+len(nftIdBz))

	copy(key, PluginRental)
	copy(key[len(PluginRental):], RentalInfoKey)
	copy(key[len(PluginRental)+len(RentalInfoKey):], classIdBz)
	copy(key[len(PluginRental)+len(RentalInfoKey)+len(classIdBz):], Delimiter)
	copy(key[len(PluginRental)+len(RentalInfoKey)+len(classIdBz)+len(Delimiter):], nftIdBz)
	return key
}

// The following functions refers to cosmos-sdk/internal/conv/string.go

// unsafeStrToBytes uses unsafe to convert string into byte array. Returned bytes
// must not be altered after this function is called as it will cause a segmentation fault.
func unsafeStrToBytes(s string) []byte {
	var buf []byte
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bufHdr := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	bufHdr.Data = sHdr.Data
	bufHdr.Cap = sHdr.Len
	bufHdr.Len = sHdr.Len
	return buf
}
