package keeper_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/farm/keeper"
	"github.com/irisnet/irismod/modules/farm/types"
	"github.com/irisnet/irismod/simapp"
	"github.com/stretchr/testify/suite"
)

var (
	testInitCoinAmt     = sdk.NewInt(100000000_000_000)
	testPoolName        = "USDT-IRIS"
	testPoolDescription = "USDT/IRIS Farm Pool"
	testBeginHeight     = uint64(1)
	testLPTokenDenom    = sdk.DefaultBondDenom
	testRewardPerBlock  = sdk.NewCoins(
		sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1_000_000)),
	)
	testTotalReward = sdk.NewCoins(
		sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000)),
	)
	testDestructible = true

	testCreator sdk.AccAddress
	testFarmer1 sdk.AccAddress
	testFarmer2 sdk.AccAddress
	testFarmer3 sdk.AccAddress
	testFarmer4 sdk.AccAddress
	testFarmer5 sdk.AccAddress

	isCheckTx = false
)

type KeeperTestSuite struct {
	suite.Suite

	cdc    codec.Marshaler
	ctx    sdk.Context
	keeper *keeper.Keeper
	app    *simapp.SimApp
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(isCheckTx)
	suite.cdc = codec.NewAminoCodec(app.LegacyAmino())
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
	suite.keeper = &app.Farmkeeper
	suite.keeper.SetParams(suite.ctx, types.DefaultParams())
	suite.setTestAddrs()
}

func (suite *KeeperTestSuite) setTestAddrs() {
	testAddrs := simapp.AddTestAddrs(suite.app, suite.ctx, 6, testInitCoinAmt)

	testCreator = testAddrs[0]
	testFarmer1 = testAddrs[1]
	testFarmer2 = testAddrs[2]
	testFarmer3 = testAddrs[3]
	testFarmer4 = testAddrs[4]
	testFarmer5 = testAddrs[5]
}

func (suite *KeeperTestSuite) TestCreatePool() {
	ctx := suite.ctx
	err := suite.keeper.CreatePool(ctx,
		testPoolName,
		testPoolDescription,
		testLPTokenDenom,
		testBeginHeight,
		testRewardPerBlock,
		testTotalReward,
		testDestructible,
		testCreator,
	)
	suite.NoError(err)

	//check farm pool
	pool, exist := suite.keeper.GetPool(ctx, testPoolName)
	suite.True(exist)

	suite.Equal(testPoolName, pool.Name)
	suite.Equal(testPoolDescription, pool.Description)
	suite.Equal(testLPTokenDenom, pool.TotalLpTokenLocked.Denom)
	suite.Equal(testBeginHeight, pool.BeginHeight)
	suite.Equal(testDestructible, pool.Destructible)
	suite.Equal(testCreator.String(), pool.Creator)

	//check reward rules
	rules := suite.keeper.GetRewardRules(ctx, testPoolName)
	suite.Len(rules, len(testRewardPerBlock))

	for _, r := range rules {
		suite.Equal(testTotalReward.AmountOf(r.Reward), r.RemainingReward)
		suite.Equal(testTotalReward.AmountOf(r.Reward), r.TotalReward)
		suite.Equal(testRewardPerBlock.AmountOf(r.Reward), r.RewardPerBlock)
		suite.Equal(sdk.ZeroDec(), r.RewardPerShare)
	}

	pool.Rules = rules
	suite.Equal(pool.ExpiredHeight(), pool.EndHeight)

	//check queue
	suite.keeper.IteratorExpiredPool(ctx, pool.EndHeight, func(pool *types.FarmPool) {
		suite.Equal(testPoolName, pool.Name)
	})

	//check balance
	expectedBal := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, testInitCoinAmt)).
		Sub(sdk.NewCoins(suite.keeper.CreatePoolFee(ctx))).
		Sub(testTotalReward)
	actualBal := suite.app.BankKeeper.GetAllBalances(ctx, testCreator)
	suite.Equal(expectedBal, actualBal)
}

func (suite *KeeperTestSuite) TestDestroyPool() {
	ctx := suite.ctx
	err := suite.keeper.CreatePool(ctx,
		testPoolName,
		testPoolDescription,
		testLPTokenDenom,
		testBeginHeight,
		testRewardPerBlock,
		testTotalReward,
		testDestructible,
		testCreator,
	)
	suite.NoError(err)

	newCtx := suite.app.BaseApp.NewContext(isCheckTx, tmproto.Header{
		Height: 10,
	})
	err = suite.keeper.DestroyPool(newCtx, testPoolName, testCreator)
	suite.NoError(err)

	//check farm pool
	pool, exist := suite.keeper.GetPool(newCtx, testPoolName)
	suite.True(exist)

	suite.EqualValues(newCtx.BlockHeight(), pool.LastHeightDistrRewards)
	suite.EqualValues(newCtx.BlockHeight(), pool.EndHeight)

	//check balance
	expectedBal := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, testInitCoinAmt)).
		Sub(sdk.NewCoins(suite.keeper.CreatePoolFee(ctx)))
	actualBal := suite.app.BankKeeper.GetAllBalances(ctx, testCreator)
	suite.Equal(expectedBal, actualBal)
}

func (suite *KeeperTestSuite) TestAppendReward() {
	ctx := suite.ctx
	err := suite.keeper.CreatePool(ctx,
		testPoolName,
		testPoolDescription,
		testLPTokenDenom,
		testBeginHeight,
		testRewardPerBlock,
		testTotalReward,
		testDestructible,
		testCreator,
	)
	suite.NoError(err)

	//check farm pool
	pool, exist := suite.keeper.GetPool(ctx, testPoolName)
	suite.True(exist)

	rewardAdded := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10_000_000))
	remaining, err := suite.keeper.AppendReward(ctx,
		testPoolName,
		sdk.NewCoins(rewardAdded),
		testCreator,
	)
	suite.NoError(err)
	suite.Equal(testTotalReward.Add(rewardAdded), remaining)

	//check farm pool
	pool2, exist := suite.keeper.GetPool(ctx, testPoolName)
	suite.True(exist)
	suite.EqualValues(pool.EndHeight+10, pool2.EndHeight)

	//check reward rules
	rules := suite.keeper.GetRewardRules(ctx, testPoolName)
	suite.Len(rules, len(testRewardPerBlock))

	for _, r := range rules {
		suite.Equal(
			testTotalReward.AmountOf(r.Reward).Add(rewardAdded.Amount),
			r.RemainingReward,
		)
		suite.Equal(
			testTotalReward.AmountOf(r.Reward).Add(rewardAdded.Amount),
			r.TotalReward,
		)
	}
}

func (suite *KeeperTestSuite) TestUpdatePool() {

}

func (suite *KeeperTestSuite) TestRefund() {

}

func (suite *KeeperTestSuite) TestStake() {

}

func (suite *KeeperTestSuite) TestUnstake() {

}

func (suite *KeeperTestSuite) TestHarvest() {

}
