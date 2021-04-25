package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/irisnet/irismod/modules/coinswap/types"
	tokentypes "github.com/irisnet/irismod/modules/token/types"
	"math/rand"
	"strings"
	"sync"
)

var once sync.Once
var tokens []tokentypes.Token

// RandomizedGenState generates a random GenesisState for coinswap
func RandomizedGenState(simState *module.SimulationState) {
	var fee sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, "fee", &fee, simState.Rand,
		func(r *rand.Rand) {
			fee = sdk.NewDecWithPrec(r.Int63n(3),3)
			once.Do(func() {
				for i := 0; i < 5; i++ {
					tokens = append(tokens, randToken(r, simState.Accounts))
				}
				tokens = append(tokens, tokentypes.GetNativeToken())
				tokenGenesis := tokentypes.NewGenesisState(
					tokentypes.DefaultParams(),
					tokens,
				)
				simState.GenState[tokentypes.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenGenesis)
			})

		},
	)
	genesis := types.NewGenesisState(types.NewParams(fee),tokentypes.GetNativeToken().MinUnit)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(genesis)
}


func randToken(r *rand.Rand, accs []simtypes.Account) tokentypes.Token {
	symbol := randStringBetween(r, tokentypes.MinimumSymbolLen, tokentypes.MaximumSymbolLen)
	minUint := randStringBetween(r, tokentypes.MinimumMinUnitLen, tokentypes.MaximumMinUnitLen)
	name := randStringBetween(r, 1, tokentypes.MaximumNameLen)
	scale := simtypes.RandIntBetween(r, 1, int(tokentypes.MaximumScale))
	initialSupply := r.Int63n(int64(tokentypes.MaximumInitSupply))
	maxSupply := 2 * initialSupply
	simAccount, _ := simtypes.RandomAcc(r, accs)

	return tokentypes.Token{
		Symbol:        strings.ToLower(symbol),
		Name:          name,
		Scale:         uint32(scale),
		MinUnit:       strings.ToLower(minUint),
		InitialSupply: uint64(initialSupply),
		MaxSupply:     uint64(maxSupply),
		Mintable:      true,
		Owner:         simAccount.Address.String(),
	}
}

func randStringBetween(r *rand.Rand, min, max int) string {
	strLen := simtypes.RandIntBetween(r, min, max)
	randStr := simtypes.RandStringOfLength(r, strLen)
	return randStr
}