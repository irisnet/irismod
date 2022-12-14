package module

import (
	"encoding/json"
	"math/rand"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/irisnet/irismod/modules/rental/keeper"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic defines the basic application module used by the rental module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the Rental module's name.
func (AppModuleBasic) Name() string { panic("Fixme!") }

// RegisterLegacyAminoCodec registers the rental module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// fixme
	panic("Fixme!")
	//types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the rental module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	// fixme
	panic("Fixme!")
	//return cdc.MustMarshalJSON(types.d)
}

// ValidateGenesis performs genesis state validation for the rental module.
func (AppModuleBasic) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	panic("Fixme!")
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the rental module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {
	panic("Fixme!")
}

// GetTxCmd returns the root tx command for the rental module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	panic("Fixme!")
}

// GetQueryCmd returns the root query command for the rental  module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	panic("Fixme!")
}

// RegisterInterfaces registers interfaces and implementations of the rental module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	panic("Fixme!")
}

// AppModule implements an application module for the rental module
type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
	// fixme: add keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	panic("Fixme!")
}

// Name returns the rental module's name.
func (AppModule) Name() string { panic("Fixme!") }

// RegisterInvariants registers the rental module invariants.
func (AppModule) RegisterInvariants(sdk.InvariantRegistry) {
	panic("Fixme!")
}

// Deprecated: Route returns the message routing key for the rental module.
func (AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute returns the rental module's querier route name.
func (AppModule) QuerierRoute() string {
	panic("Fixme!")
}

// LegacyQuerierHandler returns the rental module sdk.Querier.
func (AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	panic("Fixme!")
}

// RegisterServices registers module services.
func (AppModule) RegisterServices(cfg module.Configurator) {
	panic("Fixme!")
}

// InitGenesis performs genesis initialization for the rental module. It returns
// no validator updates.
func (AppModule) InitGenesis(sdk.Context, codec.JSONCodec, json.RawMessage) []abci.ValidatorUpdate {
	panic("Fixme!")
}

// ExportGenesis returns the exported genesis state as raw bytes for the NFT module.
func (AppModule) ExportGenesis(sdk.Context, codec.JSONCodec) json.RawMessage {
	panic("Fixme!")
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// AppModuleSimulation function

// GenerateGenesisState creates a randomized GenState of the rental module.
func (AppModule) GenerateGenesisState(input *module.SimulationState) {
	panic("Fixme!")
}

// content functions used to simulate governance proposals
func (AppModule) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	panic("Fixme!")
}

// randomized module parameters for param change proposals
func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	panic("Fixme!")
}

// register a func to decode the each module's defined types from their corresponding store key
func (AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// simulation operations (i.e msgs) with their respective weight
func (AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	panic("Fixme!")
}
