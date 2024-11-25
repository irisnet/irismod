package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// BankKeeper defines the expected bank keeper (noalias)
type BankKeeper interface {
	SendCoinsFromModuleToAccount(
		ctx context.Context,
		senderModule string,
		recipientAddr sdk.AccAddress,
		amt sdk.Coins,
	) error
	SendCoinsFromModuleToModule(
		ctx context.Context,
		senderModule string,
		recipientModule string,
		amt sdk.Coins,
	) error
	SendCoinsFromAccountToModule(
		ctx context.Context,
		senderAddr sdk.AccAddress,
		recipientModule string,
		amt sdk.Coins,
	) error
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	BurnCoins(ctx context.Context, name string, amt sdk.Coins) error
}

type ValidateLPToken func(ctx context.Context, lpTokenDenom string) error

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI
	SetModuleAccount(ctx context.Context, macc sdk.ModuleAccountI)
}

// DistrKeeper defines the expected distribution keeper (noalias)
type DistrKeeper interface {
	GetFeePool(ctx context.Context) (feePool distrtypes.FeePool, err error)
	SetFeePool(ctx context.Context, feePool distrtypes.FeePool) error
}

// GovKeeper defines the expected gov keeper (noalias)
type GovKeeper interface {
	SubmitProposal(ctx context.Context, messages []sdk.Msg, metadata, title, summary string, proposer sdk.AccAddress, expedited bool) (v1.Proposal, error)
	AddDeposit(ctx context.Context, proposalID uint64, depositorAddr sdk.AccAddress, depositAmount sdk.Coins) (bool, error)
	GetProposal(ctx context.Context, proposalID uint64) (v1.Proposal, error)
	GetGovernanceAccount(ctx context.Context) sdk.ModuleAccountI
}

// CoinswapKeeper defines the expected coinswap keeper (noalias)
type CoinswapKeeper interface {
	ValidatePool(ctx sdk.Context, lptDenom string) error
}
