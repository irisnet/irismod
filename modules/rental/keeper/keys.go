package keeper

import (
	"github.com/irisnet/irismod/modules/rental/types"
	"reflect"
	"unsafe"
)

var (
	RentalInfoKey = []byte{0x01}

	Delimiter = []byte{0x00}
)

// StoreKey is the store key for rental module
const StoreKey = types.ModuleName

// rentalInfoKey returns the byte representation of the rental info key.
// This key comprises of <0x01><class-id><delimiter><nft-id>
func rentalInfoKey(classId, nftId string) []byte {
	classIdBz := unsafeStrToBytes(classId)
	nftIdBz := unsafeStrToBytes(nftId)

	key := make([]byte, len(RentalInfoKey)+len(classIdBz)+len(Delimiter)+len(nftIdBz))

	copy(key, RentalInfoKey)
	copy(key[len(RentalInfoKey):], classIdBz)
	copy(key[len(RentalInfoKey)+len(classIdBz):], Delimiter)
	copy(key[len(RentalInfoKey)+len(classIdBz)+len(Delimiter):], nftIdBz)

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

// unsafeBytesToStr is meant to make a zero allocation conversion
// from []byte -> string to speed up operations, it is not meant
// to be used generally, but for a specific pattern to delete keys
// from a map.
func unsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
