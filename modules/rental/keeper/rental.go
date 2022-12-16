package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/modules/rental/types"
)

// SetRentalInfo sets the rental info for an nft.
func (k Keeper) SetRentalInfo(ctx sdk.Context,
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

// GetRentalInfo returns the rental info for an nft.
func (k Keeper) GetRentalInfo(ctx sdk.Context,
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

// GetRentalInfos returns all rental infos.
func (k Keeper) GetRentalInfos(ctx sdk.Context) (ris []types.RentalInfo) {
	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var rental types.RentalInfo
		k.cdc.MustUnmarshal(iterator.Value(), &rental)
		ris = append(ris, rental)
	}
	return ris
}
