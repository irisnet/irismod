package token

import (
	"context"
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"mods.irisnet.org/modules/token/client/cli"
	"mods.irisnet.org/modules/token/keeper"
	"mods.irisnet.org/modules/token/simulation"
	"mods.irisnet.org/modules/token/types"
	v1 "mods.irisnet.org/modules/token/types/v1"
	"mods.irisnet.org/modules/token/types/v1beta1"
)

// ConsensusVersion defines the current token module consensus version.
const ConsensusVersion = 2

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic defines the basic application module used by the token module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the token module's name.
func (AppModuleBasic) Name() string { return types.ModuleName }

// RegisterLegacyAminoCodec registers the token module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	v1.RegisterLegacyAminoCodec(cdc)
	v1beta1.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the token module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the token module.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	config client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	var data v1.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return v1.ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the token module.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the token module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = v1.RegisterQueryHandlerClient(context.Background(), mux, v1.NewQueryClient(clientCtx))
	_ = v1beta1.RegisterQueryHandlerClient(
		context.Background(),
		mux,
		v1beta1.NewQueryClient(clientCtx),
	)
}

// GetTxCmd returns the root tx command for the token module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns no root query command for the token module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers interfaces and implementations of the token module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1.RegisterInterfaces(registry)
	v1beta1.RegisterInterfaces(registry)
}

// ____________________________________________________________________________

// AppModule implements an application module for the token module.
type AppModule struct {
	AppModuleBasic

	keeper         keeper.Keeper
	accountKeeper  types.AccountKeeper
	bankKeeper     types.BankKeeper
	legacySubspace types.Subspace
}

// NewAppModule creates a new AppModule object
func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	legacySubspace types.Subspace,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
		accountKeeper:  ak,
		bankKeeper:     bk,
		legacySubspace: legacySubspace,
	}
}

// Name returns the token module's name.
func (AppModule) Name() string { return types.ModuleName }

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	v1MsgServer := keeper.NewMsgServerImpl(am.keeper)
	v1.RegisterMsgServer(cfg.MsgServer(), v1MsgServer)
	v1.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	v1beta1.RegisterMsgServer(
		cfg.MsgServer(),
		keeper.NewLegacyMsgServerImpl(v1MsgServer, am.keeper),
	)
	v1beta1.RegisterQueryServer(
		cfg.QueryServer(),
		keeper.NewLegacyQueryServer(am.keeper, am.keeper.Codec()),
	)

	m := keeper.NewMigrator(am.keeper, am.legacySubspace)
	if err := cfg.RegisterMigration(types.ModuleName, 1, m.Migrate1to2); err != nil {
		panic(err)
	}
}

// RegisterInvariants registers the token module invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// InitGenesis performs genesis initialization for the token module. It returns
// no validator updates.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	cdc codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	var genesisState v1.GenesisState

	cdc.MustUnmarshalJSON(data, &genesisState)

	InitGenesis(ctx, am.keeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the token module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// BeginBlock performs a no-op.
func (AppModule) BeginBlock(_ context.Context) error { return nil }

// EndBlock returns the end blocker for the token module. It returns no validator updates.
func (AppModule) EndBlock(_ context.Context) error { return nil }

// ____________________________________________________________________________

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the token module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// RegisterStoreDecoder registers a decoder for token module's types
func (am AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	sdr[types.StoreKey] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the token module operations with their respective weights.
func (am AppModule) WeightedOperations(
	simState module.SimulationState,
) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(
		simState.AppParams,
		simState.Cdc,
		am.keeper,
		am.accountKeeper,
		am.bankKeeper,
	)
}
