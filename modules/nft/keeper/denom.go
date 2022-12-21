package keeper

import (
	"encoding/json"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/irisnet/irismod/modules/nft/types"
)

// SaveDenom issues a denom according to the given params
func (k Keeper) SaveDenom(ctx sdk.Context, id,
	name,
	schema,
	symbol string,
	creator sdk.AccAddress,
	mintRestricted,
	updateRestricted bool,
	description,
	uri,
	uriHash,
	data string,
) error {

	// make sure that plugin has default value if failed to convert input data
	denomPlugin := k.DenomDataToDenomPlugin(data)

	denomMetadata := &types.DenomMetadata{
		Creator:          creator.String(),
		Schema:           schema,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
		Data:             "",
		RentalPlugin:     denomPlugin.RentalPlugin,
	}
	metadata, err := codectypes.NewAnyWithValue(denomMetadata)
	if err != nil {
		return err
	}
	return k.nk.SaveClass(ctx, nft.Class{
		Id:          id,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         uri,
		UriHash:     uriHash,
		Data:        metadata,
	})
}

// TransferDenomOwner transfers the ownership of the given denom to the new owner
func (k Keeper) TransferDenomOwner(
	ctx sdk.Context,
	denomID string,
	srcOwner,
	dstOwner sdk.AccAddress,
) error {
	denom, err := k.GetDenomInfo(ctx, denomID)
	if err != nil {
		return err
	}

	// authorize
	if srcOwner.String() != denom.Creator {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to transfer denom %s", srcOwner.String(), denomID)
	}

	denomMetadata := &types.DenomMetadata{
		Creator:          dstOwner.String(),
		Schema:           denom.Schema,
		MintRestricted:   denom.MintRestricted,
		UpdateRestricted: denom.UpdateRestricted,
		Data:             denom.Data,
	}
	data, err := codectypes.NewAnyWithValue(denomMetadata)
	if err != nil {
		return err
	}
	class := nft.Class{
		Id:     denom.Id,
		Name:   denom.Name,
		Symbol: denom.Symbol,
		Data:   data,

		Description: denom.Description,
		Uri:         denom.Uri,
		UriHash:     denom.UriHash,
	}

	return k.nk.UpdateClass(ctx, class)
}

// GetDenomInfo return the denom information
func (k Keeper) GetDenomInfo(ctx sdk.Context, denomID string) (*types.Denom, error) {
	class, has := k.nk.GetClass(ctx, denomID)
	if !has {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", denomID)
	}

	var denomMetadata types.DenomMetadata
	if err := k.cdc.Unmarshal(class.Data.GetValue(), &denomMetadata); err != nil {
		return nil, err
	}

	denomPlugin := k.DenomMetadataToDenomPlugin(denomMetadata)
	data, err := json.Marshal(denomPlugin)
	if err != nil {
		return nil, err
	}

	return &types.Denom{
		Id:               class.Id,
		Name:             class.Name,
		Schema:           denomMetadata.Schema,
		Creator:          denomMetadata.Creator,
		Symbol:           class.Symbol,
		MintRestricted:   denomMetadata.MintRestricted,
		UpdateRestricted: denomMetadata.UpdateRestricted,
		Description:      class.Description,
		Uri:              class.Uri,
		UriHash:          class.UriHash,
		Data:             string(data),
	}, nil
}

// HasDenom determine whether denom exists
func (k Keeper) HasDenom(ctx sdk.Context, denomID string) bool {
	return k.nk.HasClass(ctx, denomID)
}

// DefaultDenomPlugin returns a default DenomPlugin
func (k Keeper) DefaultDenomPlugin() types.DenomPlugin {
	rentalPlugin := k.DefaultRentalPlugin()
	return types.DenomPlugin{
		RentalPlugin: &rentalPlugin,
	}
}

// DenomDataToDenomPlugin converts user denom data to denom plugin struct
func (k Keeper) DenomDataToDenomPlugin(data string) types.DenomPlugin {
	var denomPlugin types.DenomPlugin
	if err := json.Unmarshal([]byte(data), &denomPlugin); err != nil {
		denomPlugin = k.DefaultDenomPlugin()
	}
	return denomPlugin
}

// DenomMetadataToDenomPlugin extracts plugin config from denom metadata
func (k Keeper) DenomMetadataToDenomPlugin(denomMetadata types.DenomMetadata) types.DenomPlugin {
	denomPlugin := types.DenomPlugin{
		RentalPlugin: denomMetadata.RentalPlugin,
	}
	return denomPlugin
}
