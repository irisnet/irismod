package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

// SaveRegisteredDenom saves the registered denom to the store
func (k Keeper) SaveRegisteredDenom(ctx sdk.Context, denom string) error {

	// Deployed contract address is used as the key to save the denom
	addr, err := k.DeployERC721Contract(ctx, denom)
	if err != nil {
		return errorsmod.Wrap(
			err, "failed to create wrapped coin denom metadata for ERC20",
		)
	}

	pair := types.NewTokenPair(addr, denom, types.OWNER_MODULE)
	k.SetTokenPair(ctx, pair)
	k.SetDenomMap(ctx, pair.Denom, pair.GetID())
	k.SetERC721Map(ctx, common.HexToAddress(pair.Erc721Address), pair.GetID())
	return nil
}

// SaveRegisteredERC721 saves the registered ERC721 to the store
func (k Keeper) SaveRegisteredERC721(ctx sdk.Context, contract common.Address) error {

	pair := types.NewTokenPair(contract, "", types.OWNER_MODULE)
	k.SetTokenPair(ctx, pair)
	k.SetDenomMap(ctx, pair.Denom, pair.GetID())
	k.SetERC721Map(ctx, common.HexToAddress(pair.Erc721Address), pair.GetID())
	return nil
}
