package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/rental/types"
)

// InitGenesis stores the NFT genesis.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	if err := types.ValidateGenesis(gs); err != nil {
		panic(err.Error())
	}

	for _, v := range gs.RenterInfos {
		user, err := sdk.AccAddressFromBech32(v.User)
		if err != nil {
			panic(err)
		}
		k.SetRentalInfo(ctx, v.ClassId, v.NftId, user, v.Expires)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	ris := k.GetRentalInfos(ctx)
	return types.NewGenesisState(ris)
}
