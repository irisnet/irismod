package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/mt/exported"
	"github.com/irisnet/irismod/modules/mt/types"
)

// GetMT gets the the specified MT
func (k Keeper) GetMT(ctx sdk.Context, denomID, tokenID string) (mt exported.MT, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyMT(denomID, tokenID))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownCollection, "not found MT: %s", denomID)
	}

	var baseMT types.BaseMT
	k.cdc.MustUnmarshal(bz, &baseMT)

	return baseMT, nil
}

// GetMTs returns all MTs by the specified denom ID
func (k Keeper) GetMTs(ctx sdk.Context, denom string) (mts []exported.MT) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyMT(denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var baseMT types.BaseMT
		k.cdc.MustUnmarshal(iterator.Value(), &baseMT)
		mts = append(mts, baseMT)
	}

	return mts
}

// Authorize checks if the sender is the owner of the given MT
// Return the MT if true, an error otherwise
func (k Keeper) Authorize(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) (types.BaseMT, error) {
	mt, err := k.GetMT(ctx, denomID, tokenID)
	if err != nil {
		return types.BaseMT{}, err
	}

	if !owner.Equals(mt.GetOwner()) {
		return types.BaseMT{}, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}

	return mt.(types.BaseMT), nil
}

// HasMT checks if the specified MT exists
func (k Keeper) HasMT(ctx sdk.Context, denomID, tokenID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyMT(denomID, tokenID))
}

func (k Keeper) setMT(ctx sdk.Context, denomID string, mt types.BaseMT) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&mt)
	store.Set(types.KeyMT(denomID, mt.GetID()), bz)
}

// deleteMT deletes an existing MT from store
func (k Keeper) deleteMT(ctx sdk.Context, denomID string, mt exported.MT) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyMT(denomID, mt.GetID()))
}
