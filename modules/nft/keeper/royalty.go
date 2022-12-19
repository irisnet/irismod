package keeper

import (
	"math/big"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/nft"

	errorsmod "cosmossdk.io/errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdkmath "cosmossdk.io/math"
	"github.com/irisnet/irismod/modules/nft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const DefaultFeeDenominator = 10000

// SaveDefaultRoyalty sets the default royalty information of a class
func (k Keeper) SaveDefaultRoyalty(ctx sdk.Context, denomId string, receiver string, feeNumerator sdkmath.Uint, srcOwner sdk.AccAddress) error {
	denom, err := k.GetDenomInfo(ctx, denomId)
	if err != nil {
		return err
	}

	// authorize
	if srcOwner.String() != denom.Creator {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to transfer denom %s", srcOwner.String(), denomId)
	}

	// 1. Unmarshal denomInfo
	var royaltyMetadata types.RoyaltyMetadata
	if err := k.cdc.Unmarshal([]byte(denom.Data), &royaltyMetadata); err != nil {
		return err
	}

	if !royaltyMetadata.Enabled {
		return types.ErrNotEnabledRoyalty
	}

	royaltyMetadata.DefaultRoyaltyInfo = &types.RoyaltyInfo{
		Receiver:        receiver,
		RoyaltyFraction: feeNumerator,
	}

	denomDataBytes, err := codectypes.NewAnyWithValue(&royaltyMetadata)
	if err != nil {
		return err
	}

	denomMetadata := &types.DenomMetadata{
		Creator:          denom.Creator,
		Schema:           denom.Schema,
		MintRestricted:   denom.MintRestricted,
		UpdateRestricted: denom.UpdateRestricted,
		Data:             denomDataBytes.String(),
	}
	data, err := codectypes.NewAnyWithValue(denomMetadata)
	if err != nil {
		return err
	}
	class := nft.Class{
		Id:     denom.Id,
		Name:   denom.Name,
		Symbol: denom.Symbol,
		Data:   data,

		Description: denom.Description,
		Uri:         denom.Uri,
		UriHash:     denom.UriHash,
	}
	return k.nk.UpdateClass(ctx, class)
}

// SaveTokenRoyalty sets the royalty information of a token under a class
func (k Keeper) SaveTokenRoyalty(ctx sdk.Context, denomId string, tokenId string, receiver string, feeNumerator sdkmath.Uint, owner sdk.AccAddress) error {
	// just the owner of NFT can edit
	if err := k.Authorize(ctx, denomId, tokenId, owner); err != nil {
		return err
	}

	nftM, err := k.GetNFT(ctx, denomId, tokenId)
	if err != nil {
		return errorsmod.Wrapf(types.ErrUnknownNFT, "not found NFT: %s", denomId)
	}

	srcNftData := nftM.GetData()
	var royaltyInfo types.RoyaltyInfo
	if err := k.cdc.Unmarshal([]byte(srcNftData), &royaltyInfo); err != nil {
		return err
	}
	royaltyInfo.Receiver = receiver
	royaltyInfo.RoyaltyFraction = feeNumerator

	dstData, err := codectypes.NewAnyWithValue(&royaltyInfo)
	if err != nil {
		return err
	}

	// modify nftMetadata
	return k.UpdateNFT(
		ctx,
		denomId,
		nftM.GetID(),
		nftM.GetName(),
		nftM.GetURI(),
		nftM.GetURIHash(),
		dstData.String(),
		nftM.GetOwner(),
	)
}

// RemoveTokenRoyalty deletes the royalty information of a token under a class
func (k Keeper) RemoveTokenRoyalty(ctx sdk.Context, denomId string, tokenId string, owner sdk.AccAddress) error {
	// just the owner of NFT can edit
	if err := k.Authorize(ctx, denomId, tokenId, owner); err != nil {
		return err
	}

	nftM, err := k.GetNFT(ctx, denomId, tokenId)
	if err != nil {
		return errorsmod.Wrapf(types.ErrUnknownNFT, "not found NFT: %s", denomId)
	}

	srcNftData := nftM.GetData()
	var royaltyInfo types.RoyaltyInfo
	if err := k.cdc.Unmarshal([]byte(srcNftData), &royaltyInfo); err != nil {
		return err
	}

	// modify nftMetadata
	return k.UpdateNFT(
		ctx,
		denomId,
		nftM.GetID(),
		nftM.GetName(),
		nftM.GetURI(),
		nftM.GetURIHash(),
		"",
		nftM.GetOwner(),
	)

}

// RemoveDefaultRoyalty deletes the default royalty information of a class
func (k Keeper) RemoveDefaultRoyalty(ctx sdk.Context, denomId string, srcOwner sdk.AccAddress) error {

	denom, err := k.GetDenomInfo(ctx, denomId)
	if err != nil {
		return err
	}

	// authorize
	if srcOwner.String() != denom.Creator {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to transfer denom %s", srcOwner.String(), denomId)
	}

	// 1. Unmarshal denomInfo
	var royaltyMetadata types.RoyaltyMetadata
	if err := k.cdc.Unmarshal([]byte(denom.Data), &royaltyMetadata); err != nil {
		return err
	}

	if !royaltyMetadata.Enabled {
		return types.ErrNotEnabledRoyalty
	}

	royaltyMetadata.DefaultRoyaltyInfo = nil
	royaltyMetadata.Enabled = false
	denomDataBytes, err := codectypes.NewAnyWithValue(&royaltyMetadata)
	if err != nil {
		return err
	}

	denomMetadata := &types.DenomMetadata{
		Creator:          denom.Creator,
		Schema:           denom.Schema,
		MintRestricted:   denom.MintRestricted,
		UpdateRestricted: denom.UpdateRestricted,
		Data:             denomDataBytes.String(),
	}
	data, err := codectypes.NewAnyWithValue(denomMetadata)
	if err != nil {
		return err
	}
	class := nft.Class{
		Id:     denom.Id,
		Name:   denom.Name,
		Symbol: denom.Symbol,
		Data:   data,

		Description: denom.Description,
		Uri:         denom.Uri,
		UriHash:     denom.UriHash,
	}
	return k.nk.UpdateClass(ctx, class)
}

// GetFeeDenominator returns the denominator of the fee
func (k Keeper) GetFeeDenominator(ctx sdk.Context) (feeNumerator sdkmath.Uint) {
	feeNumerator = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(DefaultFeeDenominator))
	return
}

