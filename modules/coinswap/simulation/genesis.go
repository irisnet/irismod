package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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
			fee = sdk.NewDecWithPrec(int64(simtypes.RandIntBetween(r, 1, 3)), 3)

			for i := 0; i < 5; i++ {
				tokens = append(tokens, randToken(r, simState.Accounts))
			}
			tokens = append(tokens, tokentypes.GetNativeToken())
			tokenGenesis := tokentypes.NewGenesisState(
				tokentypes.DefaultParams(),
				tokens,
			)
			simState.GenState[tokentypes.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenGenesis)

		},
	)
	genesis := types.NewGenesisState(types.NewParams(fee), tokentypes.GetNativeToken().MinUnit)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(genesis)
	RandomizedGenBankState(simState)
}

// RandomGenesisBalances returns a slice of account balances. Each account has
// a balance of simState.InitialStake for sdk.DefaultBondDenom.
func RandomGenesisBalances(simState *module.SimulationState) ([]banktypes.Balance,sdk.Coins) {
	genesisBalances := []banktypes.Balance{}
	supply := sdk.NewCoins()
	for _, acc := range simState.Accounts {
		coins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(simState.InitialStake)))
		for _,token := range tokens {
			//amt := sdk.NewIntWithDecimal(int64(token.InitialSupply),int(token.Scale))
			coins = coins.Add(sdk.NewCoin(token.MinUnit,sdk.NewInt(1000)))
		}
		genesisBalances = append(genesisBalances, banktypes.Balance{
			Address: acc.Address.String(),
			Coins:   coins,
		})
		supply = supply.Add(coins...)
	}

	supply = supply.Add(sdk.NewCoin(sdk.DefaultBondDenom,sdk.NewInt(simState.InitialStake * simState.NumBonded)))
	return genesisBalances,supply
}

// RandomizedGenBankState generates a random GenesisState for bank
func RandomizedGenBankState(simState *module.SimulationState) {
	balance,supply := RandomGenesisBalances(simState)
	bankGenesis := banktypes.GenesisState{
		Params: banktypes.Params{
			DefaultSendEnabled: true,
		},
		Balances: balance,
		Supply:   supply,
	}

	simState.GenState[banktypes.ModuleName] = simState.Cdc.MustMarshalJSON(&bankGenesis)
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

// RandomDec randomized sdk.RandomDec
func RandomDec(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(r.Int63n(3), 3)
}
