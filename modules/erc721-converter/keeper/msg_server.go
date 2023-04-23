package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

var _ types.MsgServer = &Keeper{}

// ConvertNFT converts a native Cosmos token to an ERC721 token
func (k Keeper) ConvertNFT(goCtx context.Context, msg *types.MsgConvertNFT) (*types.MsgConvertNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Error checked during msg validation
	receiver := common.HexToAddress(msg.Receiver)
	sender := sdk.MustAccAddressFromBech32(msg.Sender)

	// Check if the token pair not exists
	if !k.IsClassRegistered(ctx, msg.ClassId) {
		// Register the token pair
		_, err := k.SaveRegisteredClass(ctx, msg.ClassId)
		if err != nil {
			return nil, err
		}
	}

	pair, err := k.ConvertNFTValidator(ctx, sender, receiver.Bytes(), msg.ClassId, msg.TokenId)
	if err != nil {
		return nil, err
	}

	erc721 := common.HexToAddress(pair.Erc721Address)
	acc := k.evmKeeper.GetAccountWithoutBalance(ctx, erc721)

	if acc == nil || !acc.IsContract() {
		k.DeleteTokenPair(ctx, pair)
		k.Logger(ctx).Debug(
			"deleting selfdestructed token pair from state",
			"contract", pair.Erc721Address,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}
	// Check ownership and execute conversion
	switch {
	case pair.IsNativeNFT():
		// Convert NFT to ERC721

	case pair.IsNativeERC721():

	default:
		return nil, types.ErrUndefinedOwner
	}

	return &types.MsgConvertNFTResponse{}, nil
}

// ConvertERC721 converts an ERC721 token to an native Cosmos token
func (k Keeper) ConvertERC721(goCtx context.Context, msg *types.MsgConvertERC721) (*types.MsgConvertERC721Response, error) {
	return &types.MsgConvertERC721Response{}, nil
}
