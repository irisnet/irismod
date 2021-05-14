package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/farm/types"
)

// Keeper of the farm store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler

	bk types.BankKeeper
}

func NewKeeper(storeKey sdk.StoreKey, cdc codec.Marshaler, bk types.BankKeeper) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
		bk:       bk,
	}
}

// CreatePool creates an new farm pool
func (k Keeper) SetPool(ctx sdk.Context, pool types.FarmPool) {
	pool.Rule = nil
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&pool)
	store.Set(types.GetFarmPoolKey(pool.Name), bz)
}

func (k Keeper) SetPoolRule(ctx sdk.Context, poolName string, rule types.FarmRule) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&rule)
	store.Set(types.GetFarmPoolRuleKey(poolName, rule.Reward), bz)
}

func (k Keeper) Enqueue(ctx sdk.Context, poolName string, expiredHeight uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.GetFarmPoolExpiredKey(expiredHeight, poolName),
		types.MustMarshalPoolName(k.cdc, poolName),
	)
}
