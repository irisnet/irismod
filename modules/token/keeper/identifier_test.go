package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/cosmos/gogoproto/types"

	"mods.irisnet.org/modules/token/types"
	v1 "mods.irisnet.org/modules/token/types/v1"
)

func (suite *KeeperTestSuite) TestAddTokenIdentifierUniqueness() {
	existing := v1.NewToken("alpha", "Alpha", "ualpha", 6, 0, 100, true, owner)
	suite.Require().NoError(suite.keeper.AddToken(suite.ctx, existing, true))
	for _, tc := range []struct {
		symbol, minUnit string
		want            error
	}{
		{"ualpha", "ubeta", types.ErrSymbolAlreadyExists},
		{"beta", "alpha", types.ErrMinUnitAlreadyExists},
		{"alpha", "ubeta", types.ErrSymbolAlreadyExists},
		{"beta", "ualpha", types.ErrMinUnitAlreadyExists},
	} {
		suite.Run(tc.symbol+"/"+tc.minUnit, func() {
			token := v1.NewToken(tc.symbol, "Beta", tc.minUnit, 6, 0, 100, true, owner)
			suite.ErrorIs(suite.keeper.AddToken(suite.ctx, token, true), tc.want)
		})
	}
	tokens := suite.keeper.GetTokens(suite.ctx, owner)
	suite.Len(tokens, 1)
	actual, err := suite.keeper.GetToken(suite.ctx, existing.MinUnit)
	suite.Require().NoError(err)
	suite.Equal(existing.Symbol, actual.GetSymbol())
}

func (suite *KeeperTestSuite) seedLegacyToken(token v1.Token) {
	store := suite.ctx.KVStore(suite.app.GetKey(types.StoreKey))
	cdc := suite.keeper.Codec()
	store.Set(types.KeySymbol(token.Symbol), cdc.MustMarshal(&token))
	index := cdc.MustMarshal(&gogotypes.StringValue{Value: token.Symbol})
	store.Set(types.KeyMinUint(token.MinUnit), index)
	if token.Owner != "" {
		store.Set(types.KeyTokens(token.GetOwner(), token.Symbol), index)
	}
	if token.Contract != "" {
		store.Set(types.KeyContract(token.Contract), index)
	}
}

func (suite *KeeperTestSuite) TestSwapFeeTokenUsesMinUnit() {
	input := v1.NewToken("input", "Input", "uinput", 6, 10, 100, true, owner)
	output := v1.NewToken("output", "Output", "uoutput", 6, 0, 100, true, add2)
	suite.issueToken(input)
	suite.setToken(output)
	suite.seedLegacyToken(v1.NewToken("uoutput", "Other", "uother", 18, 0, 100, true, owner))
	suite.keeper = suite.keeper.WithSwapRegistry(v1.SwapRegistry{
		input.MinUnit: {MinUnit: output.MinUnit, Ratio: math.LegacyOneDec()},
	})
	paid := sdk.NewInt64Coin(input.MinUnit, 1_000_000)
	burnt, minted, err := suite.keeper.SwapFeeToken(suite.ctx, paid, owner, add2)
	suite.Require().NoError(err)
	suite.Equal(paid, burnt)
	suite.Equal(sdk.NewInt64Coin(output.MinUnit, 1_000_000), minted)
	suite.Equal(sdk.NewInt64Coin(output.MinUnit, 1_000_000), suite.bk.GetBalance(suite.ctx, add2, output.MinUnit))
}

func (suite *KeeperTestSuite) TestDeployERC20IdentifierUniqueness() {
	existing := v1.NewToken("alpha", "Alpha", "ualpha", 6, 0, 100, true, owner)
	suite.setToken(existing)
	for _, tc := range []struct {
		symbol, minUnit string
		want            error
	}{
		{"ualpha", "ubeta", types.ErrSymbolAlreadyExists},
		{"beta", "alpha", types.ErrMinUnitAlreadyExists},
	} {
		suite.Run(tc.symbol+"/"+tc.minUnit, func() {
			_, err := suite.keeper.DeployERC20(suite.ctx, "Beta", tc.symbol, tc.minUnit, 6)
			suite.ErrorIs(err, tc.want)
		})
	}
	suite.Len(suite.keeper.GetTokens(suite.ctx, nil), 2)
}
