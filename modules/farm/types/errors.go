package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// farm module sentinel errors
var (
	ErrExpiredHeight  = sdkerrors.Register(ModuleName, 2, "Expired block height")
	ErrInvalidLPToken = sdkerrors.Register(ModuleName, 3, "Invalid lp token denom")
	ErrNotMatch       = sdkerrors.Register(ModuleName, 4, "The data does not match")
	ErrExpiredPool    = sdkerrors.Register(ModuleName, 5, "The farm pool has expired")
	ErrNotExistPool   = sdkerrors.Register(ModuleName, 6, "The farm pool is not exist")
	ErrInvalidOperate = sdkerrors.Register(ModuleName, 7, "Invalid operate")
	ErrNotExistFarmer = sdkerrors.Register(ModuleName, 7, "The farmer is not exist")
)
