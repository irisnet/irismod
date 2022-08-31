package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/nft/types"
)

func (k Keeper) GetDenomInfo(ctx sdk.Context, denomID string) (*types.Denom, error) {
	class, has := k.nk.GetClass(ctx, denomID)
	if !has {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", denomID)
	}

	var denomMetadata types.DenomMetadata
	if err := k.cdc.Unmarshal(class.Data.GetValue(), &denomMetadata); err != nil {
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
		Data:             denomMetadata.Data,
	}, nil
}
