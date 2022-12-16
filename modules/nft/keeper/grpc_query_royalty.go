package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/nft/types"
)

// Royalty info

// FeeDenominator queries the FeeDenominator
func (k Keeper) FeeDenominator(c context.Context, request *types.MsgFeeDenominatorRequest) (*types.MsgFeeDenominatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	feeDenominator := k.GetFeeDenominator(ctx)
	return &types.MsgFeeDenominatorResponse{RoyaltyFraction: feeDenominator}, nil
}

// RoyaltyInfo queries the RoyaltyInfo for the class of token
func (k Keeper) RoyaltyInfo(c context.Context, request *types.MsgRoyaltyInfoRequest) (*types.MsgRoyaltyInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if len(request.ClassId) == 0 {
		return nil, types.ErrEmptyClassId
	}
	receiver, amount := k.GetRoyaltyInfo(ctx, request.ClassId, request.TokenId, request.SalePrice)
	return &types.MsgRoyaltyInfoResponse{
		Receiver:      receiver,
		RoyaltyAmount: amount,
	}, nil

}

// DefaultRoyaltyInfo queries the default royalty info for the class
func (k Keeper) DefaultRoyaltyInfo(c context.Context, request *types.MsgDefaultRoyaltyInfoRequest) (*types.MsgDefaultRoyaltyInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if len(request.ClassId) == 0 {
		return nil, types.ErrEmptyClassId
	}
	receiver, amount := k.GetDefaultRoyaltyInfo(ctx, request.ClassId)
	return &types.MsgDefaultRoyaltyInfoResponse{
		Receiver:        receiver,
		RoyaltyFraction: amount,
	}, nil
}

// TokenRoyaltyInfo queries the royalty info for the class of a token
func (k Keeper) TokenRoyaltyInfo(c context.Context, request *types.MsgTokenRoyaltyInfoRequest) (*types.MsgTokenRoyaltyInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if len(request.ClassId) == 0 {
		return nil, types.ErrEmptyClassId
	}
	if len(request.TokenId) == 0 {
		return nil, types.ErrEmptyTokenId
	}
	receiver, amount := k.GetTokenRoyaltyInfo(ctx, request.ClassId, request.TokenId)
	return &types.MsgTokenRoyaltyInfoResponse{
		Receiver:        receiver,
		RoyaltyFraction: amount,
	}, nil
}
