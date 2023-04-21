package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName is the name of the module
	ModuleName = "erc721converter"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the module
	RouterKey = ModuleName
)

// ModuleAddress is the native module address for the module
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

// prefix bytes for the EVM persistent store
const (
	prefixTokenPair = iota + 1
	prefixTokenPairByERC721
	prefixTokenPairByDenom
	prefixERC721TokenIDByNativeTokenID
	prefixNativeTokenIDByERC721TokenID
)

// KVStore key prefixes
var (
	KeyPrefixTokenPair                    = []byte{prefixTokenPair}
	KeyPrefixTokenPairByERC721            = []byte{prefixTokenPairByERC721}
	KeyPrefixTokenPairByDenom             = []byte{prefixTokenPairByDenom}
	KeyPrefixERC721TokenIDByNativeTokenID = []byte{prefixERC721TokenIDByNativeTokenID}
	KeyPrefixNativeTokenIDByERC721TokenID = []byte{prefixNativeTokenIDByERC721TokenID}
)
