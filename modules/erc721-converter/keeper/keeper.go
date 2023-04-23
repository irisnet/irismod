package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/modules/erc721-converter/types"

	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of this module maintains collections of erc721.
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	accountKeeper    types.AccountKeeper
	evmKeeper        types.EVMKeeper
	nftKeeper        types.NFTKeeper
	irisModNFTKeeper types.IRISModNFTKeeper
}

func NewKeeper(
	storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
	ak types.AccountKeeper,
	evmKeeper types.EVMKeeper,
	nftKeeper types.NFTKeeper,
	irisModNFTKeeper types.IRISModNFTKeeper,
) Keeper {

	return Keeper{
		storeKey:         storeKey,
		cdc:              cdc,
		accountKeeper:    ak,
		evmKeeper:        evmKeeper,
		nftKeeper:        nftKeeper,
		irisModNFTKeeper: irisModNFTKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("irismod/%s", types.ModuleName))
}
