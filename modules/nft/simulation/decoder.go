package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/irisnet/irismod/modules/nft/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding gov type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.PrefixNFT):
			var nftA, nftB types.BaseNFT
			cdc.MustUnmarshalBinaryBare(kvA.Value, &nftA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &nftB)
			return fmt.Sprintf("%v\n%v", nftA, nftB)
		case bytes.Equal(kvA.Key[:1], types.PrefixOwners):
			idA := types.MustUnMarshalTokenID(cdc, kvA.Value)
			idB := types.MustUnMarshalTokenID(cdc, kvB.Value)
			return fmt.Sprintf("%v\n%v", idA, idB)
		case bytes.Equal(kvA.Key[:1], types.PrefixCollection):
			supplyA := types.MustUnMarshalSupply(cdc, kvA.Value)
			supplyB := types.MustUnMarshalSupply(cdc, kvB.Value)
			return fmt.Sprintf("%d\n%d", supplyA, supplyB)
		case bytes.Equal(kvA.Key[:1], types.PrefixDenom):
			var denomA, denomB types.Denom
			cdc.MustUnmarshalBinaryBare(kvA.Value, &denomA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &denomB)
			return fmt.Sprintf("%v\n%v", denomA, denomB)

		default:
			panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
		}
	}
}
