package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected bank keeper (noalias)
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context,
		senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context,
		senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type CoinswapKeeper interface {
	ValidatePool(ctx sdk.Context, lpTokenDenom string) error
}
