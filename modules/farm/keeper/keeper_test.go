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

	lpToken := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))
	type args struct {
		index          int64
		height         int64
		stakeCoin      sdk.Coin
		expectReward   sdk.Coins
		debt           sdk.Coins
		rewardPerShare sdk.Dec
	}

	var testcase = []args{
		{
			index:          1,
			height:         100,
			stakeCoin:      lpToken,
			expectReward:   nil,
			debt:           nil,
			rewardPerShare: sdk.ZeroDec(),
		},
		{
			index:          2,
			height:         200,
			stakeCoin:      lpToken,
			expectReward:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))),
			debt:           sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(200_000_000))),
			rewardPerShare: sdk.NewDecFromIntWithPrec(sdk.NewInt(1), 0),
		},
		{
			index:          3,
			height:         300,
			stakeCoin:      lpToken,
			expectReward:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))),
			debt:           sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(450_000_000))),
			rewardPerShare: sdk.NewDecFromIntWithPrec(sdk.NewInt(15), 1),
		},
		{
			index:          4,
			height:         400,
			stakeCoin:      lpToken,
			expectReward:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(99999999))),
			debt:           sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(733_333_333))),
			rewardPerShare: sdk.NewDecFromIntWithPrec(sdk.NewInt(1_833_333_333_333_333_333), 18),
		},
	}

	for _, tc := range testcase {
		suite.AssertStake(tc.index,
			tc.height,
			tc.stakeCoin,
			tc.expectReward,
			tc.debt,
			tc.rewardPerShare)
	}
}

func (suite *KeeperTestSuite) TestUnstake() {
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

	lpToken := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))
	suite.AssertStake(1, 100, lpToken, nil, nil, sdk.ZeroDec())

	newCtx := suite.app.BaseApp.NewContext(isCheckTx, tmproto.Header{
		Height: 200,
	})
	reward, err := suite.keeper.Unstake(newCtx, testPoolName, lpToken, testFarmer1)
	suite.Require().NoError(err)

	expect := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000)))
	suite.Require().Equal(expect, reward)

	//check farm information
	_, exist := suite.keeper.GetFarmInfo(ctx, testPoolName, testFarmer1.String())
	suite.Require().False(exist)

	//check farm pool
	pool, exist := suite.keeper.GetPool(ctx, testPoolName)
	suite.Require().True(exist)
	suite.Require().Equal(pool.TotalLpTokenLocked.Amount, sdk.ZeroInt())
}

func (suite *KeeperTestSuite) TestHarvest() {
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

	lpToken := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))
	suite.AssertStake(1, 100, lpToken, nil, nil, sdk.ZeroDec())

	type args struct {
		index          int64
		height         int64
		expectReward   sdk.Coins
		debt           sdk.Coins
		rewardPerShare sdk.Dec
	}

	var testcase = []args{
		{
			index:          1,
			height:         200,
			expectReward:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))),
			debt:           sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))),
			rewardPerShare: sdk.NewDecFromIntWithPrec(sdk.NewInt(1), 0),
		},
		{
			index:          2,
			height:         300,
			expectReward:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))),
			debt:           sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(200_000_000))),
			rewardPerShare: sdk.NewDecFromIntWithPrec(sdk.NewInt(2), 0),
		},
		{
			index:          3,
			height:         400,
			expectReward:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))),
			debt:           sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(300_000_000))),
			rewardPerShare: sdk.NewDecFromIntWithPrec(sdk.NewInt(3), 0),
		},
		{
			index:          4,
			height:         500,
			expectReward:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000))),
			debt:           sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(400_000_000))),
			rewardPerShare: sdk.NewDecFromIntWithPrec(sdk.NewInt(4), 0),
		},
	}

	for _, tc := range testcase {
		suite.AssertHarvest(tc.index,
			tc.height,
			tc.expectReward,
			tc.debt,
			tc.rewardPerShare)
	}
}

func (suite *KeeperTestSuite) AssertStake(index, height int64,
	stakeCoin sdk.Coin,
	expectReward, debt sdk.Coins,
	RewardPerShare sdk.Dec,
) {
	ctx := suite.app.BaseApp.NewContext(isCheckTx, tmproto.Header{
		Height: height,
	})
	reward, err := suite.keeper.Stake(ctx, testPoolName, stakeCoin, testFarmer1)

	suite.Require().NoError(err)
	suite.Require().Equal(expectReward, reward)

	info, exist := suite.keeper.GetFarmInfo(ctx, testPoolName, testFarmer1.String())
	suite.Require().True(exist)
	suite.Require().Equal(debt, info.RewardDebt)
	suite.Require().Equal(stakeCoin.Amount.MulRaw(index), info.Locked)

	//check reward rules again
	rules := suite.keeper.GetRewardRules(ctx, testPoolName)
	suite.Require().Len(rules, len(testRewardPerBlock))
	for _, r := range rules {
		suite.Require().Equal(RewardPerShare, r.RewardPerShare)
	}
}

func (suite *KeeperTestSuite) AssertHarvest(index, height int64,
	expectReward, debt sdk.Coins,
	RewardPerShare sdk.Dec,
) {
	ctx := suite.app.BaseApp.NewContext(isCheckTx, tmproto.Header{
		Height: height,
	})
	reward, err := suite.keeper.Harvest(ctx, testPoolName, testFarmer1)

	suite.Require().NoError(err)
	suite.Require().Equal(expectReward, reward)

	info, exist := suite.keeper.GetFarmInfo(ctx, testPoolName, testFarmer1.String())
	suite.Require().True(exist)
	suite.Require().Equal(debt, info.RewardDebt)

	//check reward rules again
	rules := suite.keeper.GetRewardRules(ctx, testPoolName)
	suite.Require().Len(rules, len(testRewardPerBlock))
	for _, r := range rules {
		suite.Require().Equal(RewardPerShare, r.RewardPerShare)
	}
}
