package keeper

import (
	"math/big"

	"github.com/irisnet/irismod/modules/nft/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const DefaultFeeDenominator = 10000

// GetFeeDenominator returns the denominator of the fee
func (k Keeper) GetFeeDenominator(ctx sdk.Context) (feeNumerator sdkmath.Uint) {
	feeNumerator = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(DefaultFeeDenominator))
	return
}

// GetRoyaltyInfo returns the royalty information of a token under a class
func (k Keeper) GetRoyaltyInfo(ctx sdk.Context, classId string, tokenId string, salePrice sdkmath.Uint) (receiver string, royaltyAmount sdkmath.Uint) {
	receiver, feeNumerator := k.GetTokenRoyaltyInfo(ctx, classId, tokenId)
	if len(receiver) == 0 {
		receiver, feeNumerator = k.GetDefaultRoyaltyInfo(ctx, classId)
	}
	royaltyAmount = salePrice.Mul(feeNumerator).Quo(k.GetFeeDenominator(ctx))
	return
}

// GetDefaultRoyaltyInfo returns the default royalty information of a class
func (k Keeper) GetDefaultRoyaltyInfo(ctx sdk.Context, classId string) (address string, royaltyAmount sdkmath.Uint) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyDefaultRoyalty(classId))
	if len(bz) == 0 {
		return "", sdkmath.Uint{}
	}
	royaltyInfo := types.RoyaltyInfo{}
	k.cdc.MustMarshalJSON(&royaltyInfo)
	return royaltyInfo.Address, royaltyInfo.RoyaltyFraction
}

// GetTokenRoyaltyInfo returns the royalty information of a token under a class
func (k Keeper) GetTokenRoyaltyInfo(ctx sdk.Context, classId string, tokenId string) (address string, royaltyAmount sdkmath.Uint) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyTokenRoyalty(classId, tokenId))
	if len(bz) == 0 {
		return "", sdkmath.Uint{}
	}
	royaltyInfo := types.RoyaltyInfo{}
	k.cdc.MustMarshalJSON(&royaltyInfo)
	return royaltyInfo.Address, royaltyInfo.RoyaltyFraction
}

// SetDefaultRoyalty sets the default royalty information of a class
func (k Keeper) SetDefaultRoyalty(ctx sdk.Context, classId string, receiver string, feeNumerator sdkmath.Uint) {
	store := ctx.KVStore(k.storeKey)

	royaltyInfo := types.RoyaltyInfo{
		Address:         receiver,
		RoyaltyFraction: feeNumerator,
	}

	bz := k.cdc.MustMarshal(&royaltyInfo)
	store.Set(types.KeyDefaultRoyalty(classId), bz)
}

// SetTokenRoyalty sets the royalty information of a token under a class
func (k Keeper) SetTokenRoyalty(ctx sdk.Context, classId string, tokenId string, receiver string, feeNumerator sdkmath.Uint) {
	store := ctx.KVStore(k.storeKey)

	royaltyInfo := types.RoyaltyInfo{
		Address:         receiver,
		RoyaltyFraction: feeNumerator,
	}

	bz := k.cdc.MustMarshal(&royaltyInfo)
	store.Set(types.KeyTokenRoyalty(classId, tokenId), bz)
}

// ResetTokenRoyalty deletes the royalty information of a token under a class
func (k Keeper) ResetTokenRoyalty(ctx sdk.Context, classId string, tokenId string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyTokenRoyalty(classId, tokenId))
}

// DeleteDefaultRoyalty deletes the default royalty information of a class
func (k Keeper) DeleteDefaultRoyalty(ctx sdk.Context, classId string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyDefaultRoyalty(classId))
}
