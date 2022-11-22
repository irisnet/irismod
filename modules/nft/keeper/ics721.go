package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"

	"github.com/irisnet/irismod/modules/nft/types"
)

var _ NFTKeeper = (*ICS721Keeper)(nil)

// NFTKeeper is a inteface for ics721 protocol
type NFTKeeper interface {
	SaveClass(ctx sdk.Context, class nft.Class) error
	Mint(ctx sdk.Context, token nft.NFT, receiver sdk.AccAddress) error
	Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error
	Burn(ctx sdk.Context, classID string, nftID string) error

	GetOwner(ctx sdk.Context, classID string, nftID string) sdk.AccAddress
	HasClass(ctx sdk.Context, classID string) bool
	GetClass(ctx sdk.Context, classID string) (nft.Class, bool)
	GetNFT(ctx sdk.Context, classID, nftID string) (nft.NFT, bool)
}

// AccountKeeper defines the contract required for account APIs.
type AccountKeeper interface {
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Set an account in the store.
	SetAccount(sdk.Context, authtypes.AccountI)
	GetModuleAddress(name string) sdk.AccAddress
}

// ICS721Keeper defines the ICS721 Keeper
type ICS721Keeper struct {
	cdc codec.Codec
	nk  nftkeeper.Keeper
	ak  AccountKeeper
}

// NewISC721Keeper creates a new ics721 Keeper instance
func (k Keeper) NewISC721Keeper(ak AccountKeeper) ICS721Keeper {
	return ICS721Keeper{
		cdc: k.cdc,
		nk:  k.nk,
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
		Creator:        moduleAddress.String(),
		MintRestricted: false,
		// NOTICE: UpdateRestricted is set to false to prevent nft from being edited on the destination chain,
		// but when transferring to the original chain, the edited information is lost
		UpdateRestricted: false,
	}
	metadata, err := codectypes.NewAnyWithValue(denomMetadata)
	if err != nil {
		return err
	}
	class.Data = metadata
	return ik.nk.SaveClass(ctx, class)
}

// Mint implement the method of ICS721Keeper.Mint
func (ik ICS721Keeper) Mint(ctx sdk.Context, token nft.NFT, receiver sdk.AccAddress) error {
	metadata, err := codectypes.NewAnyWithValue(&types.NFTMetadata{})
	if err != nil {
		return err
	}
	token.Data = metadata
	return ik.nk.Mint(ctx, token, receiver)
}

// Transfer implement the method of ICS721Keeper.Transfer
func (ik ICS721Keeper) Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error {
	return ik.nk.Transfer(ctx, classID, nftID, receiver)
}

// Burn implement the method of ICS721Keeper.Burn
func (ik ICS721Keeper) Burn(ctx sdk.Context, classID string, nftID string) error {
	return ik.nk.Burn(ctx, classID, nftID)
}

// GetOwner implement the method of ICS721Keeper.GetOwner
func (ik ICS721Keeper) GetOwner(ctx sdk.Context, classID string, nftID string) sdk.AccAddress {
	return ik.nk.GetOwner(ctx, classID, nftID)
}

// HasClass implement the method of ICS721Keeper.HasClass
func (ik ICS721Keeper) HasClass(ctx sdk.Context, classID string) bool {
	return ik.nk.HasClass(ctx, classID)
}

// GetClass implement the method of ICS721Keeper.GetClass
func (ik ICS721Keeper) GetClass(ctx sdk.Context, classID string) (nft.Class, bool) {
	return ik.nk.GetClass(ctx, classID)
}

// GetNFT implement the method of ICS721Keeper.GetNFT
func (ik ICS721Keeper) GetNFT(ctx sdk.Context, classID, nftID string) (nft.NFT, bool) {
	return ik.nk.GetNFT(ctx, classID, nftID)
}
