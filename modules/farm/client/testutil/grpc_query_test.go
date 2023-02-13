package testutil_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/testutil/rest"
	sdk "github.com/cosmos/cosmos-sdk/types"

	farmcli "github.com/irisnet/irismod/modules/farm/client/cli"
	"github.com/irisnet/irismod/modules/farm/client/testutil"
	farmtypes "github.com/irisnet/irismod/modules/farm/types"
	"github.com/irisnet/irismod/simapp"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := simapp.NewConfig()
	cfg.NumValidators = 1

	s.cfg = cfg
	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestRest() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	baseURL := val.APIAddress

	// ---------------------------------------------------------------------------

	creator := val.Address
	description := "iris-atom farm pool"
	startHeight := s.LatestHeight() + 1
	rewardPerBlock := sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))
	lpTokenDenom := s.cfg.BondDenom
	totalReward := sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(1000)))
	editable := true

	globalFlags := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
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
	respType := proto.Message(&sdk.TxResponse{})
	expectedCode := uint32(0)

	bz, err := testutil.CreateFarmPoolExec(clientCtx,
		creator.String(),
		args...,
	)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp := respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	poolId := gjson.Get(txResp.RawLog, "0.events.3.attributes.1.value").String()
	expectedContents := farmtypes.FarmPoolEntry{
		Id:              poolId,
		Description:     description,
		Creator:         creator.String(),
		StartHeight:     startHeight,
		EndHeight:       startHeight + 100,
		Editable:        editable,
		Expired:         false,
		TotalLptLocked:  sdk.NewCoin(lpTokenDenom, sdk.ZeroInt()),
		TotalReward:     totalReward,
		RemainingReward: totalReward,
		RewardPerBlock:  rewardPerBlock,
	}

	respType = proto.Message(&farmtypes.QueryFarmPoolsResponse{})
	queryPoolURL := fmt.Sprintf("%s/irismod/farm/pools", baseURL)
	resp, err := rest.GetRequest(queryPoolURL)

	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, respType))
	result := respType.(*farmtypes.QueryFarmPoolsResponse)
	s.Require().EqualValues(expectedContents, *result.Pools[0])

	_, err = s.network.WaitForHeight(startHeight)
	s.Require().NoError(err)

	respType = proto.Message(&sdk.TxResponse{})
	lpToken := sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(100))
	bz, err = testutil.StakeExec(clientCtx,
		creator.String(),
		poolId,
		lpToken.String(),
		globalFlags...,
	)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	expectFarmer := farmtypes.LockedInfo{
		PoolId:        poolId,
		Locked:        lpToken,
		PendingReward: sdk.Coins{},
	}

	queryFarmerRespType := proto.Message(&farmtypes.QueryFarmerResponse{})
	queryFarmInfoURL := fmt.Sprintf("%s/irismod/farm/farmers/%s", baseURL, creator.String())
	resp, err = rest.GetRequest(queryFarmInfoURL)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryFarmerRespType))
	farmer := queryFarmerRespType.(*farmtypes.QueryFarmerResponse)
	s.Require().EqualValues(expectFarmer, *farmer.List[0])
}

func (s *IntegrationTestSuite) LatestHeight() int64 {
	height, err := s.network.LatestHeight()
	s.Require().NoError(err)
	return height
}
