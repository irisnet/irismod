package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.Codec
	nk       nftkeeper.Keeper
}

// NewKeeper creates a new instance of rental keeper
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, nk nftkeeper.Keeper) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
		nk:       nk,
	}
}