// GetRoyaltyInfo returns the royalty information of a token under a class
func (k Keeper) GetRoyaltyInfo(ctx sdk.Context, denomId string, nftId string, salePrice sdkmath.Uint) (string, sdkmath.Uint, error) {

	var receiver string
	var feeNumerator, royaltyAmount sdkmath.Uint
	var err error
	receiver, feeNumerator, err = k.GetTokenRoyaltyInfo(ctx, denomId, nftId)
	if len(receiver) == 0 {
		receiver, feeNumerator, err = k.GetDefaultRoyaltyInfo(ctx, denomId)
		if err != nil {
			return "", sdkmath.Uint{}, err
		}
	}
	royaltyAmount = salePrice.Mul(feeNumerator).Quo(k.GetFeeDenominator(ctx))
	return receiver, royaltyAmount, nil
}

// GetDefaultRoyaltyInfo returns the default royalty information of a class
func (k Keeper) GetDefaultRoyaltyInfo(ctx sdk.Context, denomId string) (string, sdkmath.Uint, error) {
	denom, err := k.GetDenomInfo(ctx, denomId)
	if err != nil {
		return "", sdkmath.Uint{}, err
	}
	var royaltyMetadata types.RoyaltyMetadata
	if err := k.cdc.Unmarshal([]byte(denom.Data), &royaltyMetadata); err != nil {
		return "", sdkmath.Uint{}, err
	}

	if !royaltyMetadata.Enabled {
		return "", sdkmath.Uint{}, types.ErrNotEnabledRoyalty
	}
	return royaltyMetadata.DefaultRoyaltyInfo.Receiver, royaltyMetadata.DefaultRoyaltyInfo.RoyaltyFraction, nil

}

// GetTokenRoyaltyInfo returns the royalty information of a token under a class
func (k Keeper) GetTokenRoyaltyInfo(ctx sdk.Context, denomId string, tokenId string) (string, sdkmath.Uint, error) {
	nftM, err := k.GetNFT(ctx, denomId, tokenId)
	if err != nil {
		return "", sdkmath.Uint{}, errorsmod.Wrapf(types.ErrUnknownNFT, "not found NFT: %s", denomId)
	}
	nftData := nftM.GetData()
	var royaltyInfo types.RoyaltyInfo
	if err := k.cdc.Unmarshal([]byte(nftData), &royaltyInfo); err != nil {
		return "", sdkmath.Uint{}, err
	}
	return royaltyInfo.Receiver, royaltyInfo.RoyaltyFraction, nil
}
