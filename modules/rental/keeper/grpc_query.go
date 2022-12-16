package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irisnet/irismod/modules/rental/types"
)

var _ types.QueryServer = Keeper{}

// User queries the user of an nft
func (k Keeper) User(goCtx context.Context, msg *types.QueryUserRequest) (*types.QueryUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.ClassId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistentNFT, "%s-%s is not existent", msg.ClassId, msg.NftId)
	}

	rental, exist := k.getRentalInfo(ctx, msg.ClassId, msg.NftId)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistentRentalInfo, "rental info is not existent", msg.ClassId, msg.NftId)
	}

	return &types.QueryUserResponse{User: rental.User}, nil
}

// Expires queries the expires of an nft
func (k Keeper) Expires(goCtx context.Context, msg *types.QueryExpiresRequest) (*types.QueryExpiresResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.ClassId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistentNFT, "%s-%s is not existent", msg.ClassId, msg.NftId)
	}

	rental, exist := k.getRentalInfo(ctx, msg.ClassId, msg.NftId)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistentRentalInfo, "rental info is not existent", msg.ClassId, msg.NftId)
	}

	return &types.QueryExpiresResponse{Expires: rental.Expires}, nil
}

// Expires queries if an nft has the user
// WARNING: it deson't check if this rental has expires
func (k Keeper) HasUser(goCtx context.Context, msg *types.QueryHasUserRequest) (*types.QueryHasUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.ClassId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistentNFT, "%s-%s is not existent", msg.ClassId, msg.NftId)
	}

	_, exist := k.getRentalInfo(ctx, msg.ClassId, msg.NftId)
	return &types.QueryHasUserResponse{HasUser: exist}, nil
}
