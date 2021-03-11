package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// HTLC module sentinel errors
var (
	ErrInvalidID                   = sdkerrors.Register(ModuleName, 2, "invalid htlc id")
	ErrInvalidTimeLock             = sdkerrors.Register(ModuleName, 3, "invalid time lock")
	ErrInvalidSecret               = sdkerrors.Register(ModuleName, 4, "invalid secret")
	ErrHTLCExists                  = sdkerrors.Register(ModuleName, 5, "htlc already exists")
	ErrUnknownHTLC                 = sdkerrors.Register(ModuleName, 6, "unknown htlc")
	ErrHTLCNotOpen                 = sdkerrors.Register(ModuleName, 7, "htlc not open")
	ErrHTLCNotExpired              = sdkerrors.Register(ModuleName, 8, "htlc not expired")
	ErrAssetNotSupported           = sdkerrors.Register(ModuleName, 9, "asset not found")
	ErrAssetNotActive              = sdkerrors.Register(ModuleName, 10, "asset is currently inactive")
	ErrInvalidAccount              = sdkerrors.Register(ModuleName, 11, "atomic swap has invalid account")
	ErrInvalidAmount               = sdkerrors.Register(ModuleName, 12, "amount must contain exactly one coin")
	ErrInvalidSwapDirection        = sdkerrors.Register(ModuleName, 13, "invalid swap direction")
	ErrInsufficientAmount          = sdkerrors.Register(ModuleName, 15, "amount cannot cover the deputy fixed fee")
	ErrExceedsSupplyLimit          = sdkerrors.Register(ModuleName, 16, "asset supply over limit")
	ErrExceedsTimeBasedSupplyLimit = sdkerrors.Register(ModuleName, 17, "asset supply over limit for current time period")
	ErrInvalidCurrentSupply        = sdkerrors.Register(ModuleName, 18, "supply decrease puts current asset supply below 0")
	ErrInvalidIncomingSupply       = sdkerrors.Register(ModuleName, 19, "supply decrease puts incoming asset supply below 0")
	ErrInvalidOutgoingSupply       = sdkerrors.Register(ModuleName, 20, "supply decrease puts outgoing asset supply below 0")
	ErrExceedsAvailableSupply      = sdkerrors.Register(ModuleName, 21, "outgoing swap exceeds total available supply")
	ErrAssetSupplyNotFound         = sdkerrors.Register(ModuleName, 22, "asset supply not found in store")
)
