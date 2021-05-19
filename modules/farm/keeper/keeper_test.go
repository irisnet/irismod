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
		sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000_000_000)),
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
	suite.Require().NoError(err)

	//check farm pool
	pool, exist := suite.keeper.GetPool(ctx, testPoolName)
	suite.Require().True(exist)

	suite.Require().Equal(testPoolName, pool.Name)
	suite.Require().Equal(testPoolDescription, pool.Description)
	suite.Require().Equal(testLPTokenDenom, pool.TotalLpTokenLocked.Denom)
	suite.Require().Equal(testBeginHeight, pool.BeginHeight)
	suite.Require().Equal(testDestructible, pool.Destructible)
	suite.Require().Equal(testCreator.String(), pool.Creator)

	//check reward rules
	rules := suite.keeper.GetRewardRules(ctx, testPoolName)
	suite.Require().Len(rules, len(testRewardPerBlock))

	for _, r := range rules {
		suite.Require().Equal(testTotalReward.AmountOf(r.Reward), r.RemainingReward)
		suite.Require().Equal(testTotalReward.AmountOf(r.Reward), r.TotalReward)
		suite.Require().Equal(testRewardPerBlock.AmountOf(r.Reward), r.RewardPerBlock)
		suite.Require().Equal(sdk.ZeroDec(), r.RewardPerShare)
	}

	pool.Rules = rules
	suite.Require().Equal(pool.ExpiredHeight(), pool.EndHeight)

	//check queue
	suite.keeper.IteratorExpiredPool(ctx, pool.EndHeight, func(pool *types.FarmPool) {
		suite.Require().Equal(testPoolName, pool.Name)
	})

	//check balance
	expectedBal := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, testInitCoinAmt)).
		Sub(sdk.NewCoins(suite.keeper.CreatePoolFee(ctx))).
		Sub(testTotalReward)
	actualBal := suite.app.BankKeeper.GetAllBalances(ctx, testCreator)
	suite.Require().Equal(expectedBal, actualBal)
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
	suite.Require().NoError(err)

	newCtx := suite.app.BaseApp.NewContext(isCheckTx, tmproto.Header{
		Height: 10,
	})
	err = suite.keeper.DestroyPool(newCtx, testPoolName, testCreator)
	suite.Require().NoError(err)

	//check farm pool
	pool, exist := suite.keeper.GetPool(newCtx, testPoolName)
	suite.Require().True(exist)

	suite.Require().EqualValues(newCtx.BlockHeight(), pool.LastHeightDistrRewards)
	suite.Require().EqualValues(newCtx.BlockHeight(), pool.EndHeight)

	//check balance
	expectedBal := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, testInitCoinAmt)).
		Sub(sdk.NewCoins(suite.keeper.CreatePoolFee(ctx)))
	actualBal := suite.app.BankKeeper.GetAllBalances(ctx, testCreator)
	suite.Require().Equal(expectedBal, actualBal)
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
	suite.Require().NoError(err)

	//check farm pool
	pool, exist := suite.keeper.GetPool(ctx, testPoolName)
	suite.Require().True(exist)

	//panic with adding new token as reward
	rewardAdded := sdk.NewCoins(
		sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10_000_000)),
		sdk.NewCoin("uiris", sdk.NewInt(10_000_000)),
	)
	_, err = suite.keeper.AppendReward(ctx,
		testPoolName,
		rewardAdded,
		testCreator,
	)
	suite.Require().Error(err)

	rewardAdded = sdk.NewCoins(
		sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10_000_000)),
	)
	remaining, err := suite.keeper.AppendReward(ctx,
		testPoolName,
		rewardAdded,
		testCreator,
	)
	suite.Require().NoError(err)
	suite.Require().Equal(testTotalReward.Add(rewardAdded...), remaining)

	//check farm pool
	pool2, exist := suite.keeper.GetPool(ctx, testPoolName)
	suite.Require().True(exist)
	suite.Require().EqualValues(pool.EndHeight+10, pool2.EndHeight)

	//check reward rules
	rules := suite.keeper.GetRewardRules(ctx, testPoolName)
	suite.Require().Len(rules, len(testRewardPerBlock))

	for _, r := range rules {
		suite.Equal(
			testTotalReward.AmountOf(r.Reward).Add(rewardAdded.AmountOf(r.Reward)),
			r.RemainingReward,
		)
		suite.Equal(
			testTotalReward.AmountOf(r.Reward).Add(rewardAdded.AmountOf(r.Reward)),
			r.TotalReward,
		)
	}
}

