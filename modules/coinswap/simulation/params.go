package simulation

import (
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/irisnet/irismod/modules/coinswap/types"
	"math/rand"
)

const (
	keyFee = "Fee"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyFee,
			func(r *rand.Rand) string {
				return RandomDec(r).String()
			},
		),
	}
}
