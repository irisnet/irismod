package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	"mods.irisnet.org/modules/token/keeper"
	tokentypes "mods.irisnet.org/modules/token/types"
	v1 "mods.irisnet.org/modules/token/types/v1"
	"mods.irisnet.org/modules/token/types/v1beta1"
	"mods.irisnet.org/simapp"
)

const (
	isCheckTx = false
)

var (
	denom    = v1.GetNativeToken().Symbol
	owner    = sdk.AccAddress(tmhash.SumTruncated([]byte("tokenTest")))
	add2     = sdk.AccAddress(tmhash.SumTruncated([]byte("tokenTest1")))
	initAmt  = sdkmath.NewIntWithDecimal(100000000, int(6))
	initCoin = sdk.Coins{sdk.NewCoin(denom, initAmt)}
	beacon   = common.BytesToAddress(owner.Bytes())
)

type KeeperTestSuite struct {
	suite.Suite

	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	keeper      keeper.Keeper
	bk          bankkeeper.Keeper
	app         *simapp.SimApp
}

func (suite *KeeperTestSuite) SetupTest() {
	depInjectOptions := simapp.DepinjectOptions{
		Config: AppConfig,
		Providers: []interface{}{
			keeper.ProvideMockEVM(),
			keeper.ProvideMockICS20(),
		},
		Consumers: []interface{}{&suite.keeper},
	}

	app := simapp.Setup(suite.T(), isCheckTx, depInjectOptions)

	suite.legacyAmino = app.LegacyAmino()
	suite.ctx = app.BaseApp.NewContext(isCheckTx)
	suite.bk = app.BankKeeper
	suite.app = app

	// set params
	params := v1.DefaultParams()
	params.Beacon = beacon.String()
	suite.NoError(suite.keeper.SetParams(suite.ctx, params), "SetParams failed")

	// init tokens to addr
	err := suite.bk.MintCoins(suite.ctx, tokentypes.ModuleName, initCoin)
	suite.NoError(err)
	err = suite.bk.SendCoinsFromModuleToAccount(suite.ctx, tokentypes.ModuleName, owner, initCoin)
	suite.NoError(err)
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) setToken(token v1.Token) {
	err := suite.keeper.AddToken(suite.ctx, token, true)
	suite.NoError(err)
}

func (suite *KeeperTestSuite) issueToken(token v1.Token) {
	suite.setToken(token)

	mintCoins := sdk.NewCoins(
		sdk.NewCoin(
			token.MinUnit,
			sdkmath.NewIntWithDecimal(int64(token.InitialSupply), int(token.Scale)),
		),
	)

	err := suite.bk.MintCoins(suite.ctx, tokentypes.ModuleName, mintCoins)
	suite.NoError(err)

	err = suite.bk.SendCoinsFromModuleToAccount(suite.ctx, tokentypes.ModuleName, owner, mintCoins)
	suite.NoError(err)
}

func (suite *KeeperTestSuite) TestIssueToken() {
	token := v1.NewToken("btc", "Bitcoin Network", "satoshi", 18, 21000000, 21000000, false, owner)
	ownerBalance := suite.bk.GetBalance(suite.ctx, token.GetOwner(), token.MinUnit)
	totalSupply := suite.bk.GetSupply(suite.ctx, token.MinUnit)

	err := suite.keeper.IssueToken(
		suite.ctx, token.Symbol, token.Name,
		token.MinUnit, token.Scale, token.InitialSupply,
		token.MaxSupply, token.Mintable, token.GetOwner(),
	)
	suite.ErrorIs(err, tokentypes.ErrIssueTokenDisabled)
	suite.False(suite.keeper.HasToken(suite.ctx, token.Symbol))
	suite.Equal(ownerBalance, suite.bk.GetBalance(suite.ctx, token.GetOwner(), token.MinUnit))
	suite.Equal(totalSupply, suite.bk.GetSupply(suite.ctx, token.MinUnit))

	// A zero-value keeper proves the guard returns before touching any dependency.
	err = (keeper.Keeper{}).IssueToken(
		sdk.Context{}, token.Symbol, token.Name,
		token.MinUnit, token.Scale, token.InitialSupply,
		token.MaxSupply, token.Mintable, token.GetOwner(),
	)
	suite.ErrorIs(err, tokentypes.ErrIssueTokenDisabled)
}

