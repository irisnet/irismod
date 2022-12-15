package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irisnet/irismod/modules/rental/keeper"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	registry cdctypes.InterfaceRegistry,
	appParams simtypes.AppParams,
	cdc codec.JSONCodec,
	k keeper.Keeper,
) simulation.WeightedOperations {
	panic("Fixme")
}
