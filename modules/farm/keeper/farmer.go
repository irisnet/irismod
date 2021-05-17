package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/modules/farm/types"
)

// GetFarmer return the specified farmer
func (k Keeper) GetFarmer(ctx sdk.Context, poolName, address string) (farmer *types.Farmer, exist bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetFarmerKey(address, poolName))
	if len(bz) == 0 {
		return nil, false
	}

	k.cdc.MustUnmarshalBinaryBare(bz, farmer)
	return farmer, true
}

// SetFarmer save the farmer information
func (k Keeper) SetFarmer(ctx sdk.Context, poolName string, farmer types.Farmer) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&farmer)
	store.Set(types.GetFarmerKey(farmer.Address, poolName), bz)
}

func (k Keeper) DeleteFarmer(ctx sdk.Context, poolName, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetFarmerKey(address, poolName))
}
