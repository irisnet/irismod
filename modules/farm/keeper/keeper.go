package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the farm store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler
}
