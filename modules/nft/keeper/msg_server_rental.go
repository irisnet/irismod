package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/irisnet/irismod/modules/nft/types"
)

// SetUser set a user and expires time for an existent nft
func (k Keeper) SetUser(goCtx context.Context, msg *types.MsgSetUser) (*types.MsgSetUserResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	user, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// nft must exist
	if exist := k.HasNFT(ctx, msg.ClassId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "%s-%s is not existent", msg.ClassId, msg.NftId)
	}

	// sender must own or be approved for this nft
	if owner := k.nk.GetOwner(ctx, msg.ClassId, msg.NftId); !owner.Equals(sender) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "%s is not owner of the nft", msg.Sender)
	}

	if err := k.Rent(ctx, types.RentalInfo{
		User:    user.String(),
		ClassId: msg.ClassId,
		NftId:   msg.NftId,
		Expires: msg.Expires,
	}); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetUser,
			sdk.NewAttribute(types.AttributeKeyDenomID, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenID, msg.NftId),
			sdk.NewAttribute(types.AttributeKeyExpires, strconv.FormatInt(msg.Expires, 10)),
			sdk.NewAttribute(types.AttributeKeyUser, msg.User),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgSetUserResponse{}, nil
}
