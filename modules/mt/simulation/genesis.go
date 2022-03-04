package simulation

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"

	"github.com/irisnet/irismod/modules/mt/types"
)

const (
	kitties = "kitties"
	doggos  = "doggos"
)

// RandomizedGenState generates a random GenesisState for mt
func RandomizedGenState(simState *module.SimulationState) {
	collections := types.NewCollections(
		types.NewCollection(
			types.Denom{
				Id:   doggos,
				Name: doggos,
			},
			types.MTs{types.MT{
				Id:     RandMTID(simState.Rand, 1, 10),
				Supply: 100,
				Data:   nil,
			}, types.MT{
				Id:     RandMTID(simState.Rand, 1, 10),
				Supply: 100,
				Data:   nil,
			}},
		),
		types.NewCollection(
			types.Denom{
				Id:   kitties,
				Name: kitties,
			},
			types.MTs{types.MT{
				Id:     RandMTID(simState.Rand, 1, 10),
				Supply: 100,
				Data:   nil,
			}},
		),
	)

	mtGenesis := types.NewGenesisState(collections, []types.Owner{})

	bz, err := json.MarshalIndent(mtGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mtGenesis)
}

func RandMTID(r *rand.Rand, min, max int) string {
	n := simtypes.RandIntBetween(r, min, max)
	id := simtypes.RandStringOfLength(r, n)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(id)))
}
