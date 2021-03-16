package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the HTLC module
	ModuleName = "htlc"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the HTLC module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the HTLC module
	RouterKey string = ModuleName

	// DefaultParamspace is the default name for parameter store
	DefaultParamspace = ModuleName
)

var (
	// ModulePermissionsUpgradeTime is the block time after which the htlc module account's permissions are synced with the supply module.
	ModulePermissionsUpgradeTime time.Time = time.Date(2020, 11, 3, 10, 0, 0, 0, time.UTC)

	// Keys for store prefixes
	HTLCKey              = []byte{0x01} // prefix for HTLC
	HTLCExpiredQueueKey  = []byte{0x02} // prefix for the HTLC expiration queue
	AssetSupplyPrefix    = []byte{0x03} // prefix for the HTLT supply
	PreviousBlockTimeKey = []byte{0x04} // prefix for the HTLT supply previous block time
)

// GetHTLCKey returns the key for the HTLC with the specified hash lock
// VALUE: htlc/HTLC
func GetHTLCKey(id []byte) []byte {
	return append(HTLCKey, id...)
}

// GetHTLCExpiredQueueKey returns the key for the HTLC expiration queue by the specified height and hash lock
// VALUE: []byte{}
func GetHTLCExpiredQueueKey(expirationHeight uint64, id []byte) []byte {
	return append(append(HTLCExpiredQueueKey, sdk.Uint64ToBigEndian(expirationHeight)...), id...)
}

// GetHTLCExpiredQueueSubspace returns the key prefix for the HTLC expiration queue by the given height
func GetHTLCExpiredQueueSubspace(expirationHeight uint64) []byte {
	return append(HTLCExpiredQueueKey, sdk.Uint64ToBigEndian(expirationHeight)...)
}

// GetAssetSupplyKey returns the key prefix for the asset supply by the given denom
func GetAssetSupplyKey(denom string) []byte {
	return append(AssetSupplyPrefix, []byte(denom)...)
}