func (suite *KeeperTestSuite) TestIssueTokenMsgServers() {
	msgServer := keeper.NewMsgServerImpl(suite.keeper)
	legacyMsgServer := keeper.NewLegacyMsgServerImpl(msgServer, suite.keeper)
	ownerBalance := suite.bk.GetBalance(suite.ctx, owner, denom)

	tests := []struct {
		name string
		call func(context.Context) error
	}{
		{
			name: "v1",
			call: func(ctx context.Context) error {
				_, err := msgServer.IssueToken(ctx, v1.NewMsgIssueToken(
					"btc", "satoshi", "Bitcoin Network", 18, 1, 1, false, owner.String(),
				))
				return err
			},
		},
		{
			name: "v1beta1",
			call: func(ctx context.Context) error {
				_, err := legacyMsgServer.IssueToken(ctx, v1beta1.NewMsgIssueToken(
					"btc", "satoshi", "Bitcoin Network", 18, 1, 1, false, owner.String(),
				))
				return err
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			cacheCtx, _ := suite.ctx.CacheContext()
			err := tc.call(sdk.WrapSDKContext(cacheCtx))
			suite.ErrorIs(err, tokentypes.ErrIssueTokenDisabled)
			suite.False(suite.keeper.HasToken(cacheCtx, "btc"))
			suite.True(suite.bk.GetBalance(cacheCtx, owner, denom).IsLT(ownerBalance))

			// BaseApp discards the message cache on error, rolling back the issuance fee.
			// AnteHandler transaction fees live in an earlier cache and are unaffected.
			suite.Equal(ownerBalance, suite.bk.GetBalance(suite.ctx, owner, denom))
		})
	}
}

func (suite *KeeperTestSuite) TestEditToken() {
	token := v1.NewToken("btc", "Bitcoin Network", "satoshi", 18, 21000000, 21000000, false, owner)
	suite.setToken(token)

	symbol := "btc"
	name := "Bitcoin Token"
	mintable := tokentypes.True
	maxSupply := uint64(22000000)

	err := suite.keeper.EditToken(suite.ctx, symbol, name, maxSupply, mintable, owner)
	suite.NoError(err)

	newToken, err := suite.keeper.GetToken(suite.ctx, symbol)
	suite.NoError(err)

	expToken := v1.NewToken(
		"btc",
		"Bitcoin Token",
		"satoshi",
		18,
		21000000,
		22000000,
		mintable.ToBool(),
		owner,
	)

	suite.EqualValues(newToken.(*v1.Token), &expToken)
}

func (suite *KeeperTestSuite) TestMintToken() {
	token := v1.NewToken("btc", "Bitcoin Network", "satoshi", 18, 1000, 2000, true, owner)
	suite.issueToken(token)

	amt := suite.bk.GetBalance(suite.ctx, token.GetOwner(), token.MinUnit)
	suite.Equal("1000000000000000000000satoshi", amt.String())

	coinMinted := sdk.NewCoin(token.MinUnit, sdkmath.NewIntWithDecimal(1000, int(token.Scale)))
	recipient := sdk.AccAddress{}

	err := suite.keeper.MintToken(suite.ctx, coinMinted, recipient, token.GetOwner())
	suite.NoError(err)

	amt = suite.bk.GetBalance(suite.ctx, token.GetOwner(), token.MinUnit)
	suite.Equal("2000000000000000000000satoshi", amt.String())

	// mint token without owner

	err = suite.keeper.MintToken(suite.ctx, coinMinted, owner, sdk.AccAddress{})
	suite.Error(err, "can not mint token without owner when the owner exists")

	token = v1.NewToken("atom", "Cosmos Hub", "uatom", 6, 1000, 2000, true, sdk.AccAddress{})
	suite.issueToken(token)

	err = suite.keeper.MintToken(
		suite.ctx,
		sdk.NewCoin(token.MinUnit, sdkmath.OneInt()),
		owner,
		sdk.AccAddress{},
	)
	suite.NoError(err)
}

