package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MT non fungible token interface
type MT interface {
	GetID() string
	GetOwner() sdk.AccAddress
	GetData() []byte
	GetSupply() uint64
}
