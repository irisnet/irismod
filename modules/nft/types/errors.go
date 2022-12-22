package types

import (
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

	// Rental Plugin errors
	ErrRentalPluginNotExistent = sdkerrors.Register(ModuleName, 30, "rental plugin is not existent")
	ErrRentalPluginDisabled    = sdkerrors.Register(ModuleName, 31, "rental plugin is disabled")
	ErrRentalExpiryInvalid     = sdkerrors.Register(ModuleName, 32, "rental expiry is invalid")
	ErrRentalInfoNotExsitent   = sdkerrors.Register(ModuleName, 33, "rental info is not existent")
)
