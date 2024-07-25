package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"mods.irisnet.org/modules/nft/types"
)

// InitGenesis stores the NFT genesis.
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	for _, c := range data.Collections {
		if err := k.SetDenom(ctx, c.Denom); err != nil {
			panic(err)
		}
		if err := k.SetCollection(ctx, c); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(k.GetCollections(ctx))
}
