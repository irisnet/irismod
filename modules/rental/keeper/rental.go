package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// setRentalInfo sets the rental info for an nft.
func (k Keeper) setRentalInfo(ctx sdk.Context,
	classId, nftId string,
	user sdk.AccAddress,
	expires uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(rentalInfoKey(classId, nftId), rentalInfoValue(user, expires))
}

// getRentalInfo returns the rental info for an nft.
func (k Keeper) getRentalInfo(ctx sdk.Context,
	classId, nftId string) (user sdk.AccAddress, expires uint64) {
	store := ctx.KVStore(k.storeKey)
	vbz := store.Get(rentalInfoKey(classId, nftId))

	if vbz != nil {
		user, expires = splitRentalInfoValue(vbz)
	}
	return
}
