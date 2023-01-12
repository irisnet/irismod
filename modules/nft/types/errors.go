package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidCollection = sdkerrors.Register(ModuleName, 9, "invalid nft collection")
	ErrUnknownCollection = sdkerrors.Register(ModuleName, 10, "unknown nft collection")
	ErrInvalidNFT        = sdkerrors.Register(ModuleName, 11, "invalid nft")
	ErrNFTAlreadyExists  = sdkerrors.Register(ModuleName, 12, "nft already exists")
	ErrUnknownNFT        = sdkerrors.Register(ModuleName, 13, "unknown nft")
	ErrEmptyTokenData    = sdkerrors.Register(ModuleName, 14, "nft data can't be empty")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 15, "unauthorized address")
	ErrInvalidDenom      = sdkerrors.Register(ModuleName, 16, "invalid denom")
	ErrInvalidTokenID    = sdkerrors.Register(ModuleName, 17, "invalid nft id")
	ErrInvalidTokenURI   = sdkerrors.Register(ModuleName, 18, "invalid nft uri")

	ErrNotEnabledRoyalty      = errorsmod.Register(ModuleName, 21, "royalty feature is not enabled")
	ErrNullDefaultRoyaltyInfo = errorsmod.Register(ModuleName, 22, "default royalty in is null")
	ErrNullTokenRoyaltyInfo   = errorsmod.Register(ModuleName, 23, "token royalty in is null")
)
