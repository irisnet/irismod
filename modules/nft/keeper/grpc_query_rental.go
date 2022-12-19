package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irisnet/irismod/modules/nft/types"
)

var _ types.QueryServer = Keeper{}

// User queries the user of an nft
func (k Keeper) UserOf(goCtx context.Context, msg *types.QueryUserOfRequest) (*types.QueryUserOfResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.ClassId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "%s-%s is not existent", msg.ClassId, msg.NftId)
	}

	rental, exist := k.GetRentalInfo(ctx, msg.ClassId, msg.NftId)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistentRentalInfo, "rental info is not existent", msg.ClassId, msg.NftId)
	}

	return &types.QueryUserOfResponse{User: rental.User}, nil
}

// Expires queries the expires of an nft
func (k Keeper) UserExpires(goCtx context.Context, msg *types.QueryUserExpiresRequest) (*types.QueryUserExpiresResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.ClassId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "%s-%s is not existent", msg.ClassId, msg.NftId)
	}

	rental, exist := k.GetRentalInfo(ctx, msg.ClassId, msg.NftId)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistentRentalInfo, "rental info is not existent", msg.ClassId, msg.NftId)
	}

	return &types.QueryUserExpiresResponse{Expires: rental.Expires}, nil
}

// HasUser queries if an nft has the user
// WARNING: it doesn't check if this rental has expires
func (k Keeper) HasUser(goCtx context.Context, msg *types.QueryHasUserRequest) (*types.QueryHasUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.ClassId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "%s-%s is not existent", msg.ClassId, msg.NftId)
	}

	_, exist := k.GetRentalInfo(ctx, msg.ClassId, msg.NftId)
	return &types.QueryHasUserResponse{HasUser: exist}, nil
}
