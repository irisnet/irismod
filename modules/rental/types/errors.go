package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	// fixme
	ErrInvalidRentalInfos = sdkerrors.Register(ModuleName, 1, "invalid rental infos")
)
