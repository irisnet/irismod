package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/mt/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey storetypes.StoreKey // Unexposed key to access store from sdk.Context
	cdc      codec.Codec
}

// NewKeeper creates a new instance of the MT Keeper
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("irismod/%s", types.ModuleName))
}

// IssueDenom issues a denom according to the given params
func (k Keeper) IssueDenom(ctx sdk.Context,
	id, name, schema, symbol string,
	creator sdk.AccAddress,
	mintRestricted, updateRestricted bool,
	description, uri, uriHash, data string,
) error {
	return k.SetDenom(ctx, types.Denom{
		Id:               id,
		Name:             name,
		Schema:           schema,
		Creator:          creator.String(),
		Symbol:           symbol,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
		Description:      description,
		Uri:              uri,
		UriHash:          uriHash,
		Data:             data,
	})
}

// MintMT mints an MT and manages the MT's existence within Collections and Owners
func (k Keeper) MintMT(
	ctx sdk.Context, denomID, tokenID, tokenNm,
	tokenURI, uriHash, tokenData string, owner sdk.AccAddress,
) error {
	if k.HasMT(ctx, denomID, tokenID) {
		return sdkerrors.Wrapf(types.ErrMTAlreadyExists, "MT %s already exists in collection %s", tokenID, denomID)
	}

	k.setMT(
		ctx, denomID,
		types.NewBaseMT(
			tokenID,
			tokenNm,
			owner,
			tokenURI,
			uriHash,
			tokenData,
		),
	)
	k.setOwner(ctx, denomID, tokenID, owner)
	k.increaseSupply(ctx, denomID)

	return nil
}

// EditMT updates an already existing MT
func (k Keeper) EditMT(
	ctx sdk.Context, denomID, tokenID, tokenNm,
	tokenURI, tokenURIHash, tokenData string, owner sdk.AccAddress,
) error {
	denom, found := k.GetDenom(ctx, denomID)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", denomID)
	}

	if denom.UpdateRestricted {
		// if true , nobody can update the MT under this denom
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "nobody can update the MT under this denom %s", denom.Id)
	}

	// just the owner of MT can edit
	mt, err := k.Authorize(ctx, denomID, tokenID, owner)
	if err != nil {
		return err
	}

	if types.Modified(tokenNm) {
		mt.Name = tokenNm
	}

	if types.Modified(tokenURI) {
		mt.URI = tokenURI
	}

	if types.Modified(tokenURIHash) {
		mt.UriHash = tokenURIHash
	}

	if types.Modified(tokenData) {
		mt.Data = tokenData
	}

	k.setMT(ctx, denomID, mt)

	return nil
}

// TransferOwner transfers the ownership of the given MT to the new owner
func (k Keeper) TransferOwner(
	ctx sdk.Context, denomID, tokenID, tokenNm, tokenURI, tokenURIHash,
	tokenData string, srcOwner, dstOwner sdk.AccAddress,
) error {
	denom, found := k.GetDenom(ctx, denomID)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", denomID)
	}

	mt, err := k.Authorize(ctx, denomID, tokenID, srcOwner)
	if err != nil {
		return err
	}

	mt.Owner = dstOwner.String()

	if denom.UpdateRestricted && (types.Modified(tokenNm) || types.Modified(tokenURI) || types.Modified(tokenData)) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "It is restricted to update MT under this denom %s", denom.Id)
	}

	if types.Modified(tokenNm) {
		mt.Name = tokenNm
	}
	if types.Modified(tokenURI) {
		mt.URI = tokenURI
	}
	if types.Modified(tokenURIHash) {
		mt.UriHash = tokenURIHash
	}
	if types.Modified(tokenData) {
		mt.Data = tokenData
	}

	k.setMT(ctx, denomID, mt)
	k.swapOwner(ctx, denomID, tokenID, srcOwner, dstOwner)
	return nil
}

// BurnMT deletes a specified MT
func (k Keeper) BurnMT(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) error {
	if !k.HasDenomID(ctx, denomID) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", denomID)
	}

	mt, err := k.Authorize(ctx, denomID, tokenID, owner)
	if err != nil {
		return err
	}

	k.deleteMT(ctx, denomID, mt)
	k.deleteOwner(ctx, denomID, tokenID, owner)
	k.decreaseSupply(ctx, denomID)

	return nil
}

// TransferDenomOwner transfers the ownership of the given denom to the new owner
func (k Keeper) TransferDenomOwner(
	ctx sdk.Context, denomID string, srcOwner, dstOwner sdk.AccAddress,
) error {
	denom, found := k.GetDenom(ctx, denomID)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", denomID)
	}

	// authorize
	if srcOwner.String() != denom.Creator {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to transfer denom %s", srcOwner.String(), denomID)
	}

	denom.Creator = dstOwner.String()

	err := k.UpdateDenom(ctx, denom)
	if err != nil {
		return err
	}

	return nil
}
