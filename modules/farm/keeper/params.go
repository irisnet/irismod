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

// SetParams sets the params to the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) (fee sdk.Coin) {
	k.paramSpace.SetParamSet(ctx, &params)
	return
}
