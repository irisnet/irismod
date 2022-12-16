package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidNftID          = sdkerrors.Register(ModuleName, 1, "invalid nft id")
	ErrInvalidClassID        = sdkerrors.Register(ModuleName, 2, "invalid class id")
	ErrInvalidExpires        = sdkerrors.Register(ModuleName, 3, "invalid expires")
	ErrNotApprovedNorOwner   = sdkerrors.Register(ModuleName, 4, "sender is not owner nor approved")
	ErrNotArriveExpires      = sdkerrors.Register(ModuleName, 5, "rental has not expires")
	ErrNotExistentNFT        = sdkerrors.Register(ModuleName, 6, "nft is not existent")
	ErrNotExistentRentalInfo = sdkerrors.Register(ModuleName, 7, "rental info is not existent")
)
