package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MT non fungible token interface
type MT interface {
	GetID() string
	GetName() string
	GetOwner() sdk.AccAddress
	GetURI() string
	GetURIHash() string
	GetData() string
}
