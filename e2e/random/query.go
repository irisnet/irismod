package random

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"mods.irisnet.org/e2e"
	"mods.irisnet.org/e2e/service"
	randomcli "mods.irisnet.org/modules/random/client/cli"
	randomtypes "mods.irisnet.org/modules/random/types"
	servicecli "mods.irisnet.org/modules/service/client/cli"
	servicetypes "mods.irisnet.org/modules/service/types"
)

// QueryTestSuite is a suite of end-to-end tests for the nft module
type QueryTestSuite struct {
	e2e.TestSuite
}

// SetupSuite sets up test suite
func (s *QueryTestSuite) SetupSuite() {
	s.SetupSuiteWithModifyConfigFn(func(cfg *network.Config) {
		var serviceGenState servicetypes.GenesisState
		cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[servicetypes.ModuleName], &serviceGenState)

		serviceGenState.Definitions = append(
			serviceGenState.Definitions,
			servicetypes.GenOraclePriceSvcDefinition(),
			servicetypes.GetRandomSvcDefinition(),
		)
		serviceGenState.Bindings = append(
			serviceGenState.Bindings,
			servicetypes.GenOraclePriceSvcBinding(sdk.DefaultBondDenom),
		)
		cfg.GenesisState[servicetypes.ModuleName] = cfg.Codec.MustMarshalJSON(&serviceGenState)
	})
}

// TestQueryCmd tests all query command in the nft module
func (s *QueryTestSuite) TestQueryCmd() {
	val := s.Validators[0]
	clientCtx := val.ClientCtx
	expectedCode := uint32(0)

	// ---------------------------------------------------------------------------
	serviceDeposit := fmt.Sprintf("50000%s", s.BondDenom)
	servicePrices := fmt.Sprintf(`{"price": "50%s"}`, s.BondDenom)
	qos := int64(3)
	options := "{}"
	provider := val.Address
	baseURL := val.APIAddress

	from := val.Address
	blockInterval := 4
	oracle := true
	serviceFeeCap := fmt.Sprintf("50%s", s.BondDenom)

	respResult := `{"code":200,"message":""}`
	seedStr := "ABCDEF12ABCDEF12ABCDEF12ABCDEF12ABCDEF12ABCDEF12ABCDEF12ABCDEF12"
	respOutput := fmt.Sprintf(`{"header":{},"body":{"seed":"%s"}}`, seedStr)

	// ------bind random service-------------
	args := []string{
		fmt.Sprintf("--%s=%s", servicecli.FlagServiceName, randomtypes.ServiceName),
		fmt.Sprintf("--%s=%s", servicecli.FlagDeposit, serviceDeposit),
		fmt.Sprintf("--%s=%s", servicecli.FlagPricing, servicePrices),
		fmt.Sprintf("--%s=%d", servicecli.FlagQoS, qos),
		fmt.Sprintf("--%s=%s", servicecli.FlagOptions, options),
		fmt.Sprintf("--%s=%s", servicecli.FlagProvider, provider),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.Network.BondDenom, math.NewInt(10))).String(),
		),
	}

	txResult := service.BindServiceExec(
		s.T(),
		s.Network,
		clientCtx,
		provider.String(),
		args...)
	s.Require().Equal(expectedCode, txResult.Code)

	// ------test GetCmdRequestRandom()-------------
	args = []string{
		fmt.Sprintf("--%s=%s", randomcli.FlagServiceFeeCap, serviceFeeCap),
		fmt.Sprintf("--%s=%t", randomcli.FlagOracle, oracle),
		fmt.Sprintf("--%s=%d", randomcli.FlagBlockInterval, blockInterval),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.Network.BondDenom, math.NewInt(10))).String(),
		),
	}

	txResult = RequestRandomExec(s.T(), s.Network, clientCtx, from.String(), args...)
	s.Require().Equal(expectedCode, txResult.Code)

	requestID := txResult.Events[8].Attributes[0].Value
	heightStr := txResult.Events[8].Attributes[2].Value
	requestHeight, err := strconv.ParseInt(heightStr, 10, 64)
	s.Require().NoError(err)

	// ------test GetCmdQueryRandomRequestQueue()-------------
	url := fmt.Sprintf("%s/irismod/random/queue", baseURL)
	resp, err := testutil.GetRequest(url)
	respType := proto.Message(&randomtypes.QueryRandomRequestQueueResponse{})
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, respType))
	qrrResp := respType.(*randomtypes.QueryRandomRequestQueueResponse)
	s.Require().NoError(err)
	s.Require().Len(qrrResp.Requests, 1)

	// ------get service request-------------
	requestHeight++
	_, err = s.Network.WaitForHeightWithTimeout(
		requestHeight,
		time.Duration(int64(blockInterval+2)*int64(s.Network.TimeoutCommit)),
	)
	if err != nil {
		s.Require().NoError(s.Network.WaitForNBlock(2))
	}

	blockResult, err := val.RPCClient.BlockResults(context.Background(), &requestHeight)
	s.Require().NoError(err)
	var requestId string
	for _, event := range blockResult.FinalizeBlockEvents {
		if event.Type == servicetypes.EventTypeNewBatchRequestProvider {
			var found bool
			var requestIds []string
			var requestsBz []byte
			for _, attribute := range event.Attributes {
				if attribute.Key == servicetypes.AttributeKeyRequests {
					requestsBz = []byte(attribute.Value)
					found = true
				}
			}
			s.Require().True(found)
			if found {
				err := json.Unmarshal(requestsBz, &requestIds)
				s.Require().NoError(err)
			}
			s.Require().Len(requestIds, 1)
			requestId = requestIds[0]
		}
	}
	s.Require().NotEmpty(requestId)

	// ------respond service request-------------
	args = []string{
		fmt.Sprintf("--%s=%s", servicecli.FlagRequestID, requestId),
		fmt.Sprintf("--%s=%s", servicecli.FlagResult, respResult),
		fmt.Sprintf("--%s=%s", servicecli.FlagData, respOutput),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf(
			"--%s=%s",
			flags.FlagFees,
			sdk.NewCoins(sdk.NewCoin(s.Network.BondDenom, math.NewInt(10))).String(),
		),
	}

	txResult = service.RespondServiceExec(
		s.T(),
		s.Network,
		clientCtx,
		provider.String(),
		args...)
	s.Require().Equal(expectedCode, txResult.Code)

	// ------test GetCmdQueryRandom()-------------
	url = fmt.Sprintf("%s/irismod/random/randoms/%s", baseURL, requestID)
	resp, err = testutil.GetRequest(url)
	respType = proto.Message(&randomtypes.QueryRandomResponse{})
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, respType))
	randomResp := respType.(*randomtypes.QueryRandomResponse)
	s.Require().NoError(err)
	s.Require().NotNil(randomResp.Random.Value)
}
