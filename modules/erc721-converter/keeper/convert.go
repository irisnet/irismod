package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/irismod/modules/erc721-converter/contracts"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

// ConvertNFTMint converts a native Cosmos token to an ERC721 token
func (k Keeper) ConvertNFTMint(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertNFT,
	receiver common.Address,
	sender sdk.AccAddress,
) (*types.MsgConvertNFTResponse, error) {

	// Escrow native token on module account
	err := k.nftKeeper.TransferOwnership(
		ctx,
		msg.ClassId,
		msg.TokenId,
		"",
		"",
		"",
		"",
		sender,
		types.ModuleAddress.Bytes(),
	)
	if err != nil {
		return nil, err
	}
	contract := pair.GetERC721Contract()
	erc721Abi := contracts.ERC721PresetMinterPauserAutoIdContract.ABI

	// Mint ERC721 token
	_, err = k.CallEVM(ctx, erc721Abi, types.ModuleAddress, contract, true, "mint", receiver)
	if err != nil {
		return nil, err
	}

	// Check expected receiver balance after transfer
	//k.OwnerOf(ctx, erc721Abi, contract, msg.TokenId)

	return &types.MsgConvertNFTResponse{}, nil
}

// ConvertNFTBurn converts a native Cosmos token to an ERC721 token
func (k Keeper) ConvertNFTBurn() {}

// ConvertERC721Mint converts a erc721 token to an native Cosmos token
func (k Keeper) ConvertERC721Mint() {
}

// ConvertERC721Burn converts a erc721 token to an native Cosmos token
func (k Keeper) ConvertERC721Burn() {}
