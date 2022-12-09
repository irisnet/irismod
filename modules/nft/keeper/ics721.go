package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"

	"github.com/irisnet/irismod/modules/nft/types"
)

// AccountKeeper defines the contract required for account APIs.
type AccountKeeper interface {
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Set an account in the store.
	SetAccount(sdk.Context, authtypes.AccountI)
	GetModuleAddress(name string) sdk.AccAddress
}

// ICS721Keeper defines the ICS721 Keeper
type ICS721Keeper struct {
	nftkeeper.Keeper
	k   Keeper
	cdc codec.Codec
	ak  AccountKeeper
}

// NewISC721Keeper creates a new ics721 Keeper instance
func NewISC721Keeper(k Keeper, ak AccountKeeper) ICS721Keeper {
	return ICS721Keeper{
		Keeper: k.nk,
		k:      k,
		cdc:    k.cdc,
		ak:     ak,
	}
}

// SaveClass implement the method of ICS721Keeper.SaveClass
func (ik ICS721Keeper) SaveClass(ctx sdk.Context, class nft.Class) error {
	moduleAddress := ik.ak.GetModuleAddress(types.ModuleName)
	if moduleAddress == nil {
		moduleAddress = authtypes.NewModuleAddress(types.ModuleName)
		acc := ik.ak.NewAccountWithAddress(ctx, moduleAddress)
		ik.ak.SetAccount(ctx, acc)
	}

	//TODO Because ics721 protocol is not currently supported the transfer of classData ,
	// the original classData will be ignored now
	var denomMetadata = &types.DenomMetadata{
		Creator:          moduleAddress.String(),
		MintRestricted:   false,
		UpdateRestricted: false,
	}
	metadata, err := codectypes.NewAnyWithValue(denomMetadata)
	if err != nil {
		return err
	}
	class.Data = metadata
	return ik.Keeper.SaveClass(ctx, class)
}

// Transfer implement the method of ICS721Keeper.Update
func (ik ICS721Keeper) Update(ctx sdk.Context, token nft.NFT) error {
	denomInfo, err := ik.k.GetDenomInfo(ctx, token.ClassId)
	if err != nil {
		return err
	}

	if denomInfo.UpdateRestricted {
		// if true , nobody can update the NFT under this denom
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "nobody can update the NFT under this denom %s", token.ClassId)
	}
	return ik.Keeper.Update(ctx, token)
}
