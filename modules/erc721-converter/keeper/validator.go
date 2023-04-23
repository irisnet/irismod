package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

// ConvertNFTValidator is the validator for ConvertNFT
func (k Keeper) ConvertNFTValidator(
	ctx sdk.Context,
	sender, receiver sdk.AccAddress,
	classId string,
	nftId string,
) (types.TokenPair, error) {

	id := k.GetTokenPairID(ctx, classId)
	if len(id) == 0 {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "class '%s' not registered by id", classId,
		)
	}
	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrTokenPairNotFound, "class '%s' not registered", classId,
		)
	}

	if !pair.Enabled {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrERC721TokenPairDisabled, "minting token '%s' is not enabled by governance", classId,
		)
	}

	// Check if the sender has the token

	if k.nftKeeper.Authorize(ctx, classId, nftId, sender) != nil {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrUnauthorized, "sender '%s' does not own token '%s'", sender, nftId,
		)
	}
	// check sender equals receiver
	if !sender.Equals(receiver) {
		return types.TokenPair{}, errorsmod.Wrapf(
			types.ErrUnauthorized, "sender '%s' is not equal to receiver '%s'", sender, receiver,
		)
	}

	return pair, nil
}
