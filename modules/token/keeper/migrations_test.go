package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/ethereum/go-ethereum/common"

	"mods.irisnet.org/modules/token/keeper"
	v3 "mods.irisnet.org/modules/token/migrations/v3"
	"mods.irisnet.org/modules/token/types"
	v1 "mods.irisnet.org/modules/token/types/v1"
)

type alternateBankKeeper struct{ types.BankKeeper }

// Native token configuration is process-global, so these tests must remain serial.
func (suite *KeeperTestSuite) setMigrationNativeToken(token v1.Token) {
	previous := v1.GetNativeToken()
	decodeOwner := func(value string) sdk.AccAddress {
		if value == "" {
			return nil
		}
		address, err := sdk.AccAddressFromBech32(value)
		suite.Require().NoError(err)
		return address
	}
	previousOwner := decodeOwner(previous.Owner)
	tokenOwner := decodeOwner(token.Owner)
	suite.T().Cleanup(func() {
		v1.SetNativeToken(previous.Symbol, previous.Name, previous.MinUnit, previous.Scale,
			previous.InitialSupply, previous.MaxSupply, previous.Mintable, previousOwner)
	})
	v1.SetNativeToken(token.Symbol, token.Name, token.MinUnit, token.Scale,
		token.InitialSupply, token.MaxSupply, token.Mintable, tokenOwner)
}

func (suite *KeeperTestSuite) tokenStoreSnapshot() map[string]string {
	store := suite.ctx.KVStore(suite.app.GetKey(types.StoreKey))
	it := store.Iterator(nil, nil)
	defer it.Close()
	snapshot := make(map[string]string)
	for ; it.Valid(); it.Next() {
		snapshot[string(it.Key())] = string(it.Value())
	}
	return snapshot
}

