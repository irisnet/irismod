package types

import (
	"bytes"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "mt"

	// StoreKey is the default store key for MT
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the MT store.
	QuerierRoute = ModuleName

	// RouterKey is the message route for the MT module
	RouterKey = ModuleName
)

var (
	PrefixMT         = []byte{0x01}
	PrefixBalance    = []byte{0x02} // key for a owner
	PrefixCollection = []byte{0x03} // key for balance of MTs held by the denom
	PrefixDenom      = []byte{0x04} // key for denom of the mt

	delimiter = []byte("0x00")
)

// SplitKeyOwner return the address,denom,id from the key of stored owner
func SplitKeyOwner(key []byte) (address sdk.AccAddress, denomID, mtID string, err error) {
	key = key[len(PrefixBalance)+len(delimiter):]
	keys := bytes.Split(key, delimiter)
	if len(keys) != 3 {
		return address, denomID, mtID, errors.New("wrong KeyOwner")
	}

	address, _ = sdk.AccAddressFromBech32(string(keys[0]))
	denomID = string(keys[1])
	mtID = string(keys[2])
	return
}

func SplitKeyDenom(key []byte) (denomID, mtID string, err error) {
	keys := bytes.Split(key, delimiter)
	if len(keys) != 2 {
		return denomID, mtID, errors.New("wrong KeyOwner")
	}

	denomID = string(keys[0])
	mtID = string(keys[1])
	return
}

// KeyOwner gets the key of a collection owned by an account address
func KeyOwner(address sdk.AccAddress, denomID, mtID string) []byte {
	key := append(PrefixBalance, delimiter...)
	if address != nil {
		key = append(key, []byte(address.String())...)
		key = append(key, delimiter...)
	}

	if address != nil && len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}

	if address != nil && len(denomID) > 0 && len(mtID) > 0 {
		key = append(key, []byte(mtID)...)
	}
	return key
}

// KeyMT gets the key of mt stored by an denom and id
func KeyMT(denomID, mtID string) []byte {
	key := append(PrefixMT, delimiter...)
	if len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}

	if len(denomID) > 0 && len(mtID) > 0 {
		key = append(key, []byte(mtID)...)
	}
	return key
}

// KeyCollection gets the storeKey by the collection
func KeyCollection(denomID string) []byte {
	key := append(PrefixCollection, delimiter...)
	return append(key, []byte(denomID)...)
}

// KeyDenomID gets the storeKey by the denom id
func KeyDenomID(id string) []byte {
	key := append(PrefixDenom, delimiter...)
	return append(key, []byte(id)...)
}
