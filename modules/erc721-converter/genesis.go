package erc721converter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/irisnet/irismod/modules/erc721-converter/keeper"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	data types.GenesisState,
) {

	// ensure erc721 module account is set on genesis
	if acc := accountKeeper.GetModuleAccount(ctx, types.ModuleName); acc == nil {
		// NOTE: shouldn't occur
		panic("the erc721 module account has not been set")
	}

	for _, pair := range data.TokenPairs {
		id := pair.GetID()
		k.SetTokenPair(ctx, pair)
		k.SetClassMap(ctx, pair.ClassId, id)
		k.SetERC721Map(ctx, pair.GetERC721Contract(), id)
	}
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		TokenPairs: k.GetTokenPairs(ctx),
	}
}
