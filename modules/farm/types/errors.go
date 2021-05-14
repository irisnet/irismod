package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// farm module sentinel errors
var (
	ErrExpiredHeight  = sdkerrors.Register(ModuleName, 2, "expired block height")
	ErrInvalidLPToken = sdkerrors.Register(ModuleName, 3, "invalid lp token denom")
	ErrNotMatch       = sdkerrors.Register(ModuleName, 3, "The length of the arrays does not match")
)
