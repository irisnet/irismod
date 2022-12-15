package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding rental type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	panic("Fixme!")
}
