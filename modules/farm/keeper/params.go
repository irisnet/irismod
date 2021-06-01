package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/irisnet/irismod/modules/farm/types"
)

// ParamKeyTable for farm module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&types.Params{})
}

// CreatePoolFee returns the create pool fee
func (k Keeper) CreatePoolFee(ctx sdk.Context) (fee sdk.Coin) {
	k.paramSpace.Get(ctx, types.KeyCreatePoolFee, &fee)
	return
}

// MaxRewardCategory returns the maxRewardCategory
func (k Keeper) MaxRewardCategory(ctx sdk.Context) (maxRewardCategory uint32) {
	k.paramSpace.Get(ctx, types.KeyMaxRewardCategory, &maxRewardCategory)
	return
}

// SetParams sets the params to the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) (fee sdk.Coin) {
	k.paramSpace.SetParamSet(ctx, &params)
	return
}

// GetParams gets all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.CreatePoolFee(ctx),
		k.MaxRewardCategory(ctx),
	)
}
