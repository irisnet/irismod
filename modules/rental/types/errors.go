package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidNftID   = sdkerrors.Register(ModuleName, 1, "invalid nft id")
	ErrInvalidClassID = sdkerrors.Register(ModuleName, 2, "invalid class id")
	ErrInvalidExpires = sdkerrors.Register(ModuleName, 3, "invalid expires")
)
