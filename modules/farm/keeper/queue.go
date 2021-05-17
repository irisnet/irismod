package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/modules/farm/types"
)

func (k Keeper) EnqueueExpiredPool(ctx sdk.Context, poolName string, expiredHeight uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.GetFarmPoolExpiredKey(expiredHeight, poolName),
		types.MustMarshalPoolName(k.cdc, poolName),
	)
}

func (k Keeper) DequeueExpiredPool(ctx sdk.Context, poolName string, expiredHeight uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetFarmPoolExpiredKey(expiredHeight, poolName))
}

func (k Keeper) IteratorExpiredPool(ctx sdk.Context) (fps []*types.FarmPool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store,
		types.GetFarmPoolExpiredKeyPrefix(uint64(ctx.BlockHeight())))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		poolName := types.MustUnMarshalPoolName(k.cdc, iterator.Value())
		if pool, exist := k.GetPool(ctx, poolName); exist {
			fps = append(fps, pool)
		}
	}
	return fps
}
