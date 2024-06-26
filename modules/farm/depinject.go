package farm

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	modulev1 "mods.irisnet.org/api/irismod/farm/module/v1"
	"mods.irisnet.org/modules/farm/keeper"
	"mods.irisnet.org/modules/farm/types"
)

// App Wiring Setup

func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule, ProvideKeyTable),
	)
}

func ProvideKeyTable() types.KeyTable {
	return types.ParamKeyTable()
}

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

type FarmInputs struct {
	depinject.In

	Config *modulev1.Module
	Cdc    codec.Codec
	Key    *store.KVStoreKey

	AccountKeeper  types.AccountKeeper
	BankKeeper     types.BankKeeper
	DistrKeeper    types.DistrKeeper
	GovKeeper      types.GovKeeper
	CoinswapKeeper types.CoinswapKeeper

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace types.Subspace `optional:"true"`
}

type FarmOutputs struct {
	depinject.Out

	FarmKeeper keeper.Keeper
	Module     appmodule.AppModule
}

func ProvideModule(in FarmInputs) FarmOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	keeper := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.BankKeeper,
		in.AccountKeeper,
		in.DistrKeeper,
		in.GovKeeper,
		in.CoinswapKeeper,
		in.Config.FeeCollectorName,
		in.Config.CommunityPoolName,
		authority.String(),
	)
	m := NewAppModule(in.Cdc, keeper, in.AccountKeeper, in.BankKeeper, in.LegacySubspace)

	return FarmOutputs{FarmKeeper: keeper, Module: m}
}
