package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

// SaveRegisteredClass saves the registered denom to the store
func (k Keeper) SaveRegisteredClass(ctx sdk.Context, classId string) (common.Address, error) {

	denomInfo, err := k.irisModNFTKeeper.GetDenomInfo(ctx, classId)
	if err != nil {
		return common.Address{}, err
	}

	if k.IsClassRegistered(ctx, denomInfo.Id) {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrTokenPairAlreadyExists,
			"denom metadata already registered %s", denomInfo.Name,
		)
	}

	class := nfttypes.Class{
		Id:     classId,
		Name:   denomInfo.Name,
		Symbol: denomInfo.Symbol,
	}
	// Deployed contract address is used as the key to save the denom
	addr, err := k.DeployERC721Contract(ctx, class)
	if err != nil {
		return common.Address{}, errorsmod.Wrap(
			err, "failed to create wrapped coin denom metadata for ERC721",
		)
	}

	pair := types.NewTokenPair(addr, denomInfo.Id, types.OWNER_MODULE)
	k.SetTokenPair(ctx, pair)
	k.SetClassMap(ctx, pair.ClassId, pair.GetID())
	k.SetERC721Map(ctx, common.HexToAddress(pair.Erc721Address), pair.GetID())
	return addr, nil
}

// SaveRegisteredERC721 saves the registered ERC721 to the store
func (k Keeper) SaveRegisteredERC721(ctx sdk.Context, contract common.Address) (string, error) {
	classId := types.CreateClass(contract.String())
	if _, found := k.nftKeeper.GetClass(ctx, classId); found {
		return "", errorsmod.Wrap(
			types.ErrInternalTokenPair,
			"denom metadata already registered",
		)
	}

	if k.IsClassRegistered(ctx, classId) {
		return "", errorsmod.Wrap(
			types.ErrInternalTokenPair,
			"denom metadata already registered",
		)
	}
	erc721Data, err := k.QueryERC721(ctx, contract)
	if err != nil {
		return "", err
	}

	if err := k.irisModNFTKeeper.SaveDenom(
		ctx,
		classId,
		erc721Data.Name,
		"",
		erc721Data.Symbol,
		types.ModuleAddress.Bytes(),
		false,
		false,
		"",
		"",
		"",
		"",
	); err != nil {
		return "", errorsmod.Wrap(
			err, "failed to create wrapped coin denom metadata for ERC721",
		)
	}

	pair := types.NewTokenPair(contract, classId, types.OWNER_MODULE)
	k.SetTokenPair(ctx, pair)
	k.SetClassMap(ctx, pair.ClassId, pair.GetID())
	k.SetERC721Map(ctx, common.HexToAddress(pair.Erc721Address), pair.GetID())
	return classId, nil
}
