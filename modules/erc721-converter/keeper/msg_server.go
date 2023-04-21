package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

var _ types.MsgServer = &Keeper{}

// RegisterDenom registers a native Cosmos token to an ERC721 token
func (k Keeper) RegisterDenom(goCtx context.Context, msg *types.MsgRegisterDenom) (*types.MsgRegisterDenomResponse, error) {
	//ctx := sdk.UnwrapSDKContext(goCtx)
	//// Error checked during msg validation
	//sender := sdk.MustAccAddressFromBech32(msg.Sender)
	//
	//// Check if denomination is already registered
	//if k.IsDenomRegistered(ctx, msg.DenomId) {
	//	return nil, errorsmod.Wrapf(
	//		types.ErrTokenPairAlreadyExists, "coin denomination already registered: %s", msg.DenomId,
	//	)
	//}

	return &types.MsgRegisterDenomResponse{}, nil
}

// RegisterERC721 registers an ERC721 token to an native Cosmos token
func (k Keeper) RegisterERC721(goCtx context.Context, msg *types.MsgRegisterERC721) (*types.MsgRegisterERC721Response, error) {
	return &types.MsgRegisterERC721Response{}, nil
}

// ConvertNFT converts a native Cosmos token to an ERC721 token
func (k Keeper) ConvertNFT(goCtx context.Context, msg *types.MsgConvertNFT) (*types.MsgConvertNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Error checked during msg validation
	receiver := common.HexToAddress(msg.Receiver)
	sender := sdk.MustAccAddressFromBech32(msg.Sender)

	id := k.GetTokenPairID(ctx, msg.DenomId)
	if len(id) == 0 {
		return nil, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "denom '%s' not registered by id", msg.DenomId,
		)
	}
	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return nil, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "denom '%s' not registered", msg.DenomId,
		)
	}

	if !pair.Enabled {
		return nil, errorsmod.Wrapf(
			types.ErrERC721TokenPairDisabled, "minting token '%s' is not enabled by governance", msg.DenomId,
		)
	}

	if !sender.Equals(sdk.AccAddress(receiver.Bytes())) {
		return nil, errorsmod.Wrapf(
			types.ErrERC721TokenPairDisabled, "minting token '%s' is not enabled by governance", msg.DenomId,
		)
	}
	if err := k.SaveRegisteredDenom(ctx, msg.DenomId); err != nil {
		return nil, err
	}

	return &types.MsgConvertNFTResponse{}, nil
}

// ConvertERC721 converts an ERC721 token to an native Cosmos token
func (k Keeper) ConvertERC721(goCtx context.Context, msg *types.MsgConvertERC721) (*types.MsgConvertERC721Response, error) {
	return &types.MsgConvertERC721Response{}, nil
}
