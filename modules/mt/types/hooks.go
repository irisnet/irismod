package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var Protocol = "mt"

type Hook interface {
	BeforeTokenMint(ctx sdk.Context, protocol, denomID string, sender, receiver sdk.AccAddress) error
	BeforeTokenTransfer(ctx sdk.Context, protocol, denomID, tokenID string, sender, receiver sdk.AccAddress) error
	BeforeTokenBurn(ctx sdk.Context, protocol, denomID, tokenID string, sender sdk.AccAddress) error
	BeforeTokenEdit(ctx sdk.Context, protocol, denomID, tokenID string, sender sdk.AccAddress) error
	BeforeDenomTransfer(ctx sdk.Context, protocol, denomID string, sender sdk.AccAddress) error
}
