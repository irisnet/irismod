package farm

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	"mods.irisnet.org/e2e"
	coinswaptypes "mods.irisnet.org/modules/coinswap/types"
	farmcli "mods.irisnet.org/modules/farm/client/cli"
	farmtypes "mods.irisnet.org/modules/farm/types"
	tokentypes "mods.irisnet.org/modules/token/types/v1"
)

// QueryTestSuite is a suite of end-to-end tests for the farm module
type QueryTestSuite struct {
	e2e.TestSuite
}

// SetupSuite creates a new network for integration tests
func (s *QueryTestSuite) SetupSuite() {
	sdk.SetCoinDenomRegex(func() string {
		return `[a-zA-Z][a-zA-Z0-9/\-]{2,127}`
	})
	s.TestSuite.SetupSuite()
}

// TestQueryCmd tests all query command in the farm module
func (s *QueryTestSuite) TestQueryCmd() {
	val := s.Validators[0]
	clientCtx := val.ClientCtx
	baseURL := val.APIAddress

	s.setup()

	// ---------------------------------------------------------------------------

	creator := val.Address
	description := "iris-atom farm pool"
	startHeight := s.latestHeight() + 2
	rewardPerBlock := sdk.NewCoins(sdk.NewCoin(s.BondDenom, math.NewInt(10)))
	lpTokenDenom := "lpt-1"
	totalReward := sdk.NewCoins(sdk.NewCoin(s.BondDenom, math.NewInt(1000)))
	editable := true

	globalFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.BondDenom, math.NewInt(10))).String(),
		),
	}

	args := []string{
		fmt.Sprintf("--%s=%s", farmcli.FlagDescription, description),
		fmt.Sprintf("--%s=%d", farmcli.FlagStartHeight, startHeight),
		fmt.Sprintf("--%s=%s", farmcli.FlagRewardPerBlock, rewardPerBlock),
		fmt.Sprintf("--%s=%s", farmcli.FlagLPTokenDenom, lpTokenDenom),
		fmt.Sprintf("--%s=%s", farmcli.FlagTotalReward, totalReward),
		fmt.Sprintf("--%s=%v", farmcli.FlagEditable, editable),
	}

	args = append(args, globalFlags...)
	txResult := CreateFarmPoolExec(
		s.T(),
		s.Network,
		clientCtx,
		creator.String(),
		args...,
	)
	s.Require().EqualValues(txResult.Code, 0, txResult.Log)

	poolID := s.GetAttribute(
		farmtypes.EventTypeCreatePool,
		farmtypes.AttributeValuePoolId,
		txResult.Events,
	)
	expectedContents := farmtypes.FarmPoolEntry{
		Id:              poolID,
		Description:     description,
		Creator:         creator.String(),
		StartHeight:     startHeight,
		EndHeight:       startHeight + 100,
		Editable:        editable,
		Expired:         false,
		TotalLptLocked:  sdk.NewCoin(lpTokenDenom, math.ZeroInt()),
		TotalReward:     totalReward,
		RemainingReward: totalReward,
		RewardPerBlock:  rewardPerBlock,
	}

	respType := proto.Message(&farmtypes.QueryFarmPoolsResponse{})
	queryPoolURL := fmt.Sprintf("%s/irismod/farm/pools", baseURL)
	resp, err := testutil.GetRequest(queryPoolURL)

	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, respType))
	result := respType.(*farmtypes.QueryFarmPoolsResponse)
	s.Require().EqualValues(expectedContents, *result.Pools[0])

	_, err = s.WaitForHeight(startHeight)
	s.Require().NoError(err)
	s.Require().NoError(s.WaitForNextBlock())

	lpToken := sdk.NewCoin(lpTokenDenom, math.NewInt(100))
	txResult = StakeExec(
		s.T(),
		s.Network,
		clientCtx,
		creator.String(),
		poolID,
		lpToken.String(),
		globalFlags...,
	)
	s.Require().Equal(uint32(0), txResult.Code, txResult.Log)

	expectFarmer := farmtypes.LockedInfo{
		PoolId:        poolID,
		Locked:        lpToken,
		PendingReward: sdk.Coins{},
	}

	queryFarmerRespType := proto.Message(&farmtypes.QueryFarmerResponse{})
	queryFarmInfoURL := fmt.Sprintf("%s/irismod/farm/farmers/%s", baseURL, creator.String())
	resp, err = testutil.GetRequest(queryFarmInfoURL)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryFarmerRespType))
	farmer := queryFarmerRespType.(*farmtypes.QueryFarmerResponse)

	if farmer.Height-txResult.Height > 0 {
		expectFarmer.PendingReward = rewardPerBlock.MulInt(
			math.NewInt(farmer.Height - txResult.Height),
		)
	}
	s.Require().EqualValues(expectFarmer, *farmer.List[0])
}

func (s *QueryTestSuite) latestHeight() int64 {
	height, err := s.LatestHeight()
	s.Require().NoError(err)
	return height
}

func (s *QueryTestSuite) setup() {
	val := s.Validators[0]
	clientCtx := val.ClientCtx

	from := val.Address
	const symbol = "kitty"
	const name = "Kitty Token"
	const minUnit = "kitty"
	const scale = uint32(0)
	const initialSupply = uint64(100000000)
	const maxSupply = uint64(200000000)
	const mintable = true

	// issue token
	msgIssueToken := &tokentypes.MsgIssueToken{
		Symbol:        symbol,
		Name:          name,
		Scale:         scale,
		MinUnit:       minUnit,
		InitialSupply: initialSupply,
		MaxSupply:     maxSupply,
		Mintable:      mintable,
		Owner:         from.String(),
	}
	res := s.BlockSendMsgs(s.T(), msgIssueToken)
	s.Require().Equal(uint32(0), res.Code, res.Log)

	// add liquidity
	status, err := clientCtx.Client.Status(context.Background())
	s.Require().NoError(err)
	deadline := status.SyncInfo.LatestBlockTime.Add(time.Minute)

	msgAddLiquidity := &coinswaptypes.MsgAddLiquidity{
		MaxToken:         sdk.NewCoin(symbol, math.NewInt(1000)),
		ExactStandardAmt: math.NewInt(1000),
		MinLiquidity:     math.NewInt(1000),
		Deadline:         deadline.Unix(),
		Sender:           val.Address.String(),
	}
	res = s.BlockSendMsgs(s.T(), msgAddLiquidity)
	s.Require().Equal(uint32(0), res.Code, res.Log)
}
