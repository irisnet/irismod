package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/nft/types"
)

// DefaultRentalPlugin returns a default rental plugin config
func (k Keeper) DefaultRentalPlugin() types.RentalPlugin {
	return types.RentalPlugin{Enabled: false}
}

func (k Keeper) DefaultRentalInfo() types.RentalInfo {
	return types.RentalInfo{
		User:    "",
		Expires: 0,
	}
}

// Rent set or update rental info for an nft.
func (k Keeper) Rent(ctx sdk.Context, denomID, tokenID string, rental types.RentalInfo) error {
	// 1. get rental plugin info
	cfg, err := k.GetRentalPlugin(ctx, denomID)
	if err != nil {
		return err
	}

	// 2. check rental is enabled
	if !cfg.Enabled {
		return sdkerrors.Wrapf(types.ErrRentalOption, "Rental is disabled")
	}

	// 3. expiry must be greater than the current block time.
	if ctx.BlockTime().Unix() >= rental.Expires {
		return sdkerrors.Wrapf(types.ErrInvalidExpiry, "Expiry is (%d)", rental.Expires)
	}

	// 4. construct new nft data info (we have examined its existence)
	var data types.NFTMetadata
	token, _ := k.nk.GetNFT(ctx, denomID, tokenID)
	if err := k.cdc.Unmarshal(token.Data.GetValue(), &data); err != nil {
		return err
	}
	data.RentalInfo = &rental

	newData, err := codectypes.NewAnyWithValue(&data)
	if err != nil {
		return err
	}
	token.Data = newData

	// 5. set rental info with nft update.
	err = k.nk.Update(ctx, token)
	if err != nil {
		return err
	}

	return nil
}

// GetRentalPlugin returns the rental plugin config
func (k Keeper) GetRentalPlugin(ctx sdk.Context, denomID string) (*types.RentalPlugin, error) {
	var r *types.RentalPlugin
	denom, has := k.nk.GetClass(ctx, denomID)
	if !has {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", denomID)
	}

	var denomMetadata types.DenomMetadata
	if err := k.cdc.Unmarshal(denom.Data.GetValue(), &denomMetadata); err != nil {
		return nil, err
	}

	r = denomMetadata.RentalPlugin
	return r, nil
}

// GetRentalInfo returns the rental info for an nft.
func (k Keeper) GetRentalInfo(ctx sdk.Context,
	denomID, tokenID string) (*types.RentalInfo, error) {
	var r *types.RentalInfo
	token, exist := k.nk.GetNFT(ctx, denomID, tokenID)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrInvalidNFT, "token ID %s not exists", tokenID)
	}

	var nftMetadata types.NFTMetadata
	if err := k.cdc.Unmarshal(token.Data.GetValue(), &nftMetadata); err != nil {
		return nil, err
	}
	r = nftMetadata.RentalInfo
	return r, nil
}