func (suite *KeeperTestSuite) TestStake() {
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
	suite.Require().NoError(err)

	//stake first
	newCtx := suite.app.BaseApp.NewContext(isCheckTx, tmproto.Header{
		Height: 100,
	})
	lpToken := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))
	reward, err := suite.keeper.Stake(newCtx, testPoolName, lpToken, testFarmer1)

	suite.Require().NoError(err)
	suite.Require().Nil(reward)

	//check farm pool
	pool, exist := suite.keeper.GetPool(newCtx, testPoolName)
	suite.Require().True(exist)
	suite.Require().Equal(lpToken, pool.TotalLpTokenLocked)
	suite.Require().EqualValues(newCtx.BlockHeight(), pool.LastHeightDistrRewards)

	//check farm infomation
	info, exist := suite.keeper.GetFarmInfo(newCtx, testPoolName, testFarmer1.String())
	suite.Require().True(exist)
	suite.Require().Equal(lpToken.Amount, info.Locked)

	//check reward rules
	rules := suite.keeper.GetRewardRules(newCtx, testPoolName)
	suite.Require().Len(rules, len(testRewardPerBlock))
	for _, r := range rules {
		suite.Require().Equal(testTotalReward.AmountOf(r.Reward), r.RemainingReward)
		suite.Require().Equal(testTotalReward.AmountOf(r.Reward), r.TotalReward)
		suite.Require().Equal(testRewardPerBlock.AmountOf(r.Reward), r.RewardPerBlock)
		suite.Require().Equal(sdk.ZeroDec(), r.RewardPerShare)
	}

	//stake again
	newCtx1 := suite.app.BaseApp.NewContext(isCheckTx, tmproto.Header{
		Height: 200,
	})
	lpToken = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))
	reward, err = suite.keeper.Stake(newCtx1, testPoolName, lpToken, testFarmer1)

	suite.Require().NoError(err)
	suite.Require().Equal("100000000stake", reward.String())

	info, _ = suite.keeper.GetFarmInfo(newCtx, testPoolName, testFarmer1.String())
	suite.Require().Equal("200000000stake", info.RewardDebt.String())

	//stake again
	newCtx2 := suite.app.BaseApp.NewContext(isCheckTx, tmproto.Header{
		Height: 300,
	})
	lpToken = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))
	reward, err = suite.keeper.Stake(newCtx2, testPoolName, lpToken, testFarmer1)

	suite.Require().NoError(err)
	suite.Require().Equal("100000000stake", reward.String())

	info, _ = suite.keeper.GetFarmInfo(newCtx, testPoolName, testFarmer1.String())
	suite.Require().Equal("450000000stake", info.RewardDebt.String())

	//check reward rules again
	rules = suite.keeper.GetRewardRules(newCtx, testPoolName)
	suite.Require().Len(rules, len(testRewardPerBlock))
	for _, r := range rules {
		suite.Require().Equal(testTotalReward.AmountOf(r.Reward).SubRaw(200000000), r.RemainingReward)
		suite.Require().Equal(sdk.NewDecFromIntWithPrec(sdk.NewInt(15), 1), r.RewardPerShare)
	}

}

func (suite *KeeperTestSuite) TestUnstake() {

}

func (suite *KeeperTestSuite) TestHarvest() {

}
