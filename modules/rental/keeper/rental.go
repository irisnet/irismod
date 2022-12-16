package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/modules/rental/types"
)

// setRentalInfo sets the rental info for an nft.
func (k Keeper) setRentalInfo(ctx sdk.Context,
	classId, nftId string,
	user sdk.AccAddress,
	expires uint64) {
	store := ctx.KVStore(k.storeKey)
	r := types.RentalInfo{
		User:    user.String(),
		ClassId: classId,
		NftId:   nftId,
		Expires: expires,
	}
	bz := k.cdc.MustMarshal(&r)
	store.Set(rentalInfoKey(r.ClassId, r.NftId), bz)
}

// getRentalInfo returns the rental info for an nft.
func (k Keeper) getRentalInfo(ctx sdk.Context,
	classId, nftId string) (types.RentalInfo, bool) {
	var v types.RentalInfo
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(rentalInfoKey(classId, nftId))
	if bz == nil {
		return types.RentalInfo{}, false
	}
	k.cdc.MustUnmarshal(bz, &v)
	return v, true
}