func (suite *KeeperTestSuite) TestMigrate2to3PreservesOtherNativeNamespaces() {
	for _, tc := range []struct{ symbol, minUnit string }{
		{"uiris", "xuiris"}, {"uiris", "other"}, {"iris", "other"}, {"alpha", "uiris"},
	} {
		suite.Run(tc.symbol+"/"+tc.minUnit, func() {
			suite.SetupTest()
			suite.setMigrationNativeToken(v1.NewToken("stake", "Stake", "stake", 0, 0, 100, true, owner))
			token := v1.NewToken(tc.symbol, "Other Token", tc.minUnit, 6, 0, 100, true, owner)
			token.Contract = common.BytesToAddress(owner).Hex()
			suite.Require().NoError(suite.keeper.AddToken(suite.ctx, token, true))
			suite.keeper.AddBurnCoin(suite.ctx, sdk.NewInt64Coin(tc.minUnit, 7))
			before := suite.tokenStoreSnapshot()
			metadata := suite.bk.GetAllDenomMetaData(suite.ctx)
			suite.NoError(keeper.NewMigrator(suite.keeper, nil).Migrate2to3(suite.ctx))
			suite.Equal(before, suite.tokenStoreSnapshot())
			suite.Equal(metadata, suite.bk.GetAllDenomMetaData(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestMigrate2to3UsesConfiguredNativeToken() {
	suite.legacyMigrationState()
	suite.setMigrationNativeToken(v1.NewToken("stake", "Stake", "stake", 0, 0, 100, true, owner))
	before := suite.tokenStoreSnapshot()
	metadata := suite.bk.GetAllDenomMetaData(suite.ctx)
	suite.NoError(keeper.NewMigrator(suite.keeper, nil).Migrate2to3(suite.ctx))
	suite.Equal(before, suite.tokenStoreSnapshot())
	suite.Equal(metadata, suite.bk.GetAllDenomMetaData(suite.ctx))
}

func (suite *KeeperTestSuite) TestMigrate2to3RejectsInconsistentNativeNamespace() {
	for _, tc := range []string{"missing pair", "missing record", "missing index", "record symbol", "record min-unit", "index target", "malformed record", "malformed index"} {
		suite.Run(tc, func() {
			suite.SetupTest()
			suite.legacyMigrationState()
			store := suite.ctx.KVStore(suite.app.GetKey(types.StoreKey))
			cdc := suite.keeper.Codec()
			var native v1.Token
			suite.Require().NoError(cdc.Unmarshal(store.Get(types.KeySymbol("iris")), &native))
			switch tc {
			case "missing pair":
				store.Delete(types.KeySymbol("iris"))
				store.Delete(types.KeyMinUint("uiris"))
			case "missing record":
				store.Delete(types.KeySymbol("iris"))
			case "missing index":
				store.Delete(types.KeyMinUint("uiris"))
			case "record symbol":
				native.Symbol = "other"
				store.Set(types.KeySymbol("iris"), cdc.MustMarshal(&native))
			case "record min-unit":
				native.MinUnit = "other"
				store.Set(types.KeySymbol("iris"), cdc.MustMarshal(&native))
			case "index target":
				store.Set(types.KeyMinUint("uiris"), cdc.MustMarshal(&gogotypes.StringValue{Value: "other"}))
			case "malformed record":
				store.Set(types.KeySymbol("iris"), []byte{0xff})
			case "malformed index":
				store.Set(types.KeyMinUint("uiris"), []byte{0xff})
			}
			before := suite.tokenStoreSnapshot()
			metadata := suite.bk.GetAllDenomMetaData(suite.ctx)
			suite.Error(keeper.NewMigrator(suite.keeper, nil).Migrate2to3(suite.ctx))
			suite.Equal(before, suite.tokenStoreSnapshot())
			suite.Equal(metadata, suite.bk.GetAllDenomMetaData(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestMigrate2to3BankKeeperCompatibility() {
	legacy, keys := suite.legacyMigrationState()
	key := suite.app.GetKey(types.StoreKey)
	store := suite.ctx.KVStore(key)
	before := store.Get(types.KeySymbol(legacy.Symbol))
	suite.ErrorContains(v3.Migrate(suite.ctx, key, suite.keeper.Codec(), alternateBankKeeper{suite.bk}), "unsupported bank keeper")
	suite.Equal(before, store.Get(types.KeySymbol(legacy.Symbol)))
	for _, key := range keys {
		suite.True(store.Has(key))
	}
	_, found := suite.bk.GetDenomMetaData(suite.ctx, legacy.MinUnit)
	suite.True(found)

	bk, ok := suite.bk.(bankkeeper.BaseKeeper)
	suite.Require().True(ok)
	suite.Require().NoError(v3.Migrate(suite.ctx, key, suite.keeper.Codec(), &bk))
	suite.False(store.Has(types.KeySymbol(legacy.Symbol)))
	// An absent registration requires no bank metadata access.
	suite.NoError(v3.Migrate(suite.ctx, key, suite.keeper.Codec(), nil))
}

func (suite *KeeperTestSuite) TestMigrate2to3AllowsAbsentOptionalState() {
	legacy, keys := suite.legacyMigrationState()
	store := suite.ctx.KVStore(suite.app.GetKey(types.StoreKey))
	for _, key := range keys[2:] {
		store.Delete(key)
	}
	bk, ok := suite.bk.(bankkeeper.BaseKeeper)
	suite.Require().True(ok)
	suite.Require().NoError(bk.BaseViewKeeper.DenomMetadata.Remove(suite.ctx, legacy.MinUnit))
	suite.Require().NoError(keeper.NewMigrator(suite.keeper, nil).Migrate2to3(suite.ctx))
	for _, key := range keys {
		suite.False(store.Has(key))
	}
	token, err := suite.keeper.GetToken(suite.ctx, "uiris")
	suite.Require().NoError(err)
	suite.Equal("iris", token.GetSymbol())
}

func (suite *KeeperTestSuite) legacyMigrationState() (v1.Token, [][]byte) {
	native := v1.NewToken("iris", "IRIS", "uiris", 6, 0, 100, true, add2)
	suite.setMigrationNativeToken(native)
	suite.setToken(native)
	legacy := v1.NewToken("uiris", "Legacy", "xuiris", 18, 0, 100, true, owner)
	legacy.Contract = common.BytesToAddress(owner).Hex()
	suite.seedLegacyToken(legacy)
	suite.bk.SetDenomMetaData(suite.ctx, banktypes.Metadata{Base: "xuiris", Display: "uiris"})
	suite.keeper.AddBurnCoin(suite.ctx, sdk.NewInt64Coin("xuiris", 7))
	return legacy, [][]byte{
		types.KeySymbol("uiris"), types.KeyMinUint("xuiris"),
		types.KeyTokens(owner, "uiris"), types.KeyContract(legacy.Contract), types.KeyBurnTokenAmt("xuiris"),
	}
}

func (suite *KeeperTestSuite) TestMigrate2to3() {
	_, removedKeys := suite.legacyMigrationState()
	coins := sdk.NewCoins(sdk.NewInt64Coin("xuiris", 25))
	suite.Require().NoError(suite.bk.MintCoins(suite.ctx, types.ModuleName, coins))
	suite.Require().NoError(suite.bk.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, owner, coins))
	store := suite.ctx.KVStore(suite.app.GetKey(types.StoreKey))
	nativeRecord := store.Get(types.KeySymbol("iris"))
	nativeIndex := store.Get(types.KeyMinUint("uiris"))
	nativeMetadata, found := suite.bk.GetDenomMetaData(suite.ctx, "uiris")
	suite.Require().True(found)
	migrate := keeper.NewMigrator(suite.keeper, nil)
	versions := suite.app.ModuleManager.GetVersionMap()
	versions[types.ModuleName] = 2
	versions, err := suite.app.ModuleManager.RunMigrations(suite.ctx, suite.app.Configurator(), versions)
	suite.Require().NoError(err)
	suite.Equal(uint64(3), versions[types.ModuleName])
	for i := 0; i < 2; i++ {
		suite.Require().NoError(migrate.Migrate2to3(suite.ctx))
		for _, key := range removedKeys {
			suite.False(store.Has(key), "key %x remains", key)
		}
		suite.Equal(nativeRecord, store.Get(types.KeySymbol("iris")))
		suite.Equal(nativeIndex, store.Get(types.KeyMinUint("uiris")))
		token, err := suite.keeper.GetToken(suite.ctx, "uiris")
		suite.Require().NoError(err)
		suite.Equal("iris", token.GetSymbol())
		_, err = suite.keeper.GetToken(suite.ctx, "xuiris")
		suite.ErrorIs(err, types.ErrTokenNotExists)
		suite.Empty(suite.keeper.GetTokens(suite.ctx, owner))
		_, found := suite.bk.GetDenomMetaData(suite.ctx, "xuiris")
		suite.False(found)
		metadata, found := suite.bk.GetDenomMetaData(suite.ctx, "uiris")
		suite.True(found)
		suite.Equal(nativeMetadata, metadata)
		suite.Equal(coins[0], suite.bk.GetBalance(suite.ctx, owner, "xuiris"))
		suite.Equal(coins[0], suite.bk.GetSupply(suite.ctx, "xuiris"))
	}
}

func (suite *KeeperTestSuite) TestMigrate2to3RejectsInconsistentState() {
	for _, tc := range []string{"identity", "missing min-unit index", "min-unit index", "owner index", "contract index", "invalid owner", "invalid contract", "metadata"} {
		suite.Run(tc, func() {
			suite.SetupTest()
			legacy, keys := suite.legacyMigrationState()
			store := suite.ctx.KVStore(suite.app.GetKey(types.StoreKey))
			cdc := suite.keeper.Codec()
			switch tc {
			case "identity":
				legacy.MinUnit = "other"
				store.Set(types.KeySymbol("uiris"), cdc.MustMarshal(&legacy))
			case "missing min-unit index":
				store.Delete(types.KeyMinUint("xuiris"))
			case "min-unit index":
				store.Set(types.KeyMinUint("xuiris"), cdc.MustMarshal(&gogotypes.StringValue{Value: "iris"}))
			case "owner index":
				store.Set(types.KeyTokens(owner, "uiris"), cdc.MustMarshal(&gogotypes.StringValue{Value: "iris"}))
			case "contract index":
				store.Set(types.KeyContract(legacy.Contract), cdc.MustMarshal(&gogotypes.StringValue{Value: "iris"}))
			case "invalid owner":
				legacy.Owner = "invalid"
				store.Set(types.KeySymbol("uiris"), cdc.MustMarshal(&legacy))
			case "invalid contract":
				legacy.Contract = "invalid"
				store.Set(types.KeySymbol("uiris"), cdc.MustMarshal(&legacy))
			case "metadata":
				bankStore := suite.ctx.KVStore(suite.app.GetKey(banktypes.StoreKey))
				bankStore.Set(append(banktypes.DenomMetadataPrefix, []byte("xuiris")...), cdc.MustMarshal(&banktypes.Metadata{Base: "other"}))
			}
			before := make([][]byte, len(keys))
			for i, key := range keys {
				before[i] = store.Get(key)
			}
			metadata, found := suite.bk.GetDenomMetaData(suite.ctx, "xuiris")
			suite.Require().Error(keeper.NewMigrator(suite.keeper, nil).Migrate2to3(suite.ctx))
			for i, key := range keys {
				suite.Equal(before[i], store.Get(key), "key %x changed", key)
			}
			afterMetadata, afterFound := suite.bk.GetDenomMetaData(suite.ctx, "xuiris")
			suite.Equal(found, afterFound)
			suite.Equal(metadata, afterMetadata)
		})
	}
}
