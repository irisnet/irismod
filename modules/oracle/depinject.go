package oracle

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	modulev1 "github.com/irisnet/irismod/api/irismod/oracle/module/v1"
	"github.com/irisnet/irismod/modules/oracle/keeper"
	"github.com/irisnet/irismod/modules/oracle/types"
)

// App Wiring Setup
func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

type OracleInputs struct {
	depinject.In

	Config *modulev1.Module
	Cdc    codec.Codec
	Key    *store.KVStoreKey

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
	ServiceKeeper types.ServiceKeeper

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace paramtypes.Subspace `optional:"true"`
}

type OracleOutputs struct {
	depinject.Out

	OracleKeeper keeper.Keeper
	Module       appmodule.AppModule
}

func ProvideModule(in OracleInputs) OracleOutputs {
	keeper := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.LegacySubspace,
		in.ServiceKeeper,
	)
	m := NewAppModule(in.Cdc, keeper, in.AccountKeeper, in.BankKeeper)

	return OracleOutputs{OracleKeeper: keeper, Module: m}
}
