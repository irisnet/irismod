package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/nft/keeper"
	"github.com/irisnet/irismod/modules/nft/types"
)

// InitGenesis stores the NFT genesis.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	for _, c := range data.Collections {
		creator, err := sdk.AccAddressFromBech32(c.Denom.Creator)
		if err != nil {
			panic(err)
		}
		if err := k.SaveDenom(ctx,
			c.Denom.Id,
			c.Denom.Name,
			c.Denom.Schema,
			c.Denom.Symbol,
			creator,
			c.Denom.MintRestricted,
			c.Denom.UpdateRestricted,
			c.Denom.Description,
			c.Denom.Uri,
			c.Denom.UriHash,
			c.Denom.Data,
		); err != nil {
			panic(err)
		}

		if err := k.SaveCollection(ctx, c); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	collections, err := k.GetCollections(ctx)
	if err != nil {
		panic(err)
	}
	return types.NewGenesisState(collections)
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.Collection{})
}
