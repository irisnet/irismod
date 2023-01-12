package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irisnet/irismod/modules/nft/types"
)

func (k Keeper) UserDataToDenomPlugin(ctx sdk.Context, data string) *types.DenomPlugin {

	// 1. Unmarshal denomPlugin
	denomPlugin := new(types.DenomPlugin)
	if err := k.cdc.Unmarshal([]byte(data), denomPlugin); err != nil {
		return nil
	}
	return denomPlugin
}

func (k Keeper) GetDenomMetadata(ctx sdk.Context, denomId string) (*types.DenomMetadata, error) {
	class, has := k.nk.GetClass(ctx, denomId)
	if !has {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", denomId)
	}

	denomMetadata := new(types.DenomMetadata)

	if err := k.cdc.Unmarshal(class.Data.GetValue(), denomMetadata); err != nil {
		return nil, err
	}
	return denomMetadata, nil
}

func (k Keeper) GetNftMetadata(ctx sdk.Context, denomId, tokenId string) (*types.NFTMetadata, error) {
	token, exist := k.nk.GetNFT(ctx, denomId, tokenId)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "nft ID %s not exists", tokenId)
	}
	nftMetadata := new(types.NFTMetadata)
	if err := k.cdc.Unmarshal(token.Data.GetValue(), nftMetadata); err != nil {
		return nil, err
	}
	return nftMetadata, nil
}
