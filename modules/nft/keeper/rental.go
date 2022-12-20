package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/nft/types"
)

// Rent set or update rental info for an nft.
func (k Keeper) Rent(ctx sdk.Context, rental types.RentalInfo) error {
	// if disabled, return err
	if enabled := k.GetRentalOption(ctx, rental.DenomId); !enabled {
		return sdkerrors.Wrapf(types.ErrRentalOption, "Rental is disabled")
	}

	// expiry should be greater than the current time
	if ctx.BlockTime().Unix() >= int64(rental.Expires) {
		return sdkerrors.Wrapf(types.ErrInvalidExpiry, "Expiry is (%d)", rental.Expires)
	}

	// set rental info
	k.setRentalInfo(ctx, rental.DenomId, rental.NftId, rental.User, rental.Expires)
	return nil
}

// setRentalInfo sets the rental info for an nft.
func (k Keeper) setRentalInfo(ctx sdk.Context,
	classId, nftId, user string,
	expires int64) {
	store := ctx.KVStore(k.storeKey)
	r := types.RentalInfo{
		User:    user,
		DenomId: classId,
		NftId:   nftId,
		Expires: expires,
	}
	bz := k.cdc.MustMarshal(&r)
	store.Set(rentalInfoKey(r.DenomId, r.NftId), bz)
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

// setRentalOption enables the rental feature for a class.
func (k Keeper) setRentalOption(ctx sdk.Context, denomId string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(rentalOptionKey(denomId), []byte{0x01})
}

// unsetRentalOption disables the rental feature for a class.
func (k Keeper) unsetRentalOption(ctx sdk.Context, denomId string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(rentalOptionKey(denomId), []byte{0x00})
}

// GetRentalEnabled checks if a class has its rental option enabled.
func (k Keeper) GetRentalOption(ctx sdk.Context, denomId string) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(rentalOptionKey(denomId))

	if bytes.Equal(bz, []byte{0x01}) {
		return true
	}
	return false
}
