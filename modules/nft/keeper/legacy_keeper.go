package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"mods.irisnet.org/modules/nft/exported"
	"mods.irisnet.org/modules/nft/types"
)

// LegacyKeeper is an alias of Keeper
type LegacyKeeper struct {
	nk Keeper
}

// NewLegacyKeeper creates a new instance of the NFT Keeper
func NewLegacyKeeper(nk Keeper) LegacyKeeper {
	return LegacyKeeper{nk}
}

// IssueDenom issues a denom according to the given params
func (n LegacyKeeper) IssueDenom(ctx sdk.Context,
	id, name, schema, symbol string,
	creator sdk.AccAddress,
	mintRestricted, updateRestricted bool) error {
	return n.nk.IssueDenom(ctx, id, name, schema, symbol, creator, mintRestricted, updateRestricted, types.DoNotModify, types.DoNotModify, types.DoNotModify, types.DoNotModify)
}

// MintNFT mints a new NFT
func (n LegacyKeeper) MintNFT(ctx sdk.Context,
	denomID, tokenID, tokenNm, tokenURI, tokenData string,
	owner sdk.AccAddress) error {
	return n.nk.MintNFT(ctx, denomID, tokenID, tokenNm, tokenURI, "", tokenData, owner)
}

// TransferOwner transfers the ownership of an NFT
func (n LegacyKeeper) TransferOwner(ctx sdk.Context,
	denomID, tokenID, tokenNm, tokenURI, tokenData string,
	srcOwner, dstOwner sdk.AccAddress) error {
	return n.nk.TransferOwner(ctx, denomID, tokenID, tokenNm, tokenURI, types.DoNotModify, tokenData, srcOwner, dstOwner)
}

// BurnNFT burns an NFT
func (n LegacyKeeper) BurnNFT(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) error {
	return n.nk.BurnNFT(ctx, denomID, tokenID, owner)
}

// GetNFT returns an NFT
func (n LegacyKeeper) GetNFT(ctx sdk.Context, denomID, tokenID string) (nft exported.NFT, err error) {
	return n.nk.GetNFT(ctx, denomID, tokenID)
}

// GetDenom returns a denom
func (n LegacyKeeper) GetDenom(ctx sdk.Context, id string) (denom types.Denom, found bool) {
	return n.nk.GetDenom(ctx, id)
}
