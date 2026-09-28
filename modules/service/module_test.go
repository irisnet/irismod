package service

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/stretchr/testify/require"

	"mods.irisnet.org/modules/service/types"
)

func TestGenerateGenesisStateUsesValidDefaults(t *testing.T) {
	c := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	simState := module.SimulationState{
		Cdc:      c,
		GenState: make(map[string]json.RawMessage),
	}

	AppModule{}.GenerateGenesisState(&simState)

	var genesis types.GenesisState
	c.MustUnmarshalJSON(simState.GenState[types.ModuleName], &genesis)
	require.NoError(t, types.ValidateGenesis(genesis))
	require.Equal(t, types.DefaultParams(), genesis.Params)
}
