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
func (k Keeper) SaveDefaultRoyalty(ctx sdk.Context, denomId string, receiver string, fraction sdkmath.Uint, srcOwner sdk.AccAddress) error {
	denom, err := k.GetDenomInfo(ctx, denomId)
	if err != nil {
		return err
	}

	// authorize
	if srcOwner.String() != denom.Creator {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to transfer denom %s", srcOwner.String(), denomId)
	}

	royaltyPlugin := &types.RoyaltyPlugin{
		Enabled:  true,
		Receiver: receiver,
		Fraction: fraction,
	}

	denomMetadata := &types.DenomMetadata{
		Creator:          denom.Creator,
		Schema:           denom.Schema,
		MintRestricted:   denom.MintRestricted,
		UpdateRestricted: denom.UpdateRestricted,
		Data:             denom.Data,
		RoyaltyPlugin:    royaltyPlugin,
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

	if !k.IsNotEnabledRoyalty(ctx, denomId) {
		return types.ErrNotEnabledRoyalty
	}

	denomMetadata := &types.DenomMetadata{
		Creator:          denom.Creator,
		Schema:           denom.Schema,
		MintRestricted:   denom.MintRestricted,
		UpdateRestricted: denom.UpdateRestricted,
		Data:             denom.Data,
		RoyaltyPlugin:    nil,
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
func (k Keeper) SaveTokenRoyalty(ctx sdk.Context, denomId string, tokenId string, receiver string, fraction sdkmath.Uint, owner sdk.AccAddress) error {

	// just the owner of NFT can edit
	if err := k.Authorize(ctx, denomId, tokenId, owner); err != nil {
		return err
	}

	if !k.IsNotEnabledRoyalty(ctx, denomId) {
		return types.ErrNotEnabledRoyalty
	}

	nftM, err := k.GetNFT(ctx, denomId, tokenId)
	if err != nil {
		return errorsmod.Wrapf(types.ErrUnknownNFT, "not found NFT: %s", denomId)
	}

	tokenRoyaltyInfo := &types.TokenRoyaltyPlugin{
		Receiver: receiver,
		Fraction: fraction,
	}

	dstData, err := k.cdc.Marshal(tokenRoyaltyInfo)
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
		string(dstData),
		nftM.GetOwner(),
	)
}

// RemoveTokenRoyalty deletes the royalty information of a token under a class
func (k Keeper) RemoveTokenRoyalty(ctx sdk.Context, denomId string, tokenId string, owner sdk.AccAddress) error {
	// just the owner of NFT can edit
	if err := k.Authorize(ctx, denomId, tokenId, owner); err != nil {
		return err
	}

	if !k.IsNotEnabledRoyalty(ctx, denomId) {
		return types.ErrNotEnabledRoyalty
	}

	nftM, err := k.GetNFT(ctx, denomId, tokenId)
	if err != nil {
		return errorsmod.Wrapf(types.ErrUnknownNFT, "not found NFT: %s", denomId)
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

	if !k.IsNotEnabledRoyalty(ctx, denomId) {
		return "", sdkmath.Uint{}, types.ErrNotEnabledRoyalty
	}

	denomMetadata, err := k.GetDenomMetadata(ctx, denomId)
	if err != nil {
		return "", sdkmath.Uint{}, err
	}

	return denomMetadata.RoyaltyPlugin.Receiver, denomMetadata.RoyaltyPlugin.Fraction, nil

}

// GetTokenRoyaltyInfo returns the royalty information of a token under a class
func (k Keeper) GetTokenRoyaltyInfo(ctx sdk.Context, denomId string, tokenId string) (string, sdkmath.Uint, error) {

	if !k.IsNotEnabledRoyalty(ctx, denomId) {
		return "", sdkmath.Uint{}, types.ErrNotEnabledRoyalty
	}

	nftMetadata, err := k.GetNftMetadata(ctx, denomId, tokenId)
	if err != nil {
		return "", sdkmath.Uint{}, errorsmod.Wrapf(types.ErrUnknownNFT, "not found NFT: %s", denomId)
	}

	if nftMetadata.TokenRoyaltyPlugin == nil {
		return "", sdkmath.Uint{}, types.ErrNullTokenRoyaltyInfo
	}

	return nftMetadata.TokenRoyaltyPlugin.Receiver, nftMetadata.TokenRoyaltyPlugin.Fraction, nil
}

func (k Keeper) IsNotEnabledRoyalty(ctx sdk.Context, denomId string) bool {

	denomMetadata, err := k.GetDenomMetadata(ctx, denomId)
	if err != nil {
		return false
	}

	if denomMetadata.RoyaltyPlugin == nil || !denomMetadata.RoyaltyPlugin.Enabled {
		return false
	}

	return true
}

func (k Keeper) getTokenRoyaltyInfoFromTokenData(ctx sdk.Context, tokenData, denomId string) (*types.TokenRoyaltyPlugin, string) {
	// royalty option
	if k.IsNotEnabledRoyalty(ctx, denomId) {
		if len(tokenData) == 0 {
			return nil, ""
		}
		tokenRoyaltyPlugin := new(types.TokenRoyaltyPlugin)
		if err := k.cdc.Unmarshal([]byte(tokenData), tokenRoyaltyPlugin); err != nil {
			return nil, ""
		}
		if tokenRoyaltyPlugin != nil {
			return tokenRoyaltyPlugin, tokenData
		} else {
			return nil, ""
		}
	}
	return nil, tokenData
}
