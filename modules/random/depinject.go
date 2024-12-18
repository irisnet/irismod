package random

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	store "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"

	modulev1 "mods.irisnet.org/api/irismod/random/module/v1"
	"mods.irisnet.org/modules/random/keeper"
	"mods.irisnet.org/modules/random/types"
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

// Inputs define the module inputs for the depinject.
type Inputs struct {
	depinject.In

	Config *modulev1.Module
	Cdc    codec.Codec
	Key    *store.KVStoreKey

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
	ServiceKeeper types.ServiceKeeper
}

// Outputs define the module outputs for the depinject.
type Outputs struct {
	depinject.Out

	RandomKeeper keeper.Keeper
	Module       appmodule.AppModule
}

// ProvideModule creates a new AppModule and returns the Outputs.
//
// - in: the Inputs struct containing the necessary dependencies for creating the AppModule.
// Returns:
// - Outputs: the struct containing the RandomKeeper and AppModule.
func ProvideModule(in Inputs) Outputs {
	keeper := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.BankKeeper,
		in.ServiceKeeper,
	)
	m := NewAppModule(in.Cdc, keeper, in.AccountKeeper, in.BankKeeper)

	return Outputs{RandomKeeper: keeper, Module: m}
}