func (suite *KeeperTestSuite) TestBurnToken() {
	token := v1.NewToken("btc", "Bitcoin Network", "satoshi", 18, 1000, 2000, true, owner)
	suite.issueToken(token)

	amt := suite.bk.GetBalance(suite.ctx, token.GetOwner(), token.MinUnit)
	suite.Equal("1000000000000000000000satoshi", amt.String())

	coinBurnt := sdk.NewCoin(token.MinUnit, sdkmath.NewIntWithDecimal(200, int(token.Scale)))

	err := suite.keeper.BurnToken(suite.ctx, coinBurnt, token.GetOwner())
	suite.NoError(err)

	amt = suite.bk.GetBalance(suite.ctx, token.GetOwner(), token.MinUnit)
	suite.Equal("800000000000000000000satoshi", amt.String())
}

func (suite *KeeperTestSuite) TestTransferToken() {
	token := v1.NewToken("btc", "Bitcoin Network", "satoshi", 18, 21000000, 21000000, false, owner)
	suite.setToken(token)

	dstOwner := sdk.AccAddress(tmhash.SumTruncated([]byte("TokenDstOwner")))

	err := suite.keeper.TransferTokenOwner(suite.ctx, token.Symbol, token.GetOwner(), dstOwner)
	suite.NoError(err)

	newToken, err := suite.keeper.GetToken(suite.ctx, token.Symbol)
	suite.NoError(err)

	suite.Equal(dstOwner, newToken.GetOwner())
}

func (suite *KeeperTestSuite) TestSwapFeeToken() {
	token1 := v1.NewToken("token1", "Test Token1", "t1min", 6, 1000, 2000, true, owner)
	suite.issueToken(token1)

	amt1 := suite.bk.GetBalance(suite.ctx, token1.GetOwner(), token1.MinUnit)
	suite.Equal("1000000000t1min", amt1.String())

	token2 := v1.NewToken("token2", "Test Token1", "t2min", 18, 0, 2000, true, add2)
	suite.issueToken(token2)

	suite.keeper = suite.keeper.WithSwapRegistry(v1.SwapRegistry{
		token1.MinUnit: v1.SwapParams{
			MinUnit: token2.MinUnit,
			Ratio:   math.LegacyNewDec(1),
		},
		token2.MinUnit: v1.SwapParams{
			MinUnit: token1.MinUnit,
			Ratio:   math.LegacyNewDec(1),
		},
	})

	amt2 := suite.bk.GetBalance(suite.ctx, add2, token2.MinUnit)
	suite.Equal("0t2min", amt2.String())

	feePaid := sdk.NewCoin(token1.MinUnit, sdkmath.NewIntWithDecimal(100, int(token1.Scale)))

	_, feeGot, err := suite.keeper.SwapFeeToken(
		suite.ctx,
		feePaid,
		token1.GetOwner(),
		token2.GetOwner(),
	)
	suite.NoError(err)
	suite.Equal("100000000000000000000t2min", feeGot.String())

	amt := suite.bk.GetBalance(suite.ctx, token1.GetOwner(), token1.MinUnit)
	suite.Equal("900000000t1min", amt.String())

	amt = suite.bk.GetBalance(suite.ctx, token2.GetOwner(), token2.MinUnit)
	suite.Equal("100000000000000000000t2min", amt.String())

	// reverse test
	_, feeGot, err = suite.keeper.SwapFeeToken(
		suite.ctx,
		feeGot,
		token2.GetOwner(),
		token1.GetOwner(),
	)
	suite.NoError(err)
	suite.Equal("100000000t1min", feeGot.String())

	amt = suite.bk.GetBalance(suite.ctx, token1.GetOwner(), token1.MinUnit)
	suite.Equal("1000000000t1min", amt.String())

	amt = suite.bk.GetBalance(suite.ctx, token2.GetOwner(), token2.MinUnit)
	suite.Equal("0t2min", amt.String())
}
